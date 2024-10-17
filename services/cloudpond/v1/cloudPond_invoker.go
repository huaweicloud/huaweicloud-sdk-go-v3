package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpond/v1/model"
)

type CreateEdgeSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeSiteInvoker) Invoke() (*model.CreateEdgeSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeSiteResponse), nil
	}
}

type DeleteEdgeSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeSiteInvoker) Invoke() (*model.DeleteEdgeSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeSiteResponse), nil
	}
}

type ListEdgeSitesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeSitesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeSitesInvoker) Invoke() (*model.ListEdgeSitesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeSitesResponse), nil
	}
}

type ShowEdgeSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEdgeSiteInvoker) Invoke() (*model.ShowEdgeSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeSiteResponse), nil
	}
}

type UpdateEdgeSiteInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeSiteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeSiteInvoker) Invoke() (*model.UpdateEdgeSiteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeSiteResponse), nil
	}
}

type ListEdgeSiteMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeSiteMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeSiteMetricsInvoker) Invoke() (*model.ListEdgeSiteMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeSiteMetricsResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListRacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRacksInvoker) Invoke() (*model.ListRacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRacksResponse), nil
	}
}

type ShowRackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRackInvoker) Invoke() (*model.ShowRackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRackResponse), nil
	}
}

type ListSupportedRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportedRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSupportedRegionsInvoker) Invoke() (*model.ListSupportedRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportedRegionsResponse), nil
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

type ListSupportedZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportedZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSupportedZonesInvoker) Invoke() (*model.ListSupportedZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportedZonesResponse), nil
	}
}
