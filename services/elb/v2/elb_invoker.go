package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"
)

type BatchCreateListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateListenerTagsInvoker) Invoke() (*model.BatchCreateListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateListenerTagsResponse), nil
	}
}

type BatchCreateLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateLoadbalancerTagsInvoker) Invoke() (*model.BatchCreateLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateLoadbalancerTagsResponse), nil
	}
}

type BatchDeleteListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteListenerTagsInvoker) Invoke() (*model.BatchDeleteListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteListenerTagsResponse), nil
	}
}

type BatchDeleteLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteLoadbalancerTagsInvoker) Invoke() (*model.BatchDeleteLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLoadbalancerTagsResponse), nil
	}
}

type CreateHealthmonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHealthmonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHealthmonitorInvoker) Invoke() (*model.CreateHealthmonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHealthmonitorResponse), nil
	}
}

type CreateL7policyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateL7policyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateL7policyInvoker) Invoke() (*model.CreateL7policyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateL7policyResponse), nil
	}
}

type CreateL7ruleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateL7ruleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateL7ruleInvoker) Invoke() (*model.CreateL7ruleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateL7ruleResponse), nil
	}
}

type CreateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateListenerInvoker) Invoke() (*model.CreateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateListenerResponse), nil
	}
}

type CreateListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateListenerTagsInvoker) Invoke() (*model.CreateListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateListenerTagsResponse), nil
	}
}

type CreateLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLoadbalancerInvoker) Invoke() (*model.CreateLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLoadbalancerResponse), nil
	}
}

type CreateLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLoadbalancerTagsInvoker) Invoke() (*model.CreateLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLoadbalancerTagsResponse), nil
	}
}

type CreateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMemberInvoker) Invoke() (*model.CreateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMemberResponse), nil
	}
}

type CreatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePoolInvoker) Invoke() (*model.CreatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePoolResponse), nil
	}
}

type CreateWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWhitelistInvoker) Invoke() (*model.CreateWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWhitelistResponse), nil
	}
}

type DeleteHealthmonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHealthmonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHealthmonitorInvoker) Invoke() (*model.DeleteHealthmonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHealthmonitorResponse), nil
	}
}

type DeleteL7policyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteL7policyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteL7policyInvoker) Invoke() (*model.DeleteL7policyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteL7policyResponse), nil
	}
}

type DeleteL7ruleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteL7ruleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteL7ruleInvoker) Invoke() (*model.DeleteL7ruleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteL7ruleResponse), nil
	}
}

type DeleteListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteListenerInvoker) Invoke() (*model.DeleteListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteListenerResponse), nil
	}
}

type DeleteListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteListenerTagsInvoker) Invoke() (*model.DeleteListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteListenerTagsResponse), nil
	}
}

type DeleteLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLoadbalancerInvoker) Invoke() (*model.DeleteLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadbalancerResponse), nil
	}
}

type DeleteLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLoadbalancerTagsInvoker) Invoke() (*model.DeleteLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadbalancerTagsResponse), nil
	}
}

type DeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMemberInvoker) Invoke() (*model.DeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberResponse), nil
	}
}

type DeletePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePoolInvoker) Invoke() (*model.DeletePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoolResponse), nil
	}
}

type DeleteWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWhitelistInvoker) Invoke() (*model.DeleteWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWhitelistResponse), nil
	}
}

type ListHealthmonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHealthmonitorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHealthmonitorsInvoker) Invoke() (*model.ListHealthmonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHealthmonitorsResponse), nil
	}
}

type ListL7policiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListL7policiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListL7policiesInvoker) Invoke() (*model.ListL7policiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListL7policiesResponse), nil
	}
}

type ListL7rulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListL7rulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListL7rulesInvoker) Invoke() (*model.ListL7rulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListL7rulesResponse), nil
	}
}

type ListListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListListenerTagsInvoker) Invoke() (*model.ListListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenerTagsResponse), nil
	}
}

type ListListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListListenersInvoker) Invoke() (*model.ListListenersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenersResponse), nil
	}
}

type ListListenersByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenersByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListListenersByTagsInvoker) Invoke() (*model.ListListenersByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenersByTagsResponse), nil
	}
}

type ListLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadbalancerTagsInvoker) Invoke() (*model.ListLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadbalancerTagsResponse), nil
	}
}

type ListLoadbalancersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadbalancersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadbalancersInvoker) Invoke() (*model.ListLoadbalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadbalancersResponse), nil
	}
}

type ListLoadbalancersByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadbalancersByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadbalancersByTagsInvoker) Invoke() (*model.ListLoadbalancersByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadbalancersByTagsResponse), nil
	}
}

type ListMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMembersInvoker) Invoke() (*model.ListMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMembersResponse), nil
	}
}

type ListPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolsInvoker) Invoke() (*model.ListPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolsResponse), nil
	}
}

type ListWhitelistsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWhitelistsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWhitelistsInvoker) Invoke() (*model.ListWhitelistsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWhitelistsResponse), nil
	}
}

type ShowHealthmonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHealthmonitorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHealthmonitorsInvoker) Invoke() (*model.ShowHealthmonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthmonitorsResponse), nil
	}
}

type ShowL7policyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowL7policyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowL7policyInvoker) Invoke() (*model.ShowL7policyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowL7policyResponse), nil
	}
}

type ShowL7ruleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowL7ruleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowL7ruleInvoker) Invoke() (*model.ShowL7ruleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowL7ruleResponse), nil
	}
}

type ShowListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListenerInvoker) Invoke() (*model.ShowListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerResponse), nil
	}
}

type ShowListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListenerTagsInvoker) Invoke() (*model.ShowListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerTagsResponse), nil
	}
}

type ShowLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadbalancerInvoker) Invoke() (*model.ShowLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadbalancerResponse), nil
	}
}

type ShowLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadbalancerTagsInvoker) Invoke() (*model.ShowLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadbalancerTagsResponse), nil
	}
}

type ShowLoadbalancersStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadbalancersStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadbalancersStatusInvoker) Invoke() (*model.ShowLoadbalancersStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadbalancersStatusResponse), nil
	}
}

type ShowMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMemberInvoker) Invoke() (*model.ShowMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberResponse), nil
	}
}

type ShowPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolInvoker) Invoke() (*model.ShowPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolResponse), nil
	}
}

type ShowWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWhitelistInvoker) Invoke() (*model.ShowWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWhitelistResponse), nil
	}
}

type UpdateHealthmonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHealthmonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHealthmonitorInvoker) Invoke() (*model.UpdateHealthmonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHealthmonitorResponse), nil
	}
}

type UpdateL7policiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateL7policiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateL7policiesInvoker) Invoke() (*model.UpdateL7policiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateL7policiesResponse), nil
	}
}

type UpdateL7ruleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateL7ruleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateL7ruleInvoker) Invoke() (*model.UpdateL7ruleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateL7ruleResponse), nil
	}
}

type UpdateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateListenerInvoker) Invoke() (*model.UpdateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateListenerResponse), nil
	}
}

type UpdateLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLoadbalancerInvoker) Invoke() (*model.UpdateLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLoadbalancerResponse), nil
	}
}

type UpdateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMemberInvoker) Invoke() (*model.UpdateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberResponse), nil
	}
}

type UpdatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePoolInvoker) Invoke() (*model.UpdatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePoolResponse), nil
	}
}

type UpdateWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWhitelistInvoker) Invoke() (*model.UpdateWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWhitelistResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}

type UpdateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCertificateInvoker) Invoke() (*model.UpdateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateResponse), nil
	}
}
