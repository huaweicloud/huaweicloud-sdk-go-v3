package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type CreateAccountAssignmentInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAccountAssignmentsInvoker) Invoke() (*model.ListAccountAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAssignmentsResponse), nil
	}
}

type ListProfileAssociationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProfileAssociationsInvoker) Invoke() (*model.ListProfileAssociationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProfileAssociationsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type AttachCustomerManagedPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachCustomerManagedPolicyToPermissionSetInvoker) Invoke() (*model.AttachCustomerManagedPolicyToPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachCustomerManagedPolicyToPermissionSetResponse), nil
	}
}

type AttachManagedPolicyToPermissionSetInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreatePermissionSetInvoker) Invoke() (*model.CreatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePermissionSetResponse), nil
	}
}

type DeletePermissionSetInvoker struct {
	*invoker.BaseInvoker
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

func (i *DescribePermissionSetProvisioningStatusInvoker) Invoke() (*model.DescribePermissionSetProvisioningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePermissionSetProvisioningStatusResponse), nil
	}
}

type DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker) Invoke() (*model.DetachCustomerManagedPolicyReferenceFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachCustomerManagedPolicyReferenceFromPermissionSetResponse), nil
	}
}

type DetachManagedPolicyFromPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachManagedPolicyFromPermissionSetInvoker) Invoke() (*model.DetachManagedPolicyFromPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachManagedPolicyFromPermissionSetResponse), nil
	}
}

type ListAccountsForProvisionedPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsForProvisionedPermissionSetInvoker) Invoke() (*model.ListAccountsForProvisionedPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsForProvisionedPermissionSetResponse), nil
	}
}

type ListCustomerManagedPolicyReferencesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerManagedPolicyReferencesInPermissionSetInvoker) Invoke() (*model.ListCustomerManagedPolicyReferencesInPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerManagedPolicyReferencesInPermissionSetResponse), nil
	}
}

type ListManagedPoliciesInPermissionSetInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPermissionSetsProvisionedToAccountInvoker) Invoke() (*model.ListPermissionSetsProvisionedToAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionSetsProvisionedToAccountResponse), nil
	}
}

type UpdatePermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePermissionSetInvoker) Invoke() (*model.UpdatePermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePermissionSetResponse), nil
	}
}
