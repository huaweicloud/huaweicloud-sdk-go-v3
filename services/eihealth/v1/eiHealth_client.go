package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v1/model"
)

type EiHealthClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEiHealthClient(hcClient *http_client.HcHttpClient) *EiHealthClient {
	return &EiHealthClient{HcClient: hcClient}
}

func EiHealthClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchImportApp 导入应用
//
// 批量导入应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchImportApp(request *model.BatchImportAppRequest) (*model.BatchImportAppResponse, error) {
	requestDef := GenReqDefForBatchImportApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportAppResponse), nil
	}
}

// BatchImportAppInvoker 导入应用
func (c *EiHealthClient) BatchImportAppInvoker(request *model.BatchImportAppRequest) *BatchImportAppInvoker {
	requestDef := GenReqDefForBatchImportApp()
	return &BatchImportAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用
func (c *EiHealthClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *EiHealthClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApp 获取应用列表
//
// 获取应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListApp(request *model.ListAppRequest) (*model.ListAppResponse, error) {
	requestDef := GenReqDefForListApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppResponse), nil
	}
}

// ListAppInvoker 获取应用列表
func (c *EiHealthClient) ListAppInvoker(request *model.ListAppRequest) *ListAppInvoker {
	requestDef := GenReqDefForListApp()
	return &ListAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 获取应用详情
//
// 获取应用详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 获取应用详情
func (c *EiHealthClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeApp 订阅应用
//
// 订阅应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeApp(request *model.SubscribeAppRequest) (*model.SubscribeAppResponse, error) {
	requestDef := GenReqDefForSubscribeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeAppResponse), nil
	}
}

// SubscribeAppInvoker 订阅应用
func (c *EiHealthClient) SubscribeAppInvoker(request *model.SubscribeAppRequest) *SubscribeAppInvoker {
	requestDef := GenReqDefForSubscribeApp()
	return &SubscribeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 更新应用
//
// 更新应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 更新应用
func (c *EiHealthClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAsset 获取资产列表
//
// 获取资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListAsset(request *model.ListAssetRequest) (*model.ListAssetResponse, error) {
	requestDef := GenReqDefForListAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetResponse), nil
	}
}

// ListAssetInvoker 获取资产列表
func (c *EiHealthClient) ListAssetInvoker(request *model.ListAssetRequest) *ListAssetInvoker {
	requestDef := GenReqDefForListAsset()
	return &ListAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProperty 获取属性值列表
//
// 获取属性值列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListProperty(request *model.ListPropertyRequest) (*model.ListPropertyResponse, error) {
	requestDef := GenReqDefForListProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropertyResponse), nil
	}
}

// ListPropertyInvoker 获取属性值列表
func (c *EiHealthClient) ListPropertyInvoker(request *model.ListPropertyRequest) *ListPropertyInvoker {
	requestDef := GenReqDefForListProperty()
	return &ListPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsset 查询资产详情
//
// 查询资产详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 查询资产详情
func (c *EiHealthClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetVersion 查询资产版本详情
//
// 查询资产版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAssetVersion(request *model.ShowAssetVersionRequest) (*model.ShowAssetVersionResponse, error) {
	requestDef := GenReqDefForShowAssetVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetVersionResponse), nil
	}
}

// ShowAssetVersionInvoker 查询资产版本详情
func (c *EiHealthClient) ShowAssetVersionInvoker(request *model.ShowAssetVersionRequest) *ShowAssetVersionInvoker {
	requestDef := GenReqDefForShowAssetVersion()
	return &ShowAssetVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAutoJob 创建自动作业模板
//
// 创建自动作业模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateAutoJob(request *model.CreateAutoJobRequest) (*model.CreateAutoJobResponse, error) {
	requestDef := GenReqDefForCreateAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoJobResponse), nil
	}
}

// CreateAutoJobInvoker 创建自动作业模板
func (c *EiHealthClient) CreateAutoJobInvoker(request *model.CreateAutoJobRequest) *CreateAutoJobInvoker {
	requestDef := GenReqDefForCreateAutoJob()
	return &CreateAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAutoJob 删除自动作业模板
//
// 删除自动作业模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteAutoJob(request *model.DeleteAutoJobRequest) (*model.DeleteAutoJobResponse, error) {
	requestDef := GenReqDefForDeleteAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAutoJobResponse), nil
	}
}

// DeleteAutoJobInvoker 删除自动作业模板
func (c *EiHealthClient) DeleteAutoJobInvoker(request *model.DeleteAutoJobRequest) *DeleteAutoJobInvoker {
	requestDef := GenReqDefForDeleteAutoJob()
	return &DeleteAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoJob 获取自动作业模板列表
//
// 获取自动作业模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListAutoJob(request *model.ListAutoJobRequest) (*model.ListAutoJobResponse, error) {
	requestDef := GenReqDefForListAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoJobResponse), nil
	}
}

// ListAutoJobInvoker 获取自动作业模板列表
func (c *EiHealthClient) ListAutoJobInvoker(request *model.ListAutoJobRequest) *ListAutoJobInvoker {
	requestDef := GenReqDefForListAutoJob()
	return &ListAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoJob 查询自动作业模板
//
// 查询自动作业模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAutoJob(request *model.ShowAutoJobRequest) (*model.ShowAutoJobResponse, error) {
	requestDef := GenReqDefForShowAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoJobResponse), nil
	}
}

// ShowAutoJobInvoker 查询自动作业模板
func (c *EiHealthClient) ShowAutoJobInvoker(request *model.ShowAutoJobRequest) *ShowAutoJobInvoker {
	requestDef := GenReqDefForShowAutoJob()
	return &ShowAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAutoJob 启动自动作业
//
// 启动自动作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StartAutoJob(request *model.StartAutoJobRequest) (*model.StartAutoJobResponse, error) {
	requestDef := GenReqDefForStartAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAutoJobResponse), nil
	}
}

// StartAutoJobInvoker 启动自动作业
func (c *EiHealthClient) StartAutoJobInvoker(request *model.StartAutoJobRequest) *StartAutoJobInvoker {
	requestDef := GenReqDefForStartAutoJob()
	return &StartAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopAutoJob 停止自动作业
//
// 停止自动作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StopAutoJob(request *model.StopAutoJobRequest) (*model.StopAutoJobResponse, error) {
	requestDef := GenReqDefForStopAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAutoJobResponse), nil
	}
}

// StopAutoJobInvoker 停止自动作业
func (c *EiHealthClient) StopAutoJobInvoker(request *model.StopAutoJobRequest) *StopAutoJobInvoker {
	requestDef := GenReqDefForStopAutoJob()
	return &StopAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoJob 更新自动作业模板
//
// 更新自动作业模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateAutoJob(request *model.UpdateAutoJobRequest) (*model.UpdateAutoJobResponse, error) {
	requestDef := GenReqDefForUpdateAutoJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoJobResponse), nil
	}
}

// UpdateAutoJobInvoker 更新自动作业模板
func (c *EiHealthClient) UpdateAutoJobInvoker(request *model.UpdateAutoJobRequest) *UpdateAutoJobInvoker {
	requestDef := GenReqDefForUpdateAutoJob()
	return &UpdateAutoJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComputingResource 购买计算资源
//
// 购买计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateComputingResource(request *model.CreateComputingResourceRequest) (*model.CreateComputingResourceResponse, error) {
	requestDef := GenReqDefForCreateComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComputingResourceResponse), nil
	}
}

// CreateComputingResourceInvoker 购买计算资源
func (c *EiHealthClient) CreateComputingResourceInvoker(request *model.CreateComputingResourceRequest) *CreateComputingResourceInvoker {
	requestDef := GenReqDefForCreateComputingResource()
	return &CreateComputingResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComputingResource 删除计算资源
//
// 删除计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteComputingResource(request *model.DeleteComputingResourceRequest) (*model.DeleteComputingResourceResponse, error) {
	requestDef := GenReqDefForDeleteComputingResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComputingResourceResponse), nil
	}
}

// DeleteComputingResourceInvoker 删除计算资源
func (c *EiHealthClient) DeleteComputingResourceInvoker(request *model.DeleteComputingResourceRequest) *DeleteComputingResourceInvoker {
	requestDef := GenReqDefForDeleteComputingResource()
	return &DeleteComputingResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComputingResourceFlavors 查询计算资源规格
//
// 查询计算资源规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListComputingResourceFlavors(request *model.ListComputingResourceFlavorsRequest) (*model.ListComputingResourceFlavorsResponse, error) {
	requestDef := GenReqDefForListComputingResourceFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComputingResourceFlavorsResponse), nil
	}
}

// ListComputingResourceFlavorsInvoker 查询计算资源规格
func (c *EiHealthClient) ListComputingResourceFlavorsInvoker(request *model.ListComputingResourceFlavorsRequest) *ListComputingResourceFlavorsInvoker {
	requestDef := GenReqDefForListComputingResourceFlavors()
	return &ListComputingResourceFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComputingResources 查询计算资源
//
// 查询计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListComputingResources(request *model.ListComputingResourcesRequest) (*model.ListComputingResourcesResponse, error) {
	requestDef := GenReqDefForListComputingResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComputingResourcesResponse), nil
	}
}

// ListComputingResourcesInvoker 查询计算资源
func (c *EiHealthClient) ListComputingResourcesInvoker(request *model.ListComputingResourcesRequest) *ListComputingResourcesInvoker {
	requestDef := GenReqDefForListComputingResources()
	return &ListComputingResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootNode 重启计算资源
//
// 重启计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RebootNode(request *model.RebootNodeRequest) (*model.RebootNodeResponse, error) {
	requestDef := GenReqDefForRebootNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootNodeResponse), nil
	}
}

// RebootNodeInvoker 重启计算资源
func (c *EiHealthClient) RebootNodeInvoker(request *model.RebootNodeRequest) *RebootNodeInvoker {
	requestDef := GenReqDefForRebootNode()
	return &RebootNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBmsDevices 查询bms计算资源显卡id列表
//
// 查询bms计算资源显卡id列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowBmsDevices(request *model.ShowBmsDevicesRequest) (*model.ShowBmsDevicesResponse, error) {
	requestDef := GenReqDefForShowBmsDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBmsDevicesResponse), nil
	}
}

// ShowBmsDevicesInvoker 查询bms计算资源显卡id列表
func (c *EiHealthClient) ShowBmsDevicesInvoker(request *model.ShowBmsDevicesRequest) *ShowBmsDevicesInvoker {
	requestDef := GenReqDefForShowBmsDevices()
	return &ShowBmsDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvsQuota 获取EVS配额及使用情况
//
// 获取EVS配额及使用情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowEvsQuota(request *model.ShowEvsQuotaRequest) (*model.ShowEvsQuotaResponse, error) {
	requestDef := GenReqDefForShowEvsQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvsQuotaResponse), nil
	}
}

// ShowEvsQuotaInvoker 获取EVS配额及使用情况
func (c *EiHealthClient) ShowEvsQuotaInvoker(request *model.ShowEvsQuotaRequest) *ShowEvsQuotaInvoker {
	requestDef := GenReqDefForShowEvsQuota()
	return &ShowEvsQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLeftQuota 获取节点剩余配额
//
// 获取节点剩余配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowLeftQuota(request *model.ShowLeftQuotaRequest) (*model.ShowLeftQuotaResponse, error) {
	requestDef := GenReqDefForShowLeftQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLeftQuotaResponse), nil
	}
}

// ShowLeftQuotaInvoker 获取节点剩余配额
func (c *EiHealthClient) ShowLeftQuotaInvoker(request *model.ShowLeftQuotaRequest) *ShowLeftQuotaInvoker {
	requestDef := GenReqDefForShowLeftQuota()
	return &ShowLeftQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSchedule 查询计算资源调度信息
//
// 查询计算资源调度信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowSchedule(request *model.ShowScheduleRequest) (*model.ShowScheduleResponse, error) {
	requestDef := GenReqDefForShowSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScheduleResponse), nil
	}
}

// ShowScheduleInvoker 查询计算资源调度信息
func (c *EiHealthClient) ShowScheduleInvoker(request *model.ShowScheduleRequest) *ShowScheduleInvoker {
	requestDef := GenReqDefForShowSchedule()
	return &ShowScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartNode 启动计算资源
//
// 启动计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StartNode(request *model.StartNodeRequest) (*model.StartNodeResponse, error) {
	requestDef := GenReqDefForStartNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNodeResponse), nil
	}
}

// StartNodeInvoker 启动计算资源
func (c *EiHealthClient) StartNodeInvoker(request *model.StartNodeRequest) *StartNodeInvoker {
	requestDef := GenReqDefForStartNode()
	return &StartNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopNode 关闭计算资源
//
// 关闭计算资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StopNode(request *model.StopNodeRequest) (*model.StopNodeResponse, error) {
	requestDef := GenReqDefForStopNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopNodeResponse), nil
	}
}

// StopNodeInvoker 关闭计算资源
func (c *EiHealthClient) StopNodeInvoker(request *model.StopNodeRequest) *StopNodeInvoker {
	requestDef := GenReqDefForStopNode()
	return &StopNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSchedule 修改计算资源调度信息
//
// 修改计算资源调度信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateSchedule(request *model.UpdateScheduleRequest) (*model.UpdateScheduleResponse, error) {
	requestDef := GenReqDefForUpdateSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduleResponse), nil
	}
}

// UpdateScheduleInvoker 修改计算资源调度信息
func (c *EiHealthClient) UpdateScheduleInvoker(request *model.UpdateScheduleRequest) *UpdateScheduleInvoker {
	requestDef := GenReqDefForUpdateSchedule()
	return &UpdateScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBackup 归档数据
//
// 将需要归档的重要数据拷贝到数据归档桶
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateBackup(request *model.CreateBackupRequest) (*model.CreateBackupResponse, error) {
	requestDef := GenReqDefForCreateBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackupResponse), nil
	}
}

// CreateBackupInvoker 归档数据
func (c *EiHealthClient) CreateBackupInvoker(request *model.CreateBackupRequest) *CreateBackupInvoker {
	requestDef := GenReqDefForCreateBackup()
	return &CreateBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup 删除归档
//
// 删除指定的归档
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// DeleteBackupInvoker 删除归档
func (c *EiHealthClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackup 查询归档列表
//
// 分页查询用户管理的项目的所有历史归档记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListBackup(request *model.ListBackupRequest) (*model.ListBackupResponse, error) {
	requestDef := GenReqDefForListBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupResponse), nil
	}
}

// ListBackupInvoker 查询归档列表
func (c *EiHealthClient) ListBackupInvoker(request *model.ListBackupRequest) *ListBackupInvoker {
	requestDef := GenReqDefForListBackup()
	return &ListBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreBackup 恢复归档
//
// 将指定的归档数据拷贝到目标项目的某个目录下
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RestoreBackup(request *model.RestoreBackupRequest) (*model.RestoreBackupResponse, error) {
	requestDef := GenReqDefForRestoreBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreBackupResponse), nil
	}
}

// RestoreBackupInvoker 恢复归档
func (c *EiHealthClient) RestoreBackupInvoker(request *model.RestoreBackupRequest) *RestoreBackupInvoker {
	requestDef := GenReqDefForRestoreBackup()
	return &RestoreBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPath 获取指定归档的全数据清单
//
// 根据归档ID获取该归档的全数据清单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowBackupPath(request *model.ShowBackupPathRequest) (*model.ShowBackupPathResponse, error) {
	requestDef := GenReqDefForShowBackupPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPathResponse), nil
	}
}

// ShowBackupPathInvoker 获取指定归档的全数据清单
func (c *EiHealthClient) ShowBackupPathInvoker(request *model.ShowBackupPathRequest) *ShowBackupPathInvoker {
	requestDef := GenReqDefForShowBackupPath()
	return &ShowBackupPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyData 复制项目数据
//
// 复制项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CopyData(request *model.CopyDataRequest) (*model.CopyDataResponse, error) {
	requestDef := GenReqDefForCopyData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyDataResponse), nil
	}
}

// CopyDataInvoker 复制项目数据
func (c *EiHealthClient) CopyDataInvoker(request *model.CopyDataRequest) *CopyDataInvoker {
	requestDef := GenReqDefForCopyData()
	return &CopyDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateData 创建文件夹
//
// 创建文件夹
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateData(request *model.CreateDataRequest) (*model.CreateDataResponse, error) {
	requestDef := GenReqDefForCreateData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataResponse), nil
	}
}

// CreateDataInvoker 创建文件夹
func (c *EiHealthClient) CreateDataInvoker(request *model.CreateDataRequest) *CreateDataInvoker {
	requestDef := GenReqDefForCreateData()
	return &CreateDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteData 批量删除项目数据
//
// 批量删除项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteData(request *model.DeleteDataRequest) (*model.DeleteDataResponse, error) {
	requestDef := GenReqDefForDeleteData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataResponse), nil
	}
}

// DeleteDataInvoker 批量删除项目数据
func (c *EiHealthClient) DeleteDataInvoker(request *model.DeleteDataRequest) *DeleteDataInvoker {
	requestDef := GenReqDefForDeleteData()
	return &DeleteDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportData 导入项目数据
//
// 导入项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportData(request *model.ImportDataRequest) (*model.ImportDataResponse, error) {
	requestDef := GenReqDefForImportData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataResponse), nil
	}
}

// ImportDataInvoker 导入项目数据
func (c *EiHealthClient) ImportDataInvoker(request *model.ImportDataRequest) *ImportDataInvoker {
	requestDef := GenReqDefForImportData()
	return &ImportDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportNetworkData 导入网上数据
//
// 导入网上数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportNetworkData(request *model.ImportNetworkDataRequest) (*model.ImportNetworkDataResponse, error) {
	requestDef := GenReqDefForImportNetworkData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportNetworkDataResponse), nil
	}
}

// ImportNetworkDataInvoker 导入网上数据
func (c *EiHealthClient) ImportNetworkDataInvoker(request *model.ImportNetworkDataRequest) *ImportNetworkDataInvoker {
	requestDef := GenReqDefForImportNetworkData()
	return &ImportNetworkDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBucket 获取桶列表
//
// 获取桶列表(包含当前项目桶和引用项目桶)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListBucket(request *model.ListBucketRequest) (*model.ListBucketResponse, error) {
	requestDef := GenReqDefForListBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBucketResponse), nil
	}
}

// ListBucketInvoker 获取桶列表
func (c *EiHealthClient) ListBucketInvoker(request *model.ListBucketRequest) *ListBucketInvoker {
	requestDef := GenReqDefForListBucket()
	return &ListBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListData 查询数据列表
//
// 查询指定目录下的数据列表，如果不指定默认查询根目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListData(request *model.ListDataRequest) (*model.ListDataResponse, error) {
	requestDef := GenReqDefForListData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataResponse), nil
	}
}

// ListDataInvoker 查询数据列表
func (c *EiHealthClient) ListDataInvoker(request *model.ListDataRequest) *ListDataInvoker {
	requestDef := GenReqDefForListData()
	return &ListDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// QuoteData 引用项目数据
//
// 引用项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) QuoteData(request *model.QuoteDataRequest) (*model.QuoteDataResponse, error) {
	requestDef := GenReqDefForQuoteData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.QuoteDataResponse), nil
	}
}

// QuoteDataInvoker 引用项目数据
func (c *EiHealthClient) QuoteDataInvoker(request *model.QuoteDataRequest) *QuoteDataInvoker {
	requestDef := GenReqDefForQuoteData()
	return &QuoteDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBucketStorage 获取桶存量信息
//
// 获取桶存量信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowBucketStorage(request *model.ShowBucketStorageRequest) (*model.ShowBucketStorageResponse, error) {
	requestDef := GenReqDefForShowBucketStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBucketStorageResponse), nil
	}
}

// ShowBucketStorageInvoker 获取桶存量信息
func (c *EiHealthClient) ShowBucketStorageInvoker(request *model.ShowBucketStorageRequest) *ShowBucketStorageInvoker {
	requestDef := GenReqDefForShowBucketStorage()
	return &ShowBucketStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowData 获取指定数据对象
//
// 获取指定数据对象的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowData(request *model.ShowDataRequest) (*model.ShowDataResponse, error) {
	requestDef := GenReqDefForShowData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataResponse), nil
	}
}

// ShowDataInvoker 获取指定数据对象
func (c *EiHealthClient) ShowDataInvoker(request *model.ShowDataRequest) *ShowDataInvoker {
	requestDef := GenReqDefForShowData()
	return &ShowDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataPolicy 查询项目级数据权限控制策略
//
// 查询项目级数据权限控制策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDataPolicy(request *model.ShowDataPolicyRequest) (*model.ShowDataPolicyResponse, error) {
	requestDef := GenReqDefForShowDataPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataPolicyResponse), nil
	}
}

// ShowDataPolicyInvoker 查询项目级数据权限控制策略
func (c *EiHealthClient) ShowDataPolicyInvoker(request *model.ShowDataPolicyRequest) *ShowDataPolicyInvoker {
	requestDef := GenReqDefForShowDataPolicy()
	return &ShowDataPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeData 订阅资产市场数据
//
// 订阅资产市场数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeData(request *model.SubscribeDataRequest) (*model.SubscribeDataResponse, error) {
	requestDef := GenReqDefForSubscribeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeDataResponse), nil
	}
}

// SubscribeDataInvoker 订阅资产市场数据
func (c *EiHealthClient) SubscribeDataInvoker(request *model.SubscribeDataRequest) *SubscribeDataInvoker {
	requestDef := GenReqDefForSubscribeData()
	return &SubscribeDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataPolicy 设置项目级权限控制策略
//
// 设置项目级权限控制策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateDataPolicy(request *model.UpdateDataPolicyRequest) (*model.UpdateDataPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDataPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataPolicyResponse), nil
	}
}

// UpdateDataPolicyInvoker 设置项目级权限控制策略
func (c *EiHealthClient) UpdateDataPolicyInvoker(request *model.UpdateDataPolicyRequest) *UpdateDataPolicyInvoker {
	requestDef := GenReqDefForUpdateDataPolicy()
	return &UpdateDataPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadData 上传数据文件
//
// 上传数据文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UploadData(request *model.UploadDataRequest) (*model.UploadDataResponse, error) {
	requestDef := GenReqDefForUploadData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadDataResponse), nil
	}
}

// UploadDataInvoker 上传数据文件
func (c *EiHealthClient) UploadDataInvoker(request *model.UploadDataRequest) *UploadDataInvoker {
	requestDef := GenReqDefForUploadData()
	return &UploadDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelDataJob 取消数据作业
//
// 取消数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CancelDataJob(request *model.CancelDataJobRequest) (*model.CancelDataJobResponse, error) {
	requestDef := GenReqDefForCancelDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelDataJobResponse), nil
	}
}

// CancelDataJobInvoker 取消数据作业
func (c *EiHealthClient) CancelDataJobInvoker(request *model.CancelDataJobRequest) *CancelDataJobInvoker {
	requestDef := GenReqDefForCancelDataJob()
	return &CancelDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataJob 删除数据作业
//
// 删除数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDataJob(request *model.DeleteDataJobRequest) (*model.DeleteDataJobResponse, error) {
	requestDef := GenReqDefForDeleteDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataJobResponse), nil
	}
}

// DeleteDataJobInvoker 删除数据作业
func (c *EiHealthClient) DeleteDataJobInvoker(request *model.DeleteDataJobRequest) *DeleteDataJobInvoker {
	requestDef := GenReqDefForDeleteDataJob()
	return &DeleteDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadDataJobLog 下载数据作业执行日志
//
// 下载数据作业执行日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DownloadDataJobLog(request *model.DownloadDataJobLogRequest) (*model.DownloadDataJobLogResponse, error) {
	requestDef := GenReqDefForDownloadDataJobLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDataJobLogResponse), nil
	}
}

// DownloadDataJobLogInvoker 下载数据作业执行日志
func (c *EiHealthClient) DownloadDataJobLogInvoker(request *model.DownloadDataJobLogRequest) *DownloadDataJobLogInvoker {
	requestDef := GenReqDefForDownloadDataJobLog()
	return &DownloadDataJobLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCheckpoint 获取数据作业执行日志
//
// 获取数据作业执行日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListCheckpoint(request *model.ListCheckpointRequest) (*model.ListCheckpointResponse, error) {
	requestDef := GenReqDefForListCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCheckpointResponse), nil
	}
}

// ListCheckpointInvoker 获取数据作业执行日志
func (c *EiHealthClient) ListCheckpointInvoker(request *model.ListCheckpointRequest) *ListCheckpointInvoker {
	requestDef := GenReqDefForListCheckpoint()
	return &ListCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataJob 获取数据作业列表
//
// 获取数据作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDataJob(request *model.ListDataJobRequest) (*model.ListDataJobResponse, error) {
	requestDef := GenReqDefForListDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataJobResponse), nil
	}
}

// ListDataJobInvoker 获取数据作业列表
func (c *EiHealthClient) ListDataJobInvoker(request *model.ListDataJobRequest) *ListDataJobInvoker {
	requestDef := GenReqDefForListDataJob()
	return &ListDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryDataJob 重试数据作业
//
// 重试数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RetryDataJob(request *model.RetryDataJobRequest) (*model.RetryDataJobResponse, error) {
	requestDef := GenReqDefForRetryDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryDataJobResponse), nil
	}
}

// RetryDataJobInvoker 重试数据作业
func (c *EiHealthClient) RetryDataJobInvoker(request *model.RetryDataJobRequest) *RetryDataJobInvoker {
	requestDef := GenReqDefForRetryDataJob()
	return &RetryDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataJob 获取数据作业详细信息
//
// 获取数据作业详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDataJob(request *model.ShowDataJobRequest) (*model.ShowDataJobResponse, error) {
	requestDef := GenReqDefForShowDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataJobResponse), nil
	}
}

// ShowDataJobInvoker 获取数据作业详细信息
func (c *EiHealthClient) ShowDataJobInvoker(request *model.ShowDataJobRequest) *ShowDataJobInvoker {
	requestDef := GenReqDefForShowDataJob()
	return &ShowDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseData 插入单条数据
//
// 插入单条数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDatabaseData(request *model.CreateDatabaseDataRequest) (*model.CreateDatabaseDataResponse, error) {
	requestDef := GenReqDefForCreateDatabaseData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseDataResponse), nil
	}
}

// CreateDatabaseDataInvoker 插入单条数据
func (c *EiHealthClient) CreateDatabaseDataInvoker(request *model.CreateDatabaseDataRequest) *CreateDatabaseDataInvoker {
	requestDef := GenReqDefForCreateDatabaseData()
	return &CreateDatabaseDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建数据库实例
//
// 创建数据库实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建数据库实例
func (c *EiHealthClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabaseData 删除数据
//
// 删除指定行数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDatabaseData(request *model.DeleteDatabaseDataRequest) (*model.DeleteDatabaseDataResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseDataResponse), nil
	}
}

// DeleteDatabaseDataInvoker 删除数据
func (c *EiHealthClient) DeleteDatabaseDataInvoker(request *model.DeleteDatabaseDataRequest) *DeleteDatabaseDataInvoker {
	requestDef := GenReqDefForDeleteDatabaseData()
	return &DeleteDatabaseDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *EiHealthClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportDatabaseData 导入数据
//
// 导入数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportDatabaseData(request *model.ImportDatabaseDataRequest) (*model.ImportDatabaseDataResponse, error) {
	requestDef := GenReqDefForImportDatabaseData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDatabaseDataResponse), nil
	}
}

// ImportDatabaseDataInvoker 导入数据
func (c *EiHealthClient) ImportDatabaseDataInvoker(request *model.ImportDatabaseDataRequest) *ImportDatabaseDataInvoker {
	requestDef := GenReqDefForImportDatabaseData()
	return &ImportDatabaseDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseData 查询数据
//
// 查询数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDatabaseData(request *model.ListDatabaseDataRequest) (*model.ListDatabaseDataResponse, error) {
	requestDef := GenReqDefForListDatabaseData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseDataResponse), nil
	}
}

// ListDatabaseDataInvoker 查询数据
func (c *EiHealthClient) ListDatabaseDataInvoker(request *model.ListDatabaseDataRequest) *ListDatabaseDataInvoker {
	requestDef := GenReqDefForListDatabaseData()
	return &ListDatabaseDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstance 获取实例列表
//
// 获取实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListInstance(request *model.ListInstanceRequest) (*model.ListInstanceResponse, error) {
	requestDef := GenReqDefForListInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResponse), nil
	}
}

// ListInstanceInvoker 获取实例列表
func (c *EiHealthClient) ListInstanceInvoker(request *model.ListInstanceRequest) *ListInstanceInvoker {
	requestDef := GenReqDefForListInstance()
	return &ListInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// QuoteInstance 引用数据库实例
//
// 引用数据库实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) QuoteInstance(request *model.QuoteInstanceRequest) (*model.QuoteInstanceResponse, error) {
	requestDef := GenReqDefForQuoteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.QuoteInstanceResponse), nil
	}
}

// QuoteInstanceInvoker 引用数据库实例
func (c *EiHealthClient) QuoteInstanceInvoker(request *model.QuoteInstanceRequest) *QuoteInstanceInvoker {
	requestDef := GenReqDefForQuoteInstance()
	return &QuoteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询实例详情
//
// 查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询实例详情
func (c *EiHealthClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatabaseData 更新数据
//
// 更新数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateDatabaseData(request *model.UpdateDatabaseDataRequest) (*model.UpdateDatabaseDataResponse, error) {
	requestDef := GenReqDefForUpdateDatabaseData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseDataResponse), nil
	}
}

// UpdateDatabaseDataInvoker 更新数据
func (c *EiHealthClient) UpdateDatabaseDataInvoker(request *model.UpdateDatabaseDataRequest) *UpdateDatabaseDataInvoker {
	requestDef := GenReqDefForUpdateDatabaseData()
	return &UpdateDatabaseDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseResource 购买数据库资源
//
// 购买数据库资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDatabaseResource(request *model.CreateDatabaseResourceRequest) (*model.CreateDatabaseResourceResponse, error) {
	requestDef := GenReqDefForCreateDatabaseResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResourceResponse), nil
	}
}

// CreateDatabaseResourceInvoker 购买数据库资源
func (c *EiHealthClient) CreateDatabaseResourceInvoker(request *model.CreateDatabaseResourceRequest) *CreateDatabaseResourceInvoker {
	requestDef := GenReqDefForCreateDatabaseResource()
	return &CreateDatabaseResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabaseResource 删除数据库资源
//
// 删除数据库资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDatabaseResource(request *model.DeleteDatabaseResourceRequest) (*model.DeleteDatabaseResourceResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResourceResponse), nil
	}
}

// DeleteDatabaseResourceInvoker 删除数据库资源
func (c *EiHealthClient) DeleteDatabaseResourceInvoker(request *model.DeleteDatabaseResourceRequest) *DeleteDatabaseResourceInvoker {
	requestDef := GenReqDefForDeleteDatabaseResource()
	return &DeleteDatabaseResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseResource 查询数据库资源
//
// 查询数据库资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDatabaseResource(request *model.ListDatabaseResourceRequest) (*model.ListDatabaseResourceResponse, error) {
	requestDef := GenReqDefForListDatabaseResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseResourceResponse), nil
	}
}

// ListDatabaseResourceInvoker 查询数据库资源
func (c *EiHealthClient) ListDatabaseResourceInvoker(request *model.ListDatabaseResourceRequest) *ListDatabaseResourceInvoker {
	requestDef := GenReqDefForListDatabaseResource()
	return &ListDatabaseResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseResourceFlavor 获取数据库资源规格列表
//
// 获取数据库资源规格列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDatabaseResourceFlavor(request *model.ListDatabaseResourceFlavorRequest) (*model.ListDatabaseResourceFlavorResponse, error) {
	requestDef := GenReqDefForListDatabaseResourceFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseResourceFlavorResponse), nil
	}
}

// ListDatabaseResourceFlavorInvoker 获取数据库资源规格列表
func (c *EiHealthClient) ListDatabaseResourceFlavorInvoker(request *model.ListDatabaseResourceFlavorRequest) *ListDatabaseResourceFlavorInvoker {
	requestDef := GenReqDefForListDatabaseResourceFlavor()
	return &ListDatabaseResourceFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTag 批量删除镜像tag
//
// 批量删除镜像tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchDeleteTag(request *model.BatchDeleteTagRequest) (*model.BatchDeleteTagResponse, error) {
	requestDef := GenReqDefForBatchDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagResponse), nil
	}
}

// BatchDeleteTagInvoker 批量删除镜像tag
func (c *EiHealthClient) BatchDeleteTagInvoker(request *model.BatchDeleteTagRequest) *BatchDeleteTagInvoker {
	requestDef := GenReqDefForBatchDeleteTag()
	return &BatchDeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImage 创建镜像
//
// 创建镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

// CreateImageInvoker 创建镜像
func (c *EiHealthClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImage 删除镜像仓库
//
// 删除镜像仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteImage(request *model.DeleteImageRequest) (*model.DeleteImageResponse, error) {
	requestDef := GenReqDefForDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageResponse), nil
	}
}

// DeleteImageInvoker 删除镜像仓库
func (c *EiHealthClient) DeleteImageInvoker(request *model.DeleteImageRequest) *DeleteImageInvoker {
	requestDef := GenReqDefForDeleteImage()
	return &DeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除指定镜像tag
//
// 删除指定镜像tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除指定镜像tag
func (c *EiHealthClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportImage 导入镜像
//
// 导入镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportImage(request *model.ImportImageRequest) (*model.ImportImageResponse, error) {
	requestDef := GenReqDefForImportImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportImageResponse), nil
	}
}

// ImportImageInvoker 导入镜像
func (c *EiHealthClient) ImportImageInvoker(request *model.ImportImageRequest) *ImportImageInvoker {
	requestDef := GenReqDefForImportImage()
	return &ImportImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImage 获取镜像列表
//
// 获取镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListImage(request *model.ListImageRequest) (*model.ListImageResponse, error) {
	requestDef := GenReqDefForListImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageResponse), nil
	}
}

// ListImageInvoker 获取镜像列表
func (c *EiHealthClient) ListImageInvoker(request *model.ListImageRequest) *ListImageInvoker {
	requestDef := GenReqDefForListImage()
	return &ListImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageTag 获取指定镜像的tag列表
//
// 获取指定镜像的tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListImageTag(request *model.ListImageTagRequest) (*model.ListImageTagResponse, error) {
	requestDef := GenReqDefForListImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageTagResponse), nil
	}
}

// ListImageTagInvoker 获取指定镜像的tag列表
func (c *EiHealthClient) ListImageTagInvoker(request *model.ListImageTagRequest) *ListImageTagInvoker {
	requestDef := GenReqDefForListImageTag()
	return &ListImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDockerLogin 获取docker login指令
//
// 获取docker login指令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDockerLogin(request *model.ShowDockerLoginRequest) (*model.ShowDockerLoginResponse, error) {
	requestDef := GenReqDefForShowDockerLogin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDockerLoginResponse), nil
	}
}

// ShowDockerLoginInvoker 获取docker login指令
func (c *EiHealthClient) ShowDockerLoginInvoker(request *model.ShowDockerLoginRequest) *ShowDockerLoginInvoker {
	requestDef := GenReqDefForShowDockerLogin()
	return &ShowDockerLoginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeImage 订阅镜像
//
// 订阅镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeImage(request *model.SubscribeImageRequest) (*model.SubscribeImageResponse, error) {
	requestDef := GenReqDefForSubscribeImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeImageResponse), nil
	}
}

// SubscribeImageInvoker 订阅镜像
func (c *EiHealthClient) SubscribeImageInvoker(request *model.SubscribeImageRequest) *SubscribeImageInvoker {
	requestDef := GenReqDefForSubscribeImage()
	return &SubscribeImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImage 更新镜像描述信息或者类型
//
// 更新镜像描述信息或者类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateImage(request *model.UpdateImageRequest) (*model.UpdateImageResponse, error) {
	requestDef := GenReqDefForUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageResponse), nil
	}
}

// UpdateImageInvoker 更新镜像描述信息或者类型
func (c *EiHealthClient) UpdateImageInvoker(request *model.UpdateImageRequest) *UpdateImageInvoker {
	requestDef := GenReqDefForUpdateImage()
	return &UpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobConfig 获取作业配置
//
// 获取作业配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJobConfig(request *model.ShowJobConfigRequest) (*model.ShowJobConfigResponse, error) {
	requestDef := GenReqDefForShowJobConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobConfigResponse), nil
	}
}

// ShowJobConfigInvoker 获取作业配置
func (c *EiHealthClient) ShowJobConfigInvoker(request *model.ShowJobConfigRequest) *ShowJobConfigInvoker {
	requestDef := GenReqDefForShowJobConfig()
	return &ShowJobConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJobConfig 设置作业配置
//
// 设置作业配置，目前支持修改保存时长(180天 - 10年)、记录数(1W-500W)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateJobConfig(request *model.UpdateJobConfigRequest) (*model.UpdateJobConfigResponse, error) {
	requestDef := GenReqDefForUpdateJobConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobConfigResponse), nil
	}
}

// UpdateJobConfigInvoker 设置作业配置
func (c *EiHealthClient) UpdateJobConfigInvoker(request *model.UpdateJobConfigRequest) *UpdateJobConfigInvoker {
	requestDef := GenReqDefForUpdateJobConfig()
	return &UpdateJobConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelJob 取消或强制停止作业调度
//
// 取消或强制作业调度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CancelJob(request *model.CancelJobRequest) (*model.CancelJobResponse, error) {
	requestDef := GenReqDefForCancelJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelJobResponse), nil
	}
}

// CancelJobInvoker 取消或强制停止作业调度
func (c *EiHealthClient) CancelJobInvoker(request *model.CancelJobRequest) *CancelJobInvoker {
	requestDef := GenReqDefForCancelJob()
	return &CancelJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除作业
//
// 删除作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除作业
func (c *EiHealthClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteJob 启动作业
//
// 启动作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ExecuteJob(request *model.ExecuteJobRequest) (*model.ExecuteJobResponse, error) {
	requestDef := GenReqDefForExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteJobResponse), nil
	}
}

// ExecuteJobInvoker 启动作业
func (c *EiHealthClient) ExecuteJobInvoker(request *model.ExecuteJobRequest) *ExecuteJobInvoker {
	requestDef := GenReqDefForExecuteJob()
	return &ExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJob 获取作业列表
//
// 获取作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListJob(request *model.ListJobRequest) (*model.ListJobResponse, error) {
	requestDef := GenReqDefForListJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobResponse), nil
	}
}

// ListJobInvoker 获取作业列表
func (c *EiHealthClient) ListJobInvoker(request *model.ListJobRequest) *ListJobInvoker {
	requestDef := GenReqDefForListJob()
	return &ListJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryJob 重试作业
//
// 重试作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RetryJob(request *model.RetryJobRequest) (*model.RetryJobResponse, error) {
	requestDef := GenReqDefForRetryJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryJobResponse), nil
	}
}

// RetryJobInvoker 重试作业
func (c *EiHealthClient) RetryJobInvoker(request *model.RetryJobRequest) *RetryJobInvoker {
	requestDef := GenReqDefForRetryJob()
	return &RetryJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 获取作业详情
//
// 获取作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 获取作业详情
func (c *EiHealthClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobEvent 获取作业事件
//
// 获取作业事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJobEvent(request *model.ShowJobEventRequest) (*model.ShowJobEventResponse, error) {
	requestDef := GenReqDefForShowJobEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobEventResponse), nil
	}
}

// ShowJobEventInvoker 获取作业事件
func (c *EiHealthClient) ShowJobEventInvoker(request *model.ShowJobEventRequest) *ShowJobEventInvoker {
	requestDef := GenReqDefForShowJobEvent()
	return &ShowJobEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobLog 获取作业日志
//
// 获取作业日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJobLog(request *model.ShowJobLogRequest) (*model.ShowJobLogResponse, error) {
	requestDef := GenReqDefForShowJobLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobLogResponse), nil
	}
}

// ShowJobLogInvoker 获取作业日志
func (c *EiHealthClient) ShowJobLogInvoker(request *model.ShowJobLogRequest) *ShowJobLogInvoker {
	requestDef := GenReqDefForShowJobLog()
	return &ShowJobLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskEvents 获取子任务启动事件
//
// 获取子任务启动事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskEvents(request *model.ShowTaskEventsRequest) (*model.ShowTaskEventsResponse, error) {
	requestDef := GenReqDefForShowTaskEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskEventsResponse), nil
	}
}

// ShowTaskEventsInvoker 获取子任务启动事件
func (c *EiHealthClient) ShowTaskEventsInvoker(request *model.ShowTaskEventsRequest) *ShowTaskEventsInvoker {
	requestDef := GenReqDefForShowTaskEvents()
	return &ShowTaskEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstanceEvents 获取子任务中实例的事件
//
// 获取子任务中实例的事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstanceEvents(request *model.ShowTaskInstanceEventsRequest) (*model.ShowTaskInstanceEventsResponse, error) {
	requestDef := GenReqDefForShowTaskInstanceEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstanceEventsResponse), nil
	}
}

// ShowTaskInstanceEventsInvoker 获取子任务中实例的事件
func (c *EiHealthClient) ShowTaskInstanceEventsInvoker(request *model.ShowTaskInstanceEventsRequest) *ShowTaskInstanceEventsInvoker {
	requestDef := GenReqDefForShowTaskInstanceEvents()
	return &ShowTaskInstanceEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstancePod 获取子任务中实例的pod信息
//
// 获取子任务中实例的pod信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstancePod(request *model.ShowTaskInstancePodRequest) (*model.ShowTaskInstancePodResponse, error) {
	requestDef := GenReqDefForShowTaskInstancePod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstancePodResponse), nil
	}
}

// ShowTaskInstancePodInvoker 获取子任务中实例的pod信息
func (c *EiHealthClient) ShowTaskInstancePodInvoker(request *model.ShowTaskInstancePodRequest) *ShowTaskInstancePodInvoker {
	requestDef := GenReqDefForShowTaskInstancePod()
	return &ShowTaskInstancePodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstances 获取子任务实例信息
//
// 获取子任务实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstances(request *model.ShowTaskInstancesRequest) (*model.ShowTaskInstancesResponse, error) {
	requestDef := GenReqDefForShowTaskInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstancesResponse), nil
	}
}

// ShowTaskInstancesInvoker 获取子任务实例信息
func (c *EiHealthClient) ShowTaskInstancesInvoker(request *model.ShowTaskInstancesRequest) *ShowTaskInstancesInvoker {
	requestDef := GenReqDefForShowTaskInstances()
	return &ShowTaskInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLabel 创建标签
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateLabel(request *model.CreateLabelRequest) (*model.CreateLabelResponse, error) {
	requestDef := GenReqDefForCreateLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLabelResponse), nil
	}
}

// CreateLabelInvoker 创建标签
func (c *EiHealthClient) CreateLabelInvoker(request *model.CreateLabelRequest) *CreateLabelInvoker {
	requestDef := GenReqDefForCreateLabel()
	return &CreateLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLabel 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteLabel(request *model.DeleteLabelRequest) (*model.DeleteLabelResponse, error) {
	requestDef := GenReqDefForDeleteLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLabelResponse), nil
	}
}

// DeleteLabelInvoker 删除标签
func (c *EiHealthClient) DeleteLabelInvoker(request *model.DeleteLabelRequest) *DeleteLabelInvoker {
	requestDef := GenReqDefForDeleteLabel()
	return &DeleteLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLabel 获取标签列表
//
// 获取标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListLabel(request *model.ListLabelRequest) (*model.ListLabelResponse, error) {
	requestDef := GenReqDefForListLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLabelResponse), nil
	}
}

// ListLabelInvoker 获取标签列表
func (c *EiHealthClient) ListLabelInvoker(request *model.ListLabelRequest) *ListLabelInvoker {
	requestDef := GenReqDefForListLabel()
	return &ListLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLabelPage 创建标签页面
//
// 创建标签页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateLabelPage(request *model.CreateLabelPageRequest) (*model.CreateLabelPageResponse, error) {
	requestDef := GenReqDefForCreateLabelPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLabelPageResponse), nil
	}
}

// CreateLabelPageInvoker 创建标签页面
func (c *EiHealthClient) CreateLabelPageInvoker(request *model.CreateLabelPageRequest) *CreateLabelPageInvoker {
	requestDef := GenReqDefForCreateLabelPage()
	return &CreateLabelPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLabelPage 删除标签页面
//
// 删除标签页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteLabelPage(request *model.DeleteLabelPageRequest) (*model.DeleteLabelPageResponse, error) {
	requestDef := GenReqDefForDeleteLabelPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLabelPageResponse), nil
	}
}

// DeleteLabelPageInvoker 删除标签页面
func (c *EiHealthClient) DeleteLabelPageInvoker(request *model.DeleteLabelPageRequest) *DeleteLabelPageInvoker {
	requestDef := GenReqDefForDeleteLabelPage()
	return &DeleteLabelPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLabelPage 获取标签页面列表
//
// 获取标签页面列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListLabelPage(request *model.ListLabelPageRequest) (*model.ListLabelPageResponse, error) {
	requestDef := GenReqDefForListLabelPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLabelPageResponse), nil
	}
}

// ListLabelPageInvoker 获取标签页面列表
func (c *EiHealthClient) ListLabelPageInvoker(request *model.ListLabelPageRequest) *ListLabelPageInvoker {
	requestDef := GenReqDefForListLabelPage()
	return &ListLabelPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckEmailConnection 邮箱连通性测试
//
// 邮箱连通性测试
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CheckEmailConnection(request *model.CheckEmailConnectionRequest) (*model.CheckEmailConnectionResponse, error) {
	requestDef := GenReqDefForCheckEmailConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckEmailConnectionResponse), nil
	}
}

// CheckEmailConnectionInvoker 邮箱连通性测试
func (c *EiHealthClient) CheckEmailConnectionInvoker(request *model.CheckEmailConnectionRequest) *CheckEmailConnectionInvoker {
	requestDef := GenReqDefForCheckEmailConnection()
	return &CheckEmailConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMessageEmailConfig 删除消息邮件配置
//
// 删除消息邮件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteMessageEmailConfig(request *model.DeleteMessageEmailConfigRequest) (*model.DeleteMessageEmailConfigResponse, error) {
	requestDef := GenReqDefForDeleteMessageEmailConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMessageEmailConfigResponse), nil
	}
}

// DeleteMessageEmailConfigInvoker 删除消息邮件配置
func (c *EiHealthClient) DeleteMessageEmailConfigInvoker(request *model.DeleteMessageEmailConfigRequest) *DeleteMessageEmailConfigInvoker {
	requestDef := GenReqDefForDeleteMessageEmailConfig()
	return &DeleteMessageEmailConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessage 获取消息列表
//
// 从消息中心获取当前用户有权限查看的消息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListMessage(request *model.ListMessageRequest) (*model.ListMessageResponse, error) {
	requestDef := GenReqDefForListMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageResponse), nil
	}
}

// ListMessageInvoker 获取消息列表
func (c *EiHealthClient) ListMessageInvoker(request *model.ListMessageRequest) *ListMessageInvoker {
	requestDef := GenReqDefForListMessage()
	return &ListMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMessageClearRule 获取消息清理规则
//
// 获取消息清理规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowMessageClearRule(request *model.ShowMessageClearRuleRequest) (*model.ShowMessageClearRuleResponse, error) {
	requestDef := GenReqDefForShowMessageClearRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessageClearRuleResponse), nil
	}
}

// ShowMessageClearRuleInvoker 获取消息清理规则
func (c *EiHealthClient) ShowMessageClearRuleInvoker(request *model.ShowMessageClearRuleRequest) *ShowMessageClearRuleInvoker {
	requestDef := GenReqDefForShowMessageClearRule()
	return &ShowMessageClearRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMessageEmailConfig 获取消息邮件配置
//
// 获取消息邮件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowMessageEmailConfig(request *model.ShowMessageEmailConfigRequest) (*model.ShowMessageEmailConfigResponse, error) {
	requestDef := GenReqDefForShowMessageEmailConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessageEmailConfigResponse), nil
	}
}

// ShowMessageEmailConfigInvoker 获取消息邮件配置
func (c *EiHealthClient) ShowMessageEmailConfigInvoker(request *model.ShowMessageEmailConfigRequest) *ShowMessageEmailConfigInvoker {
	requestDef := GenReqDefForShowMessageEmailConfig()
	return &ShowMessageEmailConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMessageReceiveConfig 获取用户邮件配置
//
// 获取用户邮件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowMessageReceiveConfig(request *model.ShowMessageReceiveConfigRequest) (*model.ShowMessageReceiveConfigResponse, error) {
	requestDef := GenReqDefForShowMessageReceiveConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessageReceiveConfigResponse), nil
	}
}

// ShowMessageReceiveConfigInvoker 获取用户邮件配置
func (c *EiHealthClient) ShowMessageReceiveConfigInvoker(request *model.ShowMessageReceiveConfigRequest) *ShowMessageReceiveConfigInvoker {
	requestDef := GenReqDefForShowMessageReceiveConfig()
	return &ShowMessageReceiveConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMessageClearRule 设置消息清理规则
//
// 设置消息清理规则，支持修改保存时长(180天 - 10年)、记录数(1W-500W)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateMessageClearRule(request *model.UpdateMessageClearRuleRequest) (*model.UpdateMessageClearRuleResponse, error) {
	requestDef := GenReqDefForUpdateMessageClearRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMessageClearRuleResponse), nil
	}
}

// UpdateMessageClearRuleInvoker 设置消息清理规则
func (c *EiHealthClient) UpdateMessageClearRuleInvoker(request *model.UpdateMessageClearRuleRequest) *UpdateMessageClearRuleInvoker {
	requestDef := GenReqDefForUpdateMessageClearRule()
	return &UpdateMessageClearRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMessageEmailConfig 设置消息邮件配置
//
// 设置消息邮件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateMessageEmailConfig(request *model.UpdateMessageEmailConfigRequest) (*model.UpdateMessageEmailConfigResponse, error) {
	requestDef := GenReqDefForUpdateMessageEmailConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMessageEmailConfigResponse), nil
	}
}

// UpdateMessageEmailConfigInvoker 设置消息邮件配置
func (c *EiHealthClient) UpdateMessageEmailConfigInvoker(request *model.UpdateMessageEmailConfigRequest) *UpdateMessageEmailConfigInvoker {
	requestDef := GenReqDefForUpdateMessageEmailConfig()
	return &UpdateMessageEmailConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMessageReceiveConfig 设置用户邮件配置
//
// 设置用户邮件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateMessageReceiveConfig(request *model.UpdateMessageReceiveConfigRequest) (*model.UpdateMessageReceiveConfigResponse, error) {
	requestDef := GenReqDefForUpdateMessageReceiveConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMessageReceiveConfigResponse), nil
	}
}

// UpdateMessageReceiveConfigInvoker 设置用户邮件配置
func (c *EiHealthClient) UpdateMessageReceiveConfigInvoker(request *model.UpdateMessageReceiveConfigRequest) *UpdateMessageReceiveConfigInvoker {
	requestDef := GenReqDefForUpdateMessageReceiveConfig()
	return &UpdateMessageReceiveConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateNodeLabel 设置节点标签
//
// 设置节点标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchUpdateNodeLabel(request *model.BatchUpdateNodeLabelRequest) (*model.BatchUpdateNodeLabelResponse, error) {
	requestDef := GenReqDefForBatchUpdateNodeLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateNodeLabelResponse), nil
	}
}

// BatchUpdateNodeLabelInvoker 设置节点标签
func (c *EiHealthClient) BatchUpdateNodeLabelInvoker(request *model.BatchUpdateNodeLabelRequest) *BatchUpdateNodeLabelInvoker {
	requestDef := GenReqDefForBatchUpdateNodeLabel()
	return &BatchUpdateNodeLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterAllNodeLabel 获取节点标签集
//
// 获取节点标签集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListClusterAllNodeLabel(request *model.ListClusterAllNodeLabelRequest) (*model.ListClusterAllNodeLabelResponse, error) {
	requestDef := GenReqDefForListClusterAllNodeLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterAllNodeLabelResponse), nil
	}
}

// ListClusterAllNodeLabelInvoker 获取节点标签集
func (c *EiHealthClient) ListClusterAllNodeLabelInvoker(request *model.ListClusterAllNodeLabelRequest) *ListClusterAllNodeLabelInvoker {
	requestDef := GenReqDefForListClusterAllNodeLabel()
	return &ListClusterAllNodeLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodeLabel 获取节点标签集
//
// 获取节点标签集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListNodeLabel(request *model.ListNodeLabelRequest) (*model.ListNodeLabelResponse, error) {
	requestDef := GenReqDefForListNodeLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodeLabelResponse), nil
	}
}

// ListNodeLabelInvoker 获取节点标签集
func (c *EiHealthClient) ListNodeLabelInvoker(request *model.ListNodeLabelRequest) *ListNodeLabelInvoker {
	requestDef := GenReqDefForListNodeLabel()
	return &ListNodeLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPresetLabel 获取预置标签列表
//
// 获取预置标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListPresetLabel(request *model.ListPresetLabelRequest) (*model.ListPresetLabelResponse, error) {
	requestDef := GenReqDefForListPresetLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPresetLabelResponse), nil
	}
}

// ListPresetLabelInvoker 获取预置标签列表
func (c *EiHealthClient) ListPresetLabelInvoker(request *model.ListPresetLabelRequest) *ListPresetLabelInvoker {
	requestDef := GenReqDefForListPresetLabel()
	return &ListPresetLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotebook 创建notebook
//
// 创建notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateNotebook(request *model.CreateNotebookRequest) (*model.CreateNotebookResponse, error) {
	requestDef := GenReqDefForCreateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotebookResponse), nil
	}
}

// CreateNotebookInvoker 创建notebook
func (c *EiHealthClient) CreateNotebookInvoker(request *model.CreateNotebookRequest) *CreateNotebookInvoker {
	requestDef := GenReqDefForCreateNotebook()
	return &CreateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotebook 删除notebook
//
// 删除notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteNotebook(request *model.DeleteNotebookRequest) (*model.DeleteNotebookResponse, error) {
	requestDef := GenReqDefForDeleteNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotebookResponse), nil
	}
}

// DeleteNotebookInvoker 删除notebook
func (c *EiHealthClient) DeleteNotebookInvoker(request *model.DeleteNotebookRequest) *DeleteNotebookInvoker {
	requestDef := GenReqDefForDeleteNotebook()
	return &DeleteNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotebook 获取notebook列表
//
// 获取notebook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListNotebook(request *model.ListNotebookRequest) (*model.ListNotebookResponse, error) {
	requestDef := GenReqDefForListNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotebookResponse), nil
	}
}

// ListNotebookInvoker 获取notebook列表
func (c *EiHealthClient) ListNotebookInvoker(request *model.ListNotebookRequest) *ListNotebookInvoker {
	requestDef := GenReqDefForListNotebook()
	return &ListNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotebookTool 获取notebook工作环境
//
// 获取notebook工作环境
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListNotebookTool(request *model.ListNotebookToolRequest) (*model.ListNotebookToolResponse, error) {
	requestDef := GenReqDefForListNotebookTool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotebookToolResponse), nil
	}
}

// ListNotebookToolInvoker 获取notebook工作环境
func (c *EiHealthClient) ListNotebookToolInvoker(request *model.ListNotebookToolRequest) *ListNotebookToolInvoker {
	requestDef := GenReqDefForListNotebookTool()
	return &ListNotebookToolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebook 获取notebook详情
//
// 获取notebook详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowNotebook(request *model.ShowNotebookRequest) (*model.ShowNotebookResponse, error) {
	requestDef := GenReqDefForShowNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookResponse), nil
	}
}

// ShowNotebookInvoker 获取notebook详情
func (c *EiHealthClient) ShowNotebookInvoker(request *model.ShowNotebookRequest) *ShowNotebookInvoker {
	requestDef := GenReqDefForShowNotebook()
	return &ShowNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebookToken 获取notebook鉴权信息
//
// 获取notebook鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowNotebookToken(request *model.ShowNotebookTokenRequest) (*model.ShowNotebookTokenResponse, error) {
	requestDef := GenReqDefForShowNotebookToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookTokenResponse), nil
	}
}

// ShowNotebookTokenInvoker 获取notebook鉴权信息
func (c *EiHealthClient) ShowNotebookTokenInvoker(request *model.ShowNotebookTokenRequest) *ShowNotebookTokenInvoker {
	requestDef := GenReqDefForShowNotebookToken()
	return &ShowNotebookTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopOrStartNotebook 启停notebook
//
// 启停notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StopOrStartNotebook(request *model.StopOrStartNotebookRequest) (*model.StopOrStartNotebookResponse, error) {
	requestDef := GenReqDefForStopOrStartNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopOrStartNotebookResponse), nil
	}
}

// StopOrStartNotebookInvoker 启停notebook
func (c *EiHealthClient) StopOrStartNotebookInvoker(request *model.StopOrStartNotebookRequest) *StopOrStartNotebookInvoker {
	requestDef := GenReqDefForStopOrStartNotebook()
	return &StopOrStartNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotebook 更新notebook
//
// 更新notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateNotebook(request *model.UpdateNotebookRequest) (*model.UpdateNotebookResponse, error) {
	requestDef := GenReqDefForUpdateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotebookResponse), nil
	}
}

// UpdateNotebookInvoker 更新notebook
func (c *EiHealthClient) UpdateNotebookInvoker(request *model.UpdateNotebookRequest) *UpdateNotebookInvoker {
	requestDef := GenReqDefForUpdateNotebook()
	return &UpdateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBucket 获取用户OBS桶列表
//
// 获取用户OBS桶列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListObsBucket(request *model.ListObsBucketRequest) (*model.ListObsBucketResponse, error) {
	requestDef := GenReqDefForListObsBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketResponse), nil
	}
}

// ListObsBucketInvoker 获取用户OBS桶列表
func (c *EiHealthClient) ListObsBucketInvoker(request *model.ListObsBucketRequest) *ListObsBucketInvoker {
	requestDef := GenReqDefForListObsBucket()
	return &ListObsBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBucketObject 获取用户OBS桶内对象
//
// 获取用户OBS桶内对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListObsBucketObject(request *model.ListObsBucketObjectRequest) (*model.ListObsBucketObjectResponse, error) {
	requestDef := GenReqDefForListObsBucketObject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketObjectResponse), nil
	}
}

// ListObsBucketObjectInvoker 获取用户OBS桶内对象
func (c *EiHealthClient) ListObsBucketObjectInvoker(request *model.ListObsBucketObjectRequest) *ListObsBucketObjectInvoker {
	requestDef := GenReqDefForListObsBucketObject()
	return &ListObsBucketObjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOverview 获取医疗平台信息
//
// 获取医疗平台信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowOverview(request *model.ShowOverviewRequest) (*model.ShowOverviewResponse, error) {
	requestDef := GenReqDefForShowOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOverviewResponse), nil
	}
}

// ShowOverviewInvoker 获取医疗平台信息
func (c *EiHealthClient) ShowOverviewInvoker(request *model.ShowOverviewRequest) *ShowOverviewInvoker {
	requestDef := GenReqDefForShowOverview()
	return &ShowOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePerformanceResource 购买性能加速资源
//
// 购买性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreatePerformanceResource(request *model.CreatePerformanceResourceRequest) (*model.CreatePerformanceResourceResponse, error) {
	requestDef := GenReqDefForCreatePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePerformanceResourceResponse), nil
	}
}

// CreatePerformanceResourceInvoker 购买性能加速资源
func (c *EiHealthClient) CreatePerformanceResourceInvoker(request *model.CreatePerformanceResourceRequest) *CreatePerformanceResourceInvoker {
	requestDef := GenReqDefForCreatePerformanceResource()
	return &CreatePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePerformanceResource 删除性能加速资源
//
// 删除性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeletePerformanceResource(request *model.DeletePerformanceResourceRequest) (*model.DeletePerformanceResourceResponse, error) {
	requestDef := GenReqDefForDeletePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePerformanceResourceResponse), nil
	}
}

// DeletePerformanceResourceInvoker 删除性能加速资源
func (c *EiHealthClient) DeletePerformanceResourceInvoker(request *model.DeletePerformanceResourceRequest) *DeletePerformanceResourceInvoker {
	requestDef := GenReqDefForDeletePerformanceResource()
	return &DeletePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPerformanceResources 查询性能加速资源
//
// 查询性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListPerformanceResources(request *model.ListPerformanceResourcesRequest) (*model.ListPerformanceResourcesResponse, error) {
	requestDef := GenReqDefForListPerformanceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPerformanceResourcesResponse), nil
	}
}

// ListPerformanceResourcesInvoker 查询性能加速资源
func (c *EiHealthClient) ListPerformanceResourcesInvoker(request *model.ListPerformanceResourcesRequest) *ListPerformanceResourcesInvoker {
	requestDef := GenReqDefForListPerformanceResources()
	return &ListPerformanceResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePerformanceResource 更新性能加速资源配置
//
// 更新性能加速资源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdatePerformanceResource(request *model.UpdatePerformanceResourceRequest) (*model.UpdatePerformanceResourceResponse, error) {
	requestDef := GenReqDefForUpdatePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePerformanceResourceResponse), nil
	}
}

// UpdatePerformanceResourceInvoker 更新性能加速资源配置
func (c *EiHealthClient) UpdatePerformanceResourceInvoker(request *model.UpdatePerformanceResourceRequest) *UpdatePerformanceResourceInvoker {
	requestDef := GenReqDefForUpdatePerformanceResource()
	return &UpdatePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMember 批量删除项目成员
//
// 批量删除项目成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchDeleteMember(request *model.BatchDeleteMemberRequest) (*model.BatchDeleteMemberResponse, error) {
	requestDef := GenReqDefForBatchDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMemberResponse), nil
	}
}

// BatchDeleteMemberInvoker 批量删除项目成员
func (c *EiHealthClient) BatchDeleteMemberInvoker(request *model.BatchDeleteMemberRequest) *BatchDeleteMemberInvoker {
	requestDef := GenReqDefForBatchDeleteMember()
	return &BatchDeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProject 创建项目
//
// 创建项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateProject(request *model.CreateProjectRequest) (*model.CreateProjectResponse, error) {
	requestDef := GenReqDefForCreateProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectResponse), nil
	}
}

// CreateProjectInvoker 创建项目
func (c *EiHealthClient) CreateProjectInvoker(request *model.CreateProjectRequest) *CreateProjectInvoker {
	requestDef := GenReqDefForCreateProject()
	return &CreateProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 移除项目成员
//
// 移除项目成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 移除项目成员
func (c *EiHealthClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProject 删除项目
//
// 删除项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteProject(request *model.DeleteProjectRequest) (*model.DeleteProjectResponse, error) {
	requestDef := GenReqDefForDeleteProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectResponse), nil
	}
}

// DeleteProjectInvoker 删除项目
func (c *EiHealthClient) DeleteProjectInvoker(request *model.DeleteProjectRequest) *DeleteProjectInvoker {
	requestDef := GenReqDefForDeleteProject()
	return &DeleteProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProject 获取项目列表
//
// 获取项目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListProject(request *model.ListProjectRequest) (*model.ListProjectResponse, error) {
	requestDef := GenReqDefForListProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectResponse), nil
	}
}

// ListProjectInvoker 获取项目列表
func (c *EiHealthClient) ListProjectInvoker(request *model.ListProjectRequest) *ListProjectInvoker {
	requestDef := GenReqDefForListProject()
	return &ListProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecentJob 获取最近的作业列表
//
// 获取最近的作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListRecentJob(request *model.ListRecentJobRequest) (*model.ListRecentJobResponse, error) {
	requestDef := GenReqDefForListRecentJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecentJobResponse), nil
	}
}

// ListRecentJobInvoker 获取最近的作业列表
func (c *EiHealthClient) ListRecentJobInvoker(request *model.ListRecentJobRequest) *ListRecentJobInvoker {
	requestDef := GenReqDefForListRecentJob()
	return &ListRecentJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProject 获取项目详情
//
// 获取项目详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowProject(request *model.ShowProjectRequest) (*model.ShowProjectResponse, error) {
	requestDef := GenReqDefForShowProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectResponse), nil
	}
}

// ShowProjectInvoker 获取项目详情
func (c *EiHealthClient) ShowProjectInvoker(request *model.ShowProjectRequest) *ShowProjectInvoker {
	requestDef := GenReqDefForShowProject()
	return &ShowProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferProject 转移项目
//
// 转移项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) TransferProject(request *model.TransferProjectRequest) (*model.TransferProjectResponse, error) {
	requestDef := GenReqDefForTransferProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferProjectResponse), nil
	}
}

// TransferProjectInvoker 转移项目
func (c *EiHealthClient) TransferProjectInvoker(request *model.TransferProjectRequest) *TransferProjectInvoker {
	requestDef := GenReqDefForTransferProject()
	return &TransferProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMember 更新或者添加项目成员角色
//
// 更新或者添加项目成员角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

// UpdateMemberInvoker 更新或者添加项目成员角色
func (c *EiHealthClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProject 更新项目
//
// 更新项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateProject(request *model.UpdateProjectRequest) (*model.UpdateProjectResponse, error) {
	requestDef := GenReqDefForUpdateProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectResponse), nil
	}
}

// UpdateProjectInvoker 更新项目
func (c *EiHealthClient) UpdateProjectInvoker(request *model.UpdateProjectRequest) *UpdateProjectInvoker {
	requestDef := GenReqDefForUpdateProject()
	return &UpdateProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadDataTrace 下载近一万条审计日志
//
// 下载近一万条审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DownloadDataTrace(request *model.DownloadDataTraceRequest) (*model.DownloadDataTraceResponse, error) {
	requestDef := GenReqDefForDownloadDataTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDataTraceResponse), nil
	}
}

// DownloadDataTraceInvoker 下载近一万条审计日志
func (c *EiHealthClient) DownloadDataTraceInvoker(request *model.DownloadDataTraceRequest) *DownloadDataTraceInvoker {
	requestDef := GenReqDefForDownloadDataTrace()
	return &DownloadDataTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectTrace 获取项目审计日志
//
// 获取项目审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowProjectTrace(request *model.ShowProjectTraceRequest) (*model.ShowProjectTraceResponse, error) {
	requestDef := GenReqDefForShowProjectTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectTraceResponse), nil
	}
}

// ShowProjectTraceInvoker 获取项目审计日志
func (c *EiHealthClient) ShowProjectTraceInvoker(request *model.ShowProjectTraceRequest) *ShowProjectTraceInvoker {
	requestDef := GenReqDefForShowProjectTrace()
	return &ShowProjectTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectTraceData 获取指定审计日志
//
// 获取指定审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowProjectTraceData(request *model.ShowProjectTraceDataRequest) (*model.ShowProjectTraceDataResponse, error) {
	requestDef := GenReqDefForShowProjectTraceData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectTraceDataResponse), nil
	}
}

// ShowProjectTraceDataInvoker 获取指定审计日志
func (c *EiHealthClient) ShowProjectTraceDataInvoker(request *model.ShowProjectTraceDataRequest) *ShowProjectTraceDataInvoker {
	requestDef := GenReqDefForShowProjectTraceData()
	return &ShowProjectTraceDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectTracker 获取项目审计日志追踪器
//
// 获取项目审计日志追踪器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowProjectTracker(request *model.ShowProjectTrackerRequest) (*model.ShowProjectTrackerResponse, error) {
	requestDef := GenReqDefForShowProjectTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectTrackerResponse), nil
	}
}

// ShowProjectTrackerInvoker 获取项目审计日志追踪器
func (c *EiHealthClient) ShowProjectTrackerInvoker(request *model.ShowProjectTrackerRequest) *ShowProjectTrackerInvoker {
	requestDef := GenReqDefForShowProjectTracker()
	return &ShowProjectTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectTracker 更新项目审计日志追踪器配置
//
// 更新项目审计日志追踪器配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateProjectTracker(request *model.UpdateProjectTrackerRequest) (*model.UpdateProjectTrackerResponse, error) {
	requestDef := GenReqDefForUpdateProjectTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectTrackerResponse), nil
	}
}

// UpdateProjectTrackerInvoker 更新项目审计日志追踪器配置
func (c *EiHealthClient) UpdateProjectTrackerInvoker(request *model.UpdateProjectTrackerRequest) *UpdateProjectTrackerInvoker {
	requestDef := GenReqDefForUpdateProjectTracker()
	return &UpdateProjectTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceMetricData 获取资源监控数据
//
// 获取资源监控数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowResourceMetricData(request *model.ShowResourceMetricDataRequest) (*model.ShowResourceMetricDataResponse, error) {
	requestDef := GenReqDefForShowResourceMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceMetricDataResponse), nil
	}
}

// ShowResourceMetricDataInvoker 获取资源监控数据
func (c *EiHealthClient) ShowResourceMetricDataInvoker(request *model.ShowResourceMetricDataRequest) *ShowResourceMetricDataInvoker {
	requestDef := GenReqDefForShowResourceMetricData()
	return &ShowResourceMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStar 取消收藏
//
// 取消收藏
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteStar(request *model.DeleteStarRequest) (*model.DeleteStarResponse, error) {
	requestDef := GenReqDefForDeleteStar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStarResponse), nil
	}
}

// DeleteStarInvoker 取消收藏
func (c *EiHealthClient) DeleteStarInvoker(request *model.DeleteStarRequest) *DeleteStarInvoker {
	requestDef := GenReqDefForDeleteStar()
	return &DeleteStarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStar 获取收藏资产列表
//
// 获取收藏资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListStar(request *model.ListStarRequest) (*model.ListStarResponse, error) {
	requestDef := GenReqDefForListStar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarResponse), nil
	}
}

// ListStarInvoker 获取收藏资产列表
func (c *EiHealthClient) ListStarInvoker(request *model.ListStarRequest) *ListStarInvoker {
	requestDef := GenReqDefForListStar()
	return &ListStarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStar 收藏资产
//
// 收藏资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateStar(request *model.UpdateStarRequest) (*model.UpdateStarResponse, error) {
	requestDef := GenReqDefForUpdateStar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStarResponse), nil
	}
}

// UpdateStarInvoker 收藏资产
func (c *EiHealthClient) UpdateStarInvoker(request *model.UpdateStarRequest) *UpdateStarInvoker {
	requestDef := GenReqDefForUpdateStar()
	return &UpdateStarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPerformanceResourceStat 获取性能加速资源上统计信息
//
// 获取性能加速资源上统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListPerformanceResourceStat(request *model.ListPerformanceResourceStatRequest) (*model.ListPerformanceResourceStatResponse, error) {
	requestDef := GenReqDefForListPerformanceResourceStat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPerformanceResourceStatResponse), nil
	}
}

// ListPerformanceResourceStatInvoker 获取性能加速资源上统计信息
func (c *EiHealthClient) ListPerformanceResourceStatInvoker(request *model.ListPerformanceResourceStatRequest) *ListPerformanceResourceStatInvoker {
	requestDef := GenReqDefForListPerformanceResourceStat()
	return &ListPerformanceResourceStatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowStatistic 统计应用、流程、作业数目
//
// 统计应用、流程、作业数目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListWorkflowStatistic(request *model.ListWorkflowStatisticRequest) (*model.ListWorkflowStatisticResponse, error) {
	requestDef := GenReqDefForListWorkflowStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowStatisticResponse), nil
	}
}

// ListWorkflowStatisticInvoker 统计应用、流程、作业数目
func (c *EiHealthClient) ListWorkflowStatisticInvoker(request *model.ListWorkflowStatisticRequest) *ListWorkflowStatisticInvoker {
	requestDef := GenReqDefForListWorkflowStatistic()
	return &ListWorkflowStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageResources 查询存储资源
//
// 查询存储资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListStorageResources(request *model.ListStorageResourcesRequest) (*model.ListStorageResourcesResponse, error) {
	requestDef := GenReqDefForListStorageResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageResourcesResponse), nil
	}
}

// ListStorageResourcesInvoker 查询存储资源
func (c *EiHealthClient) ListStorageResourcesInvoker(request *model.ListStorageResourcesRequest) *ListStorageResourcesInvoker {
	requestDef := GenReqDefForListStorageResources()
	return &ListStorageResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStudy 创建study
//
// 创建study
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateStudy(request *model.CreateStudyRequest) (*model.CreateStudyResponse, error) {
	requestDef := GenReqDefForCreateStudy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStudyResponse), nil
	}
}

// CreateStudyInvoker 创建study
func (c *EiHealthClient) CreateStudyInvoker(request *model.CreateStudyRequest) *CreateStudyInvoker {
	requestDef := GenReqDefForCreateStudy()
	return &CreateStudyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStudyJob 创建study作业
//
// 创建study作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateStudyJob(request *model.CreateStudyJobRequest) (*model.CreateStudyJobResponse, error) {
	requestDef := GenReqDefForCreateStudyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStudyJobResponse), nil
	}
}

// CreateStudyJobInvoker 创建study作业
func (c *EiHealthClient) CreateStudyJobInvoker(request *model.CreateStudyJobRequest) *CreateStudyJobInvoker {
	requestDef := GenReqDefForCreateStudyJob()
	return &CreateStudyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStudy 删除study
//
// 删除study
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteStudy(request *model.DeleteStudyRequest) (*model.DeleteStudyResponse, error) {
	requestDef := GenReqDefForDeleteStudy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStudyResponse), nil
	}
}

// DeleteStudyInvoker 删除study
func (c *EiHealthClient) DeleteStudyInvoker(request *model.DeleteStudyRequest) *DeleteStudyInvoker {
	requestDef := GenReqDefForDeleteStudy()
	return &DeleteStudyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStudy 列举study
//
// 列举study
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListStudy(request *model.ListStudyRequest) (*model.ListStudyResponse, error) {
	requestDef := GenReqDefForListStudy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStudyResponse), nil
	}
}

// ListStudyInvoker 列举study
func (c *EiHealthClient) ListStudyInvoker(request *model.ListStudyRequest) *ListStudyInvoker {
	requestDef := GenReqDefForListStudy()
	return &ListStudyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStudyJob 列举study所有作业
//
// 列举study所有作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListStudyJob(request *model.ListStudyJobRequest) (*model.ListStudyJobResponse, error) {
	requestDef := GenReqDefForListStudyJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStudyJobResponse), nil
	}
}

// ListStudyJobInvoker 列举study所有作业
func (c *EiHealthClient) ListStudyJobInvoker(request *model.ListStudyJobRequest) *ListStudyJobInvoker {
	requestDef := GenReqDefForListStudyJob()
	return &ListStudyJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Show3dStructureContent 获取生成study作业3D结构的内容
//
// 获取生成study作业3D结构的内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) Show3dStructureContent(request *model.Show3dStructureContentRequest) (*model.Show3dStructureContentResponse, error) {
	requestDef := GenReqDefForShow3dStructureContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Show3dStructureContentResponse), nil
	}
}

// Show3dStructureContentInvoker 获取生成study作业3D结构的内容
func (c *EiHealthClient) Show3dStructureContentInvoker(request *model.Show3dStructureContentRequest) *Show3dStructureContentInvoker {
	requestDef := GenReqDefForShow3dStructureContent()
	return &Show3dStructureContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtremumInfo 获取study作业的最值信息
//
// 获取study作业的最值信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowExtremumInfo(request *model.ShowExtremumInfoRequest) (*model.ShowExtremumInfoResponse, error) {
	requestDef := GenReqDefForShowExtremumInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtremumInfoResponse), nil
	}
}

// ShowExtremumInfoInvoker 获取study作业的最值信息
func (c *EiHealthClient) ShowExtremumInfoInvoker(request *model.ShowExtremumInfoRequest) *ShowExtremumInfoInvoker {
	requestDef := GenReqDefForShowExtremumInfo()
	return &ShowExtremumInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArchiveConfigs 获取跨域归档配置
//
// 获取跨域归档配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListArchiveConfigs(request *model.ListArchiveConfigsRequest) (*model.ListArchiveConfigsResponse, error) {
	requestDef := GenReqDefForListArchiveConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArchiveConfigsResponse), nil
	}
}

// ListArchiveConfigsInvoker 获取跨域归档配置
func (c *EiHealthClient) ListArchiveConfigsInvoker(request *model.ListArchiveConfigsRequest) *ListArchiveConfigsInvoker {
	requestDef := GenReqDefForListArchiveConfigs()
	return &ListArchiveConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnv 查询系统配置列表
//
// 获取系统配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowEnv(request *model.ShowEnvRequest) (*model.ShowEnvResponse, error) {
	requestDef := GenReqDefForShowEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvResponse), nil
	}
}

// ShowEnvInvoker 查询系统配置列表
func (c *EiHealthClient) ShowEnvInvoker(request *model.ShowEnvRequest) *ShowEnvInvoker {
	requestDef := GenReqDefForShowEnv()
	return &ShowEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVendor 获取供应商配置
//
// 获取供应商配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowVendor(request *model.ShowVendorRequest) (*model.ShowVendorResponse, error) {
	requestDef := GenReqDefForShowVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVendorResponse), nil
	}
}

// ShowVendorInvoker 获取供应商配置
func (c *EiHealthClient) ShowVendorInvoker(request *model.ShowVendorRequest) *ShowVendorInvoker {
	requestDef := GenReqDefForShowVendor()
	return &ShowVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateArchiveConfig 修改跨域归档设置
//
// 修改跨域归档设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateArchiveConfig(request *model.UpdateArchiveConfigRequest) (*model.UpdateArchiveConfigResponse, error) {
	requestDef := GenReqDefForUpdateArchiveConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateArchiveConfigResponse), nil
	}
}

// UpdateArchiveConfigInvoker 修改跨域归档设置
func (c *EiHealthClient) UpdateArchiveConfigInvoker(request *model.UpdateArchiveConfigRequest) *UpdateArchiveConfigInvoker {
	requestDef := GenReqDefForUpdateArchiveConfig()
	return &UpdateArchiveConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVendor 设置供应商配置
//
// 设置供应商配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateVendor(request *model.UpdateVendorRequest) (*model.UpdateVendorResponse, error) {
	requestDef := GenReqDefForUpdateVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVendorResponse), nil
	}
}

// UpdateVendorInvoker 设置供应商配置
func (c *EiHealthClient) UpdateVendorInvoker(request *model.UpdateVendorRequest) *UpdateVendorInvoker {
	requestDef := GenReqDefForUpdateVendor()
	return &UpdateVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuota 获取当前系统配额及资源使用情况
//
// 获取当前系统配额及资源使用情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListQuota(request *model.ListQuotaRequest) (*model.ListQuotaResponse, error) {
	requestDef := GenReqDefForListQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaResponse), nil
	}
}

// ListQuotaInvoker 获取当前系统配额及资源使用情况
func (c *EiHealthClient) ListQuotaInvoker(request *model.ListQuotaRequest) *ListQuotaInvoker {
	requestDef := GenReqDefForListQuota()
	return &ListQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建模板
//
// 创建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建模板
func (c *EiHealthClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除模板
//
// 删除模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除模板
func (c *EiHealthClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportTemplate 从其他项目导入模板
//
// 从其他项目导入模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportTemplate(request *model.ImportTemplateRequest) (*model.ImportTemplateResponse, error) {
	requestDef := GenReqDefForImportTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportTemplateResponse), nil
	}
}

// ImportTemplateInvoker 从其他项目导入模板
func (c *EiHealthClient) ImportTemplateInvoker(request *model.ImportTemplateRequest) *ImportTemplateInvoker {
	requestDef := GenReqDefForImportTemplate()
	return &ImportTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplate 查询模板列表
//
// 查询模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListTemplate(request *model.ListTemplateRequest) (*model.ListTemplateResponse, error) {
	requestDef := GenReqDefForListTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateResponse), nil
	}
}

// ListTemplateInvoker 查询模板列表
func (c *EiHealthClient) ListTemplateInvoker(request *model.ListTemplateRequest) *ListTemplateInvoker {
	requestDef := GenReqDefForListTemplate()
	return &ListTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplate 查询模板详情
//
// 查询模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTemplate(request *model.ShowTemplateRequest) (*model.ShowTemplateResponse, error) {
	requestDef := GenReqDefForShowTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateResponse), nil
	}
}

// ShowTemplateInvoker 查询模板详情
func (c *EiHealthClient) ShowTemplateInvoker(request *model.ShowTemplateRequest) *ShowTemplateInvoker {
	requestDef := GenReqDefForShowTemplate()
	return &ShowTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadTemplate 上传模板
//
// 上传模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UploadTemplate(request *model.UploadTemplateRequest) (*model.UploadTemplateResponse, error) {
	requestDef := GenReqDefForUploadTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadTemplateResponse), nil
	}
}

// UploadTemplateInvoker 上传模板
func (c *EiHealthClient) UploadTemplateInvoker(request *model.UploadTemplateRequest) *UploadTemplateInvoker {
	requestDef := GenReqDefForUploadTemplate()
	return &UploadTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCode 发送验证码
//
// 发送验证码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateCode(request *model.CreateCodeRequest) (*model.CreateCodeResponse, error) {
	requestDef := GenReqDefForCreateCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCodeResponse), nil
	}
}

// CreateCodeInvoker 发送验证码
func (c *EiHealthClient) CreateCodeInvoker(request *model.CreateCodeRequest) *CreateCodeInvoker {
	requestDef := GenReqDefForCreateCode()
	return &CreateCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 创建用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *EiHealthClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 删除用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *EiHealthClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMfa 获取可用的认证方法
//
// 获取可用的认证方法
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListMfa(request *model.ListMfaRequest) (*model.ListMfaResponse, error) {
	requestDef := GenReqDefForListMfa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMfaResponse), nil
	}
}

// ListMfaInvoker 获取可用的认证方法
func (c *EiHealthClient) ListMfaInvoker(request *model.ListMfaRequest) *ListMfaInvoker {
	requestDef := GenReqDefForListMfa()
	return &ListMfaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUser 获取用户列表
//
// 获取用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

// ListUserInvoker 获取用户列表
func (c *EiHealthClient) ListUserInvoker(request *model.ListUserRequest) *ListUserInvoker {
	requestDef := GenReqDefForListUser()
	return &ListUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTokenVerification 校验token
//
// 校验token是否可访问当前环境
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTokenVerification(request *model.ShowTokenVerificationRequest) (*model.ShowTokenVerificationResponse, error) {
	requestDef := GenReqDefForShowTokenVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTokenVerificationResponse), nil
	}
}

// ShowTokenVerificationInvoker 校验token
func (c *EiHealthClient) ShowTokenVerificationInvoker(request *model.ShowTokenVerificationRequest) *ShowTokenVerificationInvoker {
	requestDef := GenReqDefForShowTokenVerification()
	return &ShowTokenVerificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUser 获取指定用户详情
//
// 获取指定用户详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserResponse), nil
	}
}

// ShowUserInvoker 获取指定用户详情
func (c *EiHealthClient) ShowUserInvoker(request *model.ShowUserRequest) *ShowUserInvoker {
	requestDef := GenReqDefForShowUser()
	return &ShowUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserSetting 查询用户设置
//
// 查询用户设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowUserSetting(request *model.ShowUserSettingRequest) (*model.ShowUserSettingResponse, error) {
	requestDef := GenReqDefForShowUserSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserSettingResponse), nil
	}
}

// ShowUserSettingInvoker 查询用户设置
func (c *EiHealthClient) ShowUserSettingInvoker(request *model.ShowUserSettingRequest) *ShowUserSettingInvoker {
	requestDef := GenReqDefForShowUserSetting()
	return &ShowUserSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInitPassword 新用户重置密码
//
// 新用户重置密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateInitPassword(request *model.UpdateInitPasswordRequest) (*model.UpdateInitPasswordResponse, error) {
	requestDef := GenReqDefForUpdateInitPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInitPasswordResponse), nil
	}
}

// UpdateInitPasswordInvoker 新用户重置密码
func (c *EiHealthClient) UpdateInitPasswordInvoker(request *model.UpdateInitPasswordRequest) *UpdateInitPasswordInvoker {
	requestDef := GenReqDefForUpdateInitPassword()
	return &UpdateInitPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePassword 修改密码
//
// 修改密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdatePassword(request *model.UpdatePasswordRequest) (*model.UpdatePasswordResponse, error) {
	requestDef := GenReqDefForUpdatePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordResponse), nil
	}
}

// UpdatePasswordInvoker 修改密码
func (c *EiHealthClient) UpdatePasswordInvoker(request *model.UpdatePasswordRequest) *UpdatePasswordInvoker {
	requestDef := GenReqDefForUpdatePassword()
	return &UpdatePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户基本信息
//
// 修改用户基本信息（邮箱，手机）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改用户基本信息
func (c *EiHealthClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserByDomain 最终租户修改子用户
//
// 最终租户修改子用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserByDomain(request *model.UpdateUserByDomainRequest) (*model.UpdateUserByDomainResponse, error) {
	requestDef := GenReqDefForUpdateUserByDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserByDomainResponse), nil
	}
}

// UpdateUserByDomainInvoker 最终租户修改子用户
func (c *EiHealthClient) UpdateUserByDomainInvoker(request *model.UpdateUserByDomainRequest) *UpdateUserByDomainInvoker {
	requestDef := GenReqDefForUpdateUserByDomain()
	return &UpdateUserByDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserRole 更新用户角色
//
// 更新用户角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserRole(request *model.UpdateUserRoleRequest) (*model.UpdateUserRoleResponse, error) {
	requestDef := GenReqDefForUpdateUserRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserRoleResponse), nil
	}
}

// UpdateUserRoleInvoker 更新用户角色
func (c *EiHealthClient) UpdateUserRoleInvoker(request *model.UpdateUserRoleRequest) *UpdateUserRoleInvoker {
	requestDef := GenReqDefForUpdateUserRole()
	return &UpdateUserRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserSetting 更新用户设置
//
// 更新用户设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserSetting(request *model.UpdateUserSettingRequest) (*model.UpdateUserSettingResponse, error) {
	requestDef := GenReqDefForUpdateUserSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserSettingResponse), nil
	}
}

// UpdateUserSettingInvoker 更新用户设置
func (c *EiHealthClient) UpdateUserSettingInvoker(request *model.UpdateUserSettingRequest) *UpdateUserSettingInvoker {
	requestDef := GenReqDefForUpdateUserSetting()
	return &UpdateUserSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateCode 预验证
//
// 预验证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ValidateCode(request *model.ValidateCodeRequest) (*model.ValidateCodeResponse, error) {
	requestDef := GenReqDefForValidateCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateCodeResponse), nil
	}
}

// ValidateCodeInvoker 预验证
func (c *EiHealthClient) ValidateCodeInvoker(request *model.ValidateCodeRequest) *ValidateCodeInvoker {
	requestDef := GenReqDefForValidateCode()
	return &ValidateCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVendor 获取供应商列表
//
// 获取供应商列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListVendor(request *model.ListVendorRequest) (*model.ListVendorResponse, error) {
	requestDef := GenReqDefForListVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVendorResponse), nil
	}
}

// ListVendorInvoker 获取供应商列表
func (c *EiHealthClient) ListVendorInvoker(request *model.ListVendorRequest) *ListVendorInvoker {
	requestDef := GenReqDefForListVendor()
	return &ListVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建流程
//
// 创建流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建流程
func (c *EiHealthClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflow 删除流程
//
// 删除流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteWorkflow(request *model.DeleteWorkflowRequest) (*model.DeleteWorkflowResponse, error) {
	requestDef := GenReqDefForDeleteWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowResponse), nil
	}
}

// DeleteWorkflowInvoker 删除流程
func (c *EiHealthClient) DeleteWorkflowInvoker(request *model.DeleteWorkflowRequest) *DeleteWorkflowInvoker {
	requestDef := GenReqDefForDeleteWorkflow()
	return &DeleteWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportWorkflow 导入流程
//
// 导入流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportWorkflow(request *model.ImportWorkflowRequest) (*model.ImportWorkflowResponse, error) {
	requestDef := GenReqDefForImportWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportWorkflowResponse), nil
	}
}

// ImportWorkflowInvoker 导入流程
func (c *EiHealthClient) ImportWorkflowInvoker(request *model.ImportWorkflowRequest) *ImportWorkflowInvoker {
	requestDef := GenReqDefForImportWorkflow()
	return &ImportWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflow 获取流程列表
//
// 获取流程列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListWorkflow(request *model.ListWorkflowRequest) (*model.ListWorkflowResponse, error) {
	requestDef := GenReqDefForListWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowResponse), nil
	}
}

// ListWorkflowInvoker 获取流程列表
func (c *EiHealthClient) ListWorkflowInvoker(request *model.ListWorkflowRequest) *ListWorkflowInvoker {
	requestDef := GenReqDefForListWorkflow()
	return &ListWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflow 获取流程详情
//
// 获取流程详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowWorkflow(request *model.ShowWorkflowRequest) (*model.ShowWorkflowResponse, error) {
	requestDef := GenReqDefForShowWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowResponse), nil
	}
}

// ShowWorkflowInvoker 获取流程详情
func (c *EiHealthClient) ShowWorkflowInvoker(request *model.ShowWorkflowRequest) *ShowWorkflowInvoker {
	requestDef := GenReqDefForShowWorkflow()
	return &ShowWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeWorkflow 订阅流程
//
// 订阅流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeWorkflow(request *model.SubscribeWorkflowRequest) (*model.SubscribeWorkflowResponse, error) {
	requestDef := GenReqDefForSubscribeWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeWorkflowResponse), nil
	}
}

// SubscribeWorkflowInvoker 订阅流程
func (c *EiHealthClient) SubscribeWorkflowInvoker(request *model.SubscribeWorkflowRequest) *SubscribeWorkflowInvoker {
	requestDef := GenReqDefForSubscribeWorkflow()
	return &SubscribeWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflow 更新流程
//
// 更新流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateWorkflow(request *model.UpdateWorkflowRequest) (*model.UpdateWorkflowResponse, error) {
	requestDef := GenReqDefForUpdateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowResponse), nil
	}
}

// UpdateWorkflowInvoker 更新流程
func (c *EiHealthClient) UpdateWorkflowInvoker(request *model.UpdateWorkflowRequest) *UpdateWorkflowInvoker {
	requestDef := GenReqDefForUpdateWorkflow()
	return &UpdateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
