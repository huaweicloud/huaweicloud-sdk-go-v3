package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotedge/v2/model"
)

type IoTEdgeClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIoTEdgeClient(hcClient *httpclient.HcHttpClient) *IoTEdgeClient {
	return &IoTEdgeClient{HcClient: hcClient}
}

func IoTEdgeClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateEdgeNode 创建边缘节点
//
// 创建边缘节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateEdgeNode(request *model.CreateEdgeNodeRequest) (*model.CreateEdgeNodeResponse, error) {
	requestDef := GenReqDefForCreateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeResponse), nil
	}
}

// CreateEdgeNodeInvoker 创建边缘节点
func (c *IoTEdgeClient) CreateEdgeNodeInvoker(request *model.CreateEdgeNodeRequest) *CreateEdgeNodeInvoker {
	requestDef := GenReqDefForCreateEdgeNode()
	return &CreateEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstallCmd 生成边缘节点安装命令
//
// 生成边缘节点安装命令，命令有效时间30分钟，超过后需要重新生成
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateInstallCmd(request *model.CreateInstallCmdRequest) (*model.CreateInstallCmdResponse, error) {
	requestDef := GenReqDefForCreateInstallCmd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstallCmdResponse), nil
	}
}

// CreateInstallCmdInvoker 生成边缘节点安装命令
func (c *IoTEdgeClient) CreateInstallCmdInvoker(request *model.CreateInstallCmdRequest) *CreateInstallCmdInvoker {
	requestDef := GenReqDefForCreateInstallCmd()
	return &CreateInstallCmdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeNode 删除边缘节点
//
// 删除指定边缘节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteEdgeNode(request *model.DeleteEdgeNodeRequest) (*model.DeleteEdgeNodeResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeResponse), nil
	}
}

// DeleteEdgeNodeInvoker 删除边缘节点
func (c *IoTEdgeClient) DeleteEdgeNodeInvoker(request *model.DeleteEdgeNodeRequest) *DeleteEdgeNodeInvoker {
	requestDef := GenReqDefForDeleteEdgeNode()
	return &DeleteEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeNodes 查询边缘节点列表
//
// 查询边缘节点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListEdgeNodes(request *model.ListEdgeNodesRequest) (*model.ListEdgeNodesResponse, error) {
	requestDef := GenReqDefForListEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodesResponse), nil
	}
}

// ListEdgeNodesInvoker 查询边缘节点列表
func (c *IoTEdgeClient) ListEdgeNodesInvoker(request *model.ListEdgeNodesRequest) *ListEdgeNodesInvoker {
	requestDef := GenReqDefForListEdgeNodes()
	return &ListEdgeNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeNode 查询边缘节点详情
//
// 查询边缘节点详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowEdgeNode(request *model.ShowEdgeNodeRequest) (*model.ShowEdgeNodeResponse, error) {
	requestDef := GenReqDefForShowEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeResponse), nil
	}
}

// ShowEdgeNodeInvoker 查询边缘节点详情
func (c *IoTEdgeClient) ShowEdgeNodeInvoker(request *model.ShowEdgeNodeRequest) *ShowEdgeNodeInvoker {
	requestDef := GenReqDefForShowEdgeNode()
	return &ShowEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeNodeHostsInfo 查询边缘节点下的主机详情
//
// 查询边缘节点下的主机详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowEdgeNodeHostsInfo(request *model.ShowEdgeNodeHostsInfoRequest) (*model.ShowEdgeNodeHostsInfoResponse, error) {
	requestDef := GenReqDefForShowEdgeNodeHostsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeHostsInfoResponse), nil
	}
}

// ShowEdgeNodeHostsInfoInvoker 查询边缘节点下的主机详情
func (c *IoTEdgeClient) ShowEdgeNodeHostsInfoInvoker(request *model.ShowEdgeNodeHostsInfoRequest) *ShowEdgeNodeHostsInfoInvoker {
	requestDef := GenReqDefForShowEdgeNodeHostsInfo()
	return &ShowEdgeNodeHostsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeNode 修改边缘节点
//
// 修改边缘节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateEdgeNode(request *model.UpdateEdgeNodeRequest) (*model.UpdateEdgeNodeResponse, error) {
	requestDef := GenReqDefForUpdateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeNodeResponse), nil
	}
}

// UpdateEdgeNodeInvoker 修改边缘节点
func (c *IoTEdgeClient) UpdateEdgeNodeInvoker(request *model.UpdateEdgeNodeRequest) *UpdateEdgeNodeInvoker {
	requestDef := GenReqDefForUpdateEdgeNode()
	return &UpdateEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDeviceControlsRelease 设备控制释放
//
// 设备控制释放
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ExecuteDeviceControlsRelease(request *model.ExecuteDeviceControlsReleaseRequest) (*model.ExecuteDeviceControlsReleaseResponse, error) {
	requestDef := GenReqDefForExecuteDeviceControlsRelease()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeviceControlsReleaseResponse), nil
	}
}

// ExecuteDeviceControlsReleaseInvoker 设备控制释放
func (c *IoTEdgeClient) ExecuteDeviceControlsReleaseInvoker(request *model.ExecuteDeviceControlsReleaseRequest) *ExecuteDeviceControlsReleaseInvoker {
	requestDef := GenReqDefForExecuteDeviceControlsRelease()
	return &ExecuteDeviceControlsReleaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDeviceControlsSet 设备控制设置
//
// 设备控制设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ExecuteDeviceControlsSet(request *model.ExecuteDeviceControlsSetRequest) (*model.ExecuteDeviceControlsSetResponse, error) {
	requestDef := GenReqDefForExecuteDeviceControlsSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeviceControlsSetResponse), nil
	}
}

// ExecuteDeviceControlsSetInvoker 设备控制设置
func (c *IoTEdgeClient) ExecuteDeviceControlsSetInvoker(request *model.ExecuteDeviceControlsSetRequest) *ExecuteDeviceControlsSetInvoker {
	requestDef := GenReqDefForExecuteDeviceControlsSet()
	return &ExecuteDeviceControlsSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetDeviceControlDefaultValues 设备控制默认值
//
// 设备控制默认值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) SetDeviceControlDefaultValues(request *model.SetDeviceControlDefaultValuesRequest) (*model.SetDeviceControlDefaultValuesResponse, error) {
	requestDef := GenReqDefForSetDeviceControlDefaultValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDeviceControlDefaultValuesResponse), nil
	}
}

// SetDeviceControlDefaultValuesInvoker 设备控制默认值
func (c *IoTEdgeClient) SetDeviceControlDefaultValuesInvoker(request *model.SetDeviceControlDefaultValuesRequest) *SetDeviceControlDefaultValuesInvoker {
	requestDef := GenReqDefForSetDeviceControlDefaultValues()
	return &SetDeviceControlDefaultValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDevice 添加设备
//
// 添加设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) AddDevice(request *model.AddDeviceRequest) (*model.AddDeviceResponse, error) {
	requestDef := GenReqDefForAddDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceResponse), nil
	}
}

// AddDeviceInvoker 添加设备
func (c *IoTEdgeClient) AddDeviceInvoker(request *model.AddDeviceRequest) *AddDeviceInvoker {
	requestDef := GenReqDefForAddDevice()
	return &AddDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDevice 删除设备
//
// 删除设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

// DeleteDeviceInvoker 删除设备
func (c *IoTEdgeClient) DeleteDeviceInvoker(request *model.DeleteDeviceRequest) *DeleteDeviceInvoker {
	requestDef := GenReqDefForDeleteDevice()
	return &DeleteDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevices 查询设备列表
//
// 查询设备列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

// ListDevicesInvoker 查询设备列表
func (c *IoTEdgeClient) ListDevicesInvoker(request *model.ListDevicesRequest) *ListDevicesInvoker {
	requestDef := GenReqDefForListDevices()
	return &ListDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProductConfig 获取协议配置
//
// 获取协议配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowProductConfig(request *model.ShowProductConfigRequest) (*model.ShowProductConfigResponse, error) {
	requestDef := GenReqDefForShowProductConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductConfigResponse), nil
	}
}

// ShowProductConfigInvoker 获取协议配置
func (c *IoTEdgeClient) ShowProductConfigInvoker(request *model.ShowProductConfigRequest) *ShowProductConfigInvoker {
	requestDef := GenReqDefForShowProductConfig()
	return &ShowProductConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevice 修改设备
//
// 修改设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// UpdateDeviceInvoker 修改设备
func (c *IoTEdgeClient) UpdateDeviceInvoker(request *model.UpdateDeviceRequest) *UpdateDeviceInvoker {
	requestDef := GenReqDefForUpdateDevice()
	return &UpdateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAppConfigsTemplates 添加应用配置模板
//
// 添加应用配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) AddAppConfigsTemplates(request *model.AddAppConfigsTemplatesRequest) (*model.AddAppConfigsTemplatesResponse, error) {
	requestDef := GenReqDefForAddAppConfigsTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAppConfigsTemplatesResponse), nil
	}
}

// AddAppConfigsTemplatesInvoker 添加应用配置模板
func (c *IoTEdgeClient) AddAppConfigsTemplatesInvoker(request *model.AddAppConfigsTemplatesRequest) *AddAppConfigsTemplatesInvoker {
	requestDef := GenReqDefForAddAppConfigsTemplates()
	return &AddAppConfigsTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGeneralAppConfigsTemplate 导入标准应用配置模板
//
// 导入标准应用配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) AddGeneralAppConfigsTemplate(request *model.AddGeneralAppConfigsTemplateRequest) (*model.AddGeneralAppConfigsTemplateResponse, error) {
	requestDef := GenReqDefForAddGeneralAppConfigsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGeneralAppConfigsTemplateResponse), nil
	}
}

// AddGeneralAppConfigsTemplateInvoker 导入标准应用配置模板
func (c *IoTEdgeClient) AddGeneralAppConfigsTemplateInvoker(request *model.AddGeneralAppConfigsTemplateRequest) *AddGeneralAppConfigsTemplateInvoker {
	requestDef := GenReqDefForAddGeneralAppConfigsTemplate()
	return &AddGeneralAppConfigsTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListAppConfigsTemplates 查询应用配置模板列表
//
// 查询应用配置模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListAppConfigsTemplates(request *model.BatchListAppConfigsTemplatesRequest) (*model.BatchListAppConfigsTemplatesResponse, error) {
	requestDef := GenReqDefForBatchListAppConfigsTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListAppConfigsTemplatesResponse), nil
	}
}

// BatchListAppConfigsTemplatesInvoker 查询应用配置模板列表
func (c *IoTEdgeClient) BatchListAppConfigsTemplatesInvoker(request *model.BatchListAppConfigsTemplatesRequest) *BatchListAppConfigsTemplatesInvoker {
	requestDef := GenReqDefForBatchListAppConfigsTemplates()
	return &BatchListAppConfigsTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppConfigsTemplate 删除应用配置模板
//
// 删除应用配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteAppConfigsTemplate(request *model.DeleteAppConfigsTemplateRequest) (*model.DeleteAppConfigsTemplateResponse, error) {
	requestDef := GenReqDefForDeleteAppConfigsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppConfigsTemplateResponse), nil
	}
}

// DeleteAppConfigsTemplateInvoker 删除应用配置模板
func (c *IoTEdgeClient) DeleteAppConfigsTemplateInvoker(request *model.DeleteAppConfigsTemplateRequest) *DeleteAppConfigsTemplateInvoker {
	requestDef := GenReqDefForDeleteAppConfigsTemplate()
	return &DeleteAppConfigsTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppConfigsTemplate 查询应用配置模板详情
//
// 查询应用配置模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowAppConfigsTemplate(request *model.ShowAppConfigsTemplateRequest) (*model.ShowAppConfigsTemplateResponse, error) {
	requestDef := GenReqDefForShowAppConfigsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppConfigsTemplateResponse), nil
	}
}

// ShowAppConfigsTemplateInvoker 查询应用配置模板详情
func (c *IoTEdgeClient) ShowAppConfigsTemplateInvoker(request *model.ShowAppConfigsTemplateRequest) *ShowAppConfigsTemplateInvoker {
	requestDef := GenReqDefForShowAppConfigsTemplate()
	return &ShowAppConfigsTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListEdgeApps 查询应用列表
//
// 查询应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListEdgeApps(request *model.BatchListEdgeAppsRequest) (*model.BatchListEdgeAppsResponse, error) {
	requestDef := GenReqDefForBatchListEdgeApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListEdgeAppsResponse), nil
	}
}

// BatchListEdgeAppsInvoker 查询应用列表
func (c *IoTEdgeClient) BatchListEdgeAppsInvoker(request *model.BatchListEdgeAppsRequest) *BatchListEdgeAppsInvoker {
	requestDef := GenReqDefForBatchListEdgeApps()
	return &BatchListEdgeAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateEdgeApp(request *model.CreateEdgeAppRequest) (*model.CreateEdgeAppResponse, error) {
	requestDef := GenReqDefForCreateEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeAppResponse), nil
	}
}

// CreateEdgeAppInvoker 创建应用
func (c *IoTEdgeClient) CreateEdgeAppInvoker(request *model.CreateEdgeAppRequest) *CreateEdgeAppInvoker {
	requestDef := GenReqDefForCreateEdgeApp()
	return &CreateEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeApp 删除应用
//
// 删除应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteEdgeApp(request *model.DeleteEdgeAppRequest) (*model.DeleteEdgeAppResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeAppResponse), nil
	}
}

// DeleteEdgeAppInvoker 删除应用
func (c *IoTEdgeClient) DeleteEdgeAppInvoker(request *model.DeleteEdgeAppRequest) *DeleteEdgeAppInvoker {
	requestDef := GenReqDefForDeleteEdgeApp()
	return &DeleteEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeApp 查询应用
//
// 查询应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowEdgeApp(request *model.ShowEdgeAppRequest) (*model.ShowEdgeAppResponse, error) {
	requestDef := GenReqDefForShowEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeAppResponse), nil
	}
}

// ShowEdgeAppInvoker 查询应用
func (c *IoTEdgeClient) ShowEdgeAppInvoker(request *model.ShowEdgeAppRequest) *ShowEdgeAppInvoker {
	requestDef := GenReqDefForShowEdgeApp()
	return &ShowEdgeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListEdgeAppVersions 查询应用版本列表
//
// 查询应用版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListEdgeAppVersions(request *model.BatchListEdgeAppVersionsRequest) (*model.BatchListEdgeAppVersionsResponse, error) {
	requestDef := GenReqDefForBatchListEdgeAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListEdgeAppVersionsResponse), nil
	}
}

// BatchListEdgeAppVersionsInvoker 查询应用版本列表
func (c *IoTEdgeClient) BatchListEdgeAppVersionsInvoker(request *model.BatchListEdgeAppVersionsRequest) *BatchListEdgeAppVersionsInvoker {
	requestDef := GenReqDefForBatchListEdgeAppVersions()
	return &BatchListEdgeAppVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeApplicationVersion 创建应用版本
//
// 创建应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateEdgeApplicationVersion(request *model.CreateEdgeApplicationVersionRequest) (*model.CreateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForCreateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeApplicationVersionResponse), nil
	}
}

// CreateEdgeApplicationVersionInvoker 创建应用版本
func (c *IoTEdgeClient) CreateEdgeApplicationVersionInvoker(request *model.CreateEdgeApplicationVersionRequest) *CreateEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForCreateEdgeApplicationVersion()
	return &CreateEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeApplicationVersion 删除应用版本
//
// 删除应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteEdgeApplicationVersion(request *model.DeleteEdgeApplicationVersionRequest) (*model.DeleteEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeApplicationVersionResponse), nil
	}
}

// DeleteEdgeApplicationVersionInvoker 删除应用版本
func (c *IoTEdgeClient) DeleteEdgeApplicationVersionInvoker(request *model.DeleteEdgeApplicationVersionRequest) *DeleteEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForDeleteEdgeApplicationVersion()
	return &DeleteEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeApplicationVersion 查询应用版本详情
//
// 查询应用版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowEdgeApplicationVersion(request *model.ShowEdgeApplicationVersionRequest) (*model.ShowEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForShowEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeApplicationVersionResponse), nil
	}
}

// ShowEdgeApplicationVersionInvoker 查询应用版本详情
func (c *IoTEdgeClient) ShowEdgeApplicationVersionInvoker(request *model.ShowEdgeApplicationVersionRequest) *ShowEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForShowEdgeApplicationVersion()
	return &ShowEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeApplicationVersion 修改应用版本
//
// 修改应用版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateEdgeApplicationVersion(request *model.UpdateEdgeApplicationVersionRequest) (*model.UpdateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionResponse), nil
	}
}

// UpdateEdgeApplicationVersionInvoker 修改应用版本
func (c *IoTEdgeClient) UpdateEdgeApplicationVersionInvoker(request *model.UpdateEdgeApplicationVersionRequest) *UpdateEdgeApplicationVersionInvoker {
	requestDef := GenReqDefForUpdateEdgeApplicationVersion()
	return &UpdateEdgeApplicationVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeApplicationVersionState 更新应用版本状态
//
// 更新应用版本状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateEdgeApplicationVersionState(request *model.UpdateEdgeApplicationVersionStateRequest) (*model.UpdateEdgeApplicationVersionStateResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersionState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionStateResponse), nil
	}
}

// UpdateEdgeApplicationVersionStateInvoker 更新应用版本状态
func (c *IoTEdgeClient) UpdateEdgeApplicationVersionStateInvoker(request *model.UpdateEdgeApplicationVersionStateRequest) *UpdateEdgeApplicationVersionStateInvoker {
	requestDef := GenReqDefForUpdateEdgeApplicationVersionState()
	return &UpdateEdgeApplicationVersionStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListDcDs 查询数据源配置列表
//
// 查询数据源配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListDcDs(request *model.BatchListDcDsRequest) (*model.BatchListDcDsResponse, error) {
	requestDef := GenReqDefForBatchListDcDs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListDcDsResponse), nil
	}
}

// BatchListDcDsInvoker 查询数据源配置列表
func (c *IoTEdgeClient) BatchListDcDsInvoker(request *model.BatchListDcDsRequest) *BatchListDcDsInvoker {
	requestDef := GenReqDefForBatchListDcDs()
	return &BatchListDcDsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDs 创建数据源配置
//
// 用户通过Console接口在指定边缘节点上创建数据源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateDs(request *model.CreateDsRequest) (*model.CreateDsResponse, error) {
	requestDef := GenReqDefForCreateDs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDsResponse), nil
	}
}

// CreateDsInvoker 创建数据源配置
func (c *IoTEdgeClient) CreateDsInvoker(request *model.CreateDsRequest) *CreateDsInvoker {
	requestDef := GenReqDefForCreateDs()
	return &CreateDsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDcDs 删除数据源配置
//
// 删除数据源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteDcDs(request *model.DeleteDcDsRequest) (*model.DeleteDcDsResponse, error) {
	requestDef := GenReqDefForDeleteDcDs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDcDsResponse), nil
	}
}

// DeleteDcDsInvoker 删除数据源配置
func (c *IoTEdgeClient) DeleteDcDsInvoker(request *model.DeleteDcDsRequest) *DeleteDcDsInvoker {
	requestDef := GenReqDefForDeleteDcDs()
	return &DeleteDcDsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDcDs 查询数据源配置
//
// 查询数据源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowDcDs(request *model.ShowDcDsRequest) (*model.ShowDcDsResponse, error) {
	requestDef := GenReqDefForShowDcDs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDcDsResponse), nil
	}
}

// ShowDcDsInvoker 查询数据源配置
func (c *IoTEdgeClient) ShowDcDsInvoker(request *model.ShowDcDsRequest) *ShowDcDsInvoker {
	requestDef := GenReqDefForShowDcDs()
	return &ShowDcDsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SynchronizeDcConfigs 下发数采配置
//
// 下发数采配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) SynchronizeDcConfigs(request *model.SynchronizeDcConfigsRequest) (*model.SynchronizeDcConfigsResponse, error) {
	requestDef := GenReqDefForSynchronizeDcConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SynchronizeDcConfigsResponse), nil
	}
}

// SynchronizeDcConfigsInvoker 下发数采配置
func (c *IoTEdgeClient) SynchronizeDcConfigsInvoker(request *model.SynchronizeDcConfigsRequest) *SynchronizeDcConfigsInvoker {
	requestDef := GenReqDefForSynchronizeDcConfigs()
	return &SynchronizeDcConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDcDs 修改数据源配置
//
// 修改数据源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateDcDs(request *model.UpdateDcDsRequest) (*model.UpdateDcDsResponse, error) {
	requestDef := GenReqDefForUpdateDcDs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDcDsResponse), nil
	}
}

// UpdateDcDsInvoker 修改数据源配置
func (c *IoTEdgeClient) UpdateDcDsInvoker(request *model.UpdateDcDsRequest) *UpdateDcDsInvoker {
	requestDef := GenReqDefForUpdateDcDs()
	return &UpdateDcDsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListDcDevices 查数采连接子设备列表
//
// 查询数采连接下子设备列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListDcDevices(request *model.BatchListDcDevicesRequest) (*model.BatchListDcDevicesResponse, error) {
	requestDef := GenReqDefForBatchListDcDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListDcDevicesResponse), nil
	}
}

// BatchListDcDevicesInvoker 查数采连接子设备列表
func (c *IoTEdgeClient) BatchListDcDevicesInvoker(request *model.BatchListDcDevicesRequest) *BatchListDcDevicesInvoker {
	requestDef := GenReqDefForBatchListDcDevices()
	return &BatchListDcDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListDcPoints 查询点位表配置列表
//
// 查询点位表配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListDcPoints(request *model.BatchListDcPointsRequest) (*model.BatchListDcPointsResponse, error) {
	requestDef := GenReqDefForBatchListDcPoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListDcPointsResponse), nil
	}
}

// BatchListDcPointsInvoker 查询点位表配置列表
func (c *IoTEdgeClient) BatchListDcPointsInvoker(request *model.BatchListDcPointsRequest) *BatchListDcPointsInvoker {
	requestDef := GenReqDefForBatchListDcPoints()
	return &BatchListDcPointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDcPoint 创建点位表配置
//
// 用户通过Console接口在指定边缘节点上点位表配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateDcPoint(request *model.CreateDcPointRequest) (*model.CreateDcPointResponse, error) {
	requestDef := GenReqDefForCreateDcPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDcPointResponse), nil
	}
}

// CreateDcPointInvoker 创建点位表配置
func (c *IoTEdgeClient) CreateDcPointInvoker(request *model.CreateDcPointRequest) *CreateDcPointInvoker {
	requestDef := GenReqDefForCreateDcPoint()
	return &CreateDcPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDcPoint 删除点位表配置
//
// 删除点位表配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteDcPoint(request *model.DeleteDcPointRequest) (*model.DeleteDcPointResponse, error) {
	requestDef := GenReqDefForDeleteDcPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDcPointResponse), nil
	}
}

// DeleteDcPointInvoker 删除点位表配置
func (c *IoTEdgeClient) DeleteDcPointInvoker(request *model.DeleteDcPointRequest) *DeleteDcPointInvoker {
	requestDef := GenReqDefForDeleteDcPoint()
	return &DeleteDcPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDcPoints 批量删除点位表配置
//
// 批量删除点位表配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteDcPoints(request *model.DeleteDcPointsRequest) (*model.DeleteDcPointsResponse, error) {
	requestDef := GenReqDefForDeleteDcPoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDcPointsResponse), nil
	}
}

// DeleteDcPointsInvoker 批量删除点位表配置
func (c *IoTEdgeClient) DeleteDcPointsInvoker(request *model.DeleteDcPointsRequest) *DeleteDcPointsInvoker {
	requestDef := GenReqDefForDeleteDcPoints()
	return &DeleteDcPointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDcPoint 查询点位表配置详情
//
// 查询点位表配置详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowDcPoint(request *model.ShowDcPointRequest) (*model.ShowDcPointResponse, error) {
	requestDef := GenReqDefForShowDcPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDcPointResponse), nil
	}
}

// ShowDcPointInvoker 查询点位表配置详情
func (c *IoTEdgeClient) ShowDcPointInvoker(request *model.ShowDcPointRequest) *ShowDcPointInvoker {
	requestDef := GenReqDefForShowDcPoint()
	return &ShowDcPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDcPoint 修改点位表配置
//
// 修改点位表配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateDcPoint(request *model.UpdateDcPointRequest) (*model.UpdateDcPointResponse, error) {
	requestDef := GenReqDefForUpdateDcPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDcPointResponse), nil
	}
}

// UpdateDcPointInvoker 修改点位表配置
func (c *IoTEdgeClient) UpdateDcPointInvoker(request *model.UpdateDcPointRequest) *UpdateDcPointInvoker {
	requestDef := GenReqDefForUpdateDcPoint()
	return &UpdateDcPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExternalEntity 在指定节点上创建外部实体
//
// 用户通过在指定边缘节点上设置外部实体的接入信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateExternalEntity(request *model.CreateExternalEntityRequest) (*model.CreateExternalEntityResponse, error) {
	requestDef := GenReqDefForCreateExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExternalEntityResponse), nil
	}
}

// CreateExternalEntityInvoker 在指定节点上创建外部实体
func (c *IoTEdgeClient) CreateExternalEntityInvoker(request *model.CreateExternalEntityRequest) *CreateExternalEntityInvoker {
	requestDef := GenReqDefForCreateExternalEntity()
	return &CreateExternalEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteExternalEntity 删除指定节点下外部实体
//
// 删除节点下外部实体
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteExternalEntity(request *model.DeleteExternalEntityRequest) (*model.DeleteExternalEntityResponse, error) {
	requestDef := GenReqDefForDeleteExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExternalEntityResponse), nil
	}
}

// DeleteExternalEntityInvoker 删除指定节点下外部实体
func (c *IoTEdgeClient) DeleteExternalEntityInvoker(request *model.DeleteExternalEntityRequest) *DeleteExternalEntityInvoker {
	requestDef := GenReqDefForDeleteExternalEntity()
	return &DeleteExternalEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExternalEntity 查询指定边缘节点下的外部实体
//
// 用户在指定边缘节点上查询外部实体列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListExternalEntity(request *model.ListExternalEntityRequest) (*model.ListExternalEntityResponse, error) {
	requestDef := GenReqDefForListExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExternalEntityResponse), nil
	}
}

// ListExternalEntityInvoker 查询指定边缘节点下的外部实体
func (c *IoTEdgeClient) ListExternalEntityInvoker(request *model.ListExternalEntityRequest) *ListExternalEntityInvoker {
	requestDef := GenReqDefForListExternalEntity()
	return &ListExternalEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateExternalEntity 修改节点下指定的外部实体信息
//
// 用户通过在指定边缘节点上修改指定外部实体的接入信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateExternalEntity(request *model.UpdateExternalEntityRequest) (*model.UpdateExternalEntityResponse, error) {
	requestDef := GenReqDefForUpdateExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExternalEntityResponse), nil
	}
}

// UpdateExternalEntityInvoker 修改节点下指定的外部实体信息
func (c *IoTEdgeClient) UpdateExternalEntityInvoker(request *model.UpdateExternalEntityRequest) *UpdateExternalEntityInvoker {
	requestDef := GenReqDefForUpdateExternalEntity()
	return &UpdateExternalEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListModules 查询边缘模块列表
//
// 用户通过Console接口查询指定边缘节点上边缘模块列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListModules(request *model.BatchListModulesRequest) (*model.BatchListModulesResponse, error) {
	requestDef := GenReqDefForBatchListModules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListModulesResponse), nil
	}
}

// BatchListModulesInvoker 查询边缘模块列表
func (c *IoTEdgeClient) BatchListModulesInvoker(request *model.BatchListModulesRequest) *BatchListModulesInvoker {
	requestDef := GenReqDefForBatchListModules()
	return &BatchListModulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModule 创建边缘模块
//
// 用户通过Console接口在指定边缘节点上创建边缘模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateModule(request *model.CreateModuleRequest) (*model.CreateModuleResponse, error) {
	requestDef := GenReqDefForCreateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModuleResponse), nil
	}
}

// CreateModuleInvoker 创建边缘模块
func (c *IoTEdgeClient) CreateModuleInvoker(request *model.CreateModuleRequest) *CreateModuleInvoker {
	requestDef := GenReqDefForCreateModule()
	return &CreateModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModule 删除边缘模块
//
// 用户通过过Console接口在指定边缘节点上删除边缘模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteModule(request *model.DeleteModuleRequest) (*model.DeleteModuleResponse, error) {
	requestDef := GenReqDefForDeleteModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModuleResponse), nil
	}
}

// DeleteModuleInvoker 删除边缘模块
func (c *IoTEdgeClient) DeleteModuleInvoker(request *model.DeleteModuleRequest) *DeleteModuleInvoker {
	requestDef := GenReqDefForDeleteModule()
	return &DeleteModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InvokeModuleMsg 代理边缘模块消息
//
// iotedge通过该接口透明代理用户到模块的请求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) InvokeModuleMsg(request *model.InvokeModuleMsgRequest) (*model.InvokeModuleMsgResponse, error) {
	requestDef := GenReqDefForInvokeModuleMsg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeModuleMsgResponse), nil
	}
}

// InvokeModuleMsgInvoker 代理边缘模块消息
func (c *IoTEdgeClient) InvokeModuleMsgInvoker(request *model.InvokeModuleMsgRequest) *InvokeModuleMsgInvoker {
	requestDef := GenReqDefForInvokeModuleMsg()
	return &InvokeModuleMsgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModule 查询边缘模块
//
// 用户通过Console接口查询指定边缘节点上指定边缘模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowModule(request *model.ShowModuleRequest) (*model.ShowModuleResponse, error) {
	requestDef := GenReqDefForShowModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModuleResponse), nil
	}
}

// ShowModuleInvoker 查询边缘模块
func (c *IoTEdgeClient) ShowModuleInvoker(request *model.ShowModuleRequest) *ShowModuleInvoker {
	requestDef := GenReqDefForShowModule()
	return &ShowModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModule 修改边缘模块
//
// 用户通过Console接口查询指定边缘节点上指定边缘模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateModule(request *model.UpdateModuleRequest) (*model.UpdateModuleResponse, error) {
	requestDef := GenReqDefForUpdateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModuleResponse), nil
	}
}

// UpdateModuleInvoker 修改边缘模块
func (c *IoTEdgeClient) UpdateModuleInvoker(request *model.UpdateModuleRequest) *UpdateModuleInvoker {
	requestDef := GenReqDefForUpdateModule()
	return &UpdateModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModuleState 修改边缘模块状态
//
// 用户通过Console接口启停数采连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateModuleState(request *model.UpdateModuleStateRequest) (*model.UpdateModuleStateResponse, error) {
	requestDef := GenReqDefForUpdateModuleState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModuleStateResponse), nil
	}
}

// UpdateModuleStateInvoker 修改边缘模块状态
func (c *IoTEdgeClient) UpdateModuleStateInvoker(request *model.UpdateModuleStateRequest) *UpdateModuleStateInvoker {
	requestDef := GenReqDefForUpdateModuleState()
	return &UpdateModuleStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModuleShadow 获取模块影子
//
// 获取模块影子信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowModuleShadow(request *model.ShowModuleShadowRequest) (*model.ShowModuleShadowResponse, error) {
	requestDef := GenReqDefForShowModuleShadow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModuleShadowResponse), nil
	}
}

// ShowModuleShadowInvoker 获取模块影子
func (c *IoTEdgeClient) ShowModuleShadowInvoker(request *model.ShowModuleShadowRequest) *ShowModuleShadowInvoker {
	requestDef := GenReqDefForShowModuleShadow()
	return &ShowModuleShadowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModuleShadow 更新模块影子
//
// 更新模块影子信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateModuleShadow(request *model.UpdateModuleShadowRequest) (*model.UpdateModuleShadowResponse, error) {
	requestDef := GenReqDefForUpdateModuleShadow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModuleShadowResponse), nil
	}
}

// UpdateModuleShadowInvoker 更新模块影子
func (c *IoTEdgeClient) UpdateModuleShadowInvoker(request *model.UpdateModuleShadowRequest) *UpdateModuleShadowInvoker {
	requestDef := GenReqDefForUpdateModuleShadow()
	return &UpdateModuleShadowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoutes 查询边缘路由列表
//
// 用户在指定边缘节点上查询边缘路由列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListRoutes(request *model.ListRoutesRequest) (*model.ListRoutesResponse, error) {
	requestDef := GenReqDefForListRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoutesResponse), nil
	}
}

// ListRoutesInvoker 查询边缘路由列表
func (c *IoTEdgeClient) ListRoutesInvoker(request *model.ListRoutesRequest) *ListRoutesInvoker {
	requestDef := GenReqDefForListRoutes()
	return &ListRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoutes 设置边缘路由
//
// 用户通过在指定边缘节点上设置边缘路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateRoutes(request *model.UpdateRoutesRequest) (*model.UpdateRoutesResponse, error) {
	requestDef := GenReqDefForUpdateRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoutesResponse), nil
	}
}

// UpdateRoutesInvoker 设置边缘路由
func (c *IoTEdgeClient) UpdateRoutesInvoker(request *model.UpdateRoutesRequest) *UpdateRoutesInvoker {
	requestDef := GenReqDefForUpdateRoutes()
	return &UpdateRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGeneralOtTemplate 导入标准数采模板
//
// 导入标准数采模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) AddGeneralOtTemplate(request *model.AddGeneralOtTemplateRequest) (*model.AddGeneralOtTemplateResponse, error) {
	requestDef := GenReqDefForAddGeneralOtTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGeneralOtTemplateResponse), nil
	}
}

// AddGeneralOtTemplateInvoker 导入标准数采模板
func (c *IoTEdgeClient) AddGeneralOtTemplateInvoker(request *model.AddGeneralOtTemplateRequest) *AddGeneralOtTemplateInvoker {
	requestDef := GenReqDefForAddGeneralOtTemplate()
	return &AddGeneralOtTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddOtTemplates 添加数采模板
//
// 添加数采模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) AddOtTemplates(request *model.AddOtTemplatesRequest) (*model.AddOtTemplatesResponse, error) {
	requestDef := GenReqDefForAddOtTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddOtTemplatesResponse), nil
	}
}

// AddOtTemplatesInvoker 添加数采模板
func (c *IoTEdgeClient) AddOtTemplatesInvoker(request *model.AddOtTemplatesRequest) *AddOtTemplatesInvoker {
	requestDef := GenReqDefForAddOtTemplates()
	return &AddOtTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListOtTemplates 查询数采模板列表
//
// 查询数采模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchListOtTemplates(request *model.BatchListOtTemplatesRequest) (*model.BatchListOtTemplatesResponse, error) {
	requestDef := GenReqDefForBatchListOtTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListOtTemplatesResponse), nil
	}
}

// BatchListOtTemplatesInvoker 查询数采模板列表
func (c *IoTEdgeClient) BatchListOtTemplatesInvoker(request *model.BatchListOtTemplatesRequest) *BatchListOtTemplatesInvoker {
	requestDef := GenReqDefForBatchListOtTemplates()
	return &BatchListOtTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOtTemplate 删除数采模板
//
// 删除数采模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteOtTemplate(request *model.DeleteOtTemplateRequest) (*model.DeleteOtTemplateResponse, error) {
	requestDef := GenReqDefForDeleteOtTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOtTemplateResponse), nil
	}
}

// DeleteOtTemplateInvoker 删除数采模板
func (c *IoTEdgeClient) DeleteOtTemplateInvoker(request *model.DeleteOtTemplateRequest) *DeleteOtTemplateInvoker {
	requestDef := GenReqDefForDeleteOtTemplate()
	return &DeleteOtTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOtTemplate 查询数采模板详情
//
// 查询数采模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowOtTemplate(request *model.ShowOtTemplateRequest) (*model.ShowOtTemplateResponse, error) {
	requestDef := GenReqDefForShowOtTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOtTemplateResponse), nil
	}
}

// ShowOtTemplateInvoker 查询数采模板详情
func (c *IoTEdgeClient) ShowOtTemplateInvoker(request *model.ShowOtTemplateRequest) *ShowOtTemplateInvoker {
	requestDef := GenReqDefForShowOtTemplate()
	return &ShowOtTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportPoints 批量导入点位表
//
// 用户通过Console接口在指定边缘节点上点位表配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ImportPoints(request *model.ImportPointsRequest) (*model.ImportPointsResponse, error) {
	requestDef := GenReqDefForImportPoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportPointsResponse), nil
	}
}

// ImportPointsInvoker 批量导入点位表
func (c *IoTEdgeClient) ImportPointsInvoker(request *model.ImportPointsRequest) *ImportPointsInvoker {
	requestDef := GenReqDefForImportPoints()
	return &ImportPointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPointTemplate 查询点位表模板文件
//
// 查询点位表模板文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowPointTemplate(request *model.ShowPointTemplateRequest) (*model.ShowPointTemplateResponse, error) {
	requestDef := GenReqDefForShowPointTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPointTemplateResponse), nil
	}
}

// ShowPointTemplateInvoker 查询点位表模板文件
func (c *IoTEdgeClient) ShowPointTemplateInvoker(request *model.ShowPointTemplateRequest) *ShowPointTemplateInvoker {
	requestDef := GenReqDefForShowPointTemplate()
	return &ShowPointTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoints 查询点位表模板文件
//
// 查询点位表模板文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowPoints(request *model.ShowPointsRequest) (*model.ShowPointsResponse, error) {
	requestDef := GenReqDefForShowPoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPointsResponse), nil
	}
}

// ShowPointsInvoker 查询点位表模板文件
func (c *IoTEdgeClient) ShowPointsInvoker(request *model.ShowPointsRequest) *ShowPointsInvoker {
	requestDef := GenReqDefForShowPoints()
	return &ShowPointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSchedule 创建调度计划
//
// 用户通过北向接口在指定边缘节点上创建调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateSchedule(request *model.CreateScheduleRequest) (*model.CreateScheduleResponse, error) {
	requestDef := GenReqDefForCreateSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduleResponse), nil
	}
}

// CreateScheduleInvoker 创建调度计划
func (c *IoTEdgeClient) CreateScheduleInvoker(request *model.CreateScheduleRequest) *CreateScheduleInvoker {
	requestDef := GenReqDefForCreateSchedule()
	return &CreateScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSchedule 删除调度计划
//
// 用户通过北向接口删除边缘节点上调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteSchedule(request *model.DeleteScheduleRequest) (*model.DeleteScheduleResponse, error) {
	requestDef := GenReqDefForDeleteSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduleResponse), nil
	}
}

// DeleteScheduleInvoker 删除调度计划
func (c *IoTEdgeClient) DeleteScheduleInvoker(request *model.DeleteScheduleRequest) *DeleteScheduleInvoker {
	requestDef := GenReqDefForDeleteSchedule()
	return &DeleteScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSchedule 更新调度计划，机机接口，全量更新字段
//
// 用户通过北向接口修改边缘节点上调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateSchedule(request *model.UpdateScheduleRequest) (*model.UpdateScheduleResponse, error) {
	requestDef := GenReqDefForUpdateSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduleResponse), nil
	}
}

// UpdateScheduleInvoker 更新调度计划，机机接口，全量更新字段
func (c *IoTEdgeClient) UpdateScheduleInvoker(request *model.UpdateScheduleRequest) *UpdateScheduleInvoker {
	requestDef := GenReqDefForUpdateSchedule()
	return &UpdateScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchConfirmConfigsNew 批量确认南向3rdIA配置项
//
// 南向3rdIA对下发的配置项进行批量确认
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchConfirmConfigsNew(request *model.BatchConfirmConfigsNewRequest) (*model.BatchConfirmConfigsNewResponse, error) {
	requestDef := GenReqDefForBatchConfirmConfigsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchConfirmConfigsNewResponse), nil
	}
}

// BatchConfirmConfigsNewInvoker 批量确认南向3rdIA配置项
func (c *IoTEdgeClient) BatchConfirmConfigsNewInvoker(request *model.BatchConfirmConfigsNewRequest) *BatchConfirmConfigsNewInvoker {
	requestDef := GenReqDefForBatchConfirmConfigsNew()
	return &BatchConfirmConfigsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchImportConfigs 批量导入南向3rdIA配置项
//
// 批量导入南向3rdIA配置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchImportConfigs(request *model.BatchImportConfigsRequest) (*model.BatchImportConfigsResponse, error) {
	requestDef := GenReqDefForBatchImportConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportConfigsResponse), nil
	}
}

// BatchImportConfigsInvoker 批量导入南向3rdIA配置项
func (c *IoTEdgeClient) BatchImportConfigsInvoker(request *model.BatchImportConfigsRequest) *BatchImportConfigsInvoker {
	requestDef := GenReqDefForBatchImportConfigs()
	return &BatchImportConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIaConfig 删除南向3rdIA配置项
//
// 删除南向3rdIA配置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteIaConfig(request *model.DeleteIaConfigRequest) (*model.DeleteIaConfigResponse, error) {
	requestDef := GenReqDefForDeleteIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIaConfigResponse), nil
	}
}

// DeleteIaConfigInvoker 删除南向3rdIA配置项
func (c *IoTEdgeClient) DeleteIaConfigInvoker(request *model.DeleteIaConfigRequest) *DeleteIaConfigInvoker {
	requestDef := GenReqDefForDeleteIaConfig()
	return &DeleteIaConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIaConfigs 查询南向3rdIA配置项列表
//
// 查询南向3rdIA配置项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListIaConfigs(request *model.ListIaConfigsRequest) (*model.ListIaConfigsResponse, error) {
	requestDef := GenReqDefForListIaConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIaConfigsResponse), nil
	}
}

// ListIaConfigsInvoker 查询南向3rdIA配置项列表
func (c *IoTEdgeClient) ListIaConfigsInvoker(request *model.ListIaConfigsRequest) *ListIaConfigsInvoker {
	requestDef := GenReqDefForListIaConfigs()
	return &ListIaConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIaConfig 查询南向3rdIA配置项详情
//
// 查询南向3rdIA配置项详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowIaConfig(request *model.ShowIaConfigRequest) (*model.ShowIaConfigResponse, error) {
	requestDef := GenReqDefForShowIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIaConfigResponse), nil
	}
}

// ShowIaConfigInvoker 查询南向3rdIA配置项详情
func (c *IoTEdgeClient) ShowIaConfigInvoker(request *model.ShowIaConfigRequest) *ShowIaConfigInvoker {
	requestDef := GenReqDefForShowIaConfig()
	return &ShowIaConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIaConfig 创建&amp;更新南向3rdIA配置项信息
//
// 创建&amp;更新南向3rdIA配置项信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateIaConfig(request *model.UpdateIaConfigRequest) (*model.UpdateIaConfigResponse, error) {
	requestDef := GenReqDefForUpdateIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIaConfigResponse), nil
	}
}

// UpdateIaConfigInvoker 创建&amp;更新南向3rdIA配置项信息
func (c *IoTEdgeClient) UpdateIaConfigInvoker(request *model.UpdateIaConfigRequest) *UpdateIaConfigInvoker {
	requestDef := GenReqDefForUpdateIaConfig()
	return &UpdateIaConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateNaToNodes 授权北向NA信息到边缘节点
//
// 批量授权北向NA信息到边缘节点。
// 已授权的边缘节点上的南向IA应用，可以通过部署在边缘节点上的api网关访问北向NA提供的接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) BatchAssociateNaToNodes(request *model.BatchAssociateNaToNodesRequest) (*model.BatchAssociateNaToNodesResponse, error) {
	requestDef := GenReqDefForBatchAssociateNaToNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateNaToNodesResponse), nil
	}
}

// BatchAssociateNaToNodesInvoker 授权北向NA信息到边缘节点
func (c *IoTEdgeClient) BatchAssociateNaToNodesInvoker(request *model.BatchAssociateNaToNodesRequest) *BatchAssociateNaToNodesInvoker {
	requestDef := GenReqDefForBatchAssociateNaToNodes()
	return &BatchAssociateNaToNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNa 删除北向NA信息
//
// 删除北向NA信息，如果有边缘节点已分配该NA信息，会通知到该边缘节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteNa(request *model.DeleteNaRequest) (*model.DeleteNaResponse, error) {
	requestDef := GenReqDefForDeleteNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNaResponse), nil
	}
}

// DeleteNaInvoker 删除北向NA信息
func (c *IoTEdgeClient) DeleteNaInvoker(request *model.DeleteNaRequest) *DeleteNaInvoker {
	requestDef := GenReqDefForDeleteNa()
	return &DeleteNaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNaAuthorizedNodes 查询该北向NA信息的已分配节点
//
// 查询该北向NA信息的已分配节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListNaAuthorizedNodes(request *model.ListNaAuthorizedNodesRequest) (*model.ListNaAuthorizedNodesResponse, error) {
	requestDef := GenReqDefForListNaAuthorizedNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNaAuthorizedNodesResponse), nil
	}
}

// ListNaAuthorizedNodesInvoker 查询该北向NA信息的已分配节点
func (c *IoTEdgeClient) ListNaAuthorizedNodesInvoker(request *model.ListNaAuthorizedNodesRequest) *ListNaAuthorizedNodesInvoker {
	requestDef := GenReqDefForListNaAuthorizedNodes()
	return &ListNaAuthorizedNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNas 查询北向NA信息列表
//
// 查询北向NA信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListNas(request *model.ListNasRequest) (*model.ListNasResponse, error) {
	requestDef := GenReqDefForListNas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNasResponse), nil
	}
}

// ListNasInvoker 查询北向NA信息列表
func (c *IoTEdgeClient) ListNasInvoker(request *model.ListNasRequest) *ListNasInvoker {
	requestDef := GenReqDefForListNas()
	return &ListNasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNa 查询北向NA信息详情
//
// 查询北向NA信息详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowNa(request *model.ShowNaRequest) (*model.ShowNaResponse, error) {
	requestDef := GenReqDefForShowNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNaResponse), nil
	}
}

// ShowNaInvoker 查询北向NA信息详情
func (c *IoTEdgeClient) ShowNaInvoker(request *model.ShowNaRequest) *ShowNaInvoker {
	requestDef := GenReqDefForShowNa()
	return &ShowNaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNa 创建&amp;更新北向NA信息
//
// 创建&amp;更新北向NA信息，当更新北向NA信息时，会通知到已分配该北向NA的所有边缘节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateNa(request *model.UpdateNaRequest) (*model.UpdateNaResponse, error) {
	requestDef := GenReqDefForUpdateNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNaResponse), nil
	}
}

// UpdateNaInvoker 创建&amp;更新北向NA信息
func (c *IoTEdgeClient) UpdateNaInvoker(request *model.UpdateNaRequest) *UpdateNaInvoker {
	requestDef := GenReqDefForUpdateNa()
	return &UpdateNaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
