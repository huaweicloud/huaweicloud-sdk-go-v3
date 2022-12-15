package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

type BatchDeleteAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAccessPoliciesInvoker) Invoke() (*model.BatchDeleteAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAccessPoliciesResponse), nil
	}
}

type CreateAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessPolicyInvoker) Invoke() (*model.CreateAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessPolicyResponse), nil
	}
}

type ListAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPoliciesInvoker) Invoke() (*model.ListAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPoliciesResponse), nil
	}
}

type ListAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPolicyObjectsInvoker) Invoke() (*model.ListAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPolicyObjectsResponse), nil
	}
}

type UpdateAccessPolicyObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessPolicyObjectsInvoker) Invoke() (*model.UpdateAccessPolicyObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessPolicyObjectsResponse), nil
	}
}

type ShowAssistAuthConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssistAuthConfigInvoker) Invoke() (*model.ShowAssistAuthConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssistAuthConfigResponse), nil
	}
}

type UpdateAssistAuthMethodConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssistAuthMethodConfigInvoker) Invoke() (*model.UpdateAssistAuthMethodConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssistAuthMethodConfigResponse), nil
	}
}

type ListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZonesInvoker) Invoke() (*model.ListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZonesResponse), nil
	}
}

type ExportUserLoginInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportUserLoginInfoNewInvoker) Invoke() (*model.ExportUserLoginInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportUserLoginInfoNewResponse), nil
	}
}

type ListHistoryOnlineInfoNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryOnlineInfoNewInvoker) Invoke() (*model.ListHistoryOnlineInfoNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryOnlineInfoNewResponse), nil
	}
}

type ListLoginRecordsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginRecordsNewInvoker) Invoke() (*model.ListLoginRecordsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginRecordsNewResponse), nil
	}
}

type BatchDeleteDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDesktopsInvoker) Invoke() (*model.BatchDeleteDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDesktopsResponse), nil
	}
}

type BatchRunDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRunDesktopsInvoker) Invoke() (*model.BatchRunDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRunDesktopsResponse), nil
	}
}

type CreateDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopInvoker) Invoke() (*model.CreateDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopResponse), nil
	}
}

type DeleteDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopInvoker) Invoke() (*model.DeleteDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopResponse), nil
	}
}

type ListDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsInvoker) Invoke() (*model.ListDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsResponse), nil
	}
}

type ListDesktopsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDesktopsDetailInvoker) Invoke() (*model.ListDesktopsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDesktopsDetailResponse), nil
	}
}

type ResizeDesktopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeDesktopInvoker) Invoke() (*model.ResizeDesktopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeDesktopResponse), nil
	}
}

type ShowDesktopDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDesktopDetailInvoker) Invoke() (*model.ShowDesktopDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDesktopDetailResponse), nil
	}
}

type ListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesInvoker) Invoke() (*model.ListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesResponse), nil
	}
}

type ListItaSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListItaSubJobsInvoker) Invoke() (*model.ListItaSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListItaSubJobsResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type CreateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTerminalsBindingDesktopsInvoker) Invoke() (*model.CreateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTerminalsBindingDesktopsResponse), nil
	}
}

type DeleteTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTerminalsBindingDesktopsInvoker) Invoke() (*model.DeleteTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTerminalsBindingDesktopsResponse), nil
	}
}

type ListTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsInvoker) Invoke() (*model.ListTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsResponse), nil
	}
}

type ListTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.ListTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTerminalsBindingDesktopsConfigResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsResponse), nil
	}
}

type UpdateTerminalsBindingDesktopsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTerminalsBindingDesktopsConfigInvoker) Invoke() (*model.UpdateTerminalsBindingDesktopsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTerminalsBindingDesktopsConfigResponse), nil
	}
}

type BatchDeleteOtpDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteOtpDevicesInvoker) Invoke() (*model.BatchDeleteOtpDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteOtpDevicesResponse), nil
	}
}

type ChangeUserStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeUserStatusInvoker) Invoke() (*model.ChangeUserStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeUserStatusResponse), nil
	}
}

type CreateDesktopUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesktopUserInvoker) Invoke() (*model.CreateDesktopUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesktopUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ListOtpDevicesByUserIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOtpDevicesByUserIdInvoker) Invoke() (*model.ListOtpDevicesByUserIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOtpDevicesByUserIdResponse), nil
	}
}

type ListUserDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserDetailInvoker) Invoke() (*model.ListUserDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserDetailResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type UpdateUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInfoInvoker) Invoke() (*model.UpdateUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserInfoResponse), nil
	}
}

type AddVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVolumesInvoker) Invoke() (*model.AddVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVolumesResponse), nil
	}
}

type DeleteDesktopVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesktopVolumesInvoker) Invoke() (*model.DeleteDesktopVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesktopVolumesResponse), nil
	}
}

type ExpandVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandVolumesInvoker) Invoke() (*model.ExpandVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandVolumesResponse), nil
	}
}

type ApplyWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyWorkspaceInvoker) Invoke() (*model.ApplyWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyWorkspaceResponse), nil
	}
}

type CancelWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelWorkspaceInvoker) Invoke() (*model.CancelWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelWorkspaceResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type ShowWorkspaceLockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceLockInvoker) Invoke() (*model.ShowWorkspaceLockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceLockResponse), nil
	}
}

type UnlockWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockWorkspaceInvoker) Invoke() (*model.UnlockWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockWorkspaceResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}
