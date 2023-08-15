package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/devsecurity/v1/model"
)

type CreateSecAppTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecAppTaskInvoker) Invoke() (*model.CreateSecAppTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecAppTaskResponse), nil
	}
}

type DeleteSecAppTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecAppTaskInvoker) Invoke() (*model.DeleteSecAppTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecAppTaskResponse), nil
	}
}

type ShowSecAppTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecAppTaskResultInvoker) Invoke() (*model.ShowSecAppTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecAppTaskResultResponse), nil
	}
}

type ShowSecAppTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecAppTaskStatusInvoker) Invoke() (*model.ShowSecAppTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecAppTaskStatusResponse), nil
	}
}
