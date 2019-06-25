package apiserver

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/samsung-cnct/cma-ssh/internal/apiserver"
	pb "github.com/samsung-cnct/cma-ssh/pkg/generated/api"
	"github.com/samsung-cnct/cma-ssh/pkg/ui/website"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type ServerOptions struct {
	PortNumber int
}

type ApiServer struct {
	Manager manager.Manager
	TcpMux  cmux.CMux
}

type MuxApiServer interface {
	AddServersToMux(options *ServerOptions)
	GetMux() cmux.CMux
}

func NewApiServer(manager manager.Manager, tcpMux cmux.CMux) MuxApiServer {
	return &ApiServer{Manager: manager, TcpMux: tcpMux}
}

func (r *ApiServer) AddServersToMux(options *ServerOptions) {
	r.addGRPCServer(r.TcpMux)
	r.addRestAndWebsite(r.TcpMux, options.PortNumber)
}

func (r *ApiServer) GetMux() cmux.CMux {
	return r.TcpMux
}

func (r *ApiServer) addGRPCServer(tcpMux cmux.CMux) {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterClusterServer(grpcServer, r.newgRPCServiceServer())
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	grpcListener := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	// Start servers
	go func() {
		klog.Info("Starting gRPC Server")
		if err := grpcServer.Serve(grpcListener); err != nil {
			klog.Errorf("Unable to start external gRPC server: %q", err)
		}
	}()
}

func (r *ApiServer) addRestAndWebsite(tcpMux cmux.CMux, grpcPortNumber int) {
	httpListener := tcpMux.Match(cmux.HTTP1Fast())

	go func() {
		router := http.NewServeMux()
		website.AddWebsiteHandles(router)
		r.addgRPCRestGateway(router, grpcPortNumber)
		httpServer := http.Server{
			Handler: router,
		}
		klog.Info("Starting HTTP/1 Server")
		err := httpServer.Serve(httpListener)
		if err != nil {
			klog.Errorf("Failed to start http server Serve(): %q", err)
		}
	}()

}

func (r *ApiServer) addgRPCRestGateway(router *http.ServeMux, grpcPortNumber int) {
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	err := pb.RegisterClusterHandlerFromEndpoint(context.Background(), gwmux, "localhost:"+strconv.Itoa(grpcPortNumber), dopts)
	if err != nil {
		klog.Errorf("Failed to register handler from endpoint: %q", err)
	}
	router.Handle("/api/", allowCORS(gwmux))
}

func (r *ApiServer) newgRPCServiceServer() *apiserver.Server {
	return &apiserver.Server{Manager: r.Manager}
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// pre-flightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}

	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	klog.Infof("pre-flight request for %s", r.URL.Path)
}
