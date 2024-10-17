package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudrtc/v1/model"
)

type ListRtcAbnormalEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcAbnormalEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcAbnormalEventInvoker) Invoke() (*model.ListRtcAbnormalEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcAbnormalEventResponse), nil
	}
}

type ListRtcEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcEventInvoker) Invoke() (*model.ListRtcEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcEventResponse), nil
	}
}

type ListRtcAbnormalEventDimensionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcAbnormalEventDimensionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcAbnormalEventDimensionInvoker) Invoke() (*model.ListRtcAbnormalEventDimensionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcAbnormalEventDimensionResponse), nil
	}
}

type ListRtcAbnormalEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcAbnormalEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcAbnormalEventsInvoker) Invoke() (*model.ListRtcAbnormalEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcAbnormalEventsResponse), nil
	}
}

type ListRtcClientQosDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcClientQosDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcClientQosDetailsInvoker) Invoke() (*model.ListRtcClientQosDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcClientQosDetailsResponse), nil
	}
}

type ListRtcHistoryQualityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcHistoryQualityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcHistoryQualityInvoker) Invoke() (*model.ListRtcHistoryQualityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcHistoryQualityResponse), nil
	}
}

type ListRtcHistoryScaleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcHistoryScaleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcHistoryScaleInvoker) Invoke() (*model.ListRtcHistoryScaleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcHistoryScaleResponse), nil
	}
}

type ListRtcHistoryUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcHistoryUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcHistoryUsageInvoker) Invoke() (*model.ListRtcHistoryUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcHistoryUsageResponse), nil
	}
}

type ListRtcRealtimeNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcRealtimeNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcRealtimeNetworkInvoker) Invoke() (*model.ListRtcRealtimeNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcRealtimeNetworkResponse), nil
	}
}

type ListRtcRealtimeQualityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcRealtimeQualityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcRealtimeQualityInvoker) Invoke() (*model.ListRtcRealtimeQualityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcRealtimeQualityResponse), nil
	}
}

type ListRtcRealtimeScaleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcRealtimeScaleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcRealtimeScaleInvoker) Invoke() (*model.ListRtcRealtimeScaleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcRealtimeScaleResponse), nil
	}
}

type ListRtcRealtimeScaleDimensionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcRealtimeScaleDimensionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcRealtimeScaleDimensionInvoker) Invoke() (*model.ListRtcRealtimeScaleDimensionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcRealtimeScaleDimensionResponse), nil
	}
}

type ListRtcRoomListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcRoomListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcRoomListInvoker) Invoke() (*model.ListRtcRoomListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcRoomListResponse), nil
	}
}

type ListRtcUserListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRtcUserListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRtcUserListInvoker) Invoke() (*model.ListRtcUserListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRtcUserListResponse), nil
	}
}
