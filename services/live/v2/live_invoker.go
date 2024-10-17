package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
)

type ListAreaDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAreaDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAreaDetailInvoker) Invoke() (*model.ListAreaDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAreaDetailResponse), nil
	}
}

type ListBandwidthDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthDetailInvoker) Invoke() (*model.ListBandwidthDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthDetailResponse), nil
	}
}

type ListDomainBandwidthPeakInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainBandwidthPeakInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainBandwidthPeakInvoker) Invoke() (*model.ListDomainBandwidthPeakResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainBandwidthPeakResponse), nil
	}
}

type ListDomainTrafficDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainTrafficDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainTrafficDetailInvoker) Invoke() (*model.ListDomainTrafficDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainTrafficDetailResponse), nil
	}
}

type ListDomainTrafficSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainTrafficSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainTrafficSummaryInvoker) Invoke() (*model.ListDomainTrafficSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainTrafficSummaryResponse), nil
	}
}

type ListHistoryStreamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryStreamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryStreamsInvoker) Invoke() (*model.ListHistoryStreamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryStreamsResponse), nil
	}
}

type ListPlayDomainStreamInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlayDomainStreamInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlayDomainStreamInfoInvoker) Invoke() (*model.ListPlayDomainStreamInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlayDomainStreamInfoResponse), nil
	}
}

type ListQueryHttpCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryHttpCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQueryHttpCodeInvoker) Invoke() (*model.ListQueryHttpCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryHttpCodeResponse), nil
	}
}

type ListRecordDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordDataInvoker) Invoke() (*model.ListRecordDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordDataResponse), nil
	}
}

type ListSnapshotDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSnapshotDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSnapshotDataInvoker) Invoke() (*model.ListSnapshotDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSnapshotDataResponse), nil
	}
}

type ListTranscodeDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTranscodeDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTranscodeDataInvoker) Invoke() (*model.ListTranscodeDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTranscodeDataResponse), nil
	}
}

type ListUsersOfStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersOfStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersOfStreamInvoker) Invoke() (*model.ListUsersOfStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersOfStreamResponse), nil
	}
}

type ShowStreamCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStreamCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStreamCountInvoker) Invoke() (*model.ShowStreamCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStreamCountResponse), nil
	}
}

type ShowStreamPortraitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStreamPortraitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStreamPortraitInvoker) Invoke() (*model.ShowStreamPortraitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStreamPortraitResponse), nil
	}
}

type ShowUpBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUpBandwidthInvoker) Invoke() (*model.ShowUpBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpBandwidthResponse), nil
	}
}

type ListSingleStreamBitrateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSingleStreamBitrateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSingleStreamBitrateInvoker) Invoke() (*model.ListSingleStreamBitrateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSingleStreamBitrateResponse), nil
	}
}

type ListSingleStreamDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSingleStreamDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSingleStreamDetailInvoker) Invoke() (*model.ListSingleStreamDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSingleStreamDetailResponse), nil
	}
}

type ListSingleStreamFramerateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSingleStreamFramerateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSingleStreamFramerateInvoker) Invoke() (*model.ListSingleStreamFramerateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSingleStreamFramerateResponse), nil
	}
}

type ListUpStreamDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUpStreamDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUpStreamDetailInvoker) Invoke() (*model.ListUpStreamDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUpStreamDetailResponse), nil
	}
}
