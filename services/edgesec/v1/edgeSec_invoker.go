package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v1/model"
)

type ListEdgeSecSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeSecSubscriptionInvoker) Invoke() (*model.ListEdgeSecSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeSecSubscriptionResponse), nil
	}
}

type CreateEdgeDDoSDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeDDoSDomainsInvoker) Invoke() (*model.CreateEdgeDDoSDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeDDoSDomainsResponse), nil
	}
}

type DeleteEdgeDDoSDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeDDoSDomainsInvoker) Invoke() (*model.DeleteEdgeDDoSDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeDDoSDomainsResponse), nil
	}
}

type ListEdgeDDoSDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeDDoSDomainsInvoker) Invoke() (*model.ListEdgeDDoSDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeDDoSDomainsResponse), nil
	}
}

type ShowStatisticsEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticsEventInvoker) Invoke() (*model.ShowStatisticsEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticsEventResponse), nil
	}
}

type ShowStatisticsTrafficInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticsTrafficInvoker) Invoke() (*model.ShowStatisticsTrafficResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticsTrafficResponse), nil
	}
}

type UpdateEdgeDDoSDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeDDoSDomainsInvoker) Invoke() (*model.UpdateEdgeDDoSDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeDDoSDomainsResponse), nil
	}
}

type ApplyWafPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyWafPolicyInvoker) Invoke() (*model.ApplyWafPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyWafPolicyResponse), nil
	}
}

type CreateEdgeWafDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeWafDomainsInvoker) Invoke() (*model.CreateEdgeWafDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeWafDomainsResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type DeleteEdgeWafDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeWafDomainsInvoker) Invoke() (*model.DeleteEdgeWafDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeWafDomainsResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type ListCdnDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainsInvoker) Invoke() (*model.ListCdnDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainsResponse), nil
	}
}

type ListEdgeWafDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeWafDomainsInvoker) Invoke() (*model.ListEdgeWafDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeWafDomainsResponse), nil
	}
}

type ListPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyInvoker) Invoke() (*model.ListPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyResponse), nil
	}
}

type ShowEdgeWafDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeWafDomainsInvoker) Invoke() (*model.ShowEdgeWafDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeWafDomainsResponse), nil
	}
}

type UpdateEdgeWafDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeWafDomainsInvoker) Invoke() (*model.UpdateEdgeWafDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeWafDomainsResponse), nil
	}
}
