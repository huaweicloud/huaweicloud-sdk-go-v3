package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ies/v1/model"
)

type CreateEdgeSiteInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSupportedRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportedRegionsInvoker) Invoke() (*model.ListSupportedRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportedRegionsResponse), nil
	}
}
