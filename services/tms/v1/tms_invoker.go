package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/tms/v1/model"
)

type ChangeAssociatedResourceOpenStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAssociatedResourceOpenStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAssociatedResourceOpenStatusInvoker) Invoke() (*model.ChangeAssociatedResourceOpenStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAssociatedResourceOpenStatusResponse), nil
	}
}

type CreateAssociatedResourceRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssociatedResourceRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssociatedResourceRulesInvoker) Invoke() (*model.CreateAssociatedResourceRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssociatedResourceRulesResponse), nil
	}
}

type CreatePredefineTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePredefineTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePredefineTagsInvoker) Invoke() (*model.CreatePredefineTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePredefineTagsResponse), nil
	}
}

type CreateResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceTagInvoker) Invoke() (*model.CreateResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceTagResponse), nil
	}
}

type DeleteAssociatedResourceRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssociatedResourceRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssociatedResourceRuleInvoker) Invoke() (*model.DeleteAssociatedResourceRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssociatedResourceRuleResponse), nil
	}
}

type DeletePredefineTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePredefineTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePredefineTagsInvoker) Invoke() (*model.DeletePredefineTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePredefineTagsResponse), nil
	}
}

type DeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceTagInvoker) Invoke() (*model.DeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListAssociatedResourceRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociatedResourceRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociatedResourceRulesInvoker) Invoke() (*model.ListAssociatedResourceRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociatedResourceRulesResponse), nil
	}
}

type ListAssociatedResourceSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociatedResourceSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociatedResourceSettingsInvoker) Invoke() (*model.ListAssociatedResourceSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociatedResourceSettingsResponse), nil
	}
}

type ListPredefineTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPredefineTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPredefineTagsInvoker) Invoke() (*model.ListPredefineTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPredefineTagsResponse), nil
	}
}

type ListProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProvidersInvoker) Invoke() (*model.ListProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProvidersResponse), nil
	}
}

type ListResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInvoker) Invoke() (*model.ListResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceResponse), nil
	}
}

type ListTagKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagKeysInvoker) Invoke() (*model.ListTagKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagKeysResponse), nil
	}
}

type ListTagValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagValuesInvoker) Invoke() (*model.ListTagValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagValuesResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}

type ShowAssociatedResourceOpenStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssociatedResourceOpenStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssociatedResourceOpenStatusInvoker) Invoke() (*model.ShowAssociatedResourceOpenStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssociatedResourceOpenStatusResponse), nil
	}
}

type ShowResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceTagInvoker) Invoke() (*model.ShowResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagResponse), nil
	}
}

type ShowTagQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagQuotaInvoker) Invoke() (*model.ShowTagQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagQuotaResponse), nil
	}
}

type UpdateAssociatedResourceRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssociatedResourceRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssociatedResourceRulesInvoker) Invoke() (*model.UpdateAssociatedResourceRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssociatedResourceRulesResponse), nil
	}
}

type UpdatePredefineTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePredefineTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePredefineTagsInvoker) Invoke() (*model.UpdatePredefineTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePredefineTagsResponse), nil
	}
}
