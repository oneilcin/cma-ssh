# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [AddNodePoolMsg](#cnct.kaas.api.AddNodePoolMsg)
    - [AddNodePoolReply](#cnct.kaas.api.AddNodePoolReply)
    - [ClusterDetailItem](#cnct.kaas.api.ClusterDetailItem)
    - [ClusterItem](#cnct.kaas.api.ClusterItem)
    - [ControlPlaneMachineSpec](#cnct.kaas.api.ControlPlaneMachineSpec)
    - [CreateClusterMsg](#cnct.kaas.api.CreateClusterMsg)
    - [CreateClusterReply](#cnct.kaas.api.CreateClusterReply)
    - [DeleteClusterMsg](#cnct.kaas.api.DeleteClusterMsg)
    - [DeleteClusterReply](#cnct.kaas.api.DeleteClusterReply)
    - [DeleteNodePoolMsg](#cnct.kaas.api.DeleteNodePoolMsg)
    - [DeleteNodePoolReply](#cnct.kaas.api.DeleteNodePoolReply)
    - [GetClusterListMsg](#cnct.kaas.api.GetClusterListMsg)
    - [GetClusterListReply](#cnct.kaas.api.GetClusterListReply)
    - [GetClusterMsg](#cnct.kaas.api.GetClusterMsg)
    - [GetClusterNodesStatusMsg](#cnct.kaas.api.GetClusterNodesStatusMsg)
    - [GetClusterNodesStatusReply](#cnct.kaas.api.GetClusterNodesStatusReply)
    - [GetClusterNodesStatusReply.MachineStatus](#cnct.kaas.api.GetClusterNodesStatusReply.MachineStatus)
    - [GetClusterReply](#cnct.kaas.api.GetClusterReply)
    - [GetUpgradeClusterInformationMsg](#cnct.kaas.api.GetUpgradeClusterInformationMsg)
    - [GetUpgradeClusterInformationReply](#cnct.kaas.api.GetUpgradeClusterInformationReply)
    - [GetVersionMsg](#cnct.kaas.api.GetVersionMsg)
    - [GetVersionReply](#cnct.kaas.api.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cnct.kaas.api.GetVersionReply.VersionInformation)
    - [KubernetesLabel](#cnct.kaas.api.KubernetesLabel)
    - [MachineSpec](#cnct.kaas.api.MachineSpec)
    - [ScaleNodePoolMsg](#cnct.kaas.api.ScaleNodePoolMsg)
    - [ScaleNodePoolReply](#cnct.kaas.api.ScaleNodePoolReply)
    - [ScaleNodePoolSpec](#cnct.kaas.api.ScaleNodePoolSpec)
    - [UpgradeClusterMsg](#cnct.kaas.api.UpgradeClusterMsg)
    - [UpgradeClusterReply](#cnct.kaas.api.UpgradeClusterReply)
  
    - [ClusterStatus](#cnct.kaas.api.ClusterStatus)
  
  
    - [Cluster](#cnct.kaas.api.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cnct.kaas.api.AddNodePoolMsg"></a>

### AddNodePoolMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterName | [string](#string) |  | What is the cluster to add node pools to |
| worker_node_pools | [MachineSpec](#cnct.kaas.api.MachineSpec) | repeated | What Machines to add to the cluster |






<a name="cnct.kaas.api.AddNodePoolReply"></a>

### AddNodePoolReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the node pool was provisioned by this request |






<a name="cnct.kaas.api.ClusterDetailItem"></a>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cnct.kaas.api.ClusterStatus) |  | The status of the cluster |






<a name="cnct.kaas.api.ClusterItem"></a>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cnct.kaas.api.ClusterStatus) |  | The status of the cluster |






<a name="cnct.kaas.api.ControlPlaneMachineSpec"></a>

### ControlPlaneMachineSpec
The specification for a set of control plane machines


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| labels | [KubernetesLabel](#cnct.kaas.api.KubernetesLabel) | repeated | The labels for the control plane machines |
| instanceType | [string](#string) |  | Type of machines to provision (standard or gpu) |
| count | [int32](#int32) |  | The number of machines |






<a name="cnct.kaas.api.CreateClusterMsg"></a>

### CreateClusterMsg
CreateClusterMsg


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| k8s_version | [string](#string) |  | The version of Kubernetes for worker nodes. Control plane versions are determined by the MachineSpec. |
| control_plane_nodes | [ControlPlaneMachineSpec](#cnct.kaas.api.ControlPlaneMachineSpec) |  | Machines which comprise the cluster control plane |
| worker_node_pools | [MachineSpec](#cnct.kaas.api.MachineSpec) | repeated | Machines which comprise the cluster |






<a name="cnct.kaas.api.CreateClusterReply"></a>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cnct.kaas.api.ClusterItem) |  | The details of the cluster request response |






<a name="cnct.kaas.api.DeleteClusterMsg"></a>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |






<a name="cnct.kaas.api.DeleteClusterReply"></a>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cnct.kaas.api.DeleteNodePoolMsg"></a>

### DeleteNodePoolMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterName | [string](#string) |  | What is the cluster to delete node pools |
| node_pool_names | [string](#string) | repeated | What is the node pool names to delete |






<a name="cnct.kaas.api.DeleteNodePoolReply"></a>

### DeleteNodePoolReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cnct.kaas.api.GetClusterListMsg"></a>

### GetClusterListMsg







<a name="cnct.kaas.api.GetClusterListReply"></a>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cnct.kaas.api.ClusterItem) | repeated | List of clusters |






<a name="cnct.kaas.api.GetClusterMsg"></a>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |






<a name="cnct.kaas.api.GetClusterNodesStatusMsg"></a>

### GetClusterNodesStatusMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterName | [string](#string) |  | What is the name of the Cluster to query |






<a name="cnct.kaas.api.GetClusterNodesStatusReply"></a>

### GetClusterNodesStatusReply
GetClusterNodesStatusReply is the response to GetNodePool
or cluster machines status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  | The name of the cluster |
| count | [int32](#int32) |  | Count of machines in cluster |
| machines | [GetClusterNodesStatusReply.MachineStatus](#cnct.kaas.api.GetClusterNodesStatusReply.MachineStatus) | repeated | Gets list of nodes in a cluster |






<a name="cnct.kaas.api.GetClusterNodesStatusReply.MachineStatus"></a>

### GetClusterNodesStatusReply.MachineStatus
The status of a machine


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| k8sNodeStatus | [string](#string) |  | Node k8s status |
| k8sVersion | [string](#string) |  | Kubernetes Version |
| maasSystemId | [string](#string) |  | MaaS Node system_id |
| maasHostname | [string](#string) |  | MaaS Node hostname |
| maasNodeStatus | [string](#string) |  | MaaS node status |
| maasIPAddr | [string](#string) |  | MaaS IP Address |






<a name="cnct.kaas.api.GetClusterReply"></a>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cnct.kaas.api.ClusterDetailItem) |  |  |






<a name="cnct.kaas.api.GetUpgradeClusterInformationMsg"></a>

### GetUpgradeClusterInformationMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |






<a name="cnct.kaas.api.GetUpgradeClusterInformationReply"></a>

### GetUpgradeClusterInformationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Can the cluster be upgraded |
| versions | [string](#string) | repeated | What versions are possible right now |






<a name="cnct.kaas.api.GetVersionMsg"></a>

### GetVersionMsg
Get version of API Server






<a name="cnct.kaas.api.GetVersionReply"></a>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cnct.kaas.api.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cnct.kaas.api.GetVersionReply.VersionInformation"></a>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |






<a name="cnct.kaas.api.KubernetesLabel"></a>

### KubernetesLabel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of a label |
| value | [string](#string) |  | The value of a label |






<a name="cnct.kaas.api.MachineSpec"></a>

### MachineSpec
The specification for a set of machines


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the machine set |
| labels | [KubernetesLabel](#cnct.kaas.api.KubernetesLabel) | repeated | The labels for the machine set |
| instanceType | [string](#string) |  | Type of machines to provision (standard or gpu) |
| count | [int32](#int32) |  | The number of machines |






<a name="cnct.kaas.api.ScaleNodePoolMsg"></a>

### ScaleNodePoolMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterName | [string](#string) |  | What is the name of the cluster to scale a node pool |
| node_pools | [ScaleNodePoolSpec](#cnct.kaas.api.ScaleNodePoolSpec) | repeated | What node pools to scale |






<a name="cnct.kaas.api.ScaleNodePoolReply"></a>

### ScaleNodePoolReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |






<a name="cnct.kaas.api.ScaleNodePoolSpec"></a>

### ScaleNodePoolSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the node pool name to scale |
| count | [int32](#int32) |  | Number of machines to scale |






<a name="cnct.kaas.api.UpgradeClusterMsg"></a>

### UpgradeClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster that we are considering for upgrade |
| version | [string](#string) |  | What version are we upgrading to? |






<a name="cnct.kaas.api.UpgradeClusterReply"></a>

### UpgradeClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Was this a successful request |





 


<a name="cnct.kaas.api.ClusterStatus"></a>

### ClusterStatus
ClusterStatus
Specifies current cluster state.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The cluster is being created. |
| RUNNING | 2 | The cluster has been created and is fully usable. |
| RECONCILING | 3 | Some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The cluster is being deleted |
| ERROR | 5 | The cluster may be unusable |
| DEGRADED | 6 | The cluster requires user action to restore full functionality |


 

 


<a name="cnct.kaas.api.Cluster"></a>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cnct.kaas.api.CreateClusterMsg) | [CreateClusterReply](#cnct.kaas.api.CreateClusterReply) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cnct.kaas.api.GetClusterMsg) | [GetClusterReply](#cnct.kaas.api.GetClusterReply) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cnct.kaas.api.DeleteClusterMsg) | [DeleteClusterReply](#cnct.kaas.api.DeleteClusterReply) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cnct.kaas.api.GetClusterListMsg) | [GetClusterListReply](#cnct.kaas.api.GetClusterListReply) | Will retrieve a list of clusters |
| GetClusterNodesStatus | [GetClusterNodesStatusMsg](#cnct.kaas.api.GetClusterNodesStatusMsg) | [GetClusterNodesStatusReply](#cnct.kaas.api.GetClusterNodesStatusReply) | Will get cluster nodes status for a provisioned cluster |
| GetVersionInformation | [GetVersionMsg](#cnct.kaas.api.GetVersionMsg) | [GetVersionReply](#cnct.kaas.api.GetVersionReply) | Will return version information about api server |
| AddNodePool | [AddNodePoolMsg](#cnct.kaas.api.AddNodePoolMsg) | [AddNodePoolReply](#cnct.kaas.api.AddNodePoolReply) | Will add node pool to a provisioned cluster |
| DeleteNodePool | [DeleteNodePoolMsg](#cnct.kaas.api.DeleteNodePoolMsg) | [DeleteNodePoolReply](#cnct.kaas.api.DeleteNodePoolReply) | Will delete a node pool from a provisioned cluster |
| ScaleNodePool | [ScaleNodePoolMsg](#cnct.kaas.api.ScaleNodePoolMsg) | [ScaleNodePoolReply](#cnct.kaas.api.ScaleNodePoolReply) | Will scale the number of machines in a node pool for a provisioned cluster |
| GetUpgradeClusterInformation | [GetUpgradeClusterInformationMsg](#cnct.kaas.api.GetUpgradeClusterInformationMsg) | [GetUpgradeClusterInformationReply](#cnct.kaas.api.GetUpgradeClusterInformationReply) | Will return upgrade options for a given cluster |
| UpgradeCluster | [UpgradeClusterMsg](#cnct.kaas.api.UpgradeClusterMsg) | [UpgradeClusterReply](#cnct.kaas.api.UpgradeClusterReply) | Will attempt to upgrade a cluster |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

