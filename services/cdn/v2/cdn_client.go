package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v2/model"
)

type CdnClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCdnClient(hcClient *http_client.HcHttpClient) *CdnClient {
	return &CdnClient{HcClient: hcClient}
}

func CdnClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// BatchCopyDomain 批量域名复制
//
// 批量域名复制接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) BatchCopyDomain(request *model.BatchCopyDomainRequest) (*model.BatchCopyDomainResponse, error) {
	requestDef := GenReqDefForBatchCopyDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCopyDomainResponse), nil
	}
}

// BatchCopyDomainInvoker 批量域名复制
func (c *CdnClient) BatchCopyDomainInvoker(request *model.BatchCopyDomainRequest) *BatchCopyDomainInvoker {
	requestDef := GenReqDefForBatchCopyDomain()
	return &BatchCopyDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadRegionCarrierExcel 下载区域运营商指标数据表格文件
//
// - 下载区域运营商指标数据表格文件。
//
// - 支持下载90天内的指标数据表格。
//
// - 时间跨度不能超过31天。
//
// - 起始时间和结束时间，左闭右开。如时间跨度为2022-10-24 00:00:00 到 2022-10-25 00:00:00，表示取 [2022-10-24 00:00:00, 2022-10-25 00:00:00)的统计数据。
//
// - 起始时间、结束时间必须传毫秒级时间戳，起始时间和结束时间必须同时指定。
//
// - 单租户调用频率：10次/min。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) DownloadRegionCarrierExcel(request *model.DownloadRegionCarrierExcelRequest) (*model.DownloadRegionCarrierExcelResponse, error) {
	requestDef := GenReqDefForDownloadRegionCarrierExcel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadRegionCarrierExcelResponse), nil
	}
}

// DownloadRegionCarrierExcelInvoker 下载区域运营商指标数据表格文件
func (c *CdnClient) DownloadRegionCarrierExcelInvoker(request *model.DownloadRegionCarrierExcelRequest) *DownloadRegionCarrierExcelInvoker {
	requestDef := GenReqDefForDownloadRegionCarrierExcel()
	return &DownloadRegionCarrierExcelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadStatisticsExcel 下载统计指标数据表格文件
//
// - 下载统计指标数据表格文件。
//
// - 支持下载90天内的指标数据。
//
// - 时间跨度不能超过31天。
//
// - 起始时间和结束时间，左闭右开。如时间跨度为2022-10-24 00:00:00 到 2022-10-25 00:00:00，表示取 [2022-10-24 00:00:00, 2022-10-25 00:00:00)的统计数据。
//
// - 起始时间、结束时间必须传毫秒级时间戳，起始时间和结束时间必须同时指定。
//
// - 单租户调用频率：10次/min。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) DownloadStatisticsExcel(request *model.DownloadStatisticsExcelRequest) (*model.DownloadStatisticsExcelResponse, error) {
	requestDef := GenReqDefForDownloadStatisticsExcel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadStatisticsExcelResponse), nil
	}
}

// DownloadStatisticsExcelInvoker 下载统计指标数据表格文件
func (c *CdnClient) DownloadStatisticsExcelInvoker(request *model.DownloadStatisticsExcelRequest) *DownloadStatisticsExcelInvoker {
	requestDef := GenReqDefForDownloadStatisticsExcel()
	return &DownloadStatisticsExcelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetChargeModes 设置用户计费模式
//
// - 设置用户计费模式。
//
// - 服务区域仅支持mainland_china（国内）
//
// - 计费模式仅支持设置flux（流量），bw（带宽）
//
// - 加速类型仅支持base（基础加速）
//
// - 单租户调用频率：10次/min。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) SetChargeModes(request *model.SetChargeModesRequest) (*model.SetChargeModesResponse, error) {
	requestDef := GenReqDefForSetChargeModes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetChargeModesResponse), nil
	}
}

// SetChargeModesInvoker 设置用户计费模式
func (c *CdnClient) SetChargeModesInvoker(request *model.SetChargeModesRequest) *SetChargeModesInvoker {
	requestDef := GenReqDefForSetChargeModes()
	return &SetChargeModesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBandwidthCalc 查询域名带宽峰值类数据
//
// - 查询域名带宽峰值类数据。
//
// - 支持查询90天内的数据。
//
// - 查询时间跨度不能超过31天。
//
// - 起始时间和结束时间，左闭右开。如查询2022-10-24 00:00:00 到 2022-10-25 00:00:00 的数据，表示取 [2022-10-24 00:00:00, 2022-10-25 00:00:00)的统计数据。
//
// - 起始时间、结束时间必须传毫秒级时间戳，起始时间和结束时间必须同时指定。
//
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、峰值类指标单位统一为bps（比特率），请求数类和状态码类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
//
// - 单租户调用频率：2次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowBandwidthCalc(request *model.ShowBandwidthCalcRequest) (*model.ShowBandwidthCalcResponse, error) {
	requestDef := GenReqDefForShowBandwidthCalc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthCalcResponse), nil
	}
}

// ShowBandwidthCalcInvoker 查询域名带宽峰值类数据
func (c *CdnClient) ShowBandwidthCalcInvoker(request *model.ShowBandwidthCalcRequest) *ShowBandwidthCalcInvoker {
	requestDef := GenReqDefForShowBandwidthCalc()
	return &ShowBandwidthCalcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowChargeModes 查询用户计费模式
//
// - 查询用户计费模式。
//
// - 服务区域仅支持mainland_china（国内，默认）和outside_mainland_china（海外）
//
// - 计费模式状态支持active（已生效），upcoming（待生效）两种状态，默认为active(已生效)
//
// - 加速类型仅支持base（基础加速）
//
// - 单租户调用频率：5次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowChargeModes(request *model.ShowChargeModesRequest) (*model.ShowChargeModesResponse, error) {
	requestDef := GenReqDefForShowChargeModes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowChargeModesResponse), nil
	}
}

// ShowChargeModesInvoker 查询用户计费模式
func (c *CdnClient) ShowChargeModesInvoker(request *model.ShowChargeModesRequest) *ShowChargeModesInvoker {
	requestDef := GenReqDefForShowChargeModes()
	return &ShowChargeModesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainLocationStats 查询域名统计区域运营商数据
//
// - 支持查询90天内的数据。
//
// - 支持多指标同时查询，不超过5个。
//
// - 最多同时指定20个域名。
//
// - 起始时间和结束时间需要同时指定，左闭右开，毫秒级时间戳，必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。统一用开始时间表示一个时间段，如：2019-01-24
// 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
//
// - action取值：location_detail,location_summary
//
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、请求数类和状态码类指标单位统一为次数。用于查询指定域名、指定统计指标的区域运营商明细数据。
//
// - 单租户调用频率：15次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowDomainLocationStats(request *model.ShowDomainLocationStatsRequest) (*model.ShowDomainLocationStatsResponse, error) {
	requestDef := GenReqDefForShowDomainLocationStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainLocationStatsResponse), nil
	}
}

// ShowDomainLocationStatsInvoker 查询域名统计区域运营商数据
func (c *CdnClient) ShowDomainLocationStatsInvoker(request *model.ShowDomainLocationStatsRequest) *ShowDomainLocationStatsInvoker {
	requestDef := GenReqDefForShowDomainLocationStats()
	return &ShowDomainLocationStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainStats 查询域名统计基础数据
//
// - 支持查询90天内的数据。
//
// - 支持多指标同时查询，不超过5个。
//
// - 最多同时指定20个域名。
//
// - 起始时间和结束时间需要同时指定，左闭右开，毫秒级时间戳，必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。统一用开始时间表示一个时间段，如：2019-01-24
// 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
//
// - action取值：detail,summary
//
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、请求数类和状态码类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
//
// - 单租户调用频率：15次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowDomainStats(request *model.ShowDomainStatsRequest) (*model.ShowDomainStatsResponse, error) {
	requestDef := GenReqDefForShowDomainStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainStatsResponse), nil
	}
}

// ShowDomainStatsInvoker 查询域名统计基础数据
func (c *CdnClient) ShowDomainStatsInvoker(request *model.ShowDomainStatsRequest) *ShowDomainStatsInvoker {
	requestDef := GenReqDefForShowDomainStats()
	return &ShowDomainStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopDomainNames 查询TOP域名
//
// - 查询TOP域名。
//
// - 支持查询90天内的数据。
//
// - 查询时间跨度不能超过1天。
//
// - 起始时间和结束时间，左闭右开，必须同时指定。如查询2022-10-24 00:00:00 到 2022-10-25 00:00:00 的数据，表示取 [2022-10-24 00:00:00, 2022-10-25 00:00:00)的统计数据。
//
// - 起始时间、结束时间必须传整点毫秒级时间戳。
//
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、请求数类和状态码类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
//
// - 单租户调用频率：5次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowTopDomainNames(request *model.ShowTopDomainNamesRequest) (*model.ShowTopDomainNamesResponse, error) {
	requestDef := GenReqDefForShowTopDomainNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopDomainNamesResponse), nil
	}
}

// ShowTopDomainNamesInvoker 查询TOP域名
func (c *CdnClient) ShowTopDomainNamesInvoker(request *model.ShowTopDomainNamesRequest) *ShowTopDomainNamesInvoker {
	requestDef := GenReqDefForShowTopDomainNames()
	return &ShowTopDomainNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopUrl 查询TOP100 URL明细
//
// - 查询TOP100 URL明细。
//
// - 支持查询90天内的数据。
//
// - 查询跨度不能超过31天。
//
// - 起始时间和结束时间，左闭右开，需要同时指定。如查询2021-10-24 00:00:00 到 2021-10-25 00:00:00 的数据，表示取 [2021-10-24 00:00:00, 2021-10-25 00:00:00)的统计数据。
//
// - 开始时间、结束时间必须传毫秒级时间戳，且必须为凌晨0点整时刻点，如果传的不是凌晨0点整时刻点，返回数据可能与预期不一致。
//
// - 流量类指标单位统一为Byte（字节）、请求数类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
//
// - 单租户调用频率：5次/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CdnClient) ShowTopUrl(request *model.ShowTopUrlRequest) (*model.ShowTopUrlResponse, error) {
	requestDef := GenReqDefForShowTopUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopUrlResponse), nil
	}
}

// ShowTopUrlInvoker 查询TOP100 URL明细
func (c *CdnClient) ShowTopUrlInvoker(request *model.ShowTopUrlRequest) *ShowTopUrlInvoker {
	requestDef := GenReqDefForShowTopUrl()
	return &ShowTopUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
