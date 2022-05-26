package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v1/model"
)

type ListCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCbhInstanceInvoker) Invoke() (*model.ListCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCbhInstanceResponse), nil
	}
}
