package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpond/v2/model"
)

type ListNetworkDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetworkDevicesInvoker) Invoke() (*model.ListNetworkDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkDevicesResponse), nil
	}
}

type ShowNetworkDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkDeviceInvoker) Invoke() (*model.ShowNetworkDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkDeviceResponse), nil
	}
}

type ListNetworkDeviceOfferingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkDeviceOfferingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetworkDeviceOfferingsInvoker) Invoke() (*model.ListNetworkDeviceOfferingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkDeviceOfferingsResponse), nil
	}
}

type ListServerOfferingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerOfferingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerOfferingsInvoker) Invoke() (*model.ListServerOfferingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerOfferingsResponse), nil
	}
}

type ListSaleCyclesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSaleCyclesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSaleCyclesInvoker) Invoke() (*model.ListSaleCyclesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSaleCyclesResponse), nil
	}
}

type ListServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersInvoker) Invoke() (*model.ListServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ListStorageGearsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageGearsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStorageGearsInvoker) Invoke() (*model.ListStorageGearsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageGearsResponse), nil
	}
}

type ListStoragePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoragePoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStoragePoolsInvoker) Invoke() (*model.ListStoragePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoragePoolsResponse), nil
	}
}

type ShowStoragePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStoragePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStoragePoolInvoker) Invoke() (*model.ShowStoragePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStoragePoolResponse), nil
	}
}

type ListStorageTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStorageTypesInvoker) Invoke() (*model.ListStorageTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageTypesResponse), nil
	}
}
