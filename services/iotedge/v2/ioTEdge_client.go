package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotedge/v2/model"
)

type IoTEdgeClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIoTEdgeClient(hcClient *http_client.HcHttpClient) *IoTEdgeClient {
	return &IoTEdgeClient{HcClient: hcClient}
}

func IoTEdgeClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateEdgeNode 创建边缘节点
//
// 创建边缘节点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// AddDevice 添加设备
//
// 添加设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// BatchListEdgeApps 查询应用列表
//
// 查询应用列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateExternalEntity 在指定节点上创建外部实体
//
// 用户通过在指定边缘节点上设置外部实体的接入信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowModule 查询边缘模块
//
// 用户通过Console接口查询指定边缘节点上指定边缘模块
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListRoutes 查询边缘路由列表
//
// 用户在指定边缘节点上查询边缘路由列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// BatchConfirmConfigsNew 批量确认南向3rdIA配置项
//
// 南向3rdIA对下发的配置项进行批量确认
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
