package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/vas/v2/model"
)

type CreateTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTasksInvoker) Invoke() (*model.CreateTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTasksResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type ListTasksDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksDetailsInvoker) Invoke() (*model.ListTasksDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksDetailsResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type StartTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartTaskInvoker) Invoke() (*model.StartTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartTaskResponse), nil
	}
}

type StopTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopTaskInvoker) Invoke() (*model.StopTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopTaskResponse), nil
	}
}

type UpdateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInvoker) Invoke() (*model.UpdateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskResponse), nil
	}
}
