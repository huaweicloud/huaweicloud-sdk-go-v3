package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sdrs/v1/model"
)

type SdrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSdrsClient(hcClient *http_client.HcHttpClient) *SdrsClient {
	return &SdrsClient{HcClient: hcClient}
}

func SdrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 保护实例添加网卡
//
// 给指定的保护实例添加网卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) AddProtectedInstanceNic(request *model.AddProtectedInstanceNicRequest) (*model.AddProtectedInstanceNicResponse, error) {
	requestDef := GenReqDefForAddProtectedInstanceNic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProtectedInstanceNicResponse), nil
	}
}

// 添加保护实例标签
//
// 一个保护实例上最多有10个标签。此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) AddProtectedInstanceTags(request *model.AddProtectedInstanceTagsRequest) (*model.AddProtectedInstanceTagsResponse, error) {
	requestDef := GenReqDefForAddProtectedInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProtectedInstanceTagsResponse), nil
	}
}

// 保护实例挂载复制对
//
// 将指定的复制对挂载到指定的保护实例上。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) AttachProtectedInstanceReplication(request *model.AttachProtectedInstanceReplicationRequest) (*model.AttachProtectedInstanceReplicationResponse, error) {
	requestDef := GenReqDefForAttachProtectedInstanceReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachProtectedInstanceReplicationResponse), nil
	}
}

// 批量添加保护实例标签
//
// 为指定保护实例批量添加或删除标签。一个资源上最多有10个标签。
// 此接口为幂等接口：
// 创建时如果请求体中存在重复key则报错。
// 创建时，不允许设置重复key数据,如果数据库已存在该key，就覆盖value的值。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) BatchAddTags(request *model.BatchAddTagsRequest) (*model.BatchAddTagsResponse, error) {
	requestDef := GenReqDefForBatchAddTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddTagsResponse), nil
	}
}

// 批量创建保护实例
//
// 典型场景：没有特殊操作场景
// 接口功能：批量创建保护实例。保护实例创建完成后，系统默认容灾站点云服务器名称与生产站点云服务器名称相同，但ID不同。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) BatchCreateProtectedInstances(request *model.BatchCreateProtectedInstancesRequest) (*model.BatchCreateProtectedInstancesResponse, error) {
	requestDef := GenReqDefForBatchCreateProtectedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateProtectedInstancesResponse), nil
	}
}

// 批量删除保护实例
//
// 典型场景：没有特殊操作场景
// 接口功能：批量删除保护实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) BatchDeleteProtectedInstances(request *model.BatchDeleteProtectedInstancesRequest) (*model.BatchDeleteProtectedInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteProtectedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProtectedInstancesResponse), nil
	}
}

// 批量删除保护实例标签
//
// 为指定保护实例批量删除标签。一个资源上最多有10个标签。
// 此接口为幂等接口：
// 删除时，如果删除的标签不存在，默认处理成功,删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) BatchDeleteTags(request *model.BatchDeleteTagsRequest) (*model.BatchDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagsResponse), nil
	}
}

// 创建容灾演练
//
// 创建容灾演练。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) CreateDisasterRecoveryDrill(request *model.CreateDisasterRecoveryDrillRequest) (*model.CreateDisasterRecoveryDrillResponse, error) {
	requestDef := GenReqDefForCreateDisasterRecoveryDrill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDisasterRecoveryDrillResponse), nil
	}
}

// 创建保护实例
//
// 创建保护实例。保护实例创建完成后，系统默认容灾站点云服务器名称与生产站点云服务器名称相同，但ID不同。如果需要修改云服务器名称，请在保护实例详情页面单击云服务器名称，进入云服务器详情页面进行修改
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) CreateProtectedInstance(request *model.CreateProtectedInstanceRequest) (*model.CreateProtectedInstanceResponse, error) {
	requestDef := GenReqDefForCreateProtectedInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProtectedInstanceResponse), nil
	}
}

// 创建保护组
//
// 创建保护组。
// 说明：
// 本接口为异步接口，调用成功只是表示请求下发，创建结果需要通过“查询job状态”接口获取
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) CreateProtectionGroup(request *model.CreateProtectionGroupRequest) (*model.CreateProtectionGroupResponse, error) {
	requestDef := GenReqDefForCreateProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProtectionGroupResponse), nil
	}
}

// 创建复制对
//
// 创建复制对，并将其添加到指定的保护组中。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) CreateReplication(request *model.CreateReplicationRequest) (*model.CreateReplicationResponse, error) {
	requestDef := GenReqDefForCreateReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReplicationResponse), nil
	}
}

// 删除所有保护组失败任务
//
// 删除所有保护组层级的失败任务，创建、删除保护组失败等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteAllServerGroupFailureJobs(request *model.DeleteAllServerGroupFailureJobsRequest) (*model.DeleteAllServerGroupFailureJobsResponse, error) {
	requestDef := GenReqDefForDeleteAllServerGroupFailureJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAllServerGroupFailureJobsResponse), nil
	}
}

// 删除容灾演练
//
// 删除指定的容灾演练。删除后：
// 容灾演练服务器、容灾演练服务器上挂载的磁盘和网卡将被一并删除。
// 演练VPC、演练VPC的子网不会被删除。您可以继续使用该VPC创建其他云服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteDisasterRecoveryDrill(request *model.DeleteDisasterRecoveryDrillRequest) (*model.DeleteDisasterRecoveryDrillResponse, error) {
	requestDef := GenReqDefForDeleteDisasterRecoveryDrill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDisasterRecoveryDrillResponse), nil
	}
}

// 删除单个失败任务
//
// 删除单个失败任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteFailureJob(request *model.DeleteFailureJobRequest) (*model.DeleteFailureJobResponse, error) {
	requestDef := GenReqDefForDeleteFailureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFailureJobResponse), nil
	}
}

// 删除保护实例
//
// 删除指定的保护实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteProtectedInstance(request *model.DeleteProtectedInstanceRequest) (*model.DeleteProtectedInstanceResponse, error) {
	requestDef := GenReqDefForDeleteProtectedInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedInstanceResponse), nil
	}
}

// 保护实例删除网卡
//
// 删除指定保护实例的指定网卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteProtectedInstanceNic(request *model.DeleteProtectedInstanceNicRequest) (*model.DeleteProtectedInstanceNicResponse, error) {
	requestDef := GenReqDefForDeleteProtectedInstanceNic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedInstanceNicResponse), nil
	}
}

// 删除保护实例标签
//
// 幂等接口：删除时，不对标签字符集做校验，调用接口前必须要做encodeURI，服务端需要对接口URI做decodeURI。
//  说明:请自行选择工具执行URI编码。
// 删除的key不存在报404，Key不能为空或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteProtectedInstanceTag(request *model.DeleteProtectedInstanceTagRequest) (*model.DeleteProtectedInstanceTagResponse, error) {
	requestDef := GenReqDefForDeleteProtectedInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectedInstanceTagResponse), nil
	}
}

// 删除保护组
//
// 删除指定的保护组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteProtectionGroup(request *model.DeleteProtectionGroupRequest) (*model.DeleteProtectionGroupResponse, error) {
	requestDef := GenReqDefForDeleteProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProtectionGroupResponse), nil
	}
}

// 删除复制对
//
// 删除指定的复制对。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteReplication(request *model.DeleteReplicationRequest) (*model.DeleteReplicationResponse, error) {
	requestDef := GenReqDefForDeleteReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReplicationResponse), nil
	}
}

// 删除指定保护组内的所有失败任务
//
// 删除指定保护组内的所有失败任务，创建保护实例失败、创建复制对失败、删除保护实例失败、删除复制对失败等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DeleteServerGroupFailureJobs(request *model.DeleteServerGroupFailureJobsRequest) (*model.DeleteServerGroupFailureJobsResponse, error) {
	requestDef := GenReqDefForDeleteServerGroupFailureJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupFailureJobsResponse), nil
	}
}

// 保护实例卸载复制对
//
// 将指定的复制对从指定的保护实例上卸载。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) DetachProtectedInstanceReplication(request *model.DetachProtectedInstanceReplicationRequest) (*model.DetachProtectedInstanceReplicationResponse, error) {
	requestDef := GenReqDefForDetachProtectedInstanceReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachProtectedInstanceReplicationResponse), nil
	}
}

// 复制对扩容
//
// 对复制对包含的两个磁盘进行扩容操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ExpandReplication(request *model.ExpandReplicationRequest) (*model.ExpandReplicationResponse, error) {
	requestDef := GenReqDefForExpandReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandReplicationResponse), nil
	}
}

// 查询双活域
//
// 查询双活域。双活域由本端存储设备、远端存储设备组成，通过双活域，应用服务器可以实现跨站点的数据访问。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListActiveActiveDomains(request *model.ListActiveActiveDomainsRequest) (*model.ListActiveActiveDomainsResponse, error) {
	requestDef := GenReqDefForListActiveActiveDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActiveActiveDomainsResponse), nil
	}
}

// 查询容灾演练列表
//
// 查询指定保护组下的所有容灾演练列表，当未指定保护组时查询当前租户下的所有容灾演练列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListDisasterRecoveryDrills(request *model.ListDisasterRecoveryDrillsRequest) (*model.ListDisasterRecoveryDrillsResponse, error) {
	requestDef := GenReqDefForListDisasterRecoveryDrills()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDisasterRecoveryDrillsResponse), nil
	}
}

// 查询失败任务列表
//
// 查询所有保护组失败任务列表或者指定保护组下的所有失败任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListFailureJobs(request *model.ListFailureJobsRequest) (*model.ListFailureJobsResponse, error) {
	requestDef := GenReqDefForListFailureJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFailureJobsResponse), nil
	}
}

// 查询保护实例标签
//
// 查询指定保护实例的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListProtectedInstanceTags(request *model.ListProtectedInstanceTagsRequest) (*model.ListProtectedInstanceTagsResponse, error) {
	requestDef := GenReqDefForListProtectedInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedInstanceTagsResponse), nil
	}
}

// 查询保护实例列表
//
// 查询当前租户下的所有保护实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListProtectedInstances(request *model.ListProtectedInstancesRequest) (*model.ListProtectedInstancesResponse, error) {
	requestDef := GenReqDefForListProtectedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedInstancesResponse), nil
	}
}

// 通过标签查询保护实例
//
// 使用标签过滤保护实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListProtectedInstancesByTags(request *model.ListProtectedInstancesByTagsRequest) (*model.ListProtectedInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListProtectedInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedInstancesByTagsResponse), nil
	}
}

// 查询保护实例项目标签
//
// 查询租户在指定Project中保护实例的所有资源标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListProtectedInstancesProjectTags(request *model.ListProtectedInstancesProjectTagsRequest) (*model.ListProtectedInstancesProjectTagsResponse, error) {
	requestDef := GenReqDefForListProtectedInstancesProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedInstancesProjectTagsResponse), nil
	}
}

// 查询保护组列表
//
// 查询当前租户所有的保护组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListProtectionGroups(request *model.ListProtectionGroupsRequest) (*model.ListProtectionGroupsResponse, error) {
	requestDef := GenReqDefForListProtectionGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionGroupsResponse), nil
	}
}

// 查询复制对列表
//
// 查询指定保护组下的所有复制对列表，如果不给定指定保护组则查询当前租户下的所有复制对列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListReplications(request *model.ListReplicationsRequest) (*model.ListReplicationsResponse, error) {
	requestDef := GenReqDefForListReplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReplicationsResponse), nil
	}
}

// 查询资源的RPO超标趋势记录列表
//
// 查询当前租户大屏显示中，资源的RPO超标趋势记录列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListRpoStatistics(request *model.ListRpoStatisticsRequest) (*model.ListRpoStatisticsResponse, error) {
	requestDef := GenReqDefForListRpoStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRpoStatisticsResponse), nil
	}
}

// 保护实例变更规格
//
// 变更指定保护实例中弹性云服务器的规格，包括：同时变更生产站点云服务器和容灾站点云服务器的规格。
// 仅变更生产站点云服务器的规格，容灾站点云服务器规格不变。
// 生产站点云服务器规格不变，仅变更容灾站点云服务器的规格。
// 当且仅当待变更规格的云服务器处于关机状态时，才能执行此操作。
//  说明：不同规格的云服务器在性能上存在差异，可能会对云服务器上运行的应用产生影响。
// 为保证切换/故障切换后云服务器的性能，建议容灾站点服务器的规格（CPU、内存）不低于生产站点云服务器的规格（CPU、内存）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ResizeProtectedInstance(request *model.ResizeProtectedInstanceRequest) (*model.ResizeProtectedInstanceResponse, error) {
	requestDef := GenReqDefForResizeProtectedInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeProtectedInstanceResponse), nil
	}
}

// 查询单个容灾演练详情
//
// 查询单个容灾演练的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowDisasterRecoveryDrill(request *model.ShowDisasterRecoveryDrillRequest) (*model.ShowDisasterRecoveryDrillResponse, error) {
	requestDef := GenReqDefForShowDisasterRecoveryDrill()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDisasterRecoveryDrillResponse), nil
	}
}

// 查询单个保护实例详情
//
// 查询单个保护实例的详细信息，如名称、ID等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowProtectedInstance(request *model.ShowProtectedInstanceRequest) (*model.ShowProtectedInstanceResponse, error) {
	requestDef := GenReqDefForShowProtectedInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectedInstanceResponse), nil
	}
}

// 查询保护组详情
//
// 查询单个保护组的详细信息，如ID、名称等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowProtectionGroup(request *model.ShowProtectionGroupRequest) (*model.ShowProtectionGroupResponse, error) {
	requestDef := GenReqDefForShowProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectionGroupResponse), nil
	}
}

// 查询租户配额
//
// 查询资源的配额相关信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// 查询单个复制对详情
//
// 查询单个复制对的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowReplication(request *model.ShowReplicationRequest) (*model.ShowReplicationResponse, error) {
	requestDef := GenReqDefForShowReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplicationResponse), nil
	}
}

// 保护组故障切换
//
// 当保护组的生产站点发生故障时，将保护组的生产站点切到当前的容灾站点，即另一端AZ，启用当前容灾站点的云硬盘以及云服务器等资源。
// 故障切换完成之后，保护组的当前生产站点变成故障切换发生之前的容灾站点，且生产站点和容灾站点之间的数据已停止保护，必须调用5.4.6-保护组开启保护/重保护接口成功后，两端的数据才会重新被保护。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) StartFailoverProtectionGroup(request *model.StartFailoverProtectionGroupRequest) (*model.StartFailoverProtectionGroupResponse, error) {
	requestDef := GenReqDefForStartFailoverProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartFailoverProtectionGroupResponse), nil
	}
}

// 保护组开启保护/重保护
//
// 对某一个保护组的“开启保护”或“重保护”操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) StartProtectionGroup(request *model.StartProtectionGroupRequest) (*model.StartProtectionGroupResponse, error) {
	requestDef := GenReqDefForStartProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartProtectionGroupResponse), nil
	}
}

// 保护组切换
//
// 对保护组进行切换操作，可以将保护组的当前生产站点，从创建保护组时指定的生产站点切换到创建保护组时指定的容灾站点，也可以从创建保护组时指定的容灾站点切换到创建保护组时指定的生产站点。切换后，生产站点和容灾站点的数据仍然处于被保护状态，只是复制方向与操作之前相反。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) StartReverseProtectionGroup(request *model.StartReverseProtectionGroupRequest) (*model.StartReverseProtectionGroupResponse, error) {
	requestDef := GenReqDefForStartReverseProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartReverseProtectionGroupResponse), nil
	}
}

// 保护组停止保护
//
// 对某一个保护组的停止保护操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) StopProtectionGroup(request *model.StopProtectionGroupRequest) (*model.StopProtectionGroupResponse, error) {
	requestDef := GenReqDefForStopProtectionGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopProtectionGroupResponse), nil
	}
}

// 更新容灾演练名称
//
// 更新容灾演练的名称。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) UpdateDisasterRecoveryDrillName(request *model.UpdateDisasterRecoveryDrillNameRequest) (*model.UpdateDisasterRecoveryDrillNameResponse, error) {
	requestDef := GenReqDefForUpdateDisasterRecoveryDrillName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDisasterRecoveryDrillNameResponse), nil
	}
}

// 更新保护实例名称
//
// 更新某一个保护实例的名称。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) UpdateProtectedInstanceName(request *model.UpdateProtectedInstanceNameRequest) (*model.UpdateProtectedInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateProtectedInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectedInstanceNameResponse), nil
	}
}

// 更新保护组名称
//
// 更新某一个保护组的名称。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) UpdateProtectionGroupName(request *model.UpdateProtectionGroupNameRequest) (*model.UpdateProtectionGroupNameResponse, error) {
	requestDef := GenReqDefForUpdateProtectionGroupName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectionGroupNameResponse), nil
	}
}

// 更新复制对名称
//
// 更新复制对名称。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) UpdateReplicationName(request *model.UpdateReplicationNameRequest) (*model.UpdateReplicationNameResponse, error) {
	requestDef := GenReqDefForUpdateReplicationName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReplicationNameResponse), nil
	}
}

// 查询API版本信息
//
// 查询存储容灾当前所有可用的版本信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// 查询指定API版本信息
//
// 查询存储容灾指定API版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowSpecifiedApiVersion(request *model.ShowSpecifiedApiVersionRequest) (*model.ShowSpecifiedApiVersionResponse, error) {
	requestDef := GenReqDefForShowSpecifiedApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpecifiedApiVersionResponse), nil
	}
}

// 查询job状态
//
// 查询job的执行状态。
// 对于创建保护组、删除保护组、创建保护实例、删除保护实例、创建复制对、删除复制对等异步API，命令下发后，会返回job_id，通过job_id可以查询任务的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SdrsClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}
