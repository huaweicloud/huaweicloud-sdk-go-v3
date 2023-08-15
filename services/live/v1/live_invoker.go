package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/live/v1/model"
)

type BatchShowIpBelongsInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateDomainMappingInvoker) Invoke() (*model.CreateDomainMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainMappingResponse), nil
	}
}

type CreateRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateRecordRuleInvoker) Invoke() (*model.CreateRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordRuleResponse), nil
	}
}

type CreateSnapshotConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateStreamForbiddenInvoker) Invoke() (*model.CreateStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamForbiddenResponse), nil
	}
}

type CreateTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTranscodingsTemplateInvoker) Invoke() (*model.CreateTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTranscodingsTemplateResponse), nil
	}
}

type DeleteDomainInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteDomainMappingInvoker) Invoke() (*model.DeleteDomainMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainMappingResponse), nil
	}
}

type DeleteRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteRecordRuleInvoker) Invoke() (*model.DeleteRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordRuleResponse), nil
	}
}

type DeleteSnapshotConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteTranscodingsTemplateInvoker) Invoke() (*model.DeleteTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTranscodingsTemplateResponse), nil
	}
}

type ListLiveSampleLogsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListLiveStreamsOnlineInvoker) Invoke() (*model.ListLiveStreamsOnlineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveStreamsOnlineResponse), nil
	}
}

type ListRecordCallbackConfigsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListRecordRulesInvoker) Invoke() (*model.ListRecordRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordRulesResponse), nil
	}
}

type ListSnapshotConfigsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListStreamForbiddenInvoker) Invoke() (*model.ListStreamForbiddenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStreamForbiddenResponse), nil
	}
}

type RunRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRecordInvoker) Invoke() (*model.RunRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRecordResponse), nil
	}
}

type ShowDomainInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowDomainKeyChainInvoker) Invoke() (*model.ShowDomainKeyChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainKeyChainResponse), nil
	}
}

type ShowRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowRecordRuleInvoker) Invoke() (*model.ShowRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordRuleResponse), nil
	}
}

type ShowTranscodingsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTranscodingsTemplateInvoker) Invoke() (*model.ShowTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTranscodingsTemplateResponse), nil
	}
}

type UpdateDomainInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateDomainKeyChainInvoker) Invoke() (*model.UpdateDomainKeyChainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainKeyChainResponse), nil
	}
}

type UpdateRecordCallbackConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateTranscodingsTemplateInvoker) Invoke() (*model.UpdateTranscodingsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTranscodingsTemplateResponse), nil
	}
}

type DeleteDomainHttpsCertInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateDomainHttpsCertInvoker) Invoke() (*model.UpdateDomainHttpsCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainHttpsCertResponse), nil
	}
}

type UpdateObsBucketAuthorityPublicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateObsBucketAuthorityPublicInvoker) Invoke() (*model.UpdateObsBucketAuthorityPublicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateObsBucketAuthorityPublicResponse), nil
	}
}
