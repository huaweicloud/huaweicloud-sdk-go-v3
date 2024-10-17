package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/deh/v1/model"
)

type BatchCreateDedicatedHostTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDedicatedHostTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDedicatedHostTagsInvoker) Invoke() (*model.BatchCreateDedicatedHostTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDedicatedHostTagsResponse), nil
	}
}

type BatchDeleteDedicatedHostTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDedicatedHostTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDedicatedHostTagsInvoker) Invoke() (*model.BatchDeleteDedicatedHostTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDedicatedHostTagsResponse), nil
	}
}

type ListDedicatedHostTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDedicatedHostTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDedicatedHostTypesInvoker) Invoke() (*model.ListDedicatedHostTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDedicatedHostTypesResponse), nil
	}
}

type ListDedicatedHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDedicatedHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDedicatedHostsInvoker) Invoke() (*model.ListDedicatedHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDedicatedHostsResponse), nil
	}
}

type ListDedicatedHostsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDedicatedHostsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDedicatedHostsByTagsInvoker) Invoke() (*model.ListDedicatedHostsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDedicatedHostsByTagsResponse), nil
	}
}

type ListServersDedicatedHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersDedicatedHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersDedicatedHostInvoker) Invoke() (*model.ListServersDedicatedHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersDedicatedHostResponse), nil
	}
}

type ShowDedicatedHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDedicatedHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDedicatedHostInvoker) Invoke() (*model.ShowDedicatedHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDedicatedHostResponse), nil
	}
}

type ShowDedicatedHostTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDedicatedHostTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDedicatedHostTagsInvoker) Invoke() (*model.ShowDedicatedHostTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDedicatedHostTagsResponse), nil
	}
}

type ShowQuotaSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaSetsInvoker) Invoke() (*model.ShowQuotaSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaSetsResponse), nil
	}
}

type UpdateDedicatedHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDedicatedHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDedicatedHostInvoker) Invoke() (*model.UpdateDedicatedHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDedicatedHostResponse), nil
	}
}
