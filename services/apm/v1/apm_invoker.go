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

type ListAlarmDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmDataInvoker) Invoke() (*model.ListAlarmDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmDataResponse), nil
	}
}

type ListAlarmNotifyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmNotifyInvoker) Invoke() (*model.ListAlarmNotifyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmNotifyResponse), nil
	}
}

type ChangeAgentStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAgentStatusInvoker) Invoke() (*model.ChangeAgentStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAgentStatusResponse), nil
	}
}

type DeleteAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgentInvoker) Invoke() (*model.DeleteAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgentResponse), nil
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

type SearchAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchAgentInvoker) Invoke() (*model.SearchAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchAgentResponse), nil
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

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
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

type ShowBusinessDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBusinessDetailInvoker) Invoke() (*model.ShowBusinessDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBusinessDetailResponse), nil
	}
}

type ShowSubBusinessDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubBusinessDetailInvoker) Invoke() (*model.ShowSubBusinessDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubBusinessDetailResponse), nil
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

type ShowFlameLineTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlameLineTreeInvoker) Invoke() (*model.ShowFlameLineTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlameLineTreeResponse), nil
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

type SearchBusinessTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchBusinessTopologyInvoker) Invoke() (*model.SearchBusinessTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchBusinessTopologyResponse), nil
	}
}

type SearchEnvTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchEnvTopologyInvoker) Invoke() (*model.SearchEnvTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchEnvTopologyResponse), nil
	}
}

type CreateBusinessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBusinessInvoker) Invoke() (*model.CreateBusinessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBusinessResponse), nil
	}
}

type ShowAccessPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessPointInvoker) Invoke() (*model.ShowAccessPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessPointResponse), nil
	}
}

type ShowTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTokenInvoker) Invoke() (*model.ShowTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTokenResponse), nil
	}
}

type ListBusinessEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBusinessEnvInvoker) Invoke() (*model.ListBusinessEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBusinessEnvResponse), nil
	}
}

type SearchTransactionInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchTransactionInvoker) Invoke() (*model.SearchTransactionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchTransactionResponse), nil
	}
}

type SearchTransactionConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchTransactionConfigInvoker) Invoke() (*model.SearchTransactionConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchTransactionConfigResponse), nil
	}
}

type ShowTransactionDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTransactionDetailInvoker) Invoke() (*model.ShowTransactionDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTransactionDetailResponse), nil
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

type ShowMonitorItemDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonitorItemDetailInvoker) Invoke() (*model.ShowMonitorItemDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonitorItemDetailResponse), nil
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
