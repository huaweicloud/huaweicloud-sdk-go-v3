package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/css/v2/model"
)

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type RestartClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartClusterInvoker) Invoke() (*model.RestartClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartClusterResponse), nil
	}
}

type RollingRestartInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollingRestartInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RollingRestartInvoker) Invoke() (*model.RollingRestartResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollingRestartResponse), nil
	}
}

type StartAutoCreateSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAutoCreateSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAutoCreateSnapshotsInvoker) Invoke() (*model.StartAutoCreateSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAutoCreateSnapshotsResponse), nil
	}
}

type StopAutoCreateSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAutoCreateSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopAutoCreateSnapshotsInvoker) Invoke() (*model.StopAutoCreateSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAutoCreateSnapshotsResponse), nil
	}
}

type DeleteLogstashConfInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogstashConfInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLogstashConfInvoker) Invoke() (*model.DeleteLogstashConfResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogstashConfResponse), nil
	}
}

type DeleteLogstashTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogstashTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLogstashTemplateInvoker) Invoke() (*model.DeleteLogstashTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogstashTemplateResponse), nil
	}
}
