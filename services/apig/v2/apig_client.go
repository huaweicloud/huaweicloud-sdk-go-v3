package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2/model"
)

type ApigClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewApigClient(hcClient *httpclient.HcHttpClient) *ApigClient {
	return &ApigClient{HcClient: hcClient}
}

func ApigClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AcceptOrRejectEndpointConnections 接受或拒绝终端节点连接
//
// 接受或拒绝实例节点连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AcceptOrRejectEndpointConnections(request *model.AcceptOrRejectEndpointConnectionsRequest) (*model.AcceptOrRejectEndpointConnectionsResponse, error) {
	requestDef := GenReqDefForAcceptOrRejectEndpointConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptOrRejectEndpointConnectionsResponse), nil
	}
}

// AcceptOrRejectEndpointConnectionsInvoker 接受或拒绝终端节点连接
func (c *ApigClient) AcceptOrRejectEndpointConnectionsInvoker(request *model.AcceptOrRejectEndpointConnectionsRequest) *AcceptOrRejectEndpointConnectionsInvoker {
	requestDef := GenReqDefForAcceptOrRejectEndpointConnections()
	return &AcceptOrRejectEndpointConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddCustomIngressPort 新增实例的自定义入方向端口
//
// 新增实例的自定义入方向端口，在同个实例中，一个端口仅能支持一种协议。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddCustomIngressPort(request *model.AddCustomIngressPortRequest) (*model.AddCustomIngressPortResponse, error) {
	requestDef := GenReqDefForAddCustomIngressPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCustomIngressPortResponse), nil
	}
}

// AddCustomIngressPortInvoker 新增实例的自定义入方向端口
func (c *ApigClient) AddCustomIngressPortInvoker(request *model.AddCustomIngressPortRequest) *AddCustomIngressPortInvoker {
	requestDef := GenReqDefForAddCustomIngressPort()
	return &AddCustomIngressPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEipV2 实例更新或绑定EIP
//
// 实例更新或绑定EIP(仅当实例为LVS类型时支持)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddEipV2(request *model.AddEipV2Request) (*model.AddEipV2Response, error) {
	requestDef := GenReqDefForAddEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEipV2Response), nil
	}
}

// AddEipV2Invoker 实例更新或绑定EIP
func (c *ApigClient) AddEipV2Invoker(request *model.AddEipV2Request) *AddEipV2Invoker {
	requestDef := GenReqDefForAddEipV2()
	return &AddEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEndpointPermissions 批量添加实例终端节点连接白名单
//
// 批量添加实例终端节点连接白名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddEndpointPermissions(request *model.AddEndpointPermissionsRequest) (*model.AddEndpointPermissionsResponse, error) {
	requestDef := GenReqDefForAddEndpointPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEndpointPermissionsResponse), nil
	}
}

// AddEndpointPermissionsInvoker 批量添加实例终端节点连接白名单
func (c *ApigClient) AddEndpointPermissionsInvoker(request *model.AddEndpointPermissionsRequest) *AddEndpointPermissionsInvoker {
	requestDef := GenReqDefForAddEndpointPermissions()
	return &AddEndpointPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEngressEipV2 开启实例公网出口
//
// 实例开启公网出口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddEngressEipV2(request *model.AddEngressEipV2Request) (*model.AddEngressEipV2Response, error) {
	requestDef := GenReqDefForAddEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEngressEipV2Response), nil
	}
}

// AddEngressEipV2Invoker 开启实例公网出口
func (c *ApigClient) AddEngressEipV2Invoker(request *model.AddEngressEipV2Request) *AddEngressEipV2Invoker {
	requestDef := GenReqDefForAddEngressEipV2()
	return &AddEngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddIngressEipV2 开启实例公网入口
//
// 开启实例开启公网入口，仅当实例为ELB类型时支持
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddIngressEipV2(request *model.AddIngressEipV2Request) (*model.AddIngressEipV2Response, error) {
	requestDef := GenReqDefForAddIngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddIngressEipV2Response), nil
	}
}

// AddIngressEipV2Invoker 开启实例公网入口
func (c *ApigClient) AddIngressEipV2Invoker(request *model.AddIngressEipV2Request) *AddIngressEipV2Invoker {
	requestDef := GenReqDefForAddIngressEipV2()
	return &AddIngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateAppsForAppQuota 凭据配额绑定凭据列表
//
// 凭据配额绑定凭据列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AssociateAppsForAppQuota(request *model.AssociateAppsForAppQuotaRequest) (*model.AssociateAppsForAppQuotaResponse, error) {
	requestDef := GenReqDefForAssociateAppsForAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateAppsForAppQuotaResponse), nil
	}
}

// AssociateAppsForAppQuotaInvoker 凭据配额绑定凭据列表
func (c *ApigClient) AssociateAppsForAppQuotaInvoker(request *model.AssociateAppsForAppQuotaRequest) *AssociateAppsForAppQuotaInvoker {
	requestDef := GenReqDefForAssociateAppsForAppQuota()
	return &AssociateAppsForAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateCertificateV2 绑定域名证书
//
// 如果创建API时，“定义API请求”使用HTTPS请求协议，那么在独立域名中需要添加SSL证书。
// 使用实例自定义入方向端口的特性时，相同的域名会同时绑定证书，注意开启/关闭客户端校验会对相同域名的不同端口同时生效。
// 本章节主要介绍为特定域名绑定证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AssociateCertificateV2(request *model.AssociateCertificateV2Request) (*model.AssociateCertificateV2Response, error) {
	requestDef := GenReqDefForAssociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateCertificateV2Response), nil
	}
}

// AssociateCertificateV2Invoker 绑定域名证书
func (c *ApigClient) AssociateCertificateV2Invoker(request *model.AssociateCertificateV2Request) *AssociateCertificateV2Invoker {
	requestDef := GenReqDefForAssociateCertificateV2()
	return &AssociateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateDomainV2 绑定域名
//
// 用户自定义的域名，需要增加A记录才能生效，具体方法请参见《云解析服务用户指南》的“添加A类型记录集”章节。
//
// 每个API分组下最多可绑定5个域名。绑定域名后，用户可通过自定义域名调用API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AssociateDomainV2(request *model.AssociateDomainV2Request) (*model.AssociateDomainV2Response, error) {
	requestDef := GenReqDefForAssociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateDomainV2Response), nil
	}
}

// AssociateDomainV2Invoker 绑定域名
func (c *ApigClient) AssociateDomainV2Invoker(request *model.AssociateDomainV2Request) *AssociateDomainV2Invoker {
	requestDef := GenReqDefForAssociateDomainV2()
	return &AssociateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateSignatureKeyV2 绑定签名密钥
//
// 签名密钥创建后，需要绑定到API才能生效。
//
// 将签名密钥绑定到API后，则API网关请求后端服务时就会使用这个签名密钥进行加密签名，后端服务可以校验这个签名来验证请求来源。
//
// 将指定的签名密钥绑定到一个或多个已发布的API上。同一个API发布到不同的环境可以绑定不同的签名密钥；一个API在发布到特定环境后只能绑定一个签名密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AssociateSignatureKeyV2(request *model.AssociateSignatureKeyV2Request) (*model.AssociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForAssociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSignatureKeyV2Response), nil
	}
}

// AssociateSignatureKeyV2Invoker 绑定签名密钥
func (c *ApigClient) AssociateSignatureKeyV2Invoker(request *model.AssociateSignatureKeyV2Request) *AssociateSignatureKeyV2Invoker {
	requestDef := GenReqDefForAssociateSignatureKeyV2()
	return &AssociateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachApiToPlugin 插件绑定API
//
// 绑定插件到API上。
// - 只能选择发布状态的API
// - 绑定以后及时生效
// - 修改插件后及时生效
// - 相同类型的插件只能绑定一个，如果再次绑定同类型的插件，那么已绑定的同类型插件将直接被覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AttachApiToPlugin(request *model.AttachApiToPluginRequest) (*model.AttachApiToPluginResponse, error) {
	requestDef := GenReqDefForAttachApiToPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachApiToPluginResponse), nil
	}
}

// AttachApiToPluginInvoker 插件绑定API
func (c *ApigClient) AttachApiToPluginInvoker(request *model.AttachApiToPluginRequest) *AttachApiToPluginInvoker {
	requestDef := GenReqDefForAttachApiToPlugin()
	return &AttachApiToPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachPluginToApi API绑定插件
//
// 绑定插件到API上。
// - 只能选择发布状态的API
// - 绑定以后及时生效
// - 修改插件后及时生效
// - 相同类型的插件只能绑定一个，如果再次绑定同类型的插件，那么已绑定的同类型插件将直接被覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AttachPluginToApi(request *model.AttachPluginToApiRequest) (*model.AttachPluginToApiResponse, error) {
	requestDef := GenReqDefForAttachPluginToApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachPluginToApiResponse), nil
	}
}

// AttachPluginToApiInvoker API绑定插件
func (c *ApigClient) AttachPluginToApiInvoker(request *model.AttachPluginToApiRequest) *AttachPluginToApiInvoker {
	requestDef := GenReqDefForAttachPluginToApi()
	return &AttachPluginToApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateOrDeleteInstanceTags 批量添加或删除单个实例的标签
//
// 批量添加或删除单个实例的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchCreateOrDeleteInstanceTags(request *model.BatchCreateOrDeleteInstanceTagsRequest) (*model.BatchCreateOrDeleteInstanceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteInstanceTagsResponse), nil
	}
}

// BatchCreateOrDeleteInstanceTagsInvoker 批量添加或删除单个实例的标签
func (c *ApigClient) BatchCreateOrDeleteInstanceTagsInvoker(request *model.BatchCreateOrDeleteInstanceTagsRequest) *BatchCreateOrDeleteInstanceTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteInstanceTags()
	return &BatchCreateOrDeleteInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckAppV2 校验APP
//
// 校验app是否存在，非APP所有者可以调用该接口校验APP是否真实存在。这个接口只展示app的基本信息id 、name、
// remark，其他信息不显示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CheckAppV2(request *model.CheckAppV2Request) (*model.CheckAppV2Response, error) {
	requestDef := GenReqDefForCheckAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAppV2Response), nil
	}
}

// CheckAppV2Invoker 校验APP
func (c *ApigClient) CheckAppV2Invoker(request *model.CheckAppV2Request) *CheckAppV2Invoker {
	requestDef := GenReqDefForCheckAppV2()
	return &CheckAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAnAppV2 创建APP
//
// APP即应用，是一个可以访问API的身份标识。将API授权给APP后，APP即可调用API。
// 创建一个APP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAnAppV2(request *model.CreateAnAppV2Request) (*model.CreateAnAppV2Response, error) {
	requestDef := GenReqDefForCreateAnAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnAppV2Response), nil
	}
}

// CreateAnAppV2Invoker 创建APP
func (c *ApigClient) CreateAnAppV2Invoker(request *model.CreateAnAppV2Request) *CreateAnAppV2Invoker {
	requestDef := GenReqDefForCreateAnAppV2()
	return &CreateAnAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppCodeAutoV2 自动生成APP Code
//
// 创建App Code时，可以不指定具体值，由后台自动生成随机字符串填充。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAppCodeAutoV2(request *model.CreateAppCodeAutoV2Request) (*model.CreateAppCodeAutoV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeAutoV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeAutoV2Response), nil
	}
}

// CreateAppCodeAutoV2Invoker 自动生成APP Code
func (c *ApigClient) CreateAppCodeAutoV2Invoker(request *model.CreateAppCodeAutoV2Request) *CreateAppCodeAutoV2Invoker {
	requestDef := GenReqDefForCreateAppCodeAutoV2()
	return &CreateAppCodeAutoV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppCodeV2 创建APP Code
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAppCodeV2(request *model.CreateAppCodeV2Request) (*model.CreateAppCodeV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeV2Response), nil
	}
}

// CreateAppCodeV2Invoker 创建APP Code
func (c *ApigClient) CreateAppCodeV2Invoker(request *model.CreateAppCodeV2Request) *CreateAppCodeV2Invoker {
	requestDef := GenReqDefForCreateAppCodeV2()
	return &CreateAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppQuota 创建凭据配额
//
// 创建凭据配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAppQuota(request *model.CreateAppQuotaRequest) (*model.CreateAppQuotaResponse, error) {
	requestDef := GenReqDefForCreateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppQuotaResponse), nil
	}
}

// CreateAppQuotaInvoker 创建凭据配额
func (c *ApigClient) CreateAppQuotaInvoker(request *model.CreateAppQuotaRequest) *CreateAppQuotaInvoker {
	requestDef := GenReqDefForCreateAppQuota()
	return &CreateAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomAuthorizerV2 创建自定义认证
//
// 创建自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateCustomAuthorizerV2(request *model.CreateCustomAuthorizerV2Request) (*model.CreateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomAuthorizerV2Response), nil
	}
}

// CreateCustomAuthorizerV2Invoker 创建自定义认证
func (c *ApigClient) CreateCustomAuthorizerV2Invoker(request *model.CreateCustomAuthorizerV2Request) *CreateCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()
	return &CreateCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironmentV2 创建环境
//
// 在实际的生产中，API提供者可能有多个环境，如开发环境、测试环境、生产环境等，用户可以自由将API发布到某个环境，供调用者调用。
//
// 对于不同的环境，API的版本、请求地址甚至于包括请求消息等均有可能不同。如：某个API，v1.0的版本为稳定版本，发布到了生产环境供生产使用，同时，该API正处于迭代中，v1.1的版本是开发人员交付测试人员进行测试的版本，发布在测试环境上，而v1.2的版本目前开发团队正处于开发过程中，可以发布到开发环境进行自测等。
//
// 为此，API网关提供多环境管理功能，使租户能够最大化的模拟实际场景，低成本的接入API网关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateEnvironmentV2(request *model.CreateEnvironmentV2Request) (*model.CreateEnvironmentV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentV2Response), nil
	}
}

// CreateEnvironmentV2Invoker 创建环境
func (c *ApigClient) CreateEnvironmentV2Invoker(request *model.CreateEnvironmentV2Request) *CreateEnvironmentV2Invoker {
	requestDef := GenReqDefForCreateEnvironmentV2()
	return &CreateEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironmentVariableV2 新建变量
//
// 将API发布到不同的环境后，对于不同的环境，可能会有不同的环境变量，比如，API的服务部署地址，请求的版本号等。
//
// 用户可以定义不同的环境变量，用户在定义API时，在API的定义中使用这些变量，当调用API时，API网关会将这些变量替换成真实的变量值，以达到不同环境的区分效果。
//
// 环境变量定义在API分组上，该分组下的所有API都可以使用这些变量。
//
// &gt; 1. 环境变量的变量名称必须保持唯一，即一个分组在同一个环境上不能有两个同名的变量。
// &gt; 2. 环境变量区分大小写，即变量ABC与变量abc是两个不同的变量。
// &gt; 3. 设置了环境变量后，使用到该变量的API的调试功能将不可使用。
// &gt; 4. 定义了环境变量后，使用到环境变量的地方应该以对称的#标识环境变量，当API发布到相应的环境后，会对环境变量的值进行替换，如：定义的API的URL为：https://#address#:8080，环境变量address在RELEASE环境上的值为：192.168.1.5，则API发布到RELEASE环境后的真实的URL为：https://192.168.1.5:8080。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateEnvironmentVariableV2(request *model.CreateEnvironmentVariableV2Request) (*model.CreateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentVariableV2Response), nil
	}
}

// CreateEnvironmentVariableV2Invoker 新建变量
func (c *ApigClient) CreateEnvironmentVariableV2Invoker(request *model.CreateEnvironmentVariableV2Request) *CreateEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForCreateEnvironmentVariableV2()
	return &CreateEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFeatureV2 实例配置特性
//
// 为实例配置需要的特性。
//
// 支持配置的特性列表及特性配置示例请参考本手册中的“附录 &gt; 实例支持的APIG特性”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateFeatureV2(request *model.CreateFeatureV2Request) (*model.CreateFeatureV2Response, error) {
	requestDef := GenReqDefForCreateFeatureV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFeatureV2Response), nil
	}
}

// CreateFeatureV2Invoker 实例配置特性
func (c *ApigClient) CreateFeatureV2Invoker(request *model.CreateFeatureV2Request) *CreateFeatureV2Invoker {
	requestDef := GenReqDefForCreateFeatureV2()
	return &CreateFeatureV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGatewayResponseV2 创建分组自定义响应
//
// 新增分组下自定义响应
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateGatewayResponseV2(request *model.CreateGatewayResponseV2Request) (*model.CreateGatewayResponseV2Response, error) {
	requestDef := GenReqDefForCreateGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGatewayResponseV2Response), nil
	}
}

// CreateGatewayResponseV2Invoker 创建分组自定义响应
func (c *ApigClient) CreateGatewayResponseV2Invoker(request *model.CreateGatewayResponseV2Request) *CreateGatewayResponseV2Invoker {
	requestDef := GenReqDefForCreateGatewayResponseV2()
	return &CreateGatewayResponseV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceV2 创建专享版实例（按需）
//
// 创建按需专享版实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateInstanceV2(request *model.CreateInstanceV2Request) (*model.CreateInstanceV2Response, error) {
	requestDef := GenReqDefForCreateInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceV2Response), nil
	}
}

// CreateInstanceV2Invoker 创建专享版实例（按需）
func (c *ApigClient) CreateInstanceV2Invoker(request *model.CreateInstanceV2Request) *CreateInstanceV2Invoker {
	requestDef := GenReqDefForCreateInstanceV2()
	return &CreateInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrchestration 创建编排规则
//
// 创建编排规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateOrchestration(request *model.CreateOrchestrationRequest) (*model.CreateOrchestrationResponse, error) {
	requestDef := GenReqDefForCreateOrchestration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrchestrationResponse), nil
	}
}

// CreateOrchestrationInvoker 创建编排规则
func (c *ApigClient) CreateOrchestrationInvoker(request *model.CreateOrchestrationRequest) *CreateOrchestrationInvoker {
	requestDef := GenReqDefForCreateOrchestration()
	return &CreateOrchestrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrder 创建专享版实例（包周期）
//
// 创建包周期专享版实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateOrder(request *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	requestDef := GenReqDefForCreateOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrderResponse), nil
	}
}

// CreateOrderInvoker 创建专享版实例（包周期）
func (c *ApigClient) CreateOrderInvoker(request *model.CreateOrderRequest) *CreateOrderInvoker {
	requestDef := GenReqDefForCreateOrder()
	return &CreateOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlugin 创建插件
//
// 创建插件信息。
// - 插件不允许重名
// - 插件创建后未绑定API前是无意义的，绑定API后，对绑定的API即时生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreatePlugin(request *model.CreatePluginRequest) (*model.CreatePluginResponse, error) {
	requestDef := GenReqDefForCreatePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePluginResponse), nil
	}
}

// CreatePluginInvoker 创建插件
func (c *ApigClient) CreatePluginInvoker(request *model.CreatePluginRequest) *CreatePluginInvoker {
	requestDef := GenReqDefForCreatePlugin()
	return &CreatePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPayResizeOrder 按需规格变更
//
// 创建按需规格变更订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreatePostPayResizeOrder(request *model.CreatePostPayResizeOrderRequest) (*model.CreatePostPayResizeOrderResponse, error) {
	requestDef := GenReqDefForCreatePostPayResizeOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPayResizeOrderResponse), nil
	}
}

// CreatePostPayResizeOrderInvoker 按需规格变更
func (c *ApigClient) CreatePostPayResizeOrderInvoker(request *model.CreatePostPayResizeOrderRequest) *CreatePostPayResizeOrderInvoker {
	requestDef := GenReqDefForCreatePostPayResizeOrder()
	return &CreatePostPayResizeOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrepayResize 创建包周期规格变更订单
//
// 创建包周期规格变更订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreatePrepayResize(request *model.CreatePrepayResizeRequest) (*model.CreatePrepayResizeResponse, error) {
	requestDef := GenReqDefForCreatePrepayResize()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrepayResizeResponse), nil
	}
}

// CreatePrepayResizeInvoker 创建包周期规格变更订单
func (c *ApigClient) CreatePrepayResizeInvoker(request *model.CreatePrepayResizeRequest) *CreatePrepayResizeInvoker {
	requestDef := GenReqDefForCreatePrepayResize()
	return &CreatePrepayResizeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRequestThrottlingPolicyV2 创建流控策略
//
// 当API上线后，系统会默认给每个API提供一个流控策略，API提供者可以根据自身API的服务能力及负载情况变更这个流控策略。
// 流控策略即限制API在一定长度的时间内，能够允许被访问的最大次数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateRequestThrottlingPolicyV2(request *model.CreateRequestThrottlingPolicyV2Request) (*model.CreateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestThrottlingPolicyV2Response), nil
	}
}

// CreateRequestThrottlingPolicyV2Invoker 创建流控策略
func (c *ApigClient) CreateRequestThrottlingPolicyV2Invoker(request *model.CreateRequestThrottlingPolicyV2Request) *CreateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()
	return &CreateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSignatureKeyV2 创建签名密钥
//
// 为了保护API的安全性，建议租户为API的访问提供一套保护机制，即租户开放的API，需要对请求来源进行认证，不符合认证的请求直接拒绝访问。
//
// 其中，签名密钥就是API安全保护机制的一种。
//
// 租户创建一个签名密钥，并将签名密钥与API进行绑定，则API网关在请求这个API时，就会使用绑定的签名密钥对请求参数进行数据加密，生成签名。当租户的后端服务收到请求时，可以校验这个签名，如果签名校验不通过，则该请求不是API网关发出的请求，租户可以拒绝这个请求，从而保证API的安全性，避免API被未知来源的请求攻击。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateSignatureKeyV2(request *model.CreateSignatureKeyV2Request) (*model.CreateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForCreateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSignatureKeyV2Response), nil
	}
}

// CreateSignatureKeyV2Invoker 创建签名密钥
func (c *ApigClient) CreateSignatureKeyV2Invoker(request *model.CreateSignatureKeyV2Request) *CreateSignatureKeyV2Invoker {
	requestDef := GenReqDefForCreateSignatureKeyV2()
	return &CreateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSpecialThrottlingConfigurationV2 创建特殊设置
//
// 流控策略可以限制一段时间内可以访问API的最大次数，也可以限制一段时间内单个租户和单个APP可以访问API的最大次数。
//
// 如果想要对某个特定的APP进行特殊设置，例如设置所有APP每分钟的访问次数为500次，但想设置APP1每分钟的访问次数为800次，可以通过在流控策略中设置特殊APP来实现该功能。
//
// 为流控策略添加一个特殊设置的对象，可以是APP，也可以是租户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateSpecialThrottlingConfigurationV2(request *model.CreateSpecialThrottlingConfigurationV2Request) (*model.CreateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpecialThrottlingConfigurationV2Response), nil
	}
}

// CreateSpecialThrottlingConfigurationV2Invoker 创建特殊设置
func (c *ApigClient) CreateSpecialThrottlingConfigurationV2Invoker(request *model.CreateSpecialThrottlingConfigurationV2Request) *CreateSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()
	return &CreateSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppAcl 删除APP的访问控制
//
// 删除凭据的访问控制信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteAppAcl(request *model.DeleteAppAclRequest) (*model.DeleteAppAclResponse, error) {
	requestDef := GenReqDefForDeleteAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppAclResponse), nil
	}
}

// DeleteAppAclInvoker 删除APP的访问控制
func (c *ApigClient) DeleteAppAclInvoker(request *model.DeleteAppAclRequest) *DeleteAppAclInvoker {
	requestDef := GenReqDefForDeleteAppAcl()
	return &DeleteAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppCodeV2 删除APP Code
//
// 删除App Code，App Code删除后，将无法再通过简易认证访问对应的API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteAppCodeV2(request *model.DeleteAppCodeV2Request) (*model.DeleteAppCodeV2Response, error) {
	requestDef := GenReqDefForDeleteAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppCodeV2Response), nil
	}
}

// DeleteAppCodeV2Invoker 删除APP Code
func (c *ApigClient) DeleteAppCodeV2Invoker(request *model.DeleteAppCodeV2Request) *DeleteAppCodeV2Invoker {
	requestDef := GenReqDefForDeleteAppCodeV2()
	return &DeleteAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppQuota 删除凭据配额
//
// 删除凭据配额。删除凭据配额时，同时删除凭据配额和凭据的关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteAppQuota(request *model.DeleteAppQuotaRequest) (*model.DeleteAppQuotaResponse, error) {
	requestDef := GenReqDefForDeleteAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppQuotaResponse), nil
	}
}

// DeleteAppQuotaInvoker 删除凭据配额
func (c *ApigClient) DeleteAppQuotaInvoker(request *model.DeleteAppQuotaRequest) *DeleteAppQuotaInvoker {
	requestDef := GenReqDefForDeleteAppQuota()
	return &DeleteAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppV2 删除APP
//
// 删除指定的APP。
// APP删除后，将无法再调用任何API[；其中，云商店自动创建的APP无法被删除](tag:hws)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteAppV2(request *model.DeleteAppV2Request) (*model.DeleteAppV2Response, error) {
	requestDef := GenReqDefForDeleteAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppV2Response), nil
	}
}

// DeleteAppV2Invoker 删除APP
func (c *ApigClient) DeleteAppV2Invoker(request *model.DeleteAppV2Request) *DeleteAppV2Invoker {
	requestDef := GenReqDefForDeleteAppV2()
	return &DeleteAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomAuthorizerV2 删除自定义认证
//
// 删除自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteCustomAuthorizerV2(request *model.DeleteCustomAuthorizerV2Request) (*model.DeleteCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomAuthorizerV2Response), nil
	}
}

// DeleteCustomAuthorizerV2Invoker 删除自定义认证
func (c *ApigClient) DeleteCustomAuthorizerV2Invoker(request *model.DeleteCustomAuthorizerV2Request) *DeleteCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()
	return &DeleteCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomIngressPort 删除实例指定的自定义入方向端口
//
// 删除实例指定的自定义入方向端口，不包含默认端口80和443。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteCustomIngressPort(request *model.DeleteCustomIngressPortRequest) (*model.DeleteCustomIngressPortResponse, error) {
	requestDef := GenReqDefForDeleteCustomIngressPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomIngressPortResponse), nil
	}
}

// DeleteCustomIngressPortInvoker 删除实例指定的自定义入方向端口
func (c *ApigClient) DeleteCustomIngressPortInvoker(request *model.DeleteCustomIngressPortRequest) *DeleteCustomIngressPortInvoker {
	requestDef := GenReqDefForDeleteCustomIngressPort()
	return &DeleteCustomIngressPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpointPermissions 批量删除实例终端节点连接白名单
//
// 批量删除实例终端节点连接白名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteEndpointPermissions(request *model.DeleteEndpointPermissionsRequest) (*model.DeleteEndpointPermissionsResponse, error) {
	requestDef := GenReqDefForDeleteEndpointPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointPermissionsResponse), nil
	}
}

// DeleteEndpointPermissionsInvoker 批量删除实例终端节点连接白名单
func (c *ApigClient) DeleteEndpointPermissionsInvoker(request *model.DeleteEndpointPermissionsRequest) *DeleteEndpointPermissionsInvoker {
	requestDef := GenReqDefForDeleteEndpointPermissions()
	return &DeleteEndpointPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironmentV2 删除环境
//
// 删除指定的环境。
//
// 该操作将导致此API在指定的环境无法被访问，可能会影响相当一部分应用和用户。请确保已经告知用户，或者确认需要强制下线。
//
// 环境上存在已发布的API时，该环境不能被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteEnvironmentV2(request *model.DeleteEnvironmentV2Request) (*model.DeleteEnvironmentV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentV2Response), nil
	}
}

// DeleteEnvironmentV2Invoker 删除环境
func (c *ApigClient) DeleteEnvironmentV2Invoker(request *model.DeleteEnvironmentV2Request) *DeleteEnvironmentV2Invoker {
	requestDef := GenReqDefForDeleteEnvironmentV2()
	return &DeleteEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironmentVariableV2 删除变量
//
// 删除指定的环境变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteEnvironmentVariableV2(request *model.DeleteEnvironmentVariableV2Request) (*model.DeleteEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentVariableV2Response), nil
	}
}

// DeleteEnvironmentVariableV2Invoker 删除变量
func (c *ApigClient) DeleteEnvironmentVariableV2Invoker(request *model.DeleteEnvironmentVariableV2Request) *DeleteEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()
	return &DeleteEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGatewayResponseTypeV2 删除分组指定错误类型的自定义响应配置
//
// 删除分组指定错误类型的自定义响应配置，还原为使用默认值的配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteGatewayResponseTypeV2(request *model.DeleteGatewayResponseTypeV2Request) (*model.DeleteGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForDeleteGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGatewayResponseTypeV2Response), nil
	}
}

// DeleteGatewayResponseTypeV2Invoker 删除分组指定错误类型的自定义响应配置
func (c *ApigClient) DeleteGatewayResponseTypeV2Invoker(request *model.DeleteGatewayResponseTypeV2Request) *DeleteGatewayResponseTypeV2Invoker {
	requestDef := GenReqDefForDeleteGatewayResponseTypeV2()
	return &DeleteGatewayResponseTypeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGatewayResponseV2 删除分组自定义响应
//
// 删除分组自定义响应
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteGatewayResponseV2(request *model.DeleteGatewayResponseV2Request) (*model.DeleteGatewayResponseV2Response, error) {
	requestDef := GenReqDefForDeleteGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGatewayResponseV2Response), nil
	}
}

// DeleteGatewayResponseV2Invoker 删除分组自定义响应
func (c *ApigClient) DeleteGatewayResponseV2Invoker(request *model.DeleteGatewayResponseV2Request) *DeleteGatewayResponseV2Invoker {
	requestDef := GenReqDefForDeleteGatewayResponseV2()
	return &DeleteGatewayResponseV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstancesV2 删除专享版实例
//
// 删除专享版实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteInstancesV2(request *model.DeleteInstancesV2Request) (*model.DeleteInstancesV2Response, error) {
	requestDef := GenReqDefForDeleteInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesV2Response), nil
	}
}

// DeleteInstancesV2Invoker 删除专享版实例
func (c *ApigClient) DeleteInstancesV2Invoker(request *model.DeleteInstancesV2Request) *DeleteInstancesV2Invoker {
	requestDef := GenReqDefForDeleteInstancesV2()
	return &DeleteInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOrchestration 删除编排规则
//
// 删除编排规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteOrchestration(request *model.DeleteOrchestrationRequest) (*model.DeleteOrchestrationResponse, error) {
	requestDef := GenReqDefForDeleteOrchestration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrchestrationResponse), nil
	}
}

// DeleteOrchestrationInvoker 删除编排规则
func (c *ApigClient) DeleteOrchestrationInvoker(request *model.DeleteOrchestrationRequest) *DeleteOrchestrationInvoker {
	requestDef := GenReqDefForDeleteOrchestration()
	return &DeleteOrchestrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlugin 删除插件
//
// 删除插件。
// - 必须先解除API和插件的绑定关系，否则删除报错
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeletePlugin(request *model.DeletePluginRequest) (*model.DeletePluginResponse, error) {
	requestDef := GenReqDefForDeletePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePluginResponse), nil
	}
}

// DeletePluginInvoker 删除插件
func (c *ApigClient) DeletePluginInvoker(request *model.DeletePluginRequest) *DeletePluginInvoker {
	requestDef := GenReqDefForDeletePlugin()
	return &DeletePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRequestThrottlingPolicyV2 删除流控策略
//
// 删除指定的流控策略，以及该流控策略与API的所有绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteRequestThrottlingPolicyV2(request *model.DeleteRequestThrottlingPolicyV2Request) (*model.DeleteRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestThrottlingPolicyV2Response), nil
	}
}

// DeleteRequestThrottlingPolicyV2Invoker 删除流控策略
func (c *ApigClient) DeleteRequestThrottlingPolicyV2Invoker(request *model.DeleteRequestThrottlingPolicyV2Request) *DeleteRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()
	return &DeleteRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSignatureKeyV2 删除签名密钥
//
// 删除指定的签名密钥，删除签名密钥时，其配置的绑定关系会一并删除，相应的签名密钥会失效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteSignatureKeyV2(request *model.DeleteSignatureKeyV2Request) (*model.DeleteSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDeleteSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSignatureKeyV2Response), nil
	}
}

// DeleteSignatureKeyV2Invoker 删除签名密钥
func (c *ApigClient) DeleteSignatureKeyV2Invoker(request *model.DeleteSignatureKeyV2Request) *DeleteSignatureKeyV2Invoker {
	requestDef := GenReqDefForDeleteSignatureKeyV2()
	return &DeleteSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSpecialThrottlingConfigurationV2 删除特殊设置
//
// 删除某个流控策略的某个特殊配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteSpecialThrottlingConfigurationV2(request *model.DeleteSpecialThrottlingConfigurationV2Request) (*model.DeleteSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSpecialThrottlingConfigurationV2Response), nil
	}
}

// DeleteSpecialThrottlingConfigurationV2Invoker 删除特殊设置
func (c *ApigClient) DeleteSpecialThrottlingConfigurationV2Invoker(request *model.DeleteSpecialThrottlingConfigurationV2Request) *DeleteSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()
	return &DeleteSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachApiFromPlugin 解除绑定插件的API
//
// 解除绑定在插件上的API。
// - 解绑及时生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DetachApiFromPlugin(request *model.DetachApiFromPluginRequest) (*model.DetachApiFromPluginResponse, error) {
	requestDef := GenReqDefForDetachApiFromPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachApiFromPluginResponse), nil
	}
}

// DetachApiFromPluginInvoker 解除绑定插件的API
func (c *ApigClient) DetachApiFromPluginInvoker(request *model.DetachApiFromPluginRequest) *DetachApiFromPluginInvoker {
	requestDef := GenReqDefForDetachApiFromPlugin()
	return &DetachApiFromPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachPluginFromApi 解除绑定API的插件
//
// 解除绑定在API上的插件。
// - 解绑及时生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DetachPluginFromApi(request *model.DetachPluginFromApiRequest) (*model.DetachPluginFromApiResponse, error) {
	requestDef := GenReqDefForDetachPluginFromApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachPluginFromApiResponse), nil
	}
}

// DetachPluginFromApiInvoker 解除绑定API的插件
func (c *ApigClient) DetachPluginFromApiInvoker(request *model.DetachPluginFromApiRequest) *DetachPluginFromApiInvoker {
	requestDef := GenReqDefForDetachPluginFromApi()
	return &DetachPluginFromApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateAppQuotaWithApp 解除凭据配额和凭据的绑定
//
// 解除凭据配额和凭据的绑定
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DisassociateAppQuotaWithApp(request *model.DisassociateAppQuotaWithAppRequest) (*model.DisassociateAppQuotaWithAppResponse, error) {
	requestDef := GenReqDefForDisassociateAppQuotaWithApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAppQuotaWithAppResponse), nil
	}
}

// DisassociateAppQuotaWithAppInvoker 解除凭据配额和凭据的绑定
func (c *ApigClient) DisassociateAppQuotaWithAppInvoker(request *model.DisassociateAppQuotaWithAppRequest) *DisassociateAppQuotaWithAppInvoker {
	requestDef := GenReqDefForDisassociateAppQuotaWithApp()
	return &DisassociateAppQuotaWithAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateCertificateV2 删除域名证书
//
// 如果域名证书不再需要或者已过期，则可以删除证书内容。在使用自定义入方向端口的特性时，相同的域名会同时解绑证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DisassociateCertificateV2(request *model.DisassociateCertificateV2Request) (*model.DisassociateCertificateV2Response, error) {
	requestDef := GenReqDefForDisassociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateCertificateV2Response), nil
	}
}

// DisassociateCertificateV2Invoker 删除域名证书
func (c *ApigClient) DisassociateCertificateV2Invoker(request *model.DisassociateCertificateV2Request) *DisassociateCertificateV2Invoker {
	requestDef := GenReqDefForDisassociateCertificateV2()
	return &DisassociateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateDomainV2 解绑域名
//
// 如果API分组不再需要绑定某个自定义域名，则可以为此API分组解绑此域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DisassociateDomainV2(request *model.DisassociateDomainV2Request) (*model.DisassociateDomainV2Response, error) {
	requestDef := GenReqDefForDisassociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateDomainV2Response), nil
	}
}

// DisassociateDomainV2Invoker 解绑域名
func (c *ApigClient) DisassociateDomainV2Invoker(request *model.DisassociateDomainV2Request) *DisassociateDomainV2Invoker {
	requestDef := GenReqDefForDisassociateDomainV2()
	return &DisassociateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateSignatureKeyV2 解除API与签名密钥的绑定关系
//
// 解除API与签名密钥的绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DisassociateSignatureKeyV2(request *model.DisassociateSignatureKeyV2Request) (*model.DisassociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateSignatureKeyV2Response), nil
	}
}

// DisassociateSignatureKeyV2Invoker 解除API与签名密钥的绑定关系
func (c *ApigClient) DisassociateSignatureKeyV2Invoker(request *model.DisassociateSignatureKeyV2Request) *DisassociateSignatureKeyV2Invoker {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()
	return &DisassociateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportApiDefinitionsAsync 异步导出API
//
// 导出分组下API的定义信息。导出文件内容符合swagger标准规范，API网关自定义扩展字段请参考《API网关用户指南》的“导入导出API：扩展定义”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ExportApiDefinitionsAsync(request *model.ExportApiDefinitionsAsyncRequest) (*model.ExportApiDefinitionsAsyncResponse, error) {
	requestDef := GenReqDefForExportApiDefinitionsAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportApiDefinitionsAsyncResponse), nil
	}
}

// ExportApiDefinitionsAsyncInvoker 异步导出API
func (c *ApigClient) ExportApiDefinitionsAsyncInvoker(request *model.ExportApiDefinitionsAsyncRequest) *ExportApiDefinitionsAsyncInvoker {
	requestDef := GenReqDefForExportApiDefinitionsAsync()
	return &ExportApiDefinitionsAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportApiDefinitionsAsync 异步导入API
//
// 导入API。导入文件内容需要符合swagger标准规范，API网关自定义扩展字段请参考《API网关用户指南》的“导入导出API：扩展定义”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ImportApiDefinitionsAsync(request *model.ImportApiDefinitionsAsyncRequest) (*model.ImportApiDefinitionsAsyncResponse, error) {
	requestDef := GenReqDefForImportApiDefinitionsAsync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApiDefinitionsAsyncResponse), nil
	}
}

// ImportApiDefinitionsAsyncInvoker 异步导入API
func (c *ApigClient) ImportApiDefinitionsAsyncInvoker(request *model.ImportApiDefinitionsAsyncRequest) *ImportApiDefinitionsAsyncInvoker {
	requestDef := GenReqDefForImportApiDefinitionsAsync()
	return &ImportApiDefinitionsAsyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportMicroservice 导入微服务
//
// 导入微服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ImportMicroservice(request *model.ImportMicroserviceRequest) (*model.ImportMicroserviceResponse, error) {
	requestDef := GenReqDefForImportMicroservice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportMicroserviceResponse), nil
	}
}

// ImportMicroserviceInvoker 导入微服务
func (c *ApigClient) ImportMicroserviceInvoker(request *model.ImportMicroserviceRequest) *ImportMicroserviceInvoker {
	requestDef := GenReqDefForImportMicroservice()
	return &ImportMicroserviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiAttachablePlugins 查询可绑定当前API的插件
//
// 查询可绑定当前API的插件信息。
// - 支持分页返回
// - 支持插件名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiAttachablePlugins(request *model.ListApiAttachablePluginsRequest) (*model.ListApiAttachablePluginsResponse, error) {
	requestDef := GenReqDefForListApiAttachablePlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiAttachablePluginsResponse), nil
	}
}

// ListApiAttachablePluginsInvoker 查询可绑定当前API的插件
func (c *ApigClient) ListApiAttachablePluginsInvoker(request *model.ListApiAttachablePluginsRequest) *ListApiAttachablePluginsInvoker {
	requestDef := GenReqDefForListApiAttachablePlugins()
	return &ListApiAttachablePluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiAttachedPlugins 查询API下绑定的插件
//
// 查询指定API下绑定的插件信息。
// - 用于查询指定API下已经绑定的插件列表信息
// - 支持分页返回
// - 支持插件名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiAttachedPlugins(request *model.ListApiAttachedPluginsRequest) (*model.ListApiAttachedPluginsResponse, error) {
	requestDef := GenReqDefForListApiAttachedPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiAttachedPluginsResponse), nil
	}
}

// ListApiAttachedPluginsInvoker 查询API下绑定的插件
func (c *ApigClient) ListApiAttachedPluginsInvoker(request *model.ListApiAttachedPluginsRequest) *ListApiAttachedPluginsInvoker {
	requestDef := GenReqDefForListApiAttachedPlugins()
	return &ListApiAttachedPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiGroupsQuantitiesV2 查询API分组概况
//
// 查询租户名下的API分组概况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiGroupsQuantitiesV2(request *model.ListApiGroupsQuantitiesV2Request) (*model.ListApiGroupsQuantitiesV2Response, error) {
	requestDef := GenReqDefForListApiGroupsQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiGroupsQuantitiesV2Response), nil
	}
}

// ListApiGroupsQuantitiesV2Invoker 查询API分组概况
func (c *ApigClient) ListApiGroupsQuantitiesV2Invoker(request *model.ListApiGroupsQuantitiesV2Request) *ListApiGroupsQuantitiesV2Invoker {
	requestDef := GenReqDefForListApiGroupsQuantitiesV2()
	return &ListApiGroupsQuantitiesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiQuantitiesV2 查询API概况
//
// 查询租户名下的API概况：已发布到RELEASE环境的API个数，未发布到RELEASE环境的API个数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiQuantitiesV2(request *model.ListApiQuantitiesV2Request) (*model.ListApiQuantitiesV2Response, error) {
	requestDef := GenReqDefForListApiQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiQuantitiesV2Response), nil
	}
}

// ListApiQuantitiesV2Invoker 查询API概况
func (c *ApigClient) ListApiQuantitiesV2Invoker(request *model.ListApiQuantitiesV2Request) *ListApiQuantitiesV2Invoker {
	requestDef := GenReqDefForListApiQuantitiesV2()
	return &ListApiQuantitiesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToSignatureKeyV2 查看签名密钥绑定的API列表
//
// 查询某个签名密钥上已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToSignatureKeyV2(request *model.ListApisBindedToSignatureKeyV2Request) (*model.ListApisBindedToSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToSignatureKeyV2Response), nil
	}
}

// ListApisBindedToSignatureKeyV2Invoker 查看签名密钥绑定的API列表
func (c *ApigClient) ListApisBindedToSignatureKeyV2Invoker(request *model.ListApisBindedToSignatureKeyV2Request) *ListApisBindedToSignatureKeyV2Invoker {
	requestDef := GenReqDefForListApisBindedToSignatureKeyV2()
	return &ListApisBindedToSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisNotBoundWithSignatureKeyV2 查看签名密钥未绑定的API列表
//
// 查询所有未绑定到该签名密钥上的API列表。需要API已经发布，未发布的API不予展示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisNotBoundWithSignatureKeyV2(request *model.ListApisNotBoundWithSignatureKeyV2Request) (*model.ListApisNotBoundWithSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisNotBoundWithSignatureKeyV2Response), nil
	}
}

// ListApisNotBoundWithSignatureKeyV2Invoker 查看签名密钥未绑定的API列表
func (c *ApigClient) ListApisNotBoundWithSignatureKeyV2Invoker(request *model.ListApisNotBoundWithSignatureKeyV2Request) *ListApisNotBoundWithSignatureKeyV2Invoker {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()
	return &ListApisNotBoundWithSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppCodesV2 查询APP Code列表
//
// 查询App Code列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppCodesV2(request *model.ListAppCodesV2Request) (*model.ListAppCodesV2Response, error) {
	requestDef := GenReqDefForListAppCodesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppCodesV2Response), nil
	}
}

// ListAppCodesV2Invoker 查询APP Code列表
func (c *ApigClient) ListAppCodesV2Invoker(request *model.ListAppCodesV2Request) *ListAppCodesV2Invoker {
	requestDef := GenReqDefForListAppCodesV2()
	return &ListAppCodesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuantitiesV2 查询APP概况
//
// 查询租户名下的APP概况：已进行API访问授权的APP个数，未进行API访问授权的APP个数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppQuantitiesV2(request *model.ListAppQuantitiesV2Request) (*model.ListAppQuantitiesV2Response, error) {
	requestDef := GenReqDefForListAppQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuantitiesV2Response), nil
	}
}

// ListAppQuantitiesV2Invoker 查询APP概况
func (c *ApigClient) ListAppQuantitiesV2Invoker(request *model.ListAppQuantitiesV2Request) *ListAppQuantitiesV2Invoker {
	requestDef := GenReqDefForListAppQuantitiesV2()
	return &ListAppQuantitiesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotaBindableApps 查询凭据配额可绑定的凭据列表
//
// 查询凭据配额可绑定的凭据列表。支持按凭据名称模糊搜索
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppQuotaBindableApps(request *model.ListAppQuotaBindableAppsRequest) (*model.ListAppQuotaBindableAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBindableApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBindableAppsResponse), nil
	}
}

// ListAppQuotaBindableAppsInvoker 查询凭据配额可绑定的凭据列表
func (c *ApigClient) ListAppQuotaBindableAppsInvoker(request *model.ListAppQuotaBindableAppsRequest) *ListAppQuotaBindableAppsInvoker {
	requestDef := GenReqDefForListAppQuotaBindableApps()
	return &ListAppQuotaBindableAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotaBoundApps 查询凭据配额已绑定的凭据列表
//
// 查询凭据配额已绑定的凭据列表。支持按凭据名称模糊匹配
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppQuotaBoundApps(request *model.ListAppQuotaBoundAppsRequest) (*model.ListAppQuotaBoundAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBoundApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBoundAppsResponse), nil
	}
}

// ListAppQuotaBoundAppsInvoker 查询凭据配额已绑定的凭据列表
func (c *ApigClient) ListAppQuotaBoundAppsInvoker(request *model.ListAppQuotaBoundAppsRequest) *ListAppQuotaBoundAppsInvoker {
	requestDef := GenReqDefForListAppQuotaBoundApps()
	return &ListAppQuotaBoundAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotas 获取凭据配额列表
//
// 获取凭据配额列表。支持根据名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppQuotas(request *model.ListAppQuotasRequest) (*model.ListAppQuotasResponse, error) {
	requestDef := GenReqDefForListAppQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotasResponse), nil
	}
}

// ListAppQuotasInvoker 获取凭据配额列表
func (c *ApigClient) ListAppQuotasInvoker(request *model.ListAppQuotasRequest) *ListAppQuotasInvoker {
	requestDef := GenReqDefForListAppQuotas()
	return &ListAppQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppsV2 查询APP列表
//
// 查询APP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppsV2(request *model.ListAppsV2Request) (*model.ListAppsV2Response, error) {
	requestDef := GenReqDefForListAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsV2Response), nil
	}
}

// ListAppsV2Invoker 查询APP列表
func (c *ApigClient) ListAppsV2Invoker(request *model.ListAppsV2Request) *ListAppsV2Invoker {
	requestDef := GenReqDefForListAppsV2()
	return &ListAppsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZonesV2 查看可用区信息
//
// 查看可用区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAvailableZonesV2(request *model.ListAvailableZonesV2Request) (*model.ListAvailableZonesV2Response, error) {
	requestDef := GenReqDefForListAvailableZonesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesV2Response), nil
	}
}

// ListAvailableZonesV2Invoker 查看可用区信息
func (c *ApigClient) ListAvailableZonesV2Invoker(request *model.ListAvailableZonesV2Request) *ListAvailableZonesV2Invoker {
	requestDef := GenReqDefForListAvailableZonesV2()
	return &ListAvailableZonesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomAuthorizersV2 查询自定义认证列表
//
// 查询自定义认证列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListCustomAuthorizersV2(request *model.ListCustomAuthorizersV2Request) (*model.ListCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForListCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomAuthorizersV2Response), nil
	}
}

// ListCustomAuthorizersV2Invoker 查询自定义认证列表
func (c *ApigClient) ListCustomAuthorizersV2Invoker(request *model.ListCustomAuthorizersV2Request) *ListCustomAuthorizersV2Invoker {
	requestDef := GenReqDefForListCustomAuthorizersV2()
	return &ListCustomAuthorizersV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomIngressPortDomains 查询实例指定的自定义入方向端口绑定的域名信息
//
// 查询实例指定的自定义入方向端口绑定的域名信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListCustomIngressPortDomains(request *model.ListCustomIngressPortDomainsRequest) (*model.ListCustomIngressPortDomainsResponse, error) {
	requestDef := GenReqDefForListCustomIngressPortDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomIngressPortDomainsResponse), nil
	}
}

// ListCustomIngressPortDomainsInvoker 查询实例指定的自定义入方向端口绑定的域名信息
func (c *ApigClient) ListCustomIngressPortDomainsInvoker(request *model.ListCustomIngressPortDomainsRequest) *ListCustomIngressPortDomainsInvoker {
	requestDef := GenReqDefForListCustomIngressPortDomains()
	return &ListCustomIngressPortDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomIngressPorts 查询实例的自定义入方向端口列表
//
// 查询实例的自定义入方向端口列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListCustomIngressPorts(request *model.ListCustomIngressPortsRequest) (*model.ListCustomIngressPortsResponse, error) {
	requestDef := GenReqDefForListCustomIngressPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomIngressPortsResponse), nil
	}
}

// ListCustomIngressPortsInvoker 查询实例的自定义入方向端口列表
func (c *ApigClient) ListCustomIngressPortsInvoker(request *model.ListCustomIngressPortsRequest) *ListCustomIngressPortsInvoker {
	requestDef := GenReqDefForListCustomIngressPorts()
	return &ListCustomIngressPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointConnections 查询实例终端节点连接列表
//
// 查询实例终端节点连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListEndpointConnections(request *model.ListEndpointConnectionsRequest) (*model.ListEndpointConnectionsResponse, error) {
	requestDef := GenReqDefForListEndpointConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointConnectionsResponse), nil
	}
}

// ListEndpointConnectionsInvoker 查询实例终端节点连接列表
func (c *ApigClient) ListEndpointConnectionsInvoker(request *model.ListEndpointConnectionsRequest) *ListEndpointConnectionsInvoker {
	requestDef := GenReqDefForListEndpointConnections()
	return &ListEndpointConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointPermissions 查询实例的终端节点服务的白名单列表
//
// 查询当前实例终端节点服务的白名单列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListEndpointPermissions(request *model.ListEndpointPermissionsRequest) (*model.ListEndpointPermissionsResponse, error) {
	requestDef := GenReqDefForListEndpointPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointPermissionsResponse), nil
	}
}

// ListEndpointPermissionsInvoker 查询实例的终端节点服务的白名单列表
func (c *ApigClient) ListEndpointPermissionsInvoker(request *model.ListEndpointPermissionsRequest) *ListEndpointPermissionsInvoker {
	requestDef := GenReqDefForListEndpointPermissions()
	return &ListEndpointPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentVariablesV2 查询变量列表
//
// 查询分组下的所有环境变量的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListEnvironmentVariablesV2(request *model.ListEnvironmentVariablesV2Request) (*model.ListEnvironmentVariablesV2Response, error) {
	requestDef := GenReqDefForListEnvironmentVariablesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentVariablesV2Response), nil
	}
}

// ListEnvironmentVariablesV2Invoker 查询变量列表
func (c *ApigClient) ListEnvironmentVariablesV2Invoker(request *model.ListEnvironmentVariablesV2Request) *ListEnvironmentVariablesV2Invoker {
	requestDef := GenReqDefForListEnvironmentVariablesV2()
	return &ListEnvironmentVariablesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentsV2 查询环境列表
//
// 查询符合条件的环境列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListEnvironmentsV2(request *model.ListEnvironmentsV2Request) (*model.ListEnvironmentsV2Response, error) {
	requestDef := GenReqDefForListEnvironmentsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsV2Response), nil
	}
}

// ListEnvironmentsV2Invoker 查询环境列表
func (c *ApigClient) ListEnvironmentsV2Invoker(request *model.ListEnvironmentsV2Request) *ListEnvironmentsV2Invoker {
	requestDef := GenReqDefForListEnvironmentsV2()
	return &ListEnvironmentsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeaturesV2 查看实例特性列表
//
// 查看实例特性列表。注意：实例不支持以下特性的需要联系技术支持升级实例版本。
//
// 支持配置的特性列表及特性配置示例请参考本手册中的“附录 &gt; 实例支持的APIG特性”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListFeaturesV2(request *model.ListFeaturesV2Request) (*model.ListFeaturesV2Response, error) {
	requestDef := GenReqDefForListFeaturesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesV2Response), nil
	}
}

// ListFeaturesV2Invoker 查看实例特性列表
func (c *ApigClient) ListFeaturesV2Invoker(request *model.ListFeaturesV2Request) *ListFeaturesV2Invoker {
	requestDef := GenReqDefForListFeaturesV2()
	return &ListFeaturesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGatewayResponsesV2 查询分组自定义响应列表
//
// 查询分组自定义响应列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListGatewayResponsesV2(request *model.ListGatewayResponsesV2Request) (*model.ListGatewayResponsesV2Response, error) {
	requestDef := GenReqDefForListGatewayResponsesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGatewayResponsesV2Response), nil
	}
}

// ListGatewayResponsesV2Invoker 查询分组自定义响应列表
func (c *ApigClient) ListGatewayResponsesV2Invoker(request *model.ListGatewayResponsesV2Request) *ListGatewayResponsesV2Invoker {
	requestDef := GenReqDefForListGatewayResponsesV2()
	return &ListGatewayResponsesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceConfigsV2 查询租户实例配置列表
//
// 查询租户实例配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListInstanceConfigsV2(request *model.ListInstanceConfigsV2Request) (*model.ListInstanceConfigsV2Response, error) {
	requestDef := GenReqDefForListInstanceConfigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConfigsV2Response), nil
	}
}

// ListInstanceConfigsV2Invoker 查询租户实例配置列表
func (c *ApigClient) ListInstanceConfigsV2Invoker(request *model.ListInstanceConfigsV2Request) *ListInstanceConfigsV2Invoker {
	requestDef := GenReqDefForListInstanceConfigsV2()
	return &ListInstanceConfigsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceFeatures 查询实例支持的特性列表
//
// 查询实例支持的特性列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListInstanceFeatures(request *model.ListInstanceFeaturesRequest) (*model.ListInstanceFeaturesResponse, error) {
	requestDef := GenReqDefForListInstanceFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceFeaturesResponse), nil
	}
}

// ListInstanceFeaturesInvoker 查询实例支持的特性列表
func (c *ApigClient) ListInstanceFeaturesInvoker(request *model.ListInstanceFeaturesRequest) *ListInstanceFeaturesInvoker {
	requestDef := GenReqDefForListInstanceFeatures()
	return &ListInstanceFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询单个实例标签
//
// 查询单个实例的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询单个实例标签
func (c *ApigClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesByTags 通过标签查询实例列表
//
// 通过标签查询实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListInstancesByTags(request *model.ListInstancesByTagsRequest) (*model.ListInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagsResponse), nil
	}
}

// ListInstancesByTagsInvoker 通过标签查询实例列表
func (c *ApigClient) ListInstancesByTagsInvoker(request *model.ListInstancesByTagsRequest) *ListInstancesByTagsInvoker {
	requestDef := GenReqDefForListInstancesByTags()
	return &ListInstancesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesV2 查询专享版实例列表
//
// 查询专享版实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListInstancesV2(request *model.ListInstancesV2Request) (*model.ListInstancesV2Response, error) {
	requestDef := GenReqDefForListInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesV2Response), nil
	}
}

// ListInstancesV2Invoker 查询专享版实例列表
func (c *ApigClient) ListInstancesV2Invoker(request *model.ListInstancesV2Request) *ListInstancesV2Invoker {
	requestDef := GenReqDefForListInstancesV2()
	return &ListInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatelyApiStatisticsV2 API统计信息查询-最近一段时间
//
// 根据API的id和最近的一段时间查询API被调用的次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListLatelyApiStatisticsV2(request *model.ListLatelyApiStatisticsV2Request) (*model.ListLatelyApiStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyApiStatisticsV2Response), nil
	}
}

// ListLatelyApiStatisticsV2Invoker API统计信息查询-最近一段时间
func (c *ApigClient) ListLatelyApiStatisticsV2Invoker(request *model.ListLatelyApiStatisticsV2Request) *ListLatelyApiStatisticsV2Invoker {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()
	return &ListLatelyApiStatisticsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatelyGroupStatisticsV2 分组统计信息查询-最近一小时内
//
// 根据API分组的编号查询该分组下所有API被调用的总次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListLatelyGroupStatisticsV2(request *model.ListLatelyGroupStatisticsV2Request) (*model.ListLatelyGroupStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyGroupStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyGroupStatisticsV2Response), nil
	}
}

// ListLatelyGroupStatisticsV2Invoker 分组统计信息查询-最近一小时内
func (c *ApigClient) ListLatelyGroupStatisticsV2Invoker(request *model.ListLatelyGroupStatisticsV2Request) *ListLatelyGroupStatisticsV2Invoker {
	requestDef := GenReqDefForListLatelyGroupStatisticsV2()
	return &ListLatelyGroupStatisticsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricData 查询监控数据
//
// 查询指定时间范围指定指标的指定粒度的监控数据，可以通过参数指定需要查询的数据维度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListMetricData(request *model.ListMetricDataRequest) (*model.ListMetricDataResponse, error) {
	requestDef := GenReqDefForListMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricDataResponse), nil
	}
}

// ListMetricDataInvoker 查询监控数据
func (c *ApigClient) ListMetricDataInvoker(request *model.ListMetricDataRequest) *ListMetricDataInvoker {
	requestDef := GenReqDefForListMetricData()
	return &ListMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrchestrationAttachedApis 查询编排规则绑定的API
//
// 查询指定插件下绑定的API信息
// - 用于查询指定插件下已经绑定的API列表信息
// - 支持分页返回
// - 支持API名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListOrchestrationAttachedApis(request *model.ListOrchestrationAttachedApisRequest) (*model.ListOrchestrationAttachedApisResponse, error) {
	requestDef := GenReqDefForListOrchestrationAttachedApis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrchestrationAttachedApisResponse), nil
	}
}

// ListOrchestrationAttachedApisInvoker 查询编排规则绑定的API
func (c *ApigClient) ListOrchestrationAttachedApisInvoker(request *model.ListOrchestrationAttachedApisRequest) *ListOrchestrationAttachedApisInvoker {
	requestDef := GenReqDefForListOrchestrationAttachedApis()
	return &ListOrchestrationAttachedApisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrchestrations 查看编排规则列表
//
// 查看编排规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListOrchestrations(request *model.ListOrchestrationsRequest) (*model.ListOrchestrationsResponse, error) {
	requestDef := GenReqDefForListOrchestrations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrchestrationsResponse), nil
	}
}

// ListOrchestrationsInvoker 查看编排规则列表
func (c *ApigClient) ListOrchestrationsInvoker(request *model.ListOrchestrationsRequest) *ListOrchestrationsInvoker {
	requestDef := GenReqDefForListOrchestrations()
	return &ListOrchestrationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPluginAttachableApis 查询可绑定当前插件的API
//
// 查询可绑定当前插件的API信息。
// - 支持分页返回
// - 支持API名称模糊查询
// - 支持已绑定其他插件的API查询返回
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListPluginAttachableApis(request *model.ListPluginAttachableApisRequest) (*model.ListPluginAttachableApisResponse, error) {
	requestDef := GenReqDefForListPluginAttachableApis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginAttachableApisResponse), nil
	}
}

// ListPluginAttachableApisInvoker 查询可绑定当前插件的API
func (c *ApigClient) ListPluginAttachableApisInvoker(request *model.ListPluginAttachableApisRequest) *ListPluginAttachableApisInvoker {
	requestDef := GenReqDefForListPluginAttachableApis()
	return &ListPluginAttachableApisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPluginAttachedApis 查询插件下绑定的API
//
// 查询指定插件下绑定的API信息。
// - 用于查询指定插件下已经绑定的API列表信息
// - 支持分页返回
// - 支持API名称模糊查询
// - 绑定关系列表中返回的API在对应的环境中可能已经下线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListPluginAttachedApis(request *model.ListPluginAttachedApisRequest) (*model.ListPluginAttachedApisResponse, error) {
	requestDef := GenReqDefForListPluginAttachedApis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginAttachedApisResponse), nil
	}
}

// ListPluginAttachedApisInvoker 查询插件下绑定的API
func (c *ApigClient) ListPluginAttachedApisInvoker(request *model.ListPluginAttachedApisRequest) *ListPluginAttachedApisInvoker {
	requestDef := GenReqDefForListPluginAttachedApis()
	return &ListPluginAttachedApisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlugins 查询插件列表
//
// 查询一组符合条件的API网关插件详情。
// - 支持分页
// - 支持根据插件类型查询
// - 支持根据插件可见范围查询
// - 支持根据插件编码查询
// - 支持根据名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListPlugins(request *model.ListPluginsRequest) (*model.ListPluginsResponse, error) {
	requestDef := GenReqDefForListPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginsResponse), nil
	}
}

// ListPluginsInvoker 查询插件列表
func (c *ApigClient) ListPluginsInvoker(request *model.ListPluginsRequest) *ListPluginsInvoker {
	requestDef := GenReqDefForListPlugins()
	return &ListPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectCofigsV2 查询某个实例的租户配置列表
//
// 查询某个实例的租户配置列表，用户可以通过此接口查看各类型资源配置及使用情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListProjectCofigsV2(request *model.ListProjectCofigsV2Request) (*model.ListProjectCofigsV2Response, error) {
	requestDef := GenReqDefForListProjectCofigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectCofigsV2Response), nil
	}
}

// ListProjectCofigsV2Invoker 查询某个实例的租户配置列表
func (c *ApigClient) ListProjectCofigsV2Invoker(request *model.ListProjectCofigsV2Request) *ListProjectCofigsV2Invoker {
	requestDef := GenReqDefForListProjectCofigsV2()
	return &ListProjectCofigsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectInstanceTags 查询项目下所有实例标签
//
// 查询项目下所有实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListProjectInstanceTags(request *model.ListProjectInstanceTagsRequest) (*model.ListProjectInstanceTagsResponse, error) {
	requestDef := GenReqDefForListProjectInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectInstanceTagsResponse), nil
	}
}

// ListProjectInstanceTagsInvoker 查询项目下所有实例标签
func (c *ApigClient) ListProjectInstanceTagsInvoker(request *model.ListProjectInstanceTagsRequest) *ListProjectInstanceTagsInvoker {
	requestDef := GenReqDefForListProjectInstanceTags()
	return &ListProjectInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRequestThrottlingPolicyV2 查询流控策略列表
//
// 查询所有流控策略的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListRequestThrottlingPolicyV2(request *model.ListRequestThrottlingPolicyV2Request) (*model.ListRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPolicyV2Response), nil
	}
}

// ListRequestThrottlingPolicyV2Invoker 查询流控策略列表
func (c *ApigClient) ListRequestThrottlingPolicyV2Invoker(request *model.ListRequestThrottlingPolicyV2Request) *ListRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()
	return &ListRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSignatureKeysBindedToApiV2 查看API绑定的签名密钥列表
//
// 查询某个API绑定的签名密钥列表。每个API在每个环境上应该最多只会绑定一个签名密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListSignatureKeysBindedToApiV2(request *model.ListSignatureKeysBindedToApiV2Request) (*model.ListSignatureKeysBindedToApiV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysBindedToApiV2Response), nil
	}
}

// ListSignatureKeysBindedToApiV2Invoker 查看API绑定的签名密钥列表
func (c *ApigClient) ListSignatureKeysBindedToApiV2Invoker(request *model.ListSignatureKeysBindedToApiV2Request) *ListSignatureKeysBindedToApiV2Invoker {
	requestDef := GenReqDefForListSignatureKeysBindedToApiV2()
	return &ListSignatureKeysBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSignatureKeysV2 查询签名密钥列表
//
// 查询所有签名密钥的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListSignatureKeysV2(request *model.ListSignatureKeysV2Request) (*model.ListSignatureKeysV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysV2Response), nil
	}
}

// ListSignatureKeysV2Invoker 查询签名密钥列表
func (c *ApigClient) ListSignatureKeysV2Invoker(request *model.ListSignatureKeysV2Request) *ListSignatureKeysV2Invoker {
	requestDef := GenReqDefForListSignatureKeysV2()
	return &ListSignatureKeysV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecialThrottlingConfigurationsV2 查看特殊设置列表
//
// 查看给流控策略设置的特殊配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListSpecialThrottlingConfigurationsV2(request *model.ListSpecialThrottlingConfigurationsV2Request) (*model.ListSpecialThrottlingConfigurationsV2Response, error) {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecialThrottlingConfigurationsV2Response), nil
	}
}

// ListSpecialThrottlingConfigurationsV2Invoker 查看特殊设置列表
func (c *ApigClient) ListSpecialThrottlingConfigurationsV2Invoker(request *model.ListSpecialThrottlingConfigurationsV2Request) *ListSpecialThrottlingConfigurationsV2Invoker {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()
	return &ListSpecialThrottlingConfigurationsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsV2 查询标签列表
//
// 查询标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListTagsV2(request *model.ListTagsV2Request) (*model.ListTagsV2Response, error) {
	requestDef := GenReqDefForListTagsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsV2Response), nil
	}
}

// ListTagsV2Invoker 查询标签列表
func (c *ApigClient) ListTagsV2Invoker(request *model.ListTagsV2Request) *ListTagsV2Invoker {
	requestDef := GenReqDefForListTagsV2()
	return &ListTagsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveEipV2 实例解绑EIP
//
// 实例解绑EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) RemoveEipV2(request *model.RemoveEipV2Request) (*model.RemoveEipV2Response, error) {
	requestDef := GenReqDefForRemoveEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveEipV2Response), nil
	}
}

// RemoveEipV2Invoker 实例解绑EIP
func (c *ApigClient) RemoveEipV2Invoker(request *model.RemoveEipV2Request) *RemoveEipV2Invoker {
	requestDef := GenReqDefForRemoveEipV2()
	return &RemoveEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveEngressEipV2 关闭实例公网出口
//
// 关闭实例公网出口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) RemoveEngressEipV2(request *model.RemoveEngressEipV2Request) (*model.RemoveEngressEipV2Response, error) {
	requestDef := GenReqDefForRemoveEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveEngressEipV2Response), nil
	}
}

// RemoveEngressEipV2Invoker 关闭实例公网出口
func (c *ApigClient) RemoveEngressEipV2Invoker(request *model.RemoveEngressEipV2Request) *RemoveEngressEipV2Invoker {
	requestDef := GenReqDefForRemoveEngressEipV2()
	return &RemoveEngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveIngressEipV2 关闭实例公网入口
//
// 关闭实例公网入口，仅当实例为ELB类型时支持
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) RemoveIngressEipV2(request *model.RemoveIngressEipV2Request) (*model.RemoveIngressEipV2Response, error) {
	requestDef := GenReqDefForRemoveIngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveIngressEipV2Response), nil
	}
}

// RemoveIngressEipV2Invoker 关闭实例公网入口
func (c *ApigClient) RemoveIngressEipV2Invoker(request *model.RemoveIngressEipV2Request) *RemoveIngressEipV2Invoker {
	requestDef := GenReqDefForRemoveIngressEipV2()
	return &RemoveIngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResettingAppSecretV2 重置密钥
//
// 重置指定APP的密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ResettingAppSecretV2(request *model.ResettingAppSecretV2Request) (*model.ResettingAppSecretV2Response, error) {
	requestDef := GenReqDefForResettingAppSecretV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResettingAppSecretV2Response), nil
	}
}

// ResettingAppSecretV2Invoker 重置密钥
func (c *ApigClient) ResettingAppSecretV2Invoker(request *model.ResettingAppSecretV2Request) *ResettingAppSecretV2Invoker {
	requestDef := GenReqDefForResettingAppSecretV2()
	return &ResettingAppSecretV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppBoundAppQuota 查询凭据关联的凭据配额
//
// 查看指定凭据关联的凭据配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowAppBoundAppQuota(request *model.ShowAppBoundAppQuotaRequest) (*model.ShowAppBoundAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppBoundAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppBoundAppQuotaResponse), nil
	}
}

// ShowAppBoundAppQuotaInvoker 查询凭据关联的凭据配额
func (c *ApigClient) ShowAppBoundAppQuotaInvoker(request *model.ShowAppBoundAppQuotaRequest) *ShowAppBoundAppQuotaInvoker {
	requestDef := GenReqDefForShowAppBoundAppQuota()
	return &ShowAppBoundAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppQuota 获取凭据配额详情
//
// 获取凭据配额详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowAppQuota(request *model.ShowAppQuotaRequest) (*model.ShowAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppQuotaResponse), nil
	}
}

// ShowAppQuotaInvoker 获取凭据配额详情
func (c *ApigClient) ShowAppQuotaInvoker(request *model.ShowAppQuotaRequest) *ShowAppQuotaInvoker {
	requestDef := GenReqDefForShowAppQuota()
	return &ShowAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsyncTaskResult 获取异步任务结果
//
// 获取异步任务结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowAsyncTaskResult(request *model.ShowAsyncTaskResultRequest) (*model.ShowAsyncTaskResultResponse, error) {
	requestDef := GenReqDefForShowAsyncTaskResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAsyncTaskResultResponse), nil
	}
}

// ShowAsyncTaskResultInvoker 获取异步任务结果
func (c *ApigClient) ShowAsyncTaskResultInvoker(request *model.ShowAsyncTaskResultRequest) *ShowAsyncTaskResultInvoker {
	requestDef := GenReqDefForShowAsyncTaskResult()
	return &ShowAsyncTaskResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppAcl 查看APP的访问控制详情
//
// 查看APP的访问控制详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAppAcl(request *model.ShowDetailsOfAppAclRequest) (*model.ShowDetailsOfAppAclResponse, error) {
	requestDef := GenReqDefForShowDetailsOfAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppAclResponse), nil
	}
}

// ShowDetailsOfAppAclInvoker 查看APP的访问控制详情
func (c *ApigClient) ShowDetailsOfAppAclInvoker(request *model.ShowDetailsOfAppAclRequest) *ShowDetailsOfAppAclInvoker {
	requestDef := GenReqDefForShowDetailsOfAppAcl()
	return &ShowDetailsOfAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppCodeV2 查看APP Code详情
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAppCodeV2(request *model.ShowDetailsOfAppCodeV2Request) (*model.ShowDetailsOfAppCodeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppCodeV2Response), nil
	}
}

// ShowDetailsOfAppCodeV2Invoker 查看APP Code详情
func (c *ApigClient) ShowDetailsOfAppCodeV2Invoker(request *model.ShowDetailsOfAppCodeV2Request) *ShowDetailsOfAppCodeV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()
	return &ShowDetailsOfAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppV2 查看APP详情
//
// 查看指定APP的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAppV2(request *model.ShowDetailsOfAppV2Request) (*model.ShowDetailsOfAppV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppV2Response), nil
	}
}

// ShowDetailsOfAppV2Invoker 查看APP详情
func (c *ApigClient) ShowDetailsOfAppV2Invoker(request *model.ShowDetailsOfAppV2Request) *ShowDetailsOfAppV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAppV2()
	return &ShowDetailsOfAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfCustomAuthorizersV2 查看自定义认证详情
//
// 查看自定义认证详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfCustomAuthorizersV2(request *model.ShowDetailsOfCustomAuthorizersV2Request) (*model.ShowDetailsOfCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfCustomAuthorizersV2Response), nil
	}
}

// ShowDetailsOfCustomAuthorizersV2Invoker 查看自定义认证详情
func (c *ApigClient) ShowDetailsOfCustomAuthorizersV2Invoker(request *model.ShowDetailsOfCustomAuthorizersV2Request) *ShowDetailsOfCustomAuthorizersV2Invoker {
	requestDef := GenReqDefForShowDetailsOfCustomAuthorizersV2()
	return &ShowDetailsOfCustomAuthorizersV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfDomainNameCertificateV2 查看域名证书
//
// 查看域名下绑定的证书详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfDomainNameCertificateV2(request *model.ShowDetailsOfDomainNameCertificateV2Request) (*model.ShowDetailsOfDomainNameCertificateV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfDomainNameCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfDomainNameCertificateV2Response), nil
	}
}

// ShowDetailsOfDomainNameCertificateV2Invoker 查看域名证书
func (c *ApigClient) ShowDetailsOfDomainNameCertificateV2Invoker(request *model.ShowDetailsOfDomainNameCertificateV2Request) *ShowDetailsOfDomainNameCertificateV2Invoker {
	requestDef := GenReqDefForShowDetailsOfDomainNameCertificateV2()
	return &ShowDetailsOfDomainNameCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfEnvironmentVariableV2 查看变量详情
//
// 查看指定的环境变量的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfEnvironmentVariableV2(request *model.ShowDetailsOfEnvironmentVariableV2Request) (*model.ShowDetailsOfEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfEnvironmentVariableV2Response), nil
	}
}

// ShowDetailsOfEnvironmentVariableV2Invoker 查看变量详情
func (c *ApigClient) ShowDetailsOfEnvironmentVariableV2Invoker(request *model.ShowDetailsOfEnvironmentVariableV2Request) *ShowDetailsOfEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()
	return &ShowDetailsOfEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfGatewayResponseTypeV2 查看分组下指定错误类型的自定义响应
//
// 查看分组下指定错误类型的自定义响应
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfGatewayResponseTypeV2(request *model.ShowDetailsOfGatewayResponseTypeV2Request) (*model.ShowDetailsOfGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfGatewayResponseTypeV2Response), nil
	}
}

// ShowDetailsOfGatewayResponseTypeV2Invoker 查看分组下指定错误类型的自定义响应
func (c *ApigClient) ShowDetailsOfGatewayResponseTypeV2Invoker(request *model.ShowDetailsOfGatewayResponseTypeV2Request) *ShowDetailsOfGatewayResponseTypeV2Invoker {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseTypeV2()
	return &ShowDetailsOfGatewayResponseTypeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfGatewayResponseV2 查询分组自定义响应详情
//
// 查询分组自定义响应详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfGatewayResponseV2(request *model.ShowDetailsOfGatewayResponseV2Request) (*model.ShowDetailsOfGatewayResponseV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfGatewayResponseV2Response), nil
	}
}

// ShowDetailsOfGatewayResponseV2Invoker 查询分组自定义响应详情
func (c *ApigClient) ShowDetailsOfGatewayResponseV2Invoker(request *model.ShowDetailsOfGatewayResponseV2Request) *ShowDetailsOfGatewayResponseV2Invoker {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseV2()
	return &ShowDetailsOfGatewayResponseV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfInstanceProgressV2 查看专享版实例创建进度
//
// 查看专享版实例创建进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfInstanceProgressV2(request *model.ShowDetailsOfInstanceProgressV2Request) (*model.ShowDetailsOfInstanceProgressV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfInstanceProgressV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfInstanceProgressV2Response), nil
	}
}

// ShowDetailsOfInstanceProgressV2Invoker 查看专享版实例创建进度
func (c *ApigClient) ShowDetailsOfInstanceProgressV2Invoker(request *model.ShowDetailsOfInstanceProgressV2Request) *ShowDetailsOfInstanceProgressV2Invoker {
	requestDef := GenReqDefForShowDetailsOfInstanceProgressV2()
	return &ShowDetailsOfInstanceProgressV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfInstanceV2 查看专享版实例详情
//
// 查看专享版实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfInstanceV2(request *model.ShowDetailsOfInstanceV2Request) (*model.ShowDetailsOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfInstanceV2Response), nil
	}
}

// ShowDetailsOfInstanceV2Invoker 查看专享版实例详情
func (c *ApigClient) ShowDetailsOfInstanceV2Invoker(request *model.ShowDetailsOfInstanceV2Request) *ShowDetailsOfInstanceV2Invoker {
	requestDef := GenReqDefForShowDetailsOfInstanceV2()
	return &ShowDetailsOfInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfOrchestration 查询编排规则详情
//
// 查询编排规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfOrchestration(request *model.ShowDetailsOfOrchestrationRequest) (*model.ShowDetailsOfOrchestrationResponse, error) {
	requestDef := GenReqDefForShowDetailsOfOrchestration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfOrchestrationResponse), nil
	}
}

// ShowDetailsOfOrchestrationInvoker 查询编排规则详情
func (c *ApigClient) ShowDetailsOfOrchestrationInvoker(request *model.ShowDetailsOfOrchestrationRequest) *ShowDetailsOfOrchestrationInvoker {
	requestDef := GenReqDefForShowDetailsOfOrchestration()
	return &ShowDetailsOfOrchestrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfRequestThrottlingPolicyV2 查看流控策略详情
//
// 查看指定流控策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfRequestThrottlingPolicyV2(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfRequestThrottlingPolicyV2Response), nil
	}
}

// ShowDetailsOfRequestThrottlingPolicyV2Invoker 查看流控策略详情
func (c *ApigClient) ShowDetailsOfRequestThrottlingPolicyV2Invoker(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) *ShowDetailsOfRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()
	return &ShowDetailsOfRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstancesNumByTags 查询包含指定标签的实例数量
//
// 查询包含指定标签的实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowInstancesNumByTags(request *model.ShowInstancesNumByTagsRequest) (*model.ShowInstancesNumByTagsResponse, error) {
	requestDef := GenReqDefForShowInstancesNumByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstancesNumByTagsResponse), nil
	}
}

// ShowInstancesNumByTagsInvoker 查询包含指定标签的实例数量
func (c *ApigClient) ShowInstancesNumByTagsInvoker(request *model.ShowInstancesNumByTagsRequest) *ShowInstancesNumByTagsInvoker {
	requestDef := GenReqDefForShowInstancesNumByTags()
	return &ShowInstancesNumByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlugin 查询插件详情
//
// 查询插件详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowPlugin(request *model.ShowPluginRequest) (*model.ShowPluginResponse, error) {
	requestDef := GenReqDefForShowPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginResponse), nil
	}
}

// ShowPluginInvoker 查询插件详情
func (c *ApigClient) ShowPluginInvoker(request *model.ShowPluginRequest) *ShowPluginInvoker {
	requestDef := GenReqDefForShowPlugin()
	return &ShowPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRestrictionOfInstanceV2 查看实例约束信息
//
// 查看实例约束信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowRestrictionOfInstanceV2(request *model.ShowRestrictionOfInstanceV2Request) (*model.ShowRestrictionOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowRestrictionOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRestrictionOfInstanceV2Response), nil
	}
}

// ShowRestrictionOfInstanceV2Invoker 查看实例约束信息
func (c *ApigClient) ShowRestrictionOfInstanceV2Invoker(request *model.ShowRestrictionOfInstanceV2Request) *ShowRestrictionOfInstanceV2Invoker {
	requestDef := GenReqDefForShowRestrictionOfInstanceV2()
	return &ShowRestrictionOfInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppAcl 设置APP的访问控制
//
// 设置凭据的访问控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateAppAcl(request *model.UpdateAppAclRequest) (*model.UpdateAppAclResponse, error) {
	requestDef := GenReqDefForUpdateAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppAclResponse), nil
	}
}

// UpdateAppAclInvoker 设置APP的访问控制
func (c *ApigClient) UpdateAppAclInvoker(request *model.UpdateAppAclRequest) *UpdateAppAclInvoker {
	requestDef := GenReqDefForUpdateAppAcl()
	return &UpdateAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppQuota 修改凭据配额
//
// 修改凭据配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateAppQuota(request *model.UpdateAppQuotaRequest) (*model.UpdateAppQuotaResponse, error) {
	requestDef := GenReqDefForUpdateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppQuotaResponse), nil
	}
}

// UpdateAppQuotaInvoker 修改凭据配额
func (c *ApigClient) UpdateAppQuotaInvoker(request *model.UpdateAppQuotaRequest) *UpdateAppQuotaInvoker {
	requestDef := GenReqDefForUpdateAppQuota()
	return &UpdateAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppV2 修改APP
//
// 修改指定APP的信息。其中可修改的属性为：name、remark，当支持用户自定义key和secret的开关开启时，app_key和app_secret也支持修改，其它属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateAppV2(request *model.UpdateAppV2Request) (*model.UpdateAppV2Response, error) {
	requestDef := GenReqDefForUpdateAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppV2Response), nil
	}
}

// UpdateAppV2Invoker 修改APP
func (c *ApigClient) UpdateAppV2Invoker(request *model.UpdateAppV2Request) *UpdateAppV2Invoker {
	requestDef := GenReqDefForUpdateAppV2()
	return &UpdateAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomAuthorizerV2 修改自定义认证
//
// 修改自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateCustomAuthorizerV2(request *model.UpdateCustomAuthorizerV2Request) (*model.UpdateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomAuthorizerV2Response), nil
	}
}

// UpdateCustomAuthorizerV2Invoker 修改自定义认证
func (c *ApigClient) UpdateCustomAuthorizerV2Invoker(request *model.UpdateCustomAuthorizerV2Request) *UpdateCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()
	return &UpdateCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainV2 修改域名
//
// 修改绑定的域名所对应的配置信息。使用实例自定义入方向端口的特性时，注意开启/关闭客户端校验会对相同域名的不同端口同时生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateDomainV2(request *model.UpdateDomainV2Request) (*model.UpdateDomainV2Response, error) {
	requestDef := GenReqDefForUpdateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainV2Response), nil
	}
}

// UpdateDomainV2Invoker 修改域名
func (c *ApigClient) UpdateDomainV2Invoker(request *model.UpdateDomainV2Request) *UpdateDomainV2Invoker {
	requestDef := GenReqDefForUpdateDomainV2()
	return &UpdateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEngressEipV2 更新实例出公网带宽
//
// 更新实例出公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateEngressEipV2(request *model.UpdateEngressEipV2Request) (*model.UpdateEngressEipV2Response, error) {
	requestDef := GenReqDefForUpdateEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEngressEipV2Response), nil
	}
}

// UpdateEngressEipV2Invoker 更新实例出公网带宽
func (c *ApigClient) UpdateEngressEipV2Invoker(request *model.UpdateEngressEipV2Request) *UpdateEngressEipV2Invoker {
	requestDef := GenReqDefForUpdateEngressEipV2()
	return &UpdateEngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironmentV2 修改环境
//
// 修改指定环境的信息。其中可修改的属性为：name、remark，其它属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateEnvironmentV2(request *model.UpdateEnvironmentV2Request) (*model.UpdateEnvironmentV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentV2Response), nil
	}
}

// UpdateEnvironmentV2Invoker 修改环境
func (c *ApigClient) UpdateEnvironmentV2Invoker(request *model.UpdateEnvironmentV2Request) *UpdateEnvironmentV2Invoker {
	requestDef := GenReqDefForUpdateEnvironmentV2()
	return &UpdateEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironmentVariableV2 修改变量
//
// 修改环境变量。环境变量引用位置为api的后端服务地址时，修改对应环境变量会将使用该变量的所有api重新发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateEnvironmentVariableV2(request *model.UpdateEnvironmentVariableV2Request) (*model.UpdateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentVariableV2Response), nil
	}
}

// UpdateEnvironmentVariableV2Invoker 修改变量
func (c *ApigClient) UpdateEnvironmentVariableV2Invoker(request *model.UpdateEnvironmentVariableV2Request) *UpdateEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForUpdateEnvironmentVariableV2()
	return &UpdateEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGatewayResponseTypeV2 修改分组下指定错误类型的自定义响应
//
// 修改分组下指定错误类型的自定义响应。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateGatewayResponseTypeV2(request *model.UpdateGatewayResponseTypeV2Request) (*model.UpdateGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForUpdateGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGatewayResponseTypeV2Response), nil
	}
}

// UpdateGatewayResponseTypeV2Invoker 修改分组下指定错误类型的自定义响应
func (c *ApigClient) UpdateGatewayResponseTypeV2Invoker(request *model.UpdateGatewayResponseTypeV2Request) *UpdateGatewayResponseTypeV2Invoker {
	requestDef := GenReqDefForUpdateGatewayResponseTypeV2()
	return &UpdateGatewayResponseTypeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGatewayResponseV2 修改分组自定义响应
//
// 修改分组自定义响应
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateGatewayResponseV2(request *model.UpdateGatewayResponseV2Request) (*model.UpdateGatewayResponseV2Response, error) {
	requestDef := GenReqDefForUpdateGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGatewayResponseV2Response), nil
	}
}

// UpdateGatewayResponseV2Invoker 修改分组自定义响应
func (c *ApigClient) UpdateGatewayResponseV2Invoker(request *model.UpdateGatewayResponseV2Request) *UpdateGatewayResponseV2Invoker {
	requestDef := GenReqDefForUpdateGatewayResponseV2()
	return &UpdateGatewayResponseV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIngressEipV2 更新实例入公网带宽
//
// 更新实例入公网带宽，仅当实例为ELB类型时支持
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateIngressEipV2(request *model.UpdateIngressEipV2Request) (*model.UpdateIngressEipV2Response, error) {
	requestDef := GenReqDefForUpdateIngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIngressEipV2Response), nil
	}
}

// UpdateIngressEipV2Invoker 更新实例入公网带宽
func (c *ApigClient) UpdateIngressEipV2Invoker(request *model.UpdateIngressEipV2Request) *UpdateIngressEipV2Invoker {
	requestDef := GenReqDefForUpdateIngressEipV2()
	return &UpdateIngressEipV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceV2 更新专享版实例
//
// 更新专享版实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateInstanceV2(request *model.UpdateInstanceV2Request) (*model.UpdateInstanceV2Response, error) {
	requestDef := GenReqDefForUpdateInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceV2Response), nil
	}
}

// UpdateInstanceV2Invoker 更新专享版实例
func (c *ApigClient) UpdateInstanceV2Invoker(request *model.UpdateInstanceV2Request) *UpdateInstanceV2Invoker {
	requestDef := GenReqDefForUpdateInstanceV2()
	return &UpdateInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOrchestration 更新编排规则
//
// 更新编排规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateOrchestration(request *model.UpdateOrchestrationRequest) (*model.UpdateOrchestrationResponse, error) {
	requestDef := GenReqDefForUpdateOrchestration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOrchestrationResponse), nil
	}
}

// UpdateOrchestrationInvoker 更新编排规则
func (c *ApigClient) UpdateOrchestrationInvoker(request *model.UpdateOrchestrationRequest) *UpdateOrchestrationInvoker {
	requestDef := GenReqDefForUpdateOrchestration()
	return &UpdateOrchestrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlugin 修改插件
//
// 修改插件信息。
// - 插件不允许重名
// - 插件不支持修改类型和可见范围
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdatePlugin(request *model.UpdatePluginRequest) (*model.UpdatePluginResponse, error) {
	requestDef := GenReqDefForUpdatePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePluginResponse), nil
	}
}

// UpdatePluginInvoker 修改插件
func (c *ApigClient) UpdatePluginInvoker(request *model.UpdatePluginRequest) *UpdatePluginInvoker {
	requestDef := GenReqDefForUpdatePlugin()
	return &UpdatePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRequestThrottlingPolicyV2 修改流控策略
//
// 修改指定流控策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateRequestThrottlingPolicyV2(request *model.UpdateRequestThrottlingPolicyV2Request) (*model.UpdateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestThrottlingPolicyV2Response), nil
	}
}

// UpdateRequestThrottlingPolicyV2Invoker 修改流控策略
func (c *ApigClient) UpdateRequestThrottlingPolicyV2Invoker(request *model.UpdateRequestThrottlingPolicyV2Request) *UpdateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()
	return &UpdateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSignatureKeyV2 修改签名密钥
//
// 修改指定签名密钥的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateSignatureKeyV2(request *model.UpdateSignatureKeyV2Request) (*model.UpdateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForUpdateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSignatureKeyV2Response), nil
	}
}

// UpdateSignatureKeyV2Invoker 修改签名密钥
func (c *ApigClient) UpdateSignatureKeyV2Invoker(request *model.UpdateSignatureKeyV2Request) *UpdateSignatureKeyV2Invoker {
	requestDef := GenReqDefForUpdateSignatureKeyV2()
	return &UpdateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSlDomainSettingV2 设置调试域名是否可以访问
//
// 禁用或启用API分组绑定的调试域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateSlDomainSettingV2(request *model.UpdateSlDomainSettingV2Request) (*model.UpdateSlDomainSettingV2Response, error) {
	requestDef := GenReqDefForUpdateSlDomainSettingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSlDomainSettingV2Response), nil
	}
}

// UpdateSlDomainSettingV2Invoker 设置调试域名是否可以访问
func (c *ApigClient) UpdateSlDomainSettingV2Invoker(request *model.UpdateSlDomainSettingV2Request) *UpdateSlDomainSettingV2Invoker {
	requestDef := GenReqDefForUpdateSlDomainSettingV2()
	return &UpdateSlDomainSettingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSpecialThrottlingConfigurationV2 修改特殊设置
//
// 修改某个流控策略下的某个特殊设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateSpecialThrottlingConfigurationV2(request *model.UpdateSpecialThrottlingConfigurationV2Request) (*model.UpdateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpecialThrottlingConfigurationV2Response), nil
	}
}

// UpdateSpecialThrottlingConfigurationV2Invoker 修改特殊设置
func (c *ApigClient) UpdateSpecialThrottlingConfigurationV2Invoker(request *model.UpdateSpecialThrottlingConfigurationV2Request) *UpdateSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()
	return &UpdateSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAclV2 批量删除ACL策略
//
// 批量删除指定的多个ACL策略。
//
// 删除ACL策略时，如果存在ACL策略与API绑定关系，则无法删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDeleteAclV2(request *model.BatchDeleteAclV2Request) (*model.BatchDeleteAclV2Response, error) {
	requestDef := GenReqDefForBatchDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclV2Response), nil
	}
}

// BatchDeleteAclV2Invoker 批量删除ACL策略
func (c *ApigClient) BatchDeleteAclV2Invoker(request *model.BatchDeleteAclV2Request) *BatchDeleteAclV2Invoker {
	requestDef := GenReqDefForBatchDeleteAclV2()
	return &BatchDeleteAclV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAclStrategyV2 创建ACL策略
//
// 增加一个ACL策略，策略类型通过字段acl_type来确定（permit或者deny），限制的对象的类型可以为IP或者DOMAIN，这里的DOMAIN对应的acl_value的值为租户名称，而非“www.exampleDomain.com”之类的网络域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAclStrategyV2(request *model.CreateAclStrategyV2Request) (*model.CreateAclStrategyV2Response, error) {
	requestDef := GenReqDefForCreateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAclStrategyV2Response), nil
	}
}

// CreateAclStrategyV2Invoker 创建ACL策略
func (c *ApigClient) CreateAclStrategyV2Invoker(request *model.CreateAclStrategyV2Request) *CreateAclStrategyV2Invoker {
	requestDef := GenReqDefForCreateAclStrategyV2()
	return &CreateAclStrategyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclV2 删除ACL策略
//
// 删除指定的ACL策略， 如果存在api与该ACL策略的绑定关系，则无法删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteAclV2(request *model.DeleteAclV2Request) (*model.DeleteAclV2Response, error) {
	requestDef := GenReqDefForDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclV2Response), nil
	}
}

// DeleteAclV2Invoker 删除ACL策略
func (c *ApigClient) DeleteAclV2Invoker(request *model.DeleteAclV2Request) *DeleteAclV2Invoker {
	requestDef := GenReqDefForDeleteAclV2()
	return &DeleteAclV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclStrategiesV2 查看ACL策略列表
//
// 查询所有的ACL策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAclStrategiesV2(request *model.ListAclStrategiesV2Request) (*model.ListAclStrategiesV2Response, error) {
	requestDef := GenReqDefForListAclStrategiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclStrategiesV2Response), nil
	}
}

// ListAclStrategiesV2Invoker 查看ACL策略列表
func (c *ApigClient) ListAclStrategiesV2Invoker(request *model.ListAclStrategiesV2Request) *ListAclStrategiesV2Invoker {
	requestDef := GenReqDefForListAclStrategiesV2()
	return &ListAclStrategiesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAclPolicyV2 查看ACL策略详情
//
// 查询指定ACL策略的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAclPolicyV2(request *model.ShowDetailsOfAclPolicyV2Request) (*model.ShowDetailsOfAclPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAclPolicyV2Response), nil
	}
}

// ShowDetailsOfAclPolicyV2Invoker 查看ACL策略详情
func (c *ApigClient) ShowDetailsOfAclPolicyV2Invoker(request *model.ShowDetailsOfAclPolicyV2Request) *ShowDetailsOfAclPolicyV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAclPolicyV2()
	return &ShowDetailsOfAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclStrategyV2 修改ACL策略
//
// 修改指定的ACL策略，其中可修改的属性为：acl_name、acl_type、acl_value，其它属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateAclStrategyV2(request *model.UpdateAclStrategyV2Request) (*model.UpdateAclStrategyV2Response, error) {
	requestDef := GenReqDefForUpdateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclStrategyV2Response), nil
	}
}

// UpdateAclStrategyV2Invoker 修改ACL策略
func (c *ApigClient) UpdateAclStrategyV2Invoker(request *model.UpdateAclStrategyV2Request) *UpdateAclStrategyV2Invoker {
	requestDef := GenReqDefForUpdateAclStrategyV2()
	return &UpdateAclStrategyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRequestThrottlingPolicyV2 绑定流控策略
//
// 将流控策略应用于API，则所有对该API的访问将会受到该流控策略的限制。
//
// 当一定时间内的访问次数超过流控策略设置的API最大访问次数限制后，后续的访问将会被拒绝，从而能够较好的保护后端API免受异常流量的冲击，保障服务的稳定运行。
//
// 为指定的API绑定流控策略，绑定时，需要指定在哪个环境上生效。同一个API发布到不同的环境可以绑定不同的流控策略；一个API在发布到特定环境后只能绑定一个默认的流控策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AssociateRequestThrottlingPolicyV2(request *model.AssociateRequestThrottlingPolicyV2Request) (*model.AssociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForAssociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRequestThrottlingPolicyV2Response), nil
	}
}

// AssociateRequestThrottlingPolicyV2Invoker 绑定流控策略
func (c *ApigClient) AssociateRequestThrottlingPolicyV2Invoker(request *model.AssociateRequestThrottlingPolicyV2Request) *AssociateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForAssociateRequestThrottlingPolicyV2()
	return &AssociateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateThrottlingPolicyV2 批量解绑流控策略
//
// 批量解除API与流控策略的绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDisassociateThrottlingPolicyV2(request *model.BatchDisassociateThrottlingPolicyV2Request) (*model.BatchDisassociateThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForBatchDisassociateThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateThrottlingPolicyV2Response), nil
	}
}

// BatchDisassociateThrottlingPolicyV2Invoker 批量解绑流控策略
func (c *ApigClient) BatchDisassociateThrottlingPolicyV2Invoker(request *model.BatchDisassociateThrottlingPolicyV2Request) *BatchDisassociateThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForBatchDisassociateThrottlingPolicyV2()
	return &BatchDisassociateThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchPublishOrOfflineApiV2 批量发布或下线API
//
// 将多个API发布到一个指定的环境，或将多个API从指定的环境下线。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchPublishOrOfflineApiV2(request *model.BatchPublishOrOfflineApiV2Request) (*model.BatchPublishOrOfflineApiV2Response, error) {
	requestDef := GenReqDefForBatchPublishOrOfflineApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPublishOrOfflineApiV2Response), nil
	}
}

// BatchPublishOrOfflineApiV2Invoker 批量发布或下线API
func (c *ApigClient) BatchPublishOrOfflineApiV2Invoker(request *model.BatchPublishOrOfflineApiV2Request) *BatchPublishOrOfflineApiV2Invoker {
	requestDef := GenReqDefForBatchPublishOrOfflineApiV2()
	return &BatchPublishOrOfflineApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeApiVersionV2 切换API版本
//
// API每次发布时，会基于当前的API定义生成一个版本。版本记录了API发布时的各种定义及状态。
//
// 多个版本之间可以进行随意切换。但一个API在一个环境上，只能有一个版本生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ChangeApiVersionV2(request *model.ChangeApiVersionV2Request) (*model.ChangeApiVersionV2Response, error) {
	requestDef := GenReqDefForChangeApiVersionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApiVersionV2Response), nil
	}
}

// ChangeApiVersionV2Invoker 切换API版本
func (c *ApigClient) ChangeApiVersionV2Invoker(request *model.ChangeApiVersionV2Request) *ChangeApiVersionV2Invoker {
	requestDef := GenReqDefForChangeApiVersionV2()
	return &ChangeApiVersionV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckApiGroupsV2 校验API分组名称是否存在
//
// 校验API分组名称是否存在。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CheckApiGroupsV2(request *model.CheckApiGroupsV2Request) (*model.CheckApiGroupsV2Response, error) {
	requestDef := GenReqDefForCheckApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApiGroupsV2Response), nil
	}
}

// CheckApiGroupsV2Invoker 校验API分组名称是否存在
func (c *ApigClient) CheckApiGroupsV2Invoker(request *model.CheckApiGroupsV2Request) *CheckApiGroupsV2Invoker {
	requestDef := GenReqDefForCheckApiGroupsV2()
	return &CheckApiGroupsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckApisV2 校验API定义
//
// 校验API定义。校验API的路径或名称是否已存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CheckApisV2(request *model.CheckApisV2Request) (*model.CheckApisV2Response, error) {
	requestDef := GenReqDefForCheckApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApisV2Response), nil
	}
}

// CheckApisV2Invoker 校验API定义
func (c *ApigClient) CheckApisV2Invoker(request *model.CheckApisV2Request) *CheckApisV2Invoker {
	requestDef := GenReqDefForCheckApisV2()
	return &CheckApisV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiGroupV2 创建API分组
//
// API分组是API的管理单元，一个API分组等同于一个服务入口，创建API分组时，返回一个子域名作为访问入口。建议一个API分组下的API具有一定的相关性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateApiGroupV2(request *model.CreateApiGroupV2Request) (*model.CreateApiGroupV2Response, error) {
	requestDef := GenReqDefForCreateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiGroupV2Response), nil
	}
}

// CreateApiGroupV2Invoker 创建API分组
func (c *ApigClient) CreateApiGroupV2Invoker(request *model.CreateApiGroupV2Request) *CreateApiGroupV2Invoker {
	requestDef := GenReqDefForCreateApiGroupV2()
	return &CreateApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiV2 创建API
//
// 添加一个API，API即一个服务接口，具体的服务能力。
//
// API分为两部分，第一部分为面向API使用者的API接口，定义了使用者如何调用这个API。第二部分面向API提供者，由API提供者定义这个API的真实的后端情况，定义了API网关如何去访问真实的后端服务。API的真实后端服务目前支持四种类型：传统的HTTP/HTTPS形式的web后端、GRPC后端、函数工作流、MOCK。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateApiV2(request *model.CreateApiV2Request) (*model.CreateApiV2Response, error) {
	requestDef := GenReqDefForCreateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiV2Response), nil
	}
}

// CreateApiV2Invoker 创建API
func (c *ApigClient) CreateApiV2Invoker(request *model.CreateApiV2Request) *CreateApiV2Invoker {
	requestDef := GenReqDefForCreateApiV2()
	return &CreateApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrDeletePublishRecordForApiV2 发布或下线API
//
// 对API进行发布或下线。
//
// 发布操作是将一个指定的API发布到一个指定的环境，API只有发布后，才能够被调用，且只能在该环境上才能被调用。未发布的API无法被调用。
//
// 下线操作是将API从某个已发布的环境上下线，下线后，API将无法再被调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateOrDeletePublishRecordForApiV2(request *model.CreateOrDeletePublishRecordForApiV2Request) (*model.CreateOrDeletePublishRecordForApiV2Response, error) {
	requestDef := GenReqDefForCreateOrDeletePublishRecordForApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrDeletePublishRecordForApiV2Response), nil
	}
}

// CreateOrDeletePublishRecordForApiV2Invoker 发布或下线API
func (c *ApigClient) CreateOrDeletePublishRecordForApiV2Invoker(request *model.CreateOrDeletePublishRecordForApiV2Request) *CreateOrDeletePublishRecordForApiV2Invoker {
	requestDef := GenReqDefForCreateOrDeletePublishRecordForApiV2()
	return &CreateOrDeletePublishRecordForApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugApiV2 调试API
//
// 调试一个API在指定运行环境下的定义，接口调用者需要具有操作该API的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DebugApiV2(request *model.DebugApiV2Request) (*model.DebugApiV2Response, error) {
	requestDef := GenReqDefForDebugApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugApiV2Response), nil
	}
}

// DebugApiV2Invoker 调试API
func (c *ApigClient) DebugApiV2Invoker(request *model.DebugApiV2Request) *DebugApiV2Invoker {
	requestDef := GenReqDefForDebugApiV2()
	return &DebugApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiByVersionIdV2 根据版本编号下线API
//
// 对某个生效中的API版本进行下线操作，下线后，API在该版本生效的环境中将不再能够被调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteApiByVersionIdV2(request *model.DeleteApiByVersionIdV2Request) (*model.DeleteApiByVersionIdV2Response, error) {
	requestDef := GenReqDefForDeleteApiByVersionIdV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiByVersionIdV2Response), nil
	}
}

// DeleteApiByVersionIdV2Invoker 根据版本编号下线API
func (c *ApigClient) DeleteApiByVersionIdV2Invoker(request *model.DeleteApiByVersionIdV2Request) *DeleteApiByVersionIdV2Invoker {
	requestDef := GenReqDefForDeleteApiByVersionIdV2()
	return &DeleteApiByVersionIdV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiGroupV2 删除API分组
//
// 删除指定的API分组。
//
// 删除API分组前，要先下线并删除分组下的所有API。
//
// 删除时，会一并删除直接或间接关联到该分组下的所有资源，包括独立域名、SSL证书信息等等。并会将外部域名与子域名的绑定关系进行解除（取决于域名cname方式）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteApiGroupV2(request *model.DeleteApiGroupV2Request) (*model.DeleteApiGroupV2Response, error) {
	requestDef := GenReqDefForDeleteApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiGroupV2Response), nil
	}
}

// DeleteApiGroupV2Invoker 删除API分组
func (c *ApigClient) DeleteApiGroupV2Invoker(request *model.DeleteApiGroupV2Request) *DeleteApiGroupV2Invoker {
	requestDef := GenReqDefForDeleteApiGroupV2()
	return &DeleteApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiV2 删除API
//
// 删除指定的API。
//
// 删除API时，会删除该API所有相关的资源信息或绑定关系，如API的发布记录，绑定的后端服务，对APP的授权信息等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteApiV2(request *model.DeleteApiV2Request) (*model.DeleteApiV2Response, error) {
	requestDef := GenReqDefForDeleteApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiV2Response), nil
	}
}

// DeleteApiV2Invoker 删除API
func (c *ApigClient) DeleteApiV2Invoker(request *model.DeleteApiV2Request) *DeleteApiV2Invoker {
	requestDef := GenReqDefForDeleteApiV2()
	return &DeleteApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRequestThrottlingPolicyV2 解除API与流控策略的绑定关系
//
// 解除API与流控策略的绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DisassociateRequestThrottlingPolicyV2(request *model.DisassociateRequestThrottlingPolicyV2Request) (*model.DisassociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDisassociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRequestThrottlingPolicyV2Response), nil
	}
}

// DisassociateRequestThrottlingPolicyV2Invoker 解除API与流控策略的绑定关系
func (c *ApigClient) DisassociateRequestThrottlingPolicyV2Invoker(request *model.DisassociateRequestThrottlingPolicyV2Request) *DisassociateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForDisassociateRequestThrottlingPolicyV2()
	return &DisassociateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiGroupsV2 查询分组列表
//
// 查询API分组列表。
//
// 如果是租户操作，则查询该租户下所有的分组；如果是管理员权限账号操作，则查询的是所有租户的分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiGroupsV2(request *model.ListApiGroupsV2Request) (*model.ListApiGroupsV2Response, error) {
	requestDef := GenReqDefForListApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiGroupsV2Response), nil
	}
}

// ListApiGroupsV2Invoker 查询分组列表
func (c *ApigClient) ListApiGroupsV2Invoker(request *model.ListApiGroupsV2Request) *ListApiGroupsV2Invoker {
	requestDef := GenReqDefForListApiGroupsV2()
	return &ListApiGroupsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiRuntimeDefinitionV2 查询API运行时定义
//
// 查看指定的API在指定的环境上的运行时定义，默认查询RELEASE环境上的运行时定义。
//
// API的定义分为临时定义和运行时定义，分别代表如下含义：
// - 临时定义：API在编辑中的定义，表示用户最后一次编辑后的API的状态
// - 运行时定义：API在发布到某个环境时，对发布时的API的临时定义进行快照，固化出来的API的状态。
//
// 访问某个环境上的API，其实访问的就是其运行时的定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiRuntimeDefinitionV2(request *model.ListApiRuntimeDefinitionV2Request) (*model.ListApiRuntimeDefinitionV2Response, error) {
	requestDef := GenReqDefForListApiRuntimeDefinitionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiRuntimeDefinitionV2Response), nil
	}
}

// ListApiRuntimeDefinitionV2Invoker 查询API运行时定义
func (c *ApigClient) ListApiRuntimeDefinitionV2Invoker(request *model.ListApiRuntimeDefinitionV2Request) *ListApiRuntimeDefinitionV2Invoker {
	requestDef := GenReqDefForListApiRuntimeDefinitionV2()
	return &ListApiRuntimeDefinitionV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersionDetailV2 查看版本详情
//
// 查询某个指定的版本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiVersionDetailV2(request *model.ListApiVersionDetailV2Request) (*model.ListApiVersionDetailV2Response, error) {
	requestDef := GenReqDefForListApiVersionDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionDetailV2Response), nil
	}
}

// ListApiVersionDetailV2Invoker 查看版本详情
func (c *ApigClient) ListApiVersionDetailV2Invoker(request *model.ListApiVersionDetailV2Request) *ListApiVersionDetailV2Invoker {
	requestDef := GenReqDefForListApiVersionDetailV2()
	return &ListApiVersionDetailV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersionsV2 查询API历史版本列表
//
// 查询某个API的历史版本。每个API在一个环境上最多存在10个历史版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApiVersionsV2(request *model.ListApiVersionsV2Request) (*model.ListApiVersionsV2Response, error) {
	requestDef := GenReqDefForListApiVersionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsV2Response), nil
	}
}

// ListApiVersionsV2Invoker 查询API历史版本列表
func (c *ApigClient) ListApiVersionsV2Invoker(request *model.ListApiVersionsV2Request) *ListApiVersionsV2Invoker {
	requestDef := GenReqDefForListApiVersionsV2()
	return &ListApiVersionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToRequestThrottlingPolicyV2 查看流控策略绑定的API列表
//
// 查询某个流控策略上已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToRequestThrottlingPolicyV2(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) (*model.ListApisBindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToRequestThrottlingPolicyV2Response), nil
	}
}

// ListApisBindedToRequestThrottlingPolicyV2Invoker 查看流控策略绑定的API列表
func (c *ApigClient) ListApisBindedToRequestThrottlingPolicyV2Invoker(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) *ListApisBindedToRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListApisBindedToRequestThrottlingPolicyV2()
	return &ListApisBindedToRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToRequestThrottlingPolicyV2 查看流控策略未绑定的API列表
//
// 查询所有未绑定到该流控策略上的自有API列表。需要API已经发布，未发布的API不予展示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToRequestThrottlingPolicyV2(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToRequestThrottlingPolicyV2Response), nil
	}
}

// ListApisUnbindedToRequestThrottlingPolicyV2Invoker 查看流控策略未绑定的API列表
func (c *ApigClient) ListApisUnbindedToRequestThrottlingPolicyV2Invoker(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) *ListApisUnbindedToRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2()
	return &ListApisUnbindedToRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisV2 查询API列表
//
// 查看API列表，返回API详细信息、发布信息等，但不能查看到后端服务信息和API请求参数信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisV2(request *model.ListApisV2Request) (*model.ListApisV2Response, error) {
	requestDef := GenReqDefForListApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisV2Response), nil
	}
}

// ListApisV2Invoker 查询API列表
func (c *ApigClient) ListApisV2Invoker(request *model.ListApisV2Request) *ListApisV2Invoker {
	requestDef := GenReqDefForListApisV2()
	return &ListApisV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRequestThrottlingPoliciesBindedToApiV2 查看API绑定的流控策略列表
//
// 查询某个API绑定的流控策略列表。每个环境上应该最多只有一个流控策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListRequestThrottlingPoliciesBindedToApiV2(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPoliciesBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPoliciesBindedToApiV2Response), nil
	}
}

// ListRequestThrottlingPoliciesBindedToApiV2Invoker 查看API绑定的流控策略列表
func (c *ApigClient) ListRequestThrottlingPoliciesBindedToApiV2Invoker(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) *ListRequestThrottlingPoliciesBindedToApiV2Invoker {
	requestDef := GenReqDefForListRequestThrottlingPoliciesBindedToApiV2()
	return &ListRequestThrottlingPoliciesBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfApiGroupV2 查询分组详情
//
// 查询指定分组的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfApiGroupV2(request *model.ShowDetailsOfApiGroupV2Request) (*model.ShowDetailsOfApiGroupV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiGroupV2Response), nil
	}
}

// ShowDetailsOfApiGroupV2Invoker 查询分组详情
func (c *ApigClient) ShowDetailsOfApiGroupV2Invoker(request *model.ShowDetailsOfApiGroupV2Request) *ShowDetailsOfApiGroupV2Invoker {
	requestDef := GenReqDefForShowDetailsOfApiGroupV2()
	return &ShowDetailsOfApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfApiV2 查询API详情
//
// 查看指定的API的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfApiV2(request *model.ShowDetailsOfApiV2Request) (*model.ShowDetailsOfApiV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiV2Response), nil
	}
}

// ShowDetailsOfApiV2Invoker 查询API详情
func (c *ApigClient) ShowDetailsOfApiV2Invoker(request *model.ShowDetailsOfApiV2Request) *ShowDetailsOfApiV2Invoker {
	requestDef := GenReqDefForShowDetailsOfApiV2()
	return &ShowDetailsOfApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApiGroupV2 修改API分组
//
// 修改API分组属性。其中name和remark可修改，其他属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateApiGroupV2(request *model.UpdateApiGroupV2Request) (*model.UpdateApiGroupV2Response, error) {
	requestDef := GenReqDefForUpdateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiGroupV2Response), nil
	}
}

// UpdateApiGroupV2Invoker 修改API分组
func (c *ApigClient) UpdateApiGroupV2Invoker(request *model.UpdateApiGroupV2Request) *UpdateApiGroupV2Invoker {
	requestDef := GenReqDefForUpdateApiGroupV2()
	return &UpdateApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApiV2 修改API
//
// 修改指定API的信息，包括后端服务信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateApiV2(request *model.UpdateApiV2Request) (*model.UpdateApiV2Response, error) {
	requestDef := GenReqDefForUpdateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiV2Response), nil
	}
}

// UpdateApiV2Invoker 修改API
func (c *ApigClient) UpdateApiV2Invoker(request *model.UpdateApiV2Request) *UpdateApiV2Invoker {
	requestDef := GenReqDefForUpdateApiV2()
	return &UpdateApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteApiAclBindingV2 批量解除API与ACL策略的绑定
//
// 批量解除API与ACL策略的绑定
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDeleteApiAclBindingV2(request *model.BatchDeleteApiAclBindingV2Request) (*model.BatchDeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForBatchDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteApiAclBindingV2Response), nil
	}
}

// BatchDeleteApiAclBindingV2Invoker 批量解除API与ACL策略的绑定
func (c *ApigClient) BatchDeleteApiAclBindingV2Invoker(request *model.BatchDeleteApiAclBindingV2Request) *BatchDeleteApiAclBindingV2Invoker {
	requestDef := GenReqDefForBatchDeleteApiAclBindingV2()
	return &BatchDeleteApiAclBindingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiAclBindingV2 将API与ACL策略进行绑定
//
// 将API与ACL策略进行绑定。
//
// 同一个API发布到不同的环境可以绑定不同的ACL策略；一个API在发布到特定环境后只能绑定一个同一种类型的ACL策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateApiAclBindingV2(request *model.CreateApiAclBindingV2Request) (*model.CreateApiAclBindingV2Response, error) {
	requestDef := GenReqDefForCreateApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiAclBindingV2Response), nil
	}
}

// CreateApiAclBindingV2Invoker 将API与ACL策略进行绑定
func (c *ApigClient) CreateApiAclBindingV2Invoker(request *model.CreateApiAclBindingV2Request) *CreateApiAclBindingV2Invoker {
	requestDef := GenReqDefForCreateApiAclBindingV2()
	return &CreateApiAclBindingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiAclBindingV2 解除API与ACL策略的绑定
//
// 解除某条API与ACL策略的绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteApiAclBindingV2(request *model.DeleteApiAclBindingV2Request) (*model.DeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiAclBindingV2Response), nil
	}
}

// DeleteApiAclBindingV2Invoker 解除API与ACL策略的绑定
func (c *ApigClient) DeleteApiAclBindingV2Invoker(request *model.DeleteApiAclBindingV2Request) *DeleteApiAclBindingV2Invoker {
	requestDef := GenReqDefForDeleteApiAclBindingV2()
	return &DeleteApiAclBindingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclPolicyBindedToApiV2 查看API绑定的ACL策略列表
//
// 查看API绑定的ACL策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAclPolicyBindedToApiV2(request *model.ListAclPolicyBindedToApiV2Request) (*model.ListAclPolicyBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAclPolicyBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclPolicyBindedToApiV2Response), nil
	}
}

// ListAclPolicyBindedToApiV2Invoker 查看API绑定的ACL策略列表
func (c *ApigClient) ListAclPolicyBindedToApiV2Invoker(request *model.ListAclPolicyBindedToApiV2Request) *ListAclPolicyBindedToApiV2Invoker {
	requestDef := GenReqDefForListAclPolicyBindedToApiV2()
	return &ListAclPolicyBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToAclPolicyV2 查看ACL策略绑定的API列表
//
// 查看ACL策略绑定的API列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToAclPolicyV2(request *model.ListApisBindedToAclPolicyV2Request) (*model.ListApisBindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAclPolicyV2Response), nil
	}
}

// ListApisBindedToAclPolicyV2Invoker 查看ACL策略绑定的API列表
func (c *ApigClient) ListApisBindedToAclPolicyV2Invoker(request *model.ListApisBindedToAclPolicyV2Request) *ListApisBindedToAclPolicyV2Invoker {
	requestDef := GenReqDefForListApisBindedToAclPolicyV2()
	return &ListApisBindedToAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToAclPolicyV2 查看ACL策略未绑定的API列表
//
// 查看ACL策略未绑定的API列表，需要API已发布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToAclPolicyV2(request *model.ListApisUnbindedToAclPolicyV2Request) (*model.ListApisUnbindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAclPolicyV2Response), nil
	}
}

// ListApisUnbindedToAclPolicyV2Invoker 查看ACL策略未绑定的API列表
func (c *ApigClient) ListApisUnbindedToAclPolicyV2Invoker(request *model.ListApisUnbindedToAclPolicyV2Request) *ListApisUnbindedToAclPolicyV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToAclPolicyV2()
	return &ListApisUnbindedToAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelingAuthorizationV2 解除授权
//
// 解除API对APP的授权关系。解除授权后，APP将不再能够调用该API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CancelingAuthorizationV2(request *model.CancelingAuthorizationV2Request) (*model.CancelingAuthorizationV2Response, error) {
	requestDef := GenReqDefForCancelingAuthorizationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelingAuthorizationV2Response), nil
	}
}

// CancelingAuthorizationV2Invoker 解除授权
func (c *ApigClient) CancelingAuthorizationV2Invoker(request *model.CancelingAuthorizationV2Request) *CancelingAuthorizationV2Invoker {
	requestDef := GenReqDefForCancelingAuthorizationV2()
	return &CancelingAuthorizationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorizingAppsV2 APP授权
//
// APP创建成功后，还不能访问API，如果想要访问某个环境上的API，需要将该API在该环境上授权给APP。授权成功后，APP即可访问该环境上的这个API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateAuthorizingAppsV2(request *model.CreateAuthorizingAppsV2Request) (*model.CreateAuthorizingAppsV2Response, error) {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizingAppsV2Response), nil
	}
}

// CreateAuthorizingAppsV2Invoker APP授权
func (c *ApigClient) CreateAuthorizingAppsV2Invoker(request *model.CreateAuthorizingAppsV2Request) *CreateAuthorizingAppsV2Invoker {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()
	return &CreateAuthorizingAppsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToAppV2 查看APP已绑定的API列表
//
// 查询APP已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToAppV2(request *model.ListApisBindedToAppV2Request) (*model.ListApisBindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAppV2Response), nil
	}
}

// ListApisBindedToAppV2Invoker 查看APP已绑定的API列表
func (c *ApigClient) ListApisBindedToAppV2Invoker(request *model.ListApisBindedToAppV2Request) *ListApisBindedToAppV2Invoker {
	requestDef := GenReqDefForListApisBindedToAppV2()
	return &ListApisBindedToAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToAppV2 查看APP未绑定的API列表
//
// 查询指定环境上某个APP未绑定的API列表[，包括自有API和从云商店购买的API](tag:hws)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToAppV2(request *model.ListApisUnbindedToAppV2Request) (*model.ListApisUnbindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAppV2Response), nil
	}
}

// ListApisUnbindedToAppV2Invoker 查看APP未绑定的API列表
func (c *ApigClient) ListApisUnbindedToAppV2Invoker(request *model.ListApisUnbindedToAppV2Request) *ListApisUnbindedToAppV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToAppV2()
	return &ListApisUnbindedToAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppsBindedToApiV2 查看API已绑定的APP列表
//
// 查询API绑定的APP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAppsBindedToApiV2(request *model.ListAppsBindedToApiV2Request) (*model.ListAppsBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAppsBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsBindedToApiV2Response), nil
	}
}

// ListAppsBindedToApiV2Invoker 查看API已绑定的APP列表
func (c *ApigClient) ListAppsBindedToApiV2Invoker(request *model.ListAppsBindedToApiV2Request) *ListAppsBindedToApiV2Invoker {
	requestDef := GenReqDefForListAppsBindedToApiV2()
	return &ListAppsBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportApiDefinitionsV2 导出API
//
// 导出分组下API的定义信息。导出文件内容符合swagger标准规范，API网关自定义扩展字段请参考《API网关用户指南》的“导入导出API：扩展定义”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ExportApiDefinitionsV2(request *model.ExportApiDefinitionsV2Request) (*model.ExportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportApiDefinitionsV2Response), nil
	}
}

// ExportApiDefinitionsV2Invoker 导出API
func (c *ApigClient) ExportApiDefinitionsV2Invoker(request *model.ExportApiDefinitionsV2Request) *ExportApiDefinitionsV2Invoker {
	requestDef := GenReqDefForExportApiDefinitionsV2()
	return &ExportApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportApiDefinitionsV2 导入API
//
// 导入API。导入文件内容需要符合swagger标准规范，API网关自定义扩展字段请参考《API网关用户指南》的“导入导出API：扩展定义”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ImportApiDefinitionsV2(request *model.ImportApiDefinitionsV2Request) (*model.ImportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApiDefinitionsV2Response), nil
	}
}

// ImportApiDefinitionsV2Invoker 导入API
func (c *ApigClient) ImportApiDefinitionsV2Invoker(request *model.ImportApiDefinitionsV2Request) *ImportApiDefinitionsV2Invoker {
	requestDef := GenReqDefForImportApiDefinitionsV2()
	return &ImportApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateCertsV2 域名绑定SSL证书
//
// 域名绑定SSL证书。目前暂时仅支持单个绑定，请求体当中的certificate_ids里面有且只能有一个证书ID。使用实例自定义入方向端口的特性时，相同的域名会同时绑定证书，注意开启/关闭客户端校验会对相同域名的不同端口同时生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchAssociateCertsV2(request *model.BatchAssociateCertsV2Request) (*model.BatchAssociateCertsV2Response, error) {
	requestDef := GenReqDefForBatchAssociateCertsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateCertsV2Response), nil
	}
}

// BatchAssociateCertsV2Invoker 域名绑定SSL证书
func (c *ApigClient) BatchAssociateCertsV2Invoker(request *model.BatchAssociateCertsV2Request) *BatchAssociateCertsV2Invoker {
	requestDef := GenReqDefForBatchAssociateCertsV2()
	return &BatchAssociateCertsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateDomainsV2 SSL证书绑定域名
//
// SSL证书绑定域名。使用实例自定义入方向端口的特性时，相同的域名会同时绑定证书，注意开启/关闭客户端校验会对相同域名的不同端口同时生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchAssociateDomainsV2(request *model.BatchAssociateDomainsV2Request) (*model.BatchAssociateDomainsV2Response, error) {
	requestDef := GenReqDefForBatchAssociateDomainsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateDomainsV2Response), nil
	}
}

// BatchAssociateDomainsV2Invoker SSL证书绑定域名
func (c *ApigClient) BatchAssociateDomainsV2Invoker(request *model.BatchAssociateDomainsV2Request) *BatchAssociateDomainsV2Invoker {
	requestDef := GenReqDefForBatchAssociateDomainsV2()
	return &BatchAssociateDomainsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateCertsV2 域名解绑SSL证书
//
// 域名解绑SSL证书。目前暂时仅支持单个解绑，请求体当中的certificate_ids里面有且只能有一个证书ID。在使用自定义入方向端口的特性时，相同的域名会同时解绑证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDisassociateCertsV2(request *model.BatchDisassociateCertsV2Request) (*model.BatchDisassociateCertsV2Response, error) {
	requestDef := GenReqDefForBatchDisassociateCertsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateCertsV2Response), nil
	}
}

// BatchDisassociateCertsV2Invoker 域名解绑SSL证书
func (c *ApigClient) BatchDisassociateCertsV2Invoker(request *model.BatchDisassociateCertsV2Request) *BatchDisassociateCertsV2Invoker {
	requestDef := GenReqDefForBatchDisassociateCertsV2()
	return &BatchDisassociateCertsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateDomainsV2 SSL证书解绑域名
//
// SSL证书解绑域名。在使用自定义入方向端口的特性时，相同的域名会同时解绑证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDisassociateDomainsV2(request *model.BatchDisassociateDomainsV2Request) (*model.BatchDisassociateDomainsV2Response, error) {
	requestDef := GenReqDefForBatchDisassociateDomainsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateDomainsV2Response), nil
	}
}

// BatchDisassociateDomainsV2Invoker SSL证书解绑域名
func (c *ApigClient) BatchDisassociateDomainsV2Invoker(request *model.BatchDisassociateDomainsV2Request) *BatchDisassociateDomainsV2Invoker {
	requestDef := GenReqDefForBatchDisassociateDomainsV2()
	return &BatchDisassociateDomainsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificateV2 创建SSL证书
//
// 创建SSL证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateCertificateV2(request *model.CreateCertificateV2Request) (*model.CreateCertificateV2Response, error) {
	requestDef := GenReqDefForCreateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateV2Response), nil
	}
}

// CreateCertificateV2Invoker 创建SSL证书
func (c *ApigClient) CreateCertificateV2Invoker(request *model.CreateCertificateV2Request) *CreateCertificateV2Invoker {
	requestDef := GenReqDefForCreateCertificateV2()
	return &CreateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificateV2 删除SSL证书
//
// 删除ssl证书接口，删除时只有没有关联域名的证书才能被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteCertificateV2(request *model.DeleteCertificateV2Request) (*model.DeleteCertificateV2Response, error) {
	requestDef := GenReqDefForDeleteCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateV2Response), nil
	}
}

// DeleteCertificateV2Invoker 删除SSL证书
func (c *ApigClient) DeleteCertificateV2Invoker(request *model.DeleteCertificateV2Request) *DeleteCertificateV2Invoker {
	requestDef := GenReqDefForDeleteCertificateV2()
	return &DeleteCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachedDomainsV2 获取SSL证书已绑定域名列表
//
// 获取SSL证书已绑定域名列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListAttachedDomainsV2(request *model.ListAttachedDomainsV2Request) (*model.ListAttachedDomainsV2Response, error) {
	requestDef := GenReqDefForListAttachedDomainsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachedDomainsV2Response), nil
	}
}

// ListAttachedDomainsV2Invoker 获取SSL证书已绑定域名列表
func (c *ApigClient) ListAttachedDomainsV2Invoker(request *model.ListAttachedDomainsV2Request) *ListAttachedDomainsV2Invoker {
	requestDef := GenReqDefForListAttachedDomainsV2()
	return &ListAttachedDomainsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificatesV2 获取SSL证书列表
//
// 获取SSL证书列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListCertificatesV2(request *model.ListCertificatesV2Request) (*model.ListCertificatesV2Response, error) {
	requestDef := GenReqDefForListCertificatesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesV2Response), nil
	}
}

// ListCertificatesV2Invoker 获取SSL证书列表
func (c *ApigClient) ListCertificatesV2Invoker(request *model.ListCertificatesV2Request) *ListCertificatesV2Invoker {
	requestDef := GenReqDefForListCertificatesV2()
	return &ListCertificatesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfCertificateV2 查看证书详情
//
// 查看证书详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfCertificateV2(request *model.ShowDetailsOfCertificateV2Request) (*model.ShowDetailsOfCertificateV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfCertificateV2Response), nil
	}
}

// ShowDetailsOfCertificateV2Invoker 查看证书详情
func (c *ApigClient) ShowDetailsOfCertificateV2Invoker(request *model.ShowDetailsOfCertificateV2Request) *ShowDetailsOfCertificateV2Invoker {
	requestDef := GenReqDefForShowDetailsOfCertificateV2()
	return &ShowDetailsOfCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificateV2 修改SSL证书
//
// 修改SSL证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateCertificateV2(request *model.UpdateCertificateV2Request) (*model.UpdateCertificateV2Response, error) {
	requestDef := GenReqDefForUpdateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateV2Response), nil
	}
}

// UpdateCertificateV2Invoker 修改SSL证书
func (c *ApigClient) UpdateCertificateV2Invoker(request *model.UpdateCertificateV2Request) *UpdateCertificateV2Invoker {
	requestDef := GenReqDefForUpdateCertificateV2()
	return &UpdateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddingBackendInstancesV2 添加或更新后端实例
//
// 为指定的VPC通道添加后端实例
//
// 如果指定地址的后端实例已存在，则更新对应后端实例信息。如果请求体中包含多个重复地址的后端实例定义，则使用第一个定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) AddingBackendInstancesV2(request *model.AddingBackendInstancesV2Request) (*model.AddingBackendInstancesV2Response, error) {
	requestDef := GenReqDefForAddingBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddingBackendInstancesV2Response), nil
	}
}

// AddingBackendInstancesV2Invoker 添加或更新后端实例
func (c *ApigClient) AddingBackendInstancesV2Invoker(request *model.AddingBackendInstancesV2Request) *AddingBackendInstancesV2Invoker {
	requestDef := GenReqDefForAddingBackendInstancesV2()
	return &AddingBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisableMembers 批量修改后端服务器状态不可用
//
// 批量修改后端服务器状态不可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchDisableMembers(request *model.BatchDisableMembersRequest) (*model.BatchDisableMembersResponse, error) {
	requestDef := GenReqDefForBatchDisableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableMembersResponse), nil
	}
}

// BatchDisableMembersInvoker 批量修改后端服务器状态不可用
func (c *ApigClient) BatchDisableMembersInvoker(request *model.BatchDisableMembersRequest) *BatchDisableMembersInvoker {
	requestDef := GenReqDefForBatchDisableMembers()
	return &BatchDisableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableMembers 批量修改后端服务器状态可用
//
// 批量修改后端服务器状态可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) BatchEnableMembers(request *model.BatchEnableMembersRequest) (*model.BatchEnableMembersResponse, error) {
	requestDef := GenReqDefForBatchEnableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableMembersResponse), nil
	}
}

// BatchEnableMembersInvoker 批量修改后端服务器状态可用
func (c *ApigClient) BatchEnableMembersInvoker(request *model.BatchEnableMembersRequest) *BatchEnableMembersInvoker {
	requestDef := GenReqDefForBatchEnableMembers()
	return &BatchEnableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMemberGroup 添加或更新VPC通道后端服务器组
//
// 在APIG中创建VPC通道后端服务器组，VPC通道后端实例可以选择是否关联后端实例服务器组，以便管理后端服务器节点。
//
// 如果指定名称的后端服务器组已存在，则更新对应后端服务器组信息。如果请求体中包含多个重复名称的后端服务器定义，则使用第一个定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateMemberGroup(request *model.CreateMemberGroupRequest) (*model.CreateMemberGroupResponse, error) {
	requestDef := GenReqDefForCreateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberGroupResponse), nil
	}
}

// CreateMemberGroupInvoker 添加或更新VPC通道后端服务器组
func (c *ApigClient) CreateMemberGroupInvoker(request *model.CreateMemberGroupRequest) *CreateMemberGroupInvoker {
	requestDef := GenReqDefForCreateMemberGroup()
	return &CreateMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcChannelV2 创建VPC通道
//
// 在API网关中创建连接私有VPC资源的通道，并在创建API时将后端节点配置为使用这些VPC通道，以便API网关直接访问私有VPC资源。
// &gt; 每个用户最多创建30个VPC通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) CreateVpcChannelV2(request *model.CreateVpcChannelV2Request) (*model.CreateVpcChannelV2Response, error) {
	requestDef := GenReqDefForCreateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcChannelV2Response), nil
	}
}

// CreateVpcChannelV2Invoker 创建VPC通道
func (c *ApigClient) CreateVpcChannelV2Invoker(request *model.CreateVpcChannelV2Request) *CreateVpcChannelV2Invoker {
	requestDef := GenReqDefForCreateVpcChannelV2()
	return &CreateVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackendInstanceV2 删除后端实例
//
// 删除指定VPC通道中的后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteBackendInstanceV2(request *model.DeleteBackendInstanceV2Request) (*model.DeleteBackendInstanceV2Response, error) {
	requestDef := GenReqDefForDeleteBackendInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackendInstanceV2Response), nil
	}
}

// DeleteBackendInstanceV2Invoker 删除后端实例
func (c *ApigClient) DeleteBackendInstanceV2Invoker(request *model.DeleteBackendInstanceV2Request) *DeleteBackendInstanceV2Invoker {
	requestDef := GenReqDefForDeleteBackendInstanceV2()
	return &DeleteBackendInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMemberGroup 删除VPC通道后端服务器组
//
// 删除指定的VPC通道后端服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteMemberGroup(request *model.DeleteMemberGroupRequest) (*model.DeleteMemberGroupResponse, error) {
	requestDef := GenReqDefForDeleteMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberGroupResponse), nil
	}
}

// DeleteMemberGroupInvoker 删除VPC通道后端服务器组
func (c *ApigClient) DeleteMemberGroupInvoker(request *model.DeleteMemberGroupRequest) *DeleteMemberGroupInvoker {
	requestDef := GenReqDefForDeleteMemberGroup()
	return &DeleteMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcChannelV2 删除VPC通道
//
// 删除指定的VPC通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) DeleteVpcChannelV2(request *model.DeleteVpcChannelV2Request) (*model.DeleteVpcChannelV2Response, error) {
	requestDef := GenReqDefForDeleteVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcChannelV2Response), nil
	}
}

// DeleteVpcChannelV2Invoker 删除VPC通道
func (c *ApigClient) DeleteVpcChannelV2Invoker(request *model.DeleteVpcChannelV2Request) *DeleteVpcChannelV2Invoker {
	requestDef := GenReqDefForDeleteVpcChannelV2()
	return &DeleteVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackendInstancesV2 查看后端实例列表
//
// 查看指定VPC通道的后端实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListBackendInstancesV2(request *model.ListBackendInstancesV2Request) (*model.ListBackendInstancesV2Response, error) {
	requestDef := GenReqDefForListBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackendInstancesV2Response), nil
	}
}

// ListBackendInstancesV2Invoker 查看后端实例列表
func (c *ApigClient) ListBackendInstancesV2Invoker(request *model.ListBackendInstancesV2Request) *ListBackendInstancesV2Invoker {
	requestDef := GenReqDefForListBackendInstancesV2()
	return &ListBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMemberGroups 查询VPC通道后端云服务组列表
//
// 查询VPC通道后端云服务组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListMemberGroups(request *model.ListMemberGroupsRequest) (*model.ListMemberGroupsResponse, error) {
	requestDef := GenReqDefForListMemberGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMemberGroupsResponse), nil
	}
}

// ListMemberGroupsInvoker 查询VPC通道后端云服务组列表
func (c *ApigClient) ListMemberGroupsInvoker(request *model.ListMemberGroupsRequest) *ListMemberGroupsInvoker {
	requestDef := GenReqDefForListMemberGroups()
	return &ListMemberGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcChannelsV2 查询VPC通道列表
//
// 查看VPC通道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ListVpcChannelsV2(request *model.ListVpcChannelsV2Request) (*model.ListVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcChannelsV2Response), nil
	}
}

// ListVpcChannelsV2Invoker 查询VPC通道列表
func (c *ApigClient) ListVpcChannelsV2Invoker(request *model.ListVpcChannelsV2Request) *ListVpcChannelsV2Invoker {
	requestDef := GenReqDefForListVpcChannelsV2()
	return &ListVpcChannelsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfMemberGroup 查看VPC通道后端服务器组详情
//
// 查看指定的VPC通道后端服务器组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfMemberGroup(request *model.ShowDetailsOfMemberGroupRequest) (*model.ShowDetailsOfMemberGroupResponse, error) {
	requestDef := GenReqDefForShowDetailsOfMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfMemberGroupResponse), nil
	}
}

// ShowDetailsOfMemberGroupInvoker 查看VPC通道后端服务器组详情
func (c *ApigClient) ShowDetailsOfMemberGroupInvoker(request *model.ShowDetailsOfMemberGroupRequest) *ShowDetailsOfMemberGroupInvoker {
	requestDef := GenReqDefForShowDetailsOfMemberGroup()
	return &ShowDetailsOfMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfVpcChannelV2 查看VPC通道详情
//
// 查看指定的VPC通道详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfVpcChannelV2(request *model.ShowDetailsOfVpcChannelV2Request) (*model.ShowDetailsOfVpcChannelV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfVpcChannelV2Response), nil
	}
}

// ShowDetailsOfVpcChannelV2Invoker 查看VPC通道详情
func (c *ApigClient) ShowDetailsOfVpcChannelV2Invoker(request *model.ShowDetailsOfVpcChannelV2Request) *ShowDetailsOfVpcChannelV2Invoker {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()
	return &ShowDetailsOfVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBackendInstancesV2 更新后端实例
//
// 更新指定的VPC通道的后端实例。更新时，使用传入的请求参数对对应云服务组的后端实例进行全量覆盖修改。如果未指定修改的云服务器组，则进行全量覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateBackendInstancesV2(request *model.UpdateBackendInstancesV2Request) (*model.UpdateBackendInstancesV2Response, error) {
	requestDef := GenReqDefForUpdateBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackendInstancesV2Response), nil
	}
}

// UpdateBackendInstancesV2Invoker 更新后端实例
func (c *ApigClient) UpdateBackendInstancesV2Invoker(request *model.UpdateBackendInstancesV2Request) *UpdateBackendInstancesV2Invoker {
	requestDef := GenReqDefForUpdateBackendInstancesV2()
	return &UpdateBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthCheck 修改VPC通道健康检查
//
// 修改VPC通道健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateHealthCheck(request *model.UpdateHealthCheckRequest) (*model.UpdateHealthCheckResponse, error) {
	requestDef := GenReqDefForUpdateHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthCheckResponse), nil
	}
}

// UpdateHealthCheckInvoker 修改VPC通道健康检查
func (c *ApigClient) UpdateHealthCheckInvoker(request *model.UpdateHealthCheckRequest) *UpdateHealthCheckInvoker {
	requestDef := GenReqDefForUpdateHealthCheck()
	return &UpdateHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMemberGroup 更新VPC通道后端服务器组
//
// 更新指定VPC通道后端服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateMemberGroup(request *model.UpdateMemberGroupRequest) (*model.UpdateMemberGroupResponse, error) {
	requestDef := GenReqDefForUpdateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberGroupResponse), nil
	}
}

// UpdateMemberGroupInvoker 更新VPC通道后端服务器组
func (c *ApigClient) UpdateMemberGroupInvoker(request *model.UpdateMemberGroupRequest) *UpdateMemberGroupInvoker {
	requestDef := GenReqDefForUpdateMemberGroup()
	return &UpdateMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpcChannelV2 更新VPC通道
//
// 更新指定VPC通道的参数
//
// 使用传入的后端实例列表对VPC通道进行全量覆盖，如果后端实例列表为空，则会全量删除已有的后端实例；
//
// 使用传入的后端服务器组列表对VPC通道进行全量覆盖，如果后端服务器组列表为空，则会全量删除已有的服务器组；
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ApigClient) UpdateVpcChannelV2(request *model.UpdateVpcChannelV2Request) (*model.UpdateVpcChannelV2Response, error) {
	requestDef := GenReqDefForUpdateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcChannelV2Response), nil
	}
}

// UpdateVpcChannelV2Invoker 更新VPC通道
func (c *ApigClient) UpdateVpcChannelV2Invoker(request *model.UpdateVpcChannelV2Request) *UpdateVpcChannelV2Invoker {
	requestDef := GenReqDefForUpdateVpcChannelV2()
	return &UpdateVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
