package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ief/v1/model"
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

// BatchAddDeleteTags 批量添加删除资源标签
//
// 为指定实例批量添加或删除标签。
// 一个资源上最多有20个标签。
//
// 说明：
// - 此接口为幂等接口，创建时如果请求体中存在重复key则报错。
// - 创建时不允许设置重复key数据,如果数据库已存在该key，就覆盖value的值。
// - 删除时不对标签字符集范围做校验，如果删除的标签不存在，默认处理成功。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) BatchAddDeleteTags(request *model.BatchAddDeleteTagsRequest) (*model.BatchAddDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchAddDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDeleteTagsResponse), nil
	}
}

// BatchAddDeleteTagsInvoker 批量添加删除资源标签
func (c *IefClient) BatchAddDeleteTagsInvoker(request *model.BatchAddDeleteTagsRequest) *BatchAddDeleteTagsInvoker {
	requestDef := GenReqDefForBatchAddDeleteTags()
	return &BatchAddDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建应用模板
//
// 该API用于创建一个应用模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用模板
func (c *IefClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppVersions 创建应用模板版本
//
// 创建一个应用模板版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateAppVersions(request *model.CreateAppVersionsRequest) (*model.CreateAppVersionsResponse, error) {
	requestDef := GenReqDefForCreateAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppVersionsResponse), nil
	}
}

// CreateAppVersionsInvoker 创建应用模板版本
func (c *IefClient) CreateAppVersionsInvoker(request *model.CreateAppVersionsRequest) *CreateAppVersionsInvoker {
	requestDef := GenReqDefForCreateAppVersions()
	return &CreateAppVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigMap 创建配置项
//
// 创建配置项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateConfigMap(request *model.CreateConfigMapRequest) (*model.CreateConfigMapResponse, error) {
	requestDef := GenReqDefForCreateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigMapResponse), nil
	}
}

// CreateConfigMapInvoker 创建配置项
func (c *IefClient) CreateConfigMapInvoker(request *model.CreateConfigMapRequest) *CreateConfigMapInvoker {
	requestDef := GenReqDefForCreateConfigMap()
	return &CreateConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeployments 创建部署
//
// 创建部署
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateDeployments(request *model.CreateDeploymentsRequest) (*model.CreateDeploymentsResponse, error) {
	requestDef := GenReqDefForCreateDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentsResponse), nil
	}
}

// CreateDeploymentsInvoker 创建部署
func (c *IefClient) CreateDeploymentsInvoker(request *model.CreateDeploymentsRequest) *CreateDeploymentsInvoker {
	requestDef := GenReqDefForCreateDeployments()
	return &CreateDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDevice 注册终端设备
//
// 该API用于注册一个终端设备。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateDevice(request *model.CreateDeviceRequest) (*model.CreateDeviceResponse, error) {
	requestDef := GenReqDefForCreateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceResponse), nil
	}
}

// CreateDeviceInvoker 注册终端设备
func (c *IefClient) CreateDeviceInvoker(request *model.CreateDeviceRequest) *CreateDeviceInvoker {
	requestDef := GenReqDefForCreateDevice()
	return &CreateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeviceTemplate 创建终端设备模板
//
// 创建一个终端设备模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateDeviceTemplate(request *model.CreateDeviceTemplateRequest) (*model.CreateDeviceTemplateResponse, error) {
	requestDef := GenReqDefForCreateDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceTemplateResponse), nil
	}
}

// CreateDeviceTemplateInvoker 创建终端设备模板
func (c *IefClient) CreateDeviceTemplateInvoker(request *model.CreateDeviceTemplateRequest) *CreateDeviceTemplateInvoker {
	requestDef := GenReqDefForCreateDeviceTemplate()
	return &CreateDeviceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeNode 注册边缘节点
//
// 该API用于注册一个边缘节点。接口调用成功后，您可以将响应消息体中node.package字段使用base64解码成tar.gz文件，并在控制台下载边缘核心软件，然后纳管边缘节点。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateEdgeNode(request *model.CreateEdgeNodeRequest) (*model.CreateEdgeNodeResponse, error) {
	requestDef := GenReqDefForCreateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeResponse), nil
	}
}

// CreateEdgeNodeInvoker 注册边缘节点
func (c *IefClient) CreateEdgeNodeInvoker(request *model.CreateEdgeNodeRequest) *CreateEdgeNodeInvoker {
	requestDef := GenReqDefForCreateEdgeNode()
	return &CreateEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeNodeCerts 创建节点证书
//
// 创建边缘节点上的应用证书和设备证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateEdgeNodeCerts(request *model.CreateEdgeNodeCertsRequest) (*model.CreateEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForCreateEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeNodeCertsResponse), nil
	}
}

// CreateEdgeNodeCertsInvoker 创建节点证书
func (c *IefClient) CreateEdgeNodeCertsInvoker(request *model.CreateEdgeNodeCertsRequest) *CreateEdgeNodeCertsInvoker {
	requestDef := GenReqDefForCreateEdgeNodeCerts()
	return &CreateEdgeNodeCertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEncryptdatas 新增加密数据
//
// 新增加密数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateEncryptdatas(request *model.CreateEncryptdatasRequest) (*model.CreateEncryptdatasResponse, error) {
	requestDef := GenReqDefForCreateEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEncryptdatasResponse), nil
	}
}

// CreateEncryptdatasInvoker 新增加密数据
func (c *IefClient) CreateEncryptdatasInvoker(request *model.CreateEncryptdatasRequest) *CreateEncryptdatasInvoker {
	requestDef := GenReqDefForCreateEncryptdatas()
	return &CreateEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建端点
//
// 创建一个端点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建端点
func (c *IefClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNodeEncryptdatas 加密数据绑定边缘节点
//
// 加密数据绑定边缘节点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateNodeEncryptdatas(request *model.CreateNodeEncryptdatasRequest) (*model.CreateNodeEncryptdatasResponse, error) {
	requestDef := GenReqDefForCreateNodeEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodeEncryptdatasResponse), nil
	}
}

// CreateNodeEncryptdatasInvoker 加密数据绑定边缘节点
func (c *IefClient) CreateNodeEncryptdatasInvoker(request *model.CreateNodeEncryptdatasRequest) *CreateNodeEncryptdatasInvoker {
	requestDef := GenReqDefForCreateNodeEncryptdatas()
	return &CreateNodeEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRule 创建规则
//
// 创建一条规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

// CreateRuleInvoker 创建规则
func (c *IefClient) CreateRuleInvoker(request *model.CreateRuleRequest) *CreateRuleInvoker {
	requestDef := GenReqDefForCreateRule()
	return &CreateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecret 创建密钥
//
// 创建密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// CreateSecretInvoker 创建密钥
func (c *IefClient) CreateSecretInvoker(request *model.CreateSecretRequest) *CreateSecretInvoker {
	requestDef := GenReqDefForCreateSecret()
	return &CreateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateService 创建服务
//
// 创建一个服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

// CreateServiceInvoker 创建服务
func (c *IefClient) CreateServiceInvoker(request *model.CreateServiceRequest) *CreateServiceInvoker {
	requestDef := GenReqDefForCreateService()
	return &CreateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 添加资源标签
//
// 为资源添加标签。
// 一个资源上最多有20个标签。
// 此接口为幂等接口，创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 添加资源标签
func (c *IefClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用模板
//
// 删除应用模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用模板
func (c *IefClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppVersion 删除应用版本
//
// 删除应用版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteAppVersion(request *model.DeleteAppVersionRequest) (*model.DeleteAppVersionResponse, error) {
	requestDef := GenReqDefForDeleteAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppVersionResponse), nil
	}
}

// DeleteAppVersionInvoker 删除应用版本
func (c *IefClient) DeleteAppVersionInvoker(request *model.DeleteAppVersionRequest) *DeleteAppVersionInvoker {
	requestDef := GenReqDefForDeleteAppVersion()
	return &DeleteAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigMap 删除配置项
//
// 删除配置项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteConfigMap(request *model.DeleteConfigMapRequest) (*model.DeleteConfigMapResponse, error) {
	requestDef := GenReqDefForDeleteConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigMapResponse), nil
	}
}

// DeleteConfigMapInvoker 删除配置项
func (c *IefClient) DeleteConfigMapInvoker(request *model.DeleteConfigMapRequest) *DeleteConfigMapInvoker {
	requestDef := GenReqDefForDeleteConfigMap()
	return &DeleteConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployment 删除部署
//
// 删除应用部署
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteDeployment(request *model.DeleteDeploymentRequest) (*model.DeleteDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentResponse), nil
	}
}

// DeleteDeploymentInvoker 删除部署
func (c *IefClient) DeleteDeploymentInvoker(request *model.DeleteDeploymentRequest) *DeleteDeploymentInvoker {
	requestDef := GenReqDefForDeleteDeployment()
	return &DeleteDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDevice 删除终端设备
//
// 该API用于删除终端设备。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

// DeleteDeviceInvoker 删除终端设备
func (c *IefClient) DeleteDeviceInvoker(request *model.DeleteDeviceRequest) *DeleteDeviceInvoker {
	requestDef := GenReqDefForDeleteDevice()
	return &DeleteDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeviceTemplate 删除终端设备模板
//
// 删除终端设备模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteDeviceTemplate(request *model.DeleteDeviceTemplateRequest) (*model.DeleteDeviceTemplateResponse, error) {
	requestDef := GenReqDefForDeleteDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceTemplateResponse), nil
	}
}

// DeleteDeviceTemplateInvoker 删除终端设备模板
func (c *IefClient) DeleteDeviceTemplateInvoker(request *model.DeleteDeviceTemplateRequest) *DeleteDeviceTemplateInvoker {
	requestDef := GenReqDefForDeleteDeviceTemplate()
	return &DeleteDeviceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeNode 删除边缘节点
//
// 删除边缘节点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteEdgeNode(request *model.DeleteEdgeNodeRequest) (*model.DeleteEdgeNodeResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeResponse), nil
	}
}

// DeleteEdgeNodeInvoker 删除边缘节点
func (c *IefClient) DeleteEdgeNodeInvoker(request *model.DeleteEdgeNodeRequest) *DeleteEdgeNodeInvoker {
	requestDef := GenReqDefForDeleteEdgeNode()
	return &DeleteEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeNodeCerts 删除节点证书
//
// 删除边缘节点上的证书（目前只支持删除应用证书和设备证书）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteEdgeNodeCerts(request *model.DeleteEdgeNodeCertsRequest) (*model.DeleteEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForDeleteEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeNodeCertsResponse), nil
	}
}

// DeleteEdgeNodeCertsInvoker 删除节点证书
func (c *IefClient) DeleteEdgeNodeCertsInvoker(request *model.DeleteEdgeNodeCertsRequest) *DeleteEdgeNodeCertsInvoker {
	requestDef := GenReqDefForDeleteEdgeNodeCerts()
	return &DeleteEdgeNodeCertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEncryptdatas 删除加密数据
//
// 删除加密数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteEncryptdatas(request *model.DeleteEncryptdatasRequest) (*model.DeleteEncryptdatasResponse, error) {
	requestDef := GenReqDefForDeleteEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEncryptdatasResponse), nil
	}
}

// DeleteEncryptdatasInvoker 删除加密数据
func (c *IefClient) DeleteEncryptdatasInvoker(request *model.DeleteEncryptdatasRequest) *DeleteEncryptdatasInvoker {
	requestDef := GenReqDefForDeleteEncryptdatas()
	return &DeleteEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndPoint 删除一个端点
//
// 删除一个端点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteEndPoint(request *model.DeleteEndPointRequest) (*model.DeleteEndPointResponse, error) {
	requestDef := GenReqDefForDeleteEndPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndPointResponse), nil
	}
}

// DeleteEndPointInvoker 删除一个端点
func (c *IefClient) DeleteEndPointInvoker(request *model.DeleteEndPointRequest) *DeleteEndPointInvoker {
	requestDef := GenReqDefForDeleteEndPoint()
	return &DeleteEndPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNodeEncryptdatas 解绑边缘节点的加密数据
//
// 解绑边缘节点的加密数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteNodeEncryptdatas(request *model.DeleteNodeEncryptdatasRequest) (*model.DeleteNodeEncryptdatasResponse, error) {
	requestDef := GenReqDefForDeleteNodeEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodeEncryptdatasResponse), nil
	}
}

// DeleteNodeEncryptdatasInvoker 解绑边缘节点的加密数据
func (c *IefClient) DeleteNodeEncryptdatasInvoker(request *model.DeleteNodeEncryptdatasRequest) *DeleteNodeEncryptdatasInvoker {
	requestDef := GenReqDefForDeleteNodeEncryptdatas()
	return &DeleteNodeEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 删除资源标签
//
// 删除资源标签。删除时不对标签字符集做校验，调用前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，Key不能为空或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 删除资源标签
func (c *IefClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRule 删除规则
//
// 删除一条规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// DeleteRuleInvoker 删除规则
func (c *IefClient) DeleteRuleInvoker(request *model.DeleteRuleRequest) *DeleteRuleInvoker {
	requestDef := GenReqDefForDeleteRule()
	return &DeleteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecret 删除密钥
//
// 删除密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

// DeleteSecretInvoker 删除密钥
func (c *IefClient) DeleteSecretInvoker(request *model.DeleteSecretRequest) *DeleteSecretInvoker {
	requestDef := GenReqDefForDeleteSecret()
	return &DeleteSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteService 删除服务
//
// 删除一个服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

// DeleteServiceInvoker 删除服务
func (c *IefClient) DeleteServiceInvoker(request *model.DeleteServiceRequest) *DeleteServiceInvoker {
	requestDef := GenReqDefForDeleteService()
	return &DeleteServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDisableEdgeNodes 启用停用边缘节点
//
// 该API用于启用停用边缘节点。被停用的边缘节点将无法连接到云端服务，可用该URI启用边缘节点恢复连接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) EnableDisableEdgeNodes(request *model.EnableDisableEdgeNodesRequest) (*model.EnableDisableEdgeNodesResponse, error) {
	requestDef := GenReqDefForEnableDisableEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDisableEdgeNodesResponse), nil
	}
}

// EnableDisableEdgeNodesInvoker 启用停用边缘节点
func (c *IefClient) EnableDisableEdgeNodesInvoker(request *model.EnableDisableEdgeNodesRequest) *EnableDisableEdgeNodesInvoker {
	requestDef := GenReqDefForEnableDisableEdgeNodes()
	return &EnableDisableEdgeNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppVersions 查询应用模板版本列表
//
// 查询应用模板版本列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListAppVersions(request *model.ListAppVersionsRequest) (*model.ListAppVersionsResponse, error) {
	requestDef := GenReqDefForListAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppVersionsResponse), nil
	}
}

// ListAppVersionsInvoker 查询应用模板版本列表
func (c *IefClient) ListAppVersionsInvoker(request *model.ListAppVersionsRequest) *ListAppVersionsInvoker {
	requestDef := GenReqDefForListAppVersions()
	return &ListAppVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 查询应用模板列表
//
// 查询应用模板列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 查询应用模板列表
func (c *IefClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigMaps 查询配置项列表
//
// 查询配置项列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListConfigMaps(request *model.ListConfigMapsRequest) (*model.ListConfigMapsResponse, error) {
	requestDef := GenReqDefForListConfigMaps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigMapsResponse), nil
	}
}

// ListConfigMapsInvoker 查询配置项列表
func (c *IefClient) ListConfigMapsInvoker(request *model.ListConfigMapsRequest) *ListConfigMapsInvoker {
	requestDef := GenReqDefForListConfigMaps()
	return &ListConfigMapsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployments 查询部署列表
//
// 查询部署列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListDeployments(request *model.ListDeploymentsRequest) (*model.ListDeploymentsResponse, error) {
	requestDef := GenReqDefForListDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeploymentsResponse), nil
	}
}

// ListDeploymentsInvoker 查询部署列表
func (c *IefClient) ListDeploymentsInvoker(request *model.ListDeploymentsRequest) *ListDeploymentsInvoker {
	requestDef := GenReqDefForListDeployments()
	return &ListDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeviceTemplates 查询终端设备模板列表
//
// 查询终端设备模板列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListDeviceTemplates(request *model.ListDeviceTemplatesRequest) (*model.ListDeviceTemplatesResponse, error) {
	requestDef := GenReqDefForListDeviceTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceTemplatesResponse), nil
	}
}

// ListDeviceTemplatesInvoker 查询终端设备模板列表
func (c *IefClient) ListDeviceTemplatesInvoker(request *model.ListDeviceTemplatesRequest) *ListDeviceTemplatesInvoker {
	requestDef := GenReqDefForListDeviceTemplates()
	return &ListDeviceTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevices 查询终端设备列表
//
// 该API用于查询终端设备列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

// ListDevicesInvoker 查询终端设备列表
func (c *IefClient) ListDevicesInvoker(request *model.ListDevicesRequest) *ListDevicesInvoker {
	requestDef := GenReqDefForListDevices()
	return &ListDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeNodeCerts 查询节点证书
//
// 查询边缘节点上的应用证书和设备证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListEdgeNodeCerts(request *model.ListEdgeNodeCertsRequest) (*model.ListEdgeNodeCertsResponse, error) {
	requestDef := GenReqDefForListEdgeNodeCerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodeCertsResponse), nil
	}
}

// ListEdgeNodeCertsInvoker 查询节点证书
func (c *IefClient) ListEdgeNodeCertsInvoker(request *model.ListEdgeNodeCertsRequest) *ListEdgeNodeCertsInvoker {
	requestDef := GenReqDefForListEdgeNodeCerts()
	return &ListEdgeNodeCertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeNodes 查询边缘节点列表
//
// 该API用于查询边缘节点。
// - 如果不携带任何检索参数，将返回该租户的所有边缘节点信息。
// - app_name和tags不支持复合查询，如果同时存在则返回tags查询结果，可以同时携带多个其他检索参数，可同时生效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListEdgeNodes(request *model.ListEdgeNodesRequest) (*model.ListEdgeNodesResponse, error) {
	requestDef := GenReqDefForListEdgeNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeNodesResponse), nil
	}
}

// ListEdgeNodesInvoker 查询边缘节点列表
func (c *IefClient) ListEdgeNodesInvoker(request *model.ListEdgeNodesRequest) *ListEdgeNodesInvoker {
	requestDef := GenReqDefForListEdgeNodes()
	return &ListEdgeNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEncryptdataNodes 获取加密数据绑定的边缘节点
//
// 获取加密数据绑定的边缘节点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListEncryptdataNodes(request *model.ListEncryptdataNodesRequest) (*model.ListEncryptdataNodesResponse, error) {
	requestDef := GenReqDefForListEncryptdataNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEncryptdataNodesResponse), nil
	}
}

// ListEncryptdataNodesInvoker 获取加密数据绑定的边缘节点
func (c *IefClient) ListEncryptdataNodesInvoker(request *model.ListEncryptdataNodesRequest) *ListEncryptdataNodesInvoker {
	requestDef := GenReqDefForListEncryptdataNodes()
	return &ListEncryptdataNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEncryptdatas 获取加密数据列表
//
// 获取加密数据列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListEncryptdatas(request *model.ListEncryptdatasRequest) (*model.ListEncryptdatasResponse, error) {
	requestDef := GenReqDefForListEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEncryptdatasResponse), nil
	}
}

// ListEncryptdatasInvoker 获取加密数据列表
func (c *IefClient) ListEncryptdatasInvoker(request *model.ListEncryptdatasRequest) *ListEncryptdatasInvoker {
	requestDef := GenReqDefForListEncryptdatas()
	return &ListEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询端点列表
//
// 获取所有的端点详情。
// 如果不携带任何检索参数，将返回该租户的所有端点信息和系统中所有的共享端点。
// 如果同时指定is_shared&#x3D;true和其他参数，同样还会对name、type进行过滤。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询端点列表
func (c *IefClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodeEncryptdatas 获取边缘节点绑定的加密数据
//
// 获取边缘节点绑定的加密数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListNodeEncryptdatas(request *model.ListNodeEncryptdatasRequest) (*model.ListNodeEncryptdatasResponse, error) {
	requestDef := GenReqDefForListNodeEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodeEncryptdatasResponse), nil
	}
}

// ListNodeEncryptdatasInvoker 获取边缘节点绑定的加密数据
func (c *IefClient) ListNodeEncryptdatasInvoker(request *model.ListNodeEncryptdatasRequest) *ListNodeEncryptdatasInvoker {
	requestDef := GenReqDefForListNodeEncryptdatas()
	return &ListNodeEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPods 查询应用实例列表
//
// 查询应用实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListPods(request *model.ListPodsRequest) (*model.ListPodsResponse, error) {
	requestDef := GenReqDefForListPods()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPodsResponse), nil
	}
}

// ListPodsInvoker 查询应用实例列表
func (c *IefClient) ListPodsInvoker(request *model.ListPodsRequest) *ListPodsInvoker {
	requestDef := GenReqDefForListPods()
	return &ListPodsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceByTags 查询资源实例
//
// 使用标签过滤实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListResourceByTags(request *model.ListResourceByTagsRequest) (*model.ListResourceByTagsResponse, error) {
	requestDef := GenReqDefForListResourceByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceByTagsResponse), nil
	}
}

// ListResourceByTagsInvoker 查询资源实例
func (c *IefClient) ListResourceByTagsInvoker(request *model.ListResourceByTagsRequest) *ListResourceByTagsInvoker {
	requestDef := GenReqDefForListResourceByTags()
	return &ListResourceByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleErrors 查询规则错误列表
//
// 查询特定规则下的所有错误列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListRuleErrors(request *model.ListRuleErrorsRequest) (*model.ListRuleErrorsResponse, error) {
	requestDef := GenReqDefForListRuleErrors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleErrorsResponse), nil
	}
}

// ListRuleErrorsInvoker 查询规则错误列表
func (c *IefClient) ListRuleErrorsInvoker(request *model.ListRuleErrorsRequest) *ListRuleErrorsInvoker {
	requestDef := GenReqDefForListRuleErrors()
	return &ListRuleErrorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRules 查询规则列表
//
// 查询到所有的规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

// ListRulesInvoker 查询规则列表
func (c *IefClient) ListRulesInvoker(request *model.ListRulesRequest) *ListRulesInvoker {
	requestDef := GenReqDefForListRules()
	return &ListRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecrets 查询密钥列表
//
// 查询密钥列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

// ListSecretsInvoker 查询密钥列表
func (c *IefClient) ListSecretsInvoker(request *model.ListSecretsRequest) *ListSecretsInvoker {
	requestDef := GenReqDefForListSecrets()
	return &ListSecretsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServices 查询服务列表
//
// 获取所有的服务详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListServices(request *model.ListServicesRequest) (*model.ListServicesResponse, error) {
	requestDef := GenReqDefForListServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesResponse), nil
	}
}

// ListServicesInvoker 查询服务列表
func (c *IefClient) ListServicesInvoker(request *model.ListServicesRequest) *ListServicesInvoker {
	requestDef := GenReqDefForListServices()
	return &ListServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询资源标签
//
// 查询指定实例的标签信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询资源标签
func (c *IefClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsByResourceType 查询项目标签
//
// 查询指定项目中实例类型的所有资源标签集合
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ListTagsByResourceType(request *model.ListTagsByResourceTypeRequest) (*model.ListTagsByResourceTypeResponse, error) {
	requestDef := GenReqDefForListTagsByResourceType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsByResourceTypeResponse), nil
	}
}

// ListTagsByResourceTypeInvoker 查询项目标签
func (c *IefClient) ListTagsByResourceTypeInvoker(request *model.ListTagsByResourceTypeRequest) *ListTagsByResourceTypeInvoker {
	requestDef := GenReqDefForListTagsByResourceType()
	return &ListTagsByResourceTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartDeploymentsPod 容器应用实例重启
//
// 重启部署下的应用实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) RestartDeploymentsPod(request *model.RestartDeploymentsPodRequest) (*model.RestartDeploymentsPodResponse, error) {
	requestDef := GenReqDefForRestartDeploymentsPod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartDeploymentsPodResponse), nil
	}
}

// RestartDeploymentsPodInvoker 容器应用实例重启
func (c *IefClient) RestartDeploymentsPodInvoker(request *model.RestartDeploymentsPodRequest) *RestartDeploymentsPodInvoker {
	requestDef := GenReqDefForRestartDeploymentsPod()
	return &RestartDeploymentsPodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppDetail 查询应用模板详情
//
// 查询应用模板详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowAppDetail(request *model.ShowAppDetailRequest) (*model.ShowAppDetailResponse, error) {
	requestDef := GenReqDefForShowAppDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppDetailResponse), nil
	}
}

// ShowAppDetailInvoker 查询应用模板详情
func (c *IefClient) ShowAppDetailInvoker(request *model.ShowAppDetailRequest) *ShowAppDetailInvoker {
	requestDef := GenReqDefForShowAppDetail()
	return &ShowAppDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppVersionDetail 查询应用模板版本详情
//
// 查询应用模板版本详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowAppVersionDetail(request *model.ShowAppVersionDetailRequest) (*model.ShowAppVersionDetailResponse, error) {
	requestDef := GenReqDefForShowAppVersionDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppVersionDetailResponse), nil
	}
}

// ShowAppVersionDetailInvoker 查询应用模板版本详情
func (c *IefClient) ShowAppVersionDetailInvoker(request *model.ShowAppVersionDetailRequest) *ShowAppVersionDetailInvoker {
	requestDef := GenReqDefForShowAppVersionDetail()
	return &ShowAppVersionDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigMap 查询配置项详情
//
// 查询一个配置项详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowConfigMap(request *model.ShowConfigMapRequest) (*model.ShowConfigMapResponse, error) {
	requestDef := GenReqDefForShowConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigMapResponse), nil
	}
}

// ShowConfigMapInvoker 查询配置项详情
func (c *IefClient) ShowConfigMapInvoker(request *model.ShowConfigMapRequest) *ShowConfigMapInvoker {
	requestDef := GenReqDefForShowConfigMap()
	return &ShowConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeployment 查询应用部署
//
// 查询应用部署
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowDeployment(request *model.ShowDeploymentRequest) (*model.ShowDeploymentResponse, error) {
	requestDef := GenReqDefForShowDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentResponse), nil
	}
}

// ShowDeploymentInvoker 查询应用部署
func (c *IefClient) ShowDeploymentInvoker(request *model.ShowDeploymentRequest) *ShowDeploymentInvoker {
	requestDef := GenReqDefForShowDeployment()
	return &ShowDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDevice 查询终端设备详情
//
// 该API用于查询终端设备详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowDevice(request *model.ShowDeviceRequest) (*model.ShowDeviceResponse, error) {
	requestDef := GenReqDefForShowDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceResponse), nil
	}
}

// ShowDeviceInvoker 查询终端设备详情
func (c *IefClient) ShowDeviceInvoker(request *model.ShowDeviceRequest) *ShowDeviceInvoker {
	requestDef := GenReqDefForShowDevice()
	return &ShowDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceTemplate 查询终端设备模板
//
// 查询一个终端设备模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowDeviceTemplate(request *model.ShowDeviceTemplateRequest) (*model.ShowDeviceTemplateResponse, error) {
	requestDef := GenReqDefForShowDeviceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTemplateResponse), nil
	}
}

// ShowDeviceTemplateInvoker 查询终端设备模板
func (c *IefClient) ShowDeviceTemplateInvoker(request *model.ShowDeviceTemplateRequest) *ShowDeviceTemplateInvoker {
	requestDef := GenReqDefForShowDeviceTemplate()
	return &ShowDeviceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceTwin 查询终端设备孪生
//
// 该API用于查询终端设备孪生。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowDeviceTwin(request *model.ShowDeviceTwinRequest) (*model.ShowDeviceTwinResponse, error) {
	requestDef := GenReqDefForShowDeviceTwin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTwinResponse), nil
	}
}

// ShowDeviceTwinInvoker 查询终端设备孪生
func (c *IefClient) ShowDeviceTwinInvoker(request *model.ShowDeviceTwinRequest) *ShowDeviceTwinInvoker {
	requestDef := GenReqDefForShowDeviceTwin()
	return &ShowDeviceTwinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeNodeDetail 查询边缘节点详情
//
// 该API用于查询边缘节点详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowEdgeNodeDetail(request *model.ShowEdgeNodeDetailRequest) (*model.ShowEdgeNodeDetailResponse, error) {
	requestDef := GenReqDefForShowEdgeNodeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeDetailResponse), nil
	}
}

// ShowEdgeNodeDetailInvoker 查询边缘节点详情
func (c *IefClient) ShowEdgeNodeDetailInvoker(request *model.ShowEdgeNodeDetailRequest) *ShowEdgeNodeDetailInvoker {
	requestDef := GenReqDefForShowEdgeNodeDetail()
	return &ShowEdgeNodeDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeNodeUpgradeDetails 查看边缘节点升级状态
//
// 该API用于查看边缘节点升级状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowEdgeNodeUpgradeDetails(request *model.ShowEdgeNodeUpgradeDetailsRequest) (*model.ShowEdgeNodeUpgradeDetailsResponse, error) {
	requestDef := GenReqDefForShowEdgeNodeUpgradeDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeNodeUpgradeDetailsResponse), nil
	}
}

// ShowEdgeNodeUpgradeDetailsInvoker 查看边缘节点升级状态
func (c *IefClient) ShowEdgeNodeUpgradeDetailsInvoker(request *model.ShowEdgeNodeUpgradeDetailsRequest) *ShowEdgeNodeUpgradeDetailsInvoker {
	requestDef := GenReqDefForShowEdgeNodeUpgradeDetails()
	return &ShowEdgeNodeUpgradeDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEncryptdatas 查询加密数据详情
//
// 查询加密数据详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowEncryptdatas(request *model.ShowEncryptdatasRequest) (*model.ShowEncryptdatasResponse, error) {
	requestDef := GenReqDefForShowEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEncryptdatasResponse), nil
	}
}

// ShowEncryptdatasInvoker 查询加密数据详情
func (c *IefClient) ShowEncryptdatasInvoker(request *model.ShowEncryptdatasRequest) *ShowEncryptdatasInvoker {
	requestDef := GenReqDefForShowEncryptdatas()
	return &ShowEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEndPointDetail 查询端点详情
//
// 查询一个端点的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowEndPointDetail(request *model.ShowEndPointDetailRequest) (*model.ShowEndPointDetailResponse, error) {
	requestDef := GenReqDefForShowEndPointDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndPointDetailResponse), nil
	}
}

// ShowEndPointDetailInvoker 查询端点详情
func (c *IefClient) ShowEndPointDetailInvoker(request *model.ShowEndPointDetailRequest) *ShowEndPointDetailInvoker {
	requestDef := GenReqDefForShowEndPointDetail()
	return &ShowEndPointDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRuleDetail 查询规则详情
//
// 获取一条规则的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowRuleDetail(request *model.ShowRuleDetailRequest) (*model.ShowRuleDetailResponse, error) {
	requestDef := GenReqDefForShowRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleDetailResponse), nil
	}
}

// ShowRuleDetailInvoker 查询规则详情
func (c *IefClient) ShowRuleDetailInvoker(request *model.ShowRuleDetailRequest) *ShowRuleDetailInvoker {
	requestDef := GenReqDefForShowRuleDetail()
	return &ShowRuleDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecret 查询密钥详情
//
// 查询一个密钥详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

// ShowSecretInvoker 查询密钥详情
func (c *IefClient) ShowSecretInvoker(request *model.ShowSecretRequest) *ShowSecretInvoker {
	requestDef := GenReqDefForShowSecret()
	return &ShowSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServiceDetail 查询服务详情
//
// 查询一个服务的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) ShowServiceDetail(request *model.ShowServiceDetailRequest) (*model.ShowServiceDetailResponse, error) {
	requestDef := GenReqDefForShowServiceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceDetailResponse), nil
	}
}

// ShowServiceDetailInvoker 查询服务详情
func (c *IefClient) ShowServiceDetailInvoker(request *model.ShowServiceDetailRequest) *ShowServiceDetailInvoker {
	requestDef := GenReqDefForShowServiceDetail()
	return &ShowServiceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartRule 启用规则
//
// 启用一条规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) StartRule(request *model.StartRuleRequest) (*model.StartRuleResponse, error) {
	requestDef := GenReqDefForStartRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartRuleResponse), nil
	}
}

// StartRuleInvoker 启用规则
func (c *IefClient) StartRuleInvoker(request *model.StartRuleRequest) *StartRuleInvoker {
	requestDef := GenReqDefForStartRule()
	return &StartRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopRule 停用规则
//
// 停用一条规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) StopRule(request *model.StopRuleRequest) (*model.StopRuleResponse, error) {
	requestDef := GenReqDefForStopRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopRuleResponse), nil
	}
}

// StopRuleInvoker 停用规则
func (c *IefClient) StopRuleInvoker(request *model.StopRuleRequest) *StopRuleInvoker {
	requestDef := GenReqDefForStopRule()
	return &StopRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 更新应用模板
//
// 更新一个应用模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 更新应用模板
func (c *IefClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppVersion 更新应用模板版本
//
// 更新一个应用模板版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateAppVersion(request *model.UpdateAppVersionRequest) (*model.UpdateAppVersionResponse, error) {
	requestDef := GenReqDefForUpdateAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppVersionResponse), nil
	}
}

// UpdateAppVersionInvoker 更新应用模板版本
func (c *IefClient) UpdateAppVersionInvoker(request *model.UpdateAppVersionRequest) *UpdateAppVersionInvoker {
	requestDef := GenReqDefForUpdateAppVersion()
	return &UpdateAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigMap 更新配置项
//
// 更新一个配置项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateConfigMap(request *model.UpdateConfigMapRequest) (*model.UpdateConfigMapResponse, error) {
	requestDef := GenReqDefForUpdateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigMapResponse), nil
	}
}

// UpdateConfigMapInvoker 更新配置项
func (c *IefClient) UpdateConfigMapInvoker(request *model.UpdateConfigMapRequest) *UpdateConfigMapInvoker {
	requestDef := GenReqDefForUpdateConfigMap()
	return &UpdateConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeployment 更新应用部署
//
// 修改应用部署
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateDeployment(request *model.UpdateDeploymentRequest) (*model.UpdateDeploymentResponse, error) {
	requestDef := GenReqDefForUpdateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentResponse), nil
	}
}

// UpdateDeploymentInvoker 更新应用部署
func (c *IefClient) UpdateDeploymentInvoker(request *model.UpdateDeploymentRequest) *UpdateDeploymentInvoker {
	requestDef := GenReqDefForUpdateDeployment()
	return &UpdateDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevice 更新终端设备
//
// 更新一个终端设备。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// UpdateDeviceInvoker 更新终端设备
func (c *IefClient) UpdateDeviceInvoker(request *model.UpdateDeviceRequest) *UpdateDeviceInvoker {
	requestDef := GenReqDefForUpdateDevice()
	return &UpdateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeviceTemplateById 更新终端设备模板
//
// 更新一个终端设备模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateDeviceTemplateById(request *model.UpdateDeviceTemplateByIdRequest) (*model.UpdateDeviceTemplateByIdResponse, error) {
	requestDef := GenReqDefForUpdateDeviceTemplateById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceTemplateByIdResponse), nil
	}
}

// UpdateDeviceTemplateByIdInvoker 更新终端设备模板
func (c *IefClient) UpdateDeviceTemplateByIdInvoker(request *model.UpdateDeviceTemplateByIdRequest) *UpdateDeviceTemplateByIdInvoker {
	requestDef := GenReqDefForUpdateDeviceTemplateById()
	return &UpdateDeviceTemplateByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeviceTwin 更新终端设备孪生
//
// 该API用于更新终端设备孪生。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateDeviceTwin(request *model.UpdateDeviceTwinRequest) (*model.UpdateDeviceTwinResponse, error) {
	requestDef := GenReqDefForUpdateDeviceTwin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceTwinResponse), nil
	}
}

// UpdateDeviceTwinInvoker 更新终端设备孪生
func (c *IefClient) UpdateDeviceTwinInvoker(request *model.UpdateDeviceTwinRequest) *UpdateDeviceTwinInvoker {
	requestDef := GenReqDefForUpdateDeviceTwin()
	return &UpdateDeviceTwinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeNode 更新边缘节点
//
// 该API用于更新边缘节点。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateEdgeNode(request *model.UpdateEdgeNodeRequest) (*model.UpdateEdgeNodeResponse, error) {
	requestDef := GenReqDefForUpdateEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeNodeResponse), nil
	}
}

// UpdateEdgeNodeInvoker 更新边缘节点
func (c *IefClient) UpdateEdgeNodeInvoker(request *model.UpdateEdgeNodeRequest) *UpdateEdgeNodeInvoker {
	requestDef := GenReqDefForUpdateEdgeNode()
	return &UpdateEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeNodeDevice 更新边缘节点的终端设备
//
// 添加或删除边缘节点的终端设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateEdgeNodeDevice(request *model.UpdateEdgeNodeDeviceRequest) (*model.UpdateEdgeNodeDeviceResponse, error) {
	requestDef := GenReqDefForUpdateEdgeNodeDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeNodeDeviceResponse), nil
	}
}

// UpdateEdgeNodeDeviceInvoker 更新边缘节点的终端设备
func (c *IefClient) UpdateEdgeNodeDeviceInvoker(request *model.UpdateEdgeNodeDeviceRequest) *UpdateEdgeNodeDeviceInvoker {
	requestDef := GenReqDefForUpdateEdgeNodeDevice()
	return &UpdateEdgeNodeDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEncryptdatas 更新加密数据
//
// 更新加密数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateEncryptdatas(request *model.UpdateEncryptdatasRequest) (*model.UpdateEncryptdatasResponse, error) {
	requestDef := GenReqDefForUpdateEncryptdatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEncryptdatasResponse), nil
	}
}

// UpdateEncryptdatasInvoker 更新加密数据
func (c *IefClient) UpdateEncryptdatasInvoker(request *model.UpdateEncryptdatasRequest) *UpdateEncryptdatasInvoker {
	requestDef := GenReqDefForUpdateEncryptdatas()
	return &UpdateEncryptdatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNodeByDeviceId 更新终端设备的边缘节点
//
// 该API用于更新终端设备的边缘节点。功能与更新边缘节点的终端设备相同，推荐使用更新边缘节点的终端设备。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateNodeByDeviceId(request *model.UpdateNodeByDeviceIdRequest) (*model.UpdateNodeByDeviceIdResponse, error) {
	requestDef := GenReqDefForUpdateNodeByDeviceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeByDeviceIdResponse), nil
	}
}

// UpdateNodeByDeviceIdInvoker 更新终端设备的边缘节点
func (c *IefClient) UpdateNodeByDeviceIdInvoker(request *model.UpdateNodeByDeviceIdRequest) *UpdateNodeByDeviceIdInvoker {
	requestDef := GenReqDefForUpdateNodeByDeviceId()
	return &UpdateNodeByDeviceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecret 更新密钥
//
// 更新一个密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

// UpdateSecretInvoker 更新密钥
func (c *IefClient) UpdateSecretInvoker(request *model.UpdateSecretRequest) *UpdateSecretInvoker {
	requestDef := GenReqDefForUpdateSecret()
	return &UpdateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateService 更新服务
//
// 更新一个服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

// UpdateServiceInvoker 更新服务
func (c *IefClient) UpdateServiceInvoker(request *model.UpdateServiceRequest) *UpdateServiceInvoker {
	requestDef := GenReqDefForUpdateService()
	return &UpdateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeEdgeNode 升级边缘节点
//
// 该API用于升级边缘节点。边缘节点将自动升级到最新的可用版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IefClient) UpgradeEdgeNode(request *model.UpgradeEdgeNodeRequest) (*model.UpgradeEdgeNodeResponse, error) {
	requestDef := GenReqDefForUpgradeEdgeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeEdgeNodeResponse), nil
	}
}

// UpgradeEdgeNodeInvoker 升级边缘节点
func (c *IefClient) UpgradeEdgeNodeInvoker(request *model.UpgradeEdgeNodeRequest) *UpgradeEdgeNodeInvoker {
	requestDef := GenReqDefForUpgradeEdgeNode()
	return &UpgradeEdgeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
