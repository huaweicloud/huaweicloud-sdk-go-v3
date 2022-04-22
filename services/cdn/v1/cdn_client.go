package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/model"
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

// 创建加速域名
//
// 创建加速域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

// 创建预热缓存任务
//
// 创建预热任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) CreatePreheatingTasks(request *model.CreatePreheatingTasksRequest) (*model.CreatePreheatingTasksResponse, error) {
	requestDef := GenReqDefForCreatePreheatingTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePreheatingTasksResponse), nil
	}
}

// 创建刷新缓存任务
//
// 创建刷新缓存任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) CreateRefreshTasks(request *model.CreateRefreshTasksRequest) (*model.CreateRefreshTasksResponse, error) {
	requestDef := GenReqDefForCreateRefreshTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRefreshTasksResponse), nil
	}
}

// 删除加速域名
//
// 删除加速域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

// 停用加速域名
//
// 停用加速域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) DisableDomain(request *model.DisableDomainRequest) (*model.DisableDomainResponse, error) {
	requestDef := GenReqDefForDisableDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDomainResponse), nil
	}
}

// 启用加速域名
//
// 启用加速域名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) EnableDomain(request *model.EnableDomainRequest) (*model.EnableDomainResponse, error) {
	requestDef := GenReqDefForEnableDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDomainResponse), nil
	}
}

// 查询加速域名
//
// 查询加速域名信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// 查询IP黑白名单
//
// 查询域名已经设置的IP黑白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowBlackWhiteList(request *model.ShowBlackWhiteListRequest) (*model.ShowBlackWhiteListResponse, error) {
	requestDef := GenReqDefForShowBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlackWhiteListResponse), nil
	}
}

// 查询缓存规则
//
// 查询缓存规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowCacheRules(request *model.ShowCacheRulesRequest) (*model.ShowCacheRulesResponse, error) {
	requestDef := GenReqDefForShowCacheRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCacheRulesResponse), nil
	}
}

// 查询所有绑定HTTPS证书的域名信息
//
// 查询所有绑定HTTPS证书的域名信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowCertificatesHttpsInfo(request *model.ShowCertificatesHttpsInfoRequest) (*model.ShowCertificatesHttpsInfoResponse, error) {
	requestDef := GenReqDefForShowCertificatesHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificatesHttpsInfoResponse), nil
	}
}

// 查询加速域名详情
//
// 查询加速域名详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainDetail(request *model.ShowDomainDetailRequest) (*model.ShowDomainDetailResponse, error) {
	requestDef := GenReqDefForShowDomainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainDetailResponse), nil
	}
}

// 查询域名全量配置
//
// 查询域名全量配置接口，支持配置回源请求头、http header配置、url鉴权、证书设置等
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainFullConfig(request *model.ShowDomainFullConfigRequest) (*model.ShowDomainFullConfigResponse, error) {
	requestDef := GenReqDefForShowDomainFullConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainFullConfigResponse), nil
	}
}

// 批量查询域名的统计明细-按域名单独返回
//
// - 支持查询90天内的数据。
// - 查询跨度不能超过7天。
// - 最多同时指定100个域名。
// - 起始时间和结束时间，左闭右开，需要同时指定。
// - 开始时间、结束时间必须传毫秒级时间戳，且必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。
// - 统一用开始时间表示一个时间段，如：2019-01-24 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、请求数类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
// - 如果传的是多个域名，则每个域名的数据分开返回。
// - 支持同时查询多个指标，不超过10个。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainItemDetails(request *model.ShowDomainItemDetailsRequest) (*model.ShowDomainItemDetailsResponse, error) {
	requestDef := GenReqDefForShowDomainItemDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainItemDetailsResponse), nil
	}
}

// 批量查询域名的区域、运营商统计明细-按域名单独返回
//
// - 支持查询90天内的数据。
// - 查询跨度是7天。
// - 最多同时指定100个域名。
// - 起始时间和结束时间，左闭右开，需要同时指定。
// - 开始时间、结束时间必须传毫秒级时间戳，且必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。
// - 统一用开始时间表示一个时间段，如：2019-01-24 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
// - 流量类指标单位统一为Byte（字节）、带宽类指标单位统一为bit/s（比特/秒）、请求数类指标单位统一为次数。
// - 用于查询指定域名、指定统计指标的明细数据。
// - 如果传的是多个域名，则每个域名的数据分开返回。
// - 支持按区域、运营商维度查询统计数据, 回源指标除外。
// - 支持同时查询多个指标，不超过10个。
// - 域名为海外加速场景不适用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainItemLocationDetails(request *model.ShowDomainItemLocationDetailsRequest) (*model.ShowDomainItemLocationDetailsResponse, error) {
	requestDef := GenReqDefForShowDomainItemLocationDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainItemLocationDetailsResponse), nil
	}
}

// 查询域名统计数据-区域运营商
//
// - 支持查询90天内的数据。
// - 支持多指标同时查询，不超过5个。
// - 最多同时指定20个域名。
// - 起始时间和结束时间需要同时指定，左闭右开，毫秒级时间戳，必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点， 返回数据可能与预期不一致。统一用开始时间表示一个时间段，如：2019-01-24 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
// - action取值：location_detail,location_summary
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainLocationStats(request *model.ShowDomainLocationStatsRequest) (*model.ShowDomainLocationStatsResponse, error) {
	requestDef := GenReqDefForShowDomainLocationStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainLocationStatsResponse), nil
	}
}

// 查询域名统计数据
//
// - 支持查询90天内的数据。
// - 支持多指标同时查询，不超过5个。
// - 最多同时指定20个域名。
// - 起始时间和结束时间需要同时指定，左闭右开，毫秒级时间戳，必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。统一用开始时间表示一个时间段，如：2019-01-24 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
// - action取值：detail,summary
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowDomainStats(request *model.ShowDomainStatsRequest) (*model.ShowDomainStatsResponse, error) {
	requestDef := GenReqDefForShowDomainStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainStatsResponse), nil
	}
}

// 查询刷新预热任务详情
//
// 查询刷新预热任务详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowHistoryTaskDetails(request *model.ShowHistoryTaskDetailsRequest) (*model.ShowHistoryTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowHistoryTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTaskDetailsResponse), nil
	}
}

// 查询刷新预热任务
//
// 查询刷新预热任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowHistoryTasks(request *model.ShowHistoryTasksRequest) (*model.ShowHistoryTasksResponse, error) {
	requestDef := GenReqDefForShowHistoryTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTasksResponse), nil
	}
}

// 查询HTTPS配置
//
// 获取加速域名证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowHttpInfo(request *model.ShowHttpInfoRequest) (*model.ShowHttpInfoResponse, error) {
	requestDef := GenReqDefForShowHttpInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpInfoResponse), nil
	}
}

// 查询IP归属信息
//
// 查询IP归属信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowIpInfo(request *model.ShowIpInfoRequest) (*model.ShowIpInfoResponse, error) {
	requestDef := GenReqDefForShowIpInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpInfoResponse), nil
	}
}

// 日志查询
//
// 日志查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowLogs(request *model.ShowLogsRequest) (*model.ShowLogsResponse, error) {
	requestDef := GenReqDefForShowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogsResponse), nil
	}
}

// 查询回源HOST
//
// 查询回源HOST。回源HOST是CDN节点在回源过程中，在源站访问的站点域名，即http请求头中的host信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowOriginHost(request *model.ShowOriginHostRequest) (*model.ShowOriginHostResponse, error) {
	requestDef := GenReqDefForShowOriginHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOriginHostResponse), nil
	}
}

// 查询用户配额
//
// 查询当前用户域名、刷新文件、刷新目录和预热的配额
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// 查询Referer过滤规则
//
// 查询Referer过滤规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowRefer(request *model.ShowReferRequest) (*model.ShowReferResponse, error) {
	requestDef := GenReqDefForShowRefer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReferResponse), nil
	}
}

// 查询响应头配置
//
// 列举header所有配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowResponseHeader(request *model.ShowResponseHeaderRequest) (*model.ShowResponseHeaderResponse, error) {
	requestDef := GenReqDefForShowResponseHeader()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResponseHeaderResponse), nil
	}
}

// 查询TOP100 URL明细
//
// - 查询TOP100 URL明细。
// - 支持查询90天内的数据。
// - 查询跨度不能超过31天。
// - 起始时间和结束时间，左闭右开，需要同时指定。如查询2021-10-24 00:00:00 到 2021-10-25 00:00:00 的数据，表示取 [2021-10-24 00:00:00, 2021-10-25 00:00:00)的统计数据。
// - 开始时间、结束时间必须传毫秒级时间戳，且必须为凌晨0点整时刻点，如果传的不是凌晨0点整时刻点，返回数据可能与预期不一致。
// - 流量类指标单位统一为Byte（字节）、请求数类指标单位统一为次数。用于查询指定域名、指定统计指标的明细数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) ShowTopUrl(request *model.ShowTopUrlRequest) (*model.ShowTopUrlResponse, error) {
	requestDef := GenReqDefForShowTopUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopUrlResponse), nil
	}
}

// 设置IP黑白名单
//
// 设置域名的IP黑白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateBlackWhiteList(request *model.UpdateBlackWhiteListRequest) (*model.UpdateBlackWhiteListResponse, error) {
	requestDef := GenReqDefForUpdateBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBlackWhiteListResponse), nil
	}
}

// 设置缓存规则
//
// 设置CDN节点上缓存资源的缓存策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateCacheRules(request *model.UpdateCacheRulesRequest) (*model.UpdateCacheRulesResponse, error) {
	requestDef := GenReqDefForUpdateCacheRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCacheRulesResponse), nil
	}
}

// 修改域名全量配置
//
// 修改域名全量配置接口，支持配置回源请求头、http header配置、url鉴权、证书设置等
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateDomainFullConfig(request *model.UpdateDomainFullConfigRequest) (*model.UpdateDomainFullConfigResponse, error) {
	requestDef := GenReqDefForUpdateDomainFullConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainFullConfigResponse), nil
	}
}

// 一个证书批量设置多个域名
//
// 一个证书配置多个域名，设置域名强制https回源参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateDomainMultiCertificates(request *model.UpdateDomainMultiCertificatesRequest) (*model.UpdateDomainMultiCertificatesResponse, error) {
	requestDef := GenReqDefForUpdateDomainMultiCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainMultiCertificatesResponse), nil
	}
}

// 修改源站信息
//
// 修改源站信息。源站IP地址或域名都可以指引CDN节点回源到对应的源站服务器，源站域名不能与加速域名相同。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateDomainOrigin(request *model.UpdateDomainOriginRequest) (*model.UpdateDomainOriginResponse, error) {
	requestDef := GenReqDefForUpdateDomainOrigin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainOriginResponse), nil
	}
}

// 开启/关闭回源跟随
//
// 开启此项配置后，当CDN节点回源请求源站返回301/302状态码时，CDN节点会先跳转到301/302对应地址获取资源并缓存后再返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateFollow302Switch(request *model.UpdateFollow302SwitchRequest) (*model.UpdateFollow302SwitchResponse, error) {
	requestDef := GenReqDefForUpdateFollow302Switch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFollow302SwitchResponse), nil
	}
}

// 配置HTTPS
//
// 设置加速域名HTTPS。通过配置加速域名的HTTPS证书，并将其部署在全网CDN节点，实现HTTPS安全加速。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateHttpsInfo(request *model.UpdateHttpsInfoRequest) (*model.UpdateHttpsInfoResponse, error) {
	requestDef := GenReqDefForUpdateHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpsInfoResponse), nil
	}
}

// 修改回源HOST
//
// 修改回源HOST。回源HOST是CDN节点在回源过程中，在源站访问的站点域名，即http请求头中的host信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateOriginHost(request *model.UpdateOriginHostRequest) (*model.UpdateOriginHostResponse, error) {
	requestDef := GenReqDefForUpdateOriginHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOriginHostResponse), nil
	}
}

// 修改私有桶开启关闭状态
//
// 修改私有桶开启关闭状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdatePrivateBucketAccess(request *model.UpdatePrivateBucketAccessRequest) (*model.UpdatePrivateBucketAccessResponse, error) {
	requestDef := GenReqDefForUpdatePrivateBucketAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateBucketAccessResponse), nil
	}
}

// 开启/关闭Range回源
//
// Range回源是指源站在收到CDN节点回源请求时，根据http请求头中的Range信息返回指定范围的数据给CDN节点。
//
// 开启Range回源前需要确认源站是否支持Range请求，若源站不支持Range请求，开启Range回源将导致资源无法缓存。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateRangeSwitch(request *model.UpdateRangeSwitchRequest) (*model.UpdateRangeSwitchResponse, error) {
	requestDef := GenReqDefForUpdateRangeSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRangeSwitchResponse), nil
	}
}

// 设置Referer过滤规则
//
// 设置Referer过滤规则。通过设置过滤策略，对访问者身份进行识别和过滤，实现限制访问来源的目的。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateRefer(request *model.UpdateReferRequest) (*model.UpdateReferResponse, error) {
	requestDef := GenReqDefForUpdateRefer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReferResponse), nil
	}
}

// 新增/修改响应头配置
//
// 新增/修改域名响应头配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdnClient) UpdateResponseHeader(request *model.UpdateResponseHeaderRequest) (*model.UpdateResponseHeaderResponse, error) {
	requestDef := GenReqDefForUpdateResponseHeader()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResponseHeaderResponse), nil
	}
}
