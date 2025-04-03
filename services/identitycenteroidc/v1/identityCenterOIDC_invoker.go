package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenteroidc/v1/model"
)

type RegisterClientInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterClientInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterClientInvoker) Invoke() (*model.RegisterClientResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterClientResponse), nil
	}
}

type StartDeviceAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDeviceAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDeviceAuthorizationInvoker) Invoke() (*model.StartDeviceAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDeviceAuthorizationResponse), nil
	}
}

type CreateTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTokenInvoker) Invoke() (*model.CreateTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTokenResponse), nil
	}
}
