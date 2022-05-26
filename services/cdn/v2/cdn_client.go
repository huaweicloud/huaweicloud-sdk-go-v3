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

// ShowDomainLocationStats 查询域名统计区域运营商数据
//
// - 支持查询90天内的数据。
// - 支持多指标同时查询，不超过5个。
// - 最多同时指定20个域名。
// - 起始时间和结束时间需要同时指定，左闭右开，毫秒级时间戳，必须为5分钟整时刻点，如：0分、5分、10分、15分等，如果传的不是5分钟时刻点，返回数据可能与预期不一致。统一用开始时间表示一个时间段，如：2019-01-24 20:15:00 表示取 [20:15:00, 20:20:00)的统计数据，且左闭右开。
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

// ShowDomainLocationStatsInvoker 查询域名统计区域运营商数据
func (c *CdnClient) ShowDomainLocationStatsInvoker(request *model.ShowDomainLocationStatsRequest) *ShowDomainLocationStatsInvoker {
	requestDef := GenReqDefForShowDomainLocationStats()
	return &ShowDomainLocationStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainStats 查询域名统计基础数据
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

// ShowDomainStatsInvoker 查询域名统计基础数据
func (c *CdnClient) ShowDomainStatsInvoker(request *model.ShowDomainStatsRequest) *ShowDomainStatsInvoker {
	requestDef := GenReqDefForShowDomainStats()
	return &ShowDomainStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopUrl 查询TOP100 URL明细
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

// ShowTopUrlInvoker 查询TOP100 URL明细
func (c *CdnClient) ShowTopUrlInvoker(request *model.ShowTopUrlRequest) *ShowTopUrlInvoker {
	requestDef := GenReqDefForShowTopUrl()
	return &ShowTopUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
