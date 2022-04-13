package v1

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/ief/v1/model"
)

type IefClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIefClient(hcClient *http_client.HcHttpClient) *IefClient {
	return &IefClient{HcClient: hcClient}
}

func IefClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//为指定实例批量添加或删除标签。 一个资源上最多有20个标签。  说明： - 此接口为幂等接口，创建时如果请求体中存在重复key则报错。 - 创建时不允许设置重复key数据,如果数据库已存在该key，就覆盖value的值。 - 删除时不对标签字符集范围做校验，如果删除的标签不存在，默认处理成功。删除时tags结构体不能缺失，key不能为空，或者空字符串。
func (c *IefClient) BatchAddDeleteTags(request *model.BatchAddDeleteTagsRequest) (*model.BatchAddDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchAddDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDeleteTagsResponse), nil
	}
}

//该API用于创建一个应用模板。
func (c *IefClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

//创建一个应用模板版本
func (c *IefClient) CreateAppVersions(request *model.CreateAppVersionsRequest) (*model.CreateAppVersionsResponse, error) {
	requestDef := GenReqDefForCreateAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppVersionsResponse), nil
	}
}

//创建配置项
func (c *IefClient) CreateConfigMap(request *model.CreateConfigMapRequest) (*model.CreateConfigMapResponse, error) {
	requestDef := GenReqDefForCreateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigMapResponse), nil
	}
}

//创建部署
func (c *IefClient) CreateDeployments(request *model.CreateDeploymentsRequest) (*model.CreateDeploymentsResponse, error) {
	requestDef := GenReqDefForCreateDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentsResponse), nil
	}
}

//该API用于注册一个终端设备。
func (c *IefClient) CreateDevice(request *model.CreateDeviceRequest) (*model.CreateDeviceResponse, error) {
	requestDef := GenReqDefForCreateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceResponse), nil
	}
}

//创建一个终端设备模板
func (c *IefClient) CreateDeviceTemplate(request *model.CreateDeviceTemplateRequest) (*model.CreateDeviceTemplateResponse, error) {
	requestDef := GenReqDefForCreateDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceTemplateResponse), nil
	}
}

//该API用于注册一个边缘节点。接口调用成功后，您可以将响应消息体中node.package字段使用base64解码成tar.gz文件，并在控制台下载边缘核心软件，然后纳管边缘节点。
func (c *IefClient) CreateEdgeNode(request *model.CreateEdgeNodeRequest) (*model.CreateEdgeNodeResponse, error) {
	requestDef := GenReqDefForCreateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeResponse), nil
	}
}

//创建边缘节点上的应用证书和设备证书。
func (c *IefClient) CreateEdgeNodeCerts(request *model.CreateEdgeNodeCertsRequest) (*model.CreateEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForCreateEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeCertsResponse), nil
	}
}

//创建一个端点
func (c *IefClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

//创建一条规则
func (c *IefClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

//创建密钥
func (c *IefClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

//创建一个服务
func (c *IefClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

//为资源添加标签。 一个资源上最多有20个标签。 此接口为幂等接口，创建时，如果创建的标签已经存在（key相同），则覆盖。
func (c *IefClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

//删除应用模板
func (c *IefClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

//删除应用版本
func (c *IefClient) DeleteAppVersion(request *model.DeleteAppVersionRequest) (*model.DeleteAppVersionResponse, error) {
	requestDef := GenReqDefForDeleteAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppVersionResponse), nil
	}
}

//删除配置项
func (c *IefClient) DeleteConfigMap(request *model.DeleteConfigMapRequest) (*model.DeleteConfigMapResponse, error) {
	requestDef := GenReqDefForDeleteConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigMapResponse), nil
	}
}

//删除应用部署
func (c *IefClient) DeleteDeployment(request *model.DeleteDeploymentRequest) (*model.DeleteDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentResponse), nil
	}
}

//该API用于删除终端设备。
func (c *IefClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

//删除终端设备模板
func (c *IefClient) DeleteDeviceTemplate(request *model.DeleteDeviceTemplateRequest) (*model.DeleteDeviceTemplateResponse, error) {
	requestDef := GenReqDefForDeleteDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceTemplateResponse), nil
	}
}

//删除边缘节点
func (c *IefClient) DeleteEdgeNode(request *model.DeleteEdgeNodeRequest) (*model.DeleteEdgeNodeResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeResponse), nil
	}
}

//删除边缘节点上的证书（目前只支持删除应用证书和设备证书）
func (c *IefClient) DeleteEdgeNodeCerts(request *model.DeleteEdgeNodeCertsRequest) (*model.DeleteEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeCertsResponse), nil
	}
}

//删除一个端点
func (c *IefClient) DeleteEndPoint(request *model.DeleteEndPointRequest) (*model.DeleteEndPointResponse, error) {
	requestDef := GenReqDefForDeleteEndPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndPointResponse), nil
	}
}

//删除资源标签。删除时不对标签字符集做校验，调用前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，Key不能为空或者空字符串。
func (c *IefClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

//删除一条规则
func (c *IefClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

//删除密钥
func (c *IefClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

//删除一个服务
func (c *IefClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

//该API用于启用停用边缘节点。被停用的边缘节点将无法连接到云端服务，可用该URI启用边缘节点恢复连接。
func (c *IefClient) EnableDisableEdgeNodes(request *model.EnableDisableEdgeNodesRequest) (*model.EnableDisableEdgeNodesResponse, error) {
	requestDef := GenReqDefForEnableDisableEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDisableEdgeNodesResponse), nil
	}
}

//查询应用模板版本列表
func (c *IefClient) ListAppVersions(request *model.ListAppVersionsRequest) (*model.ListAppVersionsResponse, error) {
	requestDef := GenReqDefForListAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppVersionsResponse), nil
	}
}

//查询应用模板列表
func (c *IefClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

//查询配置项列表
func (c *IefClient) ListConfigMaps(request *model.ListConfigMapsRequest) (*model.ListConfigMapsResponse, error) {
	requestDef := GenReqDefForListConfigMaps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigMapsResponse), nil
	}
}

//查询部署列表
func (c *IefClient) ListDeployments(request *model.ListDeploymentsRequest) (*model.ListDeploymentsResponse, error) {
	requestDef := GenReqDefForListDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeploymentsResponse), nil
	}
}

//查询终端设备模板列表
func (c *IefClient) ListDeviceTemplates(request *model.ListDeviceTemplatesRequest) (*model.ListDeviceTemplatesResponse, error) {
	requestDef := GenReqDefForListDeviceTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceTemplatesResponse), nil
	}
}

//该API用于查询终端设备列表。
func (c *IefClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

//查询边缘节点上的应用证书和设备证书。
func (c *IefClient) ListEdgeNodeCerts(request *model.ListEdgeNodeCertsRequest) (*model.ListEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForListEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodeCertsResponse), nil
	}
}

//该API用于查询边缘节点。 - 如果不携带任何检索参数，将返回该租户的所有边缘节点信息。 - app_name和tags不支持复合查询，如果同时存在则返回tags查询结果，可以同时携带多个其他检索参数，可同时生效。
func (c *IefClient) ListEdgeNodes(request *model.ListEdgeNodesRequest) (*model.ListEdgeNodesResponse, error) {
	requestDef := GenReqDefForListEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodesResponse), nil
	}
}

//获取所有的端点详情。 如果不携带任何检索参数，将返回该租户的所有端点信息和系统中所有的共享端点。 如果同时指定is_shared=true和其他参数，同样还会对name、type进行过滤。
func (c *IefClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

//查询应用实例列表
func (c *IefClient) ListPods(request *model.ListPodsRequest) (*model.ListPodsResponse, error) {
	requestDef := GenReqDefForListPods()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPodsResponse), nil
	}
}

//使用标签过滤实例
func (c *IefClient) ListResourceByTags(request *model.ListResourceByTagsRequest) (*model.ListResourceByTagsResponse, error) {
	requestDef := GenReqDefForListResourceByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceByTagsResponse), nil
	}
}

//查询特定规则下的所有错误列表
func (c *IefClient) ListRuleErrors(request *model.ListRuleErrorsRequest) (*model.ListRuleErrorsResponse, error) {
	requestDef := GenReqDefForListRuleErrors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleErrorsResponse), nil
	}
}

//查询到所有的规则
func (c *IefClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

//查询密钥列表
func (c *IefClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

//获取所有的服务详情
func (c *IefClient) ListServices(request *model.ListServicesRequest) (*model.ListServicesResponse, error) {
	requestDef := GenReqDefForListServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesResponse), nil
	}
}

//查询指定实例的标签信息
func (c *IefClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

//查询指定项目中实例类型的所有资源标签集合
func (c *IefClient) ListTagsByResourceType(request *model.ListTagsByResourceTypeRequest) (*model.ListTagsByResourceTypeResponse, error) {
	requestDef := GenReqDefForListTagsByResourceType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsByResourceTypeResponse), nil
	}
}

//查询应用模板详情。
func (c *IefClient) ShowAppDetail(request *model.ShowAppDetailRequest) (*model.ShowAppDetailResponse, error) {
	requestDef := GenReqDefForShowAppDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppDetailResponse), nil
	}
}

//查询应用模板版本详情
func (c *IefClient) ShowAppVersionDetail(request *model.ShowAppVersionDetailRequest) (*model.ShowAppVersionDetailResponse, error) {
	requestDef := GenReqDefForShowAppVersionDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppVersionDetailResponse), nil
	}
}

//查询一个配置项详情
func (c *IefClient) ShowConfigMap(request *model.ShowConfigMapRequest) (*model.ShowConfigMapResponse, error) {
	requestDef := GenReqDefForShowConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigMapResponse), nil
	}
}

//查询应用部署
func (c *IefClient) ShowDeployment(request *model.ShowDeploymentRequest) (*model.ShowDeploymentResponse, error) {
	requestDef := GenReqDefForShowDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentResponse), nil
	}
}

//该API用于查询终端设备详情。
func (c *IefClient) ShowDevice(request *model.ShowDeviceRequest) (*model.ShowDeviceResponse, error) {
	requestDef := GenReqDefForShowDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceResponse), nil
	}
}

//查询一个终端设备模板
func (c *IefClient) ShowDeviceTemplate(request *model.ShowDeviceTemplateRequest) (*model.ShowDeviceTemplateResponse, error) {
	requestDef := GenReqDefForShowDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTemplateResponse), nil
	}
}

//该API用于查询终端设备孪生。
func (c *IefClient) ShowDeviceTwin(request *model.ShowDeviceTwinRequest) (*model.ShowDeviceTwinResponse, error) {
	requestDef := GenReqDefForShowDeviceTwin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTwinResponse), nil
	}
}

//该API用于查询边缘节点详情。
func (c *IefClient) ShowEdgeNodeDetail(request *model.ShowEdgeNodeDetailRequest) (*model.ShowEdgeNodeDetailResponse, error) {
	requestDef := GenReqDefForShowEdgeNodeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeDetailResponse), nil
	}
}

//查询一个端点的详情
func (c *IefClient) ShowEndPointDetail(request *model.ShowEndPointDetailRequest) (*model.ShowEndPointDetailResponse, error) {
	requestDef := GenReqDefForShowEndPointDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndPointDetailResponse), nil
	}
}

//获取一条规则的详情
func (c *IefClient) ShowRuleDetail(request *model.ShowRuleDetailRequest) (*model.ShowRuleDetailResponse, error) {
	requestDef := GenReqDefForShowRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleDetailResponse), nil
	}
}

//查询一个密钥详情
func (c *IefClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

//查询一个服务的详情
func (c *IefClient) ShowServiceDetail(request *model.ShowServiceDetailRequest) (*model.ShowServiceDetailResponse, error) {
	requestDef := GenReqDefForShowServiceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceDetailResponse), nil
	}
}

//启用一条规则
func (c *IefClient) StartRule(request *model.StartRuleRequest) (*model.StartRuleResponse, error) {
	requestDef := GenReqDefForStartRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartRuleResponse), nil
	}
}

//停用一条规则
func (c *IefClient) StopRule(request *model.StopRuleRequest) (*model.StopRuleResponse, error) {
	requestDef := GenReqDefForStopRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopRuleResponse), nil
	}
}

//更新一个应用模板。
func (c *IefClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

//更新一个应用模板版本
func (c *IefClient) UpdateAppVersion(request *model.UpdateAppVersionRequest) (*model.UpdateAppVersionResponse, error) {
	requestDef := GenReqDefForUpdateAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppVersionResponse), nil
	}
}

//更新一个配置项
func (c *IefClient) UpdateConfigMap(request *model.UpdateConfigMapRequest) (*model.UpdateConfigMapResponse, error) {
	requestDef := GenReqDefForUpdateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigMapResponse), nil
	}
}

//修改应用部署
func (c *IefClient) UpdateDeployment(request *model.UpdateDeploymentRequest) (*model.UpdateDeploymentResponse, error) {
	requestDef := GenReqDefForUpdateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentResponse), nil
	}
}

//更新一个终端设备。
func (c *IefClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

//更新一个终端设备模板。
func (c *IefClient) UpdateDeviceTemplateById(request *model.UpdateDeviceTemplateByIdRequest) (*model.UpdateDeviceTemplateByIdResponse, error) {
	requestDef := GenReqDefForUpdateDeviceTemplateById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceTemplateByIdResponse), nil
	}
}

//该API用于更新终端设备孪生。
func (c *IefClient) UpdateDeviceTwin(request *model.UpdateDeviceTwinRequest) (*model.UpdateDeviceTwinResponse, error) {
	requestDef := GenReqDefForUpdateDeviceTwin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceTwinResponse), nil
	}
}

//该API用于更新边缘节点。
func (c *IefClient) UpdateEdgeNode(request *model.UpdateEdgeNodeRequest) (*model.UpdateEdgeNodeResponse, error) {
	requestDef := GenReqDefForUpdateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeNodeResponse), nil
	}
}

//添加或删除边缘节点的终端设备
func (c *IefClient) UpdateEdgeNodeDevice(request *model.UpdateEdgeNodeDeviceRequest) (*model.UpdateEdgeNodeDeviceResponse, error) {
	requestDef := GenReqDefForUpdateEdgeNodeDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeNodeDeviceResponse), nil
	}
}

//该API用于更新终端设备的边缘节点。功能与更新边缘节点的终端设备相同，推荐使用更新边缘节点的终端设备。
func (c *IefClient) UpdateNodeByDeviceId(request *model.UpdateNodeByDeviceIdRequest) (*model.UpdateNodeByDeviceIdResponse, error) {
	requestDef := GenReqDefForUpdateNodeByDeviceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeByDeviceIdResponse), nil
	}
}

//更新一个密钥
func (c *IefClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

//更新一个服务
func (c *IefClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}
