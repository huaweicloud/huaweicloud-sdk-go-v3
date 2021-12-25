package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
)

type LtsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewLtsClient(hcClient *http_client.HcHttpClient) *LtsClient {
	return &LtsClient{HcClient: hcClient}
}

func LtsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建日志接入
func (c *LtsClient) CreateAccessConfig(request *model.CreateAccessConfigRequest) (*model.CreateAccessConfigResponse, error) {
	requestDef := GenReqDefForCreateAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessConfigResponse), nil
	}
}

//创建主机组
func (c *LtsClient) CreateHostGroup(request *model.CreateHostGroupRequest) (*model.CreateHostGroupResponse, error) {
	requestDef := GenReqDefForCreateHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostGroupResponse), nil
	}
}

//该接口用于创建关键词告警，目前每个帐户最多可以创建共200个关键词告警与SQL告警。
func (c *LtsClient) CreateKeywordsAlarmRule(request *model.CreateKeywordsAlarmRuleRequest) (*model.CreateKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForCreateKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeywordsAlarmRuleResponse), nil
	}
}

//该接口用于将指定的一个或多个日志流的日志转储到OBS服务。
func (c *LtsClient) CreateLogDumpObs(request *model.CreateLogDumpObsRequest) (*model.CreateLogDumpObsResponse, error) {
	requestDef := GenReqDefForCreateLogDumpObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogDumpObsResponse), nil
	}
}

//该接口用于创建一个日志组
func (c *LtsClient) CreateLogGroup(request *model.CreateLogGroupRequest) (*model.CreateLogGroupResponse, error) {
	requestDef := GenReqDefForCreateLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogGroupResponse), nil
	}
}

//该接口用于创建某个指定日志组下的日志流
func (c *LtsClient) CreateLogStream(request *model.CreateLogStreamRequest) (*model.CreateLogStreamResponse, error) {
	requestDef := GenReqDefForCreateLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogStreamResponse), nil
	}
}

//该接口用于创建通知模板，目前每个帐户最多可以创建共100个通知模板，创建后名称不可修改。
func (c *LtsClient) CreateNotificationTemplate(request *model.CreateNotificationTemplateRequest) (*model.CreateNotificationTemplateResponse, error) {
	requestDef := GenReqDefForCreateNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotificationTemplateResponse), nil
	}
}

//该接口用于创建指定日志流下的结构化配置。
func (c *LtsClient) CreateStructTemplate(request *model.CreateStructTemplateRequest) (*model.CreateStructTemplateResponse, error) {
	requestDef := GenReqDefForCreateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStructTemplateResponse), nil
	}
}

//该接口用于创建OBS转储，DIS转储，DMS转储。
func (c *LtsClient) CreateTransfer(request *model.CreateTransferRequest) (*model.CreateTransferResponse, error) {
	requestDef := GenReqDefForCreateTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransferResponse), nil
	}
}

//删除日志接入
func (c *LtsClient) DeleteAccessConfig(request *model.DeleteAccessConfigRequest) (*model.DeleteAccessConfigResponse, error) {
	requestDef := GenReqDefForDeleteAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessConfigResponse), nil
	}
}

//该接口用于删除活动告警
func (c *LtsClient) DeleteActiveAlarms(request *model.DeleteActiveAlarmsRequest) (*model.DeleteActiveAlarmsResponse, error) {
	requestDef := GenReqDefForDeleteActiveAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteActiveAlarmsResponse), nil
	}
}

//删除主机组
func (c *LtsClient) DeleteHostGroup(request *model.DeleteHostGroupRequest) (*model.DeleteHostGroupResponse, error) {
	requestDef := GenReqDefForDeleteHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostGroupResponse), nil
	}
}

//该接口用于删除关键词告警。
func (c *LtsClient) DeleteKeywordsAlarmRule(request *model.DeleteKeywordsAlarmRuleRequest) (*model.DeleteKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForDeleteKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeywordsAlarmRuleResponse), nil
	}
}

//该接口用于删除指定日志组。当日志组中的日志流配置了日志转储，需要取消日志转储后才可删除。
func (c *LtsClient) DeleteLogGroup(request *model.DeleteLogGroupRequest) (*model.DeleteLogGroupResponse, error) {
	requestDef := GenReqDefForDeleteLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogGroupResponse), nil
	}
}

//该接口用于删除指定日志组下的指定日志流。当该日志流配置了日志转储，需要取消日志转储后才可删除。
func (c *LtsClient) DeleteLogStream(request *model.DeleteLogStreamRequest) (*model.DeleteLogStreamResponse, error) {
	requestDef := GenReqDefForDeleteLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogStreamResponse), nil
	}
}

//该接口用于删除通知模板。
func (c *LtsClient) DeleteNotificationTemplate(request *model.DeleteNotificationTemplateRequest) (*model.DeleteNotificationTemplateResponse, error) {
	requestDef := GenReqDefForDeleteNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotificationTemplateResponse), nil
	}
}

//该接口用于删除指定日志流下的结构化配置。
func (c *LtsClient) DeleteStructTemplate(request *model.DeleteStructTemplateRequest) (*model.DeleteStructTemplateResponse, error) {
	requestDef := GenReqDefForDeleteStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStructTemplateResponse), nil
	}
}

//该接口用于删除OBS转储，DIS转储，DMS转储。
func (c *LtsClient) DeleteTransfer(request *model.DeleteTransferRequest) (*model.DeleteTransferResponse, error) {
	requestDef := GenReqDefForDeleteTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransferResponse), nil
	}
}

//该接口用于将超额采集日志功能关闭。
func (c *LtsClient) DisableLogCollection(request *model.DisableLogCollectionRequest) (*model.DisableLogCollectionResponse, error) {
	requestDef := GenReqDefForDisableLogCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableLogCollectionResponse), nil
	}
}

//该接口用于将超额采集日志功能打开。
func (c *LtsClient) EnableLogCollection(request *model.EnableLogCollectionRequest) (*model.EnableLogCollectionResponse, error) {
	requestDef := GenReqDefForEnableLogCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLogCollectionResponse), nil
	}
}

//查询日志接入列表
func (c *LtsClient) ListAccessConfig(request *model.ListAccessConfigRequest) (*model.ListAccessConfigResponse, error) {
	requestDef := GenReqDefForListAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessConfigResponse), nil
	}
}

//该接口用于查询告警列表
func (c *LtsClient) ListActiveOrHistoryAlarms(request *model.ListActiveOrHistoryAlarmsRequest) (*model.ListActiveOrHistoryAlarmsResponse, error) {
	requestDef := GenReqDefForListActiveOrHistoryAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActiveOrHistoryAlarmsResponse), nil
	}
}

//该接口用于查询日志流图表
func (c *LtsClient) ListCharts(request *model.ListChartsRequest) (*model.ListChartsResponse, error) {
	requestDef := GenReqDefForListCharts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChartsResponse), nil
	}
}

//查询主机列表
func (c *LtsClient) ListHost(request *model.ListHostRequest) (*model.ListHostResponse, error) {
	requestDef := GenReqDefForListHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostResponse), nil
	}
}

//查询主机组列表
func (c *LtsClient) ListHostGroup(request *model.ListHostGroupRequest) (*model.ListHostGroupResponse, error) {
	requestDef := GenReqDefForListHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupResponse), nil
	}
}

//该接口用于查询关键词告警。
func (c *LtsClient) ListKeywordsAlarmRules(request *model.ListKeywordsAlarmRulesRequest) (*model.ListKeywordsAlarmRulesResponse, error) {
	requestDef := GenReqDefForListKeywordsAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeywordsAlarmRulesResponse), nil
	}
}

//该接口用于查询账号下所有日志组。
func (c *LtsClient) ListLogGroups(request *model.ListLogGroupsRequest) (*model.ListLogGroupsResponse, error) {
	requestDef := GenReqDefForListLogGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogGroupsResponse), nil
	}
}

//查询关键词搜索条数
func (c *LtsClient) ListLogHistogram(request *model.ListLogHistogramRequest) (*model.ListLogHistogramResponse, error) {
	requestDef := GenReqDefForListLogHistogram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogHistogramResponse), nil
	}
}

//该接口用于查询指定日志组下的所有日志流信息。
func (c *LtsClient) ListLogStream(request *model.ListLogStreamRequest) (*model.ListLogStreamResponse, error) {
	requestDef := GenReqDefForListLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamResponse), nil
	}
}

//该接口用于查询LTS日志流信息。
func (c *LtsClient) ListLogStreams(request *model.ListLogStreamsRequest) (*model.ListLogStreamsResponse, error) {
	requestDef := GenReqDefForListLogStreams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamsResponse), nil
	}
}

//该接口用于查询指定日志流下的日志内容。
func (c *LtsClient) ListLogs(request *model.ListLogsRequest) (*model.ListLogsResponse, error) {
	requestDef := GenReqDefForListLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsResponse), nil
	}
}

//该接口用于预览通知模板邮件格式
func (c *LtsClient) ListNotificationTemplate(request *model.ListNotificationTemplateRequest) (*model.ListNotificationTemplateResponse, error) {
	requestDef := GenReqDefForListNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTemplateResponse), nil
	}
}

//该接口用于查询通知模板。
func (c *LtsClient) ListNotificationTemplates(request *model.ListNotificationTemplatesRequest) (*model.ListNotificationTemplatesResponse, error) {
	requestDef := GenReqDefForListNotificationTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTemplatesResponse), nil
	}
}

//该接口用于查询SMN主题
func (c *LtsClient) ListNotificationTopics(request *model.ListNotificationTopicsRequest) (*model.ListNotificationTopicsResponse, error) {
	requestDef := GenReqDefForListNotificationTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationTopicsResponse), nil
	}
}

//该接口用于查询指定日志流下的结构化日志内容。
func (c *LtsClient) ListQueryStructuredLogs(request *model.ListQueryStructuredLogsRequest) (*model.ListQueryStructuredLogsResponse, error) {
	requestDef := GenReqDefForListQueryStructuredLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryStructuredLogsResponse), nil
	}
}

//该接口用于查询指定日志流下的结构化日志内容（新版）。
func (c *LtsClient) ListStructuredLogsWithTimeRange(request *model.ListStructuredLogsWithTimeRangeRequest) (*model.ListStructuredLogsWithTimeRangeResponse, error) {
	requestDef := GenReqDefForListStructuredLogsWithTimeRange()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStructuredLogsWithTimeRangeResponse), nil
	}
}

//该接口用于查询OBS转储，DIS转储，DMS转储配置。
func (c *LtsClient) ListTransfers(request *model.ListTransfersRequest) (*model.ListTransfersResponse, error) {
	requestDef := GenReqDefForListTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransfersResponse), nil
	}
}

//该接口用于注册DMS kafka实例。
func (c *LtsClient) RegisterDmsKafkaInstance(request *model.RegisterDmsKafkaInstanceRequest) (*model.RegisterDmsKafkaInstanceResponse, error) {
	requestDef := GenReqDefForRegisterDmsKafkaInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDmsKafkaInstanceResponse), nil
	}
}

//该接口用于查询单个通知模板
func (c *LtsClient) ShowNotificationTemplate(request *model.ShowNotificationTemplateRequest) (*model.ShowNotificationTemplateResponse, error) {
	requestDef := GenReqDefForShowNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotificationTemplateResponse), nil
	}
}

//该接口用于查询指定日志流下的结构化配置内容。
func (c *LtsClient) ShowStructTemplate(request *model.ShowStructTemplateRequest) (*model.ShowStructTemplateResponse, error) {
	requestDef := GenReqDefForShowStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStructTemplateResponse), nil
	}
}

//修改日志接入
func (c *LtsClient) UpdateAccessConfig(request *model.UpdateAccessConfigRequest) (*model.UpdateAccessConfigResponse, error) {
	requestDef := GenReqDefForUpdateAccessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessConfigResponse), nil
	}
}

//修改主机组
func (c *LtsClient) UpdateHostGroup(request *model.UpdateHostGroupRequest) (*model.UpdateHostGroupResponse, error) {
	requestDef := GenReqDefForUpdateHostGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostGroupResponse), nil
	}
}

//该接口用于修改关键词告警。
func (c *LtsClient) UpdateKeywordsAlarmRule(request *model.UpdateKeywordsAlarmRuleRequest) (*model.UpdateKeywordsAlarmRuleResponse, error) {
	requestDef := GenReqDefForUpdateKeywordsAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeywordsAlarmRuleResponse), nil
	}
}

//该接口用于修改指定日志组下的日志存储时长。
func (c *LtsClient) UpdateLogGroup(request *model.UpdateLogGroupRequest) (*model.UpdateLogGroupResponse, error) {
	requestDef := GenReqDefForUpdateLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogGroupResponse), nil
	}
}

//该接口用于修改通知模板,根据名称进行修改。
func (c *LtsClient) UpdateNotificationTemplate(request *model.UpdateNotificationTemplateRequest) (*model.UpdateNotificationTemplateResponse, error) {
	requestDef := GenReqDefForUpdateNotificationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationTemplateResponse), nil
	}
}

//该接口用于修改指定日志流下的结构化配置。
func (c *LtsClient) UpdateStructTemplate(request *model.UpdateStructTemplateRequest) (*model.UpdateStructTemplateResponse, error) {
	requestDef := GenReqDefForUpdateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStructTemplateResponse), nil
	}
}

//该接口用于更新OBS转储，DIS转储，DMS转储。
func (c *LtsClient) UpdateTransfer(request *model.UpdateTransferRequest) (*model.UpdateTransferResponse, error) {
	requestDef := GenReqDefForUpdateTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTransferResponse), nil
	}
}

//该接口用于创建aom日志接入lts规则
func (c *LtsClient) CreateAomMappingRules(request *model.CreateAomMappingRulesRequest) (*model.CreateAomMappingRulesResponse, error) {
	requestDef := GenReqDefForCreateAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAomMappingRulesResponse), nil
	}
}

//该接口用于删除lts接入规则。
func (c *LtsClient) DeleteAomMappingRules(request *model.DeleteAomMappingRulesRequest) (*model.DeleteAomMappingRulesResponse, error) {
	requestDef := GenReqDefForDeleteAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAomMappingRulesResponse), nil
	}
}

//该接口用于查询单个aom日志接入lts
func (c *LtsClient) ShowAomMappingRule(request *model.ShowAomMappingRuleRequest) (*model.ShowAomMappingRuleResponse, error) {
	requestDef := GenReqDefForShowAomMappingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAomMappingRuleResponse), nil
	}
}

//该接口用于查询aom日志所有接入lts规则
func (c *LtsClient) ShowAomMappingRules(request *model.ShowAomMappingRulesRequest) (*model.ShowAomMappingRulesResponse, error) {
	requestDef := GenReqDefForShowAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAomMappingRulesResponse), nil
	}
}

//该接口用于修改接入规则
func (c *LtsClient) UpdateAomMappingRules(request *model.UpdateAomMappingRulesRequest) (*model.UpdateAomMappingRulesResponse, error) {
	requestDef := GenReqDefForUpdateAomMappingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAomMappingRulesResponse), nil
	}
}

//该接口用于创建SQL告警，目前每个帐户最多可以创建共200个关键词告警与SQL告警
func (c *LtsClient) CreateSqlAlarmRule(request *model.CreateSqlAlarmRuleRequest) (*model.CreateSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForCreateSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlAlarmRuleResponse), nil
	}
}

//该接口用于删除SQL告警
func (c *LtsClient) DeleteSqlAlarmRule(request *model.DeleteSqlAlarmRuleRequest) (*model.DeleteSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForDeleteSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlAlarmRuleResponse), nil
	}
}

//该接口用于查询SQL告警
func (c *LtsClient) ListSqlAlarmRules(request *model.ListSqlAlarmRulesRequest) (*model.ListSqlAlarmRulesResponse, error) {
	requestDef := GenReqDefForListSqlAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlAlarmRulesResponse), nil
	}
}

//改变告警规则状态
func (c *LtsClient) UpdateAlarmRuleStatus(request *model.UpdateAlarmRuleStatusRequest) (*model.UpdateAlarmRuleStatusResponse, error) {
	requestDef := GenReqDefForUpdateAlarmRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmRuleStatusResponse), nil
	}
}

//该接口用于修改SQL告警
func (c *LtsClient) UpdateSqlAlarmRule(request *model.UpdateSqlAlarmRuleRequest) (*model.UpdateSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForUpdateSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlAlarmRuleResponse), nil
	}
}
