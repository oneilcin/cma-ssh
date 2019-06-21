package util

import (
	"context"
	"errors"

	clusterv1alpha1 "github.com/samsung-cnct/cma-ssh/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClusterFromNamespace assumes that there would only be one cluster per namespace
func GetClusterFromNamespace(c client.Client, namespace string) (*clusterv1alpha1.CnctCluster, error) {
	var clusterlist clusterv1alpha1.CnctClusterList
	err := c.List(context.Background(), &client.ListOptions{Namespace: namespace}, &clusterlist)
	if err != nil {
		return nil, err
	}
	if len(clusterlist.Items) == 0 {
		return nil, errors.New("No cluster Items found")
	}
	if len(clusterlist.Items) > 1 {
		return nil, errors.New("Found more than one cluster in namespace")
	}
	// Note: The item contains the Name, and possibly no other fields.
	cluster := clusterlist.Items[0]
	return &cluster, nil
}
