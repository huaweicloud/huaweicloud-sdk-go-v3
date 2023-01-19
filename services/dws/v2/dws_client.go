package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type DwsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDwsClient(hcClient *http_client.HcHttpClient) *DwsClient {
	return &DwsClient{HcClient: hcClient}
}

func DwsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddWorkloadQueue 添加工作负载队列
//
// 添加工作负载队列
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AddWorkloadQueue(request *model.AddWorkloadQueueRequest) (*model.AddWorkloadQueueResponse, error) {
	requestDef := GenReqDefForAddWorkloadQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddWorkloadQueueResponse), nil
	}
}

// AddWorkloadQueueInvoker 添加工作负载队列
func (c *DwsClient) AddWorkloadQueueInvoker(request *model.AddWorkloadQueueRequest) *AddWorkloadQueueInvoker {
	requestDef := GenReqDefForAddWorkloadQueue()
	return &AddWorkloadQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateEip 集群绑定EIP
//
// 集群绑定Eip
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AssociateEip(request *model.AssociateEipRequest) (*model.AssociateEipResponse, error) {
	requestDef := GenReqDefForAssociateEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateEipResponse), nil
	}
}

// AssociateEipInvoker 集群绑定EIP
func (c *DwsClient) AssociateEipInvoker(request *model.AssociateEipRequest) *AssociateEipInvoker {
	requestDef := GenReqDefForAssociateEip()
	return &AssociateEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateElb 集群绑定ELB
//
// 集群绑定Elb接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AssociateElb(request *model.AssociateElbRequest) (*model.AssociateElbResponse, error) {
	requestDef := GenReqDefForAssociateElb()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateElbResponse), nil
	}
}

// AssociateElbInvoker 集群绑定ELB
func (c *DwsClient) AssociateElbInvoker(request *model.AssociateElbRequest) *AssociateElbInvoker {
	requestDef := GenReqDefForAssociateElb()
	return &AssociateElbInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateClusterCn 批量增加CN节点
//
// 当用户集群创建后，实际需要的CN数量会随着业务需求而发生变化，因此管理CN节点功能的实现使用户可以根据实际需求动态调整集群CN数量。
// - 增删CN节点过程中不允许执行其他运维操作。
// - 增删CN节点过程中需要停止业务操作，建议在业务低峰期或业务中断情况下进行操作。
// - 增删CN节点时发生故障且回滚失败，需要用户登录后台进行处理，处理方案请参见《故障排除》中的“集群使用&gt;增删CN回滚失败”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) BatchCreateClusterCn(request *model.BatchCreateClusterCnRequest) (*model.BatchCreateClusterCnResponse, error) {
	requestDef := GenReqDefForBatchCreateClusterCn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateClusterCnResponse), nil
	}
}

// BatchCreateClusterCnInvoker 批量增加CN节点
func (c *DwsClient) BatchCreateClusterCnInvoker(request *model.BatchCreateClusterCnRequest) *BatchCreateClusterCnInvoker {
	requestDef := GenReqDefForBatchCreateClusterCn()
	return &BatchCreateClusterCnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceTag 批量添加标签
//
// 为指定集群批量添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) BatchCreateResourceTag(request *model.BatchCreateResourceTagRequest) (*model.BatchCreateResourceTagResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceTagResponse), nil
	}
}

// BatchCreateResourceTagInvoker 批量添加标签
func (c *DwsClient) BatchCreateResourceTagInvoker(request *model.BatchCreateResourceTagRequest) *BatchCreateResourceTagInvoker {
	requestDef := GenReqDefForBatchCreateResourceTag()
	return &BatchCreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteClusterCn 批量删除CN节点
//
// 当用户集群创建后，实际需要的CN数量会随着业务需求而发生变化，因此管理CN节点功能的实现使用户可以根据实际需求动态调整集群CN数量。
// - 增删CN节点过程中不允许执行其他运维操作。
// - 增删CN节点过程中需要停止业务操作，建议在业务低峰期或业务中断情况下进行操作。
// - 增删CN节点时发生故障且回滚失败，需要用户登录后台进行处理，处理方案请参见《故障排除》中的“集群使用&gt;增删CN回滚失败”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) BatchDeleteClusterCn(request *model.BatchDeleteClusterCnRequest) (*model.BatchDeleteClusterCnResponse, error) {
	requestDef := GenReqDefForBatchDeleteClusterCn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteClusterCnResponse), nil
	}
}

// BatchDeleteClusterCnInvoker 批量删除CN节点
func (c *DwsClient) BatchDeleteClusterCnInvoker(request *model.BatchDeleteClusterCnRequest) *BatchDeleteClusterCnInvoker {
	requestDef := GenReqDefForBatchDeleteClusterCn()
	return &BatchDeleteClusterCnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResourceTag 批量删除标签
//
// 为指定集群批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) BatchDeleteResourceTag(request *model.BatchDeleteResourceTagRequest) (*model.BatchDeleteResourceTagResponse, error) {
	requestDef := GenReqDefForBatchDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourceTagResponse), nil
	}
}

// BatchDeleteResourceTagInvoker 批量删除标签
func (c *DwsClient) BatchDeleteResourceTagInvoker(request *model.BatchDeleteResourceTagRequest) *BatchDeleteResourceTagInvoker {
	requestDef := GenReqDefForBatchDeleteResourceTag()
	return &BatchDeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelReadonlyCluster 解除只读
//
// 当集群进入只读状态时，无法进行数据库相关操作，用户可以在管理控制台解除集群的只读状态。触发只读状态可能是由于磁盘使用率过高，因此需要对集群数据进行清理或扩容。
// - 解除只读支持1.7.2及以上版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CancelReadonlyCluster(request *model.CancelReadonlyClusterRequest) (*model.CancelReadonlyClusterResponse, error) {
	requestDef := GenReqDefForCancelReadonlyCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelReadonlyClusterResponse), nil
	}
}

// CancelReadonlyClusterInvoker 解除只读
func (c *DwsClient) CancelReadonlyClusterInvoker(request *model.CancelReadonlyClusterRequest) *CancelReadonlyClusterInvoker {
	requestDef := GenReqDefForCancelReadonlyCluster()
	return &CancelReadonlyClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckCluster 创建集群前检查
//
// 创建集群前预检查
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CheckCluster(request *model.CheckClusterRequest) (*model.CheckClusterResponse, error) {
	requestDef := GenReqDefForCheckCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckClusterResponse), nil
	}
}

// CheckClusterInvoker 创建集群前检查
func (c *DwsClient) CheckClusterInvoker(request *model.CheckClusterRequest) *CheckClusterInvoker {
	requestDef := GenReqDefForCheckCluster()
	return &CheckClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopySnapshot 复制快照
//
// 该接口用于复制一个自动快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CopySnapshot(request *model.CopySnapshotRequest) (*model.CopySnapshotResponse, error) {
	requestDef := GenReqDefForCopySnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopySnapshotResponse), nil
	}
}

// CopySnapshotInvoker 复制快照
func (c *DwsClient) CopySnapshotInvoker(request *model.CopySnapshotRequest) *CopySnapshotInvoker {
	requestDef := GenReqDefForCopySnapshot()
	return &CopySnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarmSub 创建告警订阅
//
// 创建告警订阅
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateAlarmSub(request *model.CreateAlarmSubRequest) (*model.CreateAlarmSubResponse, error) {
	requestDef := GenReqDefForCreateAlarmSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmSubResponse), nil
	}
}

// CreateAlarmSubInvoker 创建告警订阅
func (c *DwsClient) CreateAlarmSubInvoker(request *model.CreateAlarmSubRequest) *CreateAlarmSubInvoker {
	requestDef := GenReqDefForCreateAlarmSub()
	return &CreateAlarmSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建集群
//
// 该接口用于创建集群。
// 集群必须要运行在VPC之内，创建集群前，您需要先创建VPC，并获取VPC和子网的id。
// 该接口为异步接口，创建集群需要10～15分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建集群
func (c *DwsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterDns 申请域名
//
// 为指定集群申请域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateClusterDns(request *model.CreateClusterDnsRequest) (*model.CreateClusterDnsResponse, error) {
	requestDef := GenReqDefForCreateClusterDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterDnsResponse), nil
	}
}

// CreateClusterDnsInvoker 申请域名
func (c *DwsClient) CreateClusterDnsInvoker(request *model.CreateClusterDnsRequest) *CreateClusterDnsInvoker {
	requestDef := GenReqDefForCreateClusterDns()
	return &CreateClusterDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterWorkload 设置资源管理
//
// 设置资源管理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateClusterWorkload(request *model.CreateClusterWorkloadRequest) (*model.CreateClusterWorkloadResponse, error) {
	requestDef := GenReqDefForCreateClusterWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterWorkloadResponse), nil
	}
}

// CreateClusterWorkloadInvoker 设置资源管理
func (c *DwsClient) CreateClusterWorkloadInvoker(request *model.CreateClusterWorkloadRequest) *CreateClusterWorkloadInvoker {
	requestDef := GenReqDefForCreateClusterWorkload()
	return &CreateClusterWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataSource 创建数据源
//
// 该接口用于创建一个数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateDataSource(request *model.CreateDataSourceRequest) (*model.CreateDataSourceResponse, error) {
	requestDef := GenReqDefForCreateDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataSourceResponse), nil
	}
}

// CreateDataSourceInvoker 创建数据源
func (c *DwsClient) CreateDataSourceInvoker(request *model.CreateDataSourceRequest) *CreateDataSourceInvoker {
	requestDef := GenReqDefForCreateDataSource()
	return &CreateDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDisasterRecovery 创建容灾
//
// 创建容灾
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateDisasterRecovery(request *model.CreateDisasterRecoveryRequest) (*model.CreateDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForCreateDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDisasterRecoveryResponse), nil
	}
}

// CreateDisasterRecoveryInvoker 创建容灾
func (c *DwsClient) CreateDisasterRecoveryInvoker(request *model.CreateDisasterRecoveryRequest) *CreateDisasterRecoveryInvoker {
	requestDef := GenReqDefForCreateDisasterRecovery()
	return &CreateDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEventSub 创建订阅事件
//
// 添加订阅的事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateEventSub(request *model.CreateEventSubRequest) (*model.CreateEventSubResponse, error) {
	requestDef := GenReqDefForCreateEventSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventSubResponse), nil
	}
}

// CreateEventSubInvoker 创建订阅事件
func (c *DwsClient) CreateEventSubInvoker(request *model.CreateEventSubRequest) *CreateEventSubInvoker {
	requestDef := GenReqDefForCreateEventSub()
	return &CreateEventSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSnapshot 创建快照
//
// 该接口用于为指定集群创建快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

// CreateSnapshotInvoker 创建快照
func (c *DwsClient) CreateSnapshotInvoker(request *model.CreateSnapshotRequest) *CreateSnapshotInvoker {
	requestDef := GenReqDefForCreateSnapshot()
	return &CreateSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSnapshotPolicy 添加快照策略
//
// 该接口用于设置快照策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateSnapshotPolicy(request *model.CreateSnapshotPolicyRequest) (*model.CreateSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForCreateSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotPolicyResponse), nil
	}
}

// CreateSnapshotPolicyInvoker 添加快照策略
func (c *DwsClient) CreateSnapshotPolicyInvoker(request *model.CreateSnapshotPolicyRequest) *CreateSnapshotPolicyInvoker {
	requestDef := GenReqDefForCreateSnapshotPolicy()
	return &CreateSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkloadPlan 添加工作负载计划
//
// 添加工作负载计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateWorkloadPlan(request *model.CreateWorkloadPlanRequest) (*model.CreateWorkloadPlanResponse, error) {
	requestDef := GenReqDefForCreateWorkloadPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkloadPlanResponse), nil
	}
}

// CreateWorkloadPlanInvoker 添加工作负载计划
func (c *DwsClient) CreateWorkloadPlanInvoker(request *model.CreateWorkloadPlanRequest) *CreateWorkloadPlanInvoker {
	requestDef := GenReqDefForCreateWorkloadPlan()
	return &CreateWorkloadPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarmSub 删除告警订阅
//
// 删除订阅的告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteAlarmSub(request *model.DeleteAlarmSubRequest) (*model.DeleteAlarmSubResponse, error) {
	requestDef := GenReqDefForDeleteAlarmSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmSubResponse), nil
	}
}

// DeleteAlarmSubInvoker 删除告警订阅
func (c *DwsClient) DeleteAlarmSubInvoker(request *model.DeleteAlarmSubRequest) *DeleteAlarmSubInvoker {
	requestDef := GenReqDefForDeleteAlarmSub()
	return &DeleteAlarmSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除集群
//
// 此接口用于删除集群。集群删除后将释放此集群的所有资源，包括客户数据。为了安全起见，请在删除集群前为这个集群创建快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除集群
func (c *DwsClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClusterDns 删除集群域名
//
// 删除指定集群域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteClusterDns(request *model.DeleteClusterDnsRequest) (*model.DeleteClusterDnsResponse, error) {
	requestDef := GenReqDefForDeleteClusterDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterDnsResponse), nil
	}
}

// DeleteClusterDnsInvoker 删除集群域名
func (c *DwsClient) DeleteClusterDnsInvoker(request *model.DeleteClusterDnsRequest) *DeleteClusterDnsInvoker {
	requestDef := GenReqDefForDeleteClusterDns()
	return &DeleteClusterDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataSource 删除数据源
//
// 该接口用于删除一个数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteDataSource(request *model.DeleteDataSourceRequest) (*model.DeleteDataSourceResponse, error) {
	requestDef := GenReqDefForDeleteDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataSourceResponse), nil
	}
}

// DeleteDataSourceInvoker 删除数据源
func (c *DwsClient) DeleteDataSourceInvoker(request *model.DeleteDataSourceRequest) *DeleteDataSourceInvoker {
	requestDef := GenReqDefForDeleteDataSource()
	return &DeleteDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDisasterRecovery 删除容灾
//
// 删除容灾。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteDisasterRecovery(request *model.DeleteDisasterRecoveryRequest) (*model.DeleteDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForDeleteDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDisasterRecoveryResponse), nil
	}
}

// DeleteDisasterRecoveryInvoker 删除容灾
func (c *DwsClient) DeleteDisasterRecoveryInvoker(request *model.DeleteDisasterRecoveryRequest) *DeleteDisasterRecoveryInvoker {
	requestDef := GenReqDefForDeleteDisasterRecovery()
	return &DeleteDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEventSub 删除订阅事件
//
// 删除订阅的事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteEventSub(request *model.DeleteEventSubRequest) (*model.DeleteEventSubResponse, error) {
	requestDef := GenReqDefForDeleteEventSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventSubResponse), nil
	}
}

// DeleteEventSubInvoker 删除订阅事件
func (c *DwsClient) DeleteEventSubInvoker(request *model.DeleteEventSubRequest) *DeleteEventSubInvoker {
	requestDef := GenReqDefForDeleteEventSub()
	return &DeleteEventSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSnapshot 删除快照
//
// 该接口用于删除一个指定手动快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

// DeleteSnapshotInvoker 删除快照
func (c *DwsClient) DeleteSnapshotInvoker(request *model.DeleteSnapshotRequest) *DeleteSnapshotInvoker {
	requestDef := GenReqDefForDeleteSnapshot()
	return &DeleteSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSnapshotPolicy 删除快照策略
//
// 该接口用于删除一个快照策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteSnapshotPolicy(request *model.DeleteSnapshotPolicyRequest) (*model.DeleteSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotPolicyResponse), nil
	}
}

// DeleteSnapshotPolicyInvoker 删除快照策略
func (c *DwsClient) DeleteSnapshotPolicyInvoker(request *model.DeleteSnapshotPolicyRequest) *DeleteSnapshotPolicyInvoker {
	requestDef := GenReqDefForDeleteSnapshotPolicy()
	return &DeleteSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkloadQueue 删除工作负载队列
//
// 该接口用于删除工作负载队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteWorkloadQueue(request *model.DeleteWorkloadQueueRequest) (*model.DeleteWorkloadQueueResponse, error) {
	requestDef := GenReqDefForDeleteWorkloadQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkloadQueueResponse), nil
	}
}

// DeleteWorkloadQueueInvoker 删除工作负载队列
func (c *DwsClient) DeleteWorkloadQueueInvoker(request *model.DeleteWorkloadQueueRequest) *DeleteWorkloadQueueInvoker {
	requestDef := GenReqDefForDeleteWorkloadQueue()
	return &DeleteWorkloadQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateEip 集群解绑EIP
//
// 集群解绑Eip
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DisassociateEip(request *model.DisassociateEipRequest) (*model.DisassociateEipResponse, error) {
	requestDef := GenReqDefForDisassociateEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateEipResponse), nil
	}
}

// DisassociateEipInvoker 集群解绑EIP
func (c *DwsClient) DisassociateEipInvoker(request *model.DisassociateEipRequest) *DisassociateEipInvoker {
	requestDef := GenReqDefForDisassociateEip()
	return &DisassociateEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateElb 集群解绑ELB
//
// 集群解绑Elb接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DisassociateElb(request *model.DisassociateElbRequest) (*model.DisassociateElbResponse, error) {
	requestDef := GenReqDefForDisassociateElb()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateElbResponse), nil
	}
}

// DisassociateElbInvoker 集群解绑ELB
func (c *DwsClient) DisassociateElbInvoker(request *model.DisassociateElbRequest) *DisassociateElbInvoker {
	requestDef := GenReqDefForDisassociateElb()
	return &DisassociateElbInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteRedistributionCluster 下发重分布
//
// 下发重分布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ExecuteRedistributionCluster(request *model.ExecuteRedistributionClusterRequest) (*model.ExecuteRedistributionClusterResponse, error) {
	requestDef := GenReqDefForExecuteRedistributionCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteRedistributionClusterResponse), nil
	}
}

// ExecuteRedistributionClusterInvoker 下发重分布
func (c *DwsClient) ExecuteRedistributionClusterInvoker(request *model.ExecuteRedistributionClusterRequest) *ExecuteRedistributionClusterInvoker {
	requestDef := GenReqDefForExecuteRedistributionCluster()
	return &ExecuteRedistributionClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandInstanceStorage 磁盘扩容
//
// 随着客户业务的发展，磁盘空间往往最先出现资源瓶颈，在其他资源尚且充足的情况下，通过磁盘扩容可快速缓解存储资源瓶颈现象，操作过程中无需暂停业务，并且不会造成CPU、内存等资源浪费。
// - 磁盘扩容功能仅8.1.1.203及以上版本支持，并且创建集群规格需要为云数仓SSD云盘或实时数仓类型。
// - 按需+折扣套餐包消费模式下，存储扩容后超出折扣套餐包部分将按需收费。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ExpandInstanceStorage(request *model.ExpandInstanceStorageRequest) (*model.ExpandInstanceStorageResponse, error) {
	requestDef := GenReqDefForExpandInstanceStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandInstanceStorageResponse), nil
	}
}

// ExpandInstanceStorageInvoker 磁盘扩容
func (c *DwsClient) ExpandInstanceStorageInvoker(request *model.ExpandInstanceStorageRequest) *ExpandInstanceStorageInvoker {
	requestDef := GenReqDefForExpandInstanceStorage()
	return &ExpandInstanceStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmConfigs 查询告警配置
//
// 查询告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAlarmConfigs(request *model.ListAlarmConfigsRequest) (*model.ListAlarmConfigsResponse, error) {
	requestDef := GenReqDefForListAlarmConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmConfigsResponse), nil
	}
}

// ListAlarmConfigsInvoker 查询告警配置
func (c *DwsClient) ListAlarmConfigsInvoker(request *model.ListAlarmConfigsRequest) *ListAlarmConfigsInvoker {
	requestDef := GenReqDefForListAlarmConfigs()
	return &ListAlarmConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmDetail 查询告警详情列表
//
// 查询告警详情列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAlarmDetail(request *model.ListAlarmDetailRequest) (*model.ListAlarmDetailResponse, error) {
	requestDef := GenReqDefForListAlarmDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmDetailResponse), nil
	}
}

// ListAlarmDetailInvoker 查询告警详情列表
func (c *DwsClient) ListAlarmDetailInvoker(request *model.ListAlarmDetailRequest) *ListAlarmDetailInvoker {
	requestDef := GenReqDefForListAlarmDetail()
	return &ListAlarmDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmStatistic 查询告警统计列表
//
// 查询告警统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAlarmStatistic(request *model.ListAlarmStatisticRequest) (*model.ListAlarmStatisticResponse, error) {
	requestDef := GenReqDefForListAlarmStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmStatisticResponse), nil
	}
}

// ListAlarmStatisticInvoker 查询告警统计列表
func (c *DwsClient) ListAlarmStatisticInvoker(request *model.ListAlarmStatisticRequest) *ListAlarmStatisticInvoker {
	requestDef := GenReqDefForListAlarmStatistic()
	return &ListAlarmStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmSubs 查询告警订阅列表
//
// 查询订阅告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAlarmSubs(request *model.ListAlarmSubsRequest) (*model.ListAlarmSubsResponse, error) {
	requestDef := GenReqDefForListAlarmSubs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmSubsResponse), nil
	}
}

// ListAlarmSubsInvoker 查询告警订阅列表
func (c *DwsClient) ListAlarmSubsInvoker(request *model.ListAlarmSubsRequest) *ListAlarmSubsInvoker {
	requestDef := GenReqDefForListAlarmSubs()
	return &ListAlarmSubsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditLog 查询日志记录
//
// 查询审计日志记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAuditLog(request *model.ListAuditLogRequest) (*model.ListAuditLogResponse, error) {
	requestDef := GenReqDefForListAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditLogResponse), nil
	}
}

// ListAuditLogInvoker 查询日志记录
func (c *DwsClient) ListAuditLogInvoker(request *model.ListAuditLogRequest) *ListAuditLogInvoker {
	requestDef := GenReqDefForListAuditLog()
	return &ListAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询可用区列表
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询可用区列表
func (c *DwsClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterCn 查询集群CN节点
//
// 查询集群的CN节点列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterCn(request *model.ListClusterCnRequest) (*model.ListClusterCnResponse, error) {
	requestDef := GenReqDefForListClusterCn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterCnResponse), nil
	}
}

// ListClusterCnInvoker 查询集群CN节点
func (c *DwsClient) ListClusterCnInvoker(request *model.ListClusterCnRequest) *ListClusterCnInvoker {
	requestDef := GenReqDefForListClusterCn()
	return &ListClusterCnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterConfigurations 查询集群参数组
//
// 查询集群所关联的参数组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterConfigurations(request *model.ListClusterConfigurationsRequest) (*model.ListClusterConfigurationsResponse, error) {
	requestDef := GenReqDefForListClusterConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterConfigurationsResponse), nil
	}
}

// ListClusterConfigurationsInvoker 查询集群参数组
func (c *DwsClient) ListClusterConfigurationsInvoker(request *model.ListClusterConfigurationsRequest) *ListClusterConfigurationsInvoker {
	requestDef := GenReqDefForListClusterConfigurations()
	return &ListClusterConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterConfigurationsParameter 查询集群参数配置
//
// 查询集群所关联的参数组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterConfigurationsParameter(request *model.ListClusterConfigurationsParameterRequest) (*model.ListClusterConfigurationsParameterResponse, error) {
	requestDef := GenReqDefForListClusterConfigurationsParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterConfigurationsParameterResponse), nil
	}
}

// ListClusterConfigurationsParameterInvoker 查询集群参数配置
func (c *DwsClient) ListClusterConfigurationsParameterInvoker(request *model.ListClusterConfigurationsParameterRequest) *ListClusterConfigurationsParameterInvoker {
	requestDef := GenReqDefForListClusterConfigurationsParameter()
	return &ListClusterConfigurationsParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterDetails 查询集群详情
//
// 该接口用于查询集群详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterDetails(request *model.ListClusterDetailsRequest) (*model.ListClusterDetailsResponse, error) {
	requestDef := GenReqDefForListClusterDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterDetailsResponse), nil
	}
}

// ListClusterDetailsInvoker 查询集群详情
func (c *DwsClient) ListClusterDetailsInvoker(request *model.ListClusterDetailsRequest) *ListClusterDetailsInvoker {
	requestDef := GenReqDefForListClusterDetails()
	return &ListClusterDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterScaleInNumbers 查询合适的缩容数
//
// 查询合适的缩容数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterScaleInNumbers(request *model.ListClusterScaleInNumbersRequest) (*model.ListClusterScaleInNumbersResponse, error) {
	requestDef := GenReqDefForListClusterScaleInNumbers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterScaleInNumbersResponse), nil
	}
}

// ListClusterScaleInNumbersInvoker 查询合适的缩容数
func (c *DwsClient) ListClusterScaleInNumbersInvoker(request *model.ListClusterScaleInNumbersRequest) *ListClusterScaleInNumbersInvoker {
	requestDef := GenReqDefForListClusterScaleInNumbers()
	return &ListClusterScaleInNumbersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterSnapshots 查询集群快照列表
//
// 该接口用于查询集群快照列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterSnapshots(request *model.ListClusterSnapshotsRequest) (*model.ListClusterSnapshotsResponse, error) {
	requestDef := GenReqDefForListClusterSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterSnapshotsResponse), nil
	}
}

// ListClusterSnapshotsInvoker 查询集群快照列表
func (c *DwsClient) ListClusterSnapshotsInvoker(request *model.ListClusterSnapshotsRequest) *ListClusterSnapshotsInvoker {
	requestDef := GenReqDefForListClusterSnapshots()
	return &ListClusterSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterTags 查询集群标签
//
// 查询指定集群的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterTags(request *model.ListClusterTagsRequest) (*model.ListClusterTagsResponse, error) {
	requestDef := GenReqDefForListClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterTagsResponse), nil
	}
}

// ListClusterTagsInvoker 查询集群标签
func (c *DwsClient) ListClusterTagsInvoker(request *model.ListClusterTagsRequest) *ListClusterTagsInvoker {
	requestDef := GenReqDefForListClusterTags()
	return &ListClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterWorkload 查询资源管理
//
// 查询资管管理开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterWorkload(request *model.ListClusterWorkloadRequest) (*model.ListClusterWorkloadResponse, error) {
	requestDef := GenReqDefForListClusterWorkload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterWorkloadResponse), nil
	}
}

// ListClusterWorkloadInvoker 查询资源管理
func (c *DwsClient) ListClusterWorkloadInvoker(request *model.ListClusterWorkloadRequest) *ListClusterWorkloadInvoker {
	requestDef := GenReqDefForListClusterWorkload()
	return &ListClusterWorkloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 查询集群列表
//
// 该接口用于查询并显示集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 查询集群列表
func (c *DwsClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataSource 查询数据源
//
// 该接口用于查询数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListDataSource(request *model.ListDataSourceRequest) (*model.ListDataSourceResponse, error) {
	requestDef := GenReqDefForListDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataSourceResponse), nil
	}
}

// ListDataSourceInvoker 查询数据源
func (c *DwsClient) ListDataSourceInvoker(request *model.ListDataSourceRequest) *ListDataSourceInvoker {
	requestDef := GenReqDefForListDataSource()
	return &ListDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDisasterRecover 查询容灾列表
//
// 查询容灾列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListDisasterRecover(request *model.ListDisasterRecoverRequest) (*model.ListDisasterRecoverResponse, error) {
	requestDef := GenReqDefForListDisasterRecover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDisasterRecoverResponse), nil
	}
}

// ListDisasterRecoverInvoker 查询容灾列表
func (c *DwsClient) ListDisasterRecoverInvoker(request *model.ListDisasterRecoverRequest) *ListDisasterRecoverInvoker {
	requestDef := GenReqDefForListDisasterRecover()
	return &ListDisasterRecoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDssPools 查询专属分布式存储池列表
//
// 获取专属分布式存储池列表，只包括用户开通的SSD专属资源池信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListDssPools(request *model.ListDssPoolsRequest) (*model.ListDssPoolsResponse, error) {
	requestDef := GenReqDefForListDssPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDssPoolsResponse), nil
	}
}

// ListDssPoolsInvoker 查询专属分布式存储池列表
func (c *DwsClient) ListDssPoolsInvoker(request *model.ListDssPoolsRequest) *ListDssPoolsInvoker {
	requestDef := GenReqDefForListDssPools()
	return &ListDssPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListElbs 获取集群可绑定的ELB列表
//
// 查询集群可以关联的Elb列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListElbs(request *model.ListElbsRequest) (*model.ListElbsResponse, error) {
	requestDef := GenReqDefForListElbs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListElbsResponse), nil
	}
}

// ListElbsInvoker 获取集群可绑定的ELB列表
func (c *DwsClient) ListElbsInvoker(request *model.ListElbsRequest) *ListElbsInvoker {
	requestDef := GenReqDefForListElbs()
	return &ListElbsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventSpecs 查询事件配置
//
// 查询事件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListEventSpecs(request *model.ListEventSpecsRequest) (*model.ListEventSpecsResponse, error) {
	requestDef := GenReqDefForListEventSpecs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventSpecsResponse), nil
	}
}

// ListEventSpecsInvoker 查询事件配置
func (c *DwsClient) ListEventSpecsInvoker(request *model.ListEventSpecsRequest) *ListEventSpecsInvoker {
	requestDef := GenReqDefForListEventSpecs()
	return &ListEventSpecsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventSubs 查询订阅事件
//
// 查询订阅的事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListEventSubs(request *model.ListEventSubsRequest) (*model.ListEventSubsResponse, error) {
	requestDef := GenReqDefForListEventSubs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventSubsResponse), nil
	}
}

// ListEventSubsInvoker 查询订阅事件
func (c *DwsClient) ListEventSubsInvoker(request *model.ListEventSubsRequest) *ListEventSubsInvoker {
	requestDef := GenReqDefForListEventSubs()
	return &ListEventSubsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 查询事件列表
//
// 查询事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 查询事件列表
func (c *DwsClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostDisk openApi查询磁盘信息
//
// openApi查询磁盘信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListHostDisk(request *model.ListHostDiskRequest) (*model.ListHostDiskResponse, error) {
	requestDef := GenReqDefForListHostDisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostDiskResponse), nil
	}
}

// ListHostDiskInvoker openApi查询磁盘信息
func (c *DwsClient) ListHostDiskInvoker(request *model.ListHostDiskRequest) *ListHostDiskInvoker {
	requestDef := GenReqDefForListHostDisk()
	return &ListHostDiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostNet openapi获取网卡状态
//
// openapi获取网卡状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListHostNet(request *model.ListHostNetRequest) (*model.ListHostNetResponse, error) {
	requestDef := GenReqDefForListHostNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostNetResponse), nil
	}
}

// ListHostNetInvoker openapi获取网卡状态
func (c *DwsClient) ListHostNetInvoker(request *model.ListHostNetRequest) *ListHostNetInvoker {
	requestDef := GenReqDefForListHostNet()
	return &ListHostNetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostOverview openApi查询主机概览
//
// openApi查询主机概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListHostOverview(request *model.ListHostOverviewRequest) (*model.ListHostOverviewResponse, error) {
	requestDef := GenReqDefForListHostOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostOverviewResponse), nil
	}
}

// ListHostOverviewInvoker openApi查询主机概览
func (c *DwsClient) ListHostOverviewInvoker(request *model.ListHostOverviewRequest) *ListHostOverviewInvoker {
	requestDef := GenReqDefForListHostOverview()
	return &ListHostOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobDetails 查询job进度
//
// 查询job进度信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListJobDetails(request *model.ListJobDetailsRequest) (*model.ListJobDetailsResponse, error) {
	requestDef := GenReqDefForListJobDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobDetailsResponse), nil
	}
}

// ListJobDetailsInvoker 查询job进度
func (c *DwsClient) ListJobDetailsInvoker(request *model.ListJobDetailsRequest) *ListJobDetailsInvoker {
	requestDef := GenReqDefForListJobDetails()
	return &ListJobDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitorIndicatorData openApi查询历史监控数据
//
// openApi查询历史监控数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListMonitorIndicatorData(request *model.ListMonitorIndicatorDataRequest) (*model.ListMonitorIndicatorDataResponse, error) {
	requestDef := GenReqDefForListMonitorIndicatorData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorIndicatorDataResponse), nil
	}
}

// ListMonitorIndicatorDataInvoker openApi查询历史监控数据
func (c *DwsClient) ListMonitorIndicatorDataInvoker(request *model.ListMonitorIndicatorDataRequest) *ListMonitorIndicatorDataInvoker {
	requestDef := GenReqDefForListMonitorIndicatorData()
	return &ListMonitorIndicatorDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitorIndicators openApi查询性能监控指标
//
// openApi查询性能监控指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListMonitorIndicators(request *model.ListMonitorIndicatorsRequest) (*model.ListMonitorIndicatorsResponse, error) {
	requestDef := GenReqDefForListMonitorIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorIndicatorsResponse), nil
	}
}

// ListMonitorIndicatorsInvoker openApi查询性能监控指标
func (c *DwsClient) ListMonitorIndicatorsInvoker(request *model.ListMonitorIndicatorsRequest) *ListMonitorIndicatorsInvoker {
	requestDef := GenReqDefForListMonitorIndicators()
	return &ListMonitorIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodeTypes 查询节点类型
//
// 该接口用于查询所有GaussDB(DWS)服务支持的节点类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListNodeTypes(request *model.ListNodeTypesRequest) (*model.ListNodeTypesResponse, error) {
	requestDef := GenReqDefForListNodeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodeTypesResponse), nil
	}
}

// ListNodeTypesInvoker 查询节点类型
func (c *DwsClient) ListNodeTypesInvoker(request *model.ListNodeTypesRequest) *ListNodeTypesInvoker {
	requestDef := GenReqDefForListNodeTypes()
	return &ListNodeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询单租户在GaussDB(DWS)服务下的配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *DwsClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotDetails 查询快照详情
//
// 该接口用于使用快照ID查询快照详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshotDetails(request *model.ListSnapshotDetailsRequest) (*model.ListSnapshotDetailsResponse, error) {
	requestDef := GenReqDefForListSnapshotDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotDetailsResponse), nil
	}
}

// ListSnapshotDetailsInvoker 查询快照详情
func (c *DwsClient) ListSnapshotDetailsInvoker(request *model.ListSnapshotDetailsRequest) *ListSnapshotDetailsInvoker {
	requestDef := GenReqDefForListSnapshotDetails()
	return &ListSnapshotDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotPolicy 查询快照策略
//
// 查询快照策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshotPolicy(request *model.ListSnapshotPolicyRequest) (*model.ListSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForListSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotPolicyResponse), nil
	}
}

// ListSnapshotPolicyInvoker 查询快照策略
func (c *DwsClient) ListSnapshotPolicyInvoker(request *model.ListSnapshotPolicyRequest) *ListSnapshotPolicyInvoker {
	requestDef := GenReqDefForListSnapshotPolicy()
	return &ListSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotStatistics 快照统计信息
//
// 快照统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshotStatistics(request *model.ListSnapshotStatisticsRequest) (*model.ListSnapshotStatisticsResponse, error) {
	requestDef := GenReqDefForListSnapshotStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotStatisticsResponse), nil
	}
}

// ListSnapshotStatisticsInvoker 快照统计信息
func (c *DwsClient) ListSnapshotStatisticsInvoker(request *model.ListSnapshotStatisticsRequest) *ListSnapshotStatisticsInvoker {
	requestDef := GenReqDefForListSnapshotStatistics()
	return &ListSnapshotStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshots 查询快照列表
//
// 该接口用于查询快照列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// ListSnapshotsInvoker 查询快照列表
func (c *DwsClient) ListSnapshotsInvoker(request *model.ListSnapshotsRequest) *ListSnapshotsInvoker {
	requestDef := GenReqDefForListSnapshots()
	return &ListSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatistics 查询资源统计信息列表
//
// 查询当前可用资源数量，其中包括“可用集群和总集群（个）”、“可用节点和总节点（个）”、“总容量（GB）”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListStatistics(request *model.ListStatisticsRequest) (*model.ListStatisticsResponse, error) {
	requestDef := GenReqDefForListStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsResponse), nil
	}
}

// ListStatisticsInvoker 查询资源统计信息列表
func (c *DwsClient) ListStatisticsInvoker(request *model.ListStatisticsRequest) *ListStatisticsInvoker {
	requestDef := GenReqDefForListStatistics()
	return &ListStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询项目标签
//
// 查询项目标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询项目标签
func (c *DwsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkloadQueue 查询工作负载队列
//
// 查询工作负载队列
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListWorkloadQueue(request *model.ListWorkloadQueueRequest) (*model.ListWorkloadQueueResponse, error) {
	requestDef := GenReqDefForListWorkloadQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkloadQueueResponse), nil
	}
}

// ListWorkloadQueueInvoker 查询工作负载队列
func (c *DwsClient) ListWorkloadQueueInvoker(request *model.ListWorkloadQueueRequest) *ListWorkloadQueueInvoker {
	requestDef := GenReqDefForListWorkloadQueue()
	return &ListWorkloadQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseDisasterRecovery 停止容灾
//
// 停止容灾
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) PauseDisasterRecovery(request *model.PauseDisasterRecoveryRequest) (*model.PauseDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForPauseDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseDisasterRecoveryResponse), nil
	}
}

// PauseDisasterRecoveryInvoker 停止容灾
func (c *DwsClient) PauseDisasterRecoveryInvoker(request *model.PauseDisasterRecoveryRequest) *PauseDisasterRecoveryInvoker {
	requestDef := GenReqDefForPauseDisasterRecovery()
	return &PauseDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 重置密码
//
// 此接口用于重置集群管理员密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 重置密码
func (c *DwsClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeCluster 扩容集群调整集群大小
//
// 此接口用于扩容集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ResizeCluster(request *model.ResizeClusterRequest) (*model.ResizeClusterResponse, error) {
	requestDef := GenReqDefForResizeCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeClusterResponse), nil
	}
}

// ResizeClusterInvoker 扩容集群调整集群大小
func (c *DwsClient) ResizeClusterInvoker(request *model.ResizeClusterRequest) *ResizeClusterInvoker {
	requestDef := GenReqDefForResizeCluster()
	return &ResizeClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCluster 重启集群
//
// 此接口用于重启集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

// RestartClusterInvoker 重启集群
func (c *DwsClient) RestartClusterInvoker(request *model.RestartClusterRequest) *RestartClusterInvoker {
	requestDef := GenReqDefForRestartCluster()
	return &RestartClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreCluster 恢复集群
//
// 该接口用于使用快照恢复集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestoreCluster(request *model.RestoreClusterRequest) (*model.RestoreClusterResponse, error) {
	requestDef := GenReqDefForRestoreCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreClusterResponse), nil
	}
}

// RestoreClusterInvoker 恢复集群
func (c *DwsClient) RestoreClusterInvoker(request *model.RestoreClusterRequest) *RestoreClusterInvoker {
	requestDef := GenReqDefForRestoreCluster()
	return &RestoreClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreDisaster 恢复容灾
//
// 恢复容灾
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestoreDisaster(request *model.RestoreDisasterRequest) (*model.RestoreDisasterResponse, error) {
	requestDef := GenReqDefForRestoreDisaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreDisasterResponse), nil
	}
}

// RestoreDisasterInvoker 恢复容灾
func (c *DwsClient) RestoreDisasterInvoker(request *model.RestoreDisasterRequest) *RestoreDisasterInvoker {
	requestDef := GenReqDefForRestoreDisaster()
	return &RestoreDisasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkCluster 集群缩容
//
// 该接口用于缩容集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShrinkCluster(request *model.ShrinkClusterRequest) (*model.ShrinkClusterResponse, error) {
	requestDef := GenReqDefForShrinkCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkClusterResponse), nil
	}
}

// ShrinkClusterInvoker 集群缩容
func (c *DwsClient) ShrinkClusterInvoker(request *model.ShrinkClusterRequest) *ShrinkClusterInvoker {
	requestDef := GenReqDefForShrinkCluster()
	return &ShrinkClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDisasterRecovery 启动容灾
//
// 启动容灾
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) StartDisasterRecovery(request *model.StartDisasterRecoveryRequest) (*model.StartDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForStartDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDisasterRecoveryResponse), nil
	}
}

// StartDisasterRecoveryInvoker 启动容灾
func (c *DwsClient) StartDisasterRecoveryInvoker(request *model.StartDisasterRecoveryRequest) *StartDisasterRecoveryInvoker {
	requestDef := GenReqDefForStartDisasterRecovery()
	return &StartDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchFailoverDisaster 容灾异常切换
//
// 容灾-异常切换
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SwitchFailoverDisaster(request *model.SwitchFailoverDisasterRequest) (*model.SwitchFailoverDisasterResponse, error) {
	requestDef := GenReqDefForSwitchFailoverDisaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchFailoverDisasterResponse), nil
	}
}

// SwitchFailoverDisasterInvoker 容灾异常切换
func (c *DwsClient) SwitchFailoverDisasterInvoker(request *model.SwitchFailoverDisasterRequest) *SwitchFailoverDisasterInvoker {
	requestDef := GenReqDefForSwitchFailoverDisaster()
	return &SwitchFailoverDisasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchOverCluster 主备恢复
//
// 当集群状态为“非均衡”时会出现某些节点主实例增多，从而负载压力较大。这种情况下集群状态是正常的，但整体性能要低于均衡状态。可进行集群主备恢复操作将集群状态切换为“可用“状态。
// - 集群主备恢复仅8.1.1.202及以上版本支持。
// - 集群主备恢复将会短暂中断业务，中断时间根据用户自身业务量所决定，建议用户在业务低峰期执行此操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SwitchOverCluster(request *model.SwitchOverClusterRequest) (*model.SwitchOverClusterResponse, error) {
	requestDef := GenReqDefForSwitchOverCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchOverClusterResponse), nil
	}
}

// SwitchOverClusterInvoker 主备恢复
func (c *DwsClient) SwitchOverClusterInvoker(request *model.SwitchOverClusterRequest) *SwitchOverClusterInvoker {
	requestDef := GenReqDefForSwitchOverCluster()
	return &SwitchOverClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchoverDisasterRecovery 灾备切换
//
// 容灾-灾备切换
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SwitchoverDisasterRecovery(request *model.SwitchoverDisasterRecoveryRequest) (*model.SwitchoverDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForSwitchoverDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchoverDisasterRecoveryResponse), nil
	}
}

// SwitchoverDisasterRecoveryInvoker 灾备切换
func (c *DwsClient) SwitchoverDisasterRecoveryInvoker(request *model.SwitchoverDisasterRecoveryRequest) *SwitchoverDisasterRecoveryInvoker {
	requestDef := GenReqDefForSwitchoverDisasterRecovery()
	return &SwitchoverDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmSub 更新告警订阅
//
// 更新订阅的告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateAlarmSub(request *model.UpdateAlarmSubRequest) (*model.UpdateAlarmSubResponse, error) {
	requestDef := GenReqDefForUpdateAlarmSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmSubResponse), nil
	}
}

// UpdateAlarmSubInvoker 更新告警订阅
func (c *DwsClient) UpdateAlarmSubInvoker(request *model.UpdateAlarmSubRequest) *UpdateAlarmSubInvoker {
	requestDef := GenReqDefForUpdateAlarmSub()
	return &UpdateAlarmSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterDns 修改集群域名
//
// 为指定集群修改域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateClusterDns(request *model.UpdateClusterDnsRequest) (*model.UpdateClusterDnsResponse, error) {
	requestDef := GenReqDefForUpdateClusterDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterDnsResponse), nil
	}
}

// UpdateClusterDnsInvoker 修改集群域名
func (c *DwsClient) UpdateClusterDnsInvoker(request *model.UpdateClusterDnsRequest) *UpdateClusterDnsInvoker {
	requestDef := GenReqDefForUpdateClusterDns()
	return &UpdateClusterDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfiguration 修改集群参数配置
//
// 修改集群使用的参数配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateConfiguration(request *model.UpdateConfigurationRequest) (*model.UpdateConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationResponse), nil
	}
}

// UpdateConfigurationInvoker 修改集群参数配置
func (c *DwsClient) UpdateConfigurationInvoker(request *model.UpdateConfigurationRequest) *UpdateConfigurationInvoker {
	requestDef := GenReqDefForUpdateConfiguration()
	return &UpdateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataSource 更新数据源
//
// 该接口用于更新一个数据源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateDataSource(request *model.UpdateDataSourceRequest) (*model.UpdateDataSourceResponse, error) {
	requestDef := GenReqDefForUpdateDataSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataSourceResponse), nil
	}
}

// UpdateDataSourceInvoker 更新数据源
func (c *DwsClient) UpdateDataSourceInvoker(request *model.UpdateDataSourceRequest) *UpdateDataSourceInvoker {
	requestDef := GenReqDefForUpdateDataSource()
	return &UpdateDataSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEventSub 更新订阅事件
//
// 更新订阅事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateEventSub(request *model.UpdateEventSubRequest) (*model.UpdateEventSubResponse, error) {
	requestDef := GenReqDefForUpdateEventSub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventSubResponse), nil
	}
}

// UpdateEventSubInvoker 更新订阅事件
func (c *DwsClient) UpdateEventSubInvoker(request *model.UpdateEventSubRequest) *UpdateEventSubInvoker {
	requestDef := GenReqDefForUpdateEventSub()
	return &UpdateEventSubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMaintenanceWindow 修改运维时间窗
//
// 您可以根据业务需求，设置可维护时间段。建议将可维护时间段设置在业务低峰期，避免业务在维护过程中异常中断。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateMaintenanceWindow(request *model.UpdateMaintenanceWindowRequest) (*model.UpdateMaintenanceWindowResponse, error) {
	requestDef := GenReqDefForUpdateMaintenanceWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMaintenanceWindowResponse), nil
	}
}

// UpdateMaintenanceWindowInvoker 修改运维时间窗
func (c *DwsClient) UpdateMaintenanceWindowInvoker(request *model.UpdateMaintenanceWindowRequest) *UpdateMaintenanceWindowInvoker {
	requestDef := GenReqDefForUpdateMaintenanceWindow()
	return &UpdateMaintenanceWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
