package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cdn/v2/model"
)

type BatchCopyDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCopyDomainInvoker) Invoke() (*model.BatchCopyDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCopyDomainResponse), nil
	}
}

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

type ListDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainsInvoker) Invoke() (*model.ListDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainsResponse), nil
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

type ShowDomainDetailByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainDetailByNameInvoker) Invoke() (*model.ShowDomainDetailByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainDetailByNameResponse), nil
	}
}

type ShowDomainFullConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainFullConfigInvoker) Invoke() (*model.ShowDomainFullConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainFullConfigResponse), nil
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

type UpdateDomainFullConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainFullConfigInvoker) Invoke() (*model.UpdateDomainFullConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainFullConfigResponse), nil
	}
}
