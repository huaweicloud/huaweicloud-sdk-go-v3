package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cph/v1/model"
)

type AddImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddImageMemberInvoker) Invoke() (*model.AddImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddImageMemberResponse), nil
	}
}

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type BatchExportCloudPhoneDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExportCloudPhoneDataInvoker) Invoke() (*model.BatchExportCloudPhoneDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExportCloudPhoneDataResponse), nil
	}
}

type BatchImportCloudPhoneDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportCloudPhoneDataInvoker) Invoke() (*model.BatchImportCloudPhoneDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportCloudPhoneDataResponse), nil
	}
}

type BatchShowPhoneConnectInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowPhoneConnectInfosInvoker) Invoke() (*model.BatchShowPhoneConnectInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowPhoneConnectInfosResponse), nil
	}
}

type ChangeCloudPhoneServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeCloudPhoneServerInvoker) Invoke() (*model.ChangeCloudPhoneServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeCloudPhoneServerResponse), nil
	}
}

type ChangeCloudPhoneServerModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeCloudPhoneServerModelInvoker) Invoke() (*model.ChangeCloudPhoneServerModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeCloudPhoneServerModelResponse), nil
	}
}

type CreateNet2CloudPhoneServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNet2CloudPhoneServerInvoker) Invoke() (*model.CreateNet2CloudPhoneServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNet2CloudPhoneServerResponse), nil
	}
}

type DeleteCloudPhoneServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudPhoneServerInvoker) Invoke() (*model.DeleteCloudPhoneServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudPhoneServerResponse), nil
	}
}

type DeleteImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageInvoker) Invoke() (*model.DeleteImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageResponse), nil
	}
}

type DeleteImageMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageMemberInvoker) Invoke() (*model.DeleteImageMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageMemberResponse), nil
	}
}

type DeleteShareAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareAppsInvoker) Invoke() (*model.DeleteShareAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareAppsResponse), nil
	}
}

type DeleteShareFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareFilesInvoker) Invoke() (*model.DeleteShareFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareFilesResponse), nil
	}
}

type ImportTrafficInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportTrafficInvoker) Invoke() (*model.ImportTrafficResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportTrafficResponse), nil
	}
}

type ListCloudPhoneImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudPhoneImagesInvoker) Invoke() (*model.ListCloudPhoneImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudPhoneImagesResponse), nil
	}
}

type ListCloudPhoneModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudPhoneModelsInvoker) Invoke() (*model.ListCloudPhoneModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudPhoneModelsResponse), nil
	}
}

type ListCloudPhoneServerModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudPhoneServerModelsInvoker) Invoke() (*model.ListCloudPhoneServerModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudPhoneServerModelsResponse), nil
	}
}

type ListCloudPhoneServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudPhoneServersInvoker) Invoke() (*model.ListCloudPhoneServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudPhoneServersResponse), nil
	}
}

type ListCloudPhonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudPhonesInvoker) Invoke() (*model.ListCloudPhonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudPhonesResponse), nil
	}
}

type ListEncodeServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEncodeServersInvoker) Invoke() (*model.ListEncodeServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEncodeServersResponse), nil
	}
}

type ListImageMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageMembersInvoker) Invoke() (*model.ListImageMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageMembersResponse), nil
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

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type ListShareFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShareFilesInvoker) Invoke() (*model.ListShareFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShareFilesResponse), nil
	}
}

type PushShareAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushShareAppsInvoker) Invoke() (*model.PushShareAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushShareAppsResponse), nil
	}
}

type PushShareFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushShareFilesInvoker) Invoke() (*model.PushShareFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushShareFilesResponse), nil
	}
}

type ResetCloudPhoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetCloudPhoneInvoker) Invoke() (*model.ResetCloudPhoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetCloudPhoneResponse), nil
	}
}

type RestartCloudPhoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartCloudPhoneInvoker) Invoke() (*model.RestartCloudPhoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartCloudPhoneResponse), nil
	}
}

type RestartCloudPhoneServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartCloudPhoneServerInvoker) Invoke() (*model.RestartCloudPhoneServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartCloudPhoneServerResponse), nil
	}
}

type RestartEncodeServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartEncodeServerInvoker) Invoke() (*model.RestartEncodeServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartEncodeServerResponse), nil
	}
}

type ShowBandwidthDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBandwidthDetailInvoker) Invoke() (*model.ShowBandwidthDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthDetailResponse), nil
	}
}

type ShowCloudPhoneDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudPhoneDetailInvoker) Invoke() (*model.ShowCloudPhoneDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudPhoneDetailResponse), nil
	}
}

type ShowCloudPhoneServerDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudPhoneServerDetailInvoker) Invoke() (*model.ShowCloudPhoneServerDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudPhoneServerDetailResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type StopCloudPhoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopCloudPhoneInvoker) Invoke() (*model.StopCloudPhoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopCloudPhoneResponse), nil
	}
}

type UpdateBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBandwidthInvoker) Invoke() (*model.UpdateBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBandwidthResponse), nil
	}
}

type UpdateCloudPhonePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCloudPhonePropertyInvoker) Invoke() (*model.UpdateCloudPhonePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCloudPhonePropertyResponse), nil
	}
}

type UpdateKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKeypairInvoker) Invoke() (*model.UpdateKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKeypairResponse), nil
	}
}

type UpdatePhoneNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePhoneNameInvoker) Invoke() (*model.UpdatePhoneNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePhoneNameResponse), nil
	}
}

type UpdateServerNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerNameInvoker) Invoke() (*model.UpdateServerNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerNameResponse), nil
	}
}

type InstallApkInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallApkInvoker) Invoke() (*model.InstallApkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallApkResponse), nil
	}
}

type PushFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushFileInvoker) Invoke() (*model.PushFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushFileResponse), nil
	}
}

type RunShellCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunShellCommandInvoker) Invoke() (*model.RunShellCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunShellCommandResponse), nil
	}
}

type RunSyncCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSyncCommandInvoker) Invoke() (*model.RunSyncCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSyncCommandResponse), nil
	}
}

type UninstallApkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallApkInvoker) Invoke() (*model.UninstallApkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallApkResponse), nil
	}
}
