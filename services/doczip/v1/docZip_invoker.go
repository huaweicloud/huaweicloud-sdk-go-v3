package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/doczip/v1/model"
)

type CreateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplicationInvoker) Invoke() (*model.CreateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type ListApplicationDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationDetailsInvoker) Invoke() (*model.ListApplicationDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationDetailsResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type PublishApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishApplicationInvoker) Invoke() (*model.PublishApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishApplicationResponse), nil
	}
}

type ShowAppResourceLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppResourceLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppResourceLimitInvoker) Invoke() (*model.ShowAppResourceLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResourceLimitResponse), nil
	}
}

type ShowAppSecretKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppSecretKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppSecretKeyInvoker) Invoke() (*model.ShowAppSecretKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppSecretKeyResponse), nil
	}
}

type ShowBlocksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlocksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBlocksInvoker) Invoke() (*model.ShowBlocksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlocksResponse), nil
	}
}

type UpdateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApplicationInvoker) Invoke() (*model.UpdateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationResponse), nil
	}
}
