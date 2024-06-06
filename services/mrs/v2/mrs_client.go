package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v2/model"
)

type MrsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMrsClient(hcClient *httpclient.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchDeleteJobs 批量删除作业
//
// 在MRS集群中批量删除作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

// BatchDeleteJobsInvoker 批量删除作业
func (c *MrsClient) BatchDeleteJobsInvoker(request *model.BatchDeleteJobsRequest) *BatchDeleteJobsInvoker {
	requestDef := GenReqDefForBatchDeleteJobs()
	return &BatchDeleteJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAutoScalingPolicy 创建弹性伸缩策略
//
// 创建弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateAutoScalingPolicy(request *model.CreateAutoScalingPolicyRequest) (*model.CreateAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoScalingPolicyResponse), nil
	}
}

// CreateAutoScalingPolicyInvoker 创建弹性伸缩策略
func (c *MrsClient) CreateAutoScalingPolicyInvoker(request *model.CreateAutoScalingPolicyRequest) *CreateAutoScalingPolicyInvoker {
	requestDef := GenReqDefForCreateAutoScalingPolicy()
	return &CreateAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建集群
//
// 创建一个MRS集群。使用接口前，您需要先获取下的资源信息。
// - 通过VPC创建或查询VPC、子网
// - 通过ECS创建或查询密钥对
// - 通过[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)获取区域信息
// - 参考[MRS服务支持的组件](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)获取MRS版本及对应版本支持的组件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建集群
func (c *MrsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExecuteJob 新增并执行作业
//
// 在MRS集群中新增并提交一个作业。
//
// 需要先在集群详情页的“概览”页签，单击“IAM用户同步”右侧的“同步”进行IAM用户同步，然后再通过该接口提交作业。
//
// 如需使用OBS加密功能，请先参考“MRS用户指南 &gt; 管理现有集群 &gt; 作业管理 &gt; 使用OBS加密数据运行作业”页面进行相关配置后，再调用API接口运行作业。
//
// 所有示例中涉及的OBS路径、样例文件及终端节点和AKSK，请提前准备并在提交请求时根据实际情况替换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateExecuteJob(request *model.CreateExecuteJobRequest) (*model.CreateExecuteJobResponse, error) {
	requestDef := GenReqDefForCreateExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExecuteJobResponse), nil
	}
}

// CreateExecuteJobInvoker 新增并执行作业
func (c *MrsClient) CreateExecuteJobInvoker(request *model.CreateExecuteJobRequest) *CreateExecuteJobInvoker {
	requestDef := GenReqDefForCreateExecuteJob()
	return &CreateExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAutoScalingPolicy 删除弹性伸缩策略
//
// 删除弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) DeleteAutoScalingPolicy(request *model.DeleteAutoScalingPolicyRequest) (*model.DeleteAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForDeleteAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAutoScalingPolicyResponse), nil
	}
}

// DeleteAutoScalingPolicyInvoker 删除弹性伸缩策略
func (c *MrsClient) DeleteAutoScalingPolicyInvoker(request *model.DeleteAutoScalingPolicyRequest) *DeleteAutoScalingPolicyInvoker {
	requestDef := GenReqDefForDeleteAutoScalingPolicy()
	return &DeleteAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunJobFlow 创建集群并提交作业
//
// 创建一个MRS集群并提交作业，并支持作业完成后删除集群，支持MRS 1.8.9及以上集群版本使用。使用接口前，您需要先获取下的资源信息。
// - 通过VPC创建或查询VPC、子网
// - 通过ECS创建或查询密钥对
// - 通过[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)获取区域信息
// - 参考[MRS服务支持的组件](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)获取MRS版本及对应版本支持的组件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) RunJobFlow(request *model.RunJobFlowRequest) (*model.RunJobFlowResponse, error) {
	requestDef := GenReqDefForRunJobFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunJobFlowResponse), nil
	}
}

// RunJobFlowInvoker 创建集群并提交作业
func (c *MrsClient) RunJobFlowInvoker(request *model.RunJobFlowRequest) *RunJobFlowInvoker {
	requestDef := GenReqDefForRunJobFlow()
	return &RunJobFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgencyMapping 查询用户（组）与IAM委托的映射关系
//
// 获取用户（组）与IAM委托之间的映射关系的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowAgencyMapping(request *model.ShowAgencyMappingRequest) (*model.ShowAgencyMappingResponse, error) {
	requestDef := GenReqDefForShowAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyMappingResponse), nil
	}
}

// ShowAgencyMappingInvoker 查询用户（组）与IAM委托的映射关系
func (c *MrsClient) ShowAgencyMappingInvoker(request *model.ShowAgencyMappingRequest) *ShowAgencyMappingInvoker {
	requestDef := GenReqDefForShowAgencyMapping()
	return &ShowAgencyMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoScalingPolicy 查看弹性伸缩策略
//
// 查看指定集群的所有的弹性伸缩策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowAutoScalingPolicy(request *model.ShowAutoScalingPolicyRequest) (*model.ShowAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForShowAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoScalingPolicyResponse), nil
	}
}

// ShowAutoScalingPolicyInvoker 查看弹性伸缩策略
func (c *MrsClient) ShowAutoScalingPolicyInvoker(request *model.ShowAutoScalingPolicyRequest) *ShowAutoScalingPolicyInvoker {
	requestDef := GenReqDefForShowAutoScalingPolicy()
	return &ShowAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobExeListNew 查询作业列表信息
//
// 在MRS指定集群中查询作业列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowJobExeListNew(request *model.ShowJobExeListNewRequest) (*model.ShowJobExeListNewResponse, error) {
	requestDef := GenReqDefForShowJobExeListNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobExeListNewResponse), nil
	}
}

// ShowJobExeListNewInvoker 查询作业列表信息
func (c *MrsClient) ShowJobExeListNewInvoker(request *model.ShowJobExeListNewRequest) *ShowJobExeListNewInvoker {
	requestDef := GenReqDefForShowJobExeListNew()
	return &ShowJobExeListNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSingleJobExe 查询单个作业信息
//
// 在MRS集群中查询指定作业的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowSingleJobExe(request *model.ShowSingleJobExeRequest) (*model.ShowSingleJobExeResponse, error) {
	requestDef := GenReqDefForShowSingleJobExe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleJobExeResponse), nil
	}
}

// ShowSingleJobExeInvoker 查询单个作业信息
func (c *MrsClient) ShowSingleJobExeInvoker(request *model.ShowSingleJobExeRequest) *ShowSingleJobExeInvoker {
	requestDef := GenReqDefForShowSingleJobExe()
	return &ShowSingleJobExeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlResultWithJob 获取SQL结果
//
// 在MRS集群中查询SparkSql和SparkScript两种类型作业的SQL语句运行完成后返回的查询结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowSqlResultWithJob(request *model.ShowSqlResultWithJobRequest) (*model.ShowSqlResultWithJobResponse, error) {
	requestDef := GenReqDefForShowSqlResultWithJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultWithJobResponse), nil
	}
}

// ShowSqlResultWithJobInvoker 获取SQL结果
func (c *MrsClient) ShowSqlResultWithJobInvoker(request *model.ShowSqlResultWithJobRequest) *ShowSqlResultWithJobInvoker {
	requestDef := GenReqDefForShowSqlResultWithJob()
	return &ShowSqlResultWithJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJob 终止作业
//
// 在MRS集群中终止指定作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

// StopJobInvoker 终止作业
func (c *MrsClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAgencyMapping 更新用户（组）与IAM委托的映射关系
//
// 更新用户（组）与IAM委托之间的映射关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateAgencyMapping(request *model.UpdateAgencyMappingRequest) (*model.UpdateAgencyMappingResponse, error) {
	requestDef := GenReqDefForUpdateAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgencyMappingResponse), nil
	}
}

// UpdateAgencyMappingInvoker 更新用户（组）与IAM委托的映射关系
func (c *MrsClient) UpdateAgencyMappingInvoker(request *model.UpdateAgencyMappingRequest) *UpdateAgencyMappingInvoker {
	requestDef := GenReqDefForUpdateAgencyMapping()
	return &UpdateAgencyMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoScalingPolicy 更新弹性伸缩策略
//
// 更新弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateAutoScalingPolicy(request *model.UpdateAutoScalingPolicyRequest) (*model.UpdateAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoScalingPolicyResponse), nil
	}
}

// UpdateAutoScalingPolicyInvoker 更新弹性伸缩策略
func (c *MrsClient) UpdateAutoScalingPolicyInvoker(request *model.UpdateAutoScalingPolicyRequest) *UpdateAutoScalingPolicyInvoker {
	requestDef := GenReqDefForUpdateAutoScalingPolicy()
	return &UpdateAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterName 修改集群名称
//
// 修改集群名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateClusterName(request *model.UpdateClusterNameRequest) (*model.UpdateClusterNameResponse, error) {
	requestDef := GenReqDefForUpdateClusterName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterNameResponse), nil
	}
}

// UpdateClusterNameInvoker 修改集群名称
func (c *MrsClient) UpdateClusterNameInvoker(request *model.UpdateClusterNameRequest) *UpdateClusterNameInvoker {
	requestDef := GenReqDefForUpdateClusterName()
	return &UpdateClusterNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddComponent 集群添加组件
//
// 集群添加组件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) AddComponent(request *model.AddComponentRequest) (*model.AddComponentResponse, error) {
	requestDef := GenReqDefForAddComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddComponentResponse), nil
	}
}

// AddComponentInvoker 集群添加组件
func (c *MrsClient) AddComponentInvoker(request *model.AddComponentRequest) *AddComponentInvoker {
	requestDef := GenReqDefForAddComponent()
	return &AddComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandCluster 扩容集群
//
// 对MRS集群进行扩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ExpandCluster(request *model.ExpandClusterRequest) (*model.ExpandClusterResponse, error) {
	requestDef := GenReqDefForExpandCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandClusterResponse), nil
	}
}

// ExpandClusterInvoker 扩容集群
func (c *MrsClient) ExpandClusterInvoker(request *model.ExpandClusterRequest) *ExpandClusterInvoker {
	requestDef := GenReqDefForExpandCluster()
	return &ExpandClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodes 查询集群节点列表
//
// 查询集群节点列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListNodes(request *model.ListNodesRequest) (*model.ListNodesResponse, error) {
	requestDef := GenReqDefForListNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodesResponse), nil
	}
}

// ListNodesInvoker 查询集群节点列表
func (c *MrsClient) ListNodesInvoker(request *model.ListNodesRequest) *ListNodesInvoker {
	requestDef := GenReqDefForListNodes()
	return &ListNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkCluster 缩容集群
//
// 对MRS集群进行缩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShrinkCluster(request *model.ShrinkClusterRequest) (*model.ShrinkClusterResponse, error) {
	requestDef := GenReqDefForShrinkCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkClusterResponse), nil
	}
}

// ShrinkClusterInvoker 缩容集群
func (c *MrsClient) ShrinkClusterInvoker(request *model.ShrinkClusterRequest) *ShrinkClusterInvoker {
	requestDef := GenReqDefForShrinkCluster()
	return &ShrinkClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataConnector 创建数据连接
//
// 创建数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateDataConnector(request *model.CreateDataConnectorRequest) (*model.CreateDataConnectorResponse, error) {
	requestDef := GenReqDefForCreateDataConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataConnectorResponse), nil
	}
}

// CreateDataConnectorInvoker 创建数据连接
func (c *MrsClient) CreateDataConnectorInvoker(request *model.CreateDataConnectorRequest) *CreateDataConnectorInvoker {
	requestDef := GenReqDefForCreateDataConnector()
	return &CreateDataConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataConnector 删除数据连接
//
// 删除数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) DeleteDataConnector(request *model.DeleteDataConnectorRequest) (*model.DeleteDataConnectorResponse, error) {
	requestDef := GenReqDefForDeleteDataConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataConnectorResponse), nil
	}
}

// DeleteDataConnectorInvoker 删除数据连接
func (c *MrsClient) DeleteDataConnectorInvoker(request *model.DeleteDataConnectorRequest) *DeleteDataConnectorInvoker {
	requestDef := GenReqDefForDeleteDataConnector()
	return &DeleteDataConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataConnector 查询数据连接列表
//
// 查询数据连接列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListDataConnector(request *model.ListDataConnectorRequest) (*model.ListDataConnectorResponse, error) {
	requestDef := GenReqDefForListDataConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataConnectorResponse), nil
	}
}

// ListDataConnectorInvoker 查询数据连接列表
func (c *MrsClient) ListDataConnectorInvoker(request *model.ListDataConnectorRequest) *ListDataConnectorInvoker {
	requestDef := GenReqDefForListDataConnector()
	return &ListDataConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataConnector 更新数据连接
//
// 更新数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateDataConnector(request *model.UpdateDataConnectorRequest) (*model.UpdateDataConnectorResponse, error) {
	requestDef := GenReqDefForUpdateDataConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataConnectorResponse), nil
	}
}

// UpdateDataConnectorInvoker 更新数据连接
func (c *MrsClient) UpdateDataConnectorInvoker(request *model.UpdateDataConnectorRequest) *UpdateDataConnectorInvoker {
	requestDef := GenReqDefForUpdateDataConnector()
	return &UpdateDataConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHdfsFileList 获取指定目录文件列表
//
// 在MRS集群中获取指定目录文件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowHdfsFileList(request *model.ShowHdfsFileListRequest) (*model.ShowHdfsFileListResponse, error) {
	requestDef := GenReqDefForShowHdfsFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHdfsFileListResponse), nil
	}
}

// ShowHdfsFileListInvoker 获取指定目录文件列表
func (c *MrsClient) ShowHdfsFileListInvoker(request *model.ShowHdfsFileListRequest) *ShowHdfsFileListInvoker {
	requestDef := GenReqDefForShowHdfsFileList()
	return &ShowHdfsFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSyncIamUser 指定用户、用户组取消同步
//
// 指定用户、用户组取消同步
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CancelSyncIamUser(request *model.CancelSyncIamUserRequest) (*model.CancelSyncIamUserResponse, error) {
	requestDef := GenReqDefForCancelSyncIamUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSyncIamUserResponse), nil
	}
}

// CancelSyncIamUserInvoker 指定用户、用户组取消同步
func (c *MrsClient) CancelSyncIamUserInvoker(request *model.CancelSyncIamUserRequest) *CancelSyncIamUserInvoker {
	requestDef := GenReqDefForCancelSyncIamUser()
	return &CancelSyncIamUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSyncIamUser 获取已经同步的IAM用户和用户组
//
// 获取已经同步的IAM用户和用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowSyncIamUser(request *model.ShowSyncIamUserRequest) (*model.ShowSyncIamUserResponse, error) {
	requestDef := GenReqDefForShowSyncIamUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSyncIamUserResponse), nil
	}
}

// ShowSyncIamUserInvoker 获取已经同步的IAM用户和用户组
func (c *MrsClient) ShowSyncIamUserInvoker(request *model.ShowSyncIamUserRequest) *ShowSyncIamUserInvoker {
	requestDef := GenReqDefForShowSyncIamUser()
	return &ShowSyncIamUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSyncIamUser IAM同步
//
// 将IAM用户和用户组同步到manager，指定用户的情况下，会将该用户关联的IAM用户组也同步到manager。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateSyncIamUser(request *model.UpdateSyncIamUserRequest) (*model.UpdateSyncIamUserResponse, error) {
	requestDef := GenReqDefForUpdateSyncIamUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSyncIamUserResponse), nil
	}
}

// UpdateSyncIamUserInvoker IAM同步
func (c *MrsClient) UpdateSyncIamUserInvoker(request *model.UpdateSyncIamUserRequest) *UpdateSyncIamUserInvoker {
	requestDef := GenReqDefForUpdateSyncIamUser()
	return &UpdateSyncIamUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSql 取消SQL执行任务
//
// 在MRS集群中取消一条SQL的执行任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CancelSql(request *model.CancelSqlRequest) (*model.CancelSqlResponse, error) {
	requestDef := GenReqDefForCancelSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSqlResponse), nil
	}
}

// CancelSqlInvoker 取消SQL执行任务
func (c *MrsClient) CancelSqlInvoker(request *model.CancelSqlRequest) *CancelSqlInvoker {
	requestDef := GenReqDefForCancelSql()
	return &CancelSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteSql 提交SQL语句
//
// 在MRS集群中提交并执行一条SQL语句。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ExecuteSql(request *model.ExecuteSqlRequest) (*model.ExecuteSqlResponse, error) {
	requestDef := GenReqDefForExecuteSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSqlResponse), nil
	}
}

// ExecuteSqlInvoker 提交SQL语句
func (c *MrsClient) ExecuteSqlInvoker(request *model.ExecuteSqlRequest) *ExecuteSqlInvoker {
	requestDef := GenReqDefForExecuteSql()
	return &ExecuteSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlResult 查询SQL结果
//
// 在MRS集群中查询一条SQL的执行结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowSqlResult(request *model.ShowSqlResultRequest) (*model.ShowSqlResultResponse, error) {
	requestDef := GenReqDefForShowSqlResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultResponse), nil
	}
}

// ShowSqlResultInvoker 查询SQL结果
func (c *MrsClient) ShowSqlResultInvoker(request *model.ShowSqlResultRequest) *ShowSqlResultInvoker {
	requestDef := GenReqDefForShowSqlResult()
	return &ShowSqlResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTagQuota 查询标签配额
//
// 查询标签配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowTagQuota(request *model.ShowTagQuotaRequest) (*model.ShowTagQuotaResponse, error) {
	requestDef := GenReqDefForShowTagQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagQuotaResponse), nil
	}
}

// ShowTagQuotaInvoker 查询标签配额
func (c *MrsClient) ShowTagQuotaInvoker(request *model.ShowTagQuotaRequest) *ShowTagQuotaInvoker {
	requestDef := GenReqDefForShowTagQuota()
	return &ShowTagQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTagStatus 查询默认标签状态
//
// 查询集群默认标签状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowTagStatus(request *model.ShowTagStatusRequest) (*model.ShowTagStatusResponse, error) {
	requestDef := GenReqDefForShowTagStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagStatusResponse), nil
	}
}

// ShowTagStatusInvoker 查询默认标签状态
func (c *MrsClient) ShowTagStatusInvoker(request *model.ShowTagStatusRequest) *ShowTagStatusInvoker {
	requestDef := GenReqDefForShowTagStatus()
	return &ShowTagStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchClusterTags 集群操作默认标签
//
// 对已有集群启用或关闭集群默认标签。开启后，集群内节点会打上集群默认标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) SwitchClusterTags(request *model.SwitchClusterTagsRequest) (*model.SwitchClusterTagsResponse, error) {
	requestDef := GenReqDefForSwitchClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchClusterTagsResponse), nil
	}
}

// SwitchClusterTagsInvoker 集群操作默认标签
func (c *MrsClient) SwitchClusterTagsInvoker(request *model.SwitchClusterTagsRequest) *SwitchClusterTagsInvoker {
	requestDef := GenReqDefForSwitchClusterTags()
	return &SwitchClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMrsFlavors 查询MRS集群版本可用的规格
//
// 查询MRS集群版本可用的规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowMrsFlavors(request *model.ShowMrsFlavorsRequest) (*model.ShowMrsFlavorsResponse, error) {
	requestDef := GenReqDefForShowMrsFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMrsFlavorsResponse), nil
	}
}

// ShowMrsFlavorsInvoker 查询MRS集群版本可用的规格
func (c *MrsClient) ShowMrsFlavorsInvoker(request *model.ShowMrsFlavorsRequest) *ShowMrsFlavorsInvoker {
	requestDef := GenReqDefForShowMrsFlavors()
	return &ShowMrsFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMrsVersionList 展示MRS版本列表
//
// 展示MRS版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowMrsVersionList(request *model.ShowMrsVersionListRequest) (*model.ShowMrsVersionListResponse, error) {
	requestDef := GenReqDefForShowMrsVersionList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMrsVersionListResponse), nil
	}
}

// ShowMrsVersionListInvoker 展示MRS版本列表
func (c *MrsClient) ShowMrsVersionListInvoker(request *model.ShowMrsVersionListRequest) *ShowMrsVersionListInvoker {
	requestDef := GenReqDefForShowMrsVersionList()
	return &ShowMrsVersionListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
