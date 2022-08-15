package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
)

type BatchCreateOrDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteTagsInvoker) Invoke() (*model.BatchCreateOrDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

type CreateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretInvoker) Invoke() (*model.CreateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretResponse), nil
	}
}

type CreateSecretTagInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateSecretVersionInvoker) Invoke() (*model.CreateSecretVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretVersionResponse), nil
	}
}

type DeleteSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretInvoker) Invoke() (*model.DeleteSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretResponse), nil
	}
}

type DeleteSecretForScheduleInvoker struct {
	*invoker.BaseInvoker
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

func (i *DownloadSecretBlobInvoker) Invoke() (*model.DownloadSecretBlobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSecretBlobResponse), nil
	}
}

type ListProjectSecretsTagsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListResourceInstancesInvoker) Invoke() (*model.ListResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstancesResponse), nil
	}
}

type ListSecretTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretTagsInvoker) Invoke() (*model.ListSecretTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretTagsResponse), nil
	}
}

type ListSecretVersionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListSecretsInvoker) Invoke() (*model.ListSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretsResponse), nil
	}
}

type RestoreSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreSecretInvoker) Invoke() (*model.RestoreSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreSecretResponse), nil
	}
}

type ShowSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretInvoker) Invoke() (*model.ShowSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretResponse), nil
	}
}

type ShowSecretStageInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowSecretVersionInvoker) Invoke() (*model.ShowSecretVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretVersionResponse), nil
	}
}

type UpdateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretInvoker) Invoke() (*model.UpdateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretResponse), nil
	}
}

type UpdateSecretStageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretStageInvoker) Invoke() (*model.UpdateSecretStageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretStageResponse), nil
	}
}

type UploadSecretBlobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadSecretBlobInvoker) Invoke() (*model.UploadSecretBlobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadSecretBlobResponse), nil
	}
}
