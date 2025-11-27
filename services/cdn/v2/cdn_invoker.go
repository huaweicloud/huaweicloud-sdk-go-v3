package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v2/model"
)

type ApplyDomainTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyDomainTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyDomainTemplateInvoker) Invoke() (*model.ApplyDomainTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyDomainTemplateResponse), nil
	}
}

type BatchCopyDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCopyDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCopyDomainInvoker) Invoke() (*model.BatchCopyDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCopyDomainResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type BatchUpdateRuleStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateRuleStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateRuleStatusInvoker) Invoke() (*model.BatchUpdateRuleStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateRuleStatusResponse), nil
	}
}

type CreateAccessControlTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessControlTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccessControlTaskInvoker) Invoke() (*model.CreateAccessControlTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessControlTaskResponse), nil
	}
}

type CreateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainInvoker) Invoke() (*model.CreateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainResponse), nil
	}
}

type CreateDomainByDuplicateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainByDuplicateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainByDuplicateInvoker) Invoke() (*model.CreateDomainByDuplicateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainByDuplicateResponse), nil
	}
}

type CreateDomainTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainTemplateInvoker) Invoke() (*model.CreateDomainTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainTemplateResponse), nil
	}
}

type CreateExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExportTaskInvoker) Invoke() (*model.CreateExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExportTaskResponse), nil
	}
}

type CreatePreheatingTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePreheatingTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePreheatingTasksInvoker) Invoke() (*model.CreatePreheatingTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePreheatingTasksResponse), nil
	}
}

type CreateRefreshTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRefreshTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRefreshTasksInvoker) Invoke() (*model.CreateRefreshTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRefreshTasksResponse), nil
	}
}

type CreateRuleNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRuleNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRuleNewInvoker) Invoke() (*model.CreateRuleNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRuleNewResponse), nil
	}
}

type CreateShareCacheGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareCacheGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShareCacheGroupsInvoker) Invoke() (*model.CreateShareCacheGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareCacheGroupsResponse), nil
	}
}

type CreateSubscriptionTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubscriptionTaskInvoker) Invoke() (*model.CreateSubscriptionTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionTaskResponse), nil
	}
}

type CreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagsInvoker) Invoke() (*model.CreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagsResponse), nil
	}
}

type DeleteDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainInvoker) Invoke() (*model.DeleteDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainResponse), nil
	}
}

type DeleteDomainTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainTemplateInvoker) Invoke() (*model.DeleteDomainTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainTemplateResponse), nil
	}
}

type DeleteRuleNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRuleNewInvoker) Invoke() (*model.DeleteRuleNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleNewResponse), nil
	}
}

type DeleteShareCacheGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareCacheGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteShareCacheGroupsInvoker) Invoke() (*model.DeleteShareCacheGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareCacheGroupsResponse), nil
	}
}

type DeleteSubscriptionTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubscriptionTaskInvoker) Invoke() (*model.DeleteSubscriptionTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionTaskResponse), nil
	}
}

type DisableDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableDomainInvoker) Invoke() (*model.DisableDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableDomainResponse), nil
	}
}

type DownloadRegionCarrierExcelInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadRegionCarrierExcelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadStatisticsExcelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadStatisticsExcelInvoker) Invoke() (*model.DownloadStatisticsExcelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadStatisticsExcelResponse), nil
	}
}

type EnableDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDomainInvoker) Invoke() (*model.EnableDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDomainResponse), nil
	}
}

type ExportStatsOpenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportStatsOpenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportStatsOpenInvoker) Invoke() (*model.ExportStatsOpenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportStatsOpenResponse), nil
	}
}

type ListAccessControlTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessControlTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessControlTaskInvoker) Invoke() (*model.ListAccessControlTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessControlTaskResponse), nil
	}
}

type ListBanUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBanUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBanUrlInvoker) Invoke() (*model.ListBanUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBanUrlResponse), nil
	}
}

type ListCdnDomainTopIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainTopIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCdnDomainTopIpsInvoker) Invoke() (*model.ListCdnDomainTopIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainTopIpsResponse), nil
	}
}

type ListCdnDomainTopOriginUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainTopOriginUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCdnDomainTopOriginUrlInvoker) Invoke() (*model.ListCdnDomainTopOriginUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainTopOriginUrlResponse), nil
	}
}

type ListCdnDomainTopPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainTopPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCdnDomainTopPathInvoker) Invoke() (*model.ListCdnDomainTopPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainTopPathResponse), nil
	}
}

type ListCdnDomainTopRefersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainTopRefersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCdnDomainTopRefersInvoker) Invoke() (*model.ListCdnDomainTopRefersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainTopRefersResponse), nil
	}
}

type ListCdnDomainTopUasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCdnDomainTopUasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCdnDomainTopUasInvoker) Invoke() (*model.ListCdnDomainTopUasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCdnDomainTopUasResponse), nil
	}
}

type ListDomainClientStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainClientStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainClientStatsInvoker) Invoke() (*model.ListDomainClientStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainClientStatsResponse), nil
	}
}

type ListDomainConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainConfigsInvoker) Invoke() (*model.ListDomainConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainConfigsResponse), nil
	}
}

type ListDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainsInvoker) Invoke() (*model.ListDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainsResponse), nil
	}
}

type ListExportTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExportTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExportTasksInvoker) Invoke() (*model.ListExportTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExportTasksResponse), nil
	}
}

type ListRuleDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRuleDetailsInvoker) Invoke() (*model.ListRuleDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleDetailsResponse), nil
	}
}

type ListShareCacheGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShareCacheGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListShareCacheGroupsInvoker) Invoke() (*model.ListShareCacheGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShareCacheGroupsResponse), nil
	}
}

type ListSpecialConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecialConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecialConfigurationInvoker) Invoke() (*model.ListSpecialConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecialConfigurationResponse), nil
	}
}

type ListSubscriptionTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionTasksInvoker) Invoke() (*model.ListSubscriptionTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionTasksResponse), nil
	}
}

type ModifyAccountInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyAccountInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyAccountInfoInvoker) Invoke() (*model.ModifyAccountInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyAccountInfoResponse), nil
	}
}

type SetChargeModesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetChargeModesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetChargeModesInvoker) Invoke() (*model.SetChargeModesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetChargeModesResponse), nil
	}
}

type SetStatsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetStatsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetStatsConfigInvoker) Invoke() (*model.SetStatsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetStatsConfigResponse), nil
	}
}

type ShowAppliedTemplateRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppliedTemplateRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppliedTemplateRecordInvoker) Invoke() (*model.ShowAppliedTemplateRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppliedTemplateRecordResponse), nil
	}
}

type ShowBandwidthCalcInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowBandwidthCalcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowBandwidthCalcInvoker) Invoke() (*model.ShowBandwidthCalcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthCalcResponse), nil
	}
}

type ShowCertificatesHttpsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificatesHttpsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificatesHttpsInfoInvoker) Invoke() (*model.ShowCertificatesHttpsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificatesHttpsInfoResponse), nil
	}
}

type ShowChargeModesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowChargeModesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowChargeModesInvoker) Invoke() (*model.ShowChargeModesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowChargeModesResponse), nil
	}
}

type ShowDomainCountryStatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainCountryStatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainCountryStatInvoker) Invoke() (*model.ShowDomainCountryStatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainCountryStatResponse), nil
	}
}

type ShowDomainDetailByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainDetailByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDomainFullConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDomainLocationStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDomainStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainStatsInvoker) Invoke() (*model.ShowDomainStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainStatsResponse), nil
	}
}

type ShowDomainTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainTemplateInvoker) Invoke() (*model.ShowDomainTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainTemplateResponse), nil
	}
}

type ShowHistoryTaskDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryTaskDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryTaskDetailsInvoker) Invoke() (*model.ShowHistoryTaskDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryTaskDetailsResponse), nil
	}
}

type ShowHistoryTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryTasksInvoker) Invoke() (*model.ShowHistoryTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryTasksResponse), nil
	}
}

type ShowIpInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpInfoInvoker) Invoke() (*model.ShowIpInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpInfoResponse), nil
	}
}

type ShowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLogsInvoker) Invoke() (*model.ShowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogsResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type ShowSpecialUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpecialUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSpecialUserInvoker) Invoke() (*model.ShowSpecialUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpecialUserResponse), nil
	}
}

type ShowStatsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatsConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatsConfigsInvoker) Invoke() (*model.ShowStatsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatsConfigsResponse), nil
	}
}

type ShowTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagsInvoker) Invoke() (*model.ShowTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagsResponse), nil
	}
}

type ShowTopDomainNamesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowTopDomainNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

func (i *ShowTopUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTopUrlInvoker) Invoke() (*model.ShowTopUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopUrlResponse), nil
	}
}

type ShowUrlTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUrlTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUrlTaskInfoInvoker) Invoke() (*model.ShowUrlTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUrlTaskInfoResponse), nil
	}
}

type ShowVerifyDomainOwnerInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVerifyDomainOwnerInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVerifyDomainOwnerInfoInvoker) Invoke() (*model.ShowVerifyDomainOwnerInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVerifyDomainOwnerInfoResponse), nil
	}
}

type UpdateDomainFullConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainFullConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainFullConfigInvoker) Invoke() (*model.UpdateDomainFullConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainFullConfigResponse), nil
	}
}

type UpdateDomainMultiCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainMultiCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainMultiCertificatesInvoker) Invoke() (*model.UpdateDomainMultiCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainMultiCertificatesResponse), nil
	}
}

type UpdateDomainTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainTemplateInvoker) Invoke() (*model.UpdateDomainTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainTemplateResponse), nil
	}
}

type UpdateFullRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFullRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFullRuleInvoker) Invoke() (*model.UpdateFullRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFullRuleResponse), nil
	}
}

type UpdatePrivateBucketAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateBucketAccessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateBucketAccessInvoker) Invoke() (*model.UpdatePrivateBucketAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateBucketAccessResponse), nil
	}
}

type UpdateRuleNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRuleNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRuleNewInvoker) Invoke() (*model.UpdateRuleNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRuleNewResponse), nil
	}
}

type UpdateShareCacheGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateShareCacheGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateShareCacheGroupsInvoker) Invoke() (*model.UpdateShareCacheGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateShareCacheGroupsResponse), nil
	}
}

type UpdateSubscriptionTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubscriptionTaskInvoker) Invoke() (*model.UpdateSubscriptionTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionTaskResponse), nil
	}
}

type VerifyDomainOwnerInvoker struct {
	*invoker.BaseInvoker
}

func (i *VerifyDomainOwnerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *VerifyDomainOwnerInvoker) Invoke() (*model.VerifyDomainOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerifyDomainOwnerResponse), nil
	}
}
