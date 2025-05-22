package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
)

type BatchCreateOrDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateOrDeleteTagsInvoker) Invoke() (*model.BatchCreateOrDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

type BatchImportSecretsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportSecretsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchImportSecretsInvoker) Invoke() (*model.BatchImportSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportSecretsResponse), nil
	}
}

type CheckSecretsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckSecretsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckSecretsInvoker) Invoke() (*model.CheckSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckSecretsResponse), nil
	}
}

type CreateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyInvoker) Invoke() (*model.CreateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyResponse), nil
	}
}

type CreateGrantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGrantsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGrantsInvoker) Invoke() (*model.CreateGrantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGrantsResponse), nil
	}
}

type CreateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretInvoker) Invoke() (*model.CreateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretResponse), nil
	}
}

type CreateSecretEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretEventInvoker) Invoke() (*model.CreateSecretEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretEventResponse), nil
	}
}

type CreateSecretTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretTagInvoker) Invoke() (*model.CreateSecretTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretTagResponse), nil
	}
}

type CreateSecretVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretVersionInvoker) Invoke() (*model.CreateSecretVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretVersionResponse), nil
	}
}

type DeleteGrantInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGrantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGrantInvoker) Invoke() (*model.DeleteGrantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGrantResponse), nil
	}
}

type DeleteSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretInvoker) Invoke() (*model.DeleteSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretResponse), nil
	}
}

type DeleteSecretEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretEventInvoker) Invoke() (*model.DeleteSecretEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretEventResponse), nil
	}
}

type DeleteSecretForScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretForScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretForScheduleInvoker) Invoke() (*model.DeleteSecretForScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretForScheduleResponse), nil
	}
}

type DeleteSecretStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretStageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretStageInvoker) Invoke() (*model.DeleteSecretStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretStageResponse), nil
	}
}

type DeleteSecretTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretTagInvoker) Invoke() (*model.DeleteSecretTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretTagResponse), nil
	}
}

type DownloadSecretBlobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadSecretBlobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadSecretBlobInvoker) Invoke() (*model.DownloadSecretBlobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSecretBlobResponse), nil
	}
}

type GenerateRandomPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *GenerateRandomPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GenerateRandomPasswordInvoker) Invoke() (*model.GenerateRandomPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenerateRandomPasswordResponse), nil
	}
}

type ListGrantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGrantsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGrantsInvoker) Invoke() (*model.ListGrantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGrantsResponse), nil
	}
}

type ListNotificationRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationRecordsInvoker) Invoke() (*model.ListNotificationRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationRecordsResponse), nil
	}
}

type ListProjectSecretsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectSecretsTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectSecretsTagsInvoker) Invoke() (*model.ListProjectSecretsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectSecretsTagsResponse), nil
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

type ListSecretEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretEventsInvoker) Invoke() (*model.ListSecretEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretEventsResponse), nil
	}
}

type ListSecretTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretTagsInvoker) Invoke() (*model.ListSecretTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretTagsResponse), nil
	}
}

type ListSecretTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretTaskInvoker) Invoke() (*model.ListSecretTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretTaskResponse), nil
	}
}

type ListSecretVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretVersionsInvoker) Invoke() (*model.ListSecretVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretVersionsResponse), nil
	}
}

type ListSecretsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretsInvoker) Invoke() (*model.ListSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretsResponse), nil
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

type RestoreSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreSecretInvoker) Invoke() (*model.RestoreSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreSecretResponse), nil
	}
}

type RotateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *RotateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RotateSecretInvoker) Invoke() (*model.RotateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotateSecretResponse), nil
	}
}

type ShowAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgencyInvoker) Invoke() (*model.ShowAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyResponse), nil
	}
}

type ShowSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretInvoker) Invoke() (*model.ShowSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretResponse), nil
	}
}

type ShowSecretEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretEventInvoker) Invoke() (*model.ShowSecretEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretEventResponse), nil
	}
}

type ShowSecretFunctionTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretFunctionTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretFunctionTemplatesInvoker) Invoke() (*model.ShowSecretFunctionTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretFunctionTemplatesResponse), nil
	}
}

type ShowSecretStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretStageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretStageInvoker) Invoke() (*model.ShowSecretStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretStageResponse), nil
	}
}

type ShowSecretVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretVersionInvoker) Invoke() (*model.ShowSecretVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretVersionResponse), nil
	}
}

type ShowSecretsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretsConfigInvoker) Invoke() (*model.ShowSecretsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretsConfigResponse), nil
	}
}

type ShowUserDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserDetailInvoker) Invoke() (*model.ShowUserDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserDetailResponse), nil
	}
}

type UpdateGrantInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGrantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGrantInvoker) Invoke() (*model.UpdateGrantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGrantResponse), nil
	}
}

type UpdateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecretInvoker) Invoke() (*model.UpdateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretResponse), nil
	}
}

type UpdateSecretEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecretEventInvoker) Invoke() (*model.UpdateSecretEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretEventResponse), nil
	}
}

type UpdateSecretStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretStageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecretStageInvoker) Invoke() (*model.UpdateSecretStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretStageResponse), nil
	}
}

type UpdateSecretsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecretsConfigInvoker) Invoke() (*model.UpdateSecretsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretsConfigResponse), nil
	}
}

type UpdateUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserPasswordInvoker) Invoke() (*model.UpdateUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserPasswordResponse), nil
	}
}

type UpdateVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVersionInvoker) Invoke() (*model.UpdateVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVersionResponse), nil
	}
}

type UploadSecretBlobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadSecretBlobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadSecretBlobInvoker) Invoke() (*model.UploadSecretBlobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadSecretBlobResponse), nil
	}
}
