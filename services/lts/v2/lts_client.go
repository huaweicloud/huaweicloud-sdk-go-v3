package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
)

type LtsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewLtsClient(hcClient *httpclient.HcHttpClient) *LtsClient {
	return &LtsClient{HcClient: hcClient}
}

func LtsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAccessConfig 创建日志接入
//
// 创建日志接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateAccessConfig(request *model.CreateAccessConfigRequest) (*model.CreateAccessConfigResponse, error) {
	requestDef := GenReqDefForCreateAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessConfigResponse), nil
	}
}

// CreateAccessConfigInvoker 创建日志接入
func (c *LtsClient) CreateAccessConfigInvoker(request *model.CreateAccessConfigRequest) *CreateAccessConfigInvoker {
	requestDef := GenReqDefForCreateAccessConfig()
	return &CreateAccessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgencyAccess 新建跨账号日志接入
//
// 新建跨账号日志接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateAgencyAccess(request *model.CreateAgencyAccessRequest) (*model.CreateAgencyAccessResponse, error) {
	requestDef := GenReqDefForCreateAgencyAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyAccessResponse), nil
	}
}

// CreateAgencyAccessInvoker 新建跨账号日志接入
func (c *LtsClient) CreateAgencyAccessInvoker(request *model.CreateAgencyAccessRequest) *CreateAgencyAccessInvoker {
	requestDef := GenReqDefForCreateAgencyAccess()
	return &CreateAgencyAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDashBoard 创建仪表盘
//
// 创建仪表盘
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateDashBoard(request *model.CreateDashBoardRequest) (*model.CreateDashBoardResponse, error) {
	requestDef := GenReqDefForCreateDashBoard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDashBoardResponse), nil
	}
}

// CreateDashBoardInvoker 创建仪表盘
func (c *LtsClient) CreateDashBoardInvoker(request *model.CreateDashBoardRequest) *CreateDashBoardInvoker {
	requestDef := GenReqDefForCreateDashBoard()
	return &CreateDashBoardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDashboardGroup 创建仪表盘分组
//
// 创建仪表盘分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateDashboardGroup(request *model.CreateDashboardGroupRequest) (*model.CreateDashboardGroupResponse, error) {
	requestDef := GenReqDefForCreateDashboardGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDashboardGroupResponse), nil
	}
}

// CreateDashboardGroupInvoker 创建仪表盘分组
func (c *LtsClient) CreateDashboardGroupInvoker(request *model.CreateDashboardGroupRequest) *CreateDashboardGroupInvoker {
	requestDef := GenReqDefForCreateDashboardGroup()
	return &CreateDashboardGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHostGroup 创建主机组
//
// 创建主机组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateHostGroup(request *model.CreateHostGroupRequest) (*model.CreateHostGroupResponse, error) {
	requestDef := GenReqDefForCreateHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostGroupResponse), nil
	}
}

// CreateHostGroupInvoker 创建主机组
func (c *LtsClient) CreateHostGroupInvoker(request *model.CreateHostGroupRequest) *CreateHostGroupInvoker {
	requestDef := GenReqDefForCreateHostGroup()
	return &CreateHostGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKeywordsAlarmRule 创建关键词告警规则
//
// 该接口用于创建关键词告警，目前每个帐户最多可以创建共200个关键词告警与SQL告警。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateKeywordsAlarmRule(request *model.CreateKeywordsAlarmRuleRequest) (*model.CreateKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForCreateKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeywordsAlarmRuleResponse), nil
	}
}

// CreateKeywordsAlarmRuleInvoker 创建关键词告警规则
func (c *LtsClient) CreateKeywordsAlarmRuleInvoker(request *model.CreateKeywordsAlarmRuleRequest) *CreateKeywordsAlarmRuleInvoker {
	requestDef := GenReqDefForCreateKeywordsAlarmRule()
	return &CreateKeywordsAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogDumpObs 创建日志转储（旧版）
//
// 该接口用于将指定的一个或多个日志流的日志转储到OBS服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateLogDumpObs(request *model.CreateLogDumpObsRequest) (*model.CreateLogDumpObsResponse, error) {
	requestDef := GenReqDefForCreateLogDumpObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogDumpObsResponse), nil
	}
}

// CreateLogDumpObsInvoker 创建日志转储（旧版）
func (c *LtsClient) CreateLogDumpObsInvoker(request *model.CreateLogDumpObsRequest) *CreateLogDumpObsInvoker {
	requestDef := GenReqDefForCreateLogDumpObs()
	return &CreateLogDumpObsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogGroup 创建日志组
//
// 该接口用于创建一个日志组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateLogGroup(request *model.CreateLogGroupRequest) (*model.CreateLogGroupResponse, error) {
	requestDef := GenReqDefForCreateLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogGroupResponse), nil
	}
}

// CreateLogGroupInvoker 创建日志组
func (c *LtsClient) CreateLogGroupInvoker(request *model.CreateLogGroupRequest) *CreateLogGroupInvoker {
	requestDef := GenReqDefForCreateLogGroup()
	return &CreateLogGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogStream 创建日志流
//
// 该接口用于创建某个指定日志组下的日志流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateLogStream(request *model.CreateLogStreamRequest) (*model.CreateLogStreamResponse, error) {
	requestDef := GenReqDefForCreateLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogStreamResponse), nil
	}
}

// CreateLogStreamInvoker 创建日志流
func (c *LtsClient) CreateLogStreamInvoker(request *model.CreateLogStreamRequest) *CreateLogStreamInvoker {
	requestDef := GenReqDefForCreateLogStream()
	return &CreateLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogStreamIndex 向指定流创建索引
//
// 向指定流创建索引
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateLogStreamIndex(request *model.CreateLogStreamIndexRequest) (*model.CreateLogStreamIndexResponse, error) {
	requestDef := GenReqDefForCreateLogStreamIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogStreamIndexResponse), nil
	}
}

// CreateLogStreamIndexInvoker 向指定流创建索引
func (c *LtsClient) CreateLogStreamIndexInvoker(request *model.CreateLogStreamIndexRequest) *CreateLogStreamIndexInvoker {
	requestDef := GenReqDefForCreateLogStreamIndex()
	return &CreateLogStreamIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotificationTemplate 创建消息模板
//
// 该接口用于创建通知模板，目前每个帐户最多可以创建共100个通知模板，创建后名称不可修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateNotificationTemplate(request *model.CreateNotificationTemplateRequest) (*model.CreateNotificationTemplateResponse, error) {
	requestDef := GenReqDefForCreateNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotificationTemplateResponse), nil
	}
}

// CreateNotificationTemplateInvoker 创建消息模板
func (c *LtsClient) CreateNotificationTemplateInvoker(request *model.CreateNotificationTemplateRequest) *CreateNotificationTemplateInvoker {
	requestDef := GenReqDefForCreateNotificationTemplate()
	return &CreateNotificationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSearchCriterias 添加快速查询
//
// 添加快速查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateSearchCriterias(request *model.CreateSearchCriteriasRequest) (*model.CreateSearchCriteriasResponse, error) {
	requestDef := GenReqDefForCreateSearchCriterias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSearchCriteriasResponse), nil
	}
}

// CreateSearchCriteriasInvoker 添加快速查询
func (c *LtsClient) CreateSearchCriteriasInvoker(request *model.CreateSearchCriteriasRequest) *CreateSearchCriteriasInvoker {
	requestDef := GenReqDefForCreateSearchCriterias()
	return &CreateSearchCriteriasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStructConfig 通过结构化模板创建结构化配置（新）
//
// 该接口通过结构化模板创建结构化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateStructConfig(request *model.CreateStructConfigRequest) (*model.CreateStructConfigResponse, error) {
	requestDef := GenReqDefForCreateStructConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStructConfigResponse), nil
	}
}

// CreateStructConfigInvoker 通过结构化模板创建结构化配置（新）
func (c *LtsClient) CreateStructConfigInvoker(request *model.CreateStructConfigRequest) *CreateStructConfigInvoker {
	requestDef := GenReqDefForCreateStructConfig()
	return &CreateStructConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStructTemplate 创建结构化配置
//
// 该接口用于创建指定日志流下的结构化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateStructTemplate(request *model.CreateStructTemplateRequest) (*model.CreateStructTemplateResponse, error) {
	requestDef := GenReqDefForCreateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStructTemplateResponse), nil
	}
}

// CreateStructTemplateInvoker 创建结构化配置
func (c *LtsClient) CreateStructTemplateInvoker(request *model.CreateStructTemplateRequest) *CreateStructTemplateInvoker {
	requestDef := GenReqDefForCreateStructTemplate()
	return &CreateStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTags
//
// 添加标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateTags(request *model.CreateTagsRequest) (*model.CreateTagsResponse, error) {
	requestDef := GenReqDefForCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagsResponse), nil
	}
}

// CreateTagsInvoker
func (c *LtsClient) CreateTagsInvoker(request *model.CreateTagsRequest) *CreateTagsInvoker {
	requestDef := GenReqDefForCreateTags()
	return &CreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTransfer 创建日志转储（新版）
//
// 该接口用于创建OBS转储，DIS转储，DMS转储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateTransfer(request *model.CreateTransferRequest) (*model.CreateTransferResponse, error) {
	requestDef := GenReqDefForCreateTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransferResponse), nil
	}
}

// CreateTransferInvoker 创建日志转储（新版）
func (c *LtsClient) CreateTransferInvoker(request *model.CreateTransferRequest) *CreateTransferInvoker {
	requestDef := GenReqDefForCreateTransfer()
	return &CreateTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Createfavorite 创建日志收藏
//
// 创建日志收藏
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) Createfavorite(request *model.CreatefavoriteRequest) (*model.CreatefavoriteResponse, error) {
	requestDef := GenReqDefForCreatefavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatefavoriteResponse), nil
	}
}

// CreatefavoriteInvoker 创建日志收藏
func (c *LtsClient) CreatefavoriteInvoker(request *model.CreatefavoriteRequest) *CreatefavoriteInvoker {
	requestDef := GenReqDefForCreatefavorite()
	return &CreatefavoriteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccessConfig 删除日志接入
//
// 删除日志接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteAccessConfig(request *model.DeleteAccessConfigRequest) (*model.DeleteAccessConfigResponse, error) {
	requestDef := GenReqDefForDeleteAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessConfigResponse), nil
	}
}

// DeleteAccessConfigInvoker 删除日志接入
func (c *LtsClient) DeleteAccessConfigInvoker(request *model.DeleteAccessConfigRequest) *DeleteAccessConfigInvoker {
	requestDef := GenReqDefForDeleteAccessConfig()
	return &DeleteAccessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteActiveAlarms 删除活动告警
//
// 该接口用于删除活动告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteActiveAlarms(request *model.DeleteActiveAlarmsRequest) (*model.DeleteActiveAlarmsResponse, error) {
	requestDef := GenReqDefForDeleteActiveAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteActiveAlarmsResponse), nil
	}
}

// DeleteActiveAlarmsInvoker 删除活动告警
func (c *LtsClient) DeleteActiveAlarmsInvoker(request *model.DeleteActiveAlarmsRequest) *DeleteActiveAlarmsInvoker {
	requestDef := GenReqDefForDeleteActiveAlarms()
	return &DeleteActiveAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDashboard 删除仪表盘
//
// 删除仪表盘
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteDashboard(request *model.DeleteDashboardRequest) (*model.DeleteDashboardResponse, error) {
	requestDef := GenReqDefForDeleteDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDashboardResponse), nil
	}
}

// DeleteDashboardInvoker 删除仪表盘
func (c *LtsClient) DeleteDashboardInvoker(request *model.DeleteDashboardRequest) *DeleteDashboardInvoker {
	requestDef := GenReqDefForDeleteDashboard()
	return &DeleteDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHostGroup 删除主机组
//
// 删除主机组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteHostGroup(request *model.DeleteHostGroupRequest) (*model.DeleteHostGroupResponse, error) {
	requestDef := GenReqDefForDeleteHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostGroupResponse), nil
	}
}

// DeleteHostGroupInvoker 删除主机组
func (c *LtsClient) DeleteHostGroupInvoker(request *model.DeleteHostGroupRequest) *DeleteHostGroupInvoker {
	requestDef := GenReqDefForDeleteHostGroup()
	return &DeleteHostGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKeywordsAlarmRule 删除关键词告警规则
//
// 该接口用于删除关键词告警。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteKeywordsAlarmRule(request *model.DeleteKeywordsAlarmRuleRequest) (*model.DeleteKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForDeleteKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeywordsAlarmRuleResponse), nil
	}
}

// DeleteKeywordsAlarmRuleInvoker 删除关键词告警规则
func (c *LtsClient) DeleteKeywordsAlarmRuleInvoker(request *model.DeleteKeywordsAlarmRuleRequest) *DeleteKeywordsAlarmRuleInvoker {
	requestDef := GenReqDefForDeleteKeywordsAlarmRule()
	return &DeleteKeywordsAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogGroup 删除日志组
//
// 该接口用于删除指定日志组。当日志组中的日志流配置了日志转储，需要取消日志转储后才可删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteLogGroup(request *model.DeleteLogGroupRequest) (*model.DeleteLogGroupResponse, error) {
	requestDef := GenReqDefForDeleteLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogGroupResponse), nil
	}
}

// DeleteLogGroupInvoker 删除日志组
func (c *LtsClient) DeleteLogGroupInvoker(request *model.DeleteLogGroupRequest) *DeleteLogGroupInvoker {
	requestDef := GenReqDefForDeleteLogGroup()
	return &DeleteLogGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogStream 删除日志流
//
// 该接口用于删除指定日志组下的指定日志流。当该日志流配置了日志转储，需要取消日志转储后才可删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteLogStream(request *model.DeleteLogStreamRequest) (*model.DeleteLogStreamResponse, error) {
	requestDef := GenReqDefForDeleteLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogStreamResponse), nil
	}
}

// DeleteLogStreamInvoker 删除日志流
func (c *LtsClient) DeleteLogStreamInvoker(request *model.DeleteLogStreamRequest) *DeleteLogStreamInvoker {
	requestDef := GenReqDefForDeleteLogStream()
	return &DeleteLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotificationTemplate 删除消息模板
//
// 该接口用于删除通知模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteNotificationTemplate(request *model.DeleteNotificationTemplateRequest) (*model.DeleteNotificationTemplateResponse, error) {
	requestDef := GenReqDefForDeleteNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotificationTemplateResponse), nil
	}
}

// DeleteNotificationTemplateInvoker 删除消息模板
func (c *LtsClient) DeleteNotificationTemplateInvoker(request *model.DeleteNotificationTemplateRequest) *DeleteNotificationTemplateInvoker {
	requestDef := GenReqDefForDeleteNotificationTemplate()
	return &DeleteNotificationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSearchCriterias 删除快速查询
//
// 删除快速查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteSearchCriterias(request *model.DeleteSearchCriteriasRequest) (*model.DeleteSearchCriteriasResponse, error) {
	requestDef := GenReqDefForDeleteSearchCriterias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSearchCriteriasResponse), nil
	}
}

// DeleteSearchCriteriasInvoker 删除快速查询
func (c *LtsClient) DeleteSearchCriteriasInvoker(request *model.DeleteSearchCriteriasRequest) *DeleteSearchCriteriasInvoker {
	requestDef := GenReqDefForDeleteSearchCriterias()
	return &DeleteSearchCriteriasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStructTemplate 删除结构化配置
//
// 该接口用于删除指定日志流下的结构化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteStructTemplate(request *model.DeleteStructTemplateRequest) (*model.DeleteStructTemplateResponse, error) {
	requestDef := GenReqDefForDeleteStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStructTemplateResponse), nil
	}
}

// DeleteStructTemplateInvoker 删除结构化配置
func (c *LtsClient) DeleteStructTemplateInvoker(request *model.DeleteStructTemplateRequest) *DeleteStructTemplateInvoker {
	requestDef := GenReqDefForDeleteStructTemplate()
	return &DeleteStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransfer 删除日志转储
//
// 该接口用于删除OBS转储，DIS转储，DMS转储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteTransfer(request *model.DeleteTransferRequest) (*model.DeleteTransferResponse, error) {
	requestDef := GenReqDefForDeleteTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransferResponse), nil
	}
}

// DeleteTransferInvoker 删除日志转储
func (c *LtsClient) DeleteTransferInvoker(request *model.DeleteTransferRequest) *DeleteTransferInvoker {
	requestDef := GenReqDefForDeleteTransfer()
	return &DeleteTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deletefavorite 取消收藏
//
// 取消收藏
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) Deletefavorite(request *model.DeletefavoriteRequest) (*model.DeletefavoriteResponse, error) {
	requestDef := GenReqDefForDeletefavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletefavoriteResponse), nil
	}
}

// DeletefavoriteInvoker 取消收藏
func (c *LtsClient) DeletefavoriteInvoker(request *model.DeletefavoriteRequest) *DeletefavoriteInvoker {
	requestDef := GenReqDefForDeletefavorite()
	return &DeletefavoriteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableLogCollection 关闭超额采集开关
//
// 该接口用于将超额采集日志功能关闭。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DisableLogCollection(request *model.DisableLogCollectionRequest) (*model.DisableLogCollectionResponse, error) {
	requestDef := GenReqDefForDisableLogCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableLogCollectionResponse), nil
	}
}

// DisableLogCollectionInvoker 关闭超额采集开关
func (c *LtsClient) DisableLogCollectionInvoker(request *model.DisableLogCollectionRequest) *DisableLogCollectionInvoker {
	requestDef := GenReqDefForDisableLogCollection()
	return &DisableLogCollectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableLogCollection 打开超额采集开关
//
// 该接口用于将超额采集日志功能打开。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) EnableLogCollection(request *model.EnableLogCollectionRequest) (*model.EnableLogCollectionResponse, error) {
	requestDef := GenReqDefForEnableLogCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLogCollectionResponse), nil
	}
}

// EnableLogCollectionInvoker 打开超额采集开关
func (c *LtsClient) EnableLogCollectionInvoker(request *model.EnableLogCollectionRequest) *EnableLogCollectionInvoker {
	requestDef := GenReqDefForEnableLogCollection()
	return &EnableLogCollectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessConfig 查询日志接入
//
// 查询日志接入列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListAccessConfig(request *model.ListAccessConfigRequest) (*model.ListAccessConfigResponse, error) {
	requestDef := GenReqDefForListAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessConfigResponse), nil
	}
}

// ListAccessConfigInvoker 查询日志接入
func (c *LtsClient) ListAccessConfigInvoker(request *model.ListAccessConfigRequest) *ListAccessConfigInvoker {
	requestDef := GenReqDefForListAccessConfig()
	return &ListAccessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListActiveOrHistoryAlarms 查询活动或历史告警列表
//
// 该接口用于查询告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListActiveOrHistoryAlarms(request *model.ListActiveOrHistoryAlarmsRequest) (*model.ListActiveOrHistoryAlarmsResponse, error) {
	requestDef := GenReqDefForListActiveOrHistoryAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActiveOrHistoryAlarmsResponse), nil
	}
}

// ListActiveOrHistoryAlarmsInvoker 查询活动或历史告警列表
func (c *LtsClient) ListActiveOrHistoryAlarmsInvoker(request *model.ListActiveOrHistoryAlarmsRequest) *ListActiveOrHistoryAlarmsInvoker {
	requestDef := GenReqDefForListActiveOrHistoryAlarms()
	return &ListActiveOrHistoryAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBreifStructTemplate 查询结构化模板简略列表
//
// 该接口用于查询结构化模板简略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListBreifStructTemplate(request *model.ListBreifStructTemplateRequest) (*model.ListBreifStructTemplateResponse, error) {
	requestDef := GenReqDefForListBreifStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBreifStructTemplateResponse), nil
	}
}

// ListBreifStructTemplateInvoker 查询结构化模板简略列表
func (c *LtsClient) ListBreifStructTemplateInvoker(request *model.ListBreifStructTemplateRequest) *ListBreifStructTemplateInvoker {
	requestDef := GenReqDefForListBreifStructTemplate()
	return &ListBreifStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCharts 查询日志流图表
//
// 该接口用于查询日志流图表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListCharts(request *model.ListChartsRequest) (*model.ListChartsResponse, error) {
	requestDef := GenReqDefForListCharts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChartsResponse), nil
	}
}

// ListChartsInvoker 查询日志流图表
func (c *LtsClient) ListChartsInvoker(request *model.ListChartsRequest) *ListChartsInvoker {
	requestDef := GenReqDefForListCharts()
	return &ListChartsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCriterias 获取快速查询
//
// 获取快速查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListCriterias(request *model.ListCriteriasRequest) (*model.ListCriteriasResponse, error) {
	requestDef := GenReqDefForListCriterias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCriteriasResponse), nil
	}
}

// ListCriteriasInvoker 获取快速查询
func (c *LtsClient) ListCriteriasInvoker(request *model.ListCriteriasRequest) *ListCriteriasInvoker {
	requestDef := GenReqDefForListCriterias()
	return &ListCriteriasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistorySql 查询用户历史sql
//
// 查询用户历史sql
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListHistorySql(request *model.ListHistorySqlRequest) (*model.ListHistorySqlResponse, error) {
	requestDef := GenReqDefForListHistorySql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistorySqlResponse), nil
	}
}

// ListHistorySqlInvoker 查询用户历史sql
func (c *LtsClient) ListHistorySqlInvoker(request *model.ListHistorySqlRequest) *ListHistorySqlInvoker {
	requestDef := GenReqDefForListHistorySql()
	return &ListHistorySqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHost 查询主机信息
//
// 查询主机列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListHost(request *model.ListHostRequest) (*model.ListHostResponse, error) {
	requestDef := GenReqDefForListHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostResponse), nil
	}
}

// ListHostInvoker 查询主机信息
func (c *LtsClient) ListHostInvoker(request *model.ListHostRequest) *ListHostInvoker {
	requestDef := GenReqDefForListHost()
	return &ListHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostGroup 查询主机组
//
// 查询主机组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListHostGroup(request *model.ListHostGroupRequest) (*model.ListHostGroupResponse, error) {
	requestDef := GenReqDefForListHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupResponse), nil
	}
}

// ListHostGroupInvoker 查询主机组
func (c *LtsClient) ListHostGroupInvoker(request *model.ListHostGroupRequest) *ListHostGroupInvoker {
	requestDef := GenReqDefForListHostGroup()
	return &ListHostGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeywordsAlarmRules 查询关键词告警规则
//
// 该接口用于查询关键词告警。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListKeywordsAlarmRules(request *model.ListKeywordsAlarmRulesRequest) (*model.ListKeywordsAlarmRulesResponse, error) {
	requestDef := GenReqDefForListKeywordsAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeywordsAlarmRulesResponse), nil
	}
}

// ListKeywordsAlarmRulesInvoker 查询关键词告警规则
func (c *LtsClient) ListKeywordsAlarmRulesInvoker(request *model.ListKeywordsAlarmRulesRequest) *ListKeywordsAlarmRulesInvoker {
	requestDef := GenReqDefForListKeywordsAlarmRules()
	return &ListKeywordsAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogGroups 查询账号下所有日志组
//
// 该接口用于查询账号下所有日志组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListLogGroups(request *model.ListLogGroupsRequest) (*model.ListLogGroupsResponse, error) {
	requestDef := GenReqDefForListLogGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogGroupsResponse), nil
	}
}

// ListLogGroupsInvoker 查询账号下所有日志组
func (c *LtsClient) ListLogGroupsInvoker(request *model.ListLogGroupsRequest) *ListLogGroupsInvoker {
	requestDef := GenReqDefForListLogGroups()
	return &ListLogGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogHistogram 查询日志直方图
//
// 查询关键词搜索条数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListLogHistogram(request *model.ListLogHistogramRequest) (*model.ListLogHistogramResponse, error) {
	requestDef := GenReqDefForListLogHistogram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogHistogramResponse), nil
	}
}

// ListLogHistogramInvoker 查询日志直方图
func (c *LtsClient) ListLogHistogramInvoker(request *model.ListLogHistogramRequest) *ListLogHistogramInvoker {
	requestDef := GenReqDefForListLogHistogram()
	return &ListLogHistogramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogStream 查询指定日志组下的所有日志流
//
// 该接口用于查询指定日志组下的所有日志流信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListLogStream(request *model.ListLogStreamRequest) (*model.ListLogStreamResponse, error) {
	requestDef := GenReqDefForListLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamResponse), nil
	}
}

// ListLogStreamInvoker 查询指定日志组下的所有日志流
func (c *LtsClient) ListLogStreamInvoker(request *model.ListLogStreamRequest) *ListLogStreamInvoker {
	requestDef := GenReqDefForListLogStream()
	return &ListLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogStreams 查询日志流信息
//
// 该接口用于查询LTS日志流信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListLogStreams(request *model.ListLogStreamsRequest) (*model.ListLogStreamsResponse, error) {
	requestDef := GenReqDefForListLogStreams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamsResponse), nil
	}
}

// ListLogStreamsInvoker 查询日志流信息
func (c *LtsClient) ListLogStreamsInvoker(request *model.ListLogStreamsRequest) *ListLogStreamsInvoker {
	requestDef := GenReqDefForListLogStreams()
	return &ListLogStreamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogs 查询日志
//
// 该接口用于查询指定日志流下的日志内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListLogs(request *model.ListLogsRequest) (*model.ListLogsResponse, error) {
	requestDef := GenReqDefForListLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsResponse), nil
	}
}

// ListLogsInvoker 查询日志
func (c *LtsClient) ListLogsInvoker(request *model.ListLogsRequest) *ListLogsInvoker {
	requestDef := GenReqDefForListLogs()
	return &ListLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotificationTemplate 预览消息模板邮件格式
//
// 该接口用于预览通知模板邮件格式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListNotificationTemplate(request *model.ListNotificationTemplateRequest) (*model.ListNotificationTemplateResponse, error) {
	requestDef := GenReqDefForListNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTemplateResponse), nil
	}
}

// ListNotificationTemplateInvoker 预览消息模板邮件格式
func (c *LtsClient) ListNotificationTemplateInvoker(request *model.ListNotificationTemplateRequest) *ListNotificationTemplateInvoker {
	requestDef := GenReqDefForListNotificationTemplate()
	return &ListNotificationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotificationTemplates 查询消息模板
//
// 该接口用于查询通知模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListNotificationTemplates(request *model.ListNotificationTemplatesRequest) (*model.ListNotificationTemplatesResponse, error) {
	requestDef := GenReqDefForListNotificationTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTemplatesResponse), nil
	}
}

// ListNotificationTemplatesInvoker 查询消息模板
func (c *LtsClient) ListNotificationTemplatesInvoker(request *model.ListNotificationTemplatesRequest) *ListNotificationTemplatesInvoker {
	requestDef := GenReqDefForListNotificationTemplates()
	return &ListNotificationTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotificationTopics 查询SMN主题
//
// 该接口用于查询SMN主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListNotificationTopics(request *model.ListNotificationTopicsRequest) (*model.ListNotificationTopicsResponse, error) {
	requestDef := GenReqDefForListNotificationTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTopicsResponse), nil
	}
}

// ListNotificationTopicsInvoker 查询SMN主题
func (c *LtsClient) ListNotificationTopicsInvoker(request *model.ListNotificationTopicsRequest) *ListNotificationTopicsInvoker {
	requestDef := GenReqDefForListNotificationTopics()
	return &ListNotificationTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryAllSearchCriterias 查询日志组下所有快速查询
//
// 查询日志组下所有快速查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListQueryAllSearchCriterias(request *model.ListQueryAllSearchCriteriasRequest) (*model.ListQueryAllSearchCriteriasResponse, error) {
	requestDef := GenReqDefForListQueryAllSearchCriterias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryAllSearchCriteriasResponse), nil
	}
}

// ListQueryAllSearchCriteriasInvoker 查询日志组下所有快速查询
func (c *LtsClient) ListQueryAllSearchCriteriasInvoker(request *model.ListQueryAllSearchCriteriasRequest) *ListQueryAllSearchCriteriasInvoker {
	requestDef := GenReqDefForListQueryAllSearchCriterias()
	return &ListQueryAllSearchCriteriasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryStructuredLogs 查询结构化日志
//
// 该接口用于查询指定日志流下的结构化日志内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListQueryStructuredLogs(request *model.ListQueryStructuredLogsRequest) (*model.ListQueryStructuredLogsResponse, error) {
	requestDef := GenReqDefForListQueryStructuredLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryStructuredLogsResponse), nil
	}
}

// ListQueryStructuredLogsInvoker 查询结构化日志
func (c *LtsClient) ListQueryStructuredLogsInvoker(request *model.ListQueryStructuredLogsRequest) *ListQueryStructuredLogsInvoker {
	requestDef := GenReqDefForListQueryStructuredLogs()
	return &ListQueryStructuredLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStructTemplate 查询结构化模板
//
// 该接口用于查询结构化模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListStructTemplate(request *model.ListStructTemplateRequest) (*model.ListStructTemplateResponse, error) {
	requestDef := GenReqDefForListStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStructTemplateResponse), nil
	}
}

// ListStructTemplateInvoker 查询结构化模板
func (c *LtsClient) ListStructTemplateInvoker(request *model.ListStructTemplateRequest) *ListStructTemplateInvoker {
	requestDef := GenReqDefForListStructTemplate()
	return &ListStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStructuredLogsWithTimeRange 查询结构化日志（新版）
//
// 该接口用于查询指定日志流下的结构化日志内容（新版）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListStructuredLogsWithTimeRange(request *model.ListStructuredLogsWithTimeRangeRequest) (*model.ListStructuredLogsWithTimeRangeResponse, error) {
	requestDef := GenReqDefForListStructuredLogsWithTimeRange()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStructuredLogsWithTimeRangeResponse), nil
	}
}

// ListStructuredLogsWithTimeRangeInvoker 查询结构化日志（新版）
func (c *LtsClient) ListStructuredLogsWithTimeRangeInvoker(request *model.ListStructuredLogsWithTimeRangeRequest) *ListStructuredLogsWithTimeRangeInvoker {
	requestDef := GenReqDefForListStructuredLogsWithTimeRange()
	return &ListStructuredLogsWithTimeRangeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTimeLineTrafficStatistics 按时间段统计查询资源
//
// 按时间段统计查询资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListTimeLineTrafficStatistics(request *model.ListTimeLineTrafficStatisticsRequest) (*model.ListTimeLineTrafficStatisticsResponse, error) {
	requestDef := GenReqDefForListTimeLineTrafficStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTimeLineTrafficStatisticsResponse), nil
	}
}

// ListTimeLineTrafficStatisticsInvoker 按时间段统计查询资源
func (c *LtsClient) ListTimeLineTrafficStatisticsInvoker(request *model.ListTimeLineTrafficStatisticsRequest) *ListTimeLineTrafficStatisticsInvoker {
	requestDef := GenReqDefForListTimeLineTrafficStatistics()
	return &ListTimeLineTrafficStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopnTrafficStatistics 统计top n的日志组或日志流流量
//
// 统计top n的日志组或日志流流量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListTopnTrafficStatistics(request *model.ListTopnTrafficStatisticsRequest) (*model.ListTopnTrafficStatisticsResponse, error) {
	requestDef := GenReqDefForListTopnTrafficStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopnTrafficStatisticsResponse), nil
	}
}

// ListTopnTrafficStatisticsInvoker 统计top n的日志组或日志流流量
func (c *LtsClient) ListTopnTrafficStatisticsInvoker(request *model.ListTopnTrafficStatisticsRequest) *ListTopnTrafficStatisticsInvoker {
	requestDef := GenReqDefForListTopnTrafficStatistics()
	return &ListTopnTrafficStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransfers 查询日志转储
//
// 该接口用于查询OBS转储，DIS转储，DMS转储配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListTransfers(request *model.ListTransfersRequest) (*model.ListTransfersResponse, error) {
	requestDef := GenReqDefForListTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransfersResponse), nil
	}
}

// ListTransfersInvoker 查询日志转储
func (c *LtsClient) ListTransfersInvoker(request *model.ListTransfersRequest) *ListTransfersInvoker {
	requestDef := GenReqDefForListTransfers()
	return &ListTransfersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterDmsKafkaInstance 注册DMS kafka实例
//
// 该接口用于注册DMS kafka实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) RegisterDmsKafkaInstance(request *model.RegisterDmsKafkaInstanceRequest) (*model.RegisterDmsKafkaInstanceResponse, error) {
	requestDef := GenReqDefForRegisterDmsKafkaInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDmsKafkaInstanceResponse), nil
	}
}

// RegisterDmsKafkaInstanceInvoker 注册DMS kafka实例
func (c *LtsClient) RegisterDmsKafkaInstanceInvoker(request *model.RegisterDmsKafkaInstanceRequest) *RegisterDmsKafkaInstanceInvoker {
	requestDef := GenReqDefForRegisterDmsKafkaInstance()
	return &RegisterDmsKafkaInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAdminConfig 获取日志汇聚开关
//
// 只能由管理员或者委托管理员调用    获取日志汇聚开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowAdminConfig(request *model.ShowAdminConfigRequest) (*model.ShowAdminConfigResponse, error) {
	requestDef := GenReqDefForShowAdminConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdminConfigResponse), nil
	}
}

// ShowAdminConfigInvoker 获取日志汇聚开关
func (c *LtsClient) ShowAdminConfigInvoker(request *model.ShowAdminConfigRequest) *ShowAdminConfigInvoker {
	requestDef := GenReqDefForShowAdminConfig()
	return &ShowAdminConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogConvergeConfig 获取组织成员汇聚配置
//
// 只能由组织管理员或者委托管理员调用    获取组织成员汇聚配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowLogConvergeConfig(request *model.ShowLogConvergeConfigRequest) (*model.ShowLogConvergeConfigResponse, error) {
	requestDef := GenReqDefForShowLogConvergeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogConvergeConfigResponse), nil
	}
}

// ShowLogConvergeConfigInvoker 获取组织成员汇聚配置
func (c *LtsClient) ShowLogConvergeConfigInvoker(request *model.ShowLogConvergeConfigRequest) *ShowLogConvergeConfigInvoker {
	requestDef := GenReqDefForShowLogConvergeConfig()
	return &ShowLogConvergeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMemberGroupAndStream 获取组织成员日志组日志流
//
// 只能由管理员或者委托管理员调用，获取组织成员日志组日志流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowMemberGroupAndStream(request *model.ShowMemberGroupAndStreamRequest) (*model.ShowMemberGroupAndStreamResponse, error) {
	requestDef := GenReqDefForShowMemberGroupAndStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberGroupAndStreamResponse), nil
	}
}

// ShowMemberGroupAndStreamInvoker 获取组织成员日志组日志流
func (c *LtsClient) ShowMemberGroupAndStreamInvoker(request *model.ShowMemberGroupAndStreamRequest) *ShowMemberGroupAndStreamInvoker {
	requestDef := GenReqDefForShowMemberGroupAndStream()
	return &ShowMemberGroupAndStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotificationTemplate 查询单个消息模板
//
// 该接口用于查询单个通知模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowNotificationTemplate(request *model.ShowNotificationTemplateRequest) (*model.ShowNotificationTemplateResponse, error) {
	requestDef := GenReqDefForShowNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationTemplateResponse), nil
	}
}

// ShowNotificationTemplateInvoker 查询单个消息模板
func (c *LtsClient) ShowNotificationTemplateInvoker(request *model.ShowNotificationTemplateRequest) *ShowNotificationTemplateInvoker {
	requestDef := GenReqDefForShowNotificationTemplate()
	return &ShowNotificationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStructTemplate 查询结构化配置
//
// 该接口用于查询指定日志流下的结构化配置内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowStructTemplate(request *model.ShowStructTemplateRequest) (*model.ShowStructTemplateResponse, error) {
	requestDef := GenReqDefForShowStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStructTemplateResponse), nil
	}
}

// ShowStructTemplateInvoker 查询结构化配置
func (c *LtsClient) ShowStructTemplateInvoker(request *model.ShowStructTemplateRequest) *ShowStructTemplateInvoker {
	requestDef := GenReqDefForShowStructTemplate()
	return &ShowStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessConfig 修改日志接入
//
// 修改日志接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateAccessConfig(request *model.UpdateAccessConfigRequest) (*model.UpdateAccessConfigResponse, error) {
	requestDef := GenReqDefForUpdateAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessConfigResponse), nil
	}
}

// UpdateAccessConfigInvoker 修改日志接入
func (c *LtsClient) UpdateAccessConfigInvoker(request *model.UpdateAccessConfigRequest) *UpdateAccessConfigInvoker {
	requestDef := GenReqDefForUpdateAccessConfig()
	return &UpdateAccessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostGroup 修改主机组
//
// 修改主机组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateHostGroup(request *model.UpdateHostGroupRequest) (*model.UpdateHostGroupResponse, error) {
	requestDef := GenReqDefForUpdateHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostGroupResponse), nil
	}
}

// UpdateHostGroupInvoker 修改主机组
func (c *LtsClient) UpdateHostGroupInvoker(request *model.UpdateHostGroupRequest) *UpdateHostGroupInvoker {
	requestDef := GenReqDefForUpdateHostGroup()
	return &UpdateHostGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeywordsAlarmRule 修改关键词告警规则
//
// 该接口用于修改关键词告警。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateKeywordsAlarmRule(request *model.UpdateKeywordsAlarmRuleRequest) (*model.UpdateKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForUpdateKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeywordsAlarmRuleResponse), nil
	}
}

// UpdateKeywordsAlarmRuleInvoker 修改关键词告警规则
func (c *LtsClient) UpdateKeywordsAlarmRuleInvoker(request *model.UpdateKeywordsAlarmRuleRequest) *UpdateKeywordsAlarmRuleInvoker {
	requestDef := GenReqDefForUpdateKeywordsAlarmRule()
	return &UpdateKeywordsAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogConvergeConfig 更新汇聚配置
//
// 只能由管理员或者委托管理员 ,更新汇聚配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateLogConvergeConfig(request *model.UpdateLogConvergeConfigRequest) (*model.UpdateLogConvergeConfigResponse, error) {
	requestDef := GenReqDefForUpdateLogConvergeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogConvergeConfigResponse), nil
	}
}

// UpdateLogConvergeConfigInvoker 更新汇聚配置
func (c *LtsClient) UpdateLogConvergeConfigInvoker(request *model.UpdateLogConvergeConfigRequest) *UpdateLogConvergeConfigInvoker {
	requestDef := GenReqDefForUpdateLogConvergeConfig()
	return &UpdateLogConvergeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogGroup 修改日志组
//
// 该接口用于修改指定日志组下的日志存储时长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateLogGroup(request *model.UpdateLogGroupRequest) (*model.UpdateLogGroupResponse, error) {
	requestDef := GenReqDefForUpdateLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogGroupResponse), nil
	}
}

// UpdateLogGroupInvoker 修改日志组
func (c *LtsClient) UpdateLogGroupInvoker(request *model.UpdateLogGroupRequest) *UpdateLogGroupInvoker {
	requestDef := GenReqDefForUpdateLogGroup()
	return &UpdateLogGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogStream 修改日志流
//
// 该接口用于修改指定日志流下的日志存储时长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateLogStream(request *model.UpdateLogStreamRequest) (*model.UpdateLogStreamResponse, error) {
	requestDef := GenReqDefForUpdateLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogStreamResponse), nil
	}
}

// UpdateLogStreamInvoker 修改日志流
func (c *LtsClient) UpdateLogStreamInvoker(request *model.UpdateLogStreamRequest) *UpdateLogStreamInvoker {
	requestDef := GenReqDefForUpdateLogStream()
	return &UpdateLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotificationTemplate 修改消息模板
//
// 该接口用于修改通知模板,根据名称进行修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateNotificationTemplate(request *model.UpdateNotificationTemplateRequest) (*model.UpdateNotificationTemplateResponse, error) {
	requestDef := GenReqDefForUpdateNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationTemplateResponse), nil
	}
}

// UpdateNotificationTemplateInvoker 修改消息模板
func (c *LtsClient) UpdateNotificationTemplateInvoker(request *model.UpdateNotificationTemplateRequest) *UpdateNotificationTemplateInvoker {
	requestDef := GenReqDefForUpdateNotificationTemplate()
	return &UpdateNotificationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStructConfig 通过结构化模板修改结构化配置（新）
//
// 该接口通过结构化模板修改结构化配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateStructConfig(request *model.UpdateStructConfigRequest) (*model.UpdateStructConfigResponse, error) {
	requestDef := GenReqDefForUpdateStructConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStructConfigResponse), nil
	}
}

// UpdateStructConfigInvoker 通过结构化模板修改结构化配置（新）
func (c *LtsClient) UpdateStructConfigInvoker(request *model.UpdateStructConfigRequest) *UpdateStructConfigInvoker {
	requestDef := GenReqDefForUpdateStructConfig()
	return &UpdateStructConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStructTemplate 修改结构化配置
//
// 该接口用于修改指定日志流下的结构化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateStructTemplate(request *model.UpdateStructTemplateRequest) (*model.UpdateStructTemplateResponse, error) {
	requestDef := GenReqDefForUpdateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStructTemplateResponse), nil
	}
}

// UpdateStructTemplateInvoker 修改结构化配置
func (c *LtsClient) UpdateStructTemplateInvoker(request *model.UpdateStructTemplateRequest) *UpdateStructTemplateInvoker {
	requestDef := GenReqDefForUpdateStructTemplate()
	return &UpdateStructTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSwitch 修改日志汇聚开关
//
// 只能由管理员或者委托管理员调用     修改日志汇聚开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateSwitch(request *model.UpdateSwitchRequest) (*model.UpdateSwitchResponse, error) {
	requestDef := GenReqDefForUpdateSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSwitchResponse), nil
	}
}

// UpdateSwitchInvoker 修改日志汇聚开关
func (c *LtsClient) UpdateSwitchInvoker(request *model.UpdateSwitchRequest) *UpdateSwitchInvoker {
	requestDef := GenReqDefForUpdateSwitch()
	return &UpdateSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTransfer 更新日志转储
//
// 该接口用于更新OBS转储，DIS转储，DMS转储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateTransfer(request *model.UpdateTransferRequest) (*model.UpdateTransferResponse, error) {
	requestDef := GenReqDefForUpdateTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTransferResponse), nil
	}
}

// UpdateTransferInvoker 更新日志转储
func (c *LtsClient) UpdateTransferInvoker(request *model.UpdateTransferRequest) *UpdateTransferInvoker {
	requestDef := GenReqDefForUpdateTransfer()
	return &UpdateTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAomMappingRules 创建接入规则
//
// 该接口用于创建aom日志接入lts规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateAomMappingRules(request *model.CreateAomMappingRulesRequest) (*model.CreateAomMappingRulesResponse, error) {
	requestDef := GenReqDefForCreateAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAomMappingRulesResponse), nil
	}
}

// CreateAomMappingRulesInvoker 创建接入规则
func (c *LtsClient) CreateAomMappingRulesInvoker(request *model.CreateAomMappingRulesRequest) *CreateAomMappingRulesInvoker {
	requestDef := GenReqDefForCreateAomMappingRules()
	return &CreateAomMappingRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAomMappingRules 删除接入规则
//
// 该接口用于删除lts接入规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteAomMappingRules(request *model.DeleteAomMappingRulesRequest) (*model.DeleteAomMappingRulesResponse, error) {
	requestDef := GenReqDefForDeleteAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAomMappingRulesResponse), nil
	}
}

// DeleteAomMappingRulesInvoker 删除接入规则
func (c *LtsClient) DeleteAomMappingRulesInvoker(request *model.DeleteAomMappingRulesRequest) *DeleteAomMappingRulesInvoker {
	requestDef := GenReqDefForDeleteAomMappingRules()
	return &DeleteAomMappingRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAomMappingRule 查询单个接入规则
//
// 该接口用于查询单个aom日志接入lts
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowAomMappingRule(request *model.ShowAomMappingRuleRequest) (*model.ShowAomMappingRuleResponse, error) {
	requestDef := GenReqDefForShowAomMappingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAomMappingRuleResponse), nil
	}
}

// ShowAomMappingRuleInvoker 查询单个接入规则
func (c *LtsClient) ShowAomMappingRuleInvoker(request *model.ShowAomMappingRuleRequest) *ShowAomMappingRuleInvoker {
	requestDef := GenReqDefForShowAomMappingRule()
	return &ShowAomMappingRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAomMappingRules 查询所有接入规则
//
// 该接口用于查询aom日志所有接入lts规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowAomMappingRules(request *model.ShowAomMappingRulesRequest) (*model.ShowAomMappingRulesResponse, error) {
	requestDef := GenReqDefForShowAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAomMappingRulesResponse), nil
	}
}

// ShowAomMappingRulesInvoker 查询所有接入规则
func (c *LtsClient) ShowAomMappingRulesInvoker(request *model.ShowAomMappingRulesRequest) *ShowAomMappingRulesInvoker {
	requestDef := GenReqDefForShowAomMappingRules()
	return &ShowAomMappingRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAomMappingRules 修改接入规则
//
// 该接口用于修改接入规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateAomMappingRules(request *model.UpdateAomMappingRulesRequest) (*model.UpdateAomMappingRulesResponse, error) {
	requestDef := GenReqDefForUpdateAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAomMappingRulesResponse), nil
	}
}

// UpdateAomMappingRulesInvoker 修改接入规则
func (c *LtsClient) UpdateAomMappingRulesInvoker(request *model.UpdateAomMappingRulesRequest) *UpdateAomMappingRulesInvoker {
	requestDef := GenReqDefForUpdateAomMappingRules()
	return &UpdateAomMappingRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConsumerGroupHeartBeat 消费者发送心跳到服务端
//
// 消费者发送心跳到服务端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ConsumerGroupHeartBeat(request *model.ConsumerGroupHeartBeatRequest) (*model.ConsumerGroupHeartBeatResponse, error) {
	requestDef := GenReqDefForConsumerGroupHeartBeat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConsumerGroupHeartBeatResponse), nil
	}
}

// ConsumerGroupHeartBeatInvoker 消费者发送心跳到服务端
func (c *LtsClient) ConsumerGroupHeartBeatInvoker(request *model.ConsumerGroupHeartBeatRequest) *ConsumerGroupHeartBeatInvoker {
	requestDef := GenReqDefForConsumerGroupHeartBeat()
	return &ConsumerGroupHeartBeatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConsumerGroup 创建消费组
//
// 创建消费组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateConsumerGroup(request *model.CreateConsumerGroupRequest) (*model.CreateConsumerGroupResponse, error) {
	requestDef := GenReqDefForCreateConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConsumerGroupResponse), nil
	}
}

// CreateConsumerGroupInvoker 创建消费组
func (c *LtsClient) CreateConsumerGroupInvoker(request *model.CreateConsumerGroupRequest) *CreateConsumerGroupInvoker {
	requestDef := GenReqDefForCreateConsumerGroup()
	return &CreateConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConsumerGroup 删除消费组
//
// 删除消费组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteConsumerGroup(request *model.DeleteConsumerGroupRequest) (*model.DeleteConsumerGroupResponse, error) {
	requestDef := GenReqDefForDeleteConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConsumerGroupResponse), nil
	}
}

// DeleteConsumerGroupInvoker 删除消费组
func (c *LtsClient) DeleteConsumerGroupInvoker(request *model.DeleteConsumerGroupRequest) *DeleteConsumerGroupInvoker {
	requestDef := GenReqDefForDeleteConsumerGroup()
	return &DeleteConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsumerGroup 查询消费组列表
//
// 查询消费组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListConsumerGroup(request *model.ListConsumerGroupRequest) (*model.ListConsumerGroupResponse, error) {
	requestDef := GenReqDefForListConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumerGroupResponse), nil
	}
}

// ListConsumerGroupInvoker 查询消费组列表
func (c *LtsClient) ListConsumerGroupInvoker(request *model.ListConsumerGroupRequest) *ListConsumerGroupInvoker {
	requestDef := GenReqDefForListConsumerGroup()
	return &ListConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDetailsConsumerGroup 查询消费组详情
//
// 查询消费组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListDetailsConsumerGroup(request *model.ListDetailsConsumerGroupRequest) (*model.ListDetailsConsumerGroupResponse, error) {
	requestDef := GenReqDefForListDetailsConsumerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDetailsConsumerGroupResponse), nil
	}
}

// ListDetailsConsumerGroupInvoker 查询消费组详情
func (c *LtsClient) ListDetailsConsumerGroupInvoker(request *model.ListDetailsConsumerGroupRequest) *ListDetailsConsumerGroupInvoker {
	requestDef := GenReqDefForListDetailsConsumerGroup()
	return &ListDetailsConsumerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCursorByTime 通过时间获取消费游标
//
// 通过时间查询cursor
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowCursorByTime(request *model.ShowCursorByTimeRequest) (*model.ShowCursorByTimeResponse, error) {
	requestDef := GenReqDefForShowCursorByTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCursorByTimeResponse), nil
	}
}

// ShowCursorByTimeInvoker 通过时间获取消费游标
func (c *LtsClient) ShowCursorByTimeInvoker(request *model.ShowCursorByTimeRequest) *ShowCursorByTimeInvoker {
	requestDef := GenReqDefForShowCursorByTime()
	return &ShowCursorByTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCursorTime 通过消费游标获取时间
//
// 通过cursor查询服务端时间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowCursorTime(request *model.ShowCursorTimeRequest) (*model.ShowCursorTimeResponse, error) {
	requestDef := GenReqDefForShowCursorTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCursorTimeResponse), nil
	}
}

// ShowCursorTimeInvoker 通过消费游标获取时间
func (c *LtsClient) ShowCursorTimeInvoker(request *model.ShowCursorTimeRequest) *ShowCursorTimeInvoker {
	requestDef := GenReqDefForShowCursorTime()
	return &ShowCursorTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogStreamShards 流消费获取Shards
//
// 流消费获取所有的query shards
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ShowLogStreamShards(request *model.ShowLogStreamShardsRequest) (*model.ShowLogStreamShardsResponse, error) {
	requestDef := GenReqDefForShowLogStreamShards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogStreamShardsResponse), nil
	}
}

// ShowLogStreamShardsInvoker 流消费获取Shards
func (c *LtsClient) ShowLogStreamShardsInvoker(request *model.ShowLogStreamShardsRequest) *ShowLogStreamShardsInvoker {
	requestDef := GenReqDefForShowLogStreamShards()
	return &ShowLogStreamShardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCheckPoint 更新消费组位点
//
// 更新消费组位点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateCheckPoint(request *model.UpdateCheckPointRequest) (*model.UpdateCheckPointResponse, error) {
	requestDef := GenReqDefForUpdateCheckPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCheckPointResponse), nil
	}
}

// UpdateCheckPointInvoker 更新消费组位点
func (c *LtsClient) UpdateCheckPointInvoker(request *model.UpdateCheckPointRequest) *UpdateCheckPointInvoker {
	requestDef := GenReqDefForUpdateCheckPoint()
	return &UpdateCheckPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlAlarmRule 创建SQL告警规则
//
// 该接口用于创建SQL告警，目前每个帐户最多可以创建共200个关键词告警与SQL告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) CreateSqlAlarmRule(request *model.CreateSqlAlarmRuleRequest) (*model.CreateSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForCreateSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlAlarmRuleResponse), nil
	}
}

// CreateSqlAlarmRuleInvoker 创建SQL告警规则
func (c *LtsClient) CreateSqlAlarmRuleInvoker(request *model.CreateSqlAlarmRuleRequest) *CreateSqlAlarmRuleInvoker {
	requestDef := GenReqDefForCreateSqlAlarmRule()
	return &CreateSqlAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlAlarmRule 删除SQL告警规则
//
// 该接口用于删除SQL告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) DeleteSqlAlarmRule(request *model.DeleteSqlAlarmRuleRequest) (*model.DeleteSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForDeleteSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlAlarmRuleResponse), nil
	}
}

// DeleteSqlAlarmRuleInvoker 删除SQL告警规则
func (c *LtsClient) DeleteSqlAlarmRuleInvoker(request *model.DeleteSqlAlarmRuleRequest) *DeleteSqlAlarmRuleInvoker {
	requestDef := GenReqDefForDeleteSqlAlarmRule()
	return &DeleteSqlAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlAlarmRules 查询SQL告警规则
//
// 该接口用于查询SQL告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) ListSqlAlarmRules(request *model.ListSqlAlarmRulesRequest) (*model.ListSqlAlarmRulesResponse, error) {
	requestDef := GenReqDefForListSqlAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlAlarmRulesResponse), nil
	}
}

// ListSqlAlarmRulesInvoker 查询SQL告警规则
func (c *LtsClient) ListSqlAlarmRulesInvoker(request *model.ListSqlAlarmRulesRequest) *ListSqlAlarmRulesInvoker {
	requestDef := GenReqDefForListSqlAlarmRules()
	return &ListSqlAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmRuleStatus 切换告警规则状态
//
// 改变告警规则状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateAlarmRuleStatus(request *model.UpdateAlarmRuleStatusRequest) (*model.UpdateAlarmRuleStatusResponse, error) {
	requestDef := GenReqDefForUpdateAlarmRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmRuleStatusResponse), nil
	}
}

// UpdateAlarmRuleStatusInvoker 切换告警规则状态
func (c *LtsClient) UpdateAlarmRuleStatusInvoker(request *model.UpdateAlarmRuleStatusRequest) *UpdateAlarmRuleStatusInvoker {
	requestDef := GenReqDefForUpdateAlarmRuleStatus()
	return &UpdateAlarmRuleStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSqlAlarmRule 修改SQL告警规则
//
// 该接口用于修改SQL告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LtsClient) UpdateSqlAlarmRule(request *model.UpdateSqlAlarmRuleRequest) (*model.UpdateSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForUpdateSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlAlarmRuleResponse), nil
	}
}

// UpdateSqlAlarmRuleInvoker 修改SQL告警规则
func (c *LtsClient) UpdateSqlAlarmRuleInvoker(request *model.UpdateSqlAlarmRuleRequest) *UpdateSqlAlarmRuleInvoker {
	requestDef := GenReqDefForUpdateSqlAlarmRule()
	return &UpdateSqlAlarmRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
