package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2/model"
)

type ApigClient struct {
	HcClient *http_client.HcHttpClient
}

func NewApigClient(hcClient *http_client.HcHttpClient) *ApigClient {
	return &ApigClient{HcClient: hcClient}
}

func ApigClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 实例更新或绑定EIP
//
// 实例更新或绑定EIP
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AddEipV2(request *model.AddEipV2Request) (*model.AddEipV2Response, error) {
	requestDef := GenReqDefForAddEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEipV2Response), nil
	}
}

// 开启实例公网出口
//
// 实例开启公网出口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AddEngressEipV2(request *model.AddEngressEipV2Request) (*model.AddEngressEipV2Response, error) {
	requestDef := GenReqDefForAddEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEngressEipV2Response), nil
	}
}

// 绑定域名证书
//
// 如果创建API时，“定义API请求”使用HTTPS请求协议，那么在独立域名中需要添加SSL证书。
// 本章节主要介绍为特定域名绑定证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AssociateCertificateV2(request *model.AssociateCertificateV2Request) (*model.AssociateCertificateV2Response, error) {
	requestDef := GenReqDefForAssociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateCertificateV2Response), nil
	}
}

// 绑定域名
//
// 用户自定义的域名，需要CNAME到API分组的子域名上才能生效，具体方法请参见《云解析服务用户指南》的“添加CANME类型记录集”章节。
// 每个API分组下最多可绑定5个域名。绑定域名后，用户可通过自定义域名调用API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AssociateDomainV2(request *model.AssociateDomainV2Request) (*model.AssociateDomainV2Response, error) {
	requestDef := GenReqDefForAssociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateDomainV2Response), nil
	}
}

// 绑定签名密钥
//
// 签名密钥创建后，需要绑定到API才能生效。
//
// 将签名密钥绑定到API后，则API网关请求后端服务时就会使用这个签名密钥进行加密签名，后端服务可以校验这个签名来验证请求来源。
//
// 将指定的签名密钥绑定到一个或多个已发布的API上。同一个API发布到不同的环境可以绑定不同的签名密钥；一个API在发布到特定环境后只能绑定一个签名密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AssociateSignatureKeyV2(request *model.AssociateSignatureKeyV2Request) (*model.AssociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForAssociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSignatureKeyV2Response), nil
	}
}

// 创建自定义认证
//
// 创建自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateCustomAuthorizerV2(request *model.CreateCustomAuthorizerV2Request) (*model.CreateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomAuthorizerV2Response), nil
	}
}

// 创建环境
//
// 在实际的生产中，API提供者可能有多个环境，如开发环境、测试环境、生产环境等，用户可以自由将API发布到某个环境，供调用者调用。
//
// 对于不同的环境，API的版本、请求地址甚至于包括请求消息等均有可能不同。如：某个API，v1.0的版本为稳定版本，发布到了生产环境供生产使用，同时，该API正处于迭代中，v1.1的版本是开发人员交付测试人员进行测试的版本，发布在测试环境上，而v1.2的版本目前开发团队正处于开发过程中，可以发布到开发环境进行自测等。
//
// 为此，API网关提供多环境管理功能，使租户能够最大化的模拟实际场景，低成本的接入API网关。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateEnvironmentV2(request *model.CreateEnvironmentV2Request) (*model.CreateEnvironmentV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentV2Response), nil
	}
}

// 新建变量
//
// 将API发布到不同的环境后，对于不同的环境，可能会有不同的环境变量，比如，API的服务部署地址，请求的版本号等。
//
// 用户可以定义不同的环境变量，用户在定义API时，在API的定义中使用这些变量，当调用API时，API网关会将这些变量替换成真实的变量值，以达到不同环境的区分效果。
//
// 环境变量定义在API分组上，该分组下的所有API都可以使用这些变量。
// &gt; 1.环境变量的变量名称必须保持唯一，即一个分组在同一个环境上不能有两个同名的变量
//   2.环境变量区分大小写，即变量ABC与变量abc是两个不同的变量
//   3.设置了环境变量后，使用到该变量的API的调试功能将不可使用。
//   4.定义了环境变量后，使用到环境变量的地方应该以对称的#标识环境变量，当API发布到相应的环境后，会对环境变量的值进行替换，如：定义的API的URL为：https://#address#:8080，环境变量address在RELEASE环境上的值为：192.168.1.5，则API发布到RELEASE环境后的真实的URL为：https://192.168.1.5:8080。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateEnvironmentVariableV2(request *model.CreateEnvironmentVariableV2Request) (*model.CreateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentVariableV2Response), nil
	}
}

// 实例配置特性
//
// 为实例配置需要的特性。
//
// 支持配置的特性列表及特性配置示例如下：
//
// | 特性名称 | 特性描述 | 特性配置示例 | 特性配置参数 |  |  |  |
// --------| :------- | :-------| :-------| :-------| :-------| :-------
// |  |  |  | 参数名称 | 参数描述 | 参数默认值 | 参数范围 |
// | lts | 是否支持shubao访问日志上报功能。|{\&quot;name\&quot;:\&quot;lts\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;group_id\\\\\&quot;: \\\&quot;\\,\\\\\&quot;topic_id\\\\\&quot;:\\\\\&quot;\\\\\&quot;,\\\\\&quot;log_group\\\\\&quot;:\\\\\&quot;\\\\\&quot;,\\\\\&quot;log_stream\\\\\&quot;:\\\\\&quot;\\\\\&quot;}\&quot;} | group_id | 日志组ID | | |
// |  |  |  | topic_id | 日志流ID | | |
// |  |  |  | log_group | 日志组名称 | | |
// |  |  |  | log_stream | 日志流名称 | | |
// | ratelimit | 是否支持自定义流控值。|{\&quot;name\&quot;:\&quot;ratelimit\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;api_limits\\\\\&quot;: 500}\&quot;} | api_limits | API全局默认流控值。注意：如果配置过小会导致业务持续被流控，请根据业务谨慎修改。 | 200 次/秒 | 1-1000000 次/秒 |
// | request_body_size | 是否支持指定最大请求Body大小。|{\&quot;name\&quot;:\&quot;request_body_size\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;104857600\&quot;} | request_body_size | 请求中允许携带的Body大小上限。 | 12 M | 1-9536 M |
// | backend_timeout | 是否支持配置后端API最大超时时间。|{\&quot;name\&quot;:\&quot;backend_timeout\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\&quot;max_timeout\\\&quot;: 500}\&quot;} | max_timeout | API网关到后端服务的超时时间上限。 | 60000 ms | 1-600000 ms |
// | app_token | 是否开启app_token认证方式。|{\&quot;name\&quot;:\&quot;app_token\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;enable\\\\\&quot;: \\\\\&quot;on\\\\\&quot;, \\\\\&quot;app_token_expire_time\\\\\&quot;: 3600, \\\\\&quot;app_token_uri\\\\\&quot;: \\\\\&quot;/v1/apigw/oauth2/token\\\\\&quot;, \\\\\&quot;refresh_token_expire_time\\\\\&quot;: 7200}\&quot;} | enable | 是否开启 | off | on/off |
// |  |  |  | app_token_expire_time | access token的有效时间 | 3600 s | 1-72000 s |
// |  |  |  | refresh_token_expire_time | refresh token的有效时间 | 7200 s | 1-72000 s |
// |  |  |  | app_token_uri | 获取token的uri | /v1/apigw/oauth2/token |  |
// |  |  |  | app_token_key | token的加密key |  |  |
// | app_api_key | 是否开启app_api_key认证方式。|{\&quot;name\&quot;:\&quot;app_api_key\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;on\&quot;} |  |  | off | on/off |
// | app_basic | 是否开启app_basic认证方式。|{\&quot;name\&quot;:\&quot;app_basic\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;on\&quot;} |  |  | off | on/off |
// | app_secret | 是否支持app_secret认证方式。|{\&quot;name\&quot;:\&quot;app_secret\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;on\&quot;} |  |  | off | on/off |
// | app_jwt | 是否支持app_jwt认证方式。|{\&quot;name\&quot;:\&quot;app_jwt\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;enable\\\\\&quot;: \\\\\&quot;on\\\\\&quot;, \\\\\&quot;auth_header\\\\\&quot;: \\\\\&quot;Authorization\\\\\&quot;}\&quot;}| enable | 是否开启app_jwt认证方式。 | off | on/off |
// |  |  |  | auth_header | app_jwt认证头 | Authorization |  |
// | public_key | 是否支持public_key类型的后端签名。|{\&quot;name\&quot;:\&quot;public_key\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;enable\\\\\&quot;: \\\\\&quot;on\\\\\&quot;, \\\\\&quot;public_key_uri_prefix\\\\\&quot;: \\\\\&quot;/apigw/authadv/v2/public-key/\\\\\&quot;}\&quot;}| enable | 是否开启app_jwt认证方式。 | off | on/off |
// |  |  |  | public_key_uri_prefix | 获取public key的uri前缀 | /apigw/authadv/v2/public-key/ |  |
// | backend_token_allow | 是否支持普通租户透传token到后端。|{\&quot;name\&quot;:\&quot;backend_token_allow\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;backend_token_allow_users\\\\\&quot;: [\\\\\&quot;paas_apig_wwx548366_01\\\\\&quot;]}\&quot;} | backend_token_allow_users | 透传token到后端普通租户白名单，匹配普通租户domain name正则表达式 |  |  |
// | backend_client_certificate | 是否开启后端双向认证。|{\&quot;name\&quot;:\&quot;backend_client_certificate\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;{\\\\\&quot;enable\\\\\&quot;: \\\\\&quot;on\\\\\&quot;,\\\\\&quot;ca\\\\\&quot;: \\\\\&quot;\\\\\&quot;,\\\\\&quot;content\\\\\&quot;: \\\\\&quot;\\\\\&quot;,\\\\\&quot;key\\\\\&quot;: \\\\\&quot;\\\\\&quot;}\&quot;} | enable | 是否开启 | off | on/off |
// |  |  |  | ca | 双向认证信任证书 |  |  |
// |  |  |  | content | 双向认证证书 |  |  |
// |  |  |  | key | 双向认证信任私钥 |  |  |
// | ssl_ciphers | 是否支持https加密套件。|{\&quot;name\&quot;:\&quot;ssl_ciphers\&quot;,\&quot;enable\&quot;:true,\&quot;config\&quot;: \&quot;config\&quot;: \&quot;{\\\\\&quot;ssl_ciphers\\\\\&quot;: [\\\\\&quot;ECDHE-ECDSA-AES256-GCM-SHA384\\\\\&quot;]}\&quot;} | ssl_ciphers | 支持的加解密套件。ssl_ciphers数组中只允许出现默认值中的字符串，且数组不能为空。 |  | ECDHE-ECDSA-AES256-GCM-SHA384,ECDHE-RSA-AES256-GCM-SHA384,ECDHE-ECDSA-AES128-GCM-SHA256,ECDHE-RSA-AES128-GCM-SHA256,ECDHE-ECDSA-AES256-SHA384,ECDHE-RSA-AES256-SHA384,ECDHE-ECDSA-AES128-SHA256,ECDHE-RSA-AES128-SHA256 |
// | real_ip_from_xff | 是否开启使用xff头作为访问控制、流控策略的源ip生效依据。|{\&quot;name\&quot;:\&quot;real_ip_from_xff\&quot;,\&quot;enable\&quot;: true,\&quot;config\&quot;: \&quot;{\\\\\&quot;enable\\\\\&quot;: \\\\\&quot;on\\\\\&quot;,\\\\\&quot;xff_index\\\\\&quot;: 1}\&quot;} | enable | 是否开启 | off | on/off |
// |  |  |  | xff_index | 源ip所在xff头的索引位置（支持负数，-1为最后一位，以此类推） | -1 | int32有效值 |
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateFeatureV2(request *model.CreateFeatureV2Request) (*model.CreateFeatureV2Response, error) {
	requestDef := GenReqDefForCreateFeatureV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFeatureV2Response), nil
	}
}

// 创建分组自定义响应
//
// 新增分组下自定义响应
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateGatewayResponseV2(request *model.CreateGatewayResponseV2Request) (*model.CreateGatewayResponseV2Response, error) {
	requestDef := GenReqDefForCreateGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGatewayResponseV2Response), nil
	}
}

// 创建专享版实例
//
// 创建专享版实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateInstanceV2(request *model.CreateInstanceV2Request) (*model.CreateInstanceV2Response, error) {
	requestDef := GenReqDefForCreateInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceV2Response), nil
	}
}

// 创建流控策略
//
// 当API上线后，系统会默认给每个API提供一个流控策略，API提供者可以根据自身API的服务能力及负载情况变更这个流控策略。 流控策略即限制API在一定长度的时间内，能够允许被访问的最大次数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateRequestThrottlingPolicyV2(request *model.CreateRequestThrottlingPolicyV2Request) (*model.CreateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestThrottlingPolicyV2Response), nil
	}
}

// 创建签名密钥
//
// 为了保护API的安全性，建议租户为API的访问提供一套保护机制，即租户开放的API，需要对请求来源进行认证，不符合认证的请求直接拒绝访问。
//
// 其中，签名密钥就是API安全保护机制的一种。
//
// 租户创建一个签名密钥，并将签名密钥与API进行绑定，则API网关在请求这个API时，就会使用绑定的签名密钥对请求参数进行数据加密，生成签名。当租户的后端服务收到请求时，可以校验这个签名，如果签名校验不通过，则该请求不是API网关发出的请求，租户可以拒绝这个请求，从而保证API的安全性，避免API被未知来源的请求攻击。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateSignatureKeyV2(request *model.CreateSignatureKeyV2Request) (*model.CreateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForCreateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSignatureKeyV2Response), nil
	}
}

// 创建特殊设置
//
// 流控策略可以限制一段时间内可以访问API的最大次数，也可以限制一段时间内单个租户和单个APP可以访问API的最大次数。
//
// 如果想要对某个特定的APP进行特殊设置，例如设置所有APP每分钟的访问次数为500次，但想设置APP1每分钟的访问次数为800次，可以通过在流控策略中设置特殊APP来实现该功能。
//
// 为流控策略添加一个特殊设置的对象，可以是APP，也可以是租户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateSpecialThrottlingConfigurationV2(request *model.CreateSpecialThrottlingConfigurationV2Request) (*model.CreateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpecialThrottlingConfigurationV2Response), nil
	}
}

// 删除自定义认证
//
// 删除自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteCustomAuthorizerV2(request *model.DeleteCustomAuthorizerV2Request) (*model.DeleteCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomAuthorizerV2Response), nil
	}
}

// 删除环境
//
// 删除指定的环境。
//
// 该操作将导致此API在指定的环境无法被访问，可能会影响相当一部分应用和用户。请确保已经告知用户，或者确认需要强制下线。
//
// 环境上存在已发布的API时，该环境不能被删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteEnvironmentV2(request *model.DeleteEnvironmentV2Request) (*model.DeleteEnvironmentV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentV2Response), nil
	}
}

// 删除变量
//
// 删除指定的环境变量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteEnvironmentVariableV2(request *model.DeleteEnvironmentVariableV2Request) (*model.DeleteEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentVariableV2Response), nil
	}
}

// 删除分组指定错误类型的自定义响应配置
//
// 删除分组指定错误类型的自定义响应配置，还原为使用默认值的配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteGatewayResponseTypeV2(request *model.DeleteGatewayResponseTypeV2Request) (*model.DeleteGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForDeleteGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGatewayResponseTypeV2Response), nil
	}
}

// 删除分组自定义响应
//
// 删除分组自定义响应
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteGatewayResponseV2(request *model.DeleteGatewayResponseV2Request) (*model.DeleteGatewayResponseV2Response, error) {
	requestDef := GenReqDefForDeleteGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGatewayResponseV2Response), nil
	}
}

// 删除专享版实例
//
// 删除专享版实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteInstancesV2(request *model.DeleteInstancesV2Request) (*model.DeleteInstancesV2Response, error) {
	requestDef := GenReqDefForDeleteInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesV2Response), nil
	}
}

// 删除流控策略
//
// 删除指定的流控策略,以及该流控策略与API的所有绑定关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteRequestThrottlingPolicyV2(request *model.DeleteRequestThrottlingPolicyV2Request) (*model.DeleteRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestThrottlingPolicyV2Response), nil
	}
}

// 删除签名密钥
//
// 删除指定的签名密钥,删除签名密钥时，其配置的绑定关系会一并删除，相应的签名密钥会失效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteSignatureKeyV2(request *model.DeleteSignatureKeyV2Request) (*model.DeleteSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDeleteSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSignatureKeyV2Response), nil
	}
}

// 删除特殊设置
//
// 删除某个流控策略的某个特殊配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteSpecialThrottlingConfigurationV2(request *model.DeleteSpecialThrottlingConfigurationV2Request) (*model.DeleteSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSpecialThrottlingConfigurationV2Response), nil
	}
}

// 删除域名证书
//
// 如果域名证书不再需要或者已过期，则可以删除证书内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DisassociateCertificateV2(request *model.DisassociateCertificateV2Request) (*model.DisassociateCertificateV2Response, error) {
	requestDef := GenReqDefForDisassociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateCertificateV2Response), nil
	}
}

// 解绑域名
//
// 如果API分组不再需要绑定某个自定义域名，则可以为此API分组解绑此域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DisassociateDomainV2(request *model.DisassociateDomainV2Request) (*model.DisassociateDomainV2Response, error) {
	requestDef := GenReqDefForDisassociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateDomainV2Response), nil
	}
}

// 解除绑定
//
// 解除API与签名密钥的绑定关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DisassociateSignatureKeyV2(request *model.DisassociateSignatureKeyV2Request) (*model.DisassociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateSignatureKeyV2Response), nil
	}
}

// 查询API分组概况
//
// 查询租户名下的API分组概况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiGroupsQuantitiesV2(request *model.ListApiGroupsQuantitiesV2Request) (*model.ListApiGroupsQuantitiesV2Response, error) {
	requestDef := GenReqDefForListApiGroupsQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiGroupsQuantitiesV2Response), nil
	}
}

// 查询API概况
//
// 查询租户名下的API概况：已发布到RELEASE环境的API个数，未发布到RELEASE环境的API个数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiQuantitiesV2(request *model.ListApiQuantitiesV2Request) (*model.ListApiQuantitiesV2Response, error) {
	requestDef := GenReqDefForListApiQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiQuantitiesV2Response), nil
	}
}

// 查看签名密钥绑定的API列表
//
// 查询某个签名密钥上已经绑定的API列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToSignatureKeyV2(request *model.ListApisBindedToSignatureKeyV2Request) (*model.ListApisBindedToSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToSignatureKeyV2Response), nil
	}
}

// 查看签名密钥未绑定的API列表
//
// 查询所有未绑定到该签名密钥上的API列表。需要API已经发布，未发布的API不予展示。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisNotBoundWithSignatureKeyV2(request *model.ListApisNotBoundWithSignatureKeyV2Request) (*model.ListApisNotBoundWithSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisNotBoundWithSignatureKeyV2Response), nil
	}
}

// 查询APP概况
//
// 查询租户名下的APP概况：已进行API访问授权的APP个数，未进行API访问授权的APP个数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAppQuantitiesV2(request *model.ListAppQuantitiesV2Request) (*model.ListAppQuantitiesV2Response, error) {
	requestDef := GenReqDefForListAppQuantitiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuantitiesV2Response), nil
	}
}

// 查看可用区信息
//
// 查看可用区信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAvailableZonesV2(request *model.ListAvailableZonesV2Request) (*model.ListAvailableZonesV2Response, error) {
	requestDef := GenReqDefForListAvailableZonesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesV2Response), nil
	}
}

// 查询自定义认证列表
//
// 查询自定义认证列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListCustomAuthorizersV2(request *model.ListCustomAuthorizersV2Request) (*model.ListCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForListCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomAuthorizersV2Response), nil
	}
}

// 查询变量列表
//
// 查询分组下的所有环境变量的列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListEnvironmentVariablesV2(request *model.ListEnvironmentVariablesV2Request) (*model.ListEnvironmentVariablesV2Response, error) {
	requestDef := GenReqDefForListEnvironmentVariablesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentVariablesV2Response), nil
	}
}

// 查询环境列表
//
// 查询符合条件的环境列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListEnvironmentsV2(request *model.ListEnvironmentsV2Request) (*model.ListEnvironmentsV2Response, error) {
	requestDef := GenReqDefForListEnvironmentsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsV2Response), nil
	}
}

// 查看实例特性列表
//
// 查看实例特性列表。注意：实例不支持以下特性的需要联系技术支持升级实例版本。
//
// 当前支持的特性列表如下：
//
// 特性名称 | 特性描述 | 特性是否可配置|
// --------| :------- | :-------|
// lts | 是否支持shubao访问日志上报功能。| 是 |
// gateway_responses | 是否支持网关自定义响应。| 否 |
// ratelimit | 是否支持自定义流控值。| 是 |
// request_body_size | 是否支持指定最大请求Body大小。| 是 |
// backend_timeout | 是否支持配置后端API最大超时时间。| 是 |
// app_token | 是否开启app_token认证方式。| 是 |
// app_api_key | 是否开启app_api_key认证方式。| 是 |
// app_basic | 是否开启app_basic认证方式。| 是 |
// app_secret | 是否支持app_secret认证方式。| 是 |
// app_jwt | 是否支持app_jwt认证方式。| 是 |
// public_key | 是否支持public_key类型的后端签名。| 是 |
// backend_token_allow | 是否支持普通租户透传token到后端。| 是 |
// sign_basic | 签名秘钥是否支持basic类型。| 否 |
// multi_auth | API是否支持双重认证方式。| 否 |
// backend_client_certificate | 是否开启后端双向认证。| 是 |
// ssl_ciphers | 是否支持https加密套件。  | 是 |
// route | 是否支持自定义路由。| 否 |
// cors | 是否支持API使用插件功能。| 否 |
// real_ip_from_xff | 是否开启使用xff头作为访问控制、流控策略的源ip生效依据。  | 是 |
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListFeaturesV2(request *model.ListFeaturesV2Request) (*model.ListFeaturesV2Response, error) {
	requestDef := GenReqDefForListFeaturesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesV2Response), nil
	}
}

// 查询分组自定义响应列表
//
// 查询分组自定义响应列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListGatewayResponsesV2(request *model.ListGatewayResponsesV2Request) (*model.ListGatewayResponsesV2Response, error) {
	requestDef := GenReqDefForListGatewayResponsesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGatewayResponsesV2Response), nil
	}
}

// 查询租户实例配置列表
//
// 查询租户实例配置列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListInstanceConfigsV2(request *model.ListInstanceConfigsV2Request) (*model.ListInstanceConfigsV2Response, error) {
	requestDef := GenReqDefForListInstanceConfigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConfigsV2Response), nil
	}
}

// 查询专享版实例列表
//
// 查询专享版实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListInstancesV2(request *model.ListInstancesV2Request) (*model.ListInstancesV2Response, error) {
	requestDef := GenReqDefForListInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesV2Response), nil
	}
}

// API统计信息查询-最近一段时间
//
// 根据API的id和最近的一段时间查询API被调用的次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListLatelyApiStatisticsV2(request *model.ListLatelyApiStatisticsV2Request) (*model.ListLatelyApiStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyApiStatisticsV2Response), nil
	}
}

// 分组统计信息查询-最近一小时内
//
// 根据API分组的编号查询该分组下所有API被调用的总次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListLatelyGroupStatisticsV2(request *model.ListLatelyGroupStatisticsV2Request) (*model.ListLatelyGroupStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyGroupStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyGroupStatisticsV2Response), nil
	}
}

// 查询某个实例的租户配置列表
//
// 查询某个实例的租户配置列表，用户可以通过此接口查看各类型资源配置及使用情况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListProjectCofigsV2(request *model.ListProjectCofigsV2Request) (*model.ListProjectCofigsV2Response, error) {
	requestDef := GenReqDefForListProjectCofigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectCofigsV2Response), nil
	}
}

// 查询流控策略列表
//
// 查询所有流控策略的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListRequestThrottlingPolicyV2(request *model.ListRequestThrottlingPolicyV2Request) (*model.ListRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPolicyV2Response), nil
	}
}

// 查看API绑定的签名密钥列表
//
// 查询某个API绑定的签名密钥列表。每个API在每个环境上应该最多只会绑定一个签名密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListSignatureKeysBindedToApiV2(request *model.ListSignatureKeysBindedToApiV2Request) (*model.ListSignatureKeysBindedToApiV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysBindedToApiV2Response), nil
	}
}

// 查询签名密钥列表
//
// 查询所有签名密钥的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListSignatureKeysV2(request *model.ListSignatureKeysV2Request) (*model.ListSignatureKeysV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysV2Response), nil
	}
}

// 查看特殊设置列表
//
// 查看给流控策略设置的特殊配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListSpecialThrottlingConfigurationsV2(request *model.ListSpecialThrottlingConfigurationsV2Request) (*model.ListSpecialThrottlingConfigurationsV2Response, error) {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecialThrottlingConfigurationsV2Response), nil
	}
}

// 查询标签列表
//
// 查询标签列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListTagsV2(request *model.ListTagsV2Request) (*model.ListTagsV2Response, error) {
	requestDef := GenReqDefForListTagsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsV2Response), nil
	}
}

// 实例解绑EIP
//
// 实例解绑EIP
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) RemoveEipV2(request *model.RemoveEipV2Request) (*model.RemoveEipV2Response, error) {
	requestDef := GenReqDefForRemoveEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveEipV2Response), nil
	}
}

// 关闭实例公网出口
//
// 关闭实例公网出口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) RemoveEngressEipV2(request *model.RemoveEngressEipV2Request) (*model.RemoveEngressEipV2Response, error) {
	requestDef := GenReqDefForRemoveEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveEngressEipV2Response), nil
	}
}

// 查看自定义认证详情
//
// 查看自定义认证详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfCustomAuthorizersV2(request *model.ShowDetailsOfCustomAuthorizersV2Request) (*model.ShowDetailsOfCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfCustomAuthorizersV2Response), nil
	}
}

// 查看域名证书
//
// 查看域名下绑定的证书详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfDomainNameCertificateV2(request *model.ShowDetailsOfDomainNameCertificateV2Request) (*model.ShowDetailsOfDomainNameCertificateV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfDomainNameCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfDomainNameCertificateV2Response), nil
	}
}

// 查看变量详情
//
// 查看指定的环境变量的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfEnvironmentVariableV2(request *model.ShowDetailsOfEnvironmentVariableV2Request) (*model.ShowDetailsOfEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfEnvironmentVariableV2Response), nil
	}
}

// 查看分组下指定错误类型的自定义响应
//
// 查看分组下指定错误类型的自定义响应
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfGatewayResponseTypeV2(request *model.ShowDetailsOfGatewayResponseTypeV2Request) (*model.ShowDetailsOfGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfGatewayResponseTypeV2Response), nil
	}
}

// 查询分组自定义响应详情
//
// 查询分组自定义响应详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfGatewayResponseV2(request *model.ShowDetailsOfGatewayResponseV2Request) (*model.ShowDetailsOfGatewayResponseV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfGatewayResponseV2Response), nil
	}
}

// 查看专享版实例创建进度
//
// 查看专享版实例创建进度
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfInstanceProgressV2(request *model.ShowDetailsOfInstanceProgressV2Request) (*model.ShowDetailsOfInstanceProgressV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfInstanceProgressV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfInstanceProgressV2Response), nil
	}
}

// 查看专享版实例详情
//
// 查看专享版实例详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfInstanceV2(request *model.ShowDetailsOfInstanceV2Request) (*model.ShowDetailsOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfInstanceV2Response), nil
	}
}

// 查看流控策略详情
//
// 查看指定流控策略的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfRequestThrottlingPolicyV2(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfRequestThrottlingPolicyV2Response), nil
	}
}

// 修改自定义认证
//
// 修改自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateCustomAuthorizerV2(request *model.UpdateCustomAuthorizerV2Request) (*model.UpdateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomAuthorizerV2Response), nil
	}
}

// 修改域名
//
// 修改绑定的域名所对应的配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateDomainV2(request *model.UpdateDomainV2Request) (*model.UpdateDomainV2Response, error) {
	requestDef := GenReqDefForUpdateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainV2Response), nil
	}
}

// 更新实例出公网带宽
//
// 更新实例出公网带宽
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateEngressEipV2(request *model.UpdateEngressEipV2Request) (*model.UpdateEngressEipV2Response, error) {
	requestDef := GenReqDefForUpdateEngressEipV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEngressEipV2Response), nil
	}
}

// 修改环境
//
// 修改指定环境的信息。其中可修改的属性为：name、remark，其它属性不可修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateEnvironmentV2(request *model.UpdateEnvironmentV2Request) (*model.UpdateEnvironmentV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentV2Response), nil
	}
}

// 修改分组下指定错误类型的自定义响应
//
// 修改分组下指定错误类型的自定义响应。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateGatewayResponseTypeV2(request *model.UpdateGatewayResponseTypeV2Request) (*model.UpdateGatewayResponseTypeV2Response, error) {
	requestDef := GenReqDefForUpdateGatewayResponseTypeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGatewayResponseTypeV2Response), nil
	}
}

// 修改分组自定义响应
//
// 修改分组自定义响应
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateGatewayResponseV2(request *model.UpdateGatewayResponseV2Request) (*model.UpdateGatewayResponseV2Response, error) {
	requestDef := GenReqDefForUpdateGatewayResponseV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGatewayResponseV2Response), nil
	}
}

// 更新专享版实例
//
// 更新专享版实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateInstanceV2(request *model.UpdateInstanceV2Request) (*model.UpdateInstanceV2Response, error) {
	requestDef := GenReqDefForUpdateInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceV2Response), nil
	}
}

// 修改流控策略
//
// 修改指定流控策略的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateRequestThrottlingPolicyV2(request *model.UpdateRequestThrottlingPolicyV2Request) (*model.UpdateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestThrottlingPolicyV2Response), nil
	}
}

// 修改签名密钥
//
// 修改指定签名密钥的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateSignatureKeyV2(request *model.UpdateSignatureKeyV2Request) (*model.UpdateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForUpdateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSignatureKeyV2Response), nil
	}
}

// 修改特殊设置
//
// 修改某个流控策略下的某个特殊设置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateSpecialThrottlingConfigurationV2(request *model.UpdateSpecialThrottlingConfigurationV2Request) (*model.UpdateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpecialThrottlingConfigurationV2Response), nil
	}
}

// 批量删除ACL策略
//
// 批量删除指定的多个ACL策略。
//
// 删除ACL策略时，如果存在ACL策略与API绑定关系，则无法删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) BatchDeleteAclV2(request *model.BatchDeleteAclV2Request) (*model.BatchDeleteAclV2Response, error) {
	requestDef := GenReqDefForBatchDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclV2Response), nil
	}
}

// 创建ACL策略
//
// 增加一个ACL策略，策略类型通过字段acl_type来确定（permit或者deny），限制的对象的类型可以为IP或者DOMAIN，这里的DOMAIN对应的acl_value的值为租户名称，而非“www.exampleDomain.com\&quot;之类的网络域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateAclStrategyV2(request *model.CreateAclStrategyV2Request) (*model.CreateAclStrategyV2Response, error) {
	requestDef := GenReqDefForCreateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAclStrategyV2Response), nil
	}
}

// 删除ACL策略
//
// 删除指定的ACL策略， 如果存在api与该ACL策略的绑定关系，则无法删除
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteAclV2(request *model.DeleteAclV2Request) (*model.DeleteAclV2Response, error) {
	requestDef := GenReqDefForDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclV2Response), nil
	}
}

// 查看ACL策略列表
//
// 查询所有的ACL策略列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAclStrategiesV2(request *model.ListAclStrategiesV2Request) (*model.ListAclStrategiesV2Response, error) {
	requestDef := GenReqDefForListAclStrategiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclStrategiesV2Response), nil
	}
}

// 查看ACL策略详情
//
// 查询指定ACL策略的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAclPolicyV2(request *model.ShowDetailsOfAclPolicyV2Request) (*model.ShowDetailsOfAclPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAclPolicyV2Response), nil
	}
}

// 修改ACL策略
//
// 修改指定的ACL策略，其中可修改的属性为：acl_name、acl_type、acl_value，其它属性不可修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateAclStrategyV2(request *model.UpdateAclStrategyV2Request) (*model.UpdateAclStrategyV2Response, error) {
	requestDef := GenReqDefForUpdateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclStrategyV2Response), nil
	}
}

// 绑定流控策略
//
// 将流控策略应用于API，则所有对该API的访问将会受到该流控策略的限制。
//
// 当一定时间内的访问次数超过流控策略设置的API最大访问次数限制后，后续的访问将会被拒绝，从而能够较好的保护后端API免受异常流量的冲击，保障服务的稳定运行。
//
// 为指定的API绑定流控策略，绑定时，需要指定在哪个环境上生效。同一个API发布到不同的环境可以绑定不同的流控策略；一个API在发布到特定环境后只能绑定一个默认的流控策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AssociateRequestThrottlingPolicyV2(request *model.AssociateRequestThrottlingPolicyV2Request) (*model.AssociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForAssociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRequestThrottlingPolicyV2Response), nil
	}
}

// 批量解绑流控策略
//
// 批量解除API与流控策略的绑定关系
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) BatchDisassociateThrottlingPolicyV2(request *model.BatchDisassociateThrottlingPolicyV2Request) (*model.BatchDisassociateThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForBatchDisassociateThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateThrottlingPolicyV2Response), nil
	}
}

// 批量发布或下线API
//
// 将多个API发布到一个指定的环境，或将多个API从指定的环境下线。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) BatchPublishOrOfflineApiV2(request *model.BatchPublishOrOfflineApiV2Request) (*model.BatchPublishOrOfflineApiV2Response, error) {
	requestDef := GenReqDefForBatchPublishOrOfflineApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPublishOrOfflineApiV2Response), nil
	}
}

// 切换API版本
//
// API每次发布时，会基于当前的API定义生成一个版本。版本记录了API发布时的各种定义及状态。
//
// 多个版本之间可以进行随意切换。但一个API在一个环境上，只能有一个版本生效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ChangeApiVersionV2(request *model.ChangeApiVersionV2Request) (*model.ChangeApiVersionV2Response, error) {
	requestDef := GenReqDefForChangeApiVersionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApiVersionV2Response), nil
	}
}

// 创建API分组
//
// API分组是API的管理单元，一个API分组等同于一个服务入口，创建API分组时，返回一个子域名作为访问入口。建议一个API分组下的API具有一定的相关性。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateApiGroupV2(request *model.CreateApiGroupV2Request) (*model.CreateApiGroupV2Response, error) {
	requestDef := GenReqDefForCreateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiGroupV2Response), nil
	}
}

// 注册API
//
// 添加一个API，API即一个服务接口，具体的服务能力。
//
// API分为两部分，第一部分为面向API使用者的API接口，定义了使用者如何调用这个API。第二部分面向API提供者，由API提供者定义这个API的真实的后端情况，定义了API网关如何去访问真实的后端服务。API的真实后端服务目前支持三种类型：传统的HTTP/HTTPS形式的web后端、函数工作流、MOCK。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateApiV2(request *model.CreateApiV2Request) (*model.CreateApiV2Response, error) {
	requestDef := GenReqDefForCreateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiV2Response), nil
	}
}

// 发布或下线API
//
// 对API进行发布或下线。
//
// 发布操作是将一个指定的API发布到一个指定的环境，API只有发布后，才能够被调用，且只能在该环境上才能被调用。未发布的API无法被调用。
//
// 下线操作是将API从某个已发布的环境上下线，下线后，API将无法再被调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateOrDeletePublishRecordForApiV2(request *model.CreateOrDeletePublishRecordForApiV2Request) (*model.CreateOrDeletePublishRecordForApiV2Response, error) {
	requestDef := GenReqDefForCreateOrDeletePublishRecordForApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrDeletePublishRecordForApiV2Response), nil
	}
}

// 调试API
//
// 调试一个API在指定运行环境下的定义，接口调用者需要具有操作该API的权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DebugApiV2(request *model.DebugApiV2Request) (*model.DebugApiV2Response, error) {
	requestDef := GenReqDefForDebugApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugApiV2Response), nil
	}
}

// 根据版本编号下线API
//
// 对某个生效中的API版本进行下线操作，下线后，API在该版本生效的环境中将不再能够被调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteApiByVersionIdV2(request *model.DeleteApiByVersionIdV2Request) (*model.DeleteApiByVersionIdV2Response, error) {
	requestDef := GenReqDefForDeleteApiByVersionIdV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiByVersionIdV2Response), nil
	}
}

// 删除API分组
//
// 删除指定的API分组。
// 删除API分组前，要先下线并删除分组下的所有API。
// 删除时，会一并删除直接或间接关联到该分组下的所有资源，包括独立域名、SSL证书信息等等。并会将外部域名与子域名的绑定关系进行解除（取决于域名cname方式）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteApiGroupV2(request *model.DeleteApiGroupV2Request) (*model.DeleteApiGroupV2Response, error) {
	requestDef := GenReqDefForDeleteApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiGroupV2Response), nil
	}
}

// 删除API
//
// 删除指定的API。
//
// 删除API时，会删除该API所有相关的资源信息或绑定关系，如API的发布记录，绑定的后端服务，对APP的授权信息等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteApiV2(request *model.DeleteApiV2Request) (*model.DeleteApiV2Response, error) {
	requestDef := GenReqDefForDeleteApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiV2Response), nil
	}
}

// 解除API与流控策略的绑定关系
//
// 解除API与流控策略的绑定关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DisassociateRequestThrottlingPolicyV2(request *model.DisassociateRequestThrottlingPolicyV2Request) (*model.DisassociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDisassociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRequestThrottlingPolicyV2Response), nil
	}
}

// 查询分组列表
//
// 查询API分组列表。
//
// 如果是租户操作，则查询该租户下所有的分组；如果是管理员操作，则查询的是所有租户的分组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiGroupsV2(request *model.ListApiGroupsV2Request) (*model.ListApiGroupsV2Response, error) {
	requestDef := GenReqDefForListApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiGroupsV2Response), nil
	}
}

// 查询API运行时定义
//
// 查看指定的API在指定的环境上的运行时定义，默认查询RELEASE环境上的运行时定义。
//
// API的定义分为临时定义和运行时定义，分别代表如下含义：
// - 临时定义：API在编辑中的定义，表示用户最后一次编辑后的API的状态
// - 运行时定义：API在发布到某个环境时，对发布时的API的临时定义进行快照，固化出来的API的状态。
//
// 访问某个环境上的API，其实访问的就是其运行时的定义
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiRuntimeDefinitionV2(request *model.ListApiRuntimeDefinitionV2Request) (*model.ListApiRuntimeDefinitionV2Response, error) {
	requestDef := GenReqDefForListApiRuntimeDefinitionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiRuntimeDefinitionV2Response), nil
	}
}

// 查看版本详情
//
// 查询某个指定的版本详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiVersionDetailV2(request *model.ListApiVersionDetailV2Request) (*model.ListApiVersionDetailV2Response, error) {
	requestDef := GenReqDefForListApiVersionDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionDetailV2Response), nil
	}
}

// 查询API历史版本列表
//
// 查询某个API的历史版本。每个API在一个环境上最多存在10个历史版本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApiVersionsV2(request *model.ListApiVersionsV2Request) (*model.ListApiVersionsV2Response, error) {
	requestDef := GenReqDefForListApiVersionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsV2Response), nil
	}
}

// 查看流控策略绑定的API列表
//
// 查询某个流控策略上已经绑定的API列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToRequestThrottlingPolicyV2(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) (*model.ListApisBindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToRequestThrottlingPolicyV2Response), nil
	}
}

// 查看流控策略未绑定的API列表
//
// 查询所有未绑定到该流控策略上的自有API列表。需要API已经发布，未发布的API不予展示。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToRequestThrottlingPolicyV2(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToRequestThrottlingPolicyV2Response), nil
	}
}

// 查询API列表
//
// 查看API列表，返回API详细信息、发布信息等，但不能查看到后端服务信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisV2(request *model.ListApisV2Request) (*model.ListApisV2Response, error) {
	requestDef := GenReqDefForListApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisV2Response), nil
	}
}

// 查看API绑定的流控策略列表
//
// 查询某个API绑定的流控策略列表。每个环境上应该最多只有一个流控策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListRequestThrottlingPoliciesBindedToApiV2(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPoliciesBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPoliciesBindedToApiV2Response), nil
	}
}

// 查询分组详情
//
// 查询指定分组的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfApiGroupV2(request *model.ShowDetailsOfApiGroupV2Request) (*model.ShowDetailsOfApiGroupV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiGroupV2Response), nil
	}
}

// 查询API详情
//
// 查看指定的API的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfApiV2(request *model.ShowDetailsOfApiV2Request) (*model.ShowDetailsOfApiV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiV2Response), nil
	}
}

// 修改API分组
//
// 修改API分组属性。其中name和remark可修改，其他属性不可修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateApiGroupV2(request *model.UpdateApiGroupV2Request) (*model.UpdateApiGroupV2Response, error) {
	requestDef := GenReqDefForUpdateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiGroupV2Response), nil
	}
}

// 修改API
//
// 修改指定API的信息，包括后端服务信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateApiV2(request *model.UpdateApiV2Request) (*model.UpdateApiV2Response, error) {
	requestDef := GenReqDefForUpdateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiV2Response), nil
	}
}

// 批量解除API与ACL策略的绑定
//
// 批量解除API与ACL策略的绑定
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) BatchDeleteApiAclBindingV2(request *model.BatchDeleteApiAclBindingV2Request) (*model.BatchDeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForBatchDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteApiAclBindingV2Response), nil
	}
}

// 将API与ACL策略进行绑定
//
// 将API与ACL策略进行绑定。
//
// 同一个API发布到不同的环境可以绑定不同的ACL策略；一个API在发布到特定环境后只能绑定一个同一种类型的ACL策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateApiAclBindingV2(request *model.CreateApiAclBindingV2Request) (*model.CreateApiAclBindingV2Response, error) {
	requestDef := GenReqDefForCreateApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiAclBindingV2Response), nil
	}
}

// 解除API与ACL策略的绑定
//
// 解除某条API与ACL策略的绑定关系
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteApiAclBindingV2(request *model.DeleteApiAclBindingV2Request) (*model.DeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiAclBindingV2Response), nil
	}
}

// 查看API绑定的ACL策略列表
//
// 查看API绑定的ACL策略列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAclPolicyBindedToApiV2(request *model.ListAclPolicyBindedToApiV2Request) (*model.ListAclPolicyBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAclPolicyBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclPolicyBindedToApiV2Response), nil
	}
}

// 查看ACL策略绑定的API列表
//
// 查看ACL策略绑定的API列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToAclPolicyV2(request *model.ListApisBindedToAclPolicyV2Request) (*model.ListApisBindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAclPolicyV2Response), nil
	}
}

// 查看ACL策略未绑定的API列表
//
// 查看ACL策略未绑定的API列表，需要API已发布
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToAclPolicyV2(request *model.ListApisUnbindedToAclPolicyV2Request) (*model.ListApisUnbindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAclPolicyV2Response), nil
	}
}

// 解除授权
//
// 解除API对APP的授权关系。解除授权后，APP将不再能够调用该API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CancelingAuthorizationV2(request *model.CancelingAuthorizationV2Request) (*model.CancelingAuthorizationV2Response, error) {
	requestDef := GenReqDefForCancelingAuthorizationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelingAuthorizationV2Response), nil
	}
}

// 校验APP
//
// 校验app是否存在，非APP所有者可以调用该接口校验APP是否真实存在。这个接口只展示app的基本信息id 、name、 remark，其他信息不显示。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CheckAppV2(request *model.CheckAppV2Request) (*model.CheckAppV2Response, error) {
	requestDef := GenReqDefForCheckAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAppV2Response), nil
	}
}

// 创建APP
//
// APP即应用，是一个可以访问API的身份标识。将API授权给APP后，APP即可调用API。
// 创建一个APP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateAnAppV2(request *model.CreateAnAppV2Request) (*model.CreateAnAppV2Response, error) {
	requestDef := GenReqDefForCreateAnAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnAppV2Response), nil
	}
}

// 自动生成APP Code
//
// 创建App Code时，可以不指定具体值，由后台自动生成随机字符串填充。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateAppCodeAutoV2(request *model.CreateAppCodeAutoV2Request) (*model.CreateAppCodeAutoV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeAutoV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeAutoV2Response), nil
	}
}

// 创建APP Code
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateAppCodeV2(request *model.CreateAppCodeV2Request) (*model.CreateAppCodeV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeV2Response), nil
	}
}

// APP授权
//
// APP创建成功后，还不能访问API，如果想要访问某个环境上的API，需要将该API在该环境上授权给APP。授权成功后，APP即可访问该环境上的这个API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateAuthorizingAppsV2(request *model.CreateAuthorizingAppsV2Request) (*model.CreateAuthorizingAppsV2Response, error) {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizingAppsV2Response), nil
	}
}

// 删除APP Code
//
// 删除App Code，App Code删除后，将无法再通过简易认证访问对应的API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteAppCodeV2(request *model.DeleteAppCodeV2Request) (*model.DeleteAppCodeV2Response, error) {
	requestDef := GenReqDefForDeleteAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppCodeV2Response), nil
	}
}

// 删除APP
//
// 删除指定的APP。
// APP删除后，将无法再调用任何API；其中，云市场自动创建的APP无法被删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteAppV2(request *model.DeleteAppV2Request) (*model.DeleteAppV2Response, error) {
	requestDef := GenReqDefForDeleteAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppV2Response), nil
	}
}

// 查看APP已绑定的API列表
//
// 查询APP已经绑定的API列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisBindedToAppV2(request *model.ListApisBindedToAppV2Request) (*model.ListApisBindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAppV2Response), nil
	}
}

// 查看APP未绑定的API列表
//
// 查询指定环境上某个APP未绑定的API列表，包括自有API和从云市场购买的API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListApisUnbindedToAppV2(request *model.ListApisUnbindedToAppV2Request) (*model.ListApisUnbindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAppV2Response), nil
	}
}

// 查询APP Code列表
//
// 查询App Code列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAppCodesV2(request *model.ListAppCodesV2Request) (*model.ListAppCodesV2Response, error) {
	requestDef := GenReqDefForListAppCodesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppCodesV2Response), nil
	}
}

// 查看API已绑定的APP列表
//
// 查询API绑定的APP列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAppsBindedToApiV2(request *model.ListAppsBindedToApiV2Request) (*model.ListAppsBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAppsBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsBindedToApiV2Response), nil
	}
}

// 查询APP列表
//
// 查询APP列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListAppsV2(request *model.ListAppsV2Request) (*model.ListAppsV2Response, error) {
	requestDef := GenReqDefForListAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsV2Response), nil
	}
}

// 重置密钥
//
// 重置指定APP的密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ResettingAppSecretV2(request *model.ResettingAppSecretV2Request) (*model.ResettingAppSecretV2Response, error) {
	requestDef := GenReqDefForResettingAppSecretV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResettingAppSecretV2Response), nil
	}
}

// 查看APP Code详情
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAppCodeV2(request *model.ShowDetailsOfAppCodeV2Request) (*model.ShowDetailsOfAppCodeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppCodeV2Response), nil
	}
}

// 查看APP详情
//
// 查看指定APP的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfAppV2(request *model.ShowDetailsOfAppV2Request) (*model.ShowDetailsOfAppV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppV2Response), nil
	}
}

// 修改APP
//
// 修改指定APP的信息。其中可修改的属性为：name、remark，当支持用户自定义key和secret的开关开启时，app_key和app_secret也支持修改，其它属性不可修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateAppV2(request *model.UpdateAppV2Request) (*model.UpdateAppV2Response, error) {
	requestDef := GenReqDefForUpdateAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppV2Response), nil
	}
}

// 导出API
//
// 导出分组下API的定义信息。导出文件内容符合swagger标准规范，API网关自定义扩展字段请参考《API网关开发指南》的“导入导出API：扩展定义”章节。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ExportApiDefinitionsV2(request *model.ExportApiDefinitionsV2Request) (*model.ExportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportApiDefinitionsV2Response), nil
	}
}

// 导入API
//
// 导入API。导入文件内容需要符合swagger标准规范，API网关自定义扩展字段请参考《API网关开发指南》的“导入导出API：扩展定义”章节。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ImportApiDefinitionsV2(request *model.ImportApiDefinitionsV2Request) (*model.ImportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApiDefinitionsV2Response), nil
	}
}

// 添加后端实例
//
// 为指定的VPC通道添加弹性云服务器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) AddingBackendInstancesV2(request *model.AddingBackendInstancesV2Request) (*model.AddingBackendInstancesV2Response, error) {
	requestDef := GenReqDefForAddingBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddingBackendInstancesV2Response), nil
	}
}

// 创建VPC通道
//
// 在API网关中创建连接私有VPC资源的通道，并在创建API时将后端节点配置为使用这些VPC通道，以便API网关直接访问私有VPC资源。
// &gt; 每个用户最多创建30个VPC通道。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) CreateVpcChannelV2(request *model.CreateVpcChannelV2Request) (*model.CreateVpcChannelV2Response, error) {
	requestDef := GenReqDefForCreateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcChannelV2Response), nil
	}
}

// 删除后端实例
//
// 删除指定VPC通道中的弹性云服务器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteBackendInstanceV2(request *model.DeleteBackendInstanceV2Request) (*model.DeleteBackendInstanceV2Response, error) {
	requestDef := GenReqDefForDeleteBackendInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackendInstanceV2Response), nil
	}
}

// 删除VPC通道
//
// 删除指定的VPC通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) DeleteVpcChannelV2(request *model.DeleteVpcChannelV2Request) (*model.DeleteVpcChannelV2Response, error) {
	requestDef := GenReqDefForDeleteVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcChannelV2Response), nil
	}
}

// 查看后端实例列表
//
// 查看指定VPC通道的弹性云服务器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListBackendInstancesV2(request *model.ListBackendInstancesV2Request) (*model.ListBackendInstancesV2Response, error) {
	requestDef := GenReqDefForListBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackendInstancesV2Response), nil
	}
}

// 查询VPC通道列表
//
// 查看VPC通道列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ListVpcChannelsV2(request *model.ListVpcChannelsV2Request) (*model.ListVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcChannelsV2Response), nil
	}
}

// 查看VPC通道详情
//
// 查看指定的VPC通道详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) ShowDetailsOfVpcChannelV2(request *model.ShowDetailsOfVpcChannelV2Request) (*model.ShowDetailsOfVpcChannelV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfVpcChannelV2Response), nil
	}
}

// 更新VPC通道
//
// 更新指定VPC通道的参数
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApigClient) UpdateVpcChannelV2(request *model.UpdateVpcChannelV2Request) (*model.UpdateVpcChannelV2Response, error) {
	requestDef := GenReqDefForUpdateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcChannelV2Response), nil
	}
}
