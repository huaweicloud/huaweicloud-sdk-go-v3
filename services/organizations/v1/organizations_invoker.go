package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/organizations/v1/model"
)

type CloseAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseAccountInvoker) Invoke() (*model.CloseAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseAccountResponse), nil
	}
}

type CreateAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccountInvoker) Invoke() (*model.CreateAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccountResponse), nil
	}
}

type InviteAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InviteAccountInvoker) Invoke() (*model.InviteAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteAccountResponse), nil
	}
}

type ListAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsInvoker) Invoke() (*model.ListAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsResponse), nil
	}
}

type ListCloseAccountStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloseAccountStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloseAccountStatusesInvoker) Invoke() (*model.ListCloseAccountStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloseAccountStatusesResponse), nil
	}
}

type ListCreateAccountStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCreateAccountStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCreateAccountStatusesInvoker) Invoke() (*model.ListCreateAccountStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCreateAccountStatusesResponse), nil
	}
}

type MoveAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *MoveAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MoveAccountInvoker) Invoke() (*model.MoveAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MoveAccountResponse), nil
	}
}

type RemoveAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveAccountInvoker) Invoke() (*model.RemoveAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveAccountResponse), nil
	}
}

type ShowAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccountInvoker) Invoke() (*model.ShowAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccountResponse), nil
	}
}

type ShowCreateAccountStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCreateAccountStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCreateAccountStatusInvoker) Invoke() (*model.ShowCreateAccountStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCreateAccountStatusResponse), nil
	}
}

type UpdateAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccountInvoker) Invoke() (*model.UpdateAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccountResponse), nil
	}
}

type DeregisterDelegatedAdministratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeregisterDelegatedAdministratorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeregisterDelegatedAdministratorInvoker) Invoke() (*model.DeregisterDelegatedAdministratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeregisterDelegatedAdministratorResponse), nil
	}
}

type ListDelegatedAdministratorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDelegatedAdministratorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDelegatedAdministratorsInvoker) Invoke() (*model.ListDelegatedAdministratorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDelegatedAdministratorsResponse), nil
	}
}

type ListDelegatedServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDelegatedServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDelegatedServicesInvoker) Invoke() (*model.ListDelegatedServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDelegatedServicesResponse), nil
	}
}

type RegisterDelegatedAdministratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterDelegatedAdministratorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterDelegatedAdministratorInvoker) Invoke() (*model.RegisterDelegatedAdministratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterDelegatedAdministratorResponse), nil
	}
}

type AcceptHandshakeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptHandshakeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptHandshakeInvoker) Invoke() (*model.AcceptHandshakeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptHandshakeResponse), nil
	}
}

type CancelHandshakeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelHandshakeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelHandshakeInvoker) Invoke() (*model.CancelHandshakeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelHandshakeResponse), nil
	}
}

type DeclineHandshakeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeclineHandshakeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeclineHandshakeInvoker) Invoke() (*model.DeclineHandshakeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeclineHandshakeResponse), nil
	}
}

type ListHandshakesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHandshakesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHandshakesInvoker) Invoke() (*model.ListHandshakesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHandshakesResponse), nil
	}
}

type ListReceivedHandshakesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReceivedHandshakesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReceivedHandshakesInvoker) Invoke() (*model.ListReceivedHandshakesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReceivedHandshakesResponse), nil
	}
}

type ShowHandshakeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHandshakeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHandshakeInvoker) Invoke() (*model.ShowHandshakeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHandshakeResponse), nil
	}
}

type ListEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEntitiesInvoker) Invoke() (*model.ListEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEntitiesResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicesInvoker) Invoke() (*model.ListServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicesResponse), nil
	}
}

type ListTagPolicyServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagPolicyServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagPolicyServicesInvoker) Invoke() (*model.ListTagPolicyServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagPolicyServicesResponse), nil
	}
}

type ShowEffectivePoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEffectivePoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEffectivePoliciesInvoker) Invoke() (*model.ShowEffectivePoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEffectivePoliciesResponse), nil
	}
}

type CreateOrganizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrganizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrganizationInvoker) Invoke() (*model.CreateOrganizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrganizationResponse), nil
	}
}

type DeleteOrganizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOrganizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOrganizationInvoker) Invoke() (*model.DeleteOrganizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOrganizationResponse), nil
	}
}

type LeaveOrganizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *LeaveOrganizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LeaveOrganizationInvoker) Invoke() (*model.LeaveOrganizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveOrganizationResponse), nil
	}
}

type ListRootsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRootsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRootsInvoker) Invoke() (*model.ListRootsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRootsResponse), nil
	}
}

type ShowOrganizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationInvoker) Invoke() (*model.ShowOrganizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationResponse), nil
	}
}

type CreateOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrganizationalUnitInvoker) Invoke() (*model.CreateOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrganizationalUnitResponse), nil
	}
}

type DeleteOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOrganizationalUnitInvoker) Invoke() (*model.DeleteOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOrganizationalUnitResponse), nil
	}
}

type ListOrganizationalUnitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationalUnitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationalUnitsInvoker) Invoke() (*model.ListOrganizationalUnitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationalUnitsResponse), nil
	}
}

type ShowOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationalUnitInvoker) Invoke() (*model.ShowOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationalUnitResponse), nil
	}
}

type UpdateOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOrganizationalUnitInvoker) Invoke() (*model.UpdateOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOrganizationalUnitResponse), nil
	}
}

type AttachPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachPolicyInvoker) Invoke() (*model.AttachPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachPolicyResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type DetachPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachPolicyInvoker) Invoke() (*model.DetachPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachPolicyResponse), nil
	}
}

type DisablePolicyTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisablePolicyTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisablePolicyTypeInvoker) Invoke() (*model.DisablePolicyTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisablePolicyTypeResponse), nil
	}
}

type EnablePolicyTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnablePolicyTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnablePolicyTypeInvoker) Invoke() (*model.EnablePolicyTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnablePolicyTypeResponse), nil
	}
}

type ListEntitiesForPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEntitiesForPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEntitiesForPolicyInvoker) Invoke() (*model.ListEntitiesForPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEntitiesForPolicyResponse), nil
	}
}

type ListPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesInvoker) Invoke() (*model.ListPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesResponse), nil
	}
}

type ShowPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyInvoker) Invoke() (*model.ShowPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyResponse), nil
	}
}

type UpdatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyInvoker) Invoke() (*model.UpdatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyResponse), nil
	}
}

type CreateTagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagResourceInvoker) Invoke() (*model.CreateTagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResourceResponse), nil
	}
}

type DeleteTagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagResourceInvoker) Invoke() (*model.DeleteTagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResourceResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstancesInvoker) Invoke() (*model.ListResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstancesResponse), nil
	}
}

type ListResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type ListTagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagResourcesInvoker) Invoke() (*model.ListTagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagResourcesResponse), nil
	}
}

type ListTagsForResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsForResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsForResourceInvoker) Invoke() (*model.ListTagsForResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsForResourceResponse), nil
	}
}

type ShowResourceInstancesCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceInstancesCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceInstancesCountInvoker) Invoke() (*model.ShowResourceInstancesCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceInstancesCountResponse), nil
	}
}

type TagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagResourceInvoker) Invoke() (*model.TagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagResourceResponse), nil
	}
}

type UntagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UntagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UntagResourceInvoker) Invoke() (*model.UntagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UntagResourceResponse), nil
	}
}

type DisableTrustedServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableTrustedServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableTrustedServiceInvoker) Invoke() (*model.DisableTrustedServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableTrustedServiceResponse), nil
	}
}

type EnableTrustedServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableTrustedServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableTrustedServiceInvoker) Invoke() (*model.EnableTrustedServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableTrustedServiceResponse), nil
	}
}

type ListTrustedServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrustedServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrustedServicesInvoker) Invoke() (*model.ListTrustedServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrustedServicesResponse), nil
	}
}
