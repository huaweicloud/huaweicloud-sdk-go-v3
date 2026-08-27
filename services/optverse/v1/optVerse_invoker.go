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
