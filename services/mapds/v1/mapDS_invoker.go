package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/mapds/v1/model"
)

type CreateCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCredentialInvoker) Invoke() (*model.CreateCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCredentialResponse), nil
	}
}

type CreateSasTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSasTokenInvoker) Invoke() (*model.CreateSasTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSasTokenResponse), nil
	}
}

type DeleteCedentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCedentialInvoker) Invoke() (*model.DeleteCedentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCedentialResponse), nil
	}
}

type ShowCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCredentialInvoker) Invoke() (*model.ShowCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCredentialResponse), nil
	}
}

type ShowMapTileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMapTileInvoker) Invoke() (*model.ShowMapTileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMapTileResponse), nil
	}
}
