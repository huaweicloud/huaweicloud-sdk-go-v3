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

//该接口用于创建指定日志流下的结构化配置。
func (c *LtsClient) CreateStructTemplate(request *model.CreateStructTemplateRequest) (*model.CreateStructTemplateResponse, error) {
	requestDef := GenReqDefForCreateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStructTemplateResponse), nil
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

//该接口用于删除指定日志流下的结构化配置。
func (c *LtsClient) DeleteStructTemplate(request *model.DeleteStructTemplateRequest) (*model.DeleteStructTemplateResponse, error) {
	requestDef := GenReqDefForDeleteStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStructTemplateResponse), nil
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

//该接口用于查询指定日志组下的所有日志流信息。
func (c *LtsClient) ListLogStream(request *model.ListLogStreamRequest) (*model.ListLogStreamResponse, error) {
	requestDef := GenReqDefForListLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamResponse), nil
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

//该接口用于查询指定日志流下的结构化配置内容。
func (c *LtsClient) ShowStructTemplate(request *model.ShowStructTemplateRequest) (*model.ShowStructTemplateResponse, error) {
	requestDef := GenReqDefForShowStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStructTemplateResponse), nil
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

//该接口用于修改指定日志流下的结构化配置。
func (c *LtsClient) UpdateStructTemplate(request *model.UpdateStructTemplateRequest) (*model.UpdateStructTemplateResponse, error) {
	requestDef := GenReqDefForUpdateStructTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStructTemplateResponse), nil
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

//该接口用于修改SQL告警
func (c *LtsClient) UpdateSqlAlarmRule(request *model.UpdateSqlAlarmRuleRequest) (*model.UpdateSqlAlarmRuleResponse, error) {
	requestDef := GenReqDefForUpdateSqlAlarmRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlAlarmRuleResponse), nil
	}
}
