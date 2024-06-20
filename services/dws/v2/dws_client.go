package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type DwsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDwsClient(hcClient *httpclient.HcHttpClient) *DwsClient {
	return &DwsClient{HcClient: hcClient}
}

func DwsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddQueueUserList 添加工作负载队列的绑定用户
//
// 添加工作负载队列的绑定用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AddQueueUserList(request *model.AddQueueUserListRequest) (*model.AddQueueUserListResponse, error) {
	requestDef := GenReqDefForAddQueueUserList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddQueueUserListResponse), nil
	}
}

// AddQueueUserListInvoker 添加工作负载队列的绑定用户
func (c *DwsClient) AddQueueUserListInvoker(request *model.AddQueueUserListRequest) *AddQueueUserListInvoker {
	requestDef := GenReqDefForAddQueueUserList()
	return &AddQueueUserListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSnapshotCrossRegionPolicy 设置跨区域备份配置
//
// 该接口用于设置跨区域备份配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AddSnapshotCrossRegionPolicy(request *model.AddSnapshotCrossRegionPolicyRequest) (*model.AddSnapshotCrossRegionPolicyResponse, error) {
	requestDef := GenReqDefForAddSnapshotCrossRegionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSnapshotCrossRegionPolicyResponse), nil
	}
}

// AddSnapshotCrossRegionPolicyInvoker 设置跨区域备份配置
func (c *DwsClient) AddSnapshotCrossRegionPolicyInvoker(request *model.AddSnapshotCrossRegionPolicyRequest) *AddSnapshotCrossRegionPolicyInvoker {
	requestDef := GenReqDefForAddSnapshotCrossRegionPolicy()
	return &AddSnapshotCrossRegionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddWorkloadPlanStage 添加工作负载计划阶段
//
// 添加工作负载计划阶段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) AddWorkloadPlanStage(request *model.AddWorkloadPlanStageRequest) (*model.AddWorkloadPlanStageResponse, error) {
	requestDef := GenReqDefForAddWorkloadPlanStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddWorkloadPlanStageResponse), nil
	}
}

// AddWorkloadPlanStageInvoker 添加工作负载计划阶段
func (c *DwsClient) AddWorkloadPlanStageInvoker(request *model.AddWorkloadPlanStageRequest) *AddWorkloadPlanStageInvoker {
	requestDef := GenReqDefForAddWorkloadPlanStage()
	return &AddWorkloadPlanStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ChangeSecurityGroup 修改集群安全组
//
// 该接口用于修改集群安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

// ChangeSecurityGroupInvoker 修改集群安全组
func (c *DwsClient) ChangeSecurityGroupInvoker(request *model.ChangeSecurityGroupRequest) *ChangeSecurityGroupInvoker {
	requestDef := GenReqDefForChangeSecurityGroup()
	return &ChangeSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CheckDisasterName 检查容灾名称
//
// 该接口用于查询容灾名称是否可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CheckDisasterName(request *model.CheckDisasterNameRequest) (*model.CheckDisasterNameResponse, error) {
	requestDef := GenReqDefForCheckDisasterName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDisasterNameResponse), nil
	}
}

// CheckDisasterNameInvoker 检查容灾名称
func (c *DwsClient) CheckDisasterNameInvoker(request *model.CheckDisasterNameRequest) *CheckDisasterNameInvoker {
	requestDef := GenReqDefForCheckDisasterName()
	return &CheckDisasterNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckTableRestore 用户恢复表名检测
//
// 该接口用于用户恢复表名检测
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CheckTableRestore(request *model.CheckTableRestoreRequest) (*model.CheckTableRestoreResponse, error) {
	requestDef := GenReqDefForCheckTableRestore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTableRestoreResponse), nil
	}
}

// CheckTableRestoreInvoker 用户恢复表名检测
func (c *DwsClient) CheckTableRestoreInvoker(request *model.CheckTableRestoreRequest) *CheckTableRestoreInvoker {
	requestDef := GenReqDefForCheckTableRestore()
	return &CheckTableRestoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConvertToLogicalCluster 物理集群转换到逻辑集群
//
// 逻辑集群-物理集群转换到逻辑集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ConvertToLogicalCluster(request *model.ConvertToLogicalClusterRequest) (*model.ConvertToLogicalClusterResponse, error) {
	requestDef := GenReqDefForConvertToLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConvertToLogicalClusterResponse), nil
	}
}

// ConvertToLogicalClusterInvoker 物理集群转换到逻辑集群
func (c *DwsClient) ConvertToLogicalClusterInvoker(request *model.ConvertToLogicalClusterRequest) *ConvertToLogicalClusterInvoker {
	requestDef := GenReqDefForConvertToLogicalCluster()
	return &ConvertToLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateClusterV2 V2创建集群
//
// 该接口用于创建集群。
// 集群必须要运行在VPC之内，创建集群前，您需要先创建VPC，并获取VPC和子网的id。
// 该接口为异步接口，创建集群需要10～15分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateClusterV2(request *model.CreateClusterV2Request) (*model.CreateClusterV2Response, error) {
	requestDef := GenReqDefForCreateClusterV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterV2Response), nil
	}
}

// CreateClusterV2Invoker V2创建集群
func (c *DwsClient) CreateClusterV2Invoker(request *model.CreateClusterV2Request) *CreateClusterV2Invoker {
	requestDef := GenReqDefForCreateClusterV2()
	return &CreateClusterV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 该接口用于创建集群间容灾。
// 集群处于可用状态或者非均衡状态才可进行创建容灾操作。
// 仅支持DWS 2.0集群。
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

// CreateLogicalCluster 创建逻辑集群
//
// 逻辑集群-创建逻辑集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) CreateLogicalCluster(request *model.CreateLogicalClusterRequest) (*model.CreateLogicalClusterResponse, error) {
	requestDef := GenReqDefForCreateLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogicalClusterResponse), nil
	}
}

// CreateLogicalClusterInvoker 创建逻辑集群
func (c *DwsClient) CreateLogicalClusterInvoker(request *model.CreateLogicalClusterRequest) *CreateLogicalClusterInvoker {
	requestDef := GenReqDefForCreateLogicalCluster()
	return &CreateLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteClusterNodes 删除节点
//
// 此接口用于删除节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteClusterNodes(request *model.DeleteClusterNodesRequest) (*model.DeleteClusterNodesResponse, error) {
	requestDef := GenReqDefForDeleteClusterNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterNodesResponse), nil
	}
}

// DeleteClusterNodesInvoker 删除节点
func (c *DwsClient) DeleteClusterNodesInvoker(request *model.DeleteClusterNodesRequest) *DeleteClusterNodesInvoker {
	requestDef := GenReqDefForDeleteClusterNodes()
	return &DeleteClusterNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 该接口用于删除容灾操作。
// 容灾状态为“创建失败”、“未启动”、“启动失败”、“已停止”、“停止失败”和“异常”时可以执行删除容灾操作。
// 删除后，将无法进行数据同步，且不可恢复，请谨慎操作。
// 仅支持DWS 2.0集群。
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

// DeleteLogicalCluster 删除逻辑集群
//
// 此接口用于删除逻辑集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteLogicalCluster(request *model.DeleteLogicalClusterRequest) (*model.DeleteLogicalClusterResponse, error) {
	requestDef := GenReqDefForDeleteLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogicalClusterResponse), nil
	}
}

// DeleteLogicalClusterInvoker 删除逻辑集群
func (c *DwsClient) DeleteLogicalClusterInvoker(request *model.DeleteLogicalClusterRequest) *DeleteLogicalClusterInvoker {
	requestDef := GenReqDefForDeleteLogicalCluster()
	return &DeleteLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQueueUserList 删除工作负载队列的绑定用户
//
// 删除工作负载队列的绑定用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteQueueUserList(request *model.DeleteQueueUserListRequest) (*model.DeleteQueueUserListResponse, error) {
	requestDef := GenReqDefForDeleteQueueUserList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueueUserListResponse), nil
	}
}

// DeleteQueueUserListInvoker 删除工作负载队列的绑定用户
func (c *DwsClient) DeleteQueueUserListInvoker(request *model.DeleteQueueUserListRequest) *DeleteQueueUserListInvoker {
	requestDef := GenReqDefForDeleteQueueUserList()
	return &DeleteQueueUserListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteSnapshotCrossRegionPolicy 删除跨区域备份配置
//
// 该接口用于删除跨区域备份配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteSnapshotCrossRegionPolicy(request *model.DeleteSnapshotCrossRegionPolicyRequest) (*model.DeleteSnapshotCrossRegionPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSnapshotCrossRegionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotCrossRegionPolicyResponse), nil
	}
}

// DeleteSnapshotCrossRegionPolicyInvoker 删除跨区域备份配置
func (c *DwsClient) DeleteSnapshotCrossRegionPolicyInvoker(request *model.DeleteSnapshotCrossRegionPolicyRequest) *DeleteSnapshotCrossRegionPolicyInvoker {
	requestDef := GenReqDefForDeleteSnapshotCrossRegionPolicy()
	return &DeleteSnapshotCrossRegionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteWorkloadPlan 删除工作负载计划
//
// 删除工作负载计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteWorkloadPlan(request *model.DeleteWorkloadPlanRequest) (*model.DeleteWorkloadPlanResponse, error) {
	requestDef := GenReqDefForDeleteWorkloadPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkloadPlanResponse), nil
	}
}

// DeleteWorkloadPlanInvoker 删除工作负载计划
func (c *DwsClient) DeleteWorkloadPlanInvoker(request *model.DeleteWorkloadPlanRequest) *DeleteWorkloadPlanInvoker {
	requestDef := GenReqDefForDeleteWorkloadPlan()
	return &DeleteWorkloadPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkloadPlanStage 删除工作负载计划阶段
//
// 删除工作负载计划删除工作负载计划阶段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DeleteWorkloadPlanStage(request *model.DeleteWorkloadPlanStageRequest) (*model.DeleteWorkloadPlanStageResponse, error) {
	requestDef := GenReqDefForDeleteWorkloadPlanStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkloadPlanStageResponse), nil
	}
}

// DeleteWorkloadPlanStageInvoker 删除工作负载计划阶段
func (c *DwsClient) DeleteWorkloadPlanStageInvoker(request *model.DeleteWorkloadPlanStageRequest) *DeleteWorkloadPlanStageInvoker {
	requestDef := GenReqDefForDeleteWorkloadPlanStage()
	return &DeleteWorkloadPlanStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkloadQueue 删除资源池
//
// 该接口用于删除资源池。
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

// DeleteWorkloadQueueInvoker 删除资源池
func (c *DwsClient) DeleteWorkloadQueueInvoker(request *model.DeleteWorkloadQueueRequest) *DeleteWorkloadQueueInvoker {
	requestDef := GenReqDefForDeleteWorkloadQueue()
	return &DeleteWorkloadQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableLtsLogs 关闭云服务日志
//
// 该接口用于关闭集群LTS云日志服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) DisableLtsLogs(request *model.DisableLtsLogsRequest) (*model.DisableLtsLogsResponse, error) {
	requestDef := GenReqDefForDisableLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableLtsLogsResponse), nil
	}
}

// DisableLtsLogsInvoker 关闭云服务日志
func (c *DwsClient) DisableLtsLogsInvoker(request *model.DisableLtsLogsRequest) *DisableLtsLogsInvoker {
	requestDef := GenReqDefForDisableLtsLogs()
	return &DisableLtsLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// EnableLogicalCluster 切换逻辑集群开关
//
// 此接口用于切换逻辑集群开关，仅用于控制逻辑集群相关功能模块是否在页面展示。在集群已经是逻辑集群的场景下，修改该接口无任何作用及影响。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) EnableLogicalCluster(request *model.EnableLogicalClusterRequest) (*model.EnableLogicalClusterResponse, error) {
	requestDef := GenReqDefForEnableLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLogicalClusterResponse), nil
	}
}

// EnableLogicalClusterInvoker 切换逻辑集群开关
func (c *DwsClient) EnableLogicalClusterInvoker(request *model.EnableLogicalClusterRequest) *EnableLogicalClusterInvoker {
	requestDef := GenReqDefForEnableLogicalCluster()
	return &EnableLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableLtsLogs 开启云服务日志
//
// 该接口用于开启集群LTS云日志服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) EnableLtsLogs(request *model.EnableLtsLogsRequest) (*model.EnableLtsLogsResponse, error) {
	requestDef := GenReqDefForEnableLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLtsLogsResponse), nil
	}
}

// EnableLtsLogsInvoker 开启云服务日志
func (c *DwsClient) EnableLtsLogsInvoker(request *model.EnableLtsLogsRequest) *EnableLtsLogsInvoker {
	requestDef := GenReqDefForEnableLtsLogs()
	return &EnableLtsLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteClusterUpgradeAction 下发集群升级相关操作
//
// 下发集群升级相关操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ExecuteClusterUpgradeAction(request *model.ExecuteClusterUpgradeActionRequest) (*model.ExecuteClusterUpgradeActionResponse, error) {
	requestDef := GenReqDefForExecuteClusterUpgradeAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteClusterUpgradeActionResponse), nil
	}
}

// ExecuteClusterUpgradeActionInvoker 下发集群升级相关操作
func (c *DwsClient) ExecuteClusterUpgradeActionInvoker(request *model.ExecuteClusterUpgradeActionRequest) *ExecuteClusterUpgradeActionInvoker {
	requestDef := GenReqDefForExecuteClusterUpgradeAction()
	return &ExecuteClusterUpgradeActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDatabaseOmUserAction 执行运维用户操作
//
// 进行数据库运维账户操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ExecuteDatabaseOmUserAction(request *model.ExecuteDatabaseOmUserActionRequest) (*model.ExecuteDatabaseOmUserActionResponse, error) {
	requestDef := GenReqDefForExecuteDatabaseOmUserAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDatabaseOmUserActionResponse), nil
	}
}

// ExecuteDatabaseOmUserActionInvoker 执行运维用户操作
func (c *DwsClient) ExecuteDatabaseOmUserActionInvoker(request *model.ExecuteDatabaseOmUserActionRequest) *ExecuteDatabaseOmUserActionInvoker {
	requestDef := GenReqDefForExecuteDatabaseOmUserAction()
	return &ExecuteDatabaseOmUserActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteRedistributionCluster 下发重分布
//
// 该接口用于集群扩容后将老节点数据均匀分布到新扩节点的数据重分布操作，数据“重分布”后将大大提升业务响应速率。
// 重分布功能DWS 2.0 8.1.1.200及以上集群版本支持。
// 离线调度重分布模式在8.2.0及以上版本将不再支持。
// 只有在扩容之后，集群任务信息为“待重分布”状态时才能手动使用“重分布”功能，其他时段该功能不可使用。
// 在扩容阶段也可以选择重分布模式等高级配置，详情参见设置高级配置。
// 重分布队列的排序依据表的relpage大小进行，为确保relpage大小正确，建议在重分布之前对需要重分布的表执行analyze操作。
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

// ListAvailableDisasterClusters 查询可用容灾集群列表
//
// 该接口用于查询可用的容灾集群列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListAvailableDisasterClusters(request *model.ListAvailableDisasterClustersRequest) (*model.ListAvailableDisasterClustersResponse, error) {
	requestDef := GenReqDefForListAvailableDisasterClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableDisasterClustersResponse), nil
	}
}

// ListAvailableDisasterClustersInvoker 查询可用容灾集群列表
func (c *DwsClient) ListAvailableDisasterClustersInvoker(request *model.ListAvailableDisasterClustersRequest) *ListAvailableDisasterClustersInvoker {
	requestDef := GenReqDefForListAvailableDisasterClusters()
	return &ListAvailableDisasterClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListClusterNodes 查询节点列表
//
// 查询节点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListClusterNodes(request *model.ListClusterNodesRequest) (*model.ListClusterNodesResponse, error) {
	requestDef := GenReqDefForListClusterNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterNodesResponse), nil
	}
}

// ListClusterNodesInvoker 查询节点列表
func (c *DwsClient) ListClusterNodesInvoker(request *model.ListClusterNodesRequest) *ListClusterNodesInvoker {
	requestDef := GenReqDefForListClusterNodes()
	return &ListClusterNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListConfigurationsAuditRecords 查询参数修改审计记录
//
// 查询参数修改审计记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListConfigurationsAuditRecords(request *model.ListConfigurationsAuditRecordsRequest) (*model.ListConfigurationsAuditRecordsResponse, error) {
	requestDef := GenReqDefForListConfigurationsAuditRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsAuditRecordsResponse), nil
	}
}

// ListConfigurationsAuditRecordsInvoker 查询参数修改审计记录
func (c *DwsClient) ListConfigurationsAuditRecordsInvoker(request *model.ListConfigurationsAuditRecordsRequest) *ListConfigurationsAuditRecordsInvoker {
	requestDef := GenReqDefForListConfigurationsAuditRecords()
	return &ListConfigurationsAuditRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListDatabaseUsers 查询所有数据库用户/角色
//
// 查询所有数据库用户/角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListDatabaseUsers(request *model.ListDatabaseUsersRequest) (*model.ListDatabaseUsersResponse, error) {
	requestDef := GenReqDefForListDatabaseUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseUsersResponse), nil
	}
}

// ListDatabaseUsersInvoker 查询所有数据库用户/角色
func (c *DwsClient) ListDatabaseUsersInvoker(request *model.ListDatabaseUsersRequest) *ListDatabaseUsersInvoker {
	requestDef := GenReqDefForListDatabaseUsers()
	return &ListDatabaseUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDisasterRecover 查询容灾列表
//
// 该接口用于查询容灾列表。
// 仅支持DWS 2.0集群。
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

// ListLogicalClusterRings 查询逻辑集群可用ring环节点信息
//
// 查询逻辑集群可用ring环节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListLogicalClusterRings(request *model.ListLogicalClusterRingsRequest) (*model.ListLogicalClusterRingsResponse, error) {
	requestDef := GenReqDefForListLogicalClusterRings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogicalClusterRingsResponse), nil
	}
}

// ListLogicalClusterRingsInvoker 查询逻辑集群可用ring环节点信息
func (c *DwsClient) ListLogicalClusterRingsInvoker(request *model.ListLogicalClusterRingsRequest) *ListLogicalClusterRingsInvoker {
	requestDef := GenReqDefForListLogicalClusterRings()
	return &ListLogicalClusterRingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogicalClusterTasks 查询逻辑集群任务信息
//
// 逻辑集群-查询逻辑集群任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListLogicalClusterTasks(request *model.ListLogicalClusterTasksRequest) (*model.ListLogicalClusterTasksResponse, error) {
	requestDef := GenReqDefForListLogicalClusterTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogicalClusterTasksResponse), nil
	}
}

// ListLogicalClusterTasksInvoker 查询逻辑集群任务信息
func (c *DwsClient) ListLogicalClusterTasksInvoker(request *model.ListLogicalClusterTasksRequest) *ListLogicalClusterTasksInvoker {
	requestDef := GenReqDefForListLogicalClusterTasks()
	return &ListLogicalClusterTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogicalClusterVolumes 查询逻辑集群磁盘信息
//
// 逻辑集群-查询逻辑集群磁盘信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListLogicalClusterVolumes(request *model.ListLogicalClusterVolumesRequest) (*model.ListLogicalClusterVolumesResponse, error) {
	requestDef := GenReqDefForListLogicalClusterVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogicalClusterVolumesResponse), nil
	}
}

// ListLogicalClusterVolumesInvoker 查询逻辑集群磁盘信息
func (c *DwsClient) ListLogicalClusterVolumesInvoker(request *model.ListLogicalClusterVolumesRequest) *ListLogicalClusterVolumesInvoker {
	requestDef := GenReqDefForListLogicalClusterVolumes()
	return &ListLogicalClusterVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogicalClusters 查询逻辑集群列表
//
// 逻辑集群-查询逻辑集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListLogicalClusters(request *model.ListLogicalClustersRequest) (*model.ListLogicalClustersResponse, error) {
	requestDef := GenReqDefForListLogicalClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogicalClustersResponse), nil
	}
}

// ListLogicalClustersInvoker 查询逻辑集群列表
func (c *DwsClient) ListLogicalClustersInvoker(request *model.ListLogicalClustersRequest) *ListLogicalClustersInvoker {
	requestDef := GenReqDefForListLogicalClusters()
	return &ListLogicalClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsLogs 获取LTS日志列表
//
// 获取LTS日志列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListLtsLogs(request *model.ListLtsLogsRequest) (*model.ListLtsLogsResponse, error) {
	requestDef := GenReqDefForListLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsLogsResponse), nil
	}
}

// ListLtsLogsInvoker 获取LTS日志列表
func (c *DwsClient) ListLtsLogsInvoker(request *model.ListLtsLogsRequest) *ListLtsLogsInvoker {
	requestDef := GenReqDefForListLtsLogs()
	return &ListLtsLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 查询集群使用指标列表
//
// 查询集群使用指标列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 查询集群使用指标列表
func (c *DwsClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricsData 获取指定指标相关采集数据
//
// 获取指定指标相关采集数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListMetricsData(request *model.ListMetricsDataRequest) (*model.ListMetricsDataResponse, error) {
	requestDef := GenReqDefForListMetricsData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsDataResponse), nil
	}
}

// ListMetricsDataInvoker 获取指定指标相关采集数据
func (c *DwsClient) ListMetricsDataInvoker(request *model.ListMetricsDataRequest) *ListMetricsDataInvoker {
	requestDef := GenReqDefForListMetricsData()
	return &ListMetricsDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPlanExecLogs 查看计划执行日志
//
// 查看计划执行日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListPlanExecLogs(request *model.ListPlanExecLogsRequest) (*model.ListPlanExecLogsResponse, error) {
	requestDef := GenReqDefForListPlanExecLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlanExecLogsResponse), nil
	}
}

// ListPlanExecLogsInvoker 查看计划执行日志
func (c *DwsClient) ListPlanExecLogsInvoker(request *model.ListPlanExecLogsRequest) *ListPlanExecLogsInvoker {
	requestDef := GenReqDefForListPlanExecLogs()
	return &ListPlanExecLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueries 查询SQL列表
//
// 该接口用于查询实时SQL列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListQueries(request *model.ListQueriesRequest) (*model.ListQueriesResponse, error) {
	requestDef := GenReqDefForListQueries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueriesResponse), nil
	}
}

// ListQueriesInvoker 查询SQL列表
func (c *DwsClient) ListQueriesInvoker(request *model.ListQueriesRequest) *ListQueriesInvoker {
	requestDef := GenReqDefForListQueries()
	return &ListQueriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListSchemas 查询集群模式空间信息
//
// 查询集群模式空间信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSchemas(request *model.ListSchemasRequest) (*model.ListSchemasResponse, error) {
	requestDef := GenReqDefForListSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemasResponse), nil
	}
}

// ListSchemasInvoker 查询集群模式空间信息
func (c *DwsClient) ListSchemasInvoker(request *model.ListSchemasRequest) *ListSchemasInvoker {
	requestDef := GenReqDefForListSchemas()
	return &ListSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotCrossRegion 获取跨区域快照可用region
//
// 该接口用于获取跨区域快照可用region
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshotCrossRegion(request *model.ListSnapshotCrossRegionRequest) (*model.ListSnapshotCrossRegionResponse, error) {
	requestDef := GenReqDefForListSnapshotCrossRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotCrossRegionResponse), nil
	}
}

// ListSnapshotCrossRegionInvoker 获取跨区域快照可用region
func (c *DwsClient) ListSnapshotCrossRegionInvoker(request *model.ListSnapshotCrossRegionRequest) *ListSnapshotCrossRegionInvoker {
	requestDef := GenReqDefForListSnapshotCrossRegion()
	return &ListSnapshotCrossRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotCrossRegionPolicy 查询所有跨区域快照配置
//
// 该接口用于查询所有跨区域快照配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListSnapshotCrossRegionPolicy(request *model.ListSnapshotCrossRegionPolicyRequest) (*model.ListSnapshotCrossRegionPolicyResponse, error) {
	requestDef := GenReqDefForListSnapshotCrossRegionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotCrossRegionPolicyResponse), nil
	}
}

// ListSnapshotCrossRegionPolicyInvoker 查询所有跨区域快照配置
func (c *DwsClient) ListSnapshotCrossRegionPolicyInvoker(request *model.ListSnapshotCrossRegionPolicyRequest) *ListSnapshotCrossRegionPolicyInvoker {
	requestDef := GenReqDefForListSnapshotCrossRegionPolicy()
	return &ListSnapshotCrossRegionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTablesStatistic 查询表倾斜或脏页率信息
//
// 该接口用于查询表倾斜或脏页率信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListTablesStatistic(request *model.ListTablesStatisticRequest) (*model.ListTablesStatisticResponse, error) {
	requestDef := GenReqDefForListTablesStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesStatisticResponse), nil
	}
}

// ListTablesStatisticInvoker 查询表倾斜或脏页率信息
func (c *DwsClient) ListTablesStatisticInvoker(request *model.ListTablesStatisticRequest) *ListTablesStatisticInvoker {
	requestDef := GenReqDefForListTablesStatistic()
	return &ListTablesStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTopoRings 查询集群拓扑ring环节点信息
//
// 查询集群拓扑ring环节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListTopoRings(request *model.ListTopoRingsRequest) (*model.ListTopoRingsResponse, error) {
	requestDef := GenReqDefForListTopoRings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopoRingsResponse), nil
	}
}

// ListTopoRingsInvoker 查询集群拓扑ring环节点信息
func (c *DwsClient) ListTopoRingsInvoker(request *model.ListTopoRingsRequest) *ListTopoRingsInvoker {
	requestDef := GenReqDefForListTopoRings()
	return &ListTopoRingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpdatableVersion 获取集群可升级的目标版本
//
// 获取集群可升级的目标版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListUpdatableVersion(request *model.ListUpdatableVersionRequest) (*model.ListUpdatableVersionResponse, error) {
	requestDef := GenReqDefForListUpdatableVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpdatableVersionResponse), nil
	}
}

// ListUpdatableVersionInvoker 获取集群可升级的目标版本
func (c *DwsClient) ListUpdatableVersionInvoker(request *model.ListUpdatableVersionRequest) *ListUpdatableVersionInvoker {
	requestDef := GenReqDefForListUpdatableVersion()
	return &ListUpdatableVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpdateRecord 获取集群升级记录
//
// 通过此api获取当前集群升级记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListUpdateRecord(request *model.ListUpdateRecordRequest) (*model.ListUpdateRecordResponse, error) {
	requestDef := GenReqDefForListUpdateRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpdateRecordResponse), nil
	}
}

// ListUpdateRecordInvoker 获取集群升级记录
func (c *DwsClient) ListUpdateRecordInvoker(request *model.ListUpdateRecordRequest) *ListUpdateRecordInvoker {
	requestDef := GenReqDefForListUpdateRecord()
	return &ListUpdateRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkloadPlans 查询资源管理计划列表
//
// 查询集群中所有资源管理计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListWorkloadPlans(request *model.ListWorkloadPlansRequest) (*model.ListWorkloadPlansResponse, error) {
	requestDef := GenReqDefForListWorkloadPlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkloadPlansResponse), nil
	}
}

// ListWorkloadPlansInvoker 查询资源管理计划列表
func (c *DwsClient) ListWorkloadPlansInvoker(request *model.ListWorkloadPlansRequest) *ListWorkloadPlansInvoker {
	requestDef := GenReqDefForListWorkloadPlans()
	return &ListWorkloadPlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListWorkloadQueueUsers 获得工作负载队列的绑定用户列表
//
// 获得工作负载队列的绑定用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ListWorkloadQueueUsers(request *model.ListWorkloadQueueUsersRequest) (*model.ListWorkloadQueueUsersResponse, error) {
	requestDef := GenReqDefForListWorkloadQueueUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkloadQueueUsersResponse), nil
	}
}

// ListWorkloadQueueUsersInvoker 获得工作负载队列的绑定用户列表
func (c *DwsClient) ListWorkloadQueueUsersInvoker(request *model.ListWorkloadQueueUsersRequest) *ListWorkloadQueueUsersInvoker {
	requestDef := GenReqDefForListWorkloadQueueUsers()
	return &ListWorkloadQueueUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseDisasterRecovery 停止容灾
//
// 该接口用于停止容灾操作。
// 容灾状态为“运行中”和“停止失败”时可以执行停止容灾操作。
// 停止后，将无法进行数据同步，请谨慎操作。
// 仅支持DWS 2.0集群。
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

// ResizeClusterWithExistedNodes 从空闲节点扩容
//
// 此接口用于从空闲节点扩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ResizeClusterWithExistedNodes(request *model.ResizeClusterWithExistedNodesRequest) (*model.ResizeClusterWithExistedNodesResponse, error) {
	requestDef := GenReqDefForResizeClusterWithExistedNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeClusterWithExistedNodesResponse), nil
	}
}

// ResizeClusterWithExistedNodesInvoker 从空闲节点扩容
func (c *DwsClient) ResizeClusterWithExistedNodesInvoker(request *model.ResizeClusterWithExistedNodesRequest) *ResizeClusterWithExistedNodesInvoker {
	requestDef := GenReqDefForResizeClusterWithExistedNodes()
	return &ResizeClusterWithExistedNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// RestartLogicalCluster 重启逻辑集群
//
// 此接口用于重启逻辑集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestartLogicalCluster(request *model.RestartLogicalClusterRequest) (*model.RestartLogicalClusterResponse, error) {
	requestDef := GenReqDefForRestartLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartLogicalClusterResponse), nil
	}
}

// RestartLogicalClusterInvoker 重启逻辑集群
func (c *DwsClient) RestartLogicalClusterInvoker(request *model.RestartLogicalClusterRequest) *RestartLogicalClusterInvoker {
	requestDef := GenReqDefForRestartLogicalCluster()
	return &RestartLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 该接口用于主备集群进行异常切换，备集群恢复可用状态后进行的容灾恢复操作。
// 容灾恢复仅8.1.2及以上集群版本支持。
// 容灾恢复会删除灾备集群数据与新生产集群重新建立容灾关系。
// 仅支持DWS 2.0集群。
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

// RestoreRedistribution 恢复重分布
//
// 此接口用于恢复暂停状态下的重分布操作,仅支持DWS2.0集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestoreRedistribution(request *model.RestoreRedistributionRequest) (*model.RestoreRedistributionResponse, error) {
	requestDef := GenReqDefForRestoreRedistribution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreRedistributionResponse), nil
	}
}

// RestoreRedistributionInvoker 恢复重分布
func (c *DwsClient) RestoreRedistributionInvoker(request *model.RestoreRedistributionRequest) *RestoreRedistributionInvoker {
	requestDef := GenReqDefForRestoreRedistribution()
	return &RestoreRedistributionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreTable 恢复表
//
// 该接口用于恢复表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) RestoreTable(request *model.RestoreTableRequest) (*model.RestoreTableResponse, error) {
	requestDef := GenReqDefForRestoreTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreTableResponse), nil
	}
}

// RestoreTableInvoker 恢复表
func (c *DwsClient) RestoreTableInvoker(request *model.RestoreTableRequest) *RestoreTableInvoker {
	requestDef := GenReqDefForRestoreTable()
	return &RestoreTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveClusterDescriptionInfo 保存集群描述信息
//
// 保存集群描述信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SaveClusterDescriptionInfo(request *model.SaveClusterDescriptionInfoRequest) (*model.SaveClusterDescriptionInfoResponse, error) {
	requestDef := GenReqDefForSaveClusterDescriptionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveClusterDescriptionInfoResponse), nil
	}
}

// SaveClusterDescriptionInfoInvoker 保存集群描述信息
func (c *DwsClient) SaveClusterDescriptionInfoInvoker(request *model.SaveClusterDescriptionInfoRequest) *SaveClusterDescriptionInfoInvoker {
	requestDef := GenReqDefForSaveClusterDescriptionInfo()
	return &SaveClusterDescriptionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterFlavor 查询集群规格详情
//
// 查询集群使用的规格详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowClusterFlavor(request *model.ShowClusterFlavorRequest) (*model.ShowClusterFlavorResponse, error) {
	requestDef := GenReqDefForShowClusterFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterFlavorResponse), nil
	}
}

// ShowClusterFlavorInvoker 查询集群规格详情
func (c *DwsClient) ShowClusterFlavorInvoker(request *model.ShowClusterFlavorRequest) *ShowClusterFlavorInvoker {
	requestDef := GenReqDefForShowClusterFlavor()
	return &ShowClusterFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterRedistribution 查询重分布详情
//
// 该接口用于查看当前集群的重分布模式、重分布进度、数据表重分布详情等监控信息。
// 查看重分布详情功能DWS 2.0 8.1.1.200及以上集群版本支持，其中数据表重分布进度详情仅DWS 2.0 8.2.1及以上集群版本支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowClusterRedistribution(request *model.ShowClusterRedistributionRequest) (*model.ShowClusterRedistributionResponse, error) {
	requestDef := GenReqDefForShowClusterRedistribution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterRedistributionResponse), nil
	}
}

// ShowClusterRedistributionInvoker 查询重分布详情
func (c *DwsClient) ShowClusterRedistributionInvoker(request *model.ShowClusterRedistributionRequest) *ShowClusterRedistributionInvoker {
	requestDef := GenReqDefForShowClusterRedistribution()
	return &ShowClusterRedistributionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabaseAuthority 查询数据库对象权限
//
// 查询数据库对象权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowDatabaseAuthority(request *model.ShowDatabaseAuthorityRequest) (*model.ShowDatabaseAuthorityResponse, error) {
	requestDef := GenReqDefForShowDatabaseAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseAuthorityResponse), nil
	}
}

// ShowDatabaseAuthorityInvoker 查询数据库对象权限
func (c *DwsClient) ShowDatabaseAuthorityInvoker(request *model.ShowDatabaseAuthorityRequest) *ShowDatabaseAuthorityInvoker {
	requestDef := GenReqDefForShowDatabaseAuthority()
	return &ShowDatabaseAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabaseOmUserStatus 获得集群运维账户状态
//
// 获得数据库运维账户状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowDatabaseOmUserStatus(request *model.ShowDatabaseOmUserStatusRequest) (*model.ShowDatabaseOmUserStatusResponse, error) {
	requestDef := GenReqDefForShowDatabaseOmUserStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseOmUserStatusResponse), nil
	}
}

// ShowDatabaseOmUserStatusInvoker 获得集群运维账户状态
func (c *DwsClient) ShowDatabaseOmUserStatusInvoker(request *model.ShowDatabaseOmUserStatusRequest) *ShowDatabaseOmUserStatusInvoker {
	requestDef := GenReqDefForShowDatabaseOmUserStatus()
	return &ShowDatabaseOmUserStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabaseUser 查询指定用户信息
//
// 查询指定用户信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowDatabaseUser(request *model.ShowDatabaseUserRequest) (*model.ShowDatabaseUserResponse, error) {
	requestDef := GenReqDefForShowDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseUserResponse), nil
	}
}

// ShowDatabaseUserInvoker 查询指定用户信息
func (c *DwsClient) ShowDatabaseUserInvoker(request *model.ShowDatabaseUserRequest) *ShowDatabaseUserInvoker {
	requestDef := GenReqDefForShowDatabaseUser()
	return &ShowDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDisasterDetail 查询容灾详情
//
// 该接口用于查询单个容灾详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowDisasterDetail(request *model.ShowDisasterDetailRequest) (*model.ShowDisasterDetailResponse, error) {
	requestDef := GenReqDefForShowDisasterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDisasterDetailResponse), nil
	}
}

// ShowDisasterDetailInvoker 查询容灾详情
func (c *DwsClient) ShowDisasterDetailInvoker(request *model.ShowDisasterDetailRequest) *ShowDisasterDetailInvoker {
	requestDef := GenReqDefForShowDisasterDetail()
	return &ShowDisasterDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDisasterProgress 容灾-查询容灾进度详情
//
// 该接口用于查询容灾进度详情信息操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowDisasterProgress(request *model.ShowDisasterProgressRequest) (*model.ShowDisasterProgressResponse, error) {
	requestDef := GenReqDefForShowDisasterProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDisasterProgressResponse), nil
	}
}

// ShowDisasterProgressInvoker 容灾-查询容灾进度详情
func (c *DwsClient) ShowDisasterProgressInvoker(request *model.ShowDisasterProgressRequest) *ShowDisasterProgressInvoker {
	requestDef := GenReqDefForShowDisasterProgress()
	return &ShowDisasterProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询单个实例
//
// 查询单个实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询单个实例
func (c *DwsClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQueryDetail 查询SQL执行信息
//
// 查询SQL执行信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowQueryDetail(request *model.ShowQueryDetailRequest) (*model.ShowQueryDetailResponse, error) {
	requestDef := GenReqDefForShowQueryDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueryDetailResponse), nil
	}
}

// ShowQueryDetailInvoker 查询SQL执行信息
func (c *DwsClient) ShowQueryDetailInvoker(request *model.ShowQueryDetailRequest) *ShowQueryDetailInvoker {
	requestDef := GenReqDefForShowQueryDetail()
	return &ShowQueryDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceStatistics 查询资源统计
//
// 该接口用于查询资源统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowResourceStatistics(request *model.ShowResourceStatisticsRequest) (*model.ShowResourceStatisticsResponse, error) {
	requestDef := GenReqDefForShowResourceStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceStatisticsResponse), nil
	}
}

// ShowResourceStatisticsInvoker 查询资源统计
func (c *DwsClient) ShowResourceStatisticsInvoker(request *model.ShowResourceStatisticsRequest) *ShowResourceStatisticsInvoker {
	requestDef := GenReqDefForShowResourceStatistics()
	return &ShowResourceStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkloadPlan 查询某个工作负载计划详细信息
//
// 查询某个工作负载计划详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowWorkloadPlan(request *model.ShowWorkloadPlanRequest) (*model.ShowWorkloadPlanResponse, error) {
	requestDef := GenReqDefForShowWorkloadPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkloadPlanResponse), nil
	}
}

// ShowWorkloadPlanInvoker 查询某个工作负载计划详细信息
func (c *DwsClient) ShowWorkloadPlanInvoker(request *model.ShowWorkloadPlanRequest) *ShowWorkloadPlanInvoker {
	requestDef := GenReqDefForShowWorkloadPlan()
	return &ShowWorkloadPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkloadPlanStage 查询工作负载计划阶段详细信息
//
// 查询工作负载计划阶段详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowWorkloadPlanStage(request *model.ShowWorkloadPlanStageRequest) (*model.ShowWorkloadPlanStageResponse, error) {
	requestDef := GenReqDefForShowWorkloadPlanStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkloadPlanStageResponse), nil
	}
}

// ShowWorkloadPlanStageInvoker 查询工作负载计划阶段详细信息
func (c *DwsClient) ShowWorkloadPlanStageInvoker(request *model.ShowWorkloadPlanStageRequest) *ShowWorkloadPlanStageInvoker {
	requestDef := GenReqDefForShowWorkloadPlanStage()
	return &ShowWorkloadPlanStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkloadQueue 获得工作负载队列详细信息
//
// 获得工作负载队列详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) ShowWorkloadQueue(request *model.ShowWorkloadQueueRequest) (*model.ShowWorkloadQueueResponse, error) {
	requestDef := GenReqDefForShowWorkloadQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkloadQueueResponse), nil
	}
}

// ShowWorkloadQueueInvoker 获得工作负载队列详细信息
func (c *DwsClient) ShowWorkloadQueueInvoker(request *model.ShowWorkloadQueueRequest) *ShowWorkloadQueueInvoker {
	requestDef := GenReqDefForShowWorkloadQueue()
	return &ShowWorkloadQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 该接口用于启动容灾操作。
// 容灾状态为“未启动”、“启动失败”和“已停止”时可以执行启动容灾操作。
// 启动容灾后，生产集群和灾备集群将无法进行恢复、扩容、升级、重启、节点替换、更新密码等操作，此外，灾备集群将无法进行备份操作，请谨慎操作。
// 当容灾启动后，如果灾备集群容灾正常运行且容灾处于恢复状态中，此状态的集群会计费。
// 仅支持DWS 2.0集群。
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

// StartWorkloadPlan 启动工作负载计划
//
// 启动工作负载计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) StartWorkloadPlan(request *model.StartWorkloadPlanRequest) (*model.StartWorkloadPlanResponse, error) {
	requestDef := GenReqDefForStartWorkloadPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartWorkloadPlanResponse), nil
	}
}

// StartWorkloadPlanInvoker 启动工作负载计划
func (c *DwsClient) StartWorkloadPlanInvoker(request *model.StartWorkloadPlanRequest) *StartWorkloadPlanInvoker {
	requestDef := GenReqDefForStartWorkloadPlan()
	return &StartWorkloadPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopRedistribution 暂停重分布
//
// 该接口用于暂停运行状态下的重分布操作，重分布暂停状态可设置重分布优先级，修改重分布并发数等操作。
// 仅支持DWS 2.0集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) StopRedistribution(request *model.StopRedistributionRequest) (*model.StopRedistributionResponse, error) {
	requestDef := GenReqDefForStopRedistribution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopRedistributionResponse), nil
	}
}

// StopRedistributionInvoker 暂停重分布
func (c *DwsClient) StopRedistributionInvoker(request *model.StopRedistributionRequest) *StopRedistributionInvoker {
	requestDef := GenReqDefForStopRedistribution()
	return &StopRedistributionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopWorkloadPlan 停止工作负载计划
//
// 停止工作负载计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) StopWorkloadPlan(request *model.StopWorkloadPlanRequest) (*model.StopWorkloadPlanResponse, error) {
	requestDef := GenReqDefForStopWorkloadPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopWorkloadPlanResponse), nil
	}
}

// StopWorkloadPlanInvoker 停止工作负载计划
func (c *DwsClient) StopWorkloadPlanInvoker(request *model.StopWorkloadPlanRequest) *StopWorkloadPlanInvoker {
	requestDef := GenReqDefForStopWorkloadPlan()
	return &StopWorkloadPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchFailoverDisaster 容灾异常切换
//
// 该接口用于容灾异常场景下进行主备集群切换操作。
// “异常切换”按钮用于容灾异常或者生产集群故障情况下主备切换操作。
// 容灾异常切换仅8.1.2及以上集群版本支持。
// 异常切换会将灾备集群升为主，若原生产集群故障后存在部分数据未同步到灾备集群，那灾备集群升主后将缺少这些数据，切换时请确认容灾最后同步时间，谨慎操作。
// 仅支持DWS 2.0集群
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

// SwitchPlanStage 切换工作负载计划阶段
//
// 切换工作负载计划阶段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SwitchPlanStage(request *model.SwitchPlanStageRequest) (*model.SwitchPlanStageResponse, error) {
	requestDef := GenReqDefForSwitchPlanStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchPlanStageResponse), nil
	}
}

// SwitchPlanStageInvoker 切换工作负载计划阶段
func (c *DwsClient) SwitchPlanStageInvoker(request *model.SwitchPlanStageRequest) *SwitchPlanStageInvoker {
	requestDef := GenReqDefForSwitchPlanStage()
	return &SwitchPlanStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchoverDisasterRecovery 灾备切换
//
// 该接口用于容灾进行灾备切换操作。
// “灾备切换”按钮用于在容灾正常情况下主备倒换操作。
// 容灾状态为“运行中”时可以执行灾备切换操作。
// 灾备切换需要一定时间，在此期间，原生产集群将可不用。
// 不同场景下进行灾备切换，RPO（Recovery Point Object，灾难发生后，系统和数据必须恢复到的时间点要求。）说明如下：
//
//	生产集群在“可用”的状态下，RPO&#x3D;0。
//	生产集群在“不可用”的状态下，无法保证RPO&#x3D;0，但数据至少可恢复到生产集群“最近容灾成功时间”。
//
// 仅支持DWS 2.0集群。
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

// SyncIamUsers 同步IAM用户到数据库
//
// 同步IAM用户到数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) SyncIamUsers(request *model.SyncIamUsersRequest) (*model.SyncIamUsersResponse, error) {
	requestDef := GenReqDefForSyncIamUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncIamUsersResponse), nil
	}
}

// SyncIamUsersInvoker 同步IAM用户到数据库
func (c *DwsClient) SyncIamUsersInvoker(request *model.SyncIamUsersRequest) *SyncIamUsersInvoker {
	requestDef := GenReqDefForSyncIamUsers()
	return &SyncIamUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateDatabaseAuthority 修改数据库对象权限
//
// 修改数据库对象权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateDatabaseAuthority(request *model.UpdateDatabaseAuthorityRequest) (*model.UpdateDatabaseAuthorityResponse, error) {
	requestDef := GenReqDefForUpdateDatabaseAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseAuthorityResponse), nil
	}
}

// UpdateDatabaseAuthorityInvoker 修改数据库对象权限
func (c *DwsClient) UpdateDatabaseAuthorityInvoker(request *model.UpdateDatabaseAuthorityRequest) *UpdateDatabaseAuthorityInvoker {
	requestDef := GenReqDefForUpdateDatabaseAuthority()
	return &UpdateDatabaseAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatabaseUserInfo 修改指定用户信息
//
// 修改指定用户信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateDatabaseUserInfo(request *model.UpdateDatabaseUserInfoRequest) (*model.UpdateDatabaseUserInfoResponse, error) {
	requestDef := GenReqDefForUpdateDatabaseUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseUserInfoResponse), nil
	}
}

// UpdateDatabaseUserInfoInvoker 修改指定用户信息
func (c *DwsClient) UpdateDatabaseUserInfoInvoker(request *model.UpdateDatabaseUserInfoRequest) *UpdateDatabaseUserInfoInvoker {
	requestDef := GenReqDefForUpdateDatabaseUserInfo()
	return &UpdateDatabaseUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDisasterInfo 更新容灾配置
//
// 该接口用于更新容灾配置操作。
// 容灾状态为“未启动”或“已停止”时，可以执行容灾配置修改操作。
// 新的配置在容灾重新启动后生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateDisasterInfo(request *model.UpdateDisasterInfoRequest) (*model.UpdateDisasterInfoResponse, error) {
	requestDef := GenReqDefForUpdateDisasterInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDisasterInfoResponse), nil
	}
}

// UpdateDisasterInfoInvoker 更新容灾配置
func (c *DwsClient) UpdateDisasterInfoInvoker(request *model.UpdateDisasterInfoRequest) *UpdateDisasterInfoInvoker {
	requestDef := GenReqDefForUpdateDisasterInfo()
	return &UpdateDisasterInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateLogicalCluster 编辑逻辑集群
//
// 此接口用于编辑修改逻辑集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateLogicalCluster(request *model.UpdateLogicalClusterRequest) (*model.UpdateLogicalClusterResponse, error) {
	requestDef := GenReqDefForUpdateLogicalCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogicalClusterResponse), nil
	}
}

// UpdateLogicalClusterInvoker 编辑逻辑集群
func (c *DwsClient) UpdateLogicalClusterInvoker(request *model.UpdateLogicalClusterRequest) *UpdateLogicalClusterInvoker {
	requestDef := GenReqDefForUpdateLogicalCluster()
	return &UpdateLogicalClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateQueueResources 更新工作负载队列资源配置信息
//
// 更新工作负载队列资源配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateQueueResources(request *model.UpdateQueueResourcesRequest) (*model.UpdateQueueResourcesResponse, error) {
	requestDef := GenReqDefForUpdateQueueResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQueueResourcesResponse), nil
	}
}

// UpdateQueueResourcesInvoker 更新工作负载队列资源配置信息
func (c *DwsClient) UpdateQueueResourcesInvoker(request *model.UpdateQueueResourcesRequest) *UpdateQueueResourcesInvoker {
	requestDef := GenReqDefForUpdateQueueResources()
	return &UpdateQueueResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSchemas 更新模式空间限额
//
// 更新模式空间限额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwsClient) UpdateSchemas(request *model.UpdateSchemasRequest) (*model.UpdateSchemasResponse, error) {
	requestDef := GenReqDefForUpdateSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSchemasResponse), nil
	}
}

// UpdateSchemasInvoker 更新模式空间限额
func (c *DwsClient) UpdateSchemasInvoker(request *model.UpdateSchemasRequest) *UpdateSchemasInvoker {
	requestDef := GenReqDefForUpdateSchemas()
	return &UpdateSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
