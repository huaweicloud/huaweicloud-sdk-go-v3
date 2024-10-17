package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v2/model"
)

type PutEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutEventsInvoker) Invoke() (*model.PutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutEventsResponse), nil
	}
}

type PutOfficialEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutOfficialEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutOfficialEventsInvoker) Invoke() (*model.PutOfficialEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutOfficialEventsResponse), nil
	}
}
