package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

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
