package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterstore/v1/model"
)

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupInvoker) Invoke() (*model.CreateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupResponse), nil
	}
}

type DeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type DescribeGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeGroupInvoker) Invoke() (*model.DescribeGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeGroupResponse), nil
	}
}

type DescribeGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeGroupsInvoker) Invoke() (*model.DescribeGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeGroupsResponse), nil
	}
}

type GetGroupIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetGroupIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetGroupIdInvoker) Invoke() (*model.GetGroupIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetGroupIdResponse), nil
	}
}

type ListGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsInvoker) Invoke() (*model.ListGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsResponse), nil
	}
}

type UpdateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGroupInvoker) Invoke() (*model.UpdateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupResponse), nil
	}
}

type CreateGroupMembershipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupMembershipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGroupMembershipInvoker) Invoke() (*model.CreateGroupMembershipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupMembershipResponse), nil
	}
}

type DeleteGroupMembershipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupMembershipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupMembershipInvoker) Invoke() (*model.DeleteGroupMembershipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupMembershipResponse), nil
	}
}

type DescribeGroupMembershipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeGroupMembershipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeGroupMembershipInvoker) Invoke() (*model.DescribeGroupMembershipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeGroupMembershipResponse), nil
	}
}

type GetGroupMembershipIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetGroupMembershipIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetGroupMembershipIdInvoker) Invoke() (*model.GetGroupMembershipIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetGroupMembershipIdResponse), nil
	}
}

type IsMemberInGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *IsMemberInGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *IsMemberInGroupsInvoker) Invoke() (*model.IsMemberInGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.IsMemberInGroupsResponse), nil
	}
}

type ListGroupMembershipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMembershipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMembershipsInvoker) Invoke() (*model.ListGroupMembershipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMembershipsResponse), nil
	}
}

type ListGroupMembershipsForMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupMembershipsForMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupMembershipsForMemberInvoker) Invoke() (*model.ListGroupMembershipsForMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupMembershipsForMemberResponse), nil
	}
}

type CreateExternalIdPConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExternalIdPConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExternalIdPConfigurationForDirectoryInvoker) Invoke() (*model.CreateExternalIdPConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExternalIdPConfigurationForDirectoryResponse), nil
	}
}

type DeleteExternalIdPCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExternalIdPCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteExternalIdPCertificateInvoker) Invoke() (*model.DeleteExternalIdPCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExternalIdPCertificateResponse), nil
	}
}

type DeleteExternalIdPConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExternalIdPConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteExternalIdPConfigurationForDirectoryInvoker) Invoke() (*model.DeleteExternalIdPConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExternalIdPConfigurationForDirectoryResponse), nil
	}
}

type DisableExternalIdPConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableExternalIdPConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableExternalIdPConfigurationForDirectoryInvoker) Invoke() (*model.DisableExternalIdPConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableExternalIdPConfigurationForDirectoryResponse), nil
	}
}

type EnableExternalIdPConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableExternalIdPConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableExternalIdPConfigurationForDirectoryInvoker) Invoke() (*model.EnableExternalIdPConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableExternalIdPConfigurationForDirectoryResponse), nil
	}
}

type ImportExternalIdPCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportExternalIdPCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportExternalIdPCertificateInvoker) Invoke() (*model.ImportExternalIdPCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportExternalIdPCertificateResponse), nil
	}
}

type ListExternalIdPCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExternalIdPCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExternalIdPCertificatesInvoker) Invoke() (*model.ListExternalIdPCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExternalIdPCertificatesResponse), nil
	}
}

type ListExternalIdPConfigurationsForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExternalIdPConfigurationsForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExternalIdPConfigurationsForDirectoryInvoker) Invoke() (*model.ListExternalIdPConfigurationsForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExternalIdPConfigurationsForDirectoryResponse), nil
	}
}

type UpdateExternalIdPConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateExternalIdPConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateExternalIdPConfigurationForDirectoryInvoker) Invoke() (*model.UpdateExternalIdPConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateExternalIdPConfigurationForDirectoryResponse), nil
	}
}

type DescribePasswordPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribePasswordPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribePasswordPolicyInvoker) Invoke() (*model.DescribePasswordPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribePasswordPolicyResponse), nil
	}
}

type UpdatePasswordPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePasswordPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePasswordPolicyInvoker) Invoke() (*model.UpdatePasswordPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePasswordPolicyResponse), nil
	}
}

type CreateSpCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSpCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSpCertificateInvoker) Invoke() (*model.CreateSpCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSpCertificateResponse), nil
	}
}

type DeleteSpCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSpCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSpCertificateInvoker) Invoke() (*model.DeleteSpCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSpCertificateResponse), nil
	}
}

type GetSpConfigurationForDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetSpConfigurationForDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetSpConfigurationForDirectoryInvoker) Invoke() (*model.GetSpConfigurationForDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetSpConfigurationForDirectoryResponse), nil
	}
}

type ListSpCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpCertificatesInvoker) Invoke() (*model.ListSpCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpCertificatesResponse), nil
	}
}

type UpdateSpActiveCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSpActiveCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSpActiveCertificateInvoker) Invoke() (*model.UpdateSpActiveCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSpActiveCertificateResponse), nil
	}
}

type GetIdentityStoreSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetIdentityStoreSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetIdentityStoreSummaryInvoker) Invoke() (*model.GetIdentityStoreSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetIdentityStoreSummaryResponse), nil
	}
}

type CreateBearerTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBearerTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBearerTokenInvoker) Invoke() (*model.CreateBearerTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBearerTokenResponse), nil
	}
}

type CreateProvisioningTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProvisioningTenantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProvisioningTenantInvoker) Invoke() (*model.CreateProvisioningTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProvisioningTenantResponse), nil
	}
}

type DeleteBearerTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBearerTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBearerTokenInvoker) Invoke() (*model.DeleteBearerTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBearerTokenResponse), nil
	}
}

type DeleteProvisioningTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProvisioningTenantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProvisioningTenantInvoker) Invoke() (*model.DeleteProvisioningTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProvisioningTenantResponse), nil
	}
}

type ListBearerTokensInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBearerTokensInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBearerTokensInvoker) Invoke() (*model.ListBearerTokensResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBearerTokensResponse), nil
	}
}

type ListProvisioningTenantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProvisioningTenantsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProvisioningTenantsInvoker) Invoke() (*model.ListProvisioningTenantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProvisioningTenantsResponse), nil
	}
}

type BatchDeleteSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSessionsInvoker) Invoke() (*model.BatchDeleteSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSessionsResponse), nil
	}
}

type BatchListMfaDevicesForUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListMfaDevicesForUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchListMfaDevicesForUserInvoker) Invoke() (*model.BatchListMfaDevicesForUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListMfaDevicesForUserResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type DeleteMfaDeviceForUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMfaDeviceForUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMfaDeviceForUserInvoker) Invoke() (*model.DeleteMfaDeviceForUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMfaDeviceForUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type DescribeUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeUserInvoker) Invoke() (*model.DescribeUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeUserResponse), nil
	}
}

type DescribeUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeUsersInvoker) Invoke() (*model.DescribeUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeUsersResponse), nil
	}
}

type DisableUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableUserInvoker) Invoke() (*model.DisableUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableUserResponse), nil
	}
}

type EnableUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableUserInvoker) Invoke() (*model.EnableUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableUserResponse), nil
	}
}

type GetUserIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetUserIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetUserIdInvoker) Invoke() (*model.GetUserIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetUserIdResponse), nil
	}
}

type ListSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSessionsInvoker) Invoke() (*model.ListSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionsResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type RegisterMfaDeviceForUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterMfaDeviceForUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterMfaDeviceForUserInvoker) Invoke() (*model.RegisterMfaDeviceForUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterMfaDeviceForUserResponse), nil
	}
}

type ResetPwdModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetPwdModeInvoker) Invoke() (*model.ResetPwdModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPwdModeResponse), nil
	}
}

type UpdateMfaDeviceForUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMfaDeviceForUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMfaDeviceForUserInvoker) Invoke() (*model.UpdateMfaDeviceForUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMfaDeviceForUserResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type VerifyEmailInvoker struct {
	*invoker.BaseInvoker
}

func (i *VerifyEmailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *VerifyEmailInvoker) Invoke() (*model.VerifyEmailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerifyEmailResponse), nil
	}
}
