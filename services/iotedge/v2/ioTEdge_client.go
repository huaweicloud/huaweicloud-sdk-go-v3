package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/iotedge/v2/model"
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

//创建边缘节点
func (c *IoTEdgeClient) CreateEdgeNode(request *model.CreateEdgeNodeRequest) (*model.CreateEdgeNodeResponse, error) {
	requestDef := GenReqDefForCreateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeResponse), nil
	}
}

//生成边缘节点安装命令，命令有效时间30分钟，超过后需要重新生成
func (c *IoTEdgeClient) CreateInstallCmd(request *model.CreateInstallCmdRequest) (*model.CreateInstallCmdResponse, error) {
	requestDef := GenReqDefForCreateInstallCmd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstallCmdResponse), nil
	}
}

//删除指定边缘节点
func (c *IoTEdgeClient) DeleteEdgeNode(request *model.DeleteEdgeNodeRequest) (*model.DeleteEdgeNodeResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeResponse), nil
	}
}

//查询边缘节点列表
func (c *IoTEdgeClient) ListEdgeNodes(request *model.ListEdgeNodesRequest) (*model.ListEdgeNodesResponse, error) {
	requestDef := GenReqDefForListEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodesResponse), nil
	}
}

//查询边缘节点详情
func (c *IoTEdgeClient) ShowEdgeNode(request *model.ShowEdgeNodeRequest) (*model.ShowEdgeNodeResponse, error) {
	requestDef := GenReqDefForShowEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeResponse), nil
	}
}

//添加设备
func (c *IoTEdgeClient) AddDevice(request *model.AddDeviceRequest) (*model.AddDeviceResponse, error) {
	requestDef := GenReqDefForAddDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceResponse), nil
	}
}

//批量修改产品关联的设备，传入product_id修改该产品下所有设备，传入device_id列表，根据device_id修改,两者互斥。
func (c *IoTEdgeClient) BatchUpdateConfigs(request *model.BatchUpdateConfigsRequest) (*model.BatchUpdateConfigsResponse, error) {
	requestDef := GenReqDefForBatchUpdateConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateConfigsResponse), nil
	}
}

//生成modbus协议设备接入码
func (c *IoTEdgeClient) CreateAccessCode(request *model.CreateAccessCodeRequest) (*model.CreateAccessCodeResponse, error) {
	requestDef := GenReqDefForCreateAccessCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessCodeResponse), nil
	}
}

//删除设备
func (c *IoTEdgeClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

//查询设备列表
func (c *IoTEdgeClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

//获取协议配置
func (c *IoTEdgeClient) ShowProductConfig(request *model.ShowProductConfigRequest) (*model.ShowProductConfigResponse, error) {
	requestDef := GenReqDefForShowProductConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductConfigResponse), nil
	}
}

//获取协议映射文件
func (c *IoTEdgeClient) ShowProtocolMappings(request *model.ShowProtocolMappingsRequest) (*model.ShowProtocolMappingsResponse, error) {
	requestDef := GenReqDefForShowProtocolMappings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtocolMappingsResponse), nil
	}
}

//修改设备
func (c *IoTEdgeClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

//上传协议映射文件
func (c *IoTEdgeClient) UploadProtocolMappings(request *model.UploadProtocolMappingsRequest) (*model.UploadProtocolMappingsResponse, error) {
	requestDef := GenReqDefForUploadProtocolMappings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadProtocolMappingsResponse), nil
	}
}

//查询应用列表
func (c *IoTEdgeClient) BatchListEdgeApps(request *model.BatchListEdgeAppsRequest) (*model.BatchListEdgeAppsResponse, error) {
	requestDef := GenReqDefForBatchListEdgeApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListEdgeAppsResponse), nil
	}
}

//创建应用
func (c *IoTEdgeClient) CreateEdgeApp(request *model.CreateEdgeAppRequest) (*model.CreateEdgeAppResponse, error) {
	requestDef := GenReqDefForCreateEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeAppResponse), nil
	}
}

//删除应用
func (c *IoTEdgeClient) DeleteEdgeApp(request *model.DeleteEdgeAppRequest) (*model.DeleteEdgeAppResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeAppResponse), nil
	}
}

//查询应用
func (c *IoTEdgeClient) ShowEdgeApp(request *model.ShowEdgeAppRequest) (*model.ShowEdgeAppResponse, error) {
	requestDef := GenReqDefForShowEdgeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeAppResponse), nil
	}
}

//查询应用版本列表
func (c *IoTEdgeClient) BatchListEdgeAppVersions(request *model.BatchListEdgeAppVersionsRequest) (*model.BatchListEdgeAppVersionsResponse, error) {
	requestDef := GenReqDefForBatchListEdgeAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListEdgeAppVersionsResponse), nil
	}
}

//创建应用版本
func (c *IoTEdgeClient) CreateEdgeApplicationVersion(request *model.CreateEdgeApplicationVersionRequest) (*model.CreateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForCreateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeApplicationVersionResponse), nil
	}
}

//删除应用版本
func (c *IoTEdgeClient) DeleteEdgeApplicationVersion(request *model.DeleteEdgeApplicationVersionRequest) (*model.DeleteEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForDeleteEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeApplicationVersionResponse), nil
	}
}

//查询应用版本详情
func (c *IoTEdgeClient) ShowEdgeApplicationVersion(request *model.ShowEdgeApplicationVersionRequest) (*model.ShowEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForShowEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeApplicationVersionResponse), nil
	}
}

//修改应用版本
func (c *IoTEdgeClient) UpdateEdgeApplicationVersion(request *model.UpdateEdgeApplicationVersionRequest) (*model.UpdateEdgeApplicationVersionResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionResponse), nil
	}
}

//更新应用版本状态。
func (c *IoTEdgeClient) UpdateEdgeApplicationVersionState(request *model.UpdateEdgeApplicationVersionStateRequest) (*model.UpdateEdgeApplicationVersionStateResponse, error) {
	requestDef := GenReqDefForUpdateEdgeApplicationVersionState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeApplicationVersionStateResponse), nil
	}
}

//用户通过在指定边缘节点上设置外部实体的接入信息
func (c *IoTEdgeClient) CreateExternalEntity(request *model.CreateExternalEntityRequest) (*model.CreateExternalEntityResponse, error) {
	requestDef := GenReqDefForCreateExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExternalEntityResponse), nil
	}
}

//删除节点下外部实体
func (c *IoTEdgeClient) DeleteExternalEntity(request *model.DeleteExternalEntityRequest) (*model.DeleteExternalEntityResponse, error) {
	requestDef := GenReqDefForDeleteExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExternalEntityResponse), nil
	}
}

//用户在指定边缘节点上查询外部实体列表
func (c *IoTEdgeClient) ListExternalEntity(request *model.ListExternalEntityRequest) (*model.ListExternalEntityResponse, error) {
	requestDef := GenReqDefForListExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExternalEntityResponse), nil
	}
}

//查询指定节点下指定外部实体的详情
func (c *IoTEdgeClient) ShowExternalEntity(request *model.ShowExternalEntityRequest) (*model.ShowExternalEntityResponse, error) {
	requestDef := GenReqDefForShowExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExternalEntityResponse), nil
	}
}

//用户通过在指定边缘节点上修改指定外部实体的接入信息
func (c *IoTEdgeClient) UpdateExternalEntity(request *model.UpdateExternalEntityRequest) (*model.UpdateExternalEntityResponse, error) {
	requestDef := GenReqDefForUpdateExternalEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExternalEntityResponse), nil
	}
}

//用户通过Console接口查询指定边缘节点上边缘模块列表
func (c *IoTEdgeClient) BatchListModules(request *model.BatchListModulesRequest) (*model.BatchListModulesResponse, error) {
	requestDef := GenReqDefForBatchListModules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListModulesResponse), nil
	}
}

//用户通过Console接口在指定边缘节点上创建边缘模块
func (c *IoTEdgeClient) CreateModule(request *model.CreateModuleRequest) (*model.CreateModuleResponse, error) {
	requestDef := GenReqDefForCreateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModuleResponse), nil
	}
}

//用户通过过Console接口在指定边缘节点上删除边缘模块
func (c *IoTEdgeClient) DeleteModule(request *model.DeleteModuleRequest) (*model.DeleteModuleResponse, error) {
	requestDef := GenReqDefForDeleteModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModuleResponse), nil
	}
}

//用户通过Console接口查询指定边缘节点上指定边缘模块
func (c *IoTEdgeClient) ShowModule(request *model.ShowModuleRequest) (*model.ShowModuleResponse, error) {
	requestDef := GenReqDefForShowModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModuleResponse), nil
	}
}

//用户通过Console接口查询指定边缘节点上指定边缘模块
func (c *IoTEdgeClient) UpdateModule(request *model.UpdateModuleRequest) (*model.UpdateModuleResponse, error) {
	requestDef := GenReqDefForUpdateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModuleResponse), nil
	}
}

//用户在指定边缘节点上查询边缘路由列表
func (c *IoTEdgeClient) ListRoutes(request *model.ListRoutesRequest) (*model.ListRoutesResponse, error) {
	requestDef := GenReqDefForListRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoutesResponse), nil
	}
}

//用户通过在指定边缘节点上设置边缘路由
func (c *IoTEdgeClient) UpdateRoutes(request *model.UpdateRoutesRequest) (*model.UpdateRoutesResponse, error) {
	requestDef := GenReqDefForUpdateRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoutesResponse), nil
	}
}

//南向3rdIA对下发的配置项进行批量确认
func (c *IoTEdgeClient) BatchConfirmConfigsNew(request *model.BatchConfirmConfigsNewRequest) (*model.BatchConfirmConfigsNewResponse, error) {
	requestDef := GenReqDefForBatchConfirmConfigsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchConfirmConfigsNewResponse), nil
	}
}

//批量导入南向3rdIA配置项
func (c *IoTEdgeClient) BatchImportConfigs(request *model.BatchImportConfigsRequest) (*model.BatchImportConfigsResponse, error) {
	requestDef := GenReqDefForBatchImportConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportConfigsResponse), nil
	}
}

//删除南向3rdIA配置项
func (c *IoTEdgeClient) DeleteIaConfig(request *model.DeleteIaConfigRequest) (*model.DeleteIaConfigResponse, error) {
	requestDef := GenReqDefForDeleteIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIaConfigResponse), nil
	}
}

//查询南向3rdIA配置项列表
func (c *IoTEdgeClient) ListIaConfigs(request *model.ListIaConfigsRequest) (*model.ListIaConfigsResponse, error) {
	requestDef := GenReqDefForListIaConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIaConfigsResponse), nil
	}
}

//查询南向3rdIA配置项详情
func (c *IoTEdgeClient) ShowIaConfig(request *model.ShowIaConfigRequest) (*model.ShowIaConfigResponse, error) {
	requestDef := GenReqDefForShowIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIaConfigResponse), nil
	}
}

//创建&更新南向3rdIA配置项信息
func (c *IoTEdgeClient) UpdateIaConfig(request *model.UpdateIaConfigRequest) (*model.UpdateIaConfigResponse, error) {
	requestDef := GenReqDefForUpdateIaConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIaConfigResponse), nil
	}
}

//批量授权北向NA信息到边缘节点。 已授权的边缘节点上的南向IA应用，可以通过部署在边缘节点上的api网关访问北向NA提供的接口。
func (c *IoTEdgeClient) BatchAssociateNaToNodes(request *model.BatchAssociateNaToNodesRequest) (*model.BatchAssociateNaToNodesResponse, error) {
	requestDef := GenReqDefForBatchAssociateNaToNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateNaToNodesResponse), nil
	}
}

//删除北向NA信息，如果有边缘节点已分配该NA信息，会通知到该边缘节点。
func (c *IoTEdgeClient) DeleteNa(request *model.DeleteNaRequest) (*model.DeleteNaResponse, error) {
	requestDef := GenReqDefForDeleteNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNaResponse), nil
	}
}

//查询该北向NA信息的已分配节点
func (c *IoTEdgeClient) ListNaAuthorizedNodes(request *model.ListNaAuthorizedNodesRequest) (*model.ListNaAuthorizedNodesResponse, error) {
	requestDef := GenReqDefForListNaAuthorizedNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNaAuthorizedNodesResponse), nil
	}
}

//查询北向NA信息列表
func (c *IoTEdgeClient) ListNas(request *model.ListNasRequest) (*model.ListNasResponse, error) {
	requestDef := GenReqDefForListNas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNasResponse), nil
	}
}

//查询北向NA信息详情
func (c *IoTEdgeClient) ShowNa(request *model.ShowNaRequest) (*model.ShowNaResponse, error) {
	requestDef := GenReqDefForShowNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNaResponse), nil
	}
}

//创建&更新北向NA信息，当更新北向NA信息时，会通知到已分配该北向NA的所有边缘节点。
func (c *IoTEdgeClient) UpdateNa(request *model.UpdateNaRequest) (*model.UpdateNaResponse, error) {
	requestDef := GenReqDefForUpdateNa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNaResponse), nil
	}
}
