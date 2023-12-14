package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v1/model"
)

type MrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMrsClient(hcClient *http_client.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateClusterTags 批量添加集群标签
//
// 为指定集群批量添加标签。
//
// 一个集群上最多有10个标签。
//
// 此接口为幂等接口：
//
// - 创建时，同一个集群不允许重复key，如果数据库存在就覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) BatchCreateClusterTags(request *model.BatchCreateClusterTagsRequest) (*model.BatchCreateClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateClusterTagsResponse), nil
	}
}

// BatchCreateClusterTagsInvoker 批量添加集群标签
func (c *MrsClient) BatchCreateClusterTagsInvoker(request *model.BatchCreateClusterTagsRequest) *BatchCreateClusterTagsInvoker {
	requestDef := GenReqDefForBatchCreateClusterTags()
	return &BatchCreateClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteClusterTags 批量删除集群标签
//
// 为指定集群批量删除标签。
//
// 一个集群上最多有10个标签。
//
// 此接口为幂等接口：
//
// -
// 删除时，如果删除的标签不存在，默认处理成功，删除时不对标签字符集范围做校验。Key长度36个unicode字符，value为43个unicode字符。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) BatchDeleteClusterTags(request *model.BatchDeleteClusterTagsRequest) (*model.BatchDeleteClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteClusterTagsResponse), nil
	}
}

// BatchDeleteClusterTagsInvoker 批量删除集群标签
func (c *MrsClient) BatchDeleteClusterTagsInvoker(request *model.BatchDeleteClusterTagsRequest) *BatchDeleteClusterTagsInvoker {
	requestDef := GenReqDefForBatchDeleteClusterTags()
	return &BatchDeleteClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateAndExecuteJob 新增作业并执行（废弃）
//
// 如需使用作业管理接口请参考apiv2接口使用，本接口后续不再进行维护。
// 在MRS集群中新增一个作业，并执行作业。该接口不兼容Sahara。
// 集群ID可参考[查询集群列表](https://support.huaweicloud.com/api-mrs/ListClusters.html)接口获取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateAndExecuteJob(request *model.CreateAndExecuteJobRequest) (*model.CreateAndExecuteJobResponse, error) {
	requestDef := GenReqDefForCreateAndExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAndExecuteJobResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateAndExecuteJobInvoker 新增作业并执行（废弃）
func (c *MrsClient) CreateAndExecuteJobInvoker(request *model.CreateAndExecuteJobRequest) *CreateAndExecuteJobInvoker {
	requestDef := GenReqDefForCreateAndExecuteJob()
	return &CreateAndExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建集群并执行作业
//
// 创建一个MRS集群，并在集群中提交一个作业。该接口不兼容Sahara。
// 支持同一时间并发创建10个集群。
// 使用接口前，您需要先获取下的资源信息。
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

// CreateClusterInvoker 创建集群并执行作业
func (c *MrsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterTag 给指定集群添加标签
//
// 为特定的集群添加一个tag。
// 一个集群上最多有10个标签，此接口为幂等接口。添加标签时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateClusterTag(request *model.CreateClusterTagRequest) (*model.CreateClusterTagResponse, error) {
	requestDef := GenReqDefForCreateClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterTagResponse), nil
	}
}

// CreateClusterTagInvoker 给指定集群添加标签
func (c *MrsClient) CreateClusterTagInvoker(request *model.CreateClusterTagRequest) *CreateClusterTagInvoker {
	requestDef := GenReqDefForCreateClusterTag()
	return &CreateClusterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScalingPolicy 配置弹性伸缩规则
//
// 对弹性伸缩规则进行编辑。
//
// 在创建集群并执行作业接口中也可以创建弹性伸缩规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) CreateScalingPolicy(request *model.CreateScalingPolicyRequest) (*model.CreateScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScalingPolicyResponse), nil
	}
}

// CreateScalingPolicyInvoker 配置弹性伸缩规则
func (c *MrsClient) CreateScalingPolicyInvoker(request *model.CreateScalingPolicyRequest) *CreateScalingPolicyInvoker {
	requestDef := GenReqDefForCreateScalingPolicy()
	return &CreateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除集群
//
// 数据完成处理分析后或者集群运行异常无法提供服务时可删除集群服务。该接口兼容Sahara。
//
// 处于如下状态的集群不允许删除：
// - scaling-out：扩容中
// - scaling-in：缩容中
// - starting：启动中
// - terminating：删除中
// - terminated：已删除
// - failed：失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除集群
func (c *MrsClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClusterTag 删除指定集群的标签
//
// 删除特定集群的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) DeleteClusterTag(request *model.DeleteClusterTagRequest) (*model.DeleteClusterTagResponse, error) {
	requestDef := GenReqDefForDeleteClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterTagResponse), nil
	}
}

// DeleteClusterTagInvoker 删除指定集群的标签
func (c *MrsClient) DeleteClusterTagInvoker(request *model.DeleteClusterTagRequest) *DeleteClusterTagInvoker {
	requestDef := GenReqDefForDeleteClusterTag()
	return &DeleteClusterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteJobExecution 删除作业执行对象（废弃）
//
// 如需使用作业管理接口请参考apiv2接口使用，本接口后续不再进行维护。
// 删除指定的作业执行对象。该接口兼容Sahara。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) DeleteJobExecution(request *model.DeleteJobExecutionRequest) (*model.DeleteJobExecutionResponse, error) {
	requestDef := GenReqDefForDeleteJobExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobExecutionResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteJobExecutionInvoker 删除作业执行对象（废弃）
func (c *MrsClient) DeleteJobExecutionInvoker(request *model.DeleteJobExecutionRequest) *DeleteJobExecutionInvoker {
	requestDef := GenReqDefForDeleteJobExecution()
	return &DeleteJobExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTags 查询所有标签
//
// 查询租户在指定Region下的所有标签集合 。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

// ListAllTagsInvoker 查询所有标签
func (c *MrsClient) ListAllTagsInvoker(request *model.ListAllTagsRequest) *ListAllTagsInvoker {
	requestDef := GenReqDefForListAllTags()
	return &ListAllTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterTags 查询指定集群的标签
//
// 查询指定集群的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListClusterTags(request *model.ListClusterTagsRequest) (*model.ListClusterTagsResponse, error) {
	requestDef := GenReqDefForListClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterTagsResponse), nil
	}
}

// ListClusterTagsInvoker 查询指定集群的标签
func (c *MrsClient) ListClusterTagsInvoker(request *model.ListClusterTagsRequest) *ListClusterTagsInvoker {
	requestDef := GenReqDefForListClusterTags()
	return &ListClusterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 查询集群列表
//
// 查看用户创建的集群列表信息。该接口不兼容Sahara。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 查询集群列表
func (c *MrsClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClustersByTags 查询特定标签的集群列表
//
// 使用标签过滤集群。
//
// 集群默认按照创建时间倒序，集群tag也按照创建时间倒序。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListClustersByTags(request *model.ListClustersByTagsRequest) (*model.ListClustersByTagsResponse, error) {
	requestDef := GenReqDefForListClustersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersByTagsResponse), nil
	}
}

// ListClustersByTagsInvoker 查询特定标签的集群列表
func (c *MrsClient) ListClustersByTagsInvoker(request *model.ListClustersByTagsRequest) *ListClustersByTagsInvoker {
	requestDef := GenReqDefForListClustersByTags()
	return &ListClustersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListExecuteJob 查询作业exe对象列表（废弃）
//
// 如需使用作业管理接口请参考apiv2接口使用，本接口后续不再进行维护。
// 查询所有作业的exe对象列表。该接口不兼容Sahara。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListExecuteJob(request *model.ListExecuteJobRequest) (*model.ListExecuteJobResponse, error) {
	requestDef := GenReqDefForListExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecuteJobResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListExecuteJobInvoker 查询作业exe对象列表（废弃）
func (c *MrsClient) ListExecuteJobInvoker(request *model.ListExecuteJobRequest) *ListExecuteJobInvoker {
	requestDef := GenReqDefForListExecuteJob()
	return &ListExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHosts 查询主机列表
//
// 该接口用于查询输入集群的主机列表详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

// ListHostsInvoker 查询主机列表
func (c *MrsClient) ListHostsInvoker(request *model.ListHostsRequest) *ListHostsInvoker {
	requestDef := GenReqDefForListHosts()
	return &ListHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterDetails 查询集群详情
//
// 查看指定集群的详细信息。该接口不兼容Sahara。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowClusterDetails(request *model.ShowClusterDetailsRequest) (*model.ShowClusterDetailsResponse, error) {
	requestDef := GenReqDefForShowClusterDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailsResponse), nil
	}
}

// ShowClusterDetailsInvoker 查询集群详情
func (c *MrsClient) ShowClusterDetailsInvoker(request *model.ShowClusterDetailsRequest) *ShowClusterDetailsInvoker {
	requestDef := GenReqDefForShowClusterDetails()
	return &ShowClusterDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowJobExes 查询作业exe对象详情（废弃）
//
// 如需使用作业管理接口请参考apiv2接口使用，本接口后续不再进行维护。
// 查询指定作业的exe对象详细信息。该接口不兼容Sahara。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowJobExes(request *model.ShowJobExesRequest) (*model.ShowJobExesResponse, error) {
	requestDef := GenReqDefForShowJobExes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobExesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowJobExesInvoker 查询作业exe对象详情（废弃）
func (c *MrsClient) ShowJobExesInvoker(request *model.ShowJobExesRequest) *ShowJobExesInvoker {
	requestDef := GenReqDefForShowJobExes()
	return &ShowJobExesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterScaling 调整集群节点
//
// 创建集群后，扩容/缩容集群Core节点或者Task节点。MRS集群创建成功后不支持调整Master节点数量，即不支持扩缩容Master节点。该接口不兼容Sahara。
// 处于running状态的集群才允许扩容/缩容，其他状态则不允许扩容/缩容。 集群状态和集群ID可参考[查询集群列表](https://support.huaweicloud.com/api-mrs/ListClusters.html)接口获取。 本章节的接口只支持流式集群、分析集群和混合集群，不支持自定义集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) UpdateClusterScaling(request *model.UpdateClusterScalingRequest) (*model.UpdateClusterScalingResponse, error) {
	requestDef := GenReqDefForUpdateClusterScaling()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterScalingResponse), nil
	}
}

// UpdateClusterScalingInvoker 调整集群节点
func (c *MrsClient) UpdateClusterScalingInvoker(request *model.UpdateClusterScalingRequest) *UpdateClusterScalingInvoker {
	requestDef := GenReqDefForUpdateClusterScaling()
	return &UpdateClusterScalingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 查询可用区信息
//
// 在创建集群时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 查询可用区信息
func (c *MrsClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMrsVersionMetadata 查询对应版本元数据
//
// 查询对应版本元数据。如果参数里指定集群id，则可查询集群更新过补丁之后的最新元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MrsClient) ShowMrsVersionMetadata(request *model.ShowMrsVersionMetadataRequest) (*model.ShowMrsVersionMetadataResponse, error) {
	requestDef := GenReqDefForShowMrsVersionMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMrsVersionMetadataResponse), nil
	}
}

// ShowMrsVersionMetadataInvoker 查询对应版本元数据
func (c *MrsClient) ShowMrsVersionMetadataInvoker(request *model.ShowMrsVersionMetadataRequest) *ShowMrsVersionMetadataInvoker {
	requestDef := GenReqDefForShowMrsVersionMetadata()
	return &ShowMrsVersionMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
