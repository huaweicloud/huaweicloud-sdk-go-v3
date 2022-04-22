package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/roma/v2/model"
)

type RomaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRomaClient(hcClient *http_client.HcHttpClient) *RomaClient {
	return &RomaClient{HcClient: hcClient}
}

func RomaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 添加子设备到网关
//
// 添加子设备到网关
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AddSubsetsToGateway(request *model.AddSubsetsToGatewayRequest) (*model.AddSubsetsToGatewayResponse, error) {
	requestDef := GenReqDefForAddSubsetsToGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubsetsToGatewayResponse), nil
	}
}

// 客户端配额绑定客户端应用列表
//
// 客户端配额绑定客户端应用列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AssociateAppsForAppQuota(request *model.AssociateAppsForAppQuotaRequest) (*model.AssociateAppsForAppQuotaResponse, error) {
	requestDef := GenReqDefForAssociateAppsForAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateAppsForAppQuotaResponse), nil
	}
}

// 绑定域名证书
//
// 如果创建API时，“定义API请求”使用HTTPS请求协议，那么在独立域名中需要添加SSL证书。
// 本章节主要介绍为特定域名绑定证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AssociateCertificateV2(request *model.AssociateCertificateV2Request) (*model.AssociateCertificateV2Response, error) {
	requestDef := GenReqDefForAssociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateCertificateV2Response), nil
	}
}

// 绑定域名
//
// 用户自定义的域名，需要CNAME到API分组的子域名上才能生效。 每个API分组下最多可绑定5个域名。绑定域名后，用户可通过自定义域名调用API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AssociateDomainV2(request *model.AssociateDomainV2Request) (*model.AssociateDomainV2Response, error) {
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
// 将签名密钥绑定到API后，则ROMA Connect APIC请求后端服务时就会使用这个签名密钥进行加密签名，后端服务可以校验这个签名来验证请求来源。
//
// 将指定的签名密钥绑定到一个或多个已发布的API上。同一个API发布到不同的环境可以绑定不同的签名密钥；一个API在发布到特定环境后只能绑定一个签名密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AssociateSignatureKeyV2(request *model.AssociateSignatureKeyV2Request) (*model.AssociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForAssociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSignatureKeyV2Response), nil
	}
}

// 批量添加设备到设备分组
//
// 批量添加设备到设备分组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchAddDeviceToGroup(request *model.BatchAddDeviceToGroupRequest) (*model.BatchAddDeviceToGroupResponse, error) {
	requestDef := GenReqDefForBatchAddDeviceToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDeviceToGroupResponse), nil
	}
}

// 批量删除设备
//
// 批量删除设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchDeleteDevices(request *model.BatchDeleteDevicesRequest) (*model.BatchDeleteDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDevicesResponse), nil
	}
}

// 批量删除Topic
//
// 批量删除Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchDeleteMqsInstanceTopic(request *model.BatchDeleteMqsInstanceTopicRequest) (*model.BatchDeleteMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForBatchDeleteMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMqsInstanceTopicResponse), nil
	}
}

// 批量删除规则
//
// 批量删除规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchDeleteRules(request *model.BatchDeleteRulesRequest) (*model.BatchDeleteRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRulesResponse), nil
	}
}

// 设备批量下线
//
// 设备批量下线
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchFreezeDevices(request *model.BatchFreezeDevicesRequest) (*model.BatchFreezeDevicesResponse, error) {
	requestDef := GenReqDefForBatchFreezeDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchFreezeDevicesResponse), nil
	}
}

// 批量启动\\停止任务
//
// 批量启动\\停止任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchStartOrStopTasks(request *model.BatchStartOrStopTasksRequest) (*model.BatchStartOrStopTasksResponse, error) {
	requestDef := GenReqDefForBatchStartOrStopTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartOrStopTasksResponse), nil
	}
}

// 校验自定义后端API定义
//
// 校验自定义后端API定义。校验自定义后端API的路径或名称是否已存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckLivedataApisV2(request *model.CheckLivedataApisV2Request) (*model.CheckLivedataApisV2Response, error) {
	requestDef := GenReqDefForCheckLivedataApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckLivedataApisV2Response), nil
	}
}

// 设备数量统计
//
// 设备数量统计
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CountDevices(request *model.CountDevicesRequest) (*model.CountDevicesResponse, error) {
	requestDef := GenReqDefForCountDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountDevicesResponse), nil
	}
}

// 统计不同类型不同状态任务数量
//
// 统计不同类型不同状态任务数量
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CountTasks(request *model.CountTasksRequest) (*model.CountTasksResponse, error) {
	requestDef := GenReqDefForCountTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTasksResponse), nil
	}
}

// 自动生成APP Code
//
// 创建App Code时，可以不指定具体值，由后台自动生成随机字符串填充。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateAppCodeAutoV2(request *model.CreateAppCodeAutoV2Request) (*model.CreateAppCodeAutoV2Response, error) {
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
func (c *RomaClient) CreateAppCodeV2(request *model.CreateAppCodeV2Request) (*model.CreateAppCodeV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeV2Response), nil
	}
}

// 创建应用配置
//
// 创建应用配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateAppConfigV2(request *model.CreateAppConfigV2Request) (*model.CreateAppConfigV2Response, error) {
	requestDef := GenReqDefForCreateAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppConfigV2Response), nil
	}
}

// 创建客户端配额
//
// 创建客户端配额
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateAppQuota(request *model.CreateAppQuotaRequest) (*model.CreateAppQuotaResponse, error) {
	requestDef := GenReqDefForCreateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppQuotaResponse), nil
	}
}

// 创建命令
//
// 创建命令
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateCommand(request *model.CreateCommandRequest) (*model.CreateCommandResponse, error) {
	requestDef := GenReqDefForCreateCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommandResponse), nil
	}
}

// 创建普通任务
//
// 创建普通任务(区别于组合任务)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateCommonTask(request *model.CreateCommonTaskRequest) (*model.CreateCommonTaskResponse, error) {
	requestDef := GenReqDefForCreateCommonTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommonTaskResponse), nil
	}
}

// 创建自定义认证
//
// 创建自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateCustomAuthorizerV2(request *model.CreateCustomAuthorizerV2Request) (*model.CreateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomAuthorizerV2Response), nil
	}
}

// 创建数据源
//
// 创建数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDatasourceInfo(request *model.CreateDatasourceInfoRequest) (*model.CreateDatasourceInfoResponse, error) {
	requestDef := GenReqDefForCreateDatasourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatasourceInfoResponse), nil
	}
}

// 添加目标数据源
//
// 添加目标数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDestination(request *model.CreateDestinationRequest) (*model.CreateDestinationResponse, error) {
	requestDef := GenReqDefForCreateDestination()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDestinationResponse), nil
	}
}

// 创建设备
//
// 创建设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDevice(request *model.CreateDeviceRequest) (*model.CreateDeviceResponse, error) {
	requestDef := GenReqDefForCreateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceResponse), nil
	}
}

// 创建设备分组
//
// 创建设备分组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDeviceGroup(request *model.CreateDeviceGroupRequest) (*model.CreateDeviceGroupResponse, error) {
	requestDef := GenReqDefForCreateDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceGroupResponse), nil
	}
}

// 创建调度计划
//
// 创建调度计划
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDispatches(request *model.CreateDispatchesRequest) (*model.CreateDispatchesResponse, error) {
	requestDef := GenReqDefForCreateDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDispatchesResponse), nil
	}
}

// 创建环境
//
// 在实际的生产中，API提供者可能有多个环境，如开发环境、测试环境、生产环境等，用户可以自由将API发布到某个环境，供调用者调用。
//
// 对于不同的环境，API的版本、请求地址甚至于包括请求消息等均有可能不同。如：某个API，v1.0的版本为稳定版本，发布到了生产环境供生产使用，同时，该API正处于迭代中，v1.1的版本是开发人员交付测试人员进行测试的版本，发布在测试环境上，而v1.2的版本目前开发团队正处于开发过程中，可以发布到开发环境进行自测等。
//
// 为此，ROMA Connect APIC提供多环境管理功能，使租户能够最大化的模拟实际场景，低成本的接入ROMA Connect APIC。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateEnvironmentV2(request *model.CreateEnvironmentV2Request) (*model.CreateEnvironmentV2Response, error) {
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
// 用户可以定义不同的环境变量，用户在定义API时，在API的定义中使用这些变量，当调用API时，ROMA Connect APIC会将这些变量替换成真实的变量值，以达到不同环境的区分效果。
//
// 环境变量定义在API分组上，该分组下的所有API都可以使用这些变量。
// &gt; 1.环境变量的变量名称必须保持唯一，即一个分组在同一个环境上不能有两个同名的变量
//   2.环境变量区分大小写，即变量ABC与变量abc是两个不同的变量
//   3.设置了环境变量后，使用到该变量的API的调试功能将不可使用。
//   4.定义了环境变量后，使用到环境变量的地方应该以对称的#标识环境变量，当API发布到相应的环境后，会对环境变量的值进行替换，如：定义的API的URL为：https://#address#:8080，环境变量address在RELEASE环境上的值为：192.168.1.5，则API发布到RELEASE环境后的真实的URL为：https://192.168.1.5:8080。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateEnvironmentVariableV2(request *model.CreateEnvironmentVariableV2Request) (*model.CreateEnvironmentVariableV2Response, error) {
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
// 支持配置的特性列表及特性配置请参考“附录 &gt; 实例支持的APIC特性”
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateFeatureV2(request *model.CreateFeatureV2Request) (*model.CreateFeatureV2Response, error) {
	requestDef := GenReqDefForCreateFeatureV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFeatureV2Response), nil
	}
}

// 创建后端API脚本
//
// 在某个实例中创建后端API脚本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateLiveDataApiScriptV2(request *model.CreateLiveDataApiScriptV2Request) (*model.CreateLiveDataApiScriptV2Response, error) {
	requestDef := GenReqDefForCreateLiveDataApiScriptV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLiveDataApiScriptV2Response), nil
	}
}

// 创建后端API
//
// 在某个实例中创建后端API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateLiveDataApiV2(request *model.CreateLiveDataApiV2Request) (*model.CreateLiveDataApiV2Response, error) {
	requestDef := GenReqDefForCreateLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLiveDataApiV2Response), nil
	}
}

// 创建Topic
//
// 创建Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateMqsInstanceTopic(request *model.CreateMqsInstanceTopicRequest) (*model.CreateMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForCreateMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMqsInstanceTopicResponse), nil
	}
}

// 创建组合任务映射
//
// 创建组合任务映射
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateMultiTaskMappings(request *model.CreateMultiTaskMappingsRequest) (*model.CreateMultiTaskMappingsResponse, error) {
	requestDef := GenReqDefForCreateMultiTaskMappings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiTaskMappingsResponse), nil
	}
}

// 创建组合任务
//
// 创建组合任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateMultiTasks(request *model.CreateMultiTasksRequest) (*model.CreateMultiTasksResponse, error) {
	requestDef := GenReqDefForCreateMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiTasksResponse), nil
	}
}

// 创建订阅管理
//
// 该接口用于创建指定实例下对应的应用下的设备操作，订阅到指定的topic
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateNotification(request *model.CreateNotificationRequest) (*model.CreateNotificationResponse, error) {
	requestDef := GenReqDefForCreateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotificationResponse), nil
	}
}

// 创建产品
//
// 创建产品
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProduct(request *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	requestDef := GenReqDefForCreateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductResponse), nil
	}
}

// 创建产品模板
//
// 创建产品模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProductTemplate(request *model.CreateProductTemplateRequest) (*model.CreateProductTemplateResponse, error) {
	requestDef := GenReqDefForCreateProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductTemplateResponse), nil
	}
}

// 添加产品主题
//
// 添加产品主题
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProductTopic(request *model.CreateProductTopicRequest) (*model.CreateProductTopicResponse, error) {
	requestDef := GenReqDefForCreateProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductTopicResponse), nil
	}
}

// 创建属性
//
// 创建属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProperty(request *model.CreatePropertyRequest) (*model.CreatePropertyResponse, error) {
	requestDef := GenReqDefForCreateProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePropertyResponse), nil
	}
}

// 创建请求属性
//
// 创建请求属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateRequestProperty(request *model.CreateRequestPropertyRequest) (*model.CreateRequestPropertyResponse, error) {
	requestDef := GenReqDefForCreateRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestPropertyResponse), nil
	}
}

// 创建流控策略
//
// 当API上线后，系统会默认给每个API提供一个流控策略，API提供者可以根据自身API的服务能力及负载情况变更这个流控策略。 流控策略即限制API在一定长度的时间内，能够允许被访问的最大次数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateRequestThrottlingPolicyV2(request *model.CreateRequestThrottlingPolicyV2Request) (*model.CreateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestThrottlingPolicyV2Response), nil
	}
}

// 创建响应属性
//
// 创建响应属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateResponseProperty(request *model.CreateResponsePropertyRequest) (*model.CreateResponsePropertyResponse, error) {
	requestDef := GenReqDefForCreateResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResponsePropertyResponse), nil
	}
}

// 创建规则
//
// 创建规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

// 创建服务
//
// 创建服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

// 创建签名密钥
//
// 为了保护API的安全性，建议租户为API的访问提供一套保护机制，即租户开放的API，需要对请求来源进行认证，不符合认证的请求直接拒绝访问。
//
// 其中，签名密钥就是API安全保护机制的一种。
//
// 租户创建一个签名密钥，并将签名密钥与API进行绑定，则ROMA Connect APIC在请求这个API时，就会使用绑定的签名密钥对请求参数进行数据加密，生成签名。当租户的后端服务收到请求时，可以校验这个签名，如果签名校验不通过，则该请求不是ROMA Connect APIC发出的请求，租户可以拒绝这个请求，从而保证API的安全性，避免API被未知来源的请求攻击。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateSignatureKeyV2(request *model.CreateSignatureKeyV2Request) (*model.CreateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForCreateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSignatureKeyV2Response), nil
	}
}

// 添加源数据源
//
// 添加源数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateSource(request *model.CreateSourceRequest) (*model.CreateSourceResponse, error) {
	requestDef := GenReqDefForCreateSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSourceResponse), nil
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
func (c *RomaClient) CreateSpecialThrottlingConfigurationV2(request *model.CreateSpecialThrottlingConfigurationV2Request) (*model.CreateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpecialThrottlingConfigurationV2Response), nil
	}
}

// 测试后端API
//
// 测试后端API是否可用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DebugLiveDataApiV2(request *model.DebugLiveDataApiV2Request) (*model.DebugLiveDataApiV2Response, error) {
	requestDef := GenReqDefForDebugLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugLiveDataApiV2Response), nil
	}
}

// 规则调试
//
// 规则调试
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DebugRule(request *model.DebugRuleRequest) (*model.DebugRuleResponse, error) {
	requestDef := GenReqDefForDebugRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugRuleResponse), nil
	}
}

// 删除APP的访问控制
//
// 删除客户端配置的访问控制信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteAppAcl(request *model.DeleteAppAclRequest) (*model.DeleteAppAclResponse, error) {
	requestDef := GenReqDefForDeleteAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppAclResponse), nil
	}
}

// 删除APP Code
//
// 删除App Code，App Code删除后，将无法再通过简易认证访问对应的API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteAppCodeV2(request *model.DeleteAppCodeV2Request) (*model.DeleteAppCodeV2Response, error) {
	requestDef := GenReqDefForDeleteAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppCodeV2Response), nil
	}
}

// 删除应用配置
//
// 删除应用配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteAppConfigV2(request *model.DeleteAppConfigV2Request) (*model.DeleteAppConfigV2Response, error) {
	requestDef := GenReqDefForDeleteAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppConfigV2Response), nil
	}
}

// 删除客户端配额
//
// 删除客户端配额。删除客户端配额时，同时删除客户端配额和客户端应用的关联关系
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteAppQuota(request *model.DeleteAppQuotaRequest) (*model.DeleteAppQuotaResponse, error) {
	requestDef := GenReqDefForDeleteAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppQuotaResponse), nil
	}
}

// 删除命令
//
// 删除命令
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteCommand(request *model.DeleteCommandRequest) (*model.DeleteCommandResponse, error) {
	requestDef := GenReqDefForDeleteCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCommandResponse), nil
	}
}

// 删除自定义认证
//
// 删除自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteCustomAuthorizerV2(request *model.DeleteCustomAuthorizerV2Request) (*model.DeleteCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomAuthorizerV2Response), nil
	}
}

// 通过数据源Id删除指定数据源信息
//
// 通过数据源Id删除指定数据源信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDatasourceInfoById(request *model.DeleteDatasourceInfoByIdRequest) (*model.DeleteDatasourceInfoByIdResponse, error) {
	requestDef := GenReqDefForDeleteDatasourceInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatasourceInfoByIdResponse), nil
	}
}

// 删除目标数据源
//
// 删除目标数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDestination(request *model.DeleteDestinationRequest) (*model.DeleteDestinationResponse, error) {
	requestDef := GenReqDefForDeleteDestination()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDestinationResponse), nil
	}
}

// 删除设备
//
// 删除指定设备ID的设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

// 删除设备分组内的设备
//
// 删除设备分组内的设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDeviceFromGroup(request *model.DeleteDeviceFromGroupRequest) (*model.DeleteDeviceFromGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeviceFromGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceFromGroupResponse), nil
	}
}

// 删除设备分组
//
// 删除分组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDeviceGroup(request *model.DeleteDeviceGroupRequest) (*model.DeleteDeviceGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceGroupResponse), nil
	}
}

// 删除环境
//
// 删除指定的环境。
// 该操作将导致此API在指定的环境无法被访问，可能会影响相当一部分应用和用户。请确保已经告知用户，或者确认需要强制下线。环境上存在已发布的API时，该环境不能被删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteEnvironmentV2(request *model.DeleteEnvironmentV2Request) (*model.DeleteEnvironmentV2Response, error) {
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
func (c *RomaClient) DeleteEnvironmentVariableV2(request *model.DeleteEnvironmentVariableV2Request) (*model.DeleteEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentVariableV2Response), nil
	}
}

// 删除后端API
//
// 在某个实例中删除后端API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteLiveDataApiV2(request *model.DeleteLiveDataApiV2Request) (*model.DeleteLiveDataApiV2Response, error) {
	requestDef := GenReqDefForDeleteLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLiveDataApiV2Response), nil
	}
}

// 删除Topic
//
// 删除Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteMqsInstanceTopic(request *model.DeleteMqsInstanceTopicRequest) (*model.DeleteMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForDeleteMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMqsInstanceTopicResponse), nil
	}
}

// 删除指定任务映射
//
// 通过映射ID删除指定任务映射
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteMultiTaskMapping(request *model.DeleteMultiTaskMappingRequest) (*model.DeleteMultiTaskMappingResponse, error) {
	requestDef := GenReqDefForDeleteMultiTaskMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMultiTaskMappingResponse), nil
	}
}

// 删除订阅管理
//
// 该接口用于删除指定订阅管理
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteNotification(request *model.DeleteNotificationRequest) (*model.DeleteNotificationResponse, error) {
	requestDef := GenReqDefForDeleteNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotificationResponse), nil
	}
}

// 删除产品
//
// 删除产品
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteProduct(request *model.DeleteProductRequest) (*model.DeleteProductResponse, error) {
	requestDef := GenReqDefForDeleteProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductResponse), nil
	}
}

// 删除产品模板
//
// 删除产品模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteProductTemplate(request *model.DeleteProductTemplateRequest) (*model.DeleteProductTemplateResponse, error) {
	requestDef := GenReqDefForDeleteProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductTemplateResponse), nil
	}
}

// 删除产品主题
//
// 删除产品主题
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteProductTopic(request *model.DeleteProductTopicRequest) (*model.DeleteProductTopicResponse, error) {
	requestDef := GenReqDefForDeleteProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductTopicResponse), nil
	}
}

// 删除服务属性
//
// 删除服务属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteProperty(request *model.DeletePropertyRequest) (*model.DeletePropertyResponse, error) {
	requestDef := GenReqDefForDeleteProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePropertyResponse), nil
	}
}

// 删除请求属性
//
// 删除请求属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteRequestProperty(request *model.DeleteRequestPropertyRequest) (*model.DeleteRequestPropertyResponse, error) {
	requestDef := GenReqDefForDeleteRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestPropertyResponse), nil
	}
}

// 删除流控策略
//
// 删除指定的流控策略。当该流控策略绑定了API时，需要先解除流控策略与API的所有绑定关系后再删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteRequestThrottlingPolicyV2(request *model.DeleteRequestThrottlingPolicyV2Request) (*model.DeleteRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestThrottlingPolicyV2Response), nil
	}
}

// 删除响应属性
//
// 删除响应属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteResponseProperty(request *model.DeleteResponsePropertyRequest) (*model.DeleteResponsePropertyResponse, error) {
	requestDef := GenReqDefForDeleteResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResponsePropertyResponse), nil
	}
}

// 删除规则
//
// 删除规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// 删除服务
//
// 删除服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

// 删除签名密钥
//
// 删除指定的签名密钥。签名密钥绑定了API时无法删除，需要先解除与API的绑定关系后删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteSignatureKeyV2(request *model.DeleteSignatureKeyV2Request) (*model.DeleteSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDeleteSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSignatureKeyV2Response), nil
	}
}

// 删除源数据源
//
// 删除源数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteSource(request *model.DeleteSourceRequest) (*model.DeleteSourceResponse, error) {
	requestDef := GenReqDefForDeleteSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSourceResponse), nil
	}
}

// 删除特殊设置
//
// 删除某个流控策略的某个特殊配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteSpecialThrottlingConfigurationV2(request *model.DeleteSpecialThrottlingConfigurationV2Request) (*model.DeleteSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSpecialThrottlingConfigurationV2Response), nil
	}
}

// 通过任务ID删除指定任务
//
// 通过任务ID删除指定任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// 解除客户端配额和客户端应用的绑定
//
// 解除客户端配额和客户端应用的绑定
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DisassociateAppQuotaWithApp(request *model.DisassociateAppQuotaWithAppRequest) (*model.DisassociateAppQuotaWithAppResponse, error) {
	requestDef := GenReqDefForDisassociateAppQuotaWithApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAppQuotaWithAppResponse), nil
	}
}

// 删除域名证书
//
// 如果域名证书不再需要或者已过期，则可以删除证书内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DisassociateCertificateV2(request *model.DisassociateCertificateV2Request) (*model.DisassociateCertificateV2Response, error) {
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
func (c *RomaClient) DisassociateDomainV2(request *model.DisassociateDomainV2Request) (*model.DisassociateDomainV2Response, error) {
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
func (c *RomaClient) DisassociateSignatureKeyV2(request *model.DisassociateSignatureKeyV2Request) (*model.DisassociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateSignatureKeyV2Response), nil
	}
}

// 导出产品
//
// 导出产品
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DownloadProducts(request *model.DownloadProductsRequest) (*model.DownloadProductsResponse, error) {
	requestDef := GenReqDefForDownloadProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadProductsResponse), nil
	}
}

// 导出Topic
//
// 导出Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ExportMqsInstanceTopic(request *model.ExportMqsInstanceTopicRequest) (*model.ExportMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForExportMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportMqsInstanceTopicResponse), nil
	}
}

// 导入Topic
//
// 导入Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ImportMqsInstanceTopic(request *model.ImportMqsInstanceTopicRequest) (*model.ImportMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForImportMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportMqsInstanceTopicResponse), nil
	}
}

// 组合任务初始化
//
// 初始化组合任务，分配任务ID，初始化映射等
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) InstallMultiTasks(request *model.InstallMultiTasksRequest) (*model.InstallMultiTasksResponse, error) {
	requestDef := GenReqDefForInstallMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallMultiTasksResponse), nil
	}
}

// 查看签名密钥绑定的API列表
//
// 查询某个签名密钥上已经绑定的API列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToSignatureKeyV2(request *model.ListApisBindedToSignatureKeyV2Request) (*model.ListApisBindedToSignatureKeyV2Response, error) {
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
func (c *RomaClient) ListApisNotBoundWithSignatureKeyV2(request *model.ListApisNotBoundWithSignatureKeyV2Request) (*model.ListApisNotBoundWithSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisNotBoundWithSignatureKeyV2Response), nil
	}
}

// 查询APP Code列表
//
// 查询App Code列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppCodesV2(request *model.ListAppCodesV2Request) (*model.ListAppCodesV2Response, error) {
	requestDef := GenReqDefForListAppCodesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppCodesV2Response), nil
	}
}

// 查询应用配置列表
//
// 查询应用配置列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppConfigsV2(request *model.ListAppConfigsV2Request) (*model.ListAppConfigsV2Response, error) {
	requestDef := GenReqDefForListAppConfigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppConfigsV2Response), nil
	}
}

// 查询客户端配额可绑定的客户端应用列表
//
// 查询客户端配额可绑定的客户端应用列表。支持按客户端应用名称模糊搜索
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppQuotaBindableApps(request *model.ListAppQuotaBindableAppsRequest) (*model.ListAppQuotaBindableAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBindableApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBindableAppsResponse), nil
	}
}

// 查询客户端配额已绑定的客户端应用列表
//
// 查询客户端配额已绑定的客户端应用列表。支持按客户端应用名称模糊匹配
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppQuotaBoundApps(request *model.ListAppQuotaBoundAppsRequest) (*model.ListAppQuotaBoundAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBoundApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBoundAppsResponse), nil
	}
}

// 获取客户端配额列表
//
// 获取客户端配额列表。支持根据名称模糊查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppQuotas(request *model.ListAppQuotasRequest) (*model.ListAppQuotasResponse, error) {
	requestDef := GenReqDefForListAppQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotasResponse), nil
	}
}

// 查询APP列表
//
// 查询APP列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppsV2(request *model.ListAppsV2Request) (*model.ListAppsV2Response, error) {
	requestDef := GenReqDefForListAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsV2Response), nil
	}
}

// 查询命令
//
// 查询命令
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListCommands(request *model.ListCommandsRequest) (*model.ListCommandsResponse, error) {
	requestDef := GenReqDefForListCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommandsResponse), nil
	}
}

// 查询自定义认证列表
//
// 查询自定义认证列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListCustomAuthorizersV2(request *model.ListCustomAuthorizersV2Request) (*model.ListCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForListCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomAuthorizersV2Response), nil
	}
}

// 获取数据源中某个表中所有字段
//
// 获取数据源中中某个表中所有字段
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDatasourceColumns(request *model.ListDatasourceColumnsRequest) (*model.ListDatasourceColumnsResponse, error) {
	requestDef := GenReqDefForListDatasourceColumns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourceColumnsResponse), nil
	}
}

// 获取数据源中所有的表
//
// 获取数据源中所有的表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDatasourceTables(request *model.ListDatasourceTablesRequest) (*model.ListDatasourceTablesResponse, error) {
	requestDef := GenReqDefForListDatasourceTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourceTablesResponse), nil
	}
}

// 查询数据源
//
// 查询数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDatasources(request *model.ListDatasourcesRequest) (*model.ListDatasourcesResponse, error) {
	requestDef := GenReqDefForListDatasources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourcesResponse), nil
	}
}

// 查询目标数据源列表
//
// 查询目标数据源列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDestinations(request *model.ListDestinationsRequest) (*model.ListDestinationsResponse, error) {
	requestDef := GenReqDefForListDestinations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDestinationsResponse), nil
	}
}

// 查询设备
//
// 查询设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

// 查询产品内设备数量
//
// 查询产品内设备数量
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDevicesInProduct(request *model.ListDevicesInProductRequest) (*model.ListDevicesInProductResponse, error) {
	requestDef := GenReqDefForListDevicesInProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesInProductResponse), nil
	}
}

// 查询变量列表
//
// 查询分组下的所有环境变量的列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListEnvironmentVariablesV2(request *model.ListEnvironmentVariablesV2Request) (*model.ListEnvironmentVariablesV2Response, error) {
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
func (c *RomaClient) ListEnvironmentsV2(request *model.ListEnvironmentsV2Request) (*model.ListEnvironmentsV2Response, error) {
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
// 支持配置的特性列表及特性配置请参考“附录 &gt; 实例支持的APIC特性”
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListFeaturesV2(request *model.ListFeaturesV2Request) (*model.ListFeaturesV2Response, error) {
	requestDef := GenReqDefForListFeaturesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesV2Response), nil
	}
}

// API指标统计值查询-最近一段时间
//
// 根据API的id和最近的一段时间查询API被调用的次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLatelyApiStatisticsV2(request *model.ListLatelyApiStatisticsV2Request) (*model.ListLatelyApiStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyApiStatisticsV2Response), nil
	}
}

// 查询后端API部署历史
//
// 在某个实例中查询后端API的部署记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiDeploymentHistoryV2(request *model.ListLiveDataApiDeploymentHistoryV2Request) (*model.ListLiveDataApiDeploymentHistoryV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiDeploymentHistoryV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiDeploymentHistoryV2Response), nil
	}
}

// 查询后端API测试结果
//
// 在某个实例中查询后端API的测试结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiTestHistoryV2(request *model.ListLiveDataApiTestHistoryV2Request) (*model.ListLiveDataApiTestHistoryV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiTestHistoryV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiTestHistoryV2Response), nil
	}
}

// 查询后端API列表
//
// 获取某个实例下的所有后端API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiV2(request *model.ListLiveDataApiV2Request) (*model.ListLiveDataApiV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiV2Response), nil
	}
}

// 查询自定义后端服务数据源列表
//
// 查询自定义后端服务数据源列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLiveDataDataSourcesV2(request *model.ListLiveDataDataSourcesV2Request) (*model.ListLiveDataDataSourcesV2Response, error) {
	requestDef := GenReqDefForListLiveDataDataSourcesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataDataSourcesV2Response), nil
	}
}

// 查询自定义后端服务配额
//
// 查询自定义后端服务配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListLiveDataQuotaV2(request *model.ListLiveDataQuotaV2Request) (*model.ListLiveDataQuotaV2Response, error) {
	requestDef := GenReqDefForListLiveDataQuotaV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataQuotaV2Response), nil
	}
}

// 任务监控信息列表查询
//
// 查询所有任务的监控信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListMonitorInfos(request *model.ListMonitorInfosRequest) (*model.ListMonitorInfosResponse, error) {
	requestDef := GenReqDefForListMonitorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorInfosResponse), nil
	}
}

// 任务监控日志查询
//
// 查询单个任务的所有日志信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListMonitorLog(request *model.ListMonitorLogRequest) (*model.ListMonitorLogResponse, error) {
	requestDef := GenReqDefForListMonitorLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorLogResponse), nil
	}
}

// 查询Topic列表
//
// 查询Topic列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListMqsInstanceTopics(request *model.ListMqsInstanceTopicsRequest) (*model.ListMqsInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListMqsInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMqsInstanceTopicsResponse), nil
	}
}

// 查询订阅管理信息
//
// 该接口用于查询指定应用订阅管理信息的数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListNotification(request *model.ListNotificationRequest) (*model.ListNotificationResponse, error) {
	requestDef := GenReqDefForListNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationResponse), nil
	}
}

// 查询产品模板
//
// 查询产品模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProductTemplates(request *model.ListProductTemplatesRequest) (*model.ListProductTemplatesResponse, error) {
	requestDef := GenReqDefForListProductTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTemplatesResponse), nil
	}
}

// 查询产品主题
//
// 查询产品主题
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProductTopics(request *model.ListProductTopicsRequest) (*model.ListProductTopicsResponse, error) {
	requestDef := GenReqDefForListProductTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTopicsResponse), nil
	}
}

// 查询产品
//
// 查询产品
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// 查询某个实例的租户配置列表
//
// 查询某个实例的租户配置列表，用户可以通过此接口查看各类型资源配置及使用情况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProjectCofigsV2(request *model.ListProjectCofigsV2Request) (*model.ListProjectCofigsV2Response, error) {
	requestDef := GenReqDefForListProjectCofigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectCofigsV2Response), nil
	}
}

// 查询属性
//
// 查询属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProperties(request *model.ListPropertiesRequest) (*model.ListPropertiesResponse, error) {
	requestDef := GenReqDefForListProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropertiesResponse), nil
	}
}

// 查询请求属性
//
// 查询请求属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListRequestProperties(request *model.ListRequestPropertiesRequest) (*model.ListRequestPropertiesResponse, error) {
	requestDef := GenReqDefForListRequestProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestPropertiesResponse), nil
	}
}

// 查询流控策略列表
//
// 查询所有流控策略的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListRequestThrottlingPolicyV2(request *model.ListRequestThrottlingPolicyV2Request) (*model.ListRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPolicyV2Response), nil
	}
}

// 查询响应属性
//
// 查询响应属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListResponseProperties(request *model.ListResponsePropertiesRequest) (*model.ListResponsePropertiesResponse, error) {
	requestDef := GenReqDefForListResponseProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResponsePropertiesResponse), nil
	}
}

// 查询规则
//
// 查询规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

// 查询服务
//
// 查询服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListServices(request *model.ListServicesRequest) (*model.ListServicesResponse, error) {
	requestDef := GenReqDefForListServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesResponse), nil
	}
}

// 查询设备影子
//
// 查询设备影子
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListShadows(request *model.ListShadowsRequest) (*model.ListShadowsResponse, error) {
	requestDef := GenReqDefForListShadows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShadowsResponse), nil
	}
}

// 查看API绑定的签名密钥列表
//
// 查询某个API绑定的签名密钥列表。每个API在每个环境上应该最多只会绑定一个签名密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListSignatureKeysBindedToApiV2(request *model.ListSignatureKeysBindedToApiV2Request) (*model.ListSignatureKeysBindedToApiV2Response, error) {
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
func (c *RomaClient) ListSignatureKeysV2(request *model.ListSignatureKeysV2Request) (*model.ListSignatureKeysV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysV2Response), nil
	}
}

// 查询源数据源列表
//
// 查询源数据源列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListSources(request *model.ListSourcesRequest) (*model.ListSourcesResponse, error) {
	requestDef := GenReqDefForListSources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSourcesResponse), nil
	}
}

// 查看特殊设置列表
//
// 查看给流控策略设置的特殊配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListSpecialThrottlingConfigurationsV2(request *model.ListSpecialThrottlingConfigurationsV2Request) (*model.ListSpecialThrottlingConfigurationsV2Response, error) {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecialThrottlingConfigurationsV2Response), nil
	}
}

// 查询API指标统计值
//
// 查询某个实例下的API统计信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListStatisticsApi(request *model.ListStatisticsApiRequest) (*model.ListStatisticsApiResponse, error) {
	requestDef := GenReqDefForListStatisticsApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsApiResponse), nil
	}
}

// 查询子设备
//
// 查询子设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListSubsets(request *model.ListSubsetsRequest) (*model.ListSubsetsResponse, error) {
	requestDef := GenReqDefForListSubsets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubsetsResponse), nil
	}
}

// 查询标签列表
//
// 查询标签列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListTagsV2(request *model.ListTagsV2Request) (*model.ListTagsV2Response, error) {
	requestDef := GenReqDefForListTagsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsV2Response), nil
	}
}

// 查询任务列表
//
// 查询任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// 查询设备主题
//
// 查询设备主题
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListTopics(request *model.ListTopicsRequest) (*model.ListTopicsResponse, error) {
	requestDef := GenReqDefForListTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicsResponse), nil
	}
}

// 部署后端API
//
// 在某个实例中部署后端API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) PublishLiveDataApiV2(request *model.PublishLiveDataApiV2Request) (*model.PublishLiveDataApiV2Response, error) {
	requestDef := GenReqDefForPublishLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishLiveDataApiV2Response), nil
	}
}

// 重置设备鉴权信息
//
// 重置设备鉴权信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ResetAuthentication(request *model.ResetAuthenticationRequest) (*model.ResetAuthenticationResponse, error) {
	requestDef := GenReqDefForResetAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetAuthenticationResponse), nil
	}
}

// 重发消息
//
// 重发消息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ResetMessages(request *model.ResetMessagesRequest) (*model.ResetMessagesResponse, error) {
	requestDef := GenReqDefForResetMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMessagesResponse), nil
	}
}

// 重置组合任务进度
//
// 重置组合任务进度
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ResetMultiTaskOffset(request *model.ResetMultiTaskOffsetRequest) (*model.ResetMultiTaskOffsetResponse, error) {
	requestDef := GenReqDefForResetMultiTaskOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMultiTaskOffsetResponse), nil
	}
}

// 重置产品鉴权信息
//
// 重置产品鉴权信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ResetProductAuthentication(request *model.ResetProductAuthenticationRequest) (*model.ResetProductAuthenticationResponse, error) {
	requestDef := GenReqDefForResetProductAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetProductAuthenticationResponse), nil
	}
}

// 手工触发单个任务
//
// 手工触发一次任务调度
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) RunTask(request *model.RunTaskRequest) (*model.RunTaskResponse, error) {
	requestDef := GenReqDefForRunTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTaskResponse), nil
	}
}

// 发送命令
//
// 发送命令
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) SendCommand(request *model.SendCommandRequest) (*model.SendCommandResponse, error) {
	requestDef := GenReqDefForSendCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendCommandResponse), nil
	}
}

// 查询客户端应用关联的应用配额
//
// 查看指定客户端应用关联的应用配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowAppBoundAppQuota(request *model.ShowAppBoundAppQuotaRequest) (*model.ShowAppBoundAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppBoundAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppBoundAppQuotaResponse), nil
	}
}

// 获取客户端配额详情
//
// 获取客户端配额详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowAppQuota(request *model.ShowAppQuotaRequest) (*model.ShowAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppQuotaResponse), nil
	}
}

// 查询设备鉴权信息
//
// 查询设备鉴权信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowAuthentication(request *model.ShowAuthenticationRequest) (*model.ShowAuthenticationResponse, error) {
	requestDef := GenReqDefForShowAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthenticationResponse), nil
	}
}

// 查询命令详情
//
// 查询命令详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowCommand(request *model.ShowCommandRequest) (*model.ShowCommandResponse, error) {
	requestDef := GenReqDefForShowCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommandResponse), nil
	}
}

// 查询指定数据源
//
// 根据数据源id查询数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDataourceDetail(request *model.ShowDataourceDetailRequest) (*model.ShowDataourceDetailResponse, error) {
	requestDef := GenReqDefForShowDataourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataourceDetailResponse), nil
	}
}

// 查看APP的访问控制详情
//
// 查看APP的访问控制详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppAcl(request *model.ShowDetailsOfAppAclRequest) (*model.ShowDetailsOfAppAclResponse, error) {
	requestDef := GenReqDefForShowDetailsOfAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppAclResponse), nil
	}
}

// 查看APP Code详情
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppCodeV2(request *model.ShowDetailsOfAppCodeV2Request) (*model.ShowDetailsOfAppCodeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppCodeV2Response), nil
	}
}

// 查看应用配置详情
//
// 查看应用配置详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppConfigV2(request *model.ShowDetailsOfAppConfigV2Request) (*model.ShowDetailsOfAppConfigV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppConfigV2Response), nil
	}
}

// 查看APP详情
//
// 查看指定APP的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppV2(request *model.ShowDetailsOfAppV2Request) (*model.ShowDetailsOfAppV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppV2Response), nil
	}
}

// 查看自定义认证详情
//
// 查看自定义认证详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfCustomAuthorizersV2(request *model.ShowDetailsOfCustomAuthorizersV2Request) (*model.ShowDetailsOfCustomAuthorizersV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfDomainNameCertificateV2(request *model.ShowDetailsOfDomainNameCertificateV2Request) (*model.ShowDetailsOfDomainNameCertificateV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfEnvironmentVariableV2(request *model.ShowDetailsOfEnvironmentVariableV2Request) (*model.ShowDetailsOfEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfEnvironmentVariableV2Response), nil
	}
}

// 查看ROMA Connect实例详情
//
// 查看ROMA Connect实例详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfInstanceV2(request *model.ShowDetailsOfInstanceV2Request) (*model.ShowDetailsOfInstanceV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfRequestThrottlingPolicyV2(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfRequestThrottlingPolicyV2Response), nil
	}
}

// 查询设备详情
//
// 查询设备详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDevice(request *model.ShowDeviceRequest) (*model.ShowDeviceResponse, error) {
	requestDef := GenReqDefForShowDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceResponse), nil
	}
}

// 查询设备分组详情
//
// 获取设备分组及下一层分组信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDeviceGroup(request *model.ShowDeviceGroupRequest) (*model.ShowDeviceGroupResponse, error) {
	requestDef := GenReqDefForShowDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceGroupResponse), nil
	}
}

// 查询所有设备分组
//
// 查询所有设备分组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDeviceGroupTree(request *model.ShowDeviceGroupTreeRequest) (*model.ShowDeviceGroupTreeResponse, error) {
	requestDef := GenReqDefForShowDeviceGroupTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceGroupTreeResponse), nil
	}
}

// 查询设备分组内设备
//
// 查询设备分组内设备
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDevicesInGroup(request *model.ShowDevicesInGroupRequest) (*model.ShowDevicesInGroupResponse, error) {
	requestDef := GenReqDefForShowDevicesInGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDevicesInGroupResponse), nil
	}
}

// 查询调度计划
//
// 查询调度计划
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDispatches(request *model.ShowDispatchesRequest) (*model.ShowDispatchesResponse, error) {
	requestDef := GenReqDefForShowDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDispatchesResponse), nil
	}
}

// 查询后端API详情
//
// 查询后端API的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowLiveDataApiV2(request *model.ShowLiveDataApiV2Request) (*model.ShowLiveDataApiV2Response, error) {
	requestDef := GenReqDefForShowLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLiveDataApiV2Response), nil
	}
}

// 查询消息
//
// 查询消息的偏移量和消息内容。
// 先根据时间戳查询消息的偏移量，再根据偏移量查询消息内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstanceMessages(request *model.ShowMqsInstanceMessagesRequest) (*model.ShowMqsInstanceMessagesResponse, error) {
	requestDef := GenReqDefForShowMqsInstanceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceMessagesResponse), nil
	}
}

// 查询Topic权限
//
// 查询Topic权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstanceTopicAccessPolicy(request *model.ShowMqsInstanceTopicAccessPolicyRequest) (*model.ShowMqsInstanceTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowMqsInstanceTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceTopicAccessPolicyResponse), nil
	}
}

// 查询产品详情
//
// 查询产品详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowProduct(request *model.ShowProductRequest) (*model.ShowProductResponse, error) {
	requestDef := GenReqDefForShowProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductResponse), nil
	}
}

// 查询产品鉴权信息
//
// 查询产品鉴权信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowProductAuthentication(request *model.ShowProductAuthenticationRequest) (*model.ShowProductAuthenticationResponse, error) {
	requestDef := GenReqDefForShowProductAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductAuthenticationResponse), nil
	}
}

// 查询产品模板详情
//
// 查询产品模板详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowProductTemplate(request *model.ShowProductTemplateRequest) (*model.ShowProductTemplateResponse, error) {
	requestDef := GenReqDefForShowProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductTemplateResponse), nil
	}
}

// 查询服务属性详情
//
// 查询服务属性详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowProperty(request *model.ShowPropertyRequest) (*model.ShowPropertyResponse, error) {
	requestDef := GenReqDefForShowProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyResponse), nil
	}
}

// 查询请求属性详情
//
// 查询请求属性详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowRequestProperty(request *model.ShowRequestPropertyRequest) (*model.ShowRequestPropertyResponse, error) {
	requestDef := GenReqDefForShowRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRequestPropertyResponse), nil
	}
}

// 查询响应属性详情
//
// 查询响应属性详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowResponseProperty(request *model.ShowResponsePropertyRequest) (*model.ShowResponsePropertyResponse, error) {
	requestDef := GenReqDefForShowResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResponsePropertyResponse), nil
	}
}

// 查看ROMA Connect实例约束信息
//
// 查看ROMA Connect实例约束信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowRestrictionOfInstanceV2(request *model.ShowRestrictionOfInstanceV2Request) (*model.ShowRestrictionOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowRestrictionOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRestrictionOfInstanceV2Response), nil
	}
}

// 查询规则详情
//
// 查询规则详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowRule(request *model.ShowRuleRequest) (*model.ShowRuleResponse, error) {
	requestDef := GenReqDefForShowRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleResponse), nil
	}
}

// 查询服务详情
//
// 查询服务详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowService(request *model.ShowServiceRequest) (*model.ShowServiceResponse, error) {
	requestDef := GenReqDefForShowService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceResponse), nil
	}
}

// 通过任务ID查询指定任务的信息
//
// 通过任务ID查询指定任务的信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// 测试数据源连通性
//
// 测试数据源连通性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) StartTestDatasource(request *model.StartTestDatasourceRequest) (*model.StartTestDatasourceResponse, error) {
	requestDef := GenReqDefForStartTestDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartTestDatasourceResponse), nil
	}
}

// 手工停止当前执行的任务
//
// 手工停止当前执行的任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) StopTask(request *model.StopTaskRequest) (*model.StopTaskResponse, error) {
	requestDef := GenReqDefForStopTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskResponse), nil
	}
}

// 撤销后端API
//
// 在某个实例中取消部署后端API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UnpublishLiveDataApiV2(request *model.UnpublishLiveDataApiV2Request) (*model.UnpublishLiveDataApiV2Response, error) {
	requestDef := GenReqDefForUnpublishLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnpublishLiveDataApiV2Response), nil
	}
}

// 设置APP的访问控制
//
// 设置客户端配置的访问控制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateAppAcl(request *model.UpdateAppAclRequest) (*model.UpdateAppAclResponse, error) {
	requestDef := GenReqDefForUpdateAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppAclResponse), nil
	}
}

// 修改应用配置
//
// 修改应用配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateAppConfigV2(request *model.UpdateAppConfigV2Request) (*model.UpdateAppConfigV2Response, error) {
	requestDef := GenReqDefForUpdateAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppConfigV2Response), nil
	}
}

// 修改客户端配额
//
// 修改客户端配额
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateAppQuota(request *model.UpdateAppQuotaRequest) (*model.UpdateAppQuotaResponse, error) {
	requestDef := GenReqDefForUpdateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppQuotaResponse), nil
	}
}

// 修改命令
//
// 修改命令
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateCommand(request *model.UpdateCommandRequest) (*model.UpdateCommandResponse, error) {
	requestDef := GenReqDefForUpdateCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCommandResponse), nil
	}
}

// 修改自定义认证
//
// 修改自定义认证
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateCustomAuthorizerV2(request *model.UpdateCustomAuthorizerV2Request) (*model.UpdateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomAuthorizerV2Response), nil
	}
}

// 修改数据源
//
// 修改数据源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDatasourceInfo(request *model.UpdateDatasourceInfoRequest) (*model.UpdateDatasourceInfoResponse, error) {
	requestDef := GenReqDefForUpdateDatasourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatasourceInfoResponse), nil
	}
}

// 修改设备
//
// 修改设备信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// 修改设备分组
//
// 修改设备分组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDeviceGroup(request *model.UpdateDeviceGroupRequest) (*model.UpdateDeviceGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceGroupResponse), nil
	}
}

// 修改调度计划
//
// 通过任务ID和调度ID修改调度计划
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDispatches(request *model.UpdateDispatchesRequest) (*model.UpdateDispatchesResponse, error) {
	requestDef := GenReqDefForUpdateDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDispatchesResponse), nil
	}
}

// 修改域名
//
// 修改绑定的域名所对应的配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDomainV2(request *model.UpdateDomainV2Request) (*model.UpdateDomainV2Response, error) {
	requestDef := GenReqDefForUpdateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainV2Response), nil
	}
}

// 修改环境
//
// 修改指定环境的信息。其中可修改的属性为：name、remark，其它属性不可修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateEnvironmentV2(request *model.UpdateEnvironmentV2Request) (*model.UpdateEnvironmentV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentV2Response), nil
	}
}

// 修改变量
//
// 修改环境变量。环境变量引用位置为api的后端服务地址时，修改对应环境变量会将使用该变量的所有api重新发布。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateEnvironmentVariableV2(request *model.UpdateEnvironmentVariableV2Request) (*model.UpdateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentVariableV2Response), nil
	}
}

// 修改后端API
//
// 在某个实例中更新后端API的参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateLiveDataApiV2(request *model.UpdateLiveDataApiV2Request) (*model.UpdateLiveDataApiV2Response, error) {
	requestDef := GenReqDefForUpdateLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLiveDataApiV2Response), nil
	}
}

// 修改Topic
//
// 修改Topic。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateMqsInstanceTopic(request *model.UpdateMqsInstanceTopicRequest) (*model.UpdateMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForUpdateMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMqsInstanceTopicResponse), nil
	}
}

// 修改组合任务
//
// 修改组合任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateMultiTasks(request *model.UpdateMultiTasksRequest) (*model.UpdateMultiTasksResponse, error) {
	requestDef := GenReqDefForUpdateMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMultiTasksResponse), nil
	}
}

// 修改订阅管理
//
// 该接口用于修改指定的订阅管理
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateNotification(request *model.UpdateNotificationRequest) (*model.UpdateNotificationResponse, error) {
	requestDef := GenReqDefForUpdateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationResponse), nil
	}
}

// 修改产品信息
//
// 修改产品信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateProduct(request *model.UpdateProductRequest) (*model.UpdateProductResponse, error) {
	requestDef := GenReqDefForUpdateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductResponse), nil
	}
}

// 修改产品模板
//
// 修改产品模板
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateProductTemplate(request *model.UpdateProductTemplateRequest) (*model.UpdateProductTemplateResponse, error) {
	requestDef := GenReqDefForUpdateProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductTemplateResponse), nil
	}
}

// 更新产品主题
//
// 更新产品主题
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateProductTopic(request *model.UpdateProductTopicRequest) (*model.UpdateProductTopicResponse, error) {
	requestDef := GenReqDefForUpdateProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductTopicResponse), nil
	}
}

// 修改服务属性
//
// 修改服务属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateProperty(request *model.UpdatePropertyRequest) (*model.UpdatePropertyResponse, error) {
	requestDef := GenReqDefForUpdateProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePropertyResponse), nil
	}
}

// 修改请求属性
//
// 修改请求属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateRequestProperty(request *model.UpdateRequestPropertyRequest) (*model.UpdateRequestPropertyResponse, error) {
	requestDef := GenReqDefForUpdateRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestPropertyResponse), nil
	}
}

// 修改流控策略
//
// 修改指定流控策略的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateRequestThrottlingPolicyV2(request *model.UpdateRequestThrottlingPolicyV2Request) (*model.UpdateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestThrottlingPolicyV2Response), nil
	}
}

// 修改响应属性
//
// 修改响应属性
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateResponseProperty(request *model.UpdateResponsePropertyRequest) (*model.UpdateResponsePropertyResponse, error) {
	requestDef := GenReqDefForUpdateResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResponsePropertyResponse), nil
	}
}

// 修改规则
//
// 修改规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
	requestDef := GenReqDefForUpdateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleResponse), nil
	}
}

// 修改服务
//
// 修改服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

// 修改签名密钥
//
// 修改指定签名密钥的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateSignatureKeyV2(request *model.UpdateSignatureKeyV2Request) (*model.UpdateSignatureKeyV2Response, error) {
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
func (c *RomaClient) UpdateSpecialThrottlingConfigurationV2(request *model.UpdateSpecialThrottlingConfigurationV2Request) (*model.UpdateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpecialThrottlingConfigurationV2Response), nil
	}
}

// 更新普通任务
//
// 更新普通任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

// 更新Topic权限
//
// 更新Topic权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateTopicAccessPolicy(request *model.UpdateTopicAccessPolicyRequest) (*model.UpdateTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

// 导入产品
//
// 导入产品
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UploadProduct(request *model.UploadProductRequest) (*model.UploadProductResponse, error) {
	requestDef := GenReqDefForUploadProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadProductResponse), nil
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
func (c *RomaClient) BatchDeleteAclV2(request *model.BatchDeleteAclV2Request) (*model.BatchDeleteAclV2Response, error) {
	requestDef := GenReqDefForBatchDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclV2Response), nil
	}
}

// 创建ACL策略
//
// 增加一个ACL策略，策略类型通过字段acl_type来确定（permit或者deny），限制的对象的类型可以为IP[或者DOMAIN，这里的DOMAIN对应的acl_value的值为租户名称，而非“www.exampleDomain.com\&quot;之类的网络域名。](tag:hws;hws_hk;hcs;fcs;g42;)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateAclStrategyV2(request *model.CreateAclStrategyV2Request) (*model.CreateAclStrategyV2Response, error) {
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
func (c *RomaClient) DeleteAclV2(request *model.DeleteAclV2Request) (*model.DeleteAclV2Response, error) {
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
func (c *RomaClient) ListAclStrategiesV2(request *model.ListAclStrategiesV2Request) (*model.ListAclStrategiesV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfAclPolicyV2(request *model.ShowDetailsOfAclPolicyV2Request) (*model.ShowDetailsOfAclPolicyV2Response, error) {
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
func (c *RomaClient) UpdateAclStrategyV2(request *model.UpdateAclStrategyV2Request) (*model.UpdateAclStrategyV2Response, error) {
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
func (c *RomaClient) AssociateRequestThrottlingPolicyV2(request *model.AssociateRequestThrottlingPolicyV2Request) (*model.AssociateRequestThrottlingPolicyV2Response, error) {
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
func (c *RomaClient) BatchDisassociateThrottlingPolicyV2(request *model.BatchDisassociateThrottlingPolicyV2Request) (*model.BatchDisassociateThrottlingPolicyV2Response, error) {
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
// 注意：当action &#x3D; online时，接口返回的响应中publish_id，version_id， publish_time字段才有含义。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchPublishOrOfflineApiV2(request *model.BatchPublishOrOfflineApiV2Request) (*model.BatchPublishOrOfflineApiV2Response, error) {
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
func (c *RomaClient) ChangeApiVersionV2(request *model.ChangeApiVersionV2Request) (*model.ChangeApiVersionV2Response, error) {
	requestDef := GenReqDefForChangeApiVersionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApiVersionV2Response), nil
	}
}

// 校验API分组名称是否存在
//
// 校验API分组名称是否存在。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckApiGroupsV2(request *model.CheckApiGroupsV2Request) (*model.CheckApiGroupsV2Response, error) {
	requestDef := GenReqDefForCheckApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApiGroupsV2Response), nil
	}
}

// 校验API定义
//
// 校验API定义。校验API的路径或名称是否已存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckApisV2(request *model.CheckApisV2Request) (*model.CheckApisV2Response, error) {
	requestDef := GenReqDefForCheckApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApisV2Response), nil
	}
}

// 创建API分组
//
// API分组是API的管理单元，一个API分组等同于一个服务入口，创建API分组时，返回一个子域名作为访问入口。建议一个API分组下的API具有一定的相关性。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateApiGroupV2(request *model.CreateApiGroupV2Request) (*model.CreateApiGroupV2Response, error) {
	requestDef := GenReqDefForCreateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiGroupV2Response), nil
	}
}

// 创建API
//
// 添加一个API，API即一个服务接口，具体的服务能力。
//
// API分为两部分，第一部分为面向API使用者的API接口，定义了使用者如何调用这个API。第二部分面向API提供者，由API提供者定义这个API的真实的后端情况，定义了ROMA Connect如何去访问真实的后端服务。API的真实后端服务目前支持三种类型：传统的HTTP/HTTPS形式的web后端、[函数工作流、](tag:hws;hws_hk;hcs;fcs;g42;)MOCK。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateApiV2(request *model.CreateApiV2Request) (*model.CreateApiV2Response, error) {
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
func (c *RomaClient) CreateOrDeletePublishRecordForApiV2(request *model.CreateOrDeletePublishRecordForApiV2Request) (*model.CreateOrDeletePublishRecordForApiV2Response, error) {
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
func (c *RomaClient) DebugApiV2(request *model.DebugApiV2Request) (*model.DebugApiV2Response, error) {
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
func (c *RomaClient) DeleteApiByVersionIdV2(request *model.DeleteApiByVersionIdV2Request) (*model.DeleteApiByVersionIdV2Response, error) {
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
// 分组下存在API时分组无法删除，需要删除所有分组下的API后，再删除分组。
// 删除分组时，会一并删除直接或间接关联到该分组下的所有资源，包括独立域名、SSL证书等等。并会将外部域名与子域名的绑定关系进行解除（取决于域名cname方式）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteApiGroupV2(request *model.DeleteApiGroupV2Request) (*model.DeleteApiGroupV2Response, error) {
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
func (c *RomaClient) DeleteApiV2(request *model.DeleteApiV2Request) (*model.DeleteApiV2Response, error) {
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
func (c *RomaClient) DisassociateRequestThrottlingPolicyV2(request *model.DisassociateRequestThrottlingPolicyV2Request) (*model.DisassociateRequestThrottlingPolicyV2Response, error) {
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
func (c *RomaClient) ListApiGroupsV2(request *model.ListApiGroupsV2Request) (*model.ListApiGroupsV2Response, error) {
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
func (c *RomaClient) ListApiRuntimeDefinitionV2(request *model.ListApiRuntimeDefinitionV2Request) (*model.ListApiRuntimeDefinitionV2Response, error) {
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
func (c *RomaClient) ListApiVersionDetailV2(request *model.ListApiVersionDetailV2Request) (*model.ListApiVersionDetailV2Response, error) {
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
func (c *RomaClient) ListApiVersionsV2(request *model.ListApiVersionsV2Request) (*model.ListApiVersionsV2Response, error) {
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
func (c *RomaClient) ListApisBindedToRequestThrottlingPolicyV2(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) (*model.ListApisBindedToRequestThrottlingPolicyV2Response, error) {
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
func (c *RomaClient) ListApisUnbindedToRequestThrottlingPolicyV2(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, error) {
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
func (c *RomaClient) ListApisV2(request *model.ListApisV2Request) (*model.ListApisV2Response, error) {
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
func (c *RomaClient) ListRequestThrottlingPoliciesBindedToApiV2(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfApiGroupV2(request *model.ShowDetailsOfApiGroupV2Request) (*model.ShowDetailsOfApiGroupV2Response, error) {
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
func (c *RomaClient) ShowDetailsOfApiV2(request *model.ShowDetailsOfApiV2Request) (*model.ShowDetailsOfApiV2Response, error) {
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
func (c *RomaClient) UpdateApiGroupV2(request *model.UpdateApiGroupV2Request) (*model.UpdateApiGroupV2Response, error) {
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
func (c *RomaClient) UpdateApiV2(request *model.UpdateApiV2Request) (*model.UpdateApiV2Response, error) {
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
func (c *RomaClient) BatchDeleteApiAclBindingV2(request *model.BatchDeleteApiAclBindingV2Request) (*model.BatchDeleteApiAclBindingV2Response, error) {
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
func (c *RomaClient) CreateApiAclBindingV2(request *model.CreateApiAclBindingV2Request) (*model.CreateApiAclBindingV2Response, error) {
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
func (c *RomaClient) DeleteApiAclBindingV2(request *model.DeleteApiAclBindingV2Request) (*model.DeleteApiAclBindingV2Response, error) {
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
func (c *RomaClient) ListAclPolicyBindedToApiV2(request *model.ListAclPolicyBindedToApiV2Request) (*model.ListAclPolicyBindedToApiV2Response, error) {
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
func (c *RomaClient) ListApisBindedToAclPolicyV2(request *model.ListApisBindedToAclPolicyV2Request) (*model.ListApisBindedToAclPolicyV2Response, error) {
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
func (c *RomaClient) ListApisUnbindedToAclPolicyV2(request *model.ListApisUnbindedToAclPolicyV2Request) (*model.ListApisUnbindedToAclPolicyV2Response, error) {
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
func (c *RomaClient) CancelingAuthorizationV2(request *model.CancelingAuthorizationV2Request) (*model.CancelingAuthorizationV2Response, error) {
	requestDef := GenReqDefForCancelingAuthorizationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelingAuthorizationV2Response), nil
	}
}

// APP授权
//
// APP创建成功后，还不能访问API，如果想要访问某个环境上的API，需要将该API在该环境上授权给APP。授权成功后，APP即可访问该环境上的这个API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateAuthorizingAppsV2(request *model.CreateAuthorizingAppsV2Request) (*model.CreateAuthorizingAppsV2Response, error) {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizingAppsV2Response), nil
	}
}

// 查看APP已绑定的API列表
//
// 查询APP已经绑定的API列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToAppV2(request *model.ListApisBindedToAppV2Request) (*model.ListApisBindedToAppV2Response, error) {
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
func (c *RomaClient) ListApisUnbindedToAppV2(request *model.ListApisUnbindedToAppV2Request) (*model.ListApisUnbindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAppV2Response), nil
	}
}

// 查看API已绑定的APP列表
//
// 查询API绑定的APP列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListAppsBindedToApiV2(request *model.ListAppsBindedToApiV2Request) (*model.ListAppsBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAppsBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsBindedToApiV2Response), nil
	}
}

// 查看APP下路径冲突的api列表
//
// 查询指定APP下路径冲突的api列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDuplicateApisForAppV2(request *model.ListDuplicateApisForAppV2Request) (*model.ListDuplicateApisForAppV2Response, error) {
	requestDef := GenReqDefForListDuplicateApisForAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDuplicateApisForAppV2Response), nil
	}
}

// 设置用户成员
//
// - 设置应用的用户成员，为空数组时会清空已有应用成员列表
// - 设置动作为全量更新非增量更新，应用的成员列表都会替换为当次请求的应用成员列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AddUserToApp(request *model.AddUserToAppRequest) (*model.AddUserToAppResponse, error) {
	requestDef := GenReqDefForAddUserToApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddUserToAppResponse), nil
	}
}

// 查询用户成员列表
//
// 查询用户成列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckAuthUsersOfApp(request *model.CheckAuthUsersOfAppRequest) (*model.CheckAuthUsersOfAppResponse, error) {
	requestDef := GenReqDefForCheckAuthUsersOfApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAuthUsersOfAppResponse), nil
	}
}

// 查询候选用户成员
//
// 查询应用的候选用户成员列表,会过滤掉异常状态用户
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckCanAuthUsersOfApp(request *model.CheckCanAuthUsersOfAppRequest) (*model.CheckCanAuthUsersOfAppResponse, error) {
	requestDef := GenReqDefForCheckCanAuthUsersOfApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCanAuthUsersOfAppResponse), nil
	}
}

// 查询应用详情
//
// 查询应用详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckRomaAppDetails(request *model.CheckRomaAppDetailsRequest) (*model.CheckRomaAppDetailsResponse, error) {
	requestDef := GenReqDefForCheckRomaAppDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaAppDetailsResponse), nil
	}
}

// 查询应用密钥
//
// 查询应用密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckRomaAppSecret(request *model.CheckRomaAppSecretRequest) (*model.CheckRomaAppSecretResponse, error) {
	requestDef := GenReqDefForCheckRomaAppSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaAppSecretResponse), nil
	}
}

// 创建应用
//
// 创建应用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateRomaApp(request *model.CreateRomaAppRequest) (*model.CreateRomaAppResponse, error) {
	requestDef := GenReqDefForCreateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRomaAppResponse), nil
	}
}

// 删除应用
//
// 删除单个应用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteRomaApp(request *model.DeleteRomaAppRequest) (*model.DeleteRomaAppResponse, error) {
	requestDef := GenReqDefForDeleteRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRomaAppResponse), nil
	}
}

// 查询应用列表
//
// 查询应用列表，支持条件查询，所有条件是并且的关系
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListRomaApp(request *model.ListRomaAppRequest) (*model.ListRomaAppResponse, error) {
	requestDef := GenReqDefForListRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRomaAppResponse), nil
	}
}

// 重置应用密钥
//
// 重置应用密钥
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ResetRomaAppSecret(request *model.ResetRomaAppSecretRequest) (*model.ResetRomaAppSecretResponse, error) {
	requestDef := GenReqDefForResetRomaAppSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetRomaAppSecretResponse), nil
	}
}

// 更新应用
//
// 更新应用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateRomaApp(request *model.UpdateRomaAppRequest) (*model.UpdateRomaAppResponse, error) {
	requestDef := GenReqDefForUpdateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRomaAppResponse), nil
	}
}

// 校验应用是否存在
//
// 校验指定条件的应用是否存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ValidateRomaApp(request *model.ValidateRomaAppRequest) (*model.ValidateRomaAppResponse, error) {
	requestDef := GenReqDefForValidateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateRomaAppResponse), nil
	}
}

// 查询作业进度
//
// 查询作业进度
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckAssetJobStatus(request *model.CheckAssetJobStatusRequest) (*model.CheckAssetJobStatusResponse, error) {
	requestDef := GenReqDefForCheckAssetJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAssetJobStatusResponse), nil
	}
}

// 批量删除资产
//
// 批量删除资产
// - 创建批量删除指定条件的资产的作业任务
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产删除作业任务，没有Running状态的作业任务存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// 下载资产包
//
// - 导出作业执行成功后，通过该接口获取导出作业产生的资产包，仅能下载一次
// - 可先压缩后存在数据库，下载后删除
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DownloadAssetArchive(request *model.DownloadAssetArchiveRequest) (*model.DownloadAssetArchiveResponse, error) {
	requestDef := GenReqDefForDownloadAssetArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAssetArchiveResponse), nil
	}
}

// 批量导出资产
//
// 批量导出资产
// - 创建批量导出指定条件的资产的作业任务
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产导出作业任务，没有Running状态的作业任务存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ExportAsset(request *model.ExportAssetRequest) (*model.ExportAssetResponse, error) {
	requestDef := GenReqDefForExportAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportAssetResponse), nil
	}
}

// 导入资产
//
// - 创建导入资产作业任务，资产版本和具体哪些资产从资产内容里读取
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产导入作业任务，没有Running状态的作业任务存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ImportAsset(request *model.ImportAssetRequest) (*model.ImportAssetResponse, error) {
	requestDef := GenReqDefForImportAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportAssetResponse), nil
	}
}

// 查询字典详情
//
// 查询字典详情,
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckDictionary(request *model.CheckDictionaryRequest) (*model.CheckDictionaryResponse, error) {
	requestDef := GenReqDefForCheckDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDictionaryResponse), nil
	}
}

// 创建字典
//
// 创建字典
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateDictionary(request *model.CreateDictionaryRequest) (*model.CreateDictionaryResponse, error) {
	requestDef := GenReqDefForCreateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDictionaryResponse), nil
	}
}

// 删除字典
//
// 删除单个字典，会同时删除该字典的所有子字典
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteDictionary(request *model.DeleteDictionaryRequest) (*model.DeleteDictionaryResponse, error) {
	requestDef := GenReqDefForDeleteDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDictionaryResponse), nil
	}
}

// 查询字典列表
//
// 查询字典列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListDictionary(request *model.ListDictionaryRequest) (*model.ListDictionaryResponse, error) {
	requestDef := GenReqDefForListDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDictionaryResponse), nil
	}
}

// 更新字典
//
// 更新字典
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateDictionary(request *model.UpdateDictionaryRequest) (*model.UpdateDictionaryResponse, error) {
	requestDef := GenReqDefForUpdateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDictionaryResponse), nil
	}
}

// 校验字典是否存在
//
// 校验指定条件的字典是否存在，支持字典名称和字典编码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ValidateDictionary(request *model.ValidateDictionaryRequest) (*model.ValidateDictionaryResponse, error) {
	requestDef := GenReqDefForValidateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateDictionaryResponse), nil
	}
}

// 查询实例列表
//
// 获取符合条件的服务实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CheckRomaInstanceListV2(request *model.CheckRomaInstanceListV2Request) (*model.CheckRomaInstanceListV2Response, error) {
	requestDef := GenReqDefForCheckRomaInstanceListV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaInstanceListV2Response), nil
	}
}

// 查询MQS实例列表
//
// 查询MQS实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListMqsInstance(request *model.ListMqsInstanceRequest) (*model.ListMqsInstanceResponse, error) {
	requestDef := GenReqDefForListMqsInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMqsInstanceResponse), nil
	}
}

// 查询MQS实例详情
//
// 查询指定MQS实例详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstance(request *model.ShowMqsInstanceRequest) (*model.ShowMqsInstanceResponse, error) {
	requestDef := GenReqDefForShowMqsInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceResponse), nil
	}
}

// 导出API
//
// 导出分组下API的定义信息，导出文件内容符合swagger标准规范。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ExportApiDefinitionsV2(request *model.ExportApiDefinitionsV2Request) (*model.ExportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportApiDefinitionsV2Response), nil
	}
}

// 导出自定义后端API
//
// 导出自定义后端API，导出文件内容符合swagger标准规范。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ExportLiveDataApiDefinitionsV2(request *model.ExportLiveDataApiDefinitionsV2Request) (*model.ExportLiveDataApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportLiveDataApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportLiveDataApiDefinitionsV2Response), nil
	}
}

// 导入API
//
// 导入API。导入文件内容需要符合swagger标准规范，自定义扩展字段请参考用户指南的“附录：前端API的Swagger扩展定义”章节。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ImportApiDefinitionsV2(request *model.ImportApiDefinitionsV2Request) (*model.ImportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApiDefinitionsV2Response), nil
	}
}

// 导入自定义后端API
//
// 导入自定义后端API。导入文件内容需要符合swagger标准规范，自定义扩展字段请参考用户指南的“附录：后端API的Swagger扩展定义”章节
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ImportLiveDataApiDefinitionsV2(request *model.ImportLiveDataApiDefinitionsV2Request) (*model.ImportLiveDataApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportLiveDataApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportLiveDataApiDefinitionsV2Response), nil
	}
}

// 添加或更新后端实例
//
// 为指定的VPC通道添加后端实例
//
// 若指定地址的后端实例已存在，则更新对应后端实例信息。若请求体中包含多个重复地址的后端实例定义，则使用第一个定义。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) AddingBackendInstancesV2(request *model.AddingBackendInstancesV2Request) (*model.AddingBackendInstancesV2Response, error) {
	requestDef := GenReqDefForAddingBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddingBackendInstancesV2Response), nil
	}
}

// 批量修改后端服务器状态不可用
//
// 批量修改后端服务器状态不可用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchDisableMembers(request *model.BatchDisableMembersRequest) (*model.BatchDisableMembersResponse, error) {
	requestDef := GenReqDefForBatchDisableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableMembersResponse), nil
	}
}

// 批量修改后端服务器状态可用
//
// 批量修改后端服务器状态可用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) BatchEnableMembers(request *model.BatchEnableMembersRequest) (*model.BatchEnableMembersResponse, error) {
	requestDef := GenReqDefForBatchEnableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableMembersResponse), nil
	}
}

// 添加或更新VPC通道后端服务器组
//
// 在ROMA Connect APIC中创建VPC通道后端服务器组，VPC通道后端实例可以选择是否关联后端实例服务器组，以便管理后端服务器节点。
//
// 若指定名称的后端服务器组已存在，则更新对应后端服务器组信息。若请求体中包含多个重复名称的后端服务器定义，则使用第一个定义。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateMemberGroup(request *model.CreateMemberGroupRequest) (*model.CreateMemberGroupResponse, error) {
	requestDef := GenReqDefForCreateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberGroupResponse), nil
	}
}

// 项目下创建VPC通道
//
// 创建相同的VPC通道关联到多个实例。同一个项目下VPC通道名称不可重复。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProjectVpcChannel(request *model.CreateProjectVpcChannelRequest) (*model.CreateProjectVpcChannelResponse, error) {
	requestDef := GenReqDefForCreateProjectVpcChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectVpcChannelResponse), nil
	}
}

// 项目下同步VPC通道
//
// 同步VPC通道到多个实例。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateProjectVpcChannelSyncs(request *model.CreateProjectVpcChannelSyncsRequest) (*model.CreateProjectVpcChannelSyncsResponse, error) {
	requestDef := GenReqDefForCreateProjectVpcChannelSyncs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectVpcChannelSyncsResponse), nil
	}
}

// 创建VPC通道
//
// 在ROMA Connect APIC中创建连接私有VPC资源的通道，并在创建API时将后端节点配置为使用这些VPC通道，以便ROMA Connect APIC直接访问私有VPC资源。
// &gt; 每个用户默认最多创建200个VPC通道，如需支持更多请联系技术支持调整配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) CreateVpcChannelV2(request *model.CreateVpcChannelV2Request) (*model.CreateVpcChannelV2Response, error) {
	requestDef := GenReqDefForCreateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcChannelV2Response), nil
	}
}

// 删除后端实例
//
// 删除指定VPC通道中的后端实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteBackendInstanceV2(request *model.DeleteBackendInstanceV2Request) (*model.DeleteBackendInstanceV2Response, error) {
	requestDef := GenReqDefForDeleteBackendInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackendInstanceV2Response), nil
	}
}

// 删除VPC通道后端服务器组
//
// 删除指定的VPC通道后端服务器组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteMemberGroup(request *model.DeleteMemberGroupRequest) (*model.DeleteMemberGroupResponse, error) {
	requestDef := GenReqDefForDeleteMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberGroupResponse), nil
	}
}

// 删除VPC通道
//
// 删除指定的VPC通道
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) DeleteVpcChannelV2(request *model.DeleteVpcChannelV2Request) (*model.DeleteVpcChannelV2Response, error) {
	requestDef := GenReqDefForDeleteVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcChannelV2Response), nil
	}
}

// 查看后端实例列表
//
// 查看指定VPC通道的后端实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListBackendInstancesV2(request *model.ListBackendInstancesV2Request) (*model.ListBackendInstancesV2Response, error) {
	requestDef := GenReqDefForListBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackendInstancesV2Response), nil
	}
}

// 查询VPC通道后端云服务组列表
//
// 查询VPC通道后端云服务组列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListMemberGroups(request *model.ListMemberGroupsRequest) (*model.ListMemberGroupsResponse, error) {
	requestDef := GenReqDefForListMemberGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMemberGroupsResponse), nil
	}
}

// 查询项目下所有实例的VPC通道列表
//
// 查询项目下所有实例的VPC通道列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListProjectVpcChannelsV2(request *model.ListProjectVpcChannelsV2Request) (*model.ListProjectVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListProjectVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectVpcChannelsV2Response), nil
	}
}

// 查询VPC通道列表
//
// 查看VPC通道列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ListVpcChannelsV2(request *model.ListVpcChannelsV2Request) (*model.ListVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcChannelsV2Response), nil
	}
}

// 查看VPC通道后端服务器组详情
//
// 查看指定的VPC通道后端服务器组详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfMemberGroup(request *model.ShowDetailsOfMemberGroupRequest) (*model.ShowDetailsOfMemberGroupResponse, error) {
	requestDef := GenReqDefForShowDetailsOfMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfMemberGroupResponse), nil
	}
}

// 查看VPC通道详情
//
// 查看指定的VPC通道详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfVpcChannelV2(request *model.ShowDetailsOfVpcChannelV2Request) (*model.ShowDetailsOfVpcChannelV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfVpcChannelV2Response), nil
	}
}

// 更新后端实例
//
// 更新指定的VPC通道的后端实例。更新时，使用传入的请求参数对对应云服务组的后端实例进行全量覆盖修改。若未指定修改的云服务器组，则进行全量覆盖。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateBackendInstancesV2(request *model.UpdateBackendInstancesV2Request) (*model.UpdateBackendInstancesV2Response, error) {
	requestDef := GenReqDefForUpdateBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackendInstancesV2Response), nil
	}
}

// 修改VPC通道健康检查
//
// 修改VPC通道健康检查。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateHealthCheck(request *model.UpdateHealthCheckRequest) (*model.UpdateHealthCheckResponse, error) {
	requestDef := GenReqDefForUpdateHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthCheckResponse), nil
	}
}

// 更新VPC通道后端服务器组
//
// 更新指定VPC通道后端服务器组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateMemberGroup(request *model.UpdateMemberGroupRequest) (*model.UpdateMemberGroupResponse, error) {
	requestDef := GenReqDefForUpdateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberGroupResponse), nil
	}
}

// 项目下批量修改VPC通道
//
// 项目下根据VPC通道名称批量修改多个多个实例下的VPC通道。若实例下不存在该VPC通道则创建。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateProjectVpcChannel(request *model.UpdateProjectVpcChannelRequest) (*model.UpdateProjectVpcChannelResponse, error) {
	requestDef := GenReqDefForUpdateProjectVpcChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectVpcChannelResponse), nil
	}
}

// 更新VPC通道
//
// 更新指定VPC通道的参数
//
// 使用传入的后端实例列表对VPC通道进行全量覆盖，若后端实例列表为空，则会全量删除已有的后端实例；
//
// 使用传入的后端服务器组列表对VPC通道进行全量覆盖，若后端服务器组列表为空，则会全量删除已有的服务器组；
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RomaClient) UpdateVpcChannelV2(request *model.UpdateVpcChannelV2Request) (*model.UpdateVpcChannelV2Response, error) {
	requestDef := GenReqDefForUpdateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcChannelV2Response), nil
	}
}
