package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

type CceClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCceClient(hcClient *http_client.HcHttpClient) *CceClient {
	return &CceClient{HcClient: hcClient}
}

func CceClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddNode 纳管节点
//
// 该API用于在指定集群下纳管节点。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) AddNode(request *model.AddNodeRequest) (*model.AddNodeResponse, error) {
	requestDef := GenReqDefForAddNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddNodeResponse), nil
	}
}

// AddNodeInvoker 纳管节点
func (c *CceClient) AddNodeInvoker(request *model.AddNodeRequest) *AddNodeInvoker {
	requestDef := GenReqDefForAddNode()
	return &AddNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AwakeCluster 集群唤醒
//
// 集群唤醒用于唤醒已休眠的集群，唤醒后，将继续收取控制节点资源费用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) AwakeCluster(request *model.AwakeClusterRequest) (*model.AwakeClusterResponse, error) {
	requestDef := GenReqDefForAwakeCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AwakeClusterResponse), nil
	}
}

// AwakeClusterInvoker 集群唤醒
func (c *CceClient) AwakeClusterInvoker(request *model.AwakeClusterRequest) *AwakeClusterInvoker {
	requestDef := GenReqDefForAwakeCluster()
	return &AwakeClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ContinueUpgradeClusterTask 继续执行集群升级任务
//
// 继续执行被暂停的集群升级任务。
// &gt; - 集群升级涉及多维度的组件升级操作，强烈建议统一通过CCE控制台执行交互式升级，降低集群升级过程的业务意外受损风险；
// &gt; - 当前集群升级相关接口受限开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ContinueUpgradeClusterTask(request *model.ContinueUpgradeClusterTaskRequest) (*model.ContinueUpgradeClusterTaskResponse, error) {
	requestDef := GenReqDefForContinueUpgradeClusterTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ContinueUpgradeClusterTaskResponse), nil
	}
}

// ContinueUpgradeClusterTaskInvoker 继续执行集群升级任务
func (c *CceClient) ContinueUpgradeClusterTaskInvoker(request *model.ContinueUpgradeClusterTaskRequest) *ContinueUpgradeClusterTaskInvoker {
	requestDef := GenReqDefForContinueUpgradeClusterTask()
	return &ContinueUpgradeClusterTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAddonInstance 创建AddonInstance
//
// 根据提供的插件模板，安装插件实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateAddonInstance(request *model.CreateAddonInstanceRequest) (*model.CreateAddonInstanceResponse, error) {
	requestDef := GenReqDefForCreateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAddonInstanceResponse), nil
	}
}

// CreateAddonInstanceInvoker 创建AddonInstance
func (c *CceClient) CreateAddonInstanceInvoker(request *model.CreateAddonInstanceRequest) *CreateAddonInstanceInvoker {
	requestDef := GenReqDefForCreateAddonInstance()
	return &CreateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCloudPersistentVolumeClaims 创建PVC（待废弃）
//
// 该API用于在指定的Namespace下通过云存储服务中的云存储（EVS、SFS、OBS）去创建PVC（PersistentVolumeClaim）。该API待废弃，请使用Kubernetes PVC相关接口。
//
//
// &gt;存储管理的URL格式为：https://{clusterid}.Endpoint/uri。其中{clusterid}为集群ID，uri为资源路径，也即API访问的路径。如果使用https://Endpoint/uri，则必须指定请求header中的X-Cluster-ID参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateCloudPersistentVolumeClaims(request *model.CreateCloudPersistentVolumeClaimsRequest) (*model.CreateCloudPersistentVolumeClaimsResponse, error) {
	requestDef := GenReqDefForCreateCloudPersistentVolumeClaims()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudPersistentVolumeClaimsResponse), nil
	}
}

// CreateCloudPersistentVolumeClaimsInvoker 创建PVC（待废弃）
func (c *CceClient) CreateCloudPersistentVolumeClaimsInvoker(request *model.CreateCloudPersistentVolumeClaimsRequest) *CreateCloudPersistentVolumeClaimsInvoker {
	requestDef := GenReqDefForCreateCloudPersistentVolumeClaims()
	return &CreateCloudPersistentVolumeClaimsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建集群
//
// 该API用于创建一个空集群（即只有控制节点Master，没有工作节点Node）。请在调用本接口完成集群创建之后，通过[创建节点](cce_02_0242.xml)添加节点。
//
//
// &gt;   - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
// &gt;   - 调用该接口创建集群时，默认不安装ICAgent，若需安装ICAgent，可在请求Body参数的annotations中加入\&quot;cluster.install.addons.external/install\&quot;:\&quot;[{\&quot;addonTemplateName\&quot;:\&quot;icagent\&quot;}]\&quot;的集群注解，将在创建集群时自动安装ICAgent。ICAgent是应用性能管理APM的采集代理，运行在应用所在的服务器上，用于实时采集探针所获取的数据，安装ICAgent是使用应用性能管理APM的前提。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建集群
func (c *CceClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKubernetesClusterCert 获取集群证书
//
// 该API用于获取指定集群的证书信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateKubernetesClusterCert(request *model.CreateKubernetesClusterCertRequest) (*model.CreateKubernetesClusterCertResponse, error) {
	requestDef := GenReqDefForCreateKubernetesClusterCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKubernetesClusterCertResponse), nil
	}
}

// CreateKubernetesClusterCertInvoker 获取集群证书
func (c *CceClient) CreateKubernetesClusterCertInvoker(request *model.CreateKubernetesClusterCertRequest) *CreateKubernetesClusterCertInvoker {
	requestDef := GenReqDefForCreateKubernetesClusterCert()
	return &CreateKubernetesClusterCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNode 创建节点
//
// 该API用于在指定集群下创建节点。
// &gt; - 若无集群，请先[创建集群](cce_02_0236.xml)。
// &gt; - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateNode(request *model.CreateNodeRequest) (*model.CreateNodeResponse, error) {
	requestDef := GenReqDefForCreateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodeResponse), nil
	}
}

// CreateNodeInvoker 创建节点
func (c *CceClient) CreateNodeInvoker(request *model.CreateNodeRequest) *CreateNodeInvoker {
	requestDef := GenReqDefForCreateNode()
	return &CreateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNodePool 创建节点池
//
// 该API用于在指定集群下创建节点池。仅支持集群在处于可用、扩容、缩容状态时调用。
//
// 1.21版本的集群创建节点池时支持绑定安全组，每个节点池最多绑定五个安全组。
//
// 更新节点池的安全组后，只针对新创的pod生效，建议驱逐节点上原有的pod。
//
// &gt; 若无集群，请先[创建集群](cce_02_0236.xml)。
// &gt; 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateNodePool(request *model.CreateNodePoolRequest) (*model.CreateNodePoolResponse, error) {
	requestDef := GenReqDefForCreateNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodePoolResponse), nil
	}
}

// CreateNodePoolInvoker 创建节点池
func (c *CceClient) CreateNodePoolInvoker(request *model.CreateNodePoolRequest) *CreateNodePoolInvoker {
	requestDef := GenReqDefForCreateNodePool()
	return &CreateNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddonInstance 删除AddonInstance
//
// 删除插件实例的功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteAddonInstance(request *model.DeleteAddonInstanceRequest) (*model.DeleteAddonInstanceResponse, error) {
	requestDef := GenReqDefForDeleteAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddonInstanceResponse), nil
	}
}

// DeleteAddonInstanceInvoker 删除AddonInstance
func (c *CceClient) DeleteAddonInstanceInvoker(request *model.DeleteAddonInstanceRequest) *DeleteAddonInstanceInvoker {
	requestDef := GenReqDefForDeleteAddonInstance()
	return &DeleteAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCloudPersistentVolumeClaims 删除PVC（待废弃）
//
// 该API用于删除指定Namespace下的PVC（PersistentVolumeClaim）对象，并可以选择保留后端的云存储。该API待废弃，请使用Kubernetes PVC相关接口。
// &gt;存储管理的URL格式为：https://{clusterid}.Endpoint/uri。其中{clusterid}为集群ID，uri为资源路径，也即API访问的路径。如果使用https://Endpoint/uri，则必须指定请求header中的X-Cluster-ID参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteCloudPersistentVolumeClaims(request *model.DeleteCloudPersistentVolumeClaimsRequest) (*model.DeleteCloudPersistentVolumeClaimsResponse, error) {
	requestDef := GenReqDefForDeleteCloudPersistentVolumeClaims()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudPersistentVolumeClaimsResponse), nil
	}
}

// DeleteCloudPersistentVolumeClaimsInvoker 删除PVC（待废弃）
func (c *CceClient) DeleteCloudPersistentVolumeClaimsInvoker(request *model.DeleteCloudPersistentVolumeClaimsRequest) *DeleteCloudPersistentVolumeClaimsInvoker {
	requestDef := GenReqDefForDeleteCloudPersistentVolumeClaims()
	return &DeleteCloudPersistentVolumeClaimsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除集群
//
// 该API用于删除一个指定的集群。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除集群
func (c *CceClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNode 删除节点
//
// 该API用于删除指定的节点。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteNode(request *model.DeleteNodeRequest) (*model.DeleteNodeResponse, error) {
	requestDef := GenReqDefForDeleteNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodeResponse), nil
	}
}

// DeleteNodeInvoker 删除节点
func (c *CceClient) DeleteNodeInvoker(request *model.DeleteNodeRequest) *DeleteNodeInvoker {
	requestDef := GenReqDefForDeleteNode()
	return &DeleteNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNodePool 删除节点池
//
// 该API用于删除指定的节点池。
// &gt; 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteNodePool(request *model.DeleteNodePoolRequest) (*model.DeleteNodePoolResponse, error) {
	requestDef := GenReqDefForDeleteNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodePoolResponse), nil
	}
}

// DeleteNodePoolInvoker 删除节点池
func (c *CceClient) DeleteNodePoolInvoker(request *model.DeleteNodePoolRequest) *DeleteNodePoolInvoker {
	requestDef := GenReqDefForDeleteNodePool()
	return &DeleteNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HibernateCluster 集群休眠
//
// 集群休眠用于将运行中的集群置于休眠状态，休眠后，将不再收取控制节点资源费用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) HibernateCluster(request *model.HibernateClusterRequest) (*model.HibernateClusterResponse, error) {
	requestDef := GenReqDefForHibernateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HibernateClusterResponse), nil
	}
}

// HibernateClusterInvoker 集群休眠
func (c *CceClient) HibernateClusterInvoker(request *model.HibernateClusterRequest) *HibernateClusterInvoker {
	requestDef := GenReqDefForHibernateCluster()
	return &HibernateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddonInstances 获取AddonInstance列表
//
// 获取集群所有已安装插件实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListAddonInstances(request *model.ListAddonInstancesRequest) (*model.ListAddonInstancesResponse, error) {
	requestDef := GenReqDefForListAddonInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonInstancesResponse), nil
	}
}

// ListAddonInstancesInvoker 获取AddonInstance列表
func (c *CceClient) ListAddonInstancesInvoker(request *model.ListAddonInstancesRequest) *ListAddonInstancesInvoker {
	requestDef := GenReqDefForListAddonInstances()
	return &ListAddonInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddonTemplates 查询AddonTemplates列表
//
// 插件模板查询接口，查询插件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListAddonTemplates(request *model.ListAddonTemplatesRequest) (*model.ListAddonTemplatesResponse, error) {
	requestDef := GenReqDefForListAddonTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddonTemplatesResponse), nil
	}
}

// ListAddonTemplatesInvoker 查询AddonTemplates列表
func (c *CceClient) ListAddonTemplatesInvoker(request *model.ListAddonTemplatesRequest) *ListAddonTemplatesInvoker {
	requestDef := GenReqDefForListAddonTemplates()
	return &ListAddonTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 获取指定项目下的集群
//
// 该API用于获取指定项目下所有集群的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 获取指定项目下的集群
func (c *CceClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodePools 获取集群下所有节点池
//
// 该API用于获取集群下所有节点池。
// &gt; - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
// &gt; - nodepool是集群中具有相同配置的节点实例的子集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListNodePools(request *model.ListNodePoolsRequest) (*model.ListNodePoolsResponse, error) {
	requestDef := GenReqDefForListNodePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodePoolsResponse), nil
	}
}

// ListNodePoolsInvoker 获取集群下所有节点池
func (c *CceClient) ListNodePoolsInvoker(request *model.ListNodePoolsRequest) *ListNodePoolsInvoker {
	requestDef := GenReqDefForListNodePools()
	return &ListNodePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodes 获取集群下所有节点
//
// 该API用于通过集群ID获取指定集群下所有节点的详细信息。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListNodes(request *model.ListNodesRequest) (*model.ListNodesResponse, error) {
	requestDef := GenReqDefForListNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodesResponse), nil
	}
}

// ListNodesInvoker 获取集群下所有节点
func (c *CceClient) ListNodesInvoker(request *model.ListNodesRequest) *ListNodesInvoker {
	requestDef := GenReqDefForListNodes()
	return &ListNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateNode 节点迁移
//
// 该API用于在指定集群下迁移节点到另一集群（仅支持在同一VPC下的不同集群之间进行迁移）。
// [CCE Turbo集群下弹性云服务-物理机类型节点不支持迁移。](tag:hws,hws_hk,dt)
//
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) MigrateNode(request *model.MigrateNodeRequest) (*model.MigrateNodeResponse, error) {
	requestDef := GenReqDefForMigrateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateNodeResponse), nil
	}
}

// MigrateNodeInvoker 节点迁移
func (c *CceClient) MigrateNodeInvoker(request *model.MigrateNodeRequest) *MigrateNodeInvoker {
	requestDef := GenReqDefForMigrateNode()
	return &MigrateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseUpgradeClusterTask 暂停集群升级任务
//
// 暂停集群升级任务。
// &gt; - 集群升级涉及多维度的组件升级操作，强烈建议统一通过CCE控制台执行交互式升级，降低集群升级过程的业务意外受损风险；
// &gt; - 当前集群升级相关接口受限开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) PauseUpgradeClusterTask(request *model.PauseUpgradeClusterTaskRequest) (*model.PauseUpgradeClusterTaskResponse, error) {
	requestDef := GenReqDefForPauseUpgradeClusterTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseUpgradeClusterTaskResponse), nil
	}
}

// PauseUpgradeClusterTaskInvoker 暂停集群升级任务
func (c *CceClient) PauseUpgradeClusterTaskInvoker(request *model.PauseUpgradeClusterTaskRequest) *PauseUpgradeClusterTaskInvoker {
	requestDef := GenReqDefForPauseUpgradeClusterTask()
	return &PauseUpgradeClusterTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveNode 节点移除
//
// 该API用于在指定集群下移除节点。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) RemoveNode(request *model.RemoveNodeRequest) (*model.RemoveNodeResponse, error) {
	requestDef := GenReqDefForRemoveNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveNodeResponse), nil
	}
}

// RemoveNodeInvoker 节点移除
func (c *CceClient) RemoveNodeInvoker(request *model.RemoveNodeRequest) *RemoveNodeInvoker {
	requestDef := GenReqDefForRemoveNode()
	return &RemoveNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetNode 重置节点
//
// 该API用于在指定集群下重置节点。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ResetNode(request *model.ResetNodeRequest) (*model.ResetNodeResponse, error) {
	requestDef := GenReqDefForResetNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetNodeResponse), nil
	}
}

// ResetNodeInvoker 重置节点
func (c *CceClient) ResetNodeInvoker(request *model.ResetNodeRequest) *ResetNodeInvoker {
	requestDef := GenReqDefForResetNode()
	return &ResetNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryUpgradeClusterTask 重试集群升级任务
//
// 重新执行失败的集群升级任务。
// &gt; - 集群升级涉及多维度的组件升级操作，强烈建议统一通过CCE控制台执行交互式升级，降低集群升级过程的业务意外受损风险；
// &gt; - 当前集群升级相关接口受限开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) RetryUpgradeClusterTask(request *model.RetryUpgradeClusterTaskRequest) (*model.RetryUpgradeClusterTaskResponse, error) {
	requestDef := GenReqDefForRetryUpgradeClusterTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryUpgradeClusterTaskResponse), nil
	}
}

// RetryUpgradeClusterTaskInvoker 重试集群升级任务
func (c *CceClient) RetryUpgradeClusterTaskInvoker(request *model.RetryUpgradeClusterTaskRequest) *RetryUpgradeClusterTaskInvoker {
	requestDef := GenReqDefForRetryUpgradeClusterTask()
	return &RetryUpgradeClusterTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAddonInstance 获取AddonInstance详情
//
// 获取插件实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowAddonInstance(request *model.ShowAddonInstanceRequest) (*model.ShowAddonInstanceResponse, error) {
	requestDef := GenReqDefForShowAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAddonInstanceResponse), nil
	}
}

// ShowAddonInstanceInvoker 获取AddonInstance详情
func (c *CceClient) ShowAddonInstanceInvoker(request *model.ShowAddonInstanceRequest) *ShowAddonInstanceInvoker {
	requestDef := GenReqDefForShowAddonInstance()
	return &ShowAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCluster 获取指定的集群
//
// 该API用于获取指定集群的详细信息。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// ShowClusterInvoker 获取指定的集群
func (c *CceClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterEndpoints 获取集群访问的地址
//
// 该API用于通过集群ID获取集群访问的地址，包括PrivateIP(HA集群返回VIP)与PublicIP
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowClusterEndpoints(request *model.ShowClusterEndpointsRequest) (*model.ShowClusterEndpointsResponse, error) {
	requestDef := GenReqDefForShowClusterEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterEndpointsResponse), nil
	}
}

// ShowClusterEndpointsInvoker 获取集群访问的地址
func (c *CceClient) ShowClusterEndpointsInvoker(request *model.ShowClusterEndpointsRequest) *ShowClusterEndpointsInvoker {
	requestDef := GenReqDefForShowClusterEndpoints()
	return &ShowClusterEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 获取任务信息
//
// 该API用于获取任务信息。通过某一任务请求下发后返回的jobID来查询指定任务的进度。
// &gt; - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
// &gt; - 该接口通常使用场景为：
// &gt;   - 创建、删除集群时，查询相应任务的进度。
// &gt;   - 创建、删除节点时，查询相应任务的进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 获取任务信息
func (c *CceClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNode 获取指定的节点
//
// 该API用于通过节点ID获取指定节点的详细信息。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowNode(request *model.ShowNodeRequest) (*model.ShowNodeResponse, error) {
	requestDef := GenReqDefForShowNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeResponse), nil
	}
}

// ShowNodeInvoker 获取指定的节点
func (c *CceClient) ShowNodeInvoker(request *model.ShowNodeRequest) *ShowNodeInvoker {
	requestDef := GenReqDefForShowNode()
	return &ShowNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodePool 获取指定的节点池
//
// 该API用于获取指定节点池的详细信息。
// &gt; 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowNodePool(request *model.ShowNodePoolRequest) (*model.ShowNodePoolResponse, error) {
	requestDef := GenReqDefForShowNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodePoolResponse), nil
	}
}

// ShowNodePoolInvoker 获取指定的节点池
func (c *CceClient) ShowNodePoolInvoker(request *model.ShowNodePoolRequest) *ShowNodePoolInvoker {
	requestDef := GenReqDefForShowNodePool()
	return &ShowNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询CCE服务下的资源配额
//
// 该API用于查询CCE服务下的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询CCE服务下的资源配额
func (c *CceClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpgradeClusterTask 获取集群升级任务详情
//
// 获取集群升级任务详情，任务ID由调用集群升级API后从响应体中uid字段获取。
// &gt; - 集群升级涉及多维度的组件升级操作，强烈建议统一通过CCE控制台执行交互式升级，降低集群升级过程的业务意外受损风险；
// &gt; - 当前集群升级相关接口受限开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowUpgradeClusterTask(request *model.ShowUpgradeClusterTaskRequest) (*model.ShowUpgradeClusterTaskResponse, error) {
	requestDef := GenReqDefForShowUpgradeClusterTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeClusterTaskResponse), nil
	}
}

// ShowUpgradeClusterTaskInvoker 获取集群升级任务详情
func (c *CceClient) ShowUpgradeClusterTaskInvoker(request *model.ShowUpgradeClusterTaskRequest) *ShowUpgradeClusterTaskInvoker {
	requestDef := GenReqDefForShowUpgradeClusterTask()
	return &ShowUpgradeClusterTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAddonInstance 更新AddonInstance
//
// 更新插件实例的功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpdateAddonInstance(request *model.UpdateAddonInstanceRequest) (*model.UpdateAddonInstanceResponse, error) {
	requestDef := GenReqDefForUpdateAddonInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddonInstanceResponse), nil
	}
}

// UpdateAddonInstanceInvoker 更新AddonInstance
func (c *CceClient) UpdateAddonInstanceInvoker(request *model.UpdateAddonInstanceRequest) *UpdateAddonInstanceInvoker {
	requestDef := GenReqDefForUpdateAddonInstance()
	return &UpdateAddonInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCluster 更新指定的集群
//
// 该API用于更新指定的集群。
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpdateCluster(request *model.UpdateClusterRequest) (*model.UpdateClusterResponse, error) {
	requestDef := GenReqDefForUpdateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterResponse), nil
	}
}

// UpdateClusterInvoker 更新指定的集群
func (c *CceClient) UpdateClusterInvoker(request *model.UpdateClusterRequest) *UpdateClusterInvoker {
	requestDef := GenReqDefForUpdateCluster()
	return &UpdateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterEip 绑定、解绑集群公网apiserver地址
//
// 该API用于通过集群ID绑定、解绑集群公网apiserver地址
// &gt;集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpdateClusterEip(request *model.UpdateClusterEipRequest) (*model.UpdateClusterEipResponse, error) {
	requestDef := GenReqDefForUpdateClusterEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterEipResponse), nil
	}
}

// UpdateClusterEipInvoker 绑定、解绑集群公网apiserver地址
func (c *CceClient) UpdateClusterEipInvoker(request *model.UpdateClusterEipRequest) *UpdateClusterEipInvoker {
	requestDef := GenReqDefForUpdateClusterEip()
	return &UpdateClusterEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNode 更新指定的节点
//
// 该API用于更新指定的节点。
// &gt; - 当前仅支持更新metadata下的name字段，即节点的名字。
// &gt; - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpdateNode(request *model.UpdateNodeRequest) (*model.UpdateNodeResponse, error) {
	requestDef := GenReqDefForUpdateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeResponse), nil
	}
}

// UpdateNodeInvoker 更新指定的节点
func (c *CceClient) UpdateNodeInvoker(request *model.UpdateNodeRequest) *UpdateNodeInvoker {
	requestDef := GenReqDefForUpdateNode()
	return &UpdateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNodePool 更新指定节点池
//
// 该API用于更新指定的节点池。仅支持集群在处于可用、扩容、缩容状态时调用。
//
//
// &gt; - 集群管理的URL格式为：https://Endpoint/uri。其中uri为资源路径，也即API访问的路径
//
// &gt; - 当前仅支持更新节点池名称，spec下的initialNodeCount，k8sTags，taints，login，userTags与节点池的扩缩容配置相关字段。若此次更新未设置相关值，默认更新为初始值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpdateNodePool(request *model.UpdateNodePoolRequest) (*model.UpdateNodePoolResponse, error) {
	requestDef := GenReqDefForUpdateNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodePoolResponse), nil
	}
}

// UpdateNodePoolInvoker 更新指定节点池
func (c *CceClient) UpdateNodePoolInvoker(request *model.UpdateNodePoolRequest) *UpdateNodePoolInvoker {
	requestDef := GenReqDefForUpdateNodePool()
	return &UpdateNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeCluster 集群升级
//
// 集群升级。
// &gt; - 集群升级涉及多维度的组件升级操作，强烈建议统一通过CCE控制台执行交互式升级，降低集群升级过程的业务意外受损风险；
// &gt; - 当前集群升级相关接口受限开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) UpgradeCluster(request *model.UpgradeClusterRequest) (*model.UpgradeClusterResponse, error) {
	requestDef := GenReqDefForUpgradeCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeClusterResponse), nil
	}
}

// UpgradeClusterInvoker 集群升级
func (c *CceClient) UpgradeClusterInvoker(request *model.UpgradeClusterRequest) *UpgradeClusterInvoker {
	requestDef := GenReqDefForUpgradeCluster()
	return &UpgradeClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 查询API版本信息列表。
//
// 该API用于查询CCE服务当前支持的API版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 查询API版本信息列表。
func (c *CceClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
