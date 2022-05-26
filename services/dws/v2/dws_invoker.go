package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSnapshotInvoker) Invoke() (*model.CreateSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSnapshotResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
	}
}

type DeleteSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSnapshotInvoker) Invoke() (*model.DeleteSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSnapshotResponse), nil
	}
}

type ListClusterDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterDetailsInvoker) Invoke() (*model.ListClusterDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterDetailsResponse), nil
	}
}

type ListClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersInvoker) Invoke() (*model.ListClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersResponse), nil
	}
}

type ListNodeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeTypesInvoker) Invoke() (*model.ListNodeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeTypesResponse), nil
	}
}

type ListSnapshotDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotDetailsInvoker) Invoke() (*model.ListSnapshotDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotDetailsResponse), nil
	}
}

type ListSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotsInvoker) Invoke() (*model.ListSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotsResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeClusterInvoker) Invoke() (*model.ResizeClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeClusterResponse), nil
	}
}

type RestartClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartClusterInvoker) Invoke() (*model.RestartClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartClusterResponse), nil
	}
}

type RestoreClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreClusterInvoker) Invoke() (*model.RestoreClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreClusterResponse), nil
	}
}
