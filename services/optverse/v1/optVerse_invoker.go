package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/optverse/v1/model"
)

type BatchDeleteEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteEvolveTaskInvoker) Invoke() (*model.BatchDeleteEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteEvolveTaskResponse), nil
	}
}

type CreateAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlgorithmInvoker) Invoke() (*model.CreateAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlgorithmResponse), nil
	}
}

type CreateEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEvolveTaskInvoker) Invoke() (*model.CreateEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEvolveTaskResponse), nil
	}
}

type DeleteAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlgorithmInvoker) Invoke() (*model.DeleteAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlgorithmResponse), nil
	}
}

type DeleteAlgorithmFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlgorithmFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlgorithmFileInvoker) Invoke() (*model.DeleteAlgorithmFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlgorithmFileResponse), nil
	}
}

type DeleteEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEvolveTaskInvoker) Invoke() (*model.DeleteEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEvolveTaskResponse), nil
	}
}

type ImportAlgorithmFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportAlgorithmFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportAlgorithmFileInvoker) Invoke() (*model.ImportAlgorithmFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportAlgorithmFileResponse), nil
	}
}

type ListAlgorithmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlgorithmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlgorithmsInvoker) Invoke() (*model.ListAlgorithmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlgorithmsResponse), nil
	}
}

type ListDirectoryByAlgorithmIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDirectoryByAlgorithmIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDirectoryByAlgorithmIdInvoker) Invoke() (*model.ListDirectoryByAlgorithmIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDirectoryByAlgorithmIdResponse), nil
	}
}

type ListDirectoryByResultCommitIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDirectoryByResultCommitIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDirectoryByResultCommitIdInvoker) Invoke() (*model.ListDirectoryByResultCommitIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDirectoryByResultCommitIdResponse), nil
	}
}

type ListEvolveTaskMetasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEvolveTaskMetasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEvolveTaskMetasInvoker) Invoke() (*model.ListEvolveTaskMetasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEvolveTaskMetasResponse), nil
	}
}

type ListEvolveTaskStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEvolveTaskStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEvolveTaskStatsInvoker) Invoke() (*model.ListEvolveTaskStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEvolveTaskStatsResponse), nil
	}
}

type SaveAlgorithmFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveAlgorithmFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveAlgorithmFileInvoker) Invoke() (*model.SaveAlgorithmFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveAlgorithmFileResponse), nil
	}
}

type ShowAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlgorithmInvoker) Invoke() (*model.ShowAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlgorithmResponse), nil
	}
}

type ShowAlgorithmFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlgorithmFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlgorithmFileInvoker) Invoke() (*model.ShowAlgorithmFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlgorithmFileResponse), nil
	}
}

type ShowTaskDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskDetailsInvoker) Invoke() (*model.ShowTaskDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskDetailsResponse), nil
	}
}

type ShowTaskResultCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskResultCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskResultCommitInvoker) Invoke() (*model.ShowTaskResultCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResultCommitResponse), nil
	}
}

type ShowTaskResultListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskResultListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskResultListInvoker) Invoke() (*model.ShowTaskResultListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResultListResponse), nil
	}
}

type ShowTaskRunningDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskRunningDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskRunningDetailsInvoker) Invoke() (*model.ShowTaskRunningDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskRunningDetailsResponse), nil
	}
}

type ShowTaskRunningLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskRunningLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskRunningLogInvoker) Invoke() (*model.ShowTaskRunningLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskRunningLogResponse), nil
	}
}

type StartEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartEvolveTaskInvoker) Invoke() (*model.StartEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartEvolveTaskResponse), nil
	}
}

type StopEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopEvolveTaskInvoker) Invoke() (*model.StopEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopEvolveTaskResponse), nil
	}
}

type UpdateAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlgorithmInvoker) Invoke() (*model.UpdateAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlgorithmResponse), nil
	}
}

type UpdateEvolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEvolveTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEvolveTaskInvoker) Invoke() (*model.UpdateEvolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEvolveTaskResponse), nil
	}
}

type DeleteModelAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModelAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteModelAssetInvoker) Invoke() (*model.DeleteModelAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModelAssetResponse), nil
	}
}

type ListModelAssetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModelAssetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModelAssetsInvoker) Invoke() (*model.ListModelAssetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModelAssetsResponse), nil
	}
}

type ShowModelAssetDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModelAssetDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowModelAssetDetailInvoker) Invoke() (*model.ShowModelAssetDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModelAssetDetailResponse), nil
	}
}

type UpdateModelAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModelAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateModelAssetInvoker) Invoke() (*model.UpdateModelAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModelAssetResponse), nil
	}
}

type CancelChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelChatInvoker) Invoke() (*model.CancelChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelChatResponse), nil
	}
}

type CreateArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateArtifactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateArtifactsInvoker) Invoke() (*model.CreateArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateArtifactsResponse), nil
	}
}

type DeleteChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteChatInvoker) Invoke() (*model.DeleteChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteChatResponse), nil
	}
}

type DownloadFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadFileInvoker) Invoke() (*model.DownloadFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadFileResponse), nil
	}
}

type ListArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArtifactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListArtifactsInvoker) Invoke() (*model.ListArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArtifactsResponse), nil
	}
}

type ListChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListChatInvoker) Invoke() (*model.ListChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChatResponse), nil
	}
}

type PublishChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishChatInvoker) Invoke() (*model.PublishChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishChatResponse), nil
	}
}

type ShowChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowChatInvoker) Invoke() (*model.ShowChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowChatResponse), nil
	}
}

type UpdateChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateChatInvoker) Invoke() (*model.UpdateChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateChatResponse), nil
	}
}

type UploadFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadFileInvoker) Invoke() (*model.UploadFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFileResponse), nil
	}
}

type CreateModelServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModelServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModelServiceInvoker) Invoke() (*model.CreateModelServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModelServiceResponse), nil
	}
}

type CreateModelServiceTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModelServiceTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModelServiceTaskInvoker) Invoke() (*model.CreateModelServiceTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModelServiceTaskResponse), nil
	}
}

type DeleteModelServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModelServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteModelServiceInvoker) Invoke() (*model.DeleteModelServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModelServiceResponse), nil
	}
}

type ListModelServiceTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModelServiceTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModelServiceTasksInvoker) Invoke() (*model.ListModelServiceTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModelServiceTasksResponse), nil
	}
}

type ShowModelServiceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModelServiceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowModelServiceDetailInvoker) Invoke() (*model.ShowModelServiceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModelServiceDetailResponse), nil
	}
}

type ShowModelServiceListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModelServiceListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowModelServiceListInvoker) Invoke() (*model.ShowModelServiceListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModelServiceListResponse), nil
	}
}

type ShowModelServiceTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModelServiceTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowModelServiceTaskInvoker) Invoke() (*model.ShowModelServiceTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModelServiceTaskResponse), nil
	}
}

type StartModelServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartModelServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartModelServiceInvoker) Invoke() (*model.StartModelServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartModelServiceResponse), nil
	}
}

type StopModelServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopModelServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopModelServiceInvoker) Invoke() (*model.StopModelServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopModelServiceResponse), nil
	}
}

type UpdateModelServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModelServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateModelServiceInvoker) Invoke() (*model.UpdateModelServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModelServiceResponse), nil
	}
}

type UploadModelServiceTaskFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadModelServiceTaskFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadModelServiceTaskFileInvoker) Invoke() (*model.UploadModelServiceTaskFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadModelServiceTaskFileResponse), nil
	}
}

type PublishModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishModelInvoker) Invoke() (*model.PublishModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishModelResponse), nil
	}
}

type AuthorizePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizePermissionInvoker) Invoke() (*model.AuthorizePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizePermissionResponse), nil
	}
}

type ListBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBucketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBucketsInvoker) Invoke() (*model.ListBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBucketsResponse), nil
	}
}

type ListObjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObjectInvoker) Invoke() (*model.ListObjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObjectResponse), nil
	}
}

type ListPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermissionInvoker) Invoke() (*model.ListPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermissionResponse), nil
	}
}

type RevokePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RevokePermissionInvoker) Invoke() (*model.RevokePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokePermissionResponse), nil
	}
}
