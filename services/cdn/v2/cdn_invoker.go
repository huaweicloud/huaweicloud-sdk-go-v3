package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v2/model"
)

type DownloadRegionCarrierExcelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadRegionCarrierExcelInvoker) Invoke() (*model.DownloadRegionCarrierExcelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadRegionCarrierExcelResponse), nil
	}
}

type DownloadStatisticsExcelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadStatisticsExcelInvoker) Invoke() (*model.DownloadStatisticsExcelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadStatisticsExcelResponse), nil
	}
}

type SetChargeModesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetChargeModesInvoker) Invoke() (*model.SetChargeModesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetChargeModesResponse), nil
	}
}

type ShowBandwidthCalcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBandwidthCalcInvoker) Invoke() (*model.ShowBandwidthCalcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthCalcResponse), nil
	}
}

type ShowChargeModesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowChargeModesInvoker) Invoke() (*model.ShowChargeModesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowChargeModesResponse), nil
	}
}

type ShowDomainLocationStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainLocationStatsInvoker) Invoke() (*model.ShowDomainLocationStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainLocationStatsResponse), nil
	}
}

type ShowDomainStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainStatsInvoker) Invoke() (*model.ShowDomainStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainStatsResponse), nil
	}
}

type ShowTopDomainNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopDomainNamesInvoker) Invoke() (*model.ShowTopDomainNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopDomainNamesResponse), nil
	}
}

type ShowTopUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopUrlInvoker) Invoke() (*model.ShowTopUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopUrlResponse), nil
	}
}
