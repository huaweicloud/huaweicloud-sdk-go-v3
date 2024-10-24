package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type CreateAccountAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccountAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccountAssignmentInvoker) Invoke() (*model.CreateAccountAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccountAssignmentResponse), nil
	}
}

type DeleteAccountAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccountAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccountAssignmentInvoker) Invoke() (*model.DeleteAccountAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccountAssignmentResponse), nil
	}
}

type DescribeAccountAssignmentCreationStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeAccountAssignmentCreationStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeAccountAssignmentCreationStatusInvoker) Invoke() (*model.DescribeAccountAssignmentCreationStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeAccountAssignmentCreationStatusResponse), nil
	}
}

type DescribeAccountAssignmentDeletionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeAccountAssignmentDeletionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeAccountAssignmentDeletionStatusInvoker) Invoke() (*model.DescribeAccountAssignmentDeletionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeAccountAssignmentDeletionStatusResponse), nil
	}
}

type ListAccountAssignmentCreationStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentCreationStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentCreationStatusInvoker) Invoke() (*model.ListAccountAssignmentCreationStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentCreationStatusResponse), nil
	}
}

type ListAccountAssignmentDeletionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentDeletionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentDeletionStatusInvoker) Invoke() (*model.ListAccountAssignmentDeletionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentDeletionStatusResponse), nil
	}
}

type ListAccountAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAssignmentsInvoker) Invoke() (*model.ListAccountAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type CreateInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.CreateInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type DeleteInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.DeleteInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type DescribeInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.DescribeInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type UpdateInstanceAccessControlAttributeConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceAccessControlAttributeConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceAccessControlAttributeConfigurationInvoker) Invoke() (*model.UpdateInstanceAccessControlAttributeConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

type AttachManagedPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachManagedPolicyToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachManagedPolicyToPermissionSetInvoker) Invoke() (*model.AttachManagedPolicyToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachManagedPolicyToPermissionSetResponse), nil
	}
}

type CreatePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePermissionSetInvoker) Invoke() (*model.CreatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePermissionSetResponse), nil
	}
}

type DeleteCustomPolicyFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomPolicyFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomPolicyFromPermissionSetInvoker) Invoke() (*model.DeleteCustomPolicyFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomPolicyFromPermissionSetResponse), nil
	}
}

type DeleteCustomRoleFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomRoleFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomRoleFromPermissionSetInvoker) Invoke() (*model.DeleteCustomRoleFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomRoleFromPermissionSetResponse), nil
	}
}

type DeletePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePermissionSetInvoker) Invoke() (*model.DeletePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePermissionSetResponse), nil
	}
}

type DescribePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribePermissionSetInvoker) Invoke() (*model.DescribePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePermissionSetResponse), nil
	}
}

type DescribePermissionSetProvisioningStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribePermissionSetProvisioningStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribePermissionSetProvisioningStatusInvoker) Invoke() (*model.DescribePermissionSetProvisioningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePermissionSetProvisioningStatusResponse), nil
	}
}

type DetachManagedPolicyFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachManagedPolicyFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachManagedPolicyFromPermissionSetInvoker) Invoke() (*model.DetachManagedPolicyFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachManagedPolicyFromPermissionSetResponse), nil
	}
}

type GetCustomPolicyForPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetCustomPolicyForPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetCustomPolicyForPermissionSetInvoker) Invoke() (*model.GetCustomPolicyForPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetCustomPolicyForPermissionSetResponse), nil
	}
}

type GetCustomRoleForPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetCustomRoleForPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetCustomRoleForPermissionSetInvoker) Invoke() (*model.GetCustomRoleForPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetCustomRoleForPermissionSetResponse), nil
	}
}

type ListAccountsForProvisionedPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsForProvisionedPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsForProvisionedPermissionSetInvoker) Invoke() (*model.ListAccountsForProvisionedPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsForProvisionedPermissionSetResponse), nil
	}
}

type ListManagedPoliciesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedPoliciesInPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedPoliciesInPermissionSetInvoker) Invoke() (*model.ListManagedPoliciesInPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedPoliciesInPermissionSetResponse), nil
	}
}

type ListPermissionSetProvisioningStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetProvisioningStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetProvisioningStatusInvoker) Invoke() (*model.ListPermissionSetProvisioningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetProvisioningStatusResponse), nil
	}
}

type ListPermissionSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetsInvoker) Invoke() (*model.ListPermissionSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetsResponse), nil
	}
}

type ListPermissionSetsProvisionedToAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionSetsProvisionedToAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionSetsProvisionedToAccountInvoker) Invoke() (*model.ListPermissionSetsProvisionedToAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetsProvisionedToAccountResponse), nil
	}
}

type ProvisionPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ProvisionPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ProvisionPermissionSetInvoker) Invoke() (*model.ProvisionPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProvisionPermissionSetResponse), nil
	}
}

type PutCustomPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutCustomPolicyToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutCustomPolicyToPermissionSetInvoker) Invoke() (*model.PutCustomPolicyToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutCustomPolicyToPermissionSetResponse), nil
	}
}

type PutCustomRoleToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutCustomRoleToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutCustomRoleToPermissionSetInvoker) Invoke() (*model.PutCustomRoleToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutCustomRoleToPermissionSetResponse), nil
	}
}

type UpdatePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePermissionSetInvoker) Invoke() (*model.UpdatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePermissionSetResponse), nil
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

type AttachManagedRoleToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachManagedRoleToPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachManagedRoleToPermissionSetInvoker) Invoke() (*model.AttachManagedRoleToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachManagedRoleToPermissionSetResponse), nil
	}
}

type DetachManagedRoleFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachManagedRoleFromPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachManagedRoleFromPermissionSetInvoker) Invoke() (*model.DetachManagedRoleFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachManagedRoleFromPermissionSetResponse), nil
	}
}

type ListManagedRolesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedRolesInPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedRolesInPermissionSetInvoker) Invoke() (*model.ListManagedRolesInPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedRolesInPermissionSetResponse), nil
	}
}
