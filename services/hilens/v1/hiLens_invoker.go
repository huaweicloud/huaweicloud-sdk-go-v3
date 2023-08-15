package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/hilens/v1/model"
)

type ListDeviceAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeviceAlarmsInvoker) Invoke() (*model.ListDeviceAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeviceAlarmsResponse), nil
	}
}

type ListDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevicesInvoker) Invoke() (*model.ListDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevicesResponse), nil
	}
}
