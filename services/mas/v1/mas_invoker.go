package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/mas/v1/model"
)

type ShowNameSpaceListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNameSpaceListInvoker) Invoke() (*model.ShowNameSpaceListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNameSpaceListResponse), nil
	}
}
