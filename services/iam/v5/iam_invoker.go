package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v5/model"
)

type GetAccountSummaryV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetAccountSummaryV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAccountSummaryV5Invoker) Invoke() (*model.GetAccountSummaryV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAccountSummaryV5Response), nil
	}
}

type GetAsymmetricSignatureSwitchV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetAsymmetricSignatureSwitchV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAsymmetricSignatureSwitchV5Invoker) Invoke() (*model.GetAsymmetricSignatureSwitchV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAsymmetricSignatureSwitchV5Response), nil
	}
}

type GetFeatureStatusV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetFeatureStatusV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetFeatureStatusV5Invoker) Invoke() (*model.GetFeatureStatusV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetFeatureStatusV5Response), nil
	}
}

type SetAsymmetricSignatureSwitchV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *SetAsymmetricSignatureSwitchV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetAsymmetricSignatureSwitchV5Invoker) Invoke() (*model.SetAsymmetricSignatureSwitchV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAsymmetricSignatureSwitchV5Response), nil
	}
}

type CreateAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyV5Invoker) Invoke() (*model.CreateAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyV5Response), nil
	}
}

type CreateServiceLinkedAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServiceLinkedAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServiceLinkedAgencyV5Invoker) Invoke() (*model.CreateServiceLinkedAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServiceLinkedAgencyV5Response), nil
	}
}

type DeleteAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgencyV5Invoker) Invoke() (*model.DeleteAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgencyV5Response), nil
	}
}

type DeleteServiceLinkedAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceLinkedAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceLinkedAgencyV5Invoker) Invoke() (*model.DeleteServiceLinkedAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceLinkedAgencyV5Response), nil
	}
}

type GetAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAgencyV5Invoker) Invoke() (*model.GetAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAgencyV5Response), nil
	}
}

type GetServiceLinkedAgencyDeletionStatusV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetServiceLinkedAgencyDeletionStatusV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetServiceLinkedAgencyDeletionStatusV5Invoker) Invoke() (*model.GetServiceLinkedAgencyDeletionStatusV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetServiceLinkedAgencyDeletionStatusV5Response), nil
	}
}

type ListAgenciesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgenciesV5Invoker) Invoke() (*model.ListAgenciesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesV5Response), nil
	}
}

type UpdateAgencyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgencyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAgencyV5Invoker) Invoke() (*model.UpdateAgencyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgencyV5Response), nil
	}
}

type UpdateTrustPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrustPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTrustPolicyV5Invoker) Invoke() (*model.UpdateTrustPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrustPolicyV5Response), nil
	}
}

type GetAuthorizationSchemaV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetAuthorizationSchemaV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAuthorizationSchemaV5Invoker) Invoke() (*model.GetAuthorizationSchemaV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAuthorizationSchemaV5Response), nil
	}
}

type ListRegisteredServicesForAuthSchemaV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegisteredServicesForAuthSchemaV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegisteredServicesForAuthSchemaV5Invoker) Invoke() (*model.ListRegisteredServicesForAuthSchemaV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegisteredServicesForAuthSchemaV5Response), nil
	}
}

type ListServicePrincipalsV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicePrincipalsV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicePrincipalsV5Invoker) Invoke() (*model.ListServicePrincipalsV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicePrincipalsV5Response), nil
	}
}

type AddUserToGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddUserToGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddUserToGroupV5Invoker) Invoke() (*model.AddUserToGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddUserToGroupV5Response), nil
	}
}

type CreateGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupV5Invoker) Invoke() (*model.CreateGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupV5Response), nil
	}
}

type DeleteGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupV5Invoker) Invoke() (*model.DeleteGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupV5Response), nil
	}
}

type ListGroupsV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsV5Invoker) Invoke() (*model.ListGroupsV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsV5Response), nil
	}
}

type RemoveUserFromGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveUserFromGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveUserFromGroupV5Invoker) Invoke() (*model.RemoveUserFromGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveUserFromGroupV5Response), nil
	}
}

type ShowGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupV5Invoker) Invoke() (*model.ShowGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupV5Response), nil
	}
}

type UpdateGroupV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupV5Invoker) Invoke() (*model.UpdateGroupV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupV5Response), nil
	}
}

type CreateVirtualMfaDeviceV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVirtualMfaDeviceV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVirtualMfaDeviceV5Invoker) Invoke() (*model.CreateVirtualMfaDeviceV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVirtualMfaDeviceV5Response), nil
	}
}

type DeleteVirtualMfaDeviceV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVirtualMfaDeviceV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVirtualMfaDeviceV5Invoker) Invoke() (*model.DeleteVirtualMfaDeviceV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVirtualMfaDeviceV5Response), nil
	}
}

type DisableMfaDeviceV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisableMfaDeviceV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableMfaDeviceV5Invoker) Invoke() (*model.DisableMfaDeviceV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableMfaDeviceV5Response), nil
	}
}

type EnableMfaDeviceV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *EnableMfaDeviceV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableMfaDeviceV5Invoker) Invoke() (*model.EnableMfaDeviceV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableMfaDeviceV5Response), nil
	}
}

type ListMfaDevicesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListMfaDevicesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMfaDevicesV5Invoker) Invoke() (*model.ListMfaDevicesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMfaDevicesV5Response), nil
	}
}

type CreatePolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyV5Invoker) Invoke() (*model.CreatePolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyV5Response), nil
	}
}

type CreatePolicyVersionV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyVersionV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyVersionV5Invoker) Invoke() (*model.CreatePolicyVersionV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyVersionV5Response), nil
	}
}

type DeletePolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyV5Invoker) Invoke() (*model.DeletePolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyV5Response), nil
	}
}

type DeletePolicyVersionV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyVersionV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyVersionV5Invoker) Invoke() (*model.DeletePolicyVersionV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyVersionV5Response), nil
	}
}

type GetPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetPolicyV5Invoker) Invoke() (*model.GetPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPolicyV5Response), nil
	}
}

type GetPolicyVersionV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *GetPolicyVersionV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetPolicyVersionV5Invoker) Invoke() (*model.GetPolicyVersionV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetPolicyVersionV5Response), nil
	}
}

type ListPoliciesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesV5Invoker) Invoke() (*model.ListPoliciesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesV5Response), nil
	}
}

type ListPolicyVersionsV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyVersionsV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyVersionsV5Invoker) Invoke() (*model.ListPolicyVersionsV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyVersionsV5Response), nil
	}
}

type SetDefaultPolicyVersionV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *SetDefaultPolicyVersionV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetDefaultPolicyVersionV5Invoker) Invoke() (*model.SetDefaultPolicyVersionV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetDefaultPolicyVersionV5Response), nil
	}
}

type AttachAgencyPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *AttachAgencyPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachAgencyPolicyV5Invoker) Invoke() (*model.AttachAgencyPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachAgencyPolicyV5Response), nil
	}
}

type AttachGroupPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *AttachGroupPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachGroupPolicyV5Invoker) Invoke() (*model.AttachGroupPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachGroupPolicyV5Response), nil
	}
}

type AttachUserPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *AttachUserPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachUserPolicyV5Invoker) Invoke() (*model.AttachUserPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachUserPolicyV5Response), nil
	}
}

type DetachAgencyPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetachAgencyPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachAgencyPolicyV5Invoker) Invoke() (*model.DetachAgencyPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachAgencyPolicyV5Response), nil
	}
}

type DetachGroupPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetachGroupPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachGroupPolicyV5Invoker) Invoke() (*model.DetachGroupPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachGroupPolicyV5Response), nil
	}
}

type DetachUserPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetachUserPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachUserPolicyV5Invoker) Invoke() (*model.DetachUserPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachUserPolicyV5Response), nil
	}
}

type ListAttachedAgencyPoliciesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachedAgencyPoliciesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttachedAgencyPoliciesV5Invoker) Invoke() (*model.ListAttachedAgencyPoliciesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachedAgencyPoliciesV5Response), nil
	}
}

type ListAttachedGroupPoliciesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachedGroupPoliciesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttachedGroupPoliciesV5Invoker) Invoke() (*model.ListAttachedGroupPoliciesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachedGroupPoliciesV5Response), nil
	}
}

type ListAttachedUserPoliciesV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachedUserPoliciesV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttachedUserPoliciesV5Invoker) Invoke() (*model.ListAttachedUserPoliciesV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachedUserPoliciesV5Response), nil
	}
}

type ListEntitiesForPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListEntitiesForPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEntitiesForPolicyV5Invoker) Invoke() (*model.ListEntitiesForPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEntitiesForPolicyV5Response), nil
	}
}

type DeleteResourceTagsV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagsV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceTagsV5Invoker) Invoke() (*model.DeleteResourceTagsV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagsV5Response), nil
	}
}

type ListResourceTagsV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagsV5Invoker) Invoke() (*model.ListResourceTagsV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsV5Response), nil
	}
}

type TagResourceV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *TagResourceV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagResourceV5Invoker) Invoke() (*model.TagResourceV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagResourceV5Response), nil
	}
}

type ShowLoginPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoginPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoginPolicyV5Invoker) Invoke() (*model.ShowLoginPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoginPolicyV5Response), nil
	}
}

type ShowPasswordPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPasswordPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPasswordPolicyV5Invoker) Invoke() (*model.ShowPasswordPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPasswordPolicyV5Response), nil
	}
}

type ShowTokenPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTokenPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTokenPolicyV5Invoker) Invoke() (*model.ShowTokenPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTokenPolicyV5Response), nil
	}
}

type UpdateLoginPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLoginPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLoginPolicyV5Invoker) Invoke() (*model.UpdateLoginPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLoginPolicyV5Response), nil
	}
}

type UpdatePasswordPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePasswordPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePasswordPolicyV5Invoker) Invoke() (*model.UpdatePasswordPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePasswordPolicyV5Response), nil
	}
}

type UpdateTokenPolicyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTokenPolicyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTokenPolicyV5Invoker) Invoke() (*model.UpdateTokenPolicyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTokenPolicyV5Response), nil
	}
}

type CreateUserV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserV5Invoker) Invoke() (*model.CreateUserV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserV5Response), nil
	}
}

type DeleteUserV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserV5Invoker) Invoke() (*model.DeleteUserV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserV5Response), nil
	}
}

type ListUsersV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersV5Invoker) Invoke() (*model.ListUsersV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersV5Response), nil
	}
}

type ShowUserLastLoginV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserLastLoginV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserLastLoginV5Invoker) Invoke() (*model.ShowUserLastLoginV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserLastLoginV5Response), nil
	}
}

type ShowUserV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserV5Invoker) Invoke() (*model.ShowUserV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserV5Response), nil
	}
}

type UpdateUserV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserV5Invoker) Invoke() (*model.UpdateUserV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserV5Response), nil
	}
}

type ChangePasswordV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePasswordV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePasswordV5Invoker) Invoke() (*model.ChangePasswordV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePasswordV5Response), nil
	}
}

type CreateAccessKeyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessKeyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccessKeyV5Invoker) Invoke() (*model.CreateAccessKeyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessKeyV5Response), nil
	}
}

type CreateLoginProfileV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLoginProfileV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLoginProfileV5Invoker) Invoke() (*model.CreateLoginProfileV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLoginProfileV5Response), nil
	}
}

type DeleteAccessKeyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccessKeyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccessKeyV5Invoker) Invoke() (*model.DeleteAccessKeyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccessKeyV5Response), nil
	}
}

type DeleteLoginProfileV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoginProfileV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLoginProfileV5Invoker) Invoke() (*model.DeleteLoginProfileV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoginProfileV5Response), nil
	}
}

type ListAccessKeysV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessKeysV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessKeysV5Invoker) Invoke() (*model.ListAccessKeysV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessKeysV5Response), nil
	}
}

type ShowAccessKeyLastUsedV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessKeyLastUsedV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessKeyLastUsedV5Invoker) Invoke() (*model.ShowAccessKeyLastUsedV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessKeyLastUsedV5Response), nil
	}
}

type ShowLoginProfileV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoginProfileV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoginProfileV5Invoker) Invoke() (*model.ShowLoginProfileV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoginProfileV5Response), nil
	}
}

type UpdateAccessKeyV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessKeyV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessKeyV5Invoker) Invoke() (*model.UpdateAccessKeyV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessKeyV5Response), nil
	}
}

type UpdateLoginProfileV5Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLoginProfileV5Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLoginProfileV5Invoker) Invoke() (*model.UpdateLoginProfileV5Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLoginProfileV5Response), nil
	}
}
