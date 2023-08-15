package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cph/v1/model"
)

type CphClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCphClient(hcClient *http_client.HcHttpClient) *CphClient {
	return &CphClient{HcClient: hcClient}
}

func CphClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateTags 批量添加标签
//
// 批量添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量添加标签
func (c *CphClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTags 批量删除标签
//
// 批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) BatchDeleteTags(request *model.BatchDeleteTagsRequest) (*model.BatchDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagsResponse), nil
	}
}

// BatchDeleteTagsInvoker 批量删除标签
func (c *CphClient) BatchDeleteTagsInvoker(request *model.BatchDeleteTagsRequest) *BatchDeleteTagsInvoker {
	requestDef := GenReqDefForBatchDeleteTags()
	return &BatchDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchExportCloudPhoneData 导出云手机数据
//
// 批量导出云手机中的数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) BatchExportCloudPhoneData(request *model.BatchExportCloudPhoneDataRequest) (*model.BatchExportCloudPhoneDataResponse, error) {
	requestDef := GenReqDefForBatchExportCloudPhoneData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchExportCloudPhoneDataResponse), nil
	}
}

// BatchExportCloudPhoneDataInvoker 导出云手机数据
func (c *CphClient) BatchExportCloudPhoneDataInvoker(request *model.BatchExportCloudPhoneDataRequest) *BatchExportCloudPhoneDataInvoker {
	requestDef := GenReqDefForBatchExportCloudPhoneData()
	return &BatchExportCloudPhoneDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchImportCloudPhoneData 恢复云手机数据
//
// 导入数据到手机中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) BatchImportCloudPhoneData(request *model.BatchImportCloudPhoneDataRequest) (*model.BatchImportCloudPhoneDataResponse, error) {
	requestDef := GenReqDefForBatchImportCloudPhoneData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportCloudPhoneDataResponse), nil
	}
}

// BatchImportCloudPhoneDataInvoker 恢复云手机数据
func (c *CphClient) BatchImportCloudPhoneDataInvoker(request *model.BatchImportCloudPhoneDataRequest) *BatchImportCloudPhoneDataInvoker {
	requestDef := GenReqDefForBatchImportCloudPhoneData()
	return &BatchImportCloudPhoneDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchMigrateCloudPhone 迁移云手机
//
// 批量迁移整台云手机，包括云手机的系统盘数据和数据盘数据。该接口为异步接口，迁移完成的时间和手机的数据量有一定关系，整机数据大小为11G时，迁移时间大约为3-5分钟。迁移前请关闭云手机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) BatchMigrateCloudPhone(request *model.BatchMigrateCloudPhoneRequest) (*model.BatchMigrateCloudPhoneResponse, error) {
	requestDef := GenReqDefForBatchMigrateCloudPhone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchMigrateCloudPhoneResponse), nil
	}
}

// BatchMigrateCloudPhoneInvoker 迁移云手机
func (c *CphClient) BatchMigrateCloudPhoneInvoker(request *model.BatchMigrateCloudPhoneRequest) *BatchMigrateCloudPhoneInvoker {
	requestDef := GenReqDefForBatchMigrateCloudPhone()
	return &BatchMigrateCloudPhoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeCloudPhoneServerModel 变更云手机服务器规格
//
// 变更云手机服务器规格。只有能使用physical.rx1.xlarge.special私有规格的租户才可使用本接口。变更的目标规格也必须为特殊的规格才可变更。接口调用成功后，大约2分钟左右规格会变更结束，在订单中心可以查看到变更的订单状态为成功，且查询服务器的详细信息，可以查看到服务器规格名称已经变成新的规格名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ChangeCloudPhoneServerModel(request *model.ChangeCloudPhoneServerModelRequest) (*model.ChangeCloudPhoneServerModelResponse, error) {
	requestDef := GenReqDefForChangeCloudPhoneServerModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeCloudPhoneServerModelResponse), nil
	}
}

// ChangeCloudPhoneServerModelInvoker 变更云手机服务器规格
func (c *CphClient) ChangeCloudPhoneServerModelInvoker(request *model.ChangeCloudPhoneServerModelRequest) *ChangeCloudPhoneServerModelInvoker {
	requestDef := GenReqDefForChangeCloudPhoneServerModel()
	return &ChangeCloudPhoneServerModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCloudPhoneServer 购买系统定义网络云手机服务器
//
// 购买系统定义网络云手机服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) CreateCloudPhoneServer(request *model.CreateCloudPhoneServerRequest) (*model.CreateCloudPhoneServerResponse, error) {
	requestDef := GenReqDefForCreateCloudPhoneServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudPhoneServerResponse), nil
	}
}

// CreateCloudPhoneServerInvoker 购买系统定义网络云手机服务器
func (c *CphClient) CreateCloudPhoneServerInvoker(request *model.CreateCloudPhoneServerRequest) *CreateCloudPhoneServerInvoker {
	requestDef := GenReqDefForCreateCloudPhoneServer()
	return &CreateCloudPhoneServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNet2CloudPhoneServer 购买自定义网络云手机服务器
//
// 购买自定义网络的云手机服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) CreateNet2CloudPhoneServer(request *model.CreateNet2CloudPhoneServerRequest) (*model.CreateNet2CloudPhoneServerResponse, error) {
	requestDef := GenReqDefForCreateNet2CloudPhoneServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNet2CloudPhoneServerResponse), nil
	}
}

// CreateNet2CloudPhoneServerInvoker 购买自定义网络云手机服务器
func (c *CphClient) CreateNet2CloudPhoneServerInvoker(request *model.CreateNet2CloudPhoneServerRequest) *CreateNet2CloudPhoneServerInvoker {
	requestDef := GenReqDefForCreateNet2CloudPhoneServer()
	return &CreateNet2CloudPhoneServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteShareApps 删除共享应用
//
// 在共享应用存储目录中删除共享应用，该功能仅在支持共享应用的云手机规格上可实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) DeleteShareApps(request *model.DeleteShareAppsRequest) (*model.DeleteShareAppsResponse, error) {
	requestDef := GenReqDefForDeleteShareApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareAppsResponse), nil
	}
}

// DeleteShareAppsInvoker 删除共享应用
func (c *CphClient) DeleteShareAppsInvoker(request *model.DeleteShareAppsRequest) *DeleteShareAppsInvoker {
	requestDef := GenReqDefForDeleteShareApps()
	return &DeleteShareAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteShareFiles 删除共享存储文件
//
// 删除共享存储目录中文件，该功能仅在支持共享存储的云手机规格上可实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) DeleteShareFiles(request *model.DeleteShareFilesRequest) (*model.DeleteShareFilesResponse, error) {
	requestDef := GenReqDefForDeleteShareFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareFilesResponse), nil
	}
}

// DeleteShareFilesInvoker 删除共享存储文件
func (c *CphClient) DeleteShareFilesInvoker(request *model.DeleteShareFilesRequest) *DeleteShareFilesInvoker {
	requestDef := GenReqDefForDeleteShareFiles()
	return &DeleteShareFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportTraffic 云手机流量导流
//
// 手机流量路由修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ImportTraffic(request *model.ImportTrafficRequest) (*model.ImportTrafficResponse, error) {
	requestDef := GenReqDefForImportTraffic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportTrafficResponse), nil
	}
}

// ImportTrafficInvoker 云手机流量导流
func (c *CphClient) ImportTrafficInvoker(request *model.ImportTrafficRequest) *ImportTrafficInvoker {
	requestDef := GenReqDefForImportTraffic()
	return &ImportTrafficInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudPhoneImages 查询手机镜像
//
// 根据项目ID查询可用的手机镜像。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListCloudPhoneImages(request *model.ListCloudPhoneImagesRequest) (*model.ListCloudPhoneImagesResponse, error) {
	requestDef := GenReqDefForListCloudPhoneImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudPhoneImagesResponse), nil
	}
}

// ListCloudPhoneImagesInvoker 查询手机镜像
func (c *CphClient) ListCloudPhoneImagesInvoker(request *model.ListCloudPhoneImagesRequest) *ListCloudPhoneImagesInvoker {
	requestDef := GenReqDefForListCloudPhoneImages()
	return &ListCloudPhoneImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudPhoneModels 查询云手机规格列表
//
// 查询或统计云手机的规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListCloudPhoneModels(request *model.ListCloudPhoneModelsRequest) (*model.ListCloudPhoneModelsResponse, error) {
	requestDef := GenReqDefForListCloudPhoneModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudPhoneModelsResponse), nil
	}
}

// ListCloudPhoneModelsInvoker 查询云手机规格列表
func (c *CphClient) ListCloudPhoneModelsInvoker(request *model.ListCloudPhoneModelsRequest) *ListCloudPhoneModelsInvoker {
	requestDef := GenReqDefForListCloudPhoneModels()
	return &ListCloudPhoneModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudPhoneServerModels 查询云手机服务器规格列表
//
// 查询云手机服务器的规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListCloudPhoneServerModels(request *model.ListCloudPhoneServerModelsRequest) (*model.ListCloudPhoneServerModelsResponse, error) {
	requestDef := GenReqDefForListCloudPhoneServerModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudPhoneServerModelsResponse), nil
	}
}

// ListCloudPhoneServerModelsInvoker 查询云手机服务器规格列表
func (c *CphClient) ListCloudPhoneServerModelsInvoker(request *model.ListCloudPhoneServerModelsRequest) *ListCloudPhoneServerModelsInvoker {
	requestDef := GenReqDefForListCloudPhoneServerModels()
	return &ListCloudPhoneServerModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudPhoneServers 查询云手机服务器列表
//
// 分页查询云手机服务器，云手机服务器列表按照创建时间进行降序排列。分页查询可以指定offset以及limit。如果不存在云手机服务器，则返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListCloudPhoneServers(request *model.ListCloudPhoneServersRequest) (*model.ListCloudPhoneServersResponse, error) {
	requestDef := GenReqDefForListCloudPhoneServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudPhoneServersResponse), nil
	}
}

// ListCloudPhoneServersInvoker 查询云手机服务器列表
func (c *CphClient) ListCloudPhoneServersInvoker(request *model.ListCloudPhoneServersRequest) *ListCloudPhoneServersInvoker {
	requestDef := GenReqDefForListCloudPhoneServers()
	return &ListCloudPhoneServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudPhones 查询云手机列表
//
// 分页查询云手机，云手机列表按照创建时间进行降序排列。分页查询可以指定offset以及limit。如果不存在云手机，则返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListCloudPhones(request *model.ListCloudPhonesRequest) (*model.ListCloudPhonesResponse, error) {
	requestDef := GenReqDefForListCloudPhones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudPhonesResponse), nil
	}
}

// ListCloudPhonesInvoker 查询云手机列表
func (c *CphClient) ListCloudPhonesInvoker(request *model.ListCloudPhonesRequest) *ListCloudPhonesInvoker {
	requestDef := GenReqDefForListCloudPhones()
	return &ListCloudPhonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEncodeServers 查询编码服务
//
// 查询编码服务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListEncodeServers(request *model.ListEncodeServersRequest) (*model.ListEncodeServersResponse, error) {
	requestDef := GenReqDefForListEncodeServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEncodeServersResponse), nil
	}
}

// ListEncodeServersInvoker 查询编码服务
func (c *CphClient) ListEncodeServersInvoker(request *model.ListEncodeServersRequest) *ListEncodeServersInvoker {
	requestDef := GenReqDefForListEncodeServers()
	return &ListEncodeServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询任务执行状态列表
//
// 查询同一个request id下的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询任务执行状态列表
func (c *CphClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询租户在指定区域和资源类型的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *CphClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 查询资源实例
//
// 查询资源实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 查询资源实例
func (c *CphClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询资源标签
//
// 查询资源标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询资源标签
func (c *CphClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShareFiles 查询共享存储文件
//
// 查询共享存储指定路径下的文件列表，该功能仅在支持共享存储的云手机规格上可实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ListShareFiles(request *model.ListShareFilesRequest) (*model.ListShareFilesResponse, error) {
	requestDef := GenReqDefForListShareFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShareFilesResponse), nil
	}
}

// ListShareFilesInvoker 查询共享存储文件
func (c *CphClient) ListShareFilesInvoker(request *model.ListShareFilesRequest) *ListShareFilesInvoker {
	requestDef := GenReqDefForListShareFiles()
	return &ListShareFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushShareApps 推送共享应用
//
// 推送应用tar文件至共享应用存储目录中，该功能仅在支持共享应用的云手机规格上可实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) PushShareApps(request *model.PushShareAppsRequest) (*model.PushShareAppsResponse, error) {
	requestDef := GenReqDefForPushShareApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushShareAppsResponse), nil
	}
}

// PushShareAppsInvoker 推送共享应用
func (c *CphClient) PushShareAppsInvoker(request *model.PushShareAppsRequest) *PushShareAppsInvoker {
	requestDef := GenReqDefForPushShareApps()
	return &PushShareAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushShareFiles 推送共享存储文件
//
// 推送文件至共享存储目录中，该功能仅在支持共享存储的云手机规格上可实现。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) PushShareFiles(request *model.PushShareFilesRequest) (*model.PushShareFilesResponse, error) {
	requestDef := GenReqDefForPushShareFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushShareFilesResponse), nil
	}
}

// PushShareFilesInvoker 推送共享存储文件
func (c *CphClient) PushShareFilesInvoker(request *model.PushShareFilesRequest) *PushShareFilesInvoker {
	requestDef := GenReqDefForPushShareFiles()
	return &PushShareFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetCloudPhone 重置云手机
//
// 批量重置云手机，将云手机恢复出厂设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ResetCloudPhone(request *model.ResetCloudPhoneRequest) (*model.ResetCloudPhoneResponse, error) {
	requestDef := GenReqDefForResetCloudPhone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetCloudPhoneResponse), nil
	}
}

// ResetCloudPhoneInvoker 重置云手机
func (c *CphClient) ResetCloudPhoneInvoker(request *model.ResetCloudPhoneRequest) *ResetCloudPhoneInvoker {
	requestDef := GenReqDefForResetCloudPhone()
	return &ResetCloudPhoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCloudPhone 重启云手机
//
// 批量重启云手机，也可用于开启云手机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) RestartCloudPhone(request *model.RestartCloudPhoneRequest) (*model.RestartCloudPhoneResponse, error) {
	requestDef := GenReqDefForRestartCloudPhone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartCloudPhoneResponse), nil
	}
}

// RestartCloudPhoneInvoker 重启云手机
func (c *CphClient) RestartCloudPhoneInvoker(request *model.RestartCloudPhoneRequest) *RestartCloudPhoneInvoker {
	requestDef := GenReqDefForRestartCloudPhone()
	return &RestartCloudPhoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCloudPhoneServer 重启云手机服务器
//
// 批量重启云手机服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) RestartCloudPhoneServer(request *model.RestartCloudPhoneServerRequest) (*model.RestartCloudPhoneServerResponse, error) {
	requestDef := GenReqDefForRestartCloudPhoneServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartCloudPhoneServerResponse), nil
	}
}

// RestartCloudPhoneServerInvoker 重启云手机服务器
func (c *CphClient) RestartCloudPhoneServerInvoker(request *model.RestartCloudPhoneServerRequest) *RestartCloudPhoneServerInvoker {
	requestDef := GenReqDefForRestartCloudPhoneServer()
	return &RestartCloudPhoneServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartEncodeServer 重启编码服务
//
// 批量重启编码服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) RestartEncodeServer(request *model.RestartEncodeServerRequest) (*model.RestartEncodeServerResponse, error) {
	requestDef := GenReqDefForRestartEncodeServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartEncodeServerResponse), nil
	}
}

// RestartEncodeServerInvoker 重启编码服务
func (c *CphClient) RestartEncodeServerInvoker(request *model.RestartEncodeServerRequest) *RestartEncodeServerInvoker {
	requestDef := GenReqDefForRestartEncodeServer()
	return &RestartEncodeServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBandwidthDetail 查询带宽信息
//
// 查询云手机使用的带宽信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ShowBandwidthDetail(request *model.ShowBandwidthDetailRequest) (*model.ShowBandwidthDetailResponse, error) {
	requestDef := GenReqDefForShowBandwidthDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthDetailResponse), nil
	}
}

// ShowBandwidthDetailInvoker 查询带宽信息
func (c *CphClient) ShowBandwidthDetailInvoker(request *model.ShowBandwidthDetailRequest) *ShowBandwidthDetailInvoker {
	requestDef := GenReqDefForShowBandwidthDetail()
	return &ShowBandwidthDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCloudPhoneDetail 查询云手机详情
//
// 查询云手机的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ShowCloudPhoneDetail(request *model.ShowCloudPhoneDetailRequest) (*model.ShowCloudPhoneDetailResponse, error) {
	requestDef := GenReqDefForShowCloudPhoneDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudPhoneDetailResponse), nil
	}
}

// ShowCloudPhoneDetailInvoker 查询云手机详情
func (c *CphClient) ShowCloudPhoneDetailInvoker(request *model.ShowCloudPhoneDetailRequest) *ShowCloudPhoneDetailInvoker {
	requestDef := GenReqDefForShowCloudPhoneDetail()
	return &ShowCloudPhoneDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCloudPhoneServerDetail 查询云手机服务器详情
//
// 根据server_id查询云手机服务器的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ShowCloudPhoneServerDetail(request *model.ShowCloudPhoneServerDetailRequest) (*model.ShowCloudPhoneServerDetailResponse, error) {
	requestDef := GenReqDefForShowCloudPhoneServerDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudPhoneServerDetailResponse), nil
	}
}

// ShowCloudPhoneServerDetailInvoker 查询云手机服务器详情
func (c *CphClient) ShowCloudPhoneServerDetailInvoker(request *model.ShowCloudPhoneServerDetailRequest) *ShowCloudPhoneServerDetailInvoker {
	requestDef := GenReqDefForShowCloudPhoneServerDetail()
	return &ShowCloudPhoneServerDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询任务执行状态
//
// 查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询任务执行状态
func (c *CphClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopCloudPhone 关闭云手机
//
// 批量关闭云手机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) StopCloudPhone(request *model.StopCloudPhoneRequest) (*model.StopCloudPhoneResponse, error) {
	requestDef := GenReqDefForStopCloudPhone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopCloudPhoneResponse), nil
	}
}

// StopCloudPhoneInvoker 关闭云手机
func (c *CphClient) StopCloudPhoneInvoker(request *model.StopCloudPhoneRequest) *StopCloudPhoneInvoker {
	requestDef := GenReqDefForStopCloudPhone()
	return &StopCloudPhoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBandwidth 修改共享带宽
//
// 修改云手机使用的共享带宽大小。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UpdateBandwidth(request *model.UpdateBandwidthRequest) (*model.UpdateBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBandwidthResponse), nil
	}
}

// UpdateBandwidthInvoker 修改共享带宽
func (c *CphClient) UpdateBandwidthInvoker(request *model.UpdateBandwidthRequest) *UpdateBandwidthInvoker {
	requestDef := GenReqDefForUpdateBandwidth()
	return &UpdateBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCloudPhoneProperty 更新云手机属性
//
// 部分云手机属性开放更新能力，部分属性无法更新，部分属性需要重启手机生效，属性约束请云手机属性列表。如果手机处于异常状态，属性更新后需恢复手机状态为运行中才可生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UpdateCloudPhoneProperty(request *model.UpdateCloudPhonePropertyRequest) (*model.UpdateCloudPhonePropertyResponse, error) {
	requestDef := GenReqDefForUpdateCloudPhoneProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCloudPhonePropertyResponse), nil
	}
}

// UpdateCloudPhonePropertyInvoker 更新云手机属性
func (c *CphClient) UpdateCloudPhonePropertyInvoker(request *model.UpdateCloudPhonePropertyRequest) *UpdateCloudPhonePropertyInvoker {
	requestDef := GenReqDefForUpdateCloudPhoneProperty()
	return &UpdateCloudPhonePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeypair 更改密钥对
//
// 修改连接云手机的密钥对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UpdateKeypair(request *model.UpdateKeypairRequest) (*model.UpdateKeypairResponse, error) {
	requestDef := GenReqDefForUpdateKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeypairResponse), nil
	}
}

// UpdateKeypairInvoker 更改密钥对
func (c *CphClient) UpdateKeypairInvoker(request *model.UpdateKeypairRequest) *UpdateKeypairInvoker {
	requestDef := GenReqDefForUpdateKeypair()
	return &UpdateKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePhoneName 修改云手机名称
//
// 根据phoneId修改phoneName。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UpdatePhoneName(request *model.UpdatePhoneNameRequest) (*model.UpdatePhoneNameResponse, error) {
	requestDef := GenReqDefForUpdatePhoneName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePhoneNameResponse), nil
	}
}

// UpdatePhoneNameInvoker 修改云手机名称
func (c *CphClient) UpdatePhoneNameInvoker(request *model.UpdatePhoneNameRequest) *UpdatePhoneNameInvoker {
	requestDef := GenReqDefForUpdatePhoneName()
	return &UpdatePhoneNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerName 修改云手机服务器名称
//
// 根据serverId修改serverName。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UpdateServerName(request *model.UpdateServerNameRequest) (*model.UpdateServerNameResponse, error) {
	requestDef := GenReqDefForUpdateServerName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerNameResponse), nil
	}
}

// UpdateServerNameInvoker 修改云手机服务器名称
func (c *CphClient) UpdateServerNameInvoker(request *model.UpdateServerNameRequest) *UpdateServerNameInvoker {
	requestDef := GenReqDefForUpdateServerName()
	return &UpdateServerNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallApk 安装apk
//
// 在云手机中安装apk。系统会将指定的apk文件下载后直接安装到云手机中。
// 支持安装单apk应用和多apk应用。可使用install命令安装单apk应用，一次只支持安装一个apk；可使用install-multiple命令安装多apk应用（多apk应用为单个应用拆分成多个apk），一次只支持同一个应用的多个apk。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) InstallApk(request *model.InstallApkRequest) (*model.InstallApkResponse, error) {
	requestDef := GenReqDefForInstallApk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallApkResponse), nil
	}
}

// InstallApkInvoker 安装apk
func (c *CphClient) InstallApkInvoker(request *model.InstallApkRequest) *InstallApkInvoker {
	requestDef := GenReqDefForInstallApk()
	return &InstallApkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushFile 推送文件
//
// 推送文件到云手机文件系统中。系统会将所指定的文件下载解压后，将解压后的内容全部推送到云手机的根目录下。只支持指定tar格式的文件进行推送，您需要将tar文件提前上传至您的OBS桶中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) PushFile(request *model.PushFileRequest) (*model.PushFileResponse, error) {
	requestDef := GenReqDefForPushFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushFileResponse), nil
	}
}

// PushFileInvoker 推送文件
func (c *CphClient) PushFileInvoker(request *model.PushFileRequest) *PushFileInvoker {
	requestDef := GenReqDefForPushFile()
	return &PushFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunShellCommand 执行异步adb命令
//
// 在云手机中执行shell命令。该接口为异步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) RunShellCommand(request *model.RunShellCommandRequest) (*model.RunShellCommandResponse, error) {
	requestDef := GenReqDefForRunShellCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunShellCommandResponse), nil
	}
}

// RunShellCommandInvoker 执行异步adb命令
func (c *CphClient) RunShellCommandInvoker(request *model.RunShellCommandRequest) *RunShellCommandInvoker {
	requestDef := GenReqDefForRunShellCommand()
	return &RunShellCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSyncCommand 执行同步adb命令
//
// 在云手机中同步执行命令并返回命令执行的输出信息，该接口仅支持adb shell命令的执行。1分钟内每个用户调用接口次数上限为6次，每个云手机允许执行命令超时时间为2秒，接口时间不超过30秒，执行云手机数越多，接口耗时相应越长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) RunSyncCommand(request *model.RunSyncCommandRequest) (*model.RunSyncCommandResponse, error) {
	requestDef := GenReqDefForRunSyncCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSyncCommandResponse), nil
	}
}

// RunSyncCommandInvoker 执行同步adb命令
func (c *CphClient) RunSyncCommandInvoker(request *model.RunSyncCommandRequest) *RunSyncCommandInvoker {
	requestDef := GenReqDefForRunSyncCommand()
	return &RunSyncCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UninstallApk 卸载apk
//
// 在云手机中卸载apk。该接口为异步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CphClient) UninstallApk(request *model.UninstallApkRequest) (*model.UninstallApkResponse, error) {
	requestDef := GenReqDefForUninstallApk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UninstallApkResponse), nil
	}
}

// UninstallApkInvoker 卸载apk
func (c *CphClient) UninstallApkInvoker(request *model.UninstallApkRequest) *UninstallApkInvoker {
	requestDef := GenReqDefForUninstallApk()
	return &UninstallApkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
