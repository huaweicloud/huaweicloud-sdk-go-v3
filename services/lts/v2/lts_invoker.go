package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
)

type CreateAccessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessConfigInvoker) Invoke() (*model.CreateAccessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessConfigResponse), nil
	}
}

type CreateAgencyAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyAccessInvoker) Invoke() (*model.CreateAgencyAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyAccessResponse), nil
	}
}

type CreateDashBoardInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDashBoardInvoker) Invoke() (*model.CreateDashBoardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDashBoardResponse), nil
	}
}

type CreateDashboardGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDashboardGroupInvoker) Invoke() (*model.CreateDashboardGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDashboardGroupResponse), nil
	}
}

type CreateHostGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostGroupInvoker) Invoke() (*model.CreateHostGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostGroupResponse), nil
	}
}

type CreateKeywordsAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKeywordsAlarmRuleInvoker) Invoke() (*model.CreateKeywordsAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKeywordsAlarmRuleResponse), nil
	}
}

type CreateLogDumpObsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogDumpObsInvoker) Invoke() (*model.CreateLogDumpObsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogDumpObsResponse), nil
	}
}

type CreateLogGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogGroupInvoker) Invoke() (*model.CreateLogGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogGroupResponse), nil
	}
}

type CreateLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogStreamInvoker) Invoke() (*model.CreateLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogStreamResponse), nil
	}
}

type CreateLogStreamIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogStreamIndexInvoker) Invoke() (*model.CreateLogStreamIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogStreamIndexResponse), nil
	}
}

type CreateNotificationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotificationTemplateInvoker) Invoke() (*model.CreateNotificationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotificationTemplateResponse), nil
	}
}

type CreateSearchCriteriasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchCriteriasInvoker) Invoke() (*model.CreateSearchCriteriasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchCriteriasResponse), nil
	}
}

type CreateStructConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStructConfigInvoker) Invoke() (*model.CreateStructConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStructConfigResponse), nil
	}
}

type CreateStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStructTemplateInvoker) Invoke() (*model.CreateStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStructTemplateResponse), nil
	}
}

type CreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagsInvoker) Invoke() (*model.CreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagsResponse), nil
	}
}

type CreateTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTransferInvoker) Invoke() (*model.CreateTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTransferResponse), nil
	}
}

type CreatefavoriteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatefavoriteInvoker) Invoke() (*model.CreatefavoriteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatefavoriteResponse), nil
	}
}

type DeleteAccessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccessConfigInvoker) Invoke() (*model.DeleteAccessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccessConfigResponse), nil
	}
}

type DeleteActiveAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteActiveAlarmsInvoker) Invoke() (*model.DeleteActiveAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteActiveAlarmsResponse), nil
	}
}

type DeleteDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDashboardInvoker) Invoke() (*model.DeleteDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDashboardResponse), nil
	}
}

type DeleteHostGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostGroupInvoker) Invoke() (*model.DeleteHostGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostGroupResponse), nil
	}
}

type DeleteKeywordsAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKeywordsAlarmRuleInvoker) Invoke() (*model.DeleteKeywordsAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKeywordsAlarmRuleResponse), nil
	}
}

type DeleteLogGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogGroupInvoker) Invoke() (*model.DeleteLogGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogGroupResponse), nil
	}
}

type DeleteLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogStreamInvoker) Invoke() (*model.DeleteLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogStreamResponse), nil
	}
}

type DeleteNotificationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotificationTemplateInvoker) Invoke() (*model.DeleteNotificationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotificationTemplateResponse), nil
	}
}

type DeleteSearchCriteriasInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSearchCriteriasInvoker) Invoke() (*model.DeleteSearchCriteriasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSearchCriteriasResponse), nil
	}
}

type DeleteStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStructTemplateInvoker) Invoke() (*model.DeleteStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStructTemplateResponse), nil
	}
}

type DeleteTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTransferInvoker) Invoke() (*model.DeleteTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTransferResponse), nil
	}
}

type DeletefavoriteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletefavoriteInvoker) Invoke() (*model.DeletefavoriteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletefavoriteResponse), nil
	}
}

type DisableLogCollectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableLogCollectionInvoker) Invoke() (*model.DisableLogCollectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableLogCollectionResponse), nil
	}
}

type EnableLogCollectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableLogCollectionInvoker) Invoke() (*model.EnableLogCollectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableLogCollectionResponse), nil
	}
}

type ListAccessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessConfigInvoker) Invoke() (*model.ListAccessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessConfigResponse), nil
	}
}

type ListActiveOrHistoryAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListActiveOrHistoryAlarmsInvoker) Invoke() (*model.ListActiveOrHistoryAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListActiveOrHistoryAlarmsResponse), nil
	}
}

type ListBreifStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBreifStructTemplateInvoker) Invoke() (*model.ListBreifStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBreifStructTemplateResponse), nil
	}
}

type ListChartsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChartsInvoker) Invoke() (*model.ListChartsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChartsResponse), nil
	}
}

type ListCriteriasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCriteriasInvoker) Invoke() (*model.ListCriteriasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCriteriasResponse), nil
	}
}

type ListHistorySqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistorySqlInvoker) Invoke() (*model.ListHistorySqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistorySqlResponse), nil
	}
}

type ListHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostInvoker) Invoke() (*model.ListHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostResponse), nil
	}
}

type ListHostGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupInvoker) Invoke() (*model.ListHostGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupResponse), nil
	}
}

type ListKeywordsAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeywordsAlarmRulesInvoker) Invoke() (*model.ListKeywordsAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeywordsAlarmRulesResponse), nil
	}
}

type ListLogGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogGroupsInvoker) Invoke() (*model.ListLogGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogGroupsResponse), nil
	}
}

type ListLogHistogramInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogHistogramInvoker) Invoke() (*model.ListLogHistogramResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogHistogramResponse), nil
	}
}

type ListLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogStreamInvoker) Invoke() (*model.ListLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogStreamResponse), nil
	}
}

type ListLogStreamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogStreamsInvoker) Invoke() (*model.ListLogStreamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogStreamsResponse), nil
	}
}

type ListLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogsInvoker) Invoke() (*model.ListLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogsResponse), nil
	}
}

type ListNotificationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationTemplateInvoker) Invoke() (*model.ListNotificationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationTemplateResponse), nil
	}
}

type ListNotificationTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationTemplatesInvoker) Invoke() (*model.ListNotificationTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationTemplatesResponse), nil
	}
}

type ListNotificationTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationTopicsInvoker) Invoke() (*model.ListNotificationTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationTopicsResponse), nil
	}
}

type ListQueryAllSearchCriteriasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryAllSearchCriteriasInvoker) Invoke() (*model.ListQueryAllSearchCriteriasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryAllSearchCriteriasResponse), nil
	}
}

type ListQueryStructuredLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryStructuredLogsInvoker) Invoke() (*model.ListQueryStructuredLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryStructuredLogsResponse), nil
	}
}

type ListStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStructTemplateInvoker) Invoke() (*model.ListStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStructTemplateResponse), nil
	}
}

type ListStructuredLogsWithTimeRangeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStructuredLogsWithTimeRangeInvoker) Invoke() (*model.ListStructuredLogsWithTimeRangeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStructuredLogsWithTimeRangeResponse), nil
	}
}

type ListTimeLineTrafficStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTimeLineTrafficStatisticsInvoker) Invoke() (*model.ListTimeLineTrafficStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTimeLineTrafficStatisticsResponse), nil
	}
}

type ListTopnTrafficStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopnTrafficStatisticsInvoker) Invoke() (*model.ListTopnTrafficStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopnTrafficStatisticsResponse), nil
	}
}

type ListTransfersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransfersInvoker) Invoke() (*model.ListTransfersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransfersResponse), nil
	}
}

type RegisterDmsKafkaInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterDmsKafkaInstanceInvoker) Invoke() (*model.RegisterDmsKafkaInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterDmsKafkaInstanceResponse), nil
	}
}

type ShowAdminConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAdminConfigInvoker) Invoke() (*model.ShowAdminConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAdminConfigResponse), nil
	}
}

type ShowLogConvergeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogConvergeConfigInvoker) Invoke() (*model.ShowLogConvergeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogConvergeConfigResponse), nil
	}
}

type ShowMemberGroupAndStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberGroupAndStreamInvoker) Invoke() (*model.ShowMemberGroupAndStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberGroupAndStreamResponse), nil
	}
}

type ShowNotificationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotificationTemplateInvoker) Invoke() (*model.ShowNotificationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotificationTemplateResponse), nil
	}
}

type ShowStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStructTemplateInvoker) Invoke() (*model.ShowStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStructTemplateResponse), nil
	}
}

type UpdateAccessConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessConfigInvoker) Invoke() (*model.UpdateAccessConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessConfigResponse), nil
	}
}

type UpdateHostGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostGroupInvoker) Invoke() (*model.UpdateHostGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostGroupResponse), nil
	}
}

type UpdateKeywordsAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKeywordsAlarmRuleInvoker) Invoke() (*model.UpdateKeywordsAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKeywordsAlarmRuleResponse), nil
	}
}

type UpdateLogConvergeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogConvergeConfigInvoker) Invoke() (*model.UpdateLogConvergeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogConvergeConfigResponse), nil
	}
}

type UpdateLogGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogGroupInvoker) Invoke() (*model.UpdateLogGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogGroupResponse), nil
	}
}

type UpdateLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogStreamInvoker) Invoke() (*model.UpdateLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogStreamResponse), nil
	}
}

type UpdateNotificationTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotificationTemplateInvoker) Invoke() (*model.UpdateNotificationTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotificationTemplateResponse), nil
	}
}

type UpdateStructConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStructConfigInvoker) Invoke() (*model.UpdateStructConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStructConfigResponse), nil
	}
}

type UpdateStructTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStructTemplateInvoker) Invoke() (*model.UpdateStructTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStructTemplateResponse), nil
	}
}

type UpdateSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSwitchInvoker) Invoke() (*model.UpdateSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSwitchResponse), nil
	}
}

type UpdateTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTransferInvoker) Invoke() (*model.UpdateTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTransferResponse), nil
	}
}

type CreateAomMappingRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAomMappingRulesInvoker) Invoke() (*model.CreateAomMappingRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAomMappingRulesResponse), nil
	}
}

type DeleteAomMappingRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAomMappingRulesInvoker) Invoke() (*model.DeleteAomMappingRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAomMappingRulesResponse), nil
	}
}

type ShowAomMappingRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAomMappingRuleInvoker) Invoke() (*model.ShowAomMappingRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAomMappingRuleResponse), nil
	}
}

type ShowAomMappingRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAomMappingRulesInvoker) Invoke() (*model.ShowAomMappingRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAomMappingRulesResponse), nil
	}
}

type UpdateAomMappingRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAomMappingRulesInvoker) Invoke() (*model.UpdateAomMappingRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAomMappingRulesResponse), nil
	}
}

type ConsumerGroupHeartBeatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConsumerGroupHeartBeatInvoker) Invoke() (*model.ConsumerGroupHeartBeatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConsumerGroupHeartBeatResponse), nil
	}
}

type CreateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConsumerGroupInvoker) Invoke() (*model.CreateConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConsumerGroupResponse), nil
	}
}

type DeleteConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConsumerGroupInvoker) Invoke() (*model.DeleteConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConsumerGroupResponse), nil
	}
}

type ListConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumerGroupInvoker) Invoke() (*model.ListConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumerGroupResponse), nil
	}
}

type ListDetailsConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDetailsConsumerGroupInvoker) Invoke() (*model.ListDetailsConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDetailsConsumerGroupResponse), nil
	}
}

type ShowCursorByTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCursorByTimeInvoker) Invoke() (*model.ShowCursorByTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCursorByTimeResponse), nil
	}
}

type ShowCursorTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCursorTimeInvoker) Invoke() (*model.ShowCursorTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCursorTimeResponse), nil
	}
}

type ShowLogStreamShardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogStreamShardsInvoker) Invoke() (*model.ShowLogStreamShardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogStreamShardsResponse), nil
	}
}

type UpdateCheckPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCheckPointInvoker) Invoke() (*model.UpdateCheckPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCheckPointResponse), nil
	}
}

type CreateSqlAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlAlarmRuleInvoker) Invoke() (*model.CreateSqlAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlAlarmRuleResponse), nil
	}
}

type DeleteSqlAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlAlarmRuleInvoker) Invoke() (*model.DeleteSqlAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlAlarmRuleResponse), nil
	}
}

type ListSqlAlarmRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlAlarmRulesInvoker) Invoke() (*model.ListSqlAlarmRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlAlarmRulesResponse), nil
	}
}

type UpdateAlarmRuleStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmRuleStatusInvoker) Invoke() (*model.UpdateAlarmRuleStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmRuleStatusResponse), nil
	}
}

type UpdateSqlAlarmRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlAlarmRuleInvoker) Invoke() (*model.UpdateSqlAlarmRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlAlarmRuleResponse), nil
	}
}
