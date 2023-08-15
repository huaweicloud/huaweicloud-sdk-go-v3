package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/aom/v3/model"
)

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComponentInvoker) Invoke() (*model.CreateComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComponentResponse), nil
	}
}

type CreateEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvInvoker) Invoke() (*model.CreateEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComponentInvoker) Invoke() (*model.DeleteComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComponentResponse), nil
	}
}

type DeleteEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvInvoker) Invoke() (*model.DeleteEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvResponse), nil
	}
}

type ListResourceUnderNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceUnderNodeInvoker) Invoke() (*model.ListResourceUnderNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceUnderNodeResponse), nil
	}
}

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type ShowAppByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppByNameInvoker) Invoke() (*model.ShowAppByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppByNameResponse), nil
	}
}

type ShowComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentInvoker) Invoke() (*model.ShowComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentResponse), nil
	}
}

type ShowComponentByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentByNameInvoker) Invoke() (*model.ShowComponentByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentByNameResponse), nil
	}
}

type ShowEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvInvoker) Invoke() (*model.ShowEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvResponse), nil
	}
}

type ShowEnvByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvByNameInvoker) Invoke() (*model.ShowEnvByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvByNameResponse), nil
	}
}

type UpdateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type UpdateComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateComponentInvoker) Invoke() (*model.UpdateComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateComponentResponse), nil
	}
}

type UpdateEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnvInvoker) Invoke() (*model.UpdateEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnvResponse), nil
	}
}
