package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// AddSubsetsToGateway 添加子设备到网关
//
// 添加子设备到网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AddSubsetsToGateway(request *model.AddSubsetsToGatewayRequest) (*model.AddSubsetsToGatewayResponse, error) {
	requestDef := GenReqDefForAddSubsetsToGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubsetsToGatewayResponse), nil
	}
}

// AddSubsetsToGatewayInvoker 添加子设备到网关
func (c *RomaClient) AddSubsetsToGatewayInvoker(request *model.AddSubsetsToGatewayRequest) *AddSubsetsToGatewayInvoker {
	requestDef := GenReqDefForAddSubsetsToGateway()
	return &AddSubsetsToGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateAppsForAppQuota 客户端配额绑定客户端应用列表
//
// 客户端配额绑定客户端应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AssociateAppsForAppQuota(request *model.AssociateAppsForAppQuotaRequest) (*model.AssociateAppsForAppQuotaResponse, error) {
	requestDef := GenReqDefForAssociateAppsForAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateAppsForAppQuotaResponse), nil
	}
}

// AssociateAppsForAppQuotaInvoker 客户端配额绑定客户端应用列表
func (c *RomaClient) AssociateAppsForAppQuotaInvoker(request *model.AssociateAppsForAppQuotaRequest) *AssociateAppsForAppQuotaInvoker {
	requestDef := GenReqDefForAssociateAppsForAppQuota()
	return &AssociateAppsForAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateCertificateV2 绑定域名证书
//
// 如果创建API时，“定义API请求”使用HTTPS请求协议，那么在独立域名中需要添加SSL证书。
// 本章节主要介绍为特定域名绑定证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AssociateCertificateV2(request *model.AssociateCertificateV2Request) (*model.AssociateCertificateV2Response, error) {
	requestDef := GenReqDefForAssociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateCertificateV2Response), nil
	}
}

// AssociateCertificateV2Invoker 绑定域名证书
func (c *RomaClient) AssociateCertificateV2Invoker(request *model.AssociateCertificateV2Request) *AssociateCertificateV2Invoker {
	requestDef := GenReqDefForAssociateCertificateV2()
	return &AssociateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateDomainV2 绑定域名
//
// 用户自定义的域名，需要CNAME到API分组的子域名上才能生效。 每个API分组下最多可绑定5个域名。绑定域名后，用户可通过自定义域名调用API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AssociateDomainV2(request *model.AssociateDomainV2Request) (*model.AssociateDomainV2Response, error) {
	requestDef := GenReqDefForAssociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateDomainV2Response), nil
	}
}

// AssociateDomainV2Invoker 绑定域名
func (c *RomaClient) AssociateDomainV2Invoker(request *model.AssociateDomainV2Request) *AssociateDomainV2Invoker {
	requestDef := GenReqDefForAssociateDomainV2()
	return &AssociateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateSignatureKeyV2 绑定签名密钥
//
// 签名密钥创建后，需要绑定到API才能生效。
//
// 将签名密钥绑定到API后，则ROMA Connect APIC请求后端服务时就会使用这个签名密钥进行加密签名，后端服务可以校验这个签名来验证请求来源。
//
// 将指定的签名密钥绑定到一个或多个已发布的API上。同一个API发布到不同的环境可以绑定不同的签名密钥；一个API在发布到特定环境后只能绑定一个签名密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AssociateSignatureKeyV2(request *model.AssociateSignatureKeyV2Request) (*model.AssociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForAssociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSignatureKeyV2Response), nil
	}
}

// AssociateSignatureKeyV2Invoker 绑定签名密钥
func (c *RomaClient) AssociateSignatureKeyV2Invoker(request *model.AssociateSignatureKeyV2Request) *AssociateSignatureKeyV2Invoker {
	requestDef := GenReqDefForAssociateSignatureKeyV2()
	return &AssociateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddDeviceToGroup 批量添加设备到设备分组
//
// 批量添加设备到设备分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchAddDeviceToGroup(request *model.BatchAddDeviceToGroupRequest) (*model.BatchAddDeviceToGroupResponse, error) {
	requestDef := GenReqDefForBatchAddDeviceToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDeviceToGroupResponse), nil
	}
}

// BatchAddDeviceToGroupInvoker 批量添加设备到设备分组
func (c *RomaClient) BatchAddDeviceToGroupInvoker(request *model.BatchAddDeviceToGroupRequest) *BatchAddDeviceToGroupInvoker {
	requestDef := GenReqDefForBatchAddDeviceToGroup()
	return &BatchAddDeviceToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDevices 批量删除设备
//
// 批量删除设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDeleteDevices(request *model.BatchDeleteDevicesRequest) (*model.BatchDeleteDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDevicesResponse), nil
	}
}

// BatchDeleteDevicesInvoker 批量删除设备
func (c *RomaClient) BatchDeleteDevicesInvoker(request *model.BatchDeleteDevicesRequest) *BatchDeleteDevicesInvoker {
	requestDef := GenReqDefForBatchDeleteDevices()
	return &BatchDeleteDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMqsInstanceTopic 批量删除Topic
//
// 批量删除Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDeleteMqsInstanceTopic(request *model.BatchDeleteMqsInstanceTopicRequest) (*model.BatchDeleteMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForBatchDeleteMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMqsInstanceTopicResponse), nil
	}
}

// BatchDeleteMqsInstanceTopicInvoker 批量删除Topic
func (c *RomaClient) BatchDeleteMqsInstanceTopicInvoker(request *model.BatchDeleteMqsInstanceTopicRequest) *BatchDeleteMqsInstanceTopicInvoker {
	requestDef := GenReqDefForBatchDeleteMqsInstanceTopic()
	return &BatchDeleteMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRules 批量删除规则
//
// 批量删除规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDeleteRules(request *model.BatchDeleteRulesRequest) (*model.BatchDeleteRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRulesResponse), nil
	}
}

// BatchDeleteRulesInvoker 批量删除规则
func (c *RomaClient) BatchDeleteRulesInvoker(request *model.BatchDeleteRulesRequest) *BatchDeleteRulesInvoker {
	requestDef := GenReqDefForBatchDeleteRules()
	return &BatchDeleteRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchFreezeDevices 设备批量下线
//
// 设备批量下线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchFreezeDevices(request *model.BatchFreezeDevicesRequest) (*model.BatchFreezeDevicesResponse, error) {
	requestDef := GenReqDefForBatchFreezeDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchFreezeDevicesResponse), nil
	}
}

// BatchFreezeDevicesInvoker 设备批量下线
func (c *RomaClient) BatchFreezeDevicesInvoker(request *model.BatchFreezeDevicesRequest) *BatchFreezeDevicesInvoker {
	requestDef := GenReqDefForBatchFreezeDevices()
	return &BatchFreezeDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartOrStopTasks 批量启动\\停止任务
//
// 批量启动\\停止任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchStartOrStopTasks(request *model.BatchStartOrStopTasksRequest) (*model.BatchStartOrStopTasksResponse, error) {
	requestDef := GenReqDefForBatchStartOrStopTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartOrStopTasksResponse), nil
	}
}

// BatchStartOrStopTasksInvoker 批量启动\\停止任务
func (c *RomaClient) BatchStartOrStopTasksInvoker(request *model.BatchStartOrStopTasksRequest) *BatchStartOrStopTasksInvoker {
	requestDef := GenReqDefForBatchStartOrStopTasks()
	return &BatchStartOrStopTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckLivedataApisV2 校验自定义后端API定义
//
// 校验自定义后端API定义。校验自定义后端API的路径或名称是否已存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckLivedataApisV2(request *model.CheckLivedataApisV2Request) (*model.CheckLivedataApisV2Response, error) {
	requestDef := GenReqDefForCheckLivedataApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckLivedataApisV2Response), nil
	}
}

// CheckLivedataApisV2Invoker 校验自定义后端API定义
func (c *RomaClient) CheckLivedataApisV2Invoker(request *model.CheckLivedataApisV2Request) *CheckLivedataApisV2Invoker {
	requestDef := GenReqDefForCheckLivedataApisV2()
	return &CheckLivedataApisV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountDevices 设备数量统计
//
// 设备数量统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CountDevices(request *model.CountDevicesRequest) (*model.CountDevicesResponse, error) {
	requestDef := GenReqDefForCountDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountDevicesResponse), nil
	}
}

// CountDevicesInvoker 设备数量统计
func (c *RomaClient) CountDevicesInvoker(request *model.CountDevicesRequest) *CountDevicesInvoker {
	requestDef := GenReqDefForCountDevices()
	return &CountDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountTasks 统计不同类型不同状态任务数量
//
// 统计不同类型不同状态任务数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CountTasks(request *model.CountTasksRequest) (*model.CountTasksResponse, error) {
	requestDef := GenReqDefForCountTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTasksResponse), nil
	}
}

// CountTasksInvoker 统计不同类型不同状态任务数量
func (c *RomaClient) CountTasksInvoker(request *model.CountTasksRequest) *CountTasksInvoker {
	requestDef := GenReqDefForCountTasks()
	return &CountTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppCodeAutoV2 自动生成APP Code
//
// 创建App Code时，可以不指定具体值，由后台自动生成随机字符串填充。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAppCodeAutoV2(request *model.CreateAppCodeAutoV2Request) (*model.CreateAppCodeAutoV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeAutoV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeAutoV2Response), nil
	}
}

// CreateAppCodeAutoV2Invoker 自动生成APP Code
func (c *RomaClient) CreateAppCodeAutoV2Invoker(request *model.CreateAppCodeAutoV2Request) *CreateAppCodeAutoV2Invoker {
	requestDef := GenReqDefForCreateAppCodeAutoV2()
	return &CreateAppCodeAutoV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppCodeV2 创建APP Code
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAppCodeV2(request *model.CreateAppCodeV2Request) (*model.CreateAppCodeV2Response, error) {
	requestDef := GenReqDefForCreateAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppCodeV2Response), nil
	}
}

// CreateAppCodeV2Invoker 创建APP Code
func (c *RomaClient) CreateAppCodeV2Invoker(request *model.CreateAppCodeV2Request) *CreateAppCodeV2Invoker {
	requestDef := GenReqDefForCreateAppCodeV2()
	return &CreateAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppConfigV2 创建应用配置
//
// 创建应用配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAppConfigV2(request *model.CreateAppConfigV2Request) (*model.CreateAppConfigV2Response, error) {
	requestDef := GenReqDefForCreateAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppConfigV2Response), nil
	}
}

// CreateAppConfigV2Invoker 创建应用配置
func (c *RomaClient) CreateAppConfigV2Invoker(request *model.CreateAppConfigV2Request) *CreateAppConfigV2Invoker {
	requestDef := GenReqDefForCreateAppConfigV2()
	return &CreateAppConfigV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppQuota 创建客户端配额
//
// 创建客户端配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAppQuota(request *model.CreateAppQuotaRequest) (*model.CreateAppQuotaResponse, error) {
	requestDef := GenReqDefForCreateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppQuotaResponse), nil
	}
}

// CreateAppQuotaInvoker 创建客户端配额
func (c *RomaClient) CreateAppQuotaInvoker(request *model.CreateAppQuotaRequest) *CreateAppQuotaInvoker {
	requestDef := GenReqDefForCreateAppQuota()
	return &CreateAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCommand 创建命令
//
// 创建命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateCommand(request *model.CreateCommandRequest) (*model.CreateCommandResponse, error) {
	requestDef := GenReqDefForCreateCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommandResponse), nil
	}
}

// CreateCommandInvoker 创建命令
func (c *RomaClient) CreateCommandInvoker(request *model.CreateCommandRequest) *CreateCommandInvoker {
	requestDef := GenReqDefForCreateCommand()
	return &CreateCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCommonTask 创建普通任务
//
// 创建普通任务(区别于组合任务)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateCommonTask(request *model.CreateCommonTaskRequest) (*model.CreateCommonTaskResponse, error) {
	requestDef := GenReqDefForCreateCommonTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommonTaskResponse), nil
	}
}

// CreateCommonTaskInvoker 创建普通任务
func (c *RomaClient) CreateCommonTaskInvoker(request *model.CreateCommonTaskRequest) *CreateCommonTaskInvoker {
	requestDef := GenReqDefForCreateCommonTask()
	return &CreateCommonTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomAuthorizerV2 创建自定义认证
//
// 创建自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateCustomAuthorizerV2(request *model.CreateCustomAuthorizerV2Request) (*model.CreateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomAuthorizerV2Response), nil
	}
}

// CreateCustomAuthorizerV2Invoker 创建自定义认证
func (c *RomaClient) CreateCustomAuthorizerV2Invoker(request *model.CreateCustomAuthorizerV2Request) *CreateCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForCreateCustomAuthorizerV2()
	return &CreateCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatasourceInfo 创建数据源
//
// 创建数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDatasourceInfo(request *model.CreateDatasourceInfoRequest) (*model.CreateDatasourceInfoResponse, error) {
	requestDef := GenReqDefForCreateDatasourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatasourceInfoResponse), nil
	}
}

// CreateDatasourceInfoInvoker 创建数据源
func (c *RomaClient) CreateDatasourceInfoInvoker(request *model.CreateDatasourceInfoRequest) *CreateDatasourceInfoInvoker {
	requestDef := GenReqDefForCreateDatasourceInfo()
	return &CreateDatasourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDestination 添加目标数据源
//
// 添加目标数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDestination(request *model.CreateDestinationRequest) (*model.CreateDestinationResponse, error) {
	requestDef := GenReqDefForCreateDestination()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDestinationResponse), nil
	}
}

// CreateDestinationInvoker 添加目标数据源
func (c *RomaClient) CreateDestinationInvoker(request *model.CreateDestinationRequest) *CreateDestinationInvoker {
	requestDef := GenReqDefForCreateDestination()
	return &CreateDestinationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDevice 创建设备
//
// 创建设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDevice(request *model.CreateDeviceRequest) (*model.CreateDeviceResponse, error) {
	requestDef := GenReqDefForCreateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceResponse), nil
	}
}

// CreateDeviceInvoker 创建设备
func (c *RomaClient) CreateDeviceInvoker(request *model.CreateDeviceRequest) *CreateDeviceInvoker {
	requestDef := GenReqDefForCreateDevice()
	return &CreateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeviceGroup 创建设备分组
//
// 创建设备分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDeviceGroup(request *model.CreateDeviceGroupRequest) (*model.CreateDeviceGroupResponse, error) {
	requestDef := GenReqDefForCreateDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeviceGroupResponse), nil
	}
}

// CreateDeviceGroupInvoker 创建设备分组
func (c *RomaClient) CreateDeviceGroupInvoker(request *model.CreateDeviceGroupRequest) *CreateDeviceGroupInvoker {
	requestDef := GenReqDefForCreateDeviceGroup()
	return &CreateDeviceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDispatches 创建调度计划
//
// 创建调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDispatches(request *model.CreateDispatchesRequest) (*model.CreateDispatchesResponse, error) {
	requestDef := GenReqDefForCreateDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDispatchesResponse), nil
	}
}

// CreateDispatchesInvoker 创建调度计划
func (c *RomaClient) CreateDispatchesInvoker(request *model.CreateDispatchesRequest) *CreateDispatchesInvoker {
	requestDef := GenReqDefForCreateDispatches()
	return &CreateDispatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironmentV2 创建环境
//
// 在实际的生产中，API提供者可能有多个环境，如开发环境、测试环境、生产环境等，用户可以自由将API发布到某个环境，供调用者调用。
//
// 对于不同的环境，API的版本、请求地址甚至于包括请求消息等均有可能不同。如：某个API，v1.0的版本为稳定版本，发布到了生产环境供生产使用，同时，该API正处于迭代中，v1.1的版本是开发人员交付测试人员进行测试的版本，发布在测试环境上，而v1.2的版本目前开发团队正处于开发过程中，可以发布到开发环境进行自测等。
//
// 为此，ROMA Connect APIC提供多环境管理功能，使租户能够最大化的模拟实际场景，低成本的接入ROMA Connect APIC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateEnvironmentV2(request *model.CreateEnvironmentV2Request) (*model.CreateEnvironmentV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentV2Response), nil
	}
}

// CreateEnvironmentV2Invoker 创建环境
func (c *RomaClient) CreateEnvironmentV2Invoker(request *model.CreateEnvironmentV2Request) *CreateEnvironmentV2Invoker {
	requestDef := GenReqDefForCreateEnvironmentV2()
	return &CreateEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironmentVariableV2 新建变量
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
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateEnvironmentVariableV2(request *model.CreateEnvironmentVariableV2Request) (*model.CreateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForCreateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentVariableV2Response), nil
	}
}

// CreateEnvironmentVariableV2Invoker 新建变量
func (c *RomaClient) CreateEnvironmentVariableV2Invoker(request *model.CreateEnvironmentVariableV2Request) *CreateEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForCreateEnvironmentVariableV2()
	return &CreateEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFeatureV2 实例配置特性
//
// 为实例配置需要的特性。
//
// 支持配置的特性列表及特性配置请参考“附录 &gt; 实例支持的APIC特性”
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateFeatureV2(request *model.CreateFeatureV2Request) (*model.CreateFeatureV2Response, error) {
	requestDef := GenReqDefForCreateFeatureV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFeatureV2Response), nil
	}
}

// CreateFeatureV2Invoker 实例配置特性
func (c *RomaClient) CreateFeatureV2Invoker(request *model.CreateFeatureV2Request) *CreateFeatureV2Invoker {
	requestDef := GenReqDefForCreateFeatureV2()
	return &CreateFeatureV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLiveDataApiScriptV2 创建后端API脚本
//
// 在某个实例中创建后端API脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateLiveDataApiScriptV2(request *model.CreateLiveDataApiScriptV2Request) (*model.CreateLiveDataApiScriptV2Response, error) {
	requestDef := GenReqDefForCreateLiveDataApiScriptV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLiveDataApiScriptV2Response), nil
	}
}

// CreateLiveDataApiScriptV2Invoker 创建后端API脚本
func (c *RomaClient) CreateLiveDataApiScriptV2Invoker(request *model.CreateLiveDataApiScriptV2Request) *CreateLiveDataApiScriptV2Invoker {
	requestDef := GenReqDefForCreateLiveDataApiScriptV2()
	return &CreateLiveDataApiScriptV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLiveDataApiV2 创建后端API
//
// 在某个实例中创建后端API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateLiveDataApiV2(request *model.CreateLiveDataApiV2Request) (*model.CreateLiveDataApiV2Response, error) {
	requestDef := GenReqDefForCreateLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLiveDataApiV2Response), nil
	}
}

// CreateLiveDataApiV2Invoker 创建后端API
func (c *RomaClient) CreateLiveDataApiV2Invoker(request *model.CreateLiveDataApiV2Request) *CreateLiveDataApiV2Invoker {
	requestDef := GenReqDefForCreateLiveDataApiV2()
	return &CreateLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMqsInstanceTopic 创建Topic
//
// 创建Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateMqsInstanceTopic(request *model.CreateMqsInstanceTopicRequest) (*model.CreateMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForCreateMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMqsInstanceTopicResponse), nil
	}
}

// CreateMqsInstanceTopicInvoker 创建Topic
func (c *RomaClient) CreateMqsInstanceTopicInvoker(request *model.CreateMqsInstanceTopicRequest) *CreateMqsInstanceTopicInvoker {
	requestDef := GenReqDefForCreateMqsInstanceTopic()
	return &CreateMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMultiTaskMappings 创建组合任务映射
//
// 创建组合任务映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateMultiTaskMappings(request *model.CreateMultiTaskMappingsRequest) (*model.CreateMultiTaskMappingsResponse, error) {
	requestDef := GenReqDefForCreateMultiTaskMappings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiTaskMappingsResponse), nil
	}
}

// CreateMultiTaskMappingsInvoker 创建组合任务映射
func (c *RomaClient) CreateMultiTaskMappingsInvoker(request *model.CreateMultiTaskMappingsRequest) *CreateMultiTaskMappingsInvoker {
	requestDef := GenReqDefForCreateMultiTaskMappings()
	return &CreateMultiTaskMappingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMultiTasks 创建组合任务
//
// 创建组合任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateMultiTasks(request *model.CreateMultiTasksRequest) (*model.CreateMultiTasksResponse, error) {
	requestDef := GenReqDefForCreateMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiTasksResponse), nil
	}
}

// CreateMultiTasksInvoker 创建组合任务
func (c *RomaClient) CreateMultiTasksInvoker(request *model.CreateMultiTasksRequest) *CreateMultiTasksInvoker {
	requestDef := GenReqDefForCreateMultiTasks()
	return &CreateMultiTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotification 创建订阅管理
//
// 该接口用于创建指定实例下对应的应用下的设备操作，订阅到指定的topic
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateNotification(request *model.CreateNotificationRequest) (*model.CreateNotificationResponse, error) {
	requestDef := GenReqDefForCreateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotificationResponse), nil
	}
}

// CreateNotificationInvoker 创建订阅管理
func (c *RomaClient) CreateNotificationInvoker(request *model.CreateNotificationRequest) *CreateNotificationInvoker {
	requestDef := GenReqDefForCreateNotification()
	return &CreateNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProduct 创建产品
//
// 创建产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProduct(request *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	requestDef := GenReqDefForCreateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductResponse), nil
	}
}

// CreateProductInvoker 创建产品
func (c *RomaClient) CreateProductInvoker(request *model.CreateProductRequest) *CreateProductInvoker {
	requestDef := GenReqDefForCreateProduct()
	return &CreateProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProductTemplate 创建产品模板
//
// 创建产品模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProductTemplate(request *model.CreateProductTemplateRequest) (*model.CreateProductTemplateResponse, error) {
	requestDef := GenReqDefForCreateProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductTemplateResponse), nil
	}
}

// CreateProductTemplateInvoker 创建产品模板
func (c *RomaClient) CreateProductTemplateInvoker(request *model.CreateProductTemplateRequest) *CreateProductTemplateInvoker {
	requestDef := GenReqDefForCreateProductTemplate()
	return &CreateProductTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProductTopic 添加产品主题
//
// 添加产品主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProductTopic(request *model.CreateProductTopicRequest) (*model.CreateProductTopicResponse, error) {
	requestDef := GenReqDefForCreateProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductTopicResponse), nil
	}
}

// CreateProductTopicInvoker 添加产品主题
func (c *RomaClient) CreateProductTopicInvoker(request *model.CreateProductTopicRequest) *CreateProductTopicInvoker {
	requestDef := GenReqDefForCreateProductTopic()
	return &CreateProductTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProperty 创建属性
//
// 创建属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProperty(request *model.CreatePropertyRequest) (*model.CreatePropertyResponse, error) {
	requestDef := GenReqDefForCreateProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePropertyResponse), nil
	}
}

// CreatePropertyInvoker 创建属性
func (c *RomaClient) CreatePropertyInvoker(request *model.CreatePropertyRequest) *CreatePropertyInvoker {
	requestDef := GenReqDefForCreateProperty()
	return &CreatePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRequestProperty 创建请求属性
//
// 创建请求属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateRequestProperty(request *model.CreateRequestPropertyRequest) (*model.CreateRequestPropertyResponse, error) {
	requestDef := GenReqDefForCreateRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestPropertyResponse), nil
	}
}

// CreateRequestPropertyInvoker 创建请求属性
func (c *RomaClient) CreateRequestPropertyInvoker(request *model.CreateRequestPropertyRequest) *CreateRequestPropertyInvoker {
	requestDef := GenReqDefForCreateRequestProperty()
	return &CreateRequestPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRequestThrottlingPolicyV2 创建流控策略
//
// 当API上线后，系统会默认给每个API提供一个流控策略，API提供者可以根据自身API的服务能力及负载情况变更这个流控策略。 流控策略即限制API在一定长度的时间内，能够允许被访问的最大次数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateRequestThrottlingPolicyV2(request *model.CreateRequestThrottlingPolicyV2Request) (*model.CreateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestThrottlingPolicyV2Response), nil
	}
}

// CreateRequestThrottlingPolicyV2Invoker 创建流控策略
func (c *RomaClient) CreateRequestThrottlingPolicyV2Invoker(request *model.CreateRequestThrottlingPolicyV2Request) *CreateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForCreateRequestThrottlingPolicyV2()
	return &CreateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResponseProperty 创建响应属性
//
// 创建响应属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateResponseProperty(request *model.CreateResponsePropertyRequest) (*model.CreateResponsePropertyResponse, error) {
	requestDef := GenReqDefForCreateResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResponsePropertyResponse), nil
	}
}

// CreateResponsePropertyInvoker 创建响应属性
func (c *RomaClient) CreateResponsePropertyInvoker(request *model.CreateResponsePropertyRequest) *CreateResponsePropertyInvoker {
	requestDef := GenReqDefForCreateResponseProperty()
	return &CreateResponsePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRule 创建规则
//
// 创建规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

// CreateRuleInvoker 创建规则
func (c *RomaClient) CreateRuleInvoker(request *model.CreateRuleRequest) *CreateRuleInvoker {
	requestDef := GenReqDefForCreateRule()
	return &CreateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateService 创建服务
//
// 创建服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

// CreateServiceInvoker 创建服务
func (c *RomaClient) CreateServiceInvoker(request *model.CreateServiceRequest) *CreateServiceInvoker {
	requestDef := GenReqDefForCreateService()
	return &CreateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSignatureKeyV2 创建签名密钥
//
// 为了保护API的安全性，建议租户为API的访问提供一套保护机制，即租户开放的API，需要对请求来源进行认证，不符合认证的请求直接拒绝访问。
//
// 其中，签名密钥就是API安全保护机制的一种。
//
// 租户创建一个签名密钥，并将签名密钥与API进行绑定，则ROMA Connect APIC在请求这个API时，就会使用绑定的签名密钥对请求参数进行数据加密，生成签名。当租户的后端服务收到请求时，可以校验这个签名，如果签名校验不通过，则该请求不是ROMA Connect APIC发出的请求，租户可以拒绝这个请求，从而保证API的安全性，避免API被未知来源的请求攻击。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateSignatureKeyV2(request *model.CreateSignatureKeyV2Request) (*model.CreateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForCreateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSignatureKeyV2Response), nil
	}
}

// CreateSignatureKeyV2Invoker 创建签名密钥
func (c *RomaClient) CreateSignatureKeyV2Invoker(request *model.CreateSignatureKeyV2Request) *CreateSignatureKeyV2Invoker {
	requestDef := GenReqDefForCreateSignatureKeyV2()
	return &CreateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSource 添加源数据源
//
// 添加源数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateSource(request *model.CreateSourceRequest) (*model.CreateSourceResponse, error) {
	requestDef := GenReqDefForCreateSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSourceResponse), nil
	}
}

// CreateSourceInvoker 添加源数据源
func (c *RomaClient) CreateSourceInvoker(request *model.CreateSourceRequest) *CreateSourceInvoker {
	requestDef := GenReqDefForCreateSource()
	return &CreateSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
func (c *RomaClient) CreateSpecialThrottlingConfigurationV2(request *model.CreateSpecialThrottlingConfigurationV2Request) (*model.CreateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpecialThrottlingConfigurationV2Response), nil
	}
}

// CreateSpecialThrottlingConfigurationV2Invoker 创建特殊设置
func (c *RomaClient) CreateSpecialThrottlingConfigurationV2Invoker(request *model.CreateSpecialThrottlingConfigurationV2Request) *CreateSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForCreateSpecialThrottlingConfigurationV2()
	return &CreateSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugLiveDataApiV2 测试后端API
//
// 测试后端API是否可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DebugLiveDataApiV2(request *model.DebugLiveDataApiV2Request) (*model.DebugLiveDataApiV2Response, error) {
	requestDef := GenReqDefForDebugLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugLiveDataApiV2Response), nil
	}
}

// DebugLiveDataApiV2Invoker 测试后端API
func (c *RomaClient) DebugLiveDataApiV2Invoker(request *model.DebugLiveDataApiV2Request) *DebugLiveDataApiV2Invoker {
	requestDef := GenReqDefForDebugLiveDataApiV2()
	return &DebugLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugRule 规则调试
//
// 规则调试
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DebugRule(request *model.DebugRuleRequest) (*model.DebugRuleResponse, error) {
	requestDef := GenReqDefForDebugRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugRuleResponse), nil
	}
}

// DebugRuleInvoker 规则调试
func (c *RomaClient) DebugRuleInvoker(request *model.DebugRuleRequest) *DebugRuleInvoker {
	requestDef := GenReqDefForDebugRule()
	return &DebugRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppAcl 删除APP的访问控制
//
// 删除客户端配置的访问控制信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAppAcl(request *model.DeleteAppAclRequest) (*model.DeleteAppAclResponse, error) {
	requestDef := GenReqDefForDeleteAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppAclResponse), nil
	}
}

// DeleteAppAclInvoker 删除APP的访问控制
func (c *RomaClient) DeleteAppAclInvoker(request *model.DeleteAppAclRequest) *DeleteAppAclInvoker {
	requestDef := GenReqDefForDeleteAppAcl()
	return &DeleteAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppCodeV2 删除APP Code
//
// 删除App Code，App Code删除后，将无法再通过简易认证访问对应的API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAppCodeV2(request *model.DeleteAppCodeV2Request) (*model.DeleteAppCodeV2Response, error) {
	requestDef := GenReqDefForDeleteAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppCodeV2Response), nil
	}
}

// DeleteAppCodeV2Invoker 删除APP Code
func (c *RomaClient) DeleteAppCodeV2Invoker(request *model.DeleteAppCodeV2Request) *DeleteAppCodeV2Invoker {
	requestDef := GenReqDefForDeleteAppCodeV2()
	return &DeleteAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppConfigV2 删除应用配置
//
// 删除应用配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAppConfigV2(request *model.DeleteAppConfigV2Request) (*model.DeleteAppConfigV2Response, error) {
	requestDef := GenReqDefForDeleteAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppConfigV2Response), nil
	}
}

// DeleteAppConfigV2Invoker 删除应用配置
func (c *RomaClient) DeleteAppConfigV2Invoker(request *model.DeleteAppConfigV2Request) *DeleteAppConfigV2Invoker {
	requestDef := GenReqDefForDeleteAppConfigV2()
	return &DeleteAppConfigV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppQuota 删除客户端配额
//
// 删除客户端配额。删除客户端配额时，同时删除客户端配额和客户端应用的关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAppQuota(request *model.DeleteAppQuotaRequest) (*model.DeleteAppQuotaResponse, error) {
	requestDef := GenReqDefForDeleteAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppQuotaResponse), nil
	}
}

// DeleteAppQuotaInvoker 删除客户端配额
func (c *RomaClient) DeleteAppQuotaInvoker(request *model.DeleteAppQuotaRequest) *DeleteAppQuotaInvoker {
	requestDef := GenReqDefForDeleteAppQuota()
	return &DeleteAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCommand 删除命令
//
// 删除命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteCommand(request *model.DeleteCommandRequest) (*model.DeleteCommandResponse, error) {
	requestDef := GenReqDefForDeleteCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCommandResponse), nil
	}
}

// DeleteCommandInvoker 删除命令
func (c *RomaClient) DeleteCommandInvoker(request *model.DeleteCommandRequest) *DeleteCommandInvoker {
	requestDef := GenReqDefForDeleteCommand()
	return &DeleteCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomAuthorizerV2 删除自定义认证
//
// 删除自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteCustomAuthorizerV2(request *model.DeleteCustomAuthorizerV2Request) (*model.DeleteCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomAuthorizerV2Response), nil
	}
}

// DeleteCustomAuthorizerV2Invoker 删除自定义认证
func (c *RomaClient) DeleteCustomAuthorizerV2Invoker(request *model.DeleteCustomAuthorizerV2Request) *DeleteCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForDeleteCustomAuthorizerV2()
	return &DeleteCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatasourceInfoById 通过数据源Id删除指定数据源信息
//
// 通过数据源Id删除指定数据源信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDatasourceInfoById(request *model.DeleteDatasourceInfoByIdRequest) (*model.DeleteDatasourceInfoByIdResponse, error) {
	requestDef := GenReqDefForDeleteDatasourceInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatasourceInfoByIdResponse), nil
	}
}

// DeleteDatasourceInfoByIdInvoker 通过数据源Id删除指定数据源信息
func (c *RomaClient) DeleteDatasourceInfoByIdInvoker(request *model.DeleteDatasourceInfoByIdRequest) *DeleteDatasourceInfoByIdInvoker {
	requestDef := GenReqDefForDeleteDatasourceInfoById()
	return &DeleteDatasourceInfoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDestination 删除目标数据源
//
// 删除目标数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDestination(request *model.DeleteDestinationRequest) (*model.DeleteDestinationResponse, error) {
	requestDef := GenReqDefForDeleteDestination()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDestinationResponse), nil
	}
}

// DeleteDestinationInvoker 删除目标数据源
func (c *RomaClient) DeleteDestinationInvoker(request *model.DeleteDestinationRequest) *DeleteDestinationInvoker {
	requestDef := GenReqDefForDeleteDestination()
	return &DeleteDestinationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDevice 删除设备
//
// 删除指定设备ID的设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

// DeleteDeviceInvoker 删除设备
func (c *RomaClient) DeleteDeviceInvoker(request *model.DeleteDeviceRequest) *DeleteDeviceInvoker {
	requestDef := GenReqDefForDeleteDevice()
	return &DeleteDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeviceFromGroup 删除设备分组内的设备
//
// 删除设备分组内的设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDeviceFromGroup(request *model.DeleteDeviceFromGroupRequest) (*model.DeleteDeviceFromGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeviceFromGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceFromGroupResponse), nil
	}
}

// DeleteDeviceFromGroupInvoker 删除设备分组内的设备
func (c *RomaClient) DeleteDeviceFromGroupInvoker(request *model.DeleteDeviceFromGroupRequest) *DeleteDeviceFromGroupInvoker {
	requestDef := GenReqDefForDeleteDeviceFromGroup()
	return &DeleteDeviceFromGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeviceGroup 删除设备分组
//
// 删除分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDeviceGroup(request *model.DeleteDeviceGroupRequest) (*model.DeleteDeviceGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceGroupResponse), nil
	}
}

// DeleteDeviceGroupInvoker 删除设备分组
func (c *RomaClient) DeleteDeviceGroupInvoker(request *model.DeleteDeviceGroupRequest) *DeleteDeviceGroupInvoker {
	requestDef := GenReqDefForDeleteDeviceGroup()
	return &DeleteDeviceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironmentV2 删除环境
//
// 删除指定的环境。
// 该操作将导致此API在指定的环境无法被访问，可能会影响相当一部分应用和用户。请确保已经告知用户，或者确认需要强制下线。环境上存在已发布的API时，该环境不能被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteEnvironmentV2(request *model.DeleteEnvironmentV2Request) (*model.DeleteEnvironmentV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentV2Response), nil
	}
}

// DeleteEnvironmentV2Invoker 删除环境
func (c *RomaClient) DeleteEnvironmentV2Invoker(request *model.DeleteEnvironmentV2Request) *DeleteEnvironmentV2Invoker {
	requestDef := GenReqDefForDeleteEnvironmentV2()
	return &DeleteEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironmentVariableV2 删除变量
//
// 删除指定的环境变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteEnvironmentVariableV2(request *model.DeleteEnvironmentVariableV2Request) (*model.DeleteEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentVariableV2Response), nil
	}
}

// DeleteEnvironmentVariableV2Invoker 删除变量
func (c *RomaClient) DeleteEnvironmentVariableV2Invoker(request *model.DeleteEnvironmentVariableV2Request) *DeleteEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForDeleteEnvironmentVariableV2()
	return &DeleteEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLiveDataApiV2 删除后端API
//
// 在某个实例中删除后端API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteLiveDataApiV2(request *model.DeleteLiveDataApiV2Request) (*model.DeleteLiveDataApiV2Response, error) {
	requestDef := GenReqDefForDeleteLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLiveDataApiV2Response), nil
	}
}

// DeleteLiveDataApiV2Invoker 删除后端API
func (c *RomaClient) DeleteLiveDataApiV2Invoker(request *model.DeleteLiveDataApiV2Request) *DeleteLiveDataApiV2Invoker {
	requestDef := GenReqDefForDeleteLiveDataApiV2()
	return &DeleteLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMqsInstanceTopic 删除Topic
//
// 删除Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteMqsInstanceTopic(request *model.DeleteMqsInstanceTopicRequest) (*model.DeleteMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForDeleteMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMqsInstanceTopicResponse), nil
	}
}

// DeleteMqsInstanceTopicInvoker 删除Topic
func (c *RomaClient) DeleteMqsInstanceTopicInvoker(request *model.DeleteMqsInstanceTopicRequest) *DeleteMqsInstanceTopicInvoker {
	requestDef := GenReqDefForDeleteMqsInstanceTopic()
	return &DeleteMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMultiTaskMapping 删除指定任务映射
//
// 通过映射ID删除指定任务映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteMultiTaskMapping(request *model.DeleteMultiTaskMappingRequest) (*model.DeleteMultiTaskMappingResponse, error) {
	requestDef := GenReqDefForDeleteMultiTaskMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMultiTaskMappingResponse), nil
	}
}

// DeleteMultiTaskMappingInvoker 删除指定任务映射
func (c *RomaClient) DeleteMultiTaskMappingInvoker(request *model.DeleteMultiTaskMappingRequest) *DeleteMultiTaskMappingInvoker {
	requestDef := GenReqDefForDeleteMultiTaskMapping()
	return &DeleteMultiTaskMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotification 删除订阅管理
//
// 该接口用于删除指定订阅管理
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteNotification(request *model.DeleteNotificationRequest) (*model.DeleteNotificationResponse, error) {
	requestDef := GenReqDefForDeleteNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotificationResponse), nil
	}
}

// DeleteNotificationInvoker 删除订阅管理
func (c *RomaClient) DeleteNotificationInvoker(request *model.DeleteNotificationRequest) *DeleteNotificationInvoker {
	requestDef := GenReqDefForDeleteNotification()
	return &DeleteNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProduct 删除产品
//
// 删除产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteProduct(request *model.DeleteProductRequest) (*model.DeleteProductResponse, error) {
	requestDef := GenReqDefForDeleteProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductResponse), nil
	}
}

// DeleteProductInvoker 删除产品
func (c *RomaClient) DeleteProductInvoker(request *model.DeleteProductRequest) *DeleteProductInvoker {
	requestDef := GenReqDefForDeleteProduct()
	return &DeleteProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProductTemplate 删除产品模板
//
// 删除产品模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteProductTemplate(request *model.DeleteProductTemplateRequest) (*model.DeleteProductTemplateResponse, error) {
	requestDef := GenReqDefForDeleteProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductTemplateResponse), nil
	}
}

// DeleteProductTemplateInvoker 删除产品模板
func (c *RomaClient) DeleteProductTemplateInvoker(request *model.DeleteProductTemplateRequest) *DeleteProductTemplateInvoker {
	requestDef := GenReqDefForDeleteProductTemplate()
	return &DeleteProductTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProductTopic 删除产品主题
//
// 删除产品主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteProductTopic(request *model.DeleteProductTopicRequest) (*model.DeleteProductTopicResponse, error) {
	requestDef := GenReqDefForDeleteProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductTopicResponse), nil
	}
}

// DeleteProductTopicInvoker 删除产品主题
func (c *RomaClient) DeleteProductTopicInvoker(request *model.DeleteProductTopicRequest) *DeleteProductTopicInvoker {
	requestDef := GenReqDefForDeleteProductTopic()
	return &DeleteProductTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProperty 删除服务属性
//
// 删除服务属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteProperty(request *model.DeletePropertyRequest) (*model.DeletePropertyResponse, error) {
	requestDef := GenReqDefForDeleteProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePropertyResponse), nil
	}
}

// DeletePropertyInvoker 删除服务属性
func (c *RomaClient) DeletePropertyInvoker(request *model.DeletePropertyRequest) *DeletePropertyInvoker {
	requestDef := GenReqDefForDeleteProperty()
	return &DeletePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRequestProperty 删除请求属性
//
// 删除请求属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteRequestProperty(request *model.DeleteRequestPropertyRequest) (*model.DeleteRequestPropertyResponse, error) {
	requestDef := GenReqDefForDeleteRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestPropertyResponse), nil
	}
}

// DeleteRequestPropertyInvoker 删除请求属性
func (c *RomaClient) DeleteRequestPropertyInvoker(request *model.DeleteRequestPropertyRequest) *DeleteRequestPropertyInvoker {
	requestDef := GenReqDefForDeleteRequestProperty()
	return &DeleteRequestPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRequestThrottlingPolicyV2 删除流控策略
//
// 删除指定的流控策略。当该流控策略绑定了API时，需要先解除流控策略与API的所有绑定关系后再删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteRequestThrottlingPolicyV2(request *model.DeleteRequestThrottlingPolicyV2Request) (*model.DeleteRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRequestThrottlingPolicyV2Response), nil
	}
}

// DeleteRequestThrottlingPolicyV2Invoker 删除流控策略
func (c *RomaClient) DeleteRequestThrottlingPolicyV2Invoker(request *model.DeleteRequestThrottlingPolicyV2Request) *DeleteRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForDeleteRequestThrottlingPolicyV2()
	return &DeleteRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResponseProperty 删除响应属性
//
// 删除响应属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteResponseProperty(request *model.DeleteResponsePropertyRequest) (*model.DeleteResponsePropertyResponse, error) {
	requestDef := GenReqDefForDeleteResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResponsePropertyResponse), nil
	}
}

// DeleteResponsePropertyInvoker 删除响应属性
func (c *RomaClient) DeleteResponsePropertyInvoker(request *model.DeleteResponsePropertyRequest) *DeleteResponsePropertyInvoker {
	requestDef := GenReqDefForDeleteResponseProperty()
	return &DeleteResponsePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRule 删除规则
//
// 删除规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

// DeleteRuleInvoker 删除规则
func (c *RomaClient) DeleteRuleInvoker(request *model.DeleteRuleRequest) *DeleteRuleInvoker {
	requestDef := GenReqDefForDeleteRule()
	return &DeleteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteService 删除服务
//
// 删除服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

// DeleteServiceInvoker 删除服务
func (c *RomaClient) DeleteServiceInvoker(request *model.DeleteServiceRequest) *DeleteServiceInvoker {
	requestDef := GenReqDefForDeleteService()
	return &DeleteServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSignatureKeyV2 删除签名密钥
//
// 删除指定的签名密钥。签名密钥绑定了API时无法删除，需要先解除与API的绑定关系后删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteSignatureKeyV2(request *model.DeleteSignatureKeyV2Request) (*model.DeleteSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDeleteSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSignatureKeyV2Response), nil
	}
}

// DeleteSignatureKeyV2Invoker 删除签名密钥
func (c *RomaClient) DeleteSignatureKeyV2Invoker(request *model.DeleteSignatureKeyV2Request) *DeleteSignatureKeyV2Invoker {
	requestDef := GenReqDefForDeleteSignatureKeyV2()
	return &DeleteSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSource 删除源数据源
//
// 删除源数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteSource(request *model.DeleteSourceRequest) (*model.DeleteSourceResponse, error) {
	requestDef := GenReqDefForDeleteSource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSourceResponse), nil
	}
}

// DeleteSourceInvoker 删除源数据源
func (c *RomaClient) DeleteSourceInvoker(request *model.DeleteSourceRequest) *DeleteSourceInvoker {
	requestDef := GenReqDefForDeleteSource()
	return &DeleteSourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSpecialThrottlingConfigurationV2 删除特殊设置
//
// 删除某个流控策略的某个特殊配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteSpecialThrottlingConfigurationV2(request *model.DeleteSpecialThrottlingConfigurationV2Request) (*model.DeleteSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSpecialThrottlingConfigurationV2Response), nil
	}
}

// DeleteSpecialThrottlingConfigurationV2Invoker 删除特殊设置
func (c *RomaClient) DeleteSpecialThrottlingConfigurationV2Invoker(request *model.DeleteSpecialThrottlingConfigurationV2Request) *DeleteSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForDeleteSpecialThrottlingConfigurationV2()
	return &DeleteSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 通过任务ID删除指定任务
//
// 通过任务ID删除指定任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 通过任务ID删除指定任务
func (c *RomaClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateAppQuotaWithApp 解除客户端配额和客户端应用的绑定
//
// 解除客户端配额和客户端应用的绑定
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DisassociateAppQuotaWithApp(request *model.DisassociateAppQuotaWithAppRequest) (*model.DisassociateAppQuotaWithAppResponse, error) {
	requestDef := GenReqDefForDisassociateAppQuotaWithApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAppQuotaWithAppResponse), nil
	}
}

// DisassociateAppQuotaWithAppInvoker 解除客户端配额和客户端应用的绑定
func (c *RomaClient) DisassociateAppQuotaWithAppInvoker(request *model.DisassociateAppQuotaWithAppRequest) *DisassociateAppQuotaWithAppInvoker {
	requestDef := GenReqDefForDisassociateAppQuotaWithApp()
	return &DisassociateAppQuotaWithAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateCertificateV2 删除域名证书
//
// 如果域名证书不再需要或者已过期，则可以删除证书内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DisassociateCertificateV2(request *model.DisassociateCertificateV2Request) (*model.DisassociateCertificateV2Response, error) {
	requestDef := GenReqDefForDisassociateCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateCertificateV2Response), nil
	}
}

// DisassociateCertificateV2Invoker 删除域名证书
func (c *RomaClient) DisassociateCertificateV2Invoker(request *model.DisassociateCertificateV2Request) *DisassociateCertificateV2Invoker {
	requestDef := GenReqDefForDisassociateCertificateV2()
	return &DisassociateCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateDomainV2 解绑域名
//
// 如果API分组不再需要绑定某个自定义域名，则可以为此API分组解绑此域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DisassociateDomainV2(request *model.DisassociateDomainV2Request) (*model.DisassociateDomainV2Response, error) {
	requestDef := GenReqDefForDisassociateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateDomainV2Response), nil
	}
}

// DisassociateDomainV2Invoker 解绑域名
func (c *RomaClient) DisassociateDomainV2Invoker(request *model.DisassociateDomainV2Request) *DisassociateDomainV2Invoker {
	requestDef := GenReqDefForDisassociateDomainV2()
	return &DisassociateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateSignatureKeyV2 解除绑定
//
// 解除API与签名密钥的绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DisassociateSignatureKeyV2(request *model.DisassociateSignatureKeyV2Request) (*model.DisassociateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateSignatureKeyV2Response), nil
	}
}

// DisassociateSignatureKeyV2Invoker 解除绑定
func (c *RomaClient) DisassociateSignatureKeyV2Invoker(request *model.DisassociateSignatureKeyV2Request) *DisassociateSignatureKeyV2Invoker {
	requestDef := GenReqDefForDisassociateSignatureKeyV2()
	return &DisassociateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadProducts 导出产品
//
// 导出产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DownloadProducts(request *model.DownloadProductsRequest) (*model.DownloadProductsResponse, error) {
	requestDef := GenReqDefForDownloadProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadProductsResponse), nil
	}
}

// DownloadProductsInvoker 导出产品
func (c *RomaClient) DownloadProductsInvoker(request *model.DownloadProductsRequest) *DownloadProductsInvoker {
	requestDef := GenReqDefForDownloadProducts()
	return &DownloadProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportMqsInstanceTopic 导出Topic
//
// 导出Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ExportMqsInstanceTopic(request *model.ExportMqsInstanceTopicRequest) (*model.ExportMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForExportMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportMqsInstanceTopicResponse), nil
	}
}

// ExportMqsInstanceTopicInvoker 导出Topic
func (c *RomaClient) ExportMqsInstanceTopicInvoker(request *model.ExportMqsInstanceTopicRequest) *ExportMqsInstanceTopicInvoker {
	requestDef := GenReqDefForExportMqsInstanceTopic()
	return &ExportMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportMqsInstanceTopic 导入Topic
//
// 导入Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ImportMqsInstanceTopic(request *model.ImportMqsInstanceTopicRequest) (*model.ImportMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForImportMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportMqsInstanceTopicResponse), nil
	}
}

// ImportMqsInstanceTopicInvoker 导入Topic
func (c *RomaClient) ImportMqsInstanceTopicInvoker(request *model.ImportMqsInstanceTopicRequest) *ImportMqsInstanceTopicInvoker {
	requestDef := GenReqDefForImportMqsInstanceTopic()
	return &ImportMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallMultiTasks 组合任务初始化
//
// 初始化组合任务，分配任务ID，初始化映射等
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) InstallMultiTasks(request *model.InstallMultiTasksRequest) (*model.InstallMultiTasksResponse, error) {
	requestDef := GenReqDefForInstallMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallMultiTasksResponse), nil
	}
}

// InstallMultiTasksInvoker 组合任务初始化
func (c *RomaClient) InstallMultiTasksInvoker(request *model.InstallMultiTasksRequest) *InstallMultiTasksInvoker {
	requestDef := GenReqDefForInstallMultiTasks()
	return &InstallMultiTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToSignatureKeyV2 查看签名密钥绑定的API列表
//
// 查询某个签名密钥上已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToSignatureKeyV2(request *model.ListApisBindedToSignatureKeyV2Request) (*model.ListApisBindedToSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToSignatureKeyV2Response), nil
	}
}

// ListApisBindedToSignatureKeyV2Invoker 查看签名密钥绑定的API列表
func (c *RomaClient) ListApisBindedToSignatureKeyV2Invoker(request *model.ListApisBindedToSignatureKeyV2Request) *ListApisBindedToSignatureKeyV2Invoker {
	requestDef := GenReqDefForListApisBindedToSignatureKeyV2()
	return &ListApisBindedToSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisNotBoundWithSignatureKeyV2 查看签名密钥未绑定的API列表
//
// 查询所有未绑定到该签名密钥上的API列表。需要API已经发布，未发布的API不予展示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisNotBoundWithSignatureKeyV2(request *model.ListApisNotBoundWithSignatureKeyV2Request) (*model.ListApisNotBoundWithSignatureKeyV2Response, error) {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisNotBoundWithSignatureKeyV2Response), nil
	}
}

// ListApisNotBoundWithSignatureKeyV2Invoker 查看签名密钥未绑定的API列表
func (c *RomaClient) ListApisNotBoundWithSignatureKeyV2Invoker(request *model.ListApisNotBoundWithSignatureKeyV2Request) *ListApisNotBoundWithSignatureKeyV2Invoker {
	requestDef := GenReqDefForListApisNotBoundWithSignatureKeyV2()
	return &ListApisNotBoundWithSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppCodesV2 查询APP Code列表
//
// 查询App Code列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppCodesV2(request *model.ListAppCodesV2Request) (*model.ListAppCodesV2Response, error) {
	requestDef := GenReqDefForListAppCodesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppCodesV2Response), nil
	}
}

// ListAppCodesV2Invoker 查询APP Code列表
func (c *RomaClient) ListAppCodesV2Invoker(request *model.ListAppCodesV2Request) *ListAppCodesV2Invoker {
	requestDef := GenReqDefForListAppCodesV2()
	return &ListAppCodesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppConfigsV2 查询应用配置列表
//
// 查询应用配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppConfigsV2(request *model.ListAppConfigsV2Request) (*model.ListAppConfigsV2Response, error) {
	requestDef := GenReqDefForListAppConfigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppConfigsV2Response), nil
	}
}

// ListAppConfigsV2Invoker 查询应用配置列表
func (c *RomaClient) ListAppConfigsV2Invoker(request *model.ListAppConfigsV2Request) *ListAppConfigsV2Invoker {
	requestDef := GenReqDefForListAppConfigsV2()
	return &ListAppConfigsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotaBindableApps 查询客户端配额可绑定的客户端应用列表
//
// 查询客户端配额可绑定的客户端应用列表。支持按客户端应用名称模糊搜索
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppQuotaBindableApps(request *model.ListAppQuotaBindableAppsRequest) (*model.ListAppQuotaBindableAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBindableApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBindableAppsResponse), nil
	}
}

// ListAppQuotaBindableAppsInvoker 查询客户端配额可绑定的客户端应用列表
func (c *RomaClient) ListAppQuotaBindableAppsInvoker(request *model.ListAppQuotaBindableAppsRequest) *ListAppQuotaBindableAppsInvoker {
	requestDef := GenReqDefForListAppQuotaBindableApps()
	return &ListAppQuotaBindableAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotaBoundApps 查询客户端配额已绑定的客户端应用列表
//
// 查询客户端配额已绑定的客户端应用列表。支持按客户端应用名称模糊匹配
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppQuotaBoundApps(request *model.ListAppQuotaBoundAppsRequest) (*model.ListAppQuotaBoundAppsResponse, error) {
	requestDef := GenReqDefForListAppQuotaBoundApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotaBoundAppsResponse), nil
	}
}

// ListAppQuotaBoundAppsInvoker 查询客户端配额已绑定的客户端应用列表
func (c *RomaClient) ListAppQuotaBoundAppsInvoker(request *model.ListAppQuotaBoundAppsRequest) *ListAppQuotaBoundAppsInvoker {
	requestDef := GenReqDefForListAppQuotaBoundApps()
	return &ListAppQuotaBoundAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppQuotas 获取客户端配额列表
//
// 获取客户端配额列表。支持根据名称模糊查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppQuotas(request *model.ListAppQuotasRequest) (*model.ListAppQuotasResponse, error) {
	requestDef := GenReqDefForListAppQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppQuotasResponse), nil
	}
}

// ListAppQuotasInvoker 获取客户端配额列表
func (c *RomaClient) ListAppQuotasInvoker(request *model.ListAppQuotasRequest) *ListAppQuotasInvoker {
	requestDef := GenReqDefForListAppQuotas()
	return &ListAppQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppsV2 查询APP列表
//
// 查询APP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppsV2(request *model.ListAppsV2Request) (*model.ListAppsV2Response, error) {
	requestDef := GenReqDefForListAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsV2Response), nil
	}
}

// ListAppsV2Invoker 查询APP列表
func (c *RomaClient) ListAppsV2Invoker(request *model.ListAppsV2Request) *ListAppsV2Invoker {
	requestDef := GenReqDefForListAppsV2()
	return &ListAppsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommands 查询命令
//
// 查询命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListCommands(request *model.ListCommandsRequest) (*model.ListCommandsResponse, error) {
	requestDef := GenReqDefForListCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommandsResponse), nil
	}
}

// ListCommandsInvoker 查询命令
func (c *RomaClient) ListCommandsInvoker(request *model.ListCommandsRequest) *ListCommandsInvoker {
	requestDef := GenReqDefForListCommands()
	return &ListCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomAuthorizersV2 查询自定义认证列表
//
// 查询自定义认证列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListCustomAuthorizersV2(request *model.ListCustomAuthorizersV2Request) (*model.ListCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForListCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomAuthorizersV2Response), nil
	}
}

// ListCustomAuthorizersV2Invoker 查询自定义认证列表
func (c *RomaClient) ListCustomAuthorizersV2Invoker(request *model.ListCustomAuthorizersV2Request) *ListCustomAuthorizersV2Invoker {
	requestDef := GenReqDefForListCustomAuthorizersV2()
	return &ListCustomAuthorizersV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatasourceColumns 获取数据源中某个表中所有字段
//
// 获取数据源中中某个表中所有字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDatasourceColumns(request *model.ListDatasourceColumnsRequest) (*model.ListDatasourceColumnsResponse, error) {
	requestDef := GenReqDefForListDatasourceColumns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourceColumnsResponse), nil
	}
}

// ListDatasourceColumnsInvoker 获取数据源中某个表中所有字段
func (c *RomaClient) ListDatasourceColumnsInvoker(request *model.ListDatasourceColumnsRequest) *ListDatasourceColumnsInvoker {
	requestDef := GenReqDefForListDatasourceColumns()
	return &ListDatasourceColumnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatasourceTables 获取数据源中所有的表
//
// 获取数据源中所有的表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDatasourceTables(request *model.ListDatasourceTablesRequest) (*model.ListDatasourceTablesResponse, error) {
	requestDef := GenReqDefForListDatasourceTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourceTablesResponse), nil
	}
}

// ListDatasourceTablesInvoker 获取数据源中所有的表
func (c *RomaClient) ListDatasourceTablesInvoker(request *model.ListDatasourceTablesRequest) *ListDatasourceTablesInvoker {
	requestDef := GenReqDefForListDatasourceTables()
	return &ListDatasourceTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatasources 查询数据源
//
// 查询数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDatasources(request *model.ListDatasourcesRequest) (*model.ListDatasourcesResponse, error) {
	requestDef := GenReqDefForListDatasources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourcesResponse), nil
	}
}

// ListDatasourcesInvoker 查询数据源
func (c *RomaClient) ListDatasourcesInvoker(request *model.ListDatasourcesRequest) *ListDatasourcesInvoker {
	requestDef := GenReqDefForListDatasources()
	return &ListDatasourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDestinations 查询目标数据源列表
//
// 查询目标数据源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDestinations(request *model.ListDestinationsRequest) (*model.ListDestinationsResponse, error) {
	requestDef := GenReqDefForListDestinations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDestinationsResponse), nil
	}
}

// ListDestinationsInvoker 查询目标数据源列表
func (c *RomaClient) ListDestinationsInvoker(request *model.ListDestinationsRequest) *ListDestinationsInvoker {
	requestDef := GenReqDefForListDestinations()
	return &ListDestinationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevices 查询设备
//
// 查询设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

// ListDevicesInvoker 查询设备
func (c *RomaClient) ListDevicesInvoker(request *model.ListDevicesRequest) *ListDevicesInvoker {
	requestDef := GenReqDefForListDevices()
	return &ListDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevicesInProduct 查询产品内设备数量
//
// 查询产品内设备数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDevicesInProduct(request *model.ListDevicesInProductRequest) (*model.ListDevicesInProductResponse, error) {
	requestDef := GenReqDefForListDevicesInProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesInProductResponse), nil
	}
}

// ListDevicesInProductInvoker 查询产品内设备数量
func (c *RomaClient) ListDevicesInProductInvoker(request *model.ListDevicesInProductRequest) *ListDevicesInProductInvoker {
	requestDef := GenReqDefForListDevicesInProduct()
	return &ListDevicesInProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentVariablesV2 查询变量列表
//
// 查询分组下的所有环境变量的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListEnvironmentVariablesV2(request *model.ListEnvironmentVariablesV2Request) (*model.ListEnvironmentVariablesV2Response, error) {
	requestDef := GenReqDefForListEnvironmentVariablesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentVariablesV2Response), nil
	}
}

// ListEnvironmentVariablesV2Invoker 查询变量列表
func (c *RomaClient) ListEnvironmentVariablesV2Invoker(request *model.ListEnvironmentVariablesV2Request) *ListEnvironmentVariablesV2Invoker {
	requestDef := GenReqDefForListEnvironmentVariablesV2()
	return &ListEnvironmentVariablesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironmentsV2 查询环境列表
//
// 查询符合条件的环境列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListEnvironmentsV2(request *model.ListEnvironmentsV2Request) (*model.ListEnvironmentsV2Response, error) {
	requestDef := GenReqDefForListEnvironmentsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsV2Response), nil
	}
}

// ListEnvironmentsV2Invoker 查询环境列表
func (c *RomaClient) ListEnvironmentsV2Invoker(request *model.ListEnvironmentsV2Request) *ListEnvironmentsV2Invoker {
	requestDef := GenReqDefForListEnvironmentsV2()
	return &ListEnvironmentsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeaturesV2 查看实例特性列表
//
// 查看实例特性列表。注意：实例不支持以下特性的需要联系技术支持升级实例版本。
//
// 支持配置的特性列表及特性配置请参考“附录 &gt; 实例支持的APIC特性”
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListFeaturesV2(request *model.ListFeaturesV2Request) (*model.ListFeaturesV2Response, error) {
	requestDef := GenReqDefForListFeaturesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesV2Response), nil
	}
}

// ListFeaturesV2Invoker 查看实例特性列表
func (c *RomaClient) ListFeaturesV2Invoker(request *model.ListFeaturesV2Request) *ListFeaturesV2Invoker {
	requestDef := GenReqDefForListFeaturesV2()
	return &ListFeaturesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatelyApiStatisticsV2 API指标统计值查询-最近一段时间
//
// 根据API的id和最近的一段时间查询API被调用的次数，统计周期为1分钟。查询范围一小时以内，一分钟一个样本，其样本值为一分钟内的累计值。
// &gt; 为了安全起见，在服务器上使用curl命令调用接口查询信息后，需要清理历史操作记录，包括但不限于“~/.bash_history”、“/var/log/messages”（如有）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLatelyApiStatisticsV2(request *model.ListLatelyApiStatisticsV2Request) (*model.ListLatelyApiStatisticsV2Response, error) {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatelyApiStatisticsV2Response), nil
	}
}

// ListLatelyApiStatisticsV2Invoker API指标统计值查询-最近一段时间
func (c *RomaClient) ListLatelyApiStatisticsV2Invoker(request *model.ListLatelyApiStatisticsV2Request) *ListLatelyApiStatisticsV2Invoker {
	requestDef := GenReqDefForListLatelyApiStatisticsV2()
	return &ListLatelyApiStatisticsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLiveDataApiDeploymentHistoryV2 查询后端API部署历史
//
// 在某个实例中查询后端API的部署记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiDeploymentHistoryV2(request *model.ListLiveDataApiDeploymentHistoryV2Request) (*model.ListLiveDataApiDeploymentHistoryV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiDeploymentHistoryV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiDeploymentHistoryV2Response), nil
	}
}

// ListLiveDataApiDeploymentHistoryV2Invoker 查询后端API部署历史
func (c *RomaClient) ListLiveDataApiDeploymentHistoryV2Invoker(request *model.ListLiveDataApiDeploymentHistoryV2Request) *ListLiveDataApiDeploymentHistoryV2Invoker {
	requestDef := GenReqDefForListLiveDataApiDeploymentHistoryV2()
	return &ListLiveDataApiDeploymentHistoryV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLiveDataApiTestHistoryV2 查询后端API测试结果
//
// 在某个实例中查询后端API的测试结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiTestHistoryV2(request *model.ListLiveDataApiTestHistoryV2Request) (*model.ListLiveDataApiTestHistoryV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiTestHistoryV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiTestHistoryV2Response), nil
	}
}

// ListLiveDataApiTestHistoryV2Invoker 查询后端API测试结果
func (c *RomaClient) ListLiveDataApiTestHistoryV2Invoker(request *model.ListLiveDataApiTestHistoryV2Request) *ListLiveDataApiTestHistoryV2Invoker {
	requestDef := GenReqDefForListLiveDataApiTestHistoryV2()
	return &ListLiveDataApiTestHistoryV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLiveDataApiV2 查询后端API列表
//
// 获取某个实例下的所有后端API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLiveDataApiV2(request *model.ListLiveDataApiV2Request) (*model.ListLiveDataApiV2Response, error) {
	requestDef := GenReqDefForListLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataApiV2Response), nil
	}
}

// ListLiveDataApiV2Invoker 查询后端API列表
func (c *RomaClient) ListLiveDataApiV2Invoker(request *model.ListLiveDataApiV2Request) *ListLiveDataApiV2Invoker {
	requestDef := GenReqDefForListLiveDataApiV2()
	return &ListLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLiveDataDataSourcesV2 查询自定义后端服务数据源列表
//
// 查询自定义后端服务数据源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLiveDataDataSourcesV2(request *model.ListLiveDataDataSourcesV2Request) (*model.ListLiveDataDataSourcesV2Response, error) {
	requestDef := GenReqDefForListLiveDataDataSourcesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataDataSourcesV2Response), nil
	}
}

// ListLiveDataDataSourcesV2Invoker 查询自定义后端服务数据源列表
func (c *RomaClient) ListLiveDataDataSourcesV2Invoker(request *model.ListLiveDataDataSourcesV2Request) *ListLiveDataDataSourcesV2Invoker {
	requestDef := GenReqDefForListLiveDataDataSourcesV2()
	return &ListLiveDataDataSourcesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLiveDataQuotaV2 查询自定义后端服务配额
//
// 查询自定义后端服务配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListLiveDataQuotaV2(request *model.ListLiveDataQuotaV2Request) (*model.ListLiveDataQuotaV2Response, error) {
	requestDef := GenReqDefForListLiveDataQuotaV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveDataQuotaV2Response), nil
	}
}

// ListLiveDataQuotaV2Invoker 查询自定义后端服务配额
func (c *RomaClient) ListLiveDataQuotaV2Invoker(request *model.ListLiveDataQuotaV2Request) *ListLiveDataQuotaV2Invoker {
	requestDef := GenReqDefForListLiveDataQuotaV2()
	return &ListLiveDataQuotaV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitorInfos 任务监控信息列表查询
//
// 查询所有任务的监控信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListMonitorInfos(request *model.ListMonitorInfosRequest) (*model.ListMonitorInfosResponse, error) {
	requestDef := GenReqDefForListMonitorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorInfosResponse), nil
	}
}

// ListMonitorInfosInvoker 任务监控信息列表查询
func (c *RomaClient) ListMonitorInfosInvoker(request *model.ListMonitorInfosRequest) *ListMonitorInfosInvoker {
	requestDef := GenReqDefForListMonitorInfos()
	return &ListMonitorInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitorLog 任务监控日志查询
//
// 查询单个任务的所有日志信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListMonitorLog(request *model.ListMonitorLogRequest) (*model.ListMonitorLogResponse, error) {
	requestDef := GenReqDefForListMonitorLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitorLogResponse), nil
	}
}

// ListMonitorLogInvoker 任务监控日志查询
func (c *RomaClient) ListMonitorLogInvoker(request *model.ListMonitorLogRequest) *ListMonitorLogInvoker {
	requestDef := GenReqDefForListMonitorLog()
	return &ListMonitorLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMqsInstanceTopics 查询Topic列表
//
// 查询Topic列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListMqsInstanceTopics(request *model.ListMqsInstanceTopicsRequest) (*model.ListMqsInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListMqsInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMqsInstanceTopicsResponse), nil
	}
}

// ListMqsInstanceTopicsInvoker 查询Topic列表
func (c *RomaClient) ListMqsInstanceTopicsInvoker(request *model.ListMqsInstanceTopicsRequest) *ListMqsInstanceTopicsInvoker {
	requestDef := GenReqDefForListMqsInstanceTopics()
	return &ListMqsInstanceTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotification 查询订阅管理信息
//
// 该接口用于查询指定应用订阅管理信息的数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListNotification(request *model.ListNotificationRequest) (*model.ListNotificationResponse, error) {
	requestDef := GenReqDefForListNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationResponse), nil
	}
}

// ListNotificationInvoker 查询订阅管理信息
func (c *RomaClient) ListNotificationInvoker(request *model.ListNotificationRequest) *ListNotificationInvoker {
	requestDef := GenReqDefForListNotification()
	return &ListNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductTemplates 查询产品模板
//
// 查询产品模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProductTemplates(request *model.ListProductTemplatesRequest) (*model.ListProductTemplatesResponse, error) {
	requestDef := GenReqDefForListProductTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTemplatesResponse), nil
	}
}

// ListProductTemplatesInvoker 查询产品模板
func (c *RomaClient) ListProductTemplatesInvoker(request *model.ListProductTemplatesRequest) *ListProductTemplatesInvoker {
	requestDef := GenReqDefForListProductTemplates()
	return &ListProductTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductTopics 查询产品主题
//
// 查询产品主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProductTopics(request *model.ListProductTopicsRequest) (*model.ListProductTopicsResponse, error) {
	requestDef := GenReqDefForListProductTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductTopicsResponse), nil
	}
}

// ListProductTopicsInvoker 查询产品主题
func (c *RomaClient) ListProductTopicsInvoker(request *model.ListProductTopicsRequest) *ListProductTopicsInvoker {
	requestDef := GenReqDefForListProductTopics()
	return &ListProductTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品
//
// 查询产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询产品
func (c *RomaClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectCofigsV2 查询某个实例的租户配置列表
//
// 查询某个实例的租户配置列表，用户可以通过此接口查看各类型资源配置及使用情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProjectCofigsV2(request *model.ListProjectCofigsV2Request) (*model.ListProjectCofigsV2Response, error) {
	requestDef := GenReqDefForListProjectCofigsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectCofigsV2Response), nil
	}
}

// ListProjectCofigsV2Invoker 查询某个实例的租户配置列表
func (c *RomaClient) ListProjectCofigsV2Invoker(request *model.ListProjectCofigsV2Request) *ListProjectCofigsV2Invoker {
	requestDef := GenReqDefForListProjectCofigsV2()
	return &ListProjectCofigsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProperties 查询属性
//
// 查询属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProperties(request *model.ListPropertiesRequest) (*model.ListPropertiesResponse, error) {
	requestDef := GenReqDefForListProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropertiesResponse), nil
	}
}

// ListPropertiesInvoker 查询属性
func (c *RomaClient) ListPropertiesInvoker(request *model.ListPropertiesRequest) *ListPropertiesInvoker {
	requestDef := GenReqDefForListProperties()
	return &ListPropertiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRequestProperties 查询请求属性
//
// 查询请求属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListRequestProperties(request *model.ListRequestPropertiesRequest) (*model.ListRequestPropertiesResponse, error) {
	requestDef := GenReqDefForListRequestProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestPropertiesResponse), nil
	}
}

// ListRequestPropertiesInvoker 查询请求属性
func (c *RomaClient) ListRequestPropertiesInvoker(request *model.ListRequestPropertiesRequest) *ListRequestPropertiesInvoker {
	requestDef := GenReqDefForListRequestProperties()
	return &ListRequestPropertiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRequestThrottlingPolicyV2 查询流控策略列表
//
// 查询所有流控策略的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListRequestThrottlingPolicyV2(request *model.ListRequestThrottlingPolicyV2Request) (*model.ListRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPolicyV2Response), nil
	}
}

// ListRequestThrottlingPolicyV2Invoker 查询流控策略列表
func (c *RomaClient) ListRequestThrottlingPolicyV2Invoker(request *model.ListRequestThrottlingPolicyV2Request) *ListRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListRequestThrottlingPolicyV2()
	return &ListRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResponseProperties 查询响应属性
//
// 查询响应属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListResponseProperties(request *model.ListResponsePropertiesRequest) (*model.ListResponsePropertiesResponse, error) {
	requestDef := GenReqDefForListResponseProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResponsePropertiesResponse), nil
	}
}

// ListResponsePropertiesInvoker 查询响应属性
func (c *RomaClient) ListResponsePropertiesInvoker(request *model.ListResponsePropertiesRequest) *ListResponsePropertiesInvoker {
	requestDef := GenReqDefForListResponseProperties()
	return &ListResponsePropertiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRules 查询规则
//
// 查询规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

// ListRulesInvoker 查询规则
func (c *RomaClient) ListRulesInvoker(request *model.ListRulesRequest) *ListRulesInvoker {
	requestDef := GenReqDefForListRules()
	return &ListRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServices 查询服务
//
// 查询服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListServices(request *model.ListServicesRequest) (*model.ListServicesResponse, error) {
	requestDef := GenReqDefForListServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesResponse), nil
	}
}

// ListServicesInvoker 查询服务
func (c *RomaClient) ListServicesInvoker(request *model.ListServicesRequest) *ListServicesInvoker {
	requestDef := GenReqDefForListServices()
	return &ListServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShadows 查询设备影子
//
// 查询设备影子
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListShadows(request *model.ListShadowsRequest) (*model.ListShadowsResponse, error) {
	requestDef := GenReqDefForListShadows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShadowsResponse), nil
	}
}

// ListShadowsInvoker 查询设备影子
func (c *RomaClient) ListShadowsInvoker(request *model.ListShadowsRequest) *ListShadowsInvoker {
	requestDef := GenReqDefForListShadows()
	return &ListShadowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSignatureKeysBindedToApiV2 查看API绑定的签名密钥列表
//
// 查询某个API绑定的签名密钥列表。每个API在每个环境上应该最多只会绑定一个签名密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListSignatureKeysBindedToApiV2(request *model.ListSignatureKeysBindedToApiV2Request) (*model.ListSignatureKeysBindedToApiV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysBindedToApiV2Response), nil
	}
}

// ListSignatureKeysBindedToApiV2Invoker 查看API绑定的签名密钥列表
func (c *RomaClient) ListSignatureKeysBindedToApiV2Invoker(request *model.ListSignatureKeysBindedToApiV2Request) *ListSignatureKeysBindedToApiV2Invoker {
	requestDef := GenReqDefForListSignatureKeysBindedToApiV2()
	return &ListSignatureKeysBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSignatureKeysV2 查询签名密钥列表
//
// 查询所有签名密钥的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListSignatureKeysV2(request *model.ListSignatureKeysV2Request) (*model.ListSignatureKeysV2Response, error) {
	requestDef := GenReqDefForListSignatureKeysV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureKeysV2Response), nil
	}
}

// ListSignatureKeysV2Invoker 查询签名密钥列表
func (c *RomaClient) ListSignatureKeysV2Invoker(request *model.ListSignatureKeysV2Request) *ListSignatureKeysV2Invoker {
	requestDef := GenReqDefForListSignatureKeysV2()
	return &ListSignatureKeysV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSources 查询源数据源列表
//
// 查询源数据源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListSources(request *model.ListSourcesRequest) (*model.ListSourcesResponse, error) {
	requestDef := GenReqDefForListSources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSourcesResponse), nil
	}
}

// ListSourcesInvoker 查询源数据源列表
func (c *RomaClient) ListSourcesInvoker(request *model.ListSourcesRequest) *ListSourcesInvoker {
	requestDef := GenReqDefForListSources()
	return &ListSourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecialThrottlingConfigurationsV2 查看特殊设置列表
//
// 查看给流控策略设置的特殊配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListSpecialThrottlingConfigurationsV2(request *model.ListSpecialThrottlingConfigurationsV2Request) (*model.ListSpecialThrottlingConfigurationsV2Response, error) {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecialThrottlingConfigurationsV2Response), nil
	}
}

// ListSpecialThrottlingConfigurationsV2Invoker 查看特殊设置列表
func (c *RomaClient) ListSpecialThrottlingConfigurationsV2Invoker(request *model.ListSpecialThrottlingConfigurationsV2Request) *ListSpecialThrottlingConfigurationsV2Invoker {
	requestDef := GenReqDefForListSpecialThrottlingConfigurationsV2()
	return &ListSpecialThrottlingConfigurationsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatisticsApi 查询API指标统计值
//
// 查询某个实例下的API统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListStatisticsApi(request *model.ListStatisticsApiRequest) (*model.ListStatisticsApiResponse, error) {
	requestDef := GenReqDefForListStatisticsApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsApiResponse), nil
	}
}

// ListStatisticsApiInvoker 查询API指标统计值
func (c *RomaClient) ListStatisticsApiInvoker(request *model.ListStatisticsApiRequest) *ListStatisticsApiInvoker {
	requestDef := GenReqDefForListStatisticsApi()
	return &ListStatisticsApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubsets 查询子设备
//
// 查询子设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListSubsets(request *model.ListSubsetsRequest) (*model.ListSubsetsResponse, error) {
	requestDef := GenReqDefForListSubsets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubsetsResponse), nil
	}
}

// ListSubsetsInvoker 查询子设备
func (c *RomaClient) ListSubsetsInvoker(request *model.ListSubsetsRequest) *ListSubsetsInvoker {
	requestDef := GenReqDefForListSubsets()
	return &ListSubsetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsV2 查询标签列表
//
// 查询标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListTagsV2(request *model.ListTagsV2Request) (*model.ListTagsV2Response, error) {
	requestDef := GenReqDefForListTagsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsV2Response), nil
	}
}

// ListTagsV2Invoker 查询标签列表
func (c *RomaClient) ListTagsV2Invoker(request *model.ListTagsV2Request) *ListTagsV2Invoker {
	requestDef := GenReqDefForListTagsV2()
	return &ListTagsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询任务列表
//
// 查询任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表
func (c *RomaClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopics 查询设备主题
//
// 查询设备主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListTopics(request *model.ListTopicsRequest) (*model.ListTopicsResponse, error) {
	requestDef := GenReqDefForListTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicsResponse), nil
	}
}

// ListTopicsInvoker 查询设备主题
func (c *RomaClient) ListTopicsInvoker(request *model.ListTopicsRequest) *ListTopicsInvoker {
	requestDef := GenReqDefForListTopics()
	return &ListTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishLiveDataApiV2 部署后端API
//
// 在某个实例中部署后端API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) PublishLiveDataApiV2(request *model.PublishLiveDataApiV2Request) (*model.PublishLiveDataApiV2Response, error) {
	requestDef := GenReqDefForPublishLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishLiveDataApiV2Response), nil
	}
}

// PublishLiveDataApiV2Invoker 部署后端API
func (c *RomaClient) PublishLiveDataApiV2Invoker(request *model.PublishLiveDataApiV2Request) *PublishLiveDataApiV2Invoker {
	requestDef := GenReqDefForPublishLiveDataApiV2()
	return &PublishLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetAuthentication 重置设备鉴权信息
//
// 重置设备鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ResetAuthentication(request *model.ResetAuthenticationRequest) (*model.ResetAuthenticationResponse, error) {
	requestDef := GenReqDefForResetAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetAuthenticationResponse), nil
	}
}

// ResetAuthenticationInvoker 重置设备鉴权信息
func (c *RomaClient) ResetAuthenticationInvoker(request *model.ResetAuthenticationRequest) *ResetAuthenticationInvoker {
	requestDef := GenReqDefForResetAuthentication()
	return &ResetAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetMessages 重发消息
//
// 重发消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ResetMessages(request *model.ResetMessagesRequest) (*model.ResetMessagesResponse, error) {
	requestDef := GenReqDefForResetMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMessagesResponse), nil
	}
}

// ResetMessagesInvoker 重发消息
func (c *RomaClient) ResetMessagesInvoker(request *model.ResetMessagesRequest) *ResetMessagesInvoker {
	requestDef := GenReqDefForResetMessages()
	return &ResetMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetMultiTaskOffset 重置组合任务进度
//
// 重置组合任务进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ResetMultiTaskOffset(request *model.ResetMultiTaskOffsetRequest) (*model.ResetMultiTaskOffsetResponse, error) {
	requestDef := GenReqDefForResetMultiTaskOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMultiTaskOffsetResponse), nil
	}
}

// ResetMultiTaskOffsetInvoker 重置组合任务进度
func (c *RomaClient) ResetMultiTaskOffsetInvoker(request *model.ResetMultiTaskOffsetRequest) *ResetMultiTaskOffsetInvoker {
	requestDef := GenReqDefForResetMultiTaskOffset()
	return &ResetMultiTaskOffsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetProductAuthentication 重置产品鉴权信息
//
// 重置产品鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ResetProductAuthentication(request *model.ResetProductAuthenticationRequest) (*model.ResetProductAuthenticationResponse, error) {
	requestDef := GenReqDefForResetProductAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetProductAuthenticationResponse), nil
	}
}

// ResetProductAuthenticationInvoker 重置产品鉴权信息
func (c *RomaClient) ResetProductAuthenticationInvoker(request *model.ResetProductAuthenticationRequest) *ResetProductAuthenticationInvoker {
	requestDef := GenReqDefForResetProductAuthentication()
	return &ResetProductAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTask 手工触发单个任务
//
// 手工触发一次任务调度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) RunTask(request *model.RunTaskRequest) (*model.RunTaskResponse, error) {
	requestDef := GenReqDefForRunTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTaskResponse), nil
	}
}

// RunTaskInvoker 手工触发单个任务
func (c *RomaClient) RunTaskInvoker(request *model.RunTaskRequest) *RunTaskInvoker {
	requestDef := GenReqDefForRunTask()
	return &RunTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendCommand 发送命令
//
// 发送命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) SendCommand(request *model.SendCommandRequest) (*model.SendCommandResponse, error) {
	requestDef := GenReqDefForSendCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendCommandResponse), nil
	}
}

// SendCommandInvoker 发送命令
func (c *RomaClient) SendCommandInvoker(request *model.SendCommandRequest) *SendCommandInvoker {
	requestDef := GenReqDefForSendCommand()
	return &SendCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppBoundAppQuota 查询客户端应用关联的应用配额
//
// 查看指定客户端应用关联的应用配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowAppBoundAppQuota(request *model.ShowAppBoundAppQuotaRequest) (*model.ShowAppBoundAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppBoundAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppBoundAppQuotaResponse), nil
	}
}

// ShowAppBoundAppQuotaInvoker 查询客户端应用关联的应用配额
func (c *RomaClient) ShowAppBoundAppQuotaInvoker(request *model.ShowAppBoundAppQuotaRequest) *ShowAppBoundAppQuotaInvoker {
	requestDef := GenReqDefForShowAppBoundAppQuota()
	return &ShowAppBoundAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppQuota 获取客户端配额详情
//
// 获取客户端配额详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowAppQuota(request *model.ShowAppQuotaRequest) (*model.ShowAppQuotaResponse, error) {
	requestDef := GenReqDefForShowAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppQuotaResponse), nil
	}
}

// ShowAppQuotaInvoker 获取客户端配额详情
func (c *RomaClient) ShowAppQuotaInvoker(request *model.ShowAppQuotaRequest) *ShowAppQuotaInvoker {
	requestDef := GenReqDefForShowAppQuota()
	return &ShowAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthentication 查询设备鉴权信息
//
// 查询设备鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowAuthentication(request *model.ShowAuthenticationRequest) (*model.ShowAuthenticationResponse, error) {
	requestDef := GenReqDefForShowAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthenticationResponse), nil
	}
}

// ShowAuthenticationInvoker 查询设备鉴权信息
func (c *RomaClient) ShowAuthenticationInvoker(request *model.ShowAuthenticationRequest) *ShowAuthenticationInvoker {
	requestDef := GenReqDefForShowAuthentication()
	return &ShowAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCommand 查询命令详情
//
// 查询命令详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowCommand(request *model.ShowCommandRequest) (*model.ShowCommandResponse, error) {
	requestDef := GenReqDefForShowCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommandResponse), nil
	}
}

// ShowCommandInvoker 查询命令详情
func (c *RomaClient) ShowCommandInvoker(request *model.ShowCommandRequest) *ShowCommandInvoker {
	requestDef := GenReqDefForShowCommand()
	return &ShowCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataourceDetail 查询指定数据源
//
// 根据数据源id查询数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDataourceDetail(request *model.ShowDataourceDetailRequest) (*model.ShowDataourceDetailResponse, error) {
	requestDef := GenReqDefForShowDataourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataourceDetailResponse), nil
	}
}

// ShowDataourceDetailInvoker 查询指定数据源
func (c *RomaClient) ShowDataourceDetailInvoker(request *model.ShowDataourceDetailRequest) *ShowDataourceDetailInvoker {
	requestDef := GenReqDefForShowDataourceDetail()
	return &ShowDataourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppAcl 查看APP的访问控制详情
//
// 查看APP的访问控制详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppAcl(request *model.ShowDetailsOfAppAclRequest) (*model.ShowDetailsOfAppAclResponse, error) {
	requestDef := GenReqDefForShowDetailsOfAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppAclResponse), nil
	}
}

// ShowDetailsOfAppAclInvoker 查看APP的访问控制详情
func (c *RomaClient) ShowDetailsOfAppAclInvoker(request *model.ShowDetailsOfAppAclRequest) *ShowDetailsOfAppAclInvoker {
	requestDef := GenReqDefForShowDetailsOfAppAcl()
	return &ShowDetailsOfAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppCodeV2 查看APP Code详情
//
// App Code为APP应用下的子模块，创建App Code之后，可以实现简易的APP认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppCodeV2(request *model.ShowDetailsOfAppCodeV2Request) (*model.ShowDetailsOfAppCodeV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppCodeV2Response), nil
	}
}

// ShowDetailsOfAppCodeV2Invoker 查看APP Code详情
func (c *RomaClient) ShowDetailsOfAppCodeV2Invoker(request *model.ShowDetailsOfAppCodeV2Request) *ShowDetailsOfAppCodeV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAppCodeV2()
	return &ShowDetailsOfAppCodeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppConfigV2 查看应用配置详情
//
// 查看应用配置详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppConfigV2(request *model.ShowDetailsOfAppConfigV2Request) (*model.ShowDetailsOfAppConfigV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppConfigV2Response), nil
	}
}

// ShowDetailsOfAppConfigV2Invoker 查看应用配置详情
func (c *RomaClient) ShowDetailsOfAppConfigV2Invoker(request *model.ShowDetailsOfAppConfigV2Request) *ShowDetailsOfAppConfigV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAppConfigV2()
	return &ShowDetailsOfAppConfigV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAppV2 查看APP详情
//
// 查看指定APP的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAppV2(request *model.ShowDetailsOfAppV2Request) (*model.ShowDetailsOfAppV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAppV2Response), nil
	}
}

// ShowDetailsOfAppV2Invoker 查看APP详情
func (c *RomaClient) ShowDetailsOfAppV2Invoker(request *model.ShowDetailsOfAppV2Request) *ShowDetailsOfAppV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAppV2()
	return &ShowDetailsOfAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfCustomAuthorizersV2 查看自定义认证详情
//
// 查看自定义认证详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfCustomAuthorizersV2(request *model.ShowDetailsOfCustomAuthorizersV2Request) (*model.ShowDetailsOfCustomAuthorizersV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfCustomAuthorizersV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfCustomAuthorizersV2Response), nil
	}
}

// ShowDetailsOfCustomAuthorizersV2Invoker 查看自定义认证详情
func (c *RomaClient) ShowDetailsOfCustomAuthorizersV2Invoker(request *model.ShowDetailsOfCustomAuthorizersV2Request) *ShowDetailsOfCustomAuthorizersV2Invoker {
	requestDef := GenReqDefForShowDetailsOfCustomAuthorizersV2()
	return &ShowDetailsOfCustomAuthorizersV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfDomainNameCertificateV2 查看域名证书
//
// 查看域名下绑定的证书详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfDomainNameCertificateV2(request *model.ShowDetailsOfDomainNameCertificateV2Request) (*model.ShowDetailsOfDomainNameCertificateV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfDomainNameCertificateV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfDomainNameCertificateV2Response), nil
	}
}

// ShowDetailsOfDomainNameCertificateV2Invoker 查看域名证书
func (c *RomaClient) ShowDetailsOfDomainNameCertificateV2Invoker(request *model.ShowDetailsOfDomainNameCertificateV2Request) *ShowDetailsOfDomainNameCertificateV2Invoker {
	requestDef := GenReqDefForShowDetailsOfDomainNameCertificateV2()
	return &ShowDetailsOfDomainNameCertificateV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfEnvironmentVariableV2 查看变量详情
//
// 查看指定的环境变量的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfEnvironmentVariableV2(request *model.ShowDetailsOfEnvironmentVariableV2Request) (*model.ShowDetailsOfEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfEnvironmentVariableV2Response), nil
	}
}

// ShowDetailsOfEnvironmentVariableV2Invoker 查看变量详情
func (c *RomaClient) ShowDetailsOfEnvironmentVariableV2Invoker(request *model.ShowDetailsOfEnvironmentVariableV2Request) *ShowDetailsOfEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForShowDetailsOfEnvironmentVariableV2()
	return &ShowDetailsOfEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfInstanceV2 查看ROMA Connect实例详情
//
// 查看ROMA Connect实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfInstanceV2(request *model.ShowDetailsOfInstanceV2Request) (*model.ShowDetailsOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfInstanceV2Response), nil
	}
}

// ShowDetailsOfInstanceV2Invoker 查看ROMA Connect实例详情
func (c *RomaClient) ShowDetailsOfInstanceV2Invoker(request *model.ShowDetailsOfInstanceV2Request) *ShowDetailsOfInstanceV2Invoker {
	requestDef := GenReqDefForShowDetailsOfInstanceV2()
	return &ShowDetailsOfInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfRequestThrottlingPolicyV2 查看流控策略详情
//
// 查看指定流控策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfRequestThrottlingPolicyV2(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfRequestThrottlingPolicyV2Response), nil
	}
}

// ShowDetailsOfRequestThrottlingPolicyV2Invoker 查看流控策略详情
func (c *RomaClient) ShowDetailsOfRequestThrottlingPolicyV2Invoker(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) *ShowDetailsOfRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForShowDetailsOfRequestThrottlingPolicyV2()
	return &ShowDetailsOfRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDevice 查询设备详情
//
// 查询设备详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDevice(request *model.ShowDeviceRequest) (*model.ShowDeviceResponse, error) {
	requestDef := GenReqDefForShowDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceResponse), nil
	}
}

// ShowDeviceInvoker 查询设备详情
func (c *RomaClient) ShowDeviceInvoker(request *model.ShowDeviceRequest) *ShowDeviceInvoker {
	requestDef := GenReqDefForShowDevice()
	return &ShowDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceGroup 查询设备分组详情
//
// 获取设备分组及下一层分组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDeviceGroup(request *model.ShowDeviceGroupRequest) (*model.ShowDeviceGroupResponse, error) {
	requestDef := GenReqDefForShowDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceGroupResponse), nil
	}
}

// ShowDeviceGroupInvoker 查询设备分组详情
func (c *RomaClient) ShowDeviceGroupInvoker(request *model.ShowDeviceGroupRequest) *ShowDeviceGroupInvoker {
	requestDef := GenReqDefForShowDeviceGroup()
	return &ShowDeviceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceGroupTree 查询所有设备分组
//
// 查询所有设备分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDeviceGroupTree(request *model.ShowDeviceGroupTreeRequest) (*model.ShowDeviceGroupTreeResponse, error) {
	requestDef := GenReqDefForShowDeviceGroupTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceGroupTreeResponse), nil
	}
}

// ShowDeviceGroupTreeInvoker 查询所有设备分组
func (c *RomaClient) ShowDeviceGroupTreeInvoker(request *model.ShowDeviceGroupTreeRequest) *ShowDeviceGroupTreeInvoker {
	requestDef := GenReqDefForShowDeviceGroupTree()
	return &ShowDeviceGroupTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDevicesInGroup 查询设备分组内设备
//
// 查询设备分组内设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDevicesInGroup(request *model.ShowDevicesInGroupRequest) (*model.ShowDevicesInGroupResponse, error) {
	requestDef := GenReqDefForShowDevicesInGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDevicesInGroupResponse), nil
	}
}

// ShowDevicesInGroupInvoker 查询设备分组内设备
func (c *RomaClient) ShowDevicesInGroupInvoker(request *model.ShowDevicesInGroupRequest) *ShowDevicesInGroupInvoker {
	requestDef := GenReqDefForShowDevicesInGroup()
	return &ShowDevicesInGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDispatches 查询调度计划
//
// 查询调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDispatches(request *model.ShowDispatchesRequest) (*model.ShowDispatchesResponse, error) {
	requestDef := GenReqDefForShowDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDispatchesResponse), nil
	}
}

// ShowDispatchesInvoker 查询调度计划
func (c *RomaClient) ShowDispatchesInvoker(request *model.ShowDispatchesRequest) *ShowDispatchesInvoker {
	requestDef := GenReqDefForShowDispatches()
	return &ShowDispatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLiveDataApiV2 查询后端API详情
//
// 查询后端API的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowLiveDataApiV2(request *model.ShowLiveDataApiV2Request) (*model.ShowLiveDataApiV2Response, error) {
	requestDef := GenReqDefForShowLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLiveDataApiV2Response), nil
	}
}

// ShowLiveDataApiV2Invoker 查询后端API详情
func (c *RomaClient) ShowLiveDataApiV2Invoker(request *model.ShowLiveDataApiV2Request) *ShowLiveDataApiV2Invoker {
	requestDef := GenReqDefForShowLiveDataApiV2()
	return &ShowLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMqsInstanceMessages 查询消息
//
// 查询消息的偏移量和消息内容。
// 先根据时间戳查询消息的偏移量，再根据偏移量查询消息内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstanceMessages(request *model.ShowMqsInstanceMessagesRequest) (*model.ShowMqsInstanceMessagesResponse, error) {
	requestDef := GenReqDefForShowMqsInstanceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceMessagesResponse), nil
	}
}

// ShowMqsInstanceMessagesInvoker 查询消息
func (c *RomaClient) ShowMqsInstanceMessagesInvoker(request *model.ShowMqsInstanceMessagesRequest) *ShowMqsInstanceMessagesInvoker {
	requestDef := GenReqDefForShowMqsInstanceMessages()
	return &ShowMqsInstanceMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMqsInstanceTopicAccessPolicy 查询Topic权限
//
// 查询Topic权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstanceTopicAccessPolicy(request *model.ShowMqsInstanceTopicAccessPolicyRequest) (*model.ShowMqsInstanceTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowMqsInstanceTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceTopicAccessPolicyResponse), nil
	}
}

// ShowMqsInstanceTopicAccessPolicyInvoker 查询Topic权限
func (c *RomaClient) ShowMqsInstanceTopicAccessPolicyInvoker(request *model.ShowMqsInstanceTopicAccessPolicyRequest) *ShowMqsInstanceTopicAccessPolicyInvoker {
	requestDef := GenReqDefForShowMqsInstanceTopicAccessPolicy()
	return &ShowMqsInstanceTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProduct 查询产品详情
//
// 查询产品详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowProduct(request *model.ShowProductRequest) (*model.ShowProductResponse, error) {
	requestDef := GenReqDefForShowProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductResponse), nil
	}
}

// ShowProductInvoker 查询产品详情
func (c *RomaClient) ShowProductInvoker(request *model.ShowProductRequest) *ShowProductInvoker {
	requestDef := GenReqDefForShowProduct()
	return &ShowProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProductAuthentication 查询产品鉴权信息
//
// 查询产品鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowProductAuthentication(request *model.ShowProductAuthenticationRequest) (*model.ShowProductAuthenticationResponse, error) {
	requestDef := GenReqDefForShowProductAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductAuthenticationResponse), nil
	}
}

// ShowProductAuthenticationInvoker 查询产品鉴权信息
func (c *RomaClient) ShowProductAuthenticationInvoker(request *model.ShowProductAuthenticationRequest) *ShowProductAuthenticationInvoker {
	requestDef := GenReqDefForShowProductAuthentication()
	return &ShowProductAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProductTemplate 查询产品模板详情
//
// 查询产品模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowProductTemplate(request *model.ShowProductTemplateRequest) (*model.ShowProductTemplateResponse, error) {
	requestDef := GenReqDefForShowProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductTemplateResponse), nil
	}
}

// ShowProductTemplateInvoker 查询产品模板详情
func (c *RomaClient) ShowProductTemplateInvoker(request *model.ShowProductTemplateRequest) *ShowProductTemplateInvoker {
	requestDef := GenReqDefForShowProductTemplate()
	return &ShowProductTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProperty 查询服务属性详情
//
// 查询服务属性详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowProperty(request *model.ShowPropertyRequest) (*model.ShowPropertyResponse, error) {
	requestDef := GenReqDefForShowProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPropertyResponse), nil
	}
}

// ShowPropertyInvoker 查询服务属性详情
func (c *RomaClient) ShowPropertyInvoker(request *model.ShowPropertyRequest) *ShowPropertyInvoker {
	requestDef := GenReqDefForShowProperty()
	return &ShowPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRequestProperty 查询请求属性详情
//
// 查询请求属性详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowRequestProperty(request *model.ShowRequestPropertyRequest) (*model.ShowRequestPropertyResponse, error) {
	requestDef := GenReqDefForShowRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRequestPropertyResponse), nil
	}
}

// ShowRequestPropertyInvoker 查询请求属性详情
func (c *RomaClient) ShowRequestPropertyInvoker(request *model.ShowRequestPropertyRequest) *ShowRequestPropertyInvoker {
	requestDef := GenReqDefForShowRequestProperty()
	return &ShowRequestPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResponseProperty 查询响应属性详情
//
// 查询响应属性详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowResponseProperty(request *model.ShowResponsePropertyRequest) (*model.ShowResponsePropertyResponse, error) {
	requestDef := GenReqDefForShowResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResponsePropertyResponse), nil
	}
}

// ShowResponsePropertyInvoker 查询响应属性详情
func (c *RomaClient) ShowResponsePropertyInvoker(request *model.ShowResponsePropertyRequest) *ShowResponsePropertyInvoker {
	requestDef := GenReqDefForShowResponseProperty()
	return &ShowResponsePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRestrictionOfInstanceV2 查看ROMA Connect实例约束信息
//
// 查看ROMA Connect实例约束信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowRestrictionOfInstanceV2(request *model.ShowRestrictionOfInstanceV2Request) (*model.ShowRestrictionOfInstanceV2Response, error) {
	requestDef := GenReqDefForShowRestrictionOfInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRestrictionOfInstanceV2Response), nil
	}
}

// ShowRestrictionOfInstanceV2Invoker 查看ROMA Connect实例约束信息
func (c *RomaClient) ShowRestrictionOfInstanceV2Invoker(request *model.ShowRestrictionOfInstanceV2Request) *ShowRestrictionOfInstanceV2Invoker {
	requestDef := GenReqDefForShowRestrictionOfInstanceV2()
	return &ShowRestrictionOfInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRule 查询规则详情
//
// 查询规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowRule(request *model.ShowRuleRequest) (*model.ShowRuleResponse, error) {
	requestDef := GenReqDefForShowRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleResponse), nil
	}
}

// ShowRuleInvoker 查询规则详情
func (c *RomaClient) ShowRuleInvoker(request *model.ShowRuleRequest) *ShowRuleInvoker {
	requestDef := GenReqDefForShowRule()
	return &ShowRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowService 查询服务详情
//
// 查询服务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowService(request *model.ShowServiceRequest) (*model.ShowServiceResponse, error) {
	requestDef := GenReqDefForShowService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceResponse), nil
	}
}

// ShowServiceInvoker 查询服务详情
func (c *RomaClient) ShowServiceInvoker(request *model.ShowServiceRequest) *ShowServiceInvoker {
	requestDef := GenReqDefForShowService()
	return &ShowServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 通过任务ID查询指定任务的信息
//
// 通过任务ID查询指定任务的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 通过任务ID查询指定任务的信息
func (c *RomaClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartTestDatasource 测试数据源连通性
//
// 测试数据源连通性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) StartTestDatasource(request *model.StartTestDatasourceRequest) (*model.StartTestDatasourceResponse, error) {
	requestDef := GenReqDefForStartTestDatasource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartTestDatasourceResponse), nil
	}
}

// StartTestDatasourceInvoker 测试数据源连通性
func (c *RomaClient) StartTestDatasourceInvoker(request *model.StartTestDatasourceRequest) *StartTestDatasourceInvoker {
	requestDef := GenReqDefForStartTestDatasource()
	return &StartTestDatasourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopTask 手工停止当前执行的任务
//
// 手工停止当前执行的任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) StopTask(request *model.StopTaskRequest) (*model.StopTaskResponse, error) {
	requestDef := GenReqDefForStopTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTaskResponse), nil
	}
}

// StopTaskInvoker 手工停止当前执行的任务
func (c *RomaClient) StopTaskInvoker(request *model.StopTaskRequest) *StopTaskInvoker {
	requestDef := GenReqDefForStopTask()
	return &StopTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnpublishLiveDataApiV2 撤销后端API
//
// 在某个实例中取消部署后端API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UnpublishLiveDataApiV2(request *model.UnpublishLiveDataApiV2Request) (*model.UnpublishLiveDataApiV2Response, error) {
	requestDef := GenReqDefForUnpublishLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnpublishLiveDataApiV2Response), nil
	}
}

// UnpublishLiveDataApiV2Invoker 撤销后端API
func (c *RomaClient) UnpublishLiveDataApiV2Invoker(request *model.UnpublishLiveDataApiV2Request) *UnpublishLiveDataApiV2Invoker {
	requestDef := GenReqDefForUnpublishLiveDataApiV2()
	return &UnpublishLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppAcl 设置APP的访问控制
//
// 设置客户端配置的访问控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateAppAcl(request *model.UpdateAppAclRequest) (*model.UpdateAppAclResponse, error) {
	requestDef := GenReqDefForUpdateAppAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppAclResponse), nil
	}
}

// UpdateAppAclInvoker 设置APP的访问控制
func (c *RomaClient) UpdateAppAclInvoker(request *model.UpdateAppAclRequest) *UpdateAppAclInvoker {
	requestDef := GenReqDefForUpdateAppAcl()
	return &UpdateAppAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppConfigV2 修改应用配置
//
// 修改应用配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateAppConfigV2(request *model.UpdateAppConfigV2Request) (*model.UpdateAppConfigV2Response, error) {
	requestDef := GenReqDefForUpdateAppConfigV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppConfigV2Response), nil
	}
}

// UpdateAppConfigV2Invoker 修改应用配置
func (c *RomaClient) UpdateAppConfigV2Invoker(request *model.UpdateAppConfigV2Request) *UpdateAppConfigV2Invoker {
	requestDef := GenReqDefForUpdateAppConfigV2()
	return &UpdateAppConfigV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppQuota 修改客户端配额
//
// 修改客户端配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateAppQuota(request *model.UpdateAppQuotaRequest) (*model.UpdateAppQuotaResponse, error) {
	requestDef := GenReqDefForUpdateAppQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppQuotaResponse), nil
	}
}

// UpdateAppQuotaInvoker 修改客户端配额
func (c *RomaClient) UpdateAppQuotaInvoker(request *model.UpdateAppQuotaRequest) *UpdateAppQuotaInvoker {
	requestDef := GenReqDefForUpdateAppQuota()
	return &UpdateAppQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCommand 修改命令
//
// 修改命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateCommand(request *model.UpdateCommandRequest) (*model.UpdateCommandResponse, error) {
	requestDef := GenReqDefForUpdateCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCommandResponse), nil
	}
}

// UpdateCommandInvoker 修改命令
func (c *RomaClient) UpdateCommandInvoker(request *model.UpdateCommandRequest) *UpdateCommandInvoker {
	requestDef := GenReqDefForUpdateCommand()
	return &UpdateCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomAuthorizerV2 修改自定义认证
//
// 修改自定义认证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateCustomAuthorizerV2(request *model.UpdateCustomAuthorizerV2Request) (*model.UpdateCustomAuthorizerV2Response, error) {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomAuthorizerV2Response), nil
	}
}

// UpdateCustomAuthorizerV2Invoker 修改自定义认证
func (c *RomaClient) UpdateCustomAuthorizerV2Invoker(request *model.UpdateCustomAuthorizerV2Request) *UpdateCustomAuthorizerV2Invoker {
	requestDef := GenReqDefForUpdateCustomAuthorizerV2()
	return &UpdateCustomAuthorizerV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatasourceInfo 修改数据源
//
// 修改数据源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDatasourceInfo(request *model.UpdateDatasourceInfoRequest) (*model.UpdateDatasourceInfoResponse, error) {
	requestDef := GenReqDefForUpdateDatasourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatasourceInfoResponse), nil
	}
}

// UpdateDatasourceInfoInvoker 修改数据源
func (c *RomaClient) UpdateDatasourceInfoInvoker(request *model.UpdateDatasourceInfoRequest) *UpdateDatasourceInfoInvoker {
	requestDef := GenReqDefForUpdateDatasourceInfo()
	return &UpdateDatasourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevice 修改设备
//
// 修改设备信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// UpdateDeviceInvoker 修改设备
func (c *RomaClient) UpdateDeviceInvoker(request *model.UpdateDeviceRequest) *UpdateDeviceInvoker {
	requestDef := GenReqDefForUpdateDevice()
	return &UpdateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeviceGroup 修改设备分组
//
// 修改设备分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDeviceGroup(request *model.UpdateDeviceGroupRequest) (*model.UpdateDeviceGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceGroupResponse), nil
	}
}

// UpdateDeviceGroupInvoker 修改设备分组
func (c *RomaClient) UpdateDeviceGroupInvoker(request *model.UpdateDeviceGroupRequest) *UpdateDeviceGroupInvoker {
	requestDef := GenReqDefForUpdateDeviceGroup()
	return &UpdateDeviceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDispatches 修改调度计划
//
// 通过任务ID和调度ID修改调度计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDispatches(request *model.UpdateDispatchesRequest) (*model.UpdateDispatchesResponse, error) {
	requestDef := GenReqDefForUpdateDispatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDispatchesResponse), nil
	}
}

// UpdateDispatchesInvoker 修改调度计划
func (c *RomaClient) UpdateDispatchesInvoker(request *model.UpdateDispatchesRequest) *UpdateDispatchesInvoker {
	requestDef := GenReqDefForUpdateDispatches()
	return &UpdateDispatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainV2 修改域名
//
// 修改绑定的域名所对应的配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDomainV2(request *model.UpdateDomainV2Request) (*model.UpdateDomainV2Response, error) {
	requestDef := GenReqDefForUpdateDomainV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainV2Response), nil
	}
}

// UpdateDomainV2Invoker 修改域名
func (c *RomaClient) UpdateDomainV2Invoker(request *model.UpdateDomainV2Request) *UpdateDomainV2Invoker {
	requestDef := GenReqDefForUpdateDomainV2()
	return &UpdateDomainV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironmentV2 修改环境
//
// 修改指定环境的信息。其中可修改的属性为：name、remark，其它属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateEnvironmentV2(request *model.UpdateEnvironmentV2Request) (*model.UpdateEnvironmentV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentV2Response), nil
	}
}

// UpdateEnvironmentV2Invoker 修改环境
func (c *RomaClient) UpdateEnvironmentV2Invoker(request *model.UpdateEnvironmentV2Request) *UpdateEnvironmentV2Invoker {
	requestDef := GenReqDefForUpdateEnvironmentV2()
	return &UpdateEnvironmentV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnvironmentVariableV2 修改变量
//
// 修改环境变量。环境变量引用位置为api的后端服务地址时，修改对应环境变量会将使用该变量的所有api重新发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateEnvironmentVariableV2(request *model.UpdateEnvironmentVariableV2Request) (*model.UpdateEnvironmentVariableV2Response, error) {
	requestDef := GenReqDefForUpdateEnvironmentVariableV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvironmentVariableV2Response), nil
	}
}

// UpdateEnvironmentVariableV2Invoker 修改变量
func (c *RomaClient) UpdateEnvironmentVariableV2Invoker(request *model.UpdateEnvironmentVariableV2Request) *UpdateEnvironmentVariableV2Invoker {
	requestDef := GenReqDefForUpdateEnvironmentVariableV2()
	return &UpdateEnvironmentVariableV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLiveDataApiV2 修改后端API
//
// 在某个实例中更新后端API的参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateLiveDataApiV2(request *model.UpdateLiveDataApiV2Request) (*model.UpdateLiveDataApiV2Response, error) {
	requestDef := GenReqDefForUpdateLiveDataApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLiveDataApiV2Response), nil
	}
}

// UpdateLiveDataApiV2Invoker 修改后端API
func (c *RomaClient) UpdateLiveDataApiV2Invoker(request *model.UpdateLiveDataApiV2Request) *UpdateLiveDataApiV2Invoker {
	requestDef := GenReqDefForUpdateLiveDataApiV2()
	return &UpdateLiveDataApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMqsInstanceTopic 修改Topic
//
// 修改Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateMqsInstanceTopic(request *model.UpdateMqsInstanceTopicRequest) (*model.UpdateMqsInstanceTopicResponse, error) {
	requestDef := GenReqDefForUpdateMqsInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMqsInstanceTopicResponse), nil
	}
}

// UpdateMqsInstanceTopicInvoker 修改Topic
func (c *RomaClient) UpdateMqsInstanceTopicInvoker(request *model.UpdateMqsInstanceTopicRequest) *UpdateMqsInstanceTopicInvoker {
	requestDef := GenReqDefForUpdateMqsInstanceTopic()
	return &UpdateMqsInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMultiTasks 修改组合任务
//
// 修改组合任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateMultiTasks(request *model.UpdateMultiTasksRequest) (*model.UpdateMultiTasksResponse, error) {
	requestDef := GenReqDefForUpdateMultiTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMultiTasksResponse), nil
	}
}

// UpdateMultiTasksInvoker 修改组合任务
func (c *RomaClient) UpdateMultiTasksInvoker(request *model.UpdateMultiTasksRequest) *UpdateMultiTasksInvoker {
	requestDef := GenReqDefForUpdateMultiTasks()
	return &UpdateMultiTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotification 修改订阅管理
//
// 该接口用于修改指定的订阅管理
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateNotification(request *model.UpdateNotificationRequest) (*model.UpdateNotificationResponse, error) {
	requestDef := GenReqDefForUpdateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationResponse), nil
	}
}

// UpdateNotificationInvoker 修改订阅管理
func (c *RomaClient) UpdateNotificationInvoker(request *model.UpdateNotificationRequest) *UpdateNotificationInvoker {
	requestDef := GenReqDefForUpdateNotification()
	return &UpdateNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProduct 修改产品信息
//
// 修改产品信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateProduct(request *model.UpdateProductRequest) (*model.UpdateProductResponse, error) {
	requestDef := GenReqDefForUpdateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductResponse), nil
	}
}

// UpdateProductInvoker 修改产品信息
func (c *RomaClient) UpdateProductInvoker(request *model.UpdateProductRequest) *UpdateProductInvoker {
	requestDef := GenReqDefForUpdateProduct()
	return &UpdateProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProductTemplate 修改产品模板
//
// 修改产品模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateProductTemplate(request *model.UpdateProductTemplateRequest) (*model.UpdateProductTemplateResponse, error) {
	requestDef := GenReqDefForUpdateProductTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductTemplateResponse), nil
	}
}

// UpdateProductTemplateInvoker 修改产品模板
func (c *RomaClient) UpdateProductTemplateInvoker(request *model.UpdateProductTemplateRequest) *UpdateProductTemplateInvoker {
	requestDef := GenReqDefForUpdateProductTemplate()
	return &UpdateProductTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProductTopic 更新产品主题
//
// 更新产品主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateProductTopic(request *model.UpdateProductTopicRequest) (*model.UpdateProductTopicResponse, error) {
	requestDef := GenReqDefForUpdateProductTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductTopicResponse), nil
	}
}

// UpdateProductTopicInvoker 更新产品主题
func (c *RomaClient) UpdateProductTopicInvoker(request *model.UpdateProductTopicRequest) *UpdateProductTopicInvoker {
	requestDef := GenReqDefForUpdateProductTopic()
	return &UpdateProductTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProperty 修改服务属性
//
// 修改服务属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateProperty(request *model.UpdatePropertyRequest) (*model.UpdatePropertyResponse, error) {
	requestDef := GenReqDefForUpdateProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePropertyResponse), nil
	}
}

// UpdatePropertyInvoker 修改服务属性
func (c *RomaClient) UpdatePropertyInvoker(request *model.UpdatePropertyRequest) *UpdatePropertyInvoker {
	requestDef := GenReqDefForUpdateProperty()
	return &UpdatePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRequestProperty 修改请求属性
//
// 修改请求属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateRequestProperty(request *model.UpdateRequestPropertyRequest) (*model.UpdateRequestPropertyResponse, error) {
	requestDef := GenReqDefForUpdateRequestProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestPropertyResponse), nil
	}
}

// UpdateRequestPropertyInvoker 修改请求属性
func (c *RomaClient) UpdateRequestPropertyInvoker(request *model.UpdateRequestPropertyRequest) *UpdateRequestPropertyInvoker {
	requestDef := GenReqDefForUpdateRequestProperty()
	return &UpdateRequestPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRequestThrottlingPolicyV2 修改流控策略
//
// 修改指定流控策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateRequestThrottlingPolicyV2(request *model.UpdateRequestThrottlingPolicyV2Request) (*model.UpdateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRequestThrottlingPolicyV2Response), nil
	}
}

// UpdateRequestThrottlingPolicyV2Invoker 修改流控策略
func (c *RomaClient) UpdateRequestThrottlingPolicyV2Invoker(request *model.UpdateRequestThrottlingPolicyV2Request) *UpdateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForUpdateRequestThrottlingPolicyV2()
	return &UpdateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResponseProperty 修改响应属性
//
// 修改响应属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateResponseProperty(request *model.UpdateResponsePropertyRequest) (*model.UpdateResponsePropertyResponse, error) {
	requestDef := GenReqDefForUpdateResponseProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResponsePropertyResponse), nil
	}
}

// UpdateResponsePropertyInvoker 修改响应属性
func (c *RomaClient) UpdateResponsePropertyInvoker(request *model.UpdateResponsePropertyRequest) *UpdateResponsePropertyInvoker {
	requestDef := GenReqDefForUpdateResponseProperty()
	return &UpdateResponsePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRule 修改规则
//
// 修改规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
	requestDef := GenReqDefForUpdateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleResponse), nil
	}
}

// UpdateRuleInvoker 修改规则
func (c *RomaClient) UpdateRuleInvoker(request *model.UpdateRuleRequest) *UpdateRuleInvoker {
	requestDef := GenReqDefForUpdateRule()
	return &UpdateRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateService 修改服务
//
// 修改服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

// UpdateServiceInvoker 修改服务
func (c *RomaClient) UpdateServiceInvoker(request *model.UpdateServiceRequest) *UpdateServiceInvoker {
	requestDef := GenReqDefForUpdateService()
	return &UpdateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSignatureKeyV2 修改签名密钥
//
// 修改指定签名密钥的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateSignatureKeyV2(request *model.UpdateSignatureKeyV2Request) (*model.UpdateSignatureKeyV2Response, error) {
	requestDef := GenReqDefForUpdateSignatureKeyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSignatureKeyV2Response), nil
	}
}

// UpdateSignatureKeyV2Invoker 修改签名密钥
func (c *RomaClient) UpdateSignatureKeyV2Invoker(request *model.UpdateSignatureKeyV2Request) *UpdateSignatureKeyV2Invoker {
	requestDef := GenReqDefForUpdateSignatureKeyV2()
	return &UpdateSignatureKeyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSpecialThrottlingConfigurationV2 修改特殊设置
//
// 修改某个流控策略下的某个特殊设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateSpecialThrottlingConfigurationV2(request *model.UpdateSpecialThrottlingConfigurationV2Request) (*model.UpdateSpecialThrottlingConfigurationV2Response, error) {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpecialThrottlingConfigurationV2Response), nil
	}
}

// UpdateSpecialThrottlingConfigurationV2Invoker 修改特殊设置
func (c *RomaClient) UpdateSpecialThrottlingConfigurationV2Invoker(request *model.UpdateSpecialThrottlingConfigurationV2Request) *UpdateSpecialThrottlingConfigurationV2Invoker {
	requestDef := GenReqDefForUpdateSpecialThrottlingConfigurationV2()
	return &UpdateSpecialThrottlingConfigurationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTask 更新普通任务
//
// 更新普通任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

// UpdateTaskInvoker 更新普通任务
func (c *RomaClient) UpdateTaskInvoker(request *model.UpdateTaskRequest) *UpdateTaskInvoker {
	requestDef := GenReqDefForUpdateTask()
	return &UpdateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopicAccessPolicy 更新Topic权限
//
// 更新Topic权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateTopicAccessPolicy(request *model.UpdateTopicAccessPolicyRequest) (*model.UpdateTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

// UpdateTopicAccessPolicyInvoker 更新Topic权限
func (c *RomaClient) UpdateTopicAccessPolicyInvoker(request *model.UpdateTopicAccessPolicyRequest) *UpdateTopicAccessPolicyInvoker {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()
	return &UpdateTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadProduct 导入产品
//
// 导入产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UploadProduct(request *model.UploadProductRequest) (*model.UploadProductResponse, error) {
	requestDef := GenReqDefForUploadProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadProductResponse), nil
	}
}

// UploadProductInvoker 导入产品
func (c *RomaClient) UploadProductInvoker(request *model.UploadProductRequest) *UploadProductInvoker {
	requestDef := GenReqDefForUploadProduct()
	return &UploadProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAclV2 批量删除ACL策略
//
// 批量删除指定的多个ACL策略。
//
// 删除ACL策略时，如果存在ACL策略与API绑定关系，则无法删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDeleteAclV2(request *model.BatchDeleteAclV2Request) (*model.BatchDeleteAclV2Response, error) {
	requestDef := GenReqDefForBatchDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclV2Response), nil
	}
}

// BatchDeleteAclV2Invoker 批量删除ACL策略
func (c *RomaClient) BatchDeleteAclV2Invoker(request *model.BatchDeleteAclV2Request) *BatchDeleteAclV2Invoker {
	requestDef := GenReqDefForBatchDeleteAclV2()
	return &BatchDeleteAclV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAclStrategyV2 创建ACL策略
//
// 增加一个ACL策略，策略类型通过字段acl_type来确定（permit或者deny），限制的对象的类型可以为IP[或者DOMAIN，这里的DOMAIN对应的acl_value的值为租户名称，而非“www.exampleDomain.com\&quot;之类的网络域名。](tag:hws;hws_hk;hcs;fcs;g42;)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAclStrategyV2(request *model.CreateAclStrategyV2Request) (*model.CreateAclStrategyV2Response, error) {
	requestDef := GenReqDefForCreateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAclStrategyV2Response), nil
	}
}

// CreateAclStrategyV2Invoker 创建ACL策略
func (c *RomaClient) CreateAclStrategyV2Invoker(request *model.CreateAclStrategyV2Request) *CreateAclStrategyV2Invoker {
	requestDef := GenReqDefForCreateAclStrategyV2()
	return &CreateAclStrategyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclV2 删除ACL策略
//
// 删除指定的ACL策略， 如果存在api与该ACL策略的绑定关系，则无法删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAclV2(request *model.DeleteAclV2Request) (*model.DeleteAclV2Response, error) {
	requestDef := GenReqDefForDeleteAclV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclV2Response), nil
	}
}

// DeleteAclV2Invoker 删除ACL策略
func (c *RomaClient) DeleteAclV2Invoker(request *model.DeleteAclV2Request) *DeleteAclV2Invoker {
	requestDef := GenReqDefForDeleteAclV2()
	return &DeleteAclV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclStrategiesV2 查看ACL策略列表
//
// 查询所有的ACL策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAclStrategiesV2(request *model.ListAclStrategiesV2Request) (*model.ListAclStrategiesV2Response, error) {
	requestDef := GenReqDefForListAclStrategiesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclStrategiesV2Response), nil
	}
}

// ListAclStrategiesV2Invoker 查看ACL策略列表
func (c *RomaClient) ListAclStrategiesV2Invoker(request *model.ListAclStrategiesV2Request) *ListAclStrategiesV2Invoker {
	requestDef := GenReqDefForListAclStrategiesV2()
	return &ListAclStrategiesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfAclPolicyV2 查看ACL策略详情
//
// 查询指定ACL策略的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfAclPolicyV2(request *model.ShowDetailsOfAclPolicyV2Request) (*model.ShowDetailsOfAclPolicyV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfAclPolicyV2Response), nil
	}
}

// ShowDetailsOfAclPolicyV2Invoker 查看ACL策略详情
func (c *RomaClient) ShowDetailsOfAclPolicyV2Invoker(request *model.ShowDetailsOfAclPolicyV2Request) *ShowDetailsOfAclPolicyV2Invoker {
	requestDef := GenReqDefForShowDetailsOfAclPolicyV2()
	return &ShowDetailsOfAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclStrategyV2 修改ACL策略
//
// 修改指定的ACL策略，其中可修改的属性为：acl_name、acl_type、acl_value，其它属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateAclStrategyV2(request *model.UpdateAclStrategyV2Request) (*model.UpdateAclStrategyV2Response, error) {
	requestDef := GenReqDefForUpdateAclStrategyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclStrategyV2Response), nil
	}
}

// UpdateAclStrategyV2Invoker 修改ACL策略
func (c *RomaClient) UpdateAclStrategyV2Invoker(request *model.UpdateAclStrategyV2Request) *UpdateAclStrategyV2Invoker {
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
func (c *RomaClient) AssociateRequestThrottlingPolicyV2(request *model.AssociateRequestThrottlingPolicyV2Request) (*model.AssociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForAssociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRequestThrottlingPolicyV2Response), nil
	}
}

// AssociateRequestThrottlingPolicyV2Invoker 绑定流控策略
func (c *RomaClient) AssociateRequestThrottlingPolicyV2Invoker(request *model.AssociateRequestThrottlingPolicyV2Request) *AssociateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForAssociateRequestThrottlingPolicyV2()
	return &AssociateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateThrottlingPolicyV2 批量解绑流控策略
//
// 批量解除API与流控策略的绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDisassociateThrottlingPolicyV2(request *model.BatchDisassociateThrottlingPolicyV2Request) (*model.BatchDisassociateThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForBatchDisassociateThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateThrottlingPolicyV2Response), nil
	}
}

// BatchDisassociateThrottlingPolicyV2Invoker 批量解绑流控策略
func (c *RomaClient) BatchDisassociateThrottlingPolicyV2Invoker(request *model.BatchDisassociateThrottlingPolicyV2Request) *BatchDisassociateThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForBatchDisassociateThrottlingPolicyV2()
	return &BatchDisassociateThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchPublishOrOfflineApiV2 批量发布或下线API
//
// 将多个API发布到一个指定的环境，或将多个API从指定的环境下线。
//
// 注意：当action &#x3D; online时，接口返回的响应中publish_id，version_id， publish_time字段才有含义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchPublishOrOfflineApiV2(request *model.BatchPublishOrOfflineApiV2Request) (*model.BatchPublishOrOfflineApiV2Response, error) {
	requestDef := GenReqDefForBatchPublishOrOfflineApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPublishOrOfflineApiV2Response), nil
	}
}

// BatchPublishOrOfflineApiV2Invoker 批量发布或下线API
func (c *RomaClient) BatchPublishOrOfflineApiV2Invoker(request *model.BatchPublishOrOfflineApiV2Request) *BatchPublishOrOfflineApiV2Invoker {
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
func (c *RomaClient) ChangeApiVersionV2(request *model.ChangeApiVersionV2Request) (*model.ChangeApiVersionV2Response, error) {
	requestDef := GenReqDefForChangeApiVersionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApiVersionV2Response), nil
	}
}

// ChangeApiVersionV2Invoker 切换API版本
func (c *RomaClient) ChangeApiVersionV2Invoker(request *model.ChangeApiVersionV2Request) *ChangeApiVersionV2Invoker {
	requestDef := GenReqDefForChangeApiVersionV2()
	return &ChangeApiVersionV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckApiGroupsV2 校验API分组名称是否存在
//
// 校验API分组名称是否存在。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckApiGroupsV2(request *model.CheckApiGroupsV2Request) (*model.CheckApiGroupsV2Response, error) {
	requestDef := GenReqDefForCheckApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApiGroupsV2Response), nil
	}
}

// CheckApiGroupsV2Invoker 校验API分组名称是否存在
func (c *RomaClient) CheckApiGroupsV2Invoker(request *model.CheckApiGroupsV2Request) *CheckApiGroupsV2Invoker {
	requestDef := GenReqDefForCheckApiGroupsV2()
	return &CheckApiGroupsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckApisV2 校验API定义
//
// 校验API定义。校验API的路径或名称是否已存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckApisV2(request *model.CheckApisV2Request) (*model.CheckApisV2Response, error) {
	requestDef := GenReqDefForCheckApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckApisV2Response), nil
	}
}

// CheckApisV2Invoker 校验API定义
func (c *RomaClient) CheckApisV2Invoker(request *model.CheckApisV2Request) *CheckApisV2Invoker {
	requestDef := GenReqDefForCheckApisV2()
	return &CheckApisV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiGroupV2 创建API分组
//
// API分组是API的管理单元，一个API分组等同于一个服务入口，创建API分组时，返回一个子域名作为访问入口。建议一个API分组下的API具有一定的相关性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateApiGroupV2(request *model.CreateApiGroupV2Request) (*model.CreateApiGroupV2Response, error) {
	requestDef := GenReqDefForCreateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiGroupV2Response), nil
	}
}

// CreateApiGroupV2Invoker 创建API分组
func (c *RomaClient) CreateApiGroupV2Invoker(request *model.CreateApiGroupV2Request) *CreateApiGroupV2Invoker {
	requestDef := GenReqDefForCreateApiGroupV2()
	return &CreateApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiV2 创建API
//
// 添加一个API，API即一个服务接口，具体的服务能力。
//
// API分为两部分，第一部分为面向API使用者的API接口，定义了使用者如何调用这个API。第二部分面向API提供者，由API提供者定义这个API的真实的后端情况，定义了ROMA Connect如何去访问真实的后端服务。API的真实后端服务目前支持三种类型：传统的HTTP/HTTPS形式的web后端、[函数工作流、](tag:hws;hws_hk;hcs;fcs;g42;)MOCK。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateApiV2(request *model.CreateApiV2Request) (*model.CreateApiV2Response, error) {
	requestDef := GenReqDefForCreateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiV2Response), nil
	}
}

// CreateApiV2Invoker 创建API
func (c *RomaClient) CreateApiV2Invoker(request *model.CreateApiV2Request) *CreateApiV2Invoker {
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
func (c *RomaClient) CreateOrDeletePublishRecordForApiV2(request *model.CreateOrDeletePublishRecordForApiV2Request) (*model.CreateOrDeletePublishRecordForApiV2Response, error) {
	requestDef := GenReqDefForCreateOrDeletePublishRecordForApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrDeletePublishRecordForApiV2Response), nil
	}
}

// CreateOrDeletePublishRecordForApiV2Invoker 发布或下线API
func (c *RomaClient) CreateOrDeletePublishRecordForApiV2Invoker(request *model.CreateOrDeletePublishRecordForApiV2Request) *CreateOrDeletePublishRecordForApiV2Invoker {
	requestDef := GenReqDefForCreateOrDeletePublishRecordForApiV2()
	return &CreateOrDeletePublishRecordForApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugApiV2 调试API
//
// 调试一个API在指定运行环境下的定义，接口调用者需要具有操作该API的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DebugApiV2(request *model.DebugApiV2Request) (*model.DebugApiV2Response, error) {
	requestDef := GenReqDefForDebugApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugApiV2Response), nil
	}
}

// DebugApiV2Invoker 调试API
func (c *RomaClient) DebugApiV2Invoker(request *model.DebugApiV2Request) *DebugApiV2Invoker {
	requestDef := GenReqDefForDebugApiV2()
	return &DebugApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiByVersionIdV2 根据版本编号下线API
//
// 对某个生效中的API版本进行下线操作，下线后，API在该版本生效的环境中将不再能够被调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteApiByVersionIdV2(request *model.DeleteApiByVersionIdV2Request) (*model.DeleteApiByVersionIdV2Response, error) {
	requestDef := GenReqDefForDeleteApiByVersionIdV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiByVersionIdV2Response), nil
	}
}

// DeleteApiByVersionIdV2Invoker 根据版本编号下线API
func (c *RomaClient) DeleteApiByVersionIdV2Invoker(request *model.DeleteApiByVersionIdV2Request) *DeleteApiByVersionIdV2Invoker {
	requestDef := GenReqDefForDeleteApiByVersionIdV2()
	return &DeleteApiByVersionIdV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiGroupV2 删除API分组
//
// 删除指定的API分组。
// 分组下存在API时分组无法删除，需要删除所有分组下的API后，再删除分组。
// 删除分组时，会一并删除直接或间接关联到该分组下的所有资源，包括独立域名、SSL证书等等。并会将外部域名与子域名的绑定关系进行解除（取决于域名cname方式）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteApiGroupV2(request *model.DeleteApiGroupV2Request) (*model.DeleteApiGroupV2Response, error) {
	requestDef := GenReqDefForDeleteApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiGroupV2Response), nil
	}
}

// DeleteApiGroupV2Invoker 删除API分组
func (c *RomaClient) DeleteApiGroupV2Invoker(request *model.DeleteApiGroupV2Request) *DeleteApiGroupV2Invoker {
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
func (c *RomaClient) DeleteApiV2(request *model.DeleteApiV2Request) (*model.DeleteApiV2Response, error) {
	requestDef := GenReqDefForDeleteApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiV2Response), nil
	}
}

// DeleteApiV2Invoker 删除API
func (c *RomaClient) DeleteApiV2Invoker(request *model.DeleteApiV2Request) *DeleteApiV2Invoker {
	requestDef := GenReqDefForDeleteApiV2()
	return &DeleteApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRequestThrottlingPolicyV2 解除API与流控策略的绑定关系
//
// 解除API与流控策略的绑定关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DisassociateRequestThrottlingPolicyV2(request *model.DisassociateRequestThrottlingPolicyV2Request) (*model.DisassociateRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForDisassociateRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRequestThrottlingPolicyV2Response), nil
	}
}

// DisassociateRequestThrottlingPolicyV2Invoker 解除API与流控策略的绑定关系
func (c *RomaClient) DisassociateRequestThrottlingPolicyV2Invoker(request *model.DisassociateRequestThrottlingPolicyV2Request) *DisassociateRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForDisassociateRequestThrottlingPolicyV2()
	return &DisassociateRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiGroupsV2 查询分组列表
//
// 查询API分组列表。
//
// 如果是租户操作，则查询该租户下所有的分组；如果是管理员操作，则查询的是所有租户的分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApiGroupsV2(request *model.ListApiGroupsV2Request) (*model.ListApiGroupsV2Response, error) {
	requestDef := GenReqDefForListApiGroupsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiGroupsV2Response), nil
	}
}

// ListApiGroupsV2Invoker 查询分组列表
func (c *RomaClient) ListApiGroupsV2Invoker(request *model.ListApiGroupsV2Request) *ListApiGroupsV2Invoker {
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
func (c *RomaClient) ListApiRuntimeDefinitionV2(request *model.ListApiRuntimeDefinitionV2Request) (*model.ListApiRuntimeDefinitionV2Response, error) {
	requestDef := GenReqDefForListApiRuntimeDefinitionV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiRuntimeDefinitionV2Response), nil
	}
}

// ListApiRuntimeDefinitionV2Invoker 查询API运行时定义
func (c *RomaClient) ListApiRuntimeDefinitionV2Invoker(request *model.ListApiRuntimeDefinitionV2Request) *ListApiRuntimeDefinitionV2Invoker {
	requestDef := GenReqDefForListApiRuntimeDefinitionV2()
	return &ListApiRuntimeDefinitionV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersionDetailV2 查看版本详情
//
// 查询某个指定的版本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApiVersionDetailV2(request *model.ListApiVersionDetailV2Request) (*model.ListApiVersionDetailV2Response, error) {
	requestDef := GenReqDefForListApiVersionDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionDetailV2Response), nil
	}
}

// ListApiVersionDetailV2Invoker 查看版本详情
func (c *RomaClient) ListApiVersionDetailV2Invoker(request *model.ListApiVersionDetailV2Request) *ListApiVersionDetailV2Invoker {
	requestDef := GenReqDefForListApiVersionDetailV2()
	return &ListApiVersionDetailV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersionsV2 查询API历史版本列表
//
// 查询某个API的历史版本。每个API在一个环境上最多存在10个历史版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApiVersionsV2(request *model.ListApiVersionsV2Request) (*model.ListApiVersionsV2Response, error) {
	requestDef := GenReqDefForListApiVersionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsV2Response), nil
	}
}

// ListApiVersionsV2Invoker 查询API历史版本列表
func (c *RomaClient) ListApiVersionsV2Invoker(request *model.ListApiVersionsV2Request) *ListApiVersionsV2Invoker {
	requestDef := GenReqDefForListApiVersionsV2()
	return &ListApiVersionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToRequestThrottlingPolicyV2 查看流控策略绑定的API列表
//
// 查询某个流控策略上已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToRequestThrottlingPolicyV2(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) (*model.ListApisBindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToRequestThrottlingPolicyV2Response), nil
	}
}

// ListApisBindedToRequestThrottlingPolicyV2Invoker 查看流控策略绑定的API列表
func (c *RomaClient) ListApisBindedToRequestThrottlingPolicyV2Invoker(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) *ListApisBindedToRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListApisBindedToRequestThrottlingPolicyV2()
	return &ListApisBindedToRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToRequestThrottlingPolicyV2 查看流控策略未绑定的API列表
//
// 查询所有未绑定到该流控策略上的自有API列表。需要API已经发布，未发布的API不予展示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisUnbindedToRequestThrottlingPolicyV2(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToRequestThrottlingPolicyV2Response), nil
	}
}

// ListApisUnbindedToRequestThrottlingPolicyV2Invoker 查看流控策略未绑定的API列表
func (c *RomaClient) ListApisUnbindedToRequestThrottlingPolicyV2Invoker(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) *ListApisUnbindedToRequestThrottlingPolicyV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2()
	return &ListApisUnbindedToRequestThrottlingPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisV2 查询API列表
//
// 查看API列表，返回API详细信息、发布信息等，但不能查看到后端服务信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisV2(request *model.ListApisV2Request) (*model.ListApisV2Response, error) {
	requestDef := GenReqDefForListApisV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisV2Response), nil
	}
}

// ListApisV2Invoker 查询API列表
func (c *RomaClient) ListApisV2Invoker(request *model.ListApisV2Request) *ListApisV2Invoker {
	requestDef := GenReqDefForListApisV2()
	return &ListApisV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRequestThrottlingPoliciesBindedToApiV2 查看API绑定的流控策略列表
//
// 查询某个API绑定的流控策略列表。每个环境上应该最多只有一个流控策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListRequestThrottlingPoliciesBindedToApiV2(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, error) {
	requestDef := GenReqDefForListRequestThrottlingPoliciesBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestThrottlingPoliciesBindedToApiV2Response), nil
	}
}

// ListRequestThrottlingPoliciesBindedToApiV2Invoker 查看API绑定的流控策略列表
func (c *RomaClient) ListRequestThrottlingPoliciesBindedToApiV2Invoker(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) *ListRequestThrottlingPoliciesBindedToApiV2Invoker {
	requestDef := GenReqDefForListRequestThrottlingPoliciesBindedToApiV2()
	return &ListRequestThrottlingPoliciesBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfApiGroupV2 查询分组详情
//
// 查询指定分组的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfApiGroupV2(request *model.ShowDetailsOfApiGroupV2Request) (*model.ShowDetailsOfApiGroupV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiGroupV2Response), nil
	}
}

// ShowDetailsOfApiGroupV2Invoker 查询分组详情
func (c *RomaClient) ShowDetailsOfApiGroupV2Invoker(request *model.ShowDetailsOfApiGroupV2Request) *ShowDetailsOfApiGroupV2Invoker {
	requestDef := GenReqDefForShowDetailsOfApiGroupV2()
	return &ShowDetailsOfApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfApiV2 查询API详情
//
// 查看指定的API的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfApiV2(request *model.ShowDetailsOfApiV2Request) (*model.ShowDetailsOfApiV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfApiV2Response), nil
	}
}

// ShowDetailsOfApiV2Invoker 查询API详情
func (c *RomaClient) ShowDetailsOfApiV2Invoker(request *model.ShowDetailsOfApiV2Request) *ShowDetailsOfApiV2Invoker {
	requestDef := GenReqDefForShowDetailsOfApiV2()
	return &ShowDetailsOfApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApiGroupV2 修改API分组
//
// 修改API分组属性。其中name和remark可修改，其他属性不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateApiGroupV2(request *model.UpdateApiGroupV2Request) (*model.UpdateApiGroupV2Response, error) {
	requestDef := GenReqDefForUpdateApiGroupV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiGroupV2Response), nil
	}
}

// UpdateApiGroupV2Invoker 修改API分组
func (c *RomaClient) UpdateApiGroupV2Invoker(request *model.UpdateApiGroupV2Request) *UpdateApiGroupV2Invoker {
	requestDef := GenReqDefForUpdateApiGroupV2()
	return &UpdateApiGroupV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApiV2 修改API
//
// 修改指定API的信息，包括后端服务信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateApiV2(request *model.UpdateApiV2Request) (*model.UpdateApiV2Response, error) {
	requestDef := GenReqDefForUpdateApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiV2Response), nil
	}
}

// UpdateApiV2Invoker 修改API
func (c *RomaClient) UpdateApiV2Invoker(request *model.UpdateApiV2Request) *UpdateApiV2Invoker {
	requestDef := GenReqDefForUpdateApiV2()
	return &UpdateApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteApiAclBindingV2 批量解除API与ACL策略的绑定
//
// 批量解除API与ACL策略的绑定
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDeleteApiAclBindingV2(request *model.BatchDeleteApiAclBindingV2Request) (*model.BatchDeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForBatchDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteApiAclBindingV2Response), nil
	}
}

// BatchDeleteApiAclBindingV2Invoker 批量解除API与ACL策略的绑定
func (c *RomaClient) BatchDeleteApiAclBindingV2Invoker(request *model.BatchDeleteApiAclBindingV2Request) *BatchDeleteApiAclBindingV2Invoker {
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
func (c *RomaClient) CreateApiAclBindingV2(request *model.CreateApiAclBindingV2Request) (*model.CreateApiAclBindingV2Response, error) {
	requestDef := GenReqDefForCreateApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiAclBindingV2Response), nil
	}
}

// CreateApiAclBindingV2Invoker 将API与ACL策略进行绑定
func (c *RomaClient) CreateApiAclBindingV2Invoker(request *model.CreateApiAclBindingV2Request) *CreateApiAclBindingV2Invoker {
	requestDef := GenReqDefForCreateApiAclBindingV2()
	return &CreateApiAclBindingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApiAclBindingV2 解除API与ACL策略的绑定
//
// 解除某条API与ACL策略的绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteApiAclBindingV2(request *model.DeleteApiAclBindingV2Request) (*model.DeleteApiAclBindingV2Response, error) {
	requestDef := GenReqDefForDeleteApiAclBindingV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiAclBindingV2Response), nil
	}
}

// DeleteApiAclBindingV2Invoker 解除API与ACL策略的绑定
func (c *RomaClient) DeleteApiAclBindingV2Invoker(request *model.DeleteApiAclBindingV2Request) *DeleteApiAclBindingV2Invoker {
	requestDef := GenReqDefForDeleteApiAclBindingV2()
	return &DeleteApiAclBindingV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclPolicyBindedToApiV2 查看API绑定的ACL策略列表
//
// 查看API绑定的ACL策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAclPolicyBindedToApiV2(request *model.ListAclPolicyBindedToApiV2Request) (*model.ListAclPolicyBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAclPolicyBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclPolicyBindedToApiV2Response), nil
	}
}

// ListAclPolicyBindedToApiV2Invoker 查看API绑定的ACL策略列表
func (c *RomaClient) ListAclPolicyBindedToApiV2Invoker(request *model.ListAclPolicyBindedToApiV2Request) *ListAclPolicyBindedToApiV2Invoker {
	requestDef := GenReqDefForListAclPolicyBindedToApiV2()
	return &ListAclPolicyBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToAclPolicyV2 查看ACL策略绑定的API列表
//
// 查看ACL策略绑定的API列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToAclPolicyV2(request *model.ListApisBindedToAclPolicyV2Request) (*model.ListApisBindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAclPolicyV2Response), nil
	}
}

// ListApisBindedToAclPolicyV2Invoker 查看ACL策略绑定的API列表
func (c *RomaClient) ListApisBindedToAclPolicyV2Invoker(request *model.ListApisBindedToAclPolicyV2Request) *ListApisBindedToAclPolicyV2Invoker {
	requestDef := GenReqDefForListApisBindedToAclPolicyV2()
	return &ListApisBindedToAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToAclPolicyV2 查看ACL策略未绑定的API列表
//
// 查看ACL策略未绑定的API列表，需要API已发布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisUnbindedToAclPolicyV2(request *model.ListApisUnbindedToAclPolicyV2Request) (*model.ListApisUnbindedToAclPolicyV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAclPolicyV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAclPolicyV2Response), nil
	}
}

// ListApisUnbindedToAclPolicyV2Invoker 查看ACL策略未绑定的API列表
func (c *RomaClient) ListApisUnbindedToAclPolicyV2Invoker(request *model.ListApisUnbindedToAclPolicyV2Request) *ListApisUnbindedToAclPolicyV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToAclPolicyV2()
	return &ListApisUnbindedToAclPolicyV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelingAuthorizationV2 解除授权
//
// 解除API对APP的授权关系。解除授权后，APP将不再能够调用该API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CancelingAuthorizationV2(request *model.CancelingAuthorizationV2Request) (*model.CancelingAuthorizationV2Response, error) {
	requestDef := GenReqDefForCancelingAuthorizationV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelingAuthorizationV2Response), nil
	}
}

// CancelingAuthorizationV2Invoker 解除授权
func (c *RomaClient) CancelingAuthorizationV2Invoker(request *model.CancelingAuthorizationV2Request) *CancelingAuthorizationV2Invoker {
	requestDef := GenReqDefForCancelingAuthorizationV2()
	return &CancelingAuthorizationV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorizingAppsV2 APP授权
//
// APP创建成功后，还不能访问API，如果想要访问某个环境上的API，需要将该API在该环境上授权给APP。授权成功后，APP即可访问该环境上的这个API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateAuthorizingAppsV2(request *model.CreateAuthorizingAppsV2Request) (*model.CreateAuthorizingAppsV2Response, error) {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizingAppsV2Response), nil
	}
}

// CreateAuthorizingAppsV2Invoker APP授权
func (c *RomaClient) CreateAuthorizingAppsV2Invoker(request *model.CreateAuthorizingAppsV2Request) *CreateAuthorizingAppsV2Invoker {
	requestDef := GenReqDefForCreateAuthorizingAppsV2()
	return &CreateAuthorizingAppsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisBindedToAppV2 查看APP已绑定的API列表
//
// 查询APP已经绑定的API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisBindedToAppV2(request *model.ListApisBindedToAppV2Request) (*model.ListApisBindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisBindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisBindedToAppV2Response), nil
	}
}

// ListApisBindedToAppV2Invoker 查看APP已绑定的API列表
func (c *RomaClient) ListApisBindedToAppV2Invoker(request *model.ListApisBindedToAppV2Request) *ListApisBindedToAppV2Invoker {
	requestDef := GenReqDefForListApisBindedToAppV2()
	return &ListApisBindedToAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisUnbindedToAppV2 查看APP未绑定的API列表
//
// 查询指定环境上某个APP未绑定的API列表，包括自有API和从云市场购买的API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListApisUnbindedToAppV2(request *model.ListApisUnbindedToAppV2Request) (*model.ListApisUnbindedToAppV2Response, error) {
	requestDef := GenReqDefForListApisUnbindedToAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisUnbindedToAppV2Response), nil
	}
}

// ListApisUnbindedToAppV2Invoker 查看APP未绑定的API列表
func (c *RomaClient) ListApisUnbindedToAppV2Invoker(request *model.ListApisUnbindedToAppV2Request) *ListApisUnbindedToAppV2Invoker {
	requestDef := GenReqDefForListApisUnbindedToAppV2()
	return &ListApisUnbindedToAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppsBindedToApiV2 查看API已绑定的APP列表
//
// 查询API绑定的APP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListAppsBindedToApiV2(request *model.ListAppsBindedToApiV2Request) (*model.ListAppsBindedToApiV2Response, error) {
	requestDef := GenReqDefForListAppsBindedToApiV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsBindedToApiV2Response), nil
	}
}

// ListAppsBindedToApiV2Invoker 查看API已绑定的APP列表
func (c *RomaClient) ListAppsBindedToApiV2Invoker(request *model.ListAppsBindedToApiV2Request) *ListAppsBindedToApiV2Invoker {
	requestDef := GenReqDefForListAppsBindedToApiV2()
	return &ListAppsBindedToApiV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDuplicateApisForAppV2 查看APP下路径冲突的api列表
//
// 查询指定APP下路径冲突的api列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDuplicateApisForAppV2(request *model.ListDuplicateApisForAppV2Request) (*model.ListDuplicateApisForAppV2Response, error) {
	requestDef := GenReqDefForListDuplicateApisForAppV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDuplicateApisForAppV2Response), nil
	}
}

// ListDuplicateApisForAppV2Invoker 查看APP下路径冲突的api列表
func (c *RomaClient) ListDuplicateApisForAppV2Invoker(request *model.ListDuplicateApisForAppV2Request) *ListDuplicateApisForAppV2Invoker {
	requestDef := GenReqDefForListDuplicateApisForAppV2()
	return &ListDuplicateApisForAppV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddUserToApp 设置用户成员
//
// - 设置应用的用户成员，为空数组时会清空已有应用成员列表
// - 设置动作为全量更新非增量更新，应用的成员列表都会替换为当次请求的应用成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AddUserToApp(request *model.AddUserToAppRequest) (*model.AddUserToAppResponse, error) {
	requestDef := GenReqDefForAddUserToApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddUserToAppResponse), nil
	}
}

// AddUserToAppInvoker 设置用户成员
func (c *RomaClient) AddUserToAppInvoker(request *model.AddUserToAppRequest) *AddUserToAppInvoker {
	requestDef := GenReqDefForAddUserToApp()
	return &AddUserToAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckAuthUsersOfApp 查询用户成员列表
//
// 查询用户成列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckAuthUsersOfApp(request *model.CheckAuthUsersOfAppRequest) (*model.CheckAuthUsersOfAppResponse, error) {
	requestDef := GenReqDefForCheckAuthUsersOfApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAuthUsersOfAppResponse), nil
	}
}

// CheckAuthUsersOfAppInvoker 查询用户成员列表
func (c *RomaClient) CheckAuthUsersOfAppInvoker(request *model.CheckAuthUsersOfAppRequest) *CheckAuthUsersOfAppInvoker {
	requestDef := GenReqDefForCheckAuthUsersOfApp()
	return &CheckAuthUsersOfAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckCanAuthUsersOfApp 查询候选用户成员
//
// 查询应用的候选用户成员列表,会过滤掉异常状态用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckCanAuthUsersOfApp(request *model.CheckCanAuthUsersOfAppRequest) (*model.CheckCanAuthUsersOfAppResponse, error) {
	requestDef := GenReqDefForCheckCanAuthUsersOfApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCanAuthUsersOfAppResponse), nil
	}
}

// CheckCanAuthUsersOfAppInvoker 查询候选用户成员
func (c *RomaClient) CheckCanAuthUsersOfAppInvoker(request *model.CheckCanAuthUsersOfAppRequest) *CheckCanAuthUsersOfAppInvoker {
	requestDef := GenReqDefForCheckCanAuthUsersOfApp()
	return &CheckCanAuthUsersOfAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRomaAppDetails 查询应用详情
//
// 查询应用详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckRomaAppDetails(request *model.CheckRomaAppDetailsRequest) (*model.CheckRomaAppDetailsResponse, error) {
	requestDef := GenReqDefForCheckRomaAppDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaAppDetailsResponse), nil
	}
}

// CheckRomaAppDetailsInvoker 查询应用详情
func (c *RomaClient) CheckRomaAppDetailsInvoker(request *model.CheckRomaAppDetailsRequest) *CheckRomaAppDetailsInvoker {
	requestDef := GenReqDefForCheckRomaAppDetails()
	return &CheckRomaAppDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRomaAppSecret 查询应用密钥
//
// 查询应用密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckRomaAppSecret(request *model.CheckRomaAppSecretRequest) (*model.CheckRomaAppSecretResponse, error) {
	requestDef := GenReqDefForCheckRomaAppSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaAppSecretResponse), nil
	}
}

// CheckRomaAppSecretInvoker 查询应用密钥
func (c *RomaClient) CheckRomaAppSecretInvoker(request *model.CheckRomaAppSecretRequest) *CheckRomaAppSecretInvoker {
	requestDef := GenReqDefForCheckRomaAppSecret()
	return &CheckRomaAppSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRomaApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateRomaApp(request *model.CreateRomaAppRequest) (*model.CreateRomaAppResponse, error) {
	requestDef := GenReqDefForCreateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRomaAppResponse), nil
	}
}

// CreateRomaAppInvoker 创建应用
func (c *RomaClient) CreateRomaAppInvoker(request *model.CreateRomaAppRequest) *CreateRomaAppInvoker {
	requestDef := GenReqDefForCreateRomaApp()
	return &CreateRomaAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRomaApp 删除应用
//
// 删除单个应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteRomaApp(request *model.DeleteRomaAppRequest) (*model.DeleteRomaAppResponse, error) {
	requestDef := GenReqDefForDeleteRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRomaAppResponse), nil
	}
}

// DeleteRomaAppInvoker 删除应用
func (c *RomaClient) DeleteRomaAppInvoker(request *model.DeleteRomaAppRequest) *DeleteRomaAppInvoker {
	requestDef := GenReqDefForDeleteRomaApp()
	return &DeleteRomaAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRomaApp 查询应用列表
//
// 查询应用列表，支持条件查询，所有条件是并且的关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListRomaApp(request *model.ListRomaAppRequest) (*model.ListRomaAppResponse, error) {
	requestDef := GenReqDefForListRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRomaAppResponse), nil
	}
}

// ListRomaAppInvoker 查询应用列表
func (c *RomaClient) ListRomaAppInvoker(request *model.ListRomaAppRequest) *ListRomaAppInvoker {
	requestDef := GenReqDefForListRomaApp()
	return &ListRomaAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetRomaAppSecret 重置应用密钥
//
// 重置应用密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ResetRomaAppSecret(request *model.ResetRomaAppSecretRequest) (*model.ResetRomaAppSecretResponse, error) {
	requestDef := GenReqDefForResetRomaAppSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetRomaAppSecretResponse), nil
	}
}

// ResetRomaAppSecretInvoker 重置应用密钥
func (c *RomaClient) ResetRomaAppSecretInvoker(request *model.ResetRomaAppSecretRequest) *ResetRomaAppSecretInvoker {
	requestDef := GenReqDefForResetRomaAppSecret()
	return &ResetRomaAppSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRomaApp 更新应用
//
// 更新应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateRomaApp(request *model.UpdateRomaAppRequest) (*model.UpdateRomaAppResponse, error) {
	requestDef := GenReqDefForUpdateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRomaAppResponse), nil
	}
}

// UpdateRomaAppInvoker 更新应用
func (c *RomaClient) UpdateRomaAppInvoker(request *model.UpdateRomaAppRequest) *UpdateRomaAppInvoker {
	requestDef := GenReqDefForUpdateRomaApp()
	return &UpdateRomaAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateRomaApp 校验应用是否存在
//
// 校验指定条件的应用是否存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ValidateRomaApp(request *model.ValidateRomaAppRequest) (*model.ValidateRomaAppResponse, error) {
	requestDef := GenReqDefForValidateRomaApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateRomaAppResponse), nil
	}
}

// ValidateRomaAppInvoker 校验应用是否存在
func (c *RomaClient) ValidateRomaAppInvoker(request *model.ValidateRomaAppRequest) *ValidateRomaAppInvoker {
	requestDef := GenReqDefForValidateRomaApp()
	return &ValidateRomaAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckAssetJobStatus 查询作业进度
//
// 查询作业进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckAssetJobStatus(request *model.CheckAssetJobStatusRequest) (*model.CheckAssetJobStatusResponse, error) {
	requestDef := GenReqDefForCheckAssetJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckAssetJobStatusResponse), nil
	}
}

// CheckAssetJobStatusInvoker 查询作业进度
func (c *RomaClient) CheckAssetJobStatusInvoker(request *model.CheckAssetJobStatusRequest) *CheckAssetJobStatusInvoker {
	requestDef := GenReqDefForCheckAssetJobStatus()
	return &CheckAssetJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAsset 批量删除资产
//
// 批量删除资产
// - 创建批量删除指定条件的资产的作业任务
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产删除作业任务，没有Running状态的作业任务存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// DeleteAssetInvoker 批量删除资产
func (c *RomaClient) DeleteAssetInvoker(request *model.DeleteAssetRequest) *DeleteAssetInvoker {
	requestDef := GenReqDefForDeleteAsset()
	return &DeleteAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAssetArchive 下载资产包
//
// - 导出作业执行成功后，通过该接口获取导出作业产生的资产包，仅能下载一次
// - 可先压缩后存在数据库，下载后删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DownloadAssetArchive(request *model.DownloadAssetArchiveRequest) (*model.DownloadAssetArchiveResponse, error) {
	requestDef := GenReqDefForDownloadAssetArchive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAssetArchiveResponse), nil
	}
}

// DownloadAssetArchiveInvoker 下载资产包
func (c *RomaClient) DownloadAssetArchiveInvoker(request *model.DownloadAssetArchiveRequest) *DownloadAssetArchiveInvoker {
	requestDef := GenReqDefForDownloadAssetArchive()
	return &DownloadAssetArchiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportAsset 批量导出资产
//
// 批量导出资产
// - 创建批量导出指定条件的资产的作业任务
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产导出作业任务，没有Running状态的作业任务存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ExportAsset(request *model.ExportAssetRequest) (*model.ExportAssetResponse, error) {
	requestDef := GenReqDefForExportAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportAssetResponse), nil
	}
}

// ExportAssetInvoker 批量导出资产
func (c *RomaClient) ExportAssetInvoker(request *model.ExportAssetRequest) *ExportAssetInvoker {
	requestDef := GenReqDefForExportAsset()
	return &ExportAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportAsset 导入资产
//
// - 创建导入资产作业任务，资产版本和具体哪些资产从资产内容里读取
// - 最大支持100个应用和任务
// - 一个用户同一时刻只能创建一个资产导入作业任务，没有Running状态的作业任务存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ImportAsset(request *model.ImportAssetRequest) (*model.ImportAssetResponse, error) {
	requestDef := GenReqDefForImportAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportAssetResponse), nil
	}
}

// ImportAssetInvoker 导入资产
func (c *RomaClient) ImportAssetInvoker(request *model.ImportAssetRequest) *ImportAssetInvoker {
	requestDef := GenReqDefForImportAsset()
	return &ImportAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDictionary 查询字典详情
//
// 查询字典详情,
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckDictionary(request *model.CheckDictionaryRequest) (*model.CheckDictionaryResponse, error) {
	requestDef := GenReqDefForCheckDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDictionaryResponse), nil
	}
}

// CheckDictionaryInvoker 查询字典详情
func (c *RomaClient) CheckDictionaryInvoker(request *model.CheckDictionaryRequest) *CheckDictionaryInvoker {
	requestDef := GenReqDefForCheckDictionary()
	return &CheckDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDictionary 创建字典
//
// 创建字典
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateDictionary(request *model.CreateDictionaryRequest) (*model.CreateDictionaryResponse, error) {
	requestDef := GenReqDefForCreateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDictionaryResponse), nil
	}
}

// CreateDictionaryInvoker 创建字典
func (c *RomaClient) CreateDictionaryInvoker(request *model.CreateDictionaryRequest) *CreateDictionaryInvoker {
	requestDef := GenReqDefForCreateDictionary()
	return &CreateDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDictionary 删除字典
//
// 删除单个字典，会同时删除该字典的所有子字典
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteDictionary(request *model.DeleteDictionaryRequest) (*model.DeleteDictionaryResponse, error) {
	requestDef := GenReqDefForDeleteDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDictionaryResponse), nil
	}
}

// DeleteDictionaryInvoker 删除字典
func (c *RomaClient) DeleteDictionaryInvoker(request *model.DeleteDictionaryRequest) *DeleteDictionaryInvoker {
	requestDef := GenReqDefForDeleteDictionary()
	return &DeleteDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDictionary 查询字典列表
//
// 查询字典列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListDictionary(request *model.ListDictionaryRequest) (*model.ListDictionaryResponse, error) {
	requestDef := GenReqDefForListDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDictionaryResponse), nil
	}
}

// ListDictionaryInvoker 查询字典列表
func (c *RomaClient) ListDictionaryInvoker(request *model.ListDictionaryRequest) *ListDictionaryInvoker {
	requestDef := GenReqDefForListDictionary()
	return &ListDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDictionary 更新字典
//
// 更新字典
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateDictionary(request *model.UpdateDictionaryRequest) (*model.UpdateDictionaryResponse, error) {
	requestDef := GenReqDefForUpdateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDictionaryResponse), nil
	}
}

// UpdateDictionaryInvoker 更新字典
func (c *RomaClient) UpdateDictionaryInvoker(request *model.UpdateDictionaryRequest) *UpdateDictionaryInvoker {
	requestDef := GenReqDefForUpdateDictionary()
	return &UpdateDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateDictionary 校验字典是否存在
//
// 校验指定条件的字典是否存在，支持字典名称和字典编码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ValidateDictionary(request *model.ValidateDictionaryRequest) (*model.ValidateDictionaryResponse, error) {
	requestDef := GenReqDefForValidateDictionary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateDictionaryResponse), nil
	}
}

// ValidateDictionaryInvoker 校验字典是否存在
func (c *RomaClient) ValidateDictionaryInvoker(request *model.ValidateDictionaryRequest) *ValidateDictionaryInvoker {
	requestDef := GenReqDefForValidateDictionary()
	return &ValidateDictionaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckRomaInstanceListV2 查询实例列表
//
// 获取符合条件的服务实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CheckRomaInstanceListV2(request *model.CheckRomaInstanceListV2Request) (*model.CheckRomaInstanceListV2Response, error) {
	requestDef := GenReqDefForCheckRomaInstanceListV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckRomaInstanceListV2Response), nil
	}
}

// CheckRomaInstanceListV2Invoker 查询实例列表
func (c *RomaClient) CheckRomaInstanceListV2Invoker(request *model.CheckRomaInstanceListV2Request) *CheckRomaInstanceListV2Invoker {
	requestDef := GenReqDefForCheckRomaInstanceListV2()
	return &CheckRomaInstanceListV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMqsInstance 查询MQS实例列表
//
// 查询MQS实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListMqsInstance(request *model.ListMqsInstanceRequest) (*model.ListMqsInstanceResponse, error) {
	requestDef := GenReqDefForListMqsInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMqsInstanceResponse), nil
	}
}

// ListMqsInstanceInvoker 查询MQS实例列表
func (c *RomaClient) ListMqsInstanceInvoker(request *model.ListMqsInstanceRequest) *ListMqsInstanceInvoker {
	requestDef := GenReqDefForListMqsInstance()
	return &ListMqsInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMqsInstance 查询MQS实例详情
//
// 查询指定MQS实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowMqsInstance(request *model.ShowMqsInstanceRequest) (*model.ShowMqsInstanceResponse, error) {
	requestDef := GenReqDefForShowMqsInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMqsInstanceResponse), nil
	}
}

// ShowMqsInstanceInvoker 查询MQS实例详情
func (c *RomaClient) ShowMqsInstanceInvoker(request *model.ShowMqsInstanceRequest) *ShowMqsInstanceInvoker {
	requestDef := GenReqDefForShowMqsInstance()
	return &ShowMqsInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportApiDefinitionsV2 导出API
//
// 导出分组下API的定义信息，导出文件内容符合swagger标准规范。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ExportApiDefinitionsV2(request *model.ExportApiDefinitionsV2Request) (*model.ExportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportApiDefinitionsV2Response), nil
	}
}

// ExportApiDefinitionsV2Invoker 导出API
func (c *RomaClient) ExportApiDefinitionsV2Invoker(request *model.ExportApiDefinitionsV2Request) *ExportApiDefinitionsV2Invoker {
	requestDef := GenReqDefForExportApiDefinitionsV2()
	return &ExportApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportLiveDataApiDefinitionsV2 导出自定义后端API
//
// 导出自定义后端API，导出文件内容符合swagger标准规范。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ExportLiveDataApiDefinitionsV2(request *model.ExportLiveDataApiDefinitionsV2Request) (*model.ExportLiveDataApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForExportLiveDataApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportLiveDataApiDefinitionsV2Response), nil
	}
}

// ExportLiveDataApiDefinitionsV2Invoker 导出自定义后端API
func (c *RomaClient) ExportLiveDataApiDefinitionsV2Invoker(request *model.ExportLiveDataApiDefinitionsV2Request) *ExportLiveDataApiDefinitionsV2Invoker {
	requestDef := GenReqDefForExportLiveDataApiDefinitionsV2()
	return &ExportLiveDataApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportApiDefinitionsV2 导入API
//
// 导入API。导入文件内容需要符合swagger标准规范，自定义扩展字段请参考用户指南的“附录：前端API的Swagger扩展定义”章节。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ImportApiDefinitionsV2(request *model.ImportApiDefinitionsV2Request) (*model.ImportApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApiDefinitionsV2Response), nil
	}
}

// ImportApiDefinitionsV2Invoker 导入API
func (c *RomaClient) ImportApiDefinitionsV2Invoker(request *model.ImportApiDefinitionsV2Request) *ImportApiDefinitionsV2Invoker {
	requestDef := GenReqDefForImportApiDefinitionsV2()
	return &ImportApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportLiveDataApiDefinitionsV2 导入自定义后端API
//
// 导入自定义后端API。导入文件内容需要符合swagger标准规范，自定义扩展字段请参考用户指南的“附录：后端API的Swagger扩展定义”章节
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ImportLiveDataApiDefinitionsV2(request *model.ImportLiveDataApiDefinitionsV2Request) (*model.ImportLiveDataApiDefinitionsV2Response, error) {
	requestDef := GenReqDefForImportLiveDataApiDefinitionsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportLiveDataApiDefinitionsV2Response), nil
	}
}

// ImportLiveDataApiDefinitionsV2Invoker 导入自定义后端API
func (c *RomaClient) ImportLiveDataApiDefinitionsV2Invoker(request *model.ImportLiveDataApiDefinitionsV2Request) *ImportLiveDataApiDefinitionsV2Invoker {
	requestDef := GenReqDefForImportLiveDataApiDefinitionsV2()
	return &ImportLiveDataApiDefinitionsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddingBackendInstancesV2 添加或更新后端实例
//
// 为指定的VPC通道添加后端实例
//
// 若指定地址的后端实例已存在，则更新对应后端实例信息。若请求体中包含多个重复地址的后端实例定义，则使用第一个定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) AddingBackendInstancesV2(request *model.AddingBackendInstancesV2Request) (*model.AddingBackendInstancesV2Response, error) {
	requestDef := GenReqDefForAddingBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddingBackendInstancesV2Response), nil
	}
}

// AddingBackendInstancesV2Invoker 添加或更新后端实例
func (c *RomaClient) AddingBackendInstancesV2Invoker(request *model.AddingBackendInstancesV2Request) *AddingBackendInstancesV2Invoker {
	requestDef := GenReqDefForAddingBackendInstancesV2()
	return &AddingBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisableMembers 批量修改后端服务器状态不可用
//
// 批量修改后端服务器状态不可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchDisableMembers(request *model.BatchDisableMembersRequest) (*model.BatchDisableMembersResponse, error) {
	requestDef := GenReqDefForBatchDisableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableMembersResponse), nil
	}
}

// BatchDisableMembersInvoker 批量修改后端服务器状态不可用
func (c *RomaClient) BatchDisableMembersInvoker(request *model.BatchDisableMembersRequest) *BatchDisableMembersInvoker {
	requestDef := GenReqDefForBatchDisableMembers()
	return &BatchDisableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableMembers 批量修改后端服务器状态可用
//
// 批量修改后端服务器状态可用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) BatchEnableMembers(request *model.BatchEnableMembersRequest) (*model.BatchEnableMembersResponse, error) {
	requestDef := GenReqDefForBatchEnableMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableMembersResponse), nil
	}
}

// BatchEnableMembersInvoker 批量修改后端服务器状态可用
func (c *RomaClient) BatchEnableMembersInvoker(request *model.BatchEnableMembersRequest) *BatchEnableMembersInvoker {
	requestDef := GenReqDefForBatchEnableMembers()
	return &BatchEnableMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMemberGroup 添加或更新VPC通道后端服务器组
//
// 在ROMA Connect APIC中创建VPC通道后端服务器组，VPC通道后端实例可以选择是否关联后端实例服务器组，以便管理后端服务器节点。
//
// 若指定名称的后端服务器组已存在，则更新对应后端服务器组信息。若请求体中包含多个重复名称的后端服务器定义，则使用第一个定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateMemberGroup(request *model.CreateMemberGroupRequest) (*model.CreateMemberGroupResponse, error) {
	requestDef := GenReqDefForCreateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberGroupResponse), nil
	}
}

// CreateMemberGroupInvoker 添加或更新VPC通道后端服务器组
func (c *RomaClient) CreateMemberGroupInvoker(request *model.CreateMemberGroupRequest) *CreateMemberGroupInvoker {
	requestDef := GenReqDefForCreateMemberGroup()
	return &CreateMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectVpcChannel 项目下创建VPC通道
//
// 创建相同的VPC通道关联到多个实例。同一个项目下VPC通道名称不可重复。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProjectVpcChannel(request *model.CreateProjectVpcChannelRequest) (*model.CreateProjectVpcChannelResponse, error) {
	requestDef := GenReqDefForCreateProjectVpcChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectVpcChannelResponse), nil
	}
}

// CreateProjectVpcChannelInvoker 项目下创建VPC通道
func (c *RomaClient) CreateProjectVpcChannelInvoker(request *model.CreateProjectVpcChannelRequest) *CreateProjectVpcChannelInvoker {
	requestDef := GenReqDefForCreateProjectVpcChannel()
	return &CreateProjectVpcChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProjectVpcChannelSyncs 项目下同步VPC通道
//
// 同步VPC通道到多个实例。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateProjectVpcChannelSyncs(request *model.CreateProjectVpcChannelSyncsRequest) (*model.CreateProjectVpcChannelSyncsResponse, error) {
	requestDef := GenReqDefForCreateProjectVpcChannelSyncs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectVpcChannelSyncsResponse), nil
	}
}

// CreateProjectVpcChannelSyncsInvoker 项目下同步VPC通道
func (c *RomaClient) CreateProjectVpcChannelSyncsInvoker(request *model.CreateProjectVpcChannelSyncsRequest) *CreateProjectVpcChannelSyncsInvoker {
	requestDef := GenReqDefForCreateProjectVpcChannelSyncs()
	return &CreateProjectVpcChannelSyncsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcChannelV2 创建VPC通道
//
// 在ROMA Connect APIC中创建连接私有VPC资源的通道，并在创建API时将后端节点配置为使用这些VPC通道，以便ROMA Connect APIC直接访问私有VPC资源。
// &gt; 每个用户默认最多创建200个VPC通道，如需支持更多请联系技术支持调整配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) CreateVpcChannelV2(request *model.CreateVpcChannelV2Request) (*model.CreateVpcChannelV2Response, error) {
	requestDef := GenReqDefForCreateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcChannelV2Response), nil
	}
}

// CreateVpcChannelV2Invoker 创建VPC通道
func (c *RomaClient) CreateVpcChannelV2Invoker(request *model.CreateVpcChannelV2Request) *CreateVpcChannelV2Invoker {
	requestDef := GenReqDefForCreateVpcChannelV2()
	return &CreateVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackendInstanceV2 删除后端实例
//
// 删除指定VPC通道中的后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteBackendInstanceV2(request *model.DeleteBackendInstanceV2Request) (*model.DeleteBackendInstanceV2Response, error) {
	requestDef := GenReqDefForDeleteBackendInstanceV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackendInstanceV2Response), nil
	}
}

// DeleteBackendInstanceV2Invoker 删除后端实例
func (c *RomaClient) DeleteBackendInstanceV2Invoker(request *model.DeleteBackendInstanceV2Request) *DeleteBackendInstanceV2Invoker {
	requestDef := GenReqDefForDeleteBackendInstanceV2()
	return &DeleteBackendInstanceV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMemberGroup 删除VPC通道后端服务器组
//
// 删除指定的VPC通道后端服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteMemberGroup(request *model.DeleteMemberGroupRequest) (*model.DeleteMemberGroupResponse, error) {
	requestDef := GenReqDefForDeleteMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberGroupResponse), nil
	}
}

// DeleteMemberGroupInvoker 删除VPC通道后端服务器组
func (c *RomaClient) DeleteMemberGroupInvoker(request *model.DeleteMemberGroupRequest) *DeleteMemberGroupInvoker {
	requestDef := GenReqDefForDeleteMemberGroup()
	return &DeleteMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcChannelV2 删除VPC通道
//
// 删除指定的VPC通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) DeleteVpcChannelV2(request *model.DeleteVpcChannelV2Request) (*model.DeleteVpcChannelV2Response, error) {
	requestDef := GenReqDefForDeleteVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcChannelV2Response), nil
	}
}

// DeleteVpcChannelV2Invoker 删除VPC通道
func (c *RomaClient) DeleteVpcChannelV2Invoker(request *model.DeleteVpcChannelV2Request) *DeleteVpcChannelV2Invoker {
	requestDef := GenReqDefForDeleteVpcChannelV2()
	return &DeleteVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackendInstancesV2 查看后端实例列表
//
// 查看指定VPC通道的后端实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListBackendInstancesV2(request *model.ListBackendInstancesV2Request) (*model.ListBackendInstancesV2Response, error) {
	requestDef := GenReqDefForListBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackendInstancesV2Response), nil
	}
}

// ListBackendInstancesV2Invoker 查看后端实例列表
func (c *RomaClient) ListBackendInstancesV2Invoker(request *model.ListBackendInstancesV2Request) *ListBackendInstancesV2Invoker {
	requestDef := GenReqDefForListBackendInstancesV2()
	return &ListBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMemberGroups 查询VPC通道后端云服务组列表
//
// 查询VPC通道后端云服务组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListMemberGroups(request *model.ListMemberGroupsRequest) (*model.ListMemberGroupsResponse, error) {
	requestDef := GenReqDefForListMemberGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMemberGroupsResponse), nil
	}
}

// ListMemberGroupsInvoker 查询VPC通道后端云服务组列表
func (c *RomaClient) ListMemberGroupsInvoker(request *model.ListMemberGroupsRequest) *ListMemberGroupsInvoker {
	requestDef := GenReqDefForListMemberGroups()
	return &ListMemberGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectVpcChannelsV2 查询项目下所有实例的VPC通道列表
//
// 查询项目下所有实例的VPC通道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListProjectVpcChannelsV2(request *model.ListProjectVpcChannelsV2Request) (*model.ListProjectVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListProjectVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectVpcChannelsV2Response), nil
	}
}

// ListProjectVpcChannelsV2Invoker 查询项目下所有实例的VPC通道列表
func (c *RomaClient) ListProjectVpcChannelsV2Invoker(request *model.ListProjectVpcChannelsV2Request) *ListProjectVpcChannelsV2Invoker {
	requestDef := GenReqDefForListProjectVpcChannelsV2()
	return &ListProjectVpcChannelsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcChannelsV2 查询VPC通道列表
//
// 查看VPC通道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ListVpcChannelsV2(request *model.ListVpcChannelsV2Request) (*model.ListVpcChannelsV2Response, error) {
	requestDef := GenReqDefForListVpcChannelsV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcChannelsV2Response), nil
	}
}

// ListVpcChannelsV2Invoker 查询VPC通道列表
func (c *RomaClient) ListVpcChannelsV2Invoker(request *model.ListVpcChannelsV2Request) *ListVpcChannelsV2Invoker {
	requestDef := GenReqDefForListVpcChannelsV2()
	return &ListVpcChannelsV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfMemberGroup 查看VPC通道后端服务器组详情
//
// 查看指定的VPC通道后端服务器组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfMemberGroup(request *model.ShowDetailsOfMemberGroupRequest) (*model.ShowDetailsOfMemberGroupResponse, error) {
	requestDef := GenReqDefForShowDetailsOfMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfMemberGroupResponse), nil
	}
}

// ShowDetailsOfMemberGroupInvoker 查看VPC通道后端服务器组详情
func (c *RomaClient) ShowDetailsOfMemberGroupInvoker(request *model.ShowDetailsOfMemberGroupRequest) *ShowDetailsOfMemberGroupInvoker {
	requestDef := GenReqDefForShowDetailsOfMemberGroup()
	return &ShowDetailsOfMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDetailsOfVpcChannelV2 查看VPC通道详情
//
// 查看指定的VPC通道详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) ShowDetailsOfVpcChannelV2(request *model.ShowDetailsOfVpcChannelV2Request) (*model.ShowDetailsOfVpcChannelV2Response, error) {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDetailsOfVpcChannelV2Response), nil
	}
}

// ShowDetailsOfVpcChannelV2Invoker 查看VPC通道详情
func (c *RomaClient) ShowDetailsOfVpcChannelV2Invoker(request *model.ShowDetailsOfVpcChannelV2Request) *ShowDetailsOfVpcChannelV2Invoker {
	requestDef := GenReqDefForShowDetailsOfVpcChannelV2()
	return &ShowDetailsOfVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBackendInstancesV2 更新后端实例
//
// 更新指定的VPC通道的后端实例。更新时，使用传入的请求参数对对应云服务组的后端实例进行全量覆盖修改。若未指定修改的云服务器组，则进行全量覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateBackendInstancesV2(request *model.UpdateBackendInstancesV2Request) (*model.UpdateBackendInstancesV2Response, error) {
	requestDef := GenReqDefForUpdateBackendInstancesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackendInstancesV2Response), nil
	}
}

// UpdateBackendInstancesV2Invoker 更新后端实例
func (c *RomaClient) UpdateBackendInstancesV2Invoker(request *model.UpdateBackendInstancesV2Request) *UpdateBackendInstancesV2Invoker {
	requestDef := GenReqDefForUpdateBackendInstancesV2()
	return &UpdateBackendInstancesV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthCheck 修改VPC通道健康检查
//
// 修改VPC通道健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateHealthCheck(request *model.UpdateHealthCheckRequest) (*model.UpdateHealthCheckResponse, error) {
	requestDef := GenReqDefForUpdateHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthCheckResponse), nil
	}
}

// UpdateHealthCheckInvoker 修改VPC通道健康检查
func (c *RomaClient) UpdateHealthCheckInvoker(request *model.UpdateHealthCheckRequest) *UpdateHealthCheckInvoker {
	requestDef := GenReqDefForUpdateHealthCheck()
	return &UpdateHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMemberGroup 更新VPC通道后端服务器组
//
// 更新指定VPC通道后端服务器组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateMemberGroup(request *model.UpdateMemberGroupRequest) (*model.UpdateMemberGroupResponse, error) {
	requestDef := GenReqDefForUpdateMemberGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberGroupResponse), nil
	}
}

// UpdateMemberGroupInvoker 更新VPC通道后端服务器组
func (c *RomaClient) UpdateMemberGroupInvoker(request *model.UpdateMemberGroupRequest) *UpdateMemberGroupInvoker {
	requestDef := GenReqDefForUpdateMemberGroup()
	return &UpdateMemberGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProjectVpcChannel 项目下批量修改VPC通道
//
// 项目下根据VPC通道名称批量修改多个多个实例下的VPC通道。若实例下不存在该VPC通道则创建。注意：实例特性vpc_name_modifiable配置为off时才可使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateProjectVpcChannel(request *model.UpdateProjectVpcChannelRequest) (*model.UpdateProjectVpcChannelResponse, error) {
	requestDef := GenReqDefForUpdateProjectVpcChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectVpcChannelResponse), nil
	}
}

// UpdateProjectVpcChannelInvoker 项目下批量修改VPC通道
func (c *RomaClient) UpdateProjectVpcChannelInvoker(request *model.UpdateProjectVpcChannelRequest) *UpdateProjectVpcChannelInvoker {
	requestDef := GenReqDefForUpdateProjectVpcChannel()
	return &UpdateProjectVpcChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpcChannelV2 更新VPC通道
//
// 更新指定VPC通道的参数
//
// 使用传入的后端实例列表对VPC通道进行全量覆盖，若后端实例列表为空，则会全量删除已有的后端实例；
//
// 使用传入的后端服务器组列表对VPC通道进行全量覆盖，若后端服务器组列表为空，则会全量删除已有的服务器组；
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RomaClient) UpdateVpcChannelV2(request *model.UpdateVpcChannelV2Request) (*model.UpdateVpcChannelV2Response, error) {
	requestDef := GenReqDefForUpdateVpcChannelV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcChannelV2Response), nil
	}
}

// UpdateVpcChannelV2Invoker 更新VPC通道
func (c *RomaClient) UpdateVpcChannelV2Invoker(request *model.UpdateVpcChannelV2Request) *UpdateVpcChannelV2Invoker {
	requestDef := GenReqDefForUpdateVpcChannelV2()
	return &UpdateVpcChannelV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
