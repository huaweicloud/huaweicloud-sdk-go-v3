package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v1/model"
)

type BatchShowIpBelongsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowIpBelongsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowIpBelongsInvoker) Invoke() (*model.BatchShowIpBelongsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowIpBelongsResponse), nil
	}
}

type CreateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainInvoker) Invoke() (*model.CreateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainResponse), nil
	}
}

type CreateDomainMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainMappingInvoker) Invoke() (*model.CreateDomainMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainMappingResponse), nil
	}
}

type CreateFlowOutputInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowOutputInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlowOutputInvoker) Invoke() (*model.CreateFlowOutputResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowOutputResponse), nil
	}
}

type CreateFlowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlowsInvoker) Invoke() (*model.CreateFlowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowsResponse), nil
	}
}

type CreateRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordCallbackConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordCallbackConfigInvoker) Invoke() (*model.CreateRecordCallbackConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordCallbackConfigResponse), nil
	}
}

type CreateRecordIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordIndexInvoker) Invoke() (*model.CreateRecordIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordIndexResponse), nil
	}
}

type CreateRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordRuleInvoker) Invoke() (*model.CreateRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordRuleResponse), nil
	}
}

type CreateScheduleRecordTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduleRecordTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScheduleRecordTasksInvoker) Invoke() (*model.CreateScheduleRecordTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduleRecordTasksResponse), nil
	}
}

type CreateSnapshotConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSnapshotConfigInvoker) Invoke() (*model.CreateSnapshotConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotConfigResponse), nil
	}
}

type CreateStreamForbiddenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStreamForbiddenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStreamForbiddenInvoker) Invoke() (*model.CreateStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamForbiddenResponse), nil
	}
}

type CreateStreamForbiddenOnceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStreamForbiddenOnceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStreamForbiddenOnceInvoker) Invoke() (*model.CreateStreamForbiddenOnceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamForbiddenOnceResponse), nil
	}
}

type CreateTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTranscodingsTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTranscodingsTemplateInvoker) Invoke() (*model.CreateTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTranscodingsTemplateResponse), nil
	}
}

type CreateUrlAuthchainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUrlAuthchainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUrlAuthchainInvoker) Invoke() (*model.CreateUrlAuthchainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUrlAuthchainResponse), nil
	}
}

type DeleteDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainInvoker) Invoke() (*model.DeleteDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainResponse), nil
	}
}

type DeleteDomainKeyChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainKeyChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainKeyChainInvoker) Invoke() (*model.DeleteDomainKeyChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainKeyChainResponse), nil
	}
}

type DeleteDomainMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainMappingInvoker) Invoke() (*model.DeleteDomainMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainMappingResponse), nil
	}
}

type DeleteFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlowInvoker) Invoke() (*model.DeleteFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlowResponse), nil
	}
}

type DeleteFlowOutputInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlowOutputInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlowOutputInvoker) Invoke() (*model.DeleteFlowOutputResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlowOutputResponse), nil
	}
}

type DeletePublishTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublishTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePublishTemplateInvoker) Invoke() (*model.DeletePublishTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublishTemplateResponse), nil
	}
}

type DeleteRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordCallbackConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecordCallbackConfigInvoker) Invoke() (*model.DeleteRecordCallbackConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordCallbackConfigResponse), nil
	}
}

type DeleteRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecordRuleInvoker) Invoke() (*model.DeleteRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordRuleResponse), nil
	}
}

type DeleteRefererChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRefererChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRefererChainInvoker) Invoke() (*model.DeleteRefererChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRefererChainResponse), nil
	}
}

type DeleteScheduleRecordTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleRecordTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduleRecordTasksInvoker) Invoke() (*model.DeleteScheduleRecordTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleRecordTasksResponse), nil
	}
}

type DeleteSnapshotConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSnapshotConfigInvoker) Invoke() (*model.DeleteSnapshotConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotConfigResponse), nil
	}
}

type DeleteStreamForbiddenInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStreamForbiddenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStreamForbiddenInvoker) Invoke() (*model.DeleteStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStreamForbiddenResponse), nil
	}
}

type DeleteTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTranscodingsTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTranscodingsTemplateInvoker) Invoke() (*model.DeleteTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTranscodingsTemplateResponse), nil
	}
}

type ListDelayConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDelayConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDelayConfigInvoker) Invoke() (*model.ListDelayConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDelayConfigResponse), nil
	}
}

type ListFlowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlowsInvoker) Invoke() (*model.ListFlowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowsResponse), nil
	}
}

type ListGeoBlockingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeoBlockingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeoBlockingConfigInvoker) Invoke() (*model.ListGeoBlockingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeoBlockingConfigResponse), nil
	}
}

type ListHlsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHlsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHlsConfigInvoker) Invoke() (*model.ListHlsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHlsConfigResponse), nil
	}
}

type ListIpAuthListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpAuthListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpAuthListInvoker) Invoke() (*model.ListIpAuthListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpAuthListResponse), nil
	}
}

type ListLiveSampleLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveSampleLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveSampleLogsInvoker) Invoke() (*model.ListLiveSampleLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveSampleLogsResponse), nil
	}
}

type ListLiveStreamsOnlineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveStreamsOnlineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveStreamsOnlineInvoker) Invoke() (*model.ListLiveStreamsOnlineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveStreamsOnlineResponse), nil
	}
}

type ListPublishTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublishTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublishTemplateInvoker) Invoke() (*model.ListPublishTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublishTemplateResponse), nil
	}
}

type ListRecordCallbackConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordCallbackConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordCallbackConfigsInvoker) Invoke() (*model.ListRecordCallbackConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordCallbackConfigsResponse), nil
	}
}

type ListRecordContentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordContentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordContentsInvoker) Invoke() (*model.ListRecordContentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordContentsResponse), nil
	}
}

type ListRecordRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordRulesInvoker) Invoke() (*model.ListRecordRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordRulesResponse), nil
	}
}

type ListScheduleRecordTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduleRecordTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduleRecordTasksInvoker) Invoke() (*model.ListScheduleRecordTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduleRecordTasksResponse), nil
	}
}

type ListSnapshotConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotConfigsInvoker) Invoke() (*model.ListSnapshotConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotConfigsResponse), nil
	}
}

type ListStreamForbiddenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStreamForbiddenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStreamForbiddenInvoker) Invoke() (*model.ListStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStreamForbiddenResponse), nil
	}
}

type ModifyFlowOutputInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyFlowOutputInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyFlowOutputInvoker) Invoke() (*model.ModifyFlowOutputResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyFlowOutputResponse), nil
	}
}

type ModifyFlowSourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyFlowSourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyFlowSourcesInvoker) Invoke() (*model.ModifyFlowSourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyFlowSourcesResponse), nil
	}
}

type ModifyFlowStartInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyFlowStartInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyFlowStartInvoker) Invoke() (*model.ModifyFlowStartResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyFlowStartResponse), nil
	}
}

type ModifyFlowStopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyFlowStopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyFlowStopInvoker) Invoke() (*model.ModifyFlowStopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyFlowStopResponse), nil
	}
}

type RunRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunRecordInvoker) Invoke() (*model.RunRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRecordResponse), nil
	}
}

type SetRefererChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRefererChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRefererChainInvoker) Invoke() (*model.SetRefererChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRefererChainResponse), nil
	}
}

type ShowDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainInvoker) Invoke() (*model.ShowDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainResponse), nil
	}
}

type ShowDomainKeyChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainKeyChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainKeyChainInvoker) Invoke() (*model.ShowDomainKeyChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainKeyChainResponse), nil
	}
}

type ShowFlowDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowDetailInvoker) Invoke() (*model.ShowFlowDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowDetailResponse), nil
	}
}

type ShowOutputInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOutputInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOutputInfoInvoker) Invoke() (*model.ShowOutputInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOutputInfoResponse), nil
	}
}

type ShowPullSourcesConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPullSourcesConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPullSourcesConfigInvoker) Invoke() (*model.ShowPullSourcesConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPullSourcesConfigResponse), nil
	}
}

type ShowRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordCallbackConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordCallbackConfigInvoker) Invoke() (*model.ShowRecordCallbackConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordCallbackConfigResponse), nil
	}
}

type ShowRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordRuleInvoker) Invoke() (*model.ShowRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordRuleResponse), nil
	}
}

type ShowRefererChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRefererChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRefererChainInvoker) Invoke() (*model.ShowRefererChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRefererChainResponse), nil
	}
}

type ShowTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTranscodingsTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTranscodingsTemplateInvoker) Invoke() (*model.ShowTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTranscodingsTemplateResponse), nil
	}
}

type UpdateDelayConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDelayConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDelayConfigInvoker) Invoke() (*model.UpdateDelayConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDelayConfigResponse), nil
	}
}

type UpdateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainInvoker) Invoke() (*model.UpdateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainResponse), nil
	}
}

type UpdateDomainIp6SwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainIp6SwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainIp6SwitchInvoker) Invoke() (*model.UpdateDomainIp6SwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainIp6SwitchResponse), nil
	}
}

type UpdateDomainKeyChainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainKeyChainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainKeyChainInvoker) Invoke() (*model.UpdateDomainKeyChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainKeyChainResponse), nil
	}
}

type UpdateGeoBlockingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGeoBlockingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGeoBlockingConfigInvoker) Invoke() (*model.UpdateGeoBlockingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGeoBlockingConfigResponse), nil
	}
}

type UpdateHlsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHlsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHlsConfigInvoker) Invoke() (*model.UpdateHlsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHlsConfigResponse), nil
	}
}

type UpdateIpAuthListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpAuthListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIpAuthListInvoker) Invoke() (*model.UpdateIpAuthListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpAuthListResponse), nil
	}
}

type UpdatePublishTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublishTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePublishTemplateInvoker) Invoke() (*model.UpdatePublishTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublishTemplateResponse), nil
	}
}

type UpdatePullSourcesConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePullSourcesConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePullSourcesConfigInvoker) Invoke() (*model.UpdatePullSourcesConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePullSourcesConfigResponse), nil
	}
}

type UpdateRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordCallbackConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecordCallbackConfigInvoker) Invoke() (*model.UpdateRecordCallbackConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordCallbackConfigResponse), nil
	}
}

type UpdateRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecordRuleInvoker) Invoke() (*model.UpdateRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordRuleResponse), nil
	}
}

type UpdateSnapshotConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSnapshotConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSnapshotConfigInvoker) Invoke() (*model.UpdateSnapshotConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSnapshotConfigResponse), nil
	}
}

type UpdateStreamForbiddenInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStreamForbiddenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStreamForbiddenInvoker) Invoke() (*model.UpdateStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStreamForbiddenResponse), nil
	}
}

type UpdateTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTranscodingsTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTranscodingsTemplateInvoker) Invoke() (*model.UpdateTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTranscodingsTemplateResponse), nil
	}
}

type ListCesDimsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCesDimsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCesDimsInfoInvoker) Invoke() (*model.ListCesDimsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCesDimsInfoResponse), nil
	}
}

type ListCesInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCesInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCesInstanceInvoker) Invoke() (*model.ListCesInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCesInstanceResponse), nil
	}
}

type DeleteDomainHttpsCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainHttpsCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainHttpsCertInvoker) Invoke() (*model.DeleteDomainHttpsCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainHttpsCertResponse), nil
	}
}

type ShowDomainHttpsCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainHttpsCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainHttpsCertInvoker) Invoke() (*model.ShowDomainHttpsCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainHttpsCertResponse), nil
	}
}

type UpdateDomainHttpsCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainHttpsCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainHttpsCertInvoker) Invoke() (*model.UpdateDomainHttpsCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainHttpsCertResponse), nil
	}
}

type CreateHarvestTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHarvestTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHarvestTaskInvoker) Invoke() (*model.CreateHarvestTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHarvestTaskResponse), nil
	}
}

type DeleteHarvestTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHarvestTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHarvestTaskInvoker) Invoke() (*model.DeleteHarvestTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHarvestTaskResponse), nil
	}
}

type ListHarvestTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHarvestTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHarvestTaskInvoker) Invoke() (*model.ListHarvestTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHarvestTaskResponse), nil
	}
}

type ModifyHarvestTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyHarvestTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyHarvestTaskInvoker) Invoke() (*model.ModifyHarvestTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyHarvestTaskResponse), nil
	}
}

type UpdateHarvestJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHarvestJobStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHarvestJobStatusInvoker) Invoke() (*model.UpdateHarvestJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHarvestJobStatusResponse), nil
	}
}

type UpdateObsBucketAuthorityPublicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateObsBucketAuthorityPublicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateObsBucketAuthorityPublicInvoker) Invoke() (*model.UpdateObsBucketAuthorityPublicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateObsBucketAuthorityPublicResponse), nil
	}
}

type CreateOttChannelInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOttChannelInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOttChannelInfoInvoker) Invoke() (*model.CreateOttChannelInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOttChannelInfoResponse), nil
	}
}

type DeleteOttChannelInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOttChannelInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOttChannelInfoInvoker) Invoke() (*model.DeleteOttChannelInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOttChannelInfoResponse), nil
	}
}

type ListOttChannelInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOttChannelInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOttChannelInfoInvoker) Invoke() (*model.ListOttChannelInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOttChannelInfoResponse), nil
	}
}

type ModifyOttChannelInfoEncoderSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoEncoderSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoEncoderSettingsInvoker) Invoke() (*model.ModifyOttChannelInfoEncoderSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoEncoderSettingsResponse), nil
	}
}

type ModifyOttChannelInfoEndPointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoEndPointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoEndPointsInvoker) Invoke() (*model.ModifyOttChannelInfoEndPointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoEndPointsResponse), nil
	}
}

type ModifyOttChannelInfoGeneralInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoGeneralInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoGeneralInvoker) Invoke() (*model.ModifyOttChannelInfoGeneralResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoGeneralResponse), nil
	}
}

type ModifyOttChannelInfoInputInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoInputInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoInputInvoker) Invoke() (*model.ModifyOttChannelInfoInputResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoInputResponse), nil
	}
}

type ModifyOttChannelInfoRecordSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoRecordSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoRecordSettingsInvoker) Invoke() (*model.ModifyOttChannelInfoRecordSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoRecordSettingsResponse), nil
	}
}

type ModifyOttChannelInfoStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyOttChannelInfoStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyOttChannelInfoStatsInvoker) Invoke() (*model.ModifyOttChannelInfoStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyOttChannelInfoStatsResponse), nil
	}
}

type ShowChannelStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowChannelStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowChannelStatisticInvoker) Invoke() (*model.ShowChannelStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowChannelStatisticResponse), nil
	}
}
