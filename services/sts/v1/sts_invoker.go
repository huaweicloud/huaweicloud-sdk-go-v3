package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sts/v1/model"
)

type DecodeAuthorizationMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DecodeAuthorizationMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DecodeAuthorizationMessageInvoker) Invoke() (*model.DecodeAuthorizationMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DecodeAuthorizationMessageResponse), nil
	}
}

type GetCallerIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetCallerIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetCallerIdentityInvoker) Invoke() (*model.GetCallerIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetCallerIdentityResponse), nil
	}
}
