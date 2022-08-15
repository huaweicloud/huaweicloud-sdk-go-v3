package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type CreateAkSkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAkSkInvoker) Invoke() (*model.CreateAkSkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAkSkResponse), nil
	}
}

type DeleteAkSkInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAkSkInvoker) Invoke() (*model.DeleteAkSkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAkSkResponse), nil
	}
}

type ShowAkSksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAkSksInvoker) Invoke() (*model.ShowAkSksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAkSksResponse), nil
	}
}

type ListAkSkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAkSkInvoker) Invoke() (*model.ListAkSkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAkSkResponse), nil
	}
}

type ListBusinessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBusinessInvoker) Invoke() (*model.ListBusinessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBusinessResponse), nil
	}
}

type ListEnvMonitorItemInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvMonitorItemInvoker) Invoke() (*model.ListEnvMonitorItemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvMonitorItemResponse), nil
	}
}

type SaveMonitorItemConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveMonitorItemConfigInvoker) Invoke() (*model.SaveMonitorItemConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveMonitorItemConfigResponse), nil
	}
}

type SearchApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchApplicationInvoker) Invoke() (*model.SearchApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchApplicationResponse), nil
	}
}

type ShowMasterAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMasterAddressInvoker) Invoke() (*model.ShowMasterAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMasterAddressResponse), nil
	}
}

type ListAppEnvsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppEnvsInvoker) Invoke() (*model.ListAppEnvsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppEnvsResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListEnvTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvTagsInvoker) Invoke() (*model.ListEnvTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvTagsResponse), nil
	}
}

type ShowTopologyTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopologyTreeInvoker) Invoke() (*model.ShowTopologyTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopologyTreeResponse), nil
	}
}

type ListOpenRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOpenRegionInvoker) Invoke() (*model.ListOpenRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOpenRegionResponse), nil
	}
}

type ListSupportedRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupportedRegionInvoker) Invoke() (*model.ListSupportedRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupportedRegionResponse), nil
	}
}

type ListEnvInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvInstancesInvoker) Invoke() (*model.ListEnvInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvInstancesResponse), nil
	}
}

type ShowClobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClobDetailInvoker) Invoke() (*model.ShowClobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClobDetailResponse), nil
	}
}

type ShowEnvMonitorItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvMonitorItemsInvoker) Invoke() (*model.ShowEnvMonitorItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvMonitorItemsResponse), nil
	}
}

type ShowEventDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventDetailInvoker) Invoke() (*model.ShowEventDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventDetailResponse), nil
	}
}

type ShowMonitorItemViewConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonitorItemViewConfigInvoker) Invoke() (*model.ShowMonitorItemViewConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonitorItemViewConfigResponse), nil
	}
}

type ShowRawTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRawTableInvoker) Invoke() (*model.ShowRawTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRawTableResponse), nil
	}
}

type ShowSpanSearchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpanSearchInvoker) Invoke() (*model.ShowSpanSearchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpanSearchResponse), nil
	}
}

type ShowSumTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSumTableInvoker) Invoke() (*model.ShowSumTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSumTableResponse), nil
	}
}

type ShowTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopologyInvoker) Invoke() (*model.ShowTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopologyResponse), nil
	}
}

type ShowTraceEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTraceEventsInvoker) Invoke() (*model.ShowTraceEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTraceEventsResponse), nil
	}
}

type ShowTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrendInvoker) Invoke() (*model.ShowTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrendResponse), nil
	}
}
