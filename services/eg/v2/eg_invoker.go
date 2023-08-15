package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/eg/v2/model"
)

type PutEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutEventsInvoker) Invoke() (*model.PutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutEventsResponse), nil
	}
}
