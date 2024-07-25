package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
)

type LiveClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewLiveClient(hcClient *httpclient.HcHttpClient) *LiveClient {
	return &LiveClient{HcClient: hcClient}
}

func LiveClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListAreaDetail 查询直播各区域指标分布接口
//
// 查询直播全球区域维度的详细数据接口。
//
// 如果不传入域名，则查询租户下所有播放域名的详细数据。
//
// 当查询租户级别数据时，参数app、stream不生效。
//
// 最大查询跨度1天，最大查询周期90天。
//
// 支持查询当天，当前数据延时少于5分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListAreaDetail(request *model.ListAreaDetailRequest) (*model.ListAreaDetailResponse, error) {
	requestDef := GenReqDefForListAreaDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAreaDetailResponse), nil
	}
}

// ListAreaDetailInvoker 查询直播各区域指标分布接口
func (c *LiveClient) ListAreaDetailInvoker(request *model.ListAreaDetailRequest) *ListAreaDetailInvoker {
	requestDef := GenReqDefForListAreaDetail()
	return &ListAreaDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidthDetail 查询播放带宽趋势接口
//
// 查询播放域名带宽数据。
//
// 如果不传入域名，则查询租户下所有播放域名的带宽数据。
//
// 当查询租户级别带宽数据时，参数app、stream不生效。
//
// 最大查询跨度31天，最大查询周期一年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListBandwidthDetail(request *model.ListBandwidthDetailRequest) (*model.ListBandwidthDetailResponse, error) {
	requestDef := GenReqDefForListBandwidthDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthDetailResponse), nil
	}
}

// ListBandwidthDetailInvoker 查询播放带宽趋势接口
func (c *LiveClient) ListBandwidthDetailInvoker(request *model.ListBandwidthDetailRequest) *ListBandwidthDetailInvoker {
	requestDef := GenReqDefForListBandwidthDetail()
	return &ListBandwidthDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainBandwidthPeak 查询播放带宽峰值接口
//
// 查询指定时间范围内播放带宽峰值。
//
// 如果不传入域名，则查询租户下所有播放域名的带宽峰值。
//
// 当查询租户级别带宽数据时，参数app、stream不生效。
//
// 最大查询跨度31天，最大查询周期一年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListDomainBandwidthPeak(request *model.ListDomainBandwidthPeakRequest) (*model.ListDomainBandwidthPeakResponse, error) {
	requestDef := GenReqDefForListDomainBandwidthPeak()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainBandwidthPeakResponse), nil
	}
}

// ListDomainBandwidthPeakInvoker 查询播放带宽峰值接口
func (c *LiveClient) ListDomainBandwidthPeakInvoker(request *model.ListDomainBandwidthPeakRequest) *ListDomainBandwidthPeakInvoker {
	requestDef := GenReqDefForListDomainBandwidthPeak()
	return &ListDomainBandwidthPeakInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainTrafficDetail 查询播放流量趋势接口
//
// 查询播放域名流量数据。
//
// 如果不传入域名，则查询租户下所有播放域名的流量数据。
//
// 当查询租户级别流量数据时，参数app、stream不生效。
//
// 最大查询跨度31天，最大查询周期一年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListDomainTrafficDetail(request *model.ListDomainTrafficDetailRequest) (*model.ListDomainTrafficDetailResponse, error) {
	requestDef := GenReqDefForListDomainTrafficDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficDetailResponse), nil
	}
}

// ListDomainTrafficDetailInvoker 查询播放流量趋势接口
func (c *LiveClient) ListDomainTrafficDetailInvoker(request *model.ListDomainTrafficDetailRequest) *ListDomainTrafficDetailInvoker {
	requestDef := GenReqDefForListDomainTrafficDetail()
	return &ListDomainTrafficDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainTrafficSummary 查询播放流量汇总接口
//
// 查询指定时间范围内流量汇总量。
//
// 如果不传入域名，则查询租户下所有播放域名的流量汇总量。
//
// 当查询租户级别流量数据时，参数app、stream不生效。
//
// 最大查询跨度31天，最大查询周期一年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListDomainTrafficSummary(request *model.ListDomainTrafficSummaryRequest) (*model.ListDomainTrafficSummaryResponse, error) {
	requestDef := GenReqDefForListDomainTrafficSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficSummaryResponse), nil
	}
}

// ListDomainTrafficSummaryInvoker 查询播放流量汇总接口
func (c *LiveClient) ListDomainTrafficSummaryInvoker(request *model.ListDomainTrafficSummaryRequest) *ListDomainTrafficSummaryInvoker {
	requestDef := GenReqDefForListDomainTrafficSummary()
	return &ListDomainTrafficSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryStreams 查询历史推流列表接口
//
// 查询历史推流列表。
//
// 不能查询现推流。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListHistoryStreams(request *model.ListHistoryStreamsRequest) (*model.ListHistoryStreamsResponse, error) {
	requestDef := GenReqDefForListHistoryStreams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryStreamsResponse), nil
	}
}

// ListHistoryStreamsInvoker 查询历史推流列表接口
func (c *LiveClient) ListHistoryStreamsInvoker(request *model.ListHistoryStreamsRequest) *ListHistoryStreamsInvoker {
	requestDef := GenReqDefForListHistoryStreams()
	return &ListHistoryStreamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlayDomainStreamInfo 查询播放域名下的流数据
//
// 查询播放域名下的监控数据，根据输入时间点，返回查询该时间点所有流的带宽、在线人数、协议。
// 返回的数据粒度为1分钟。
// 最大查询周期7天，数据延迟5分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListPlayDomainStreamInfo(request *model.ListPlayDomainStreamInfoRequest) (*model.ListPlayDomainStreamInfoResponse, error) {
	requestDef := GenReqDefForListPlayDomainStreamInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlayDomainStreamInfoResponse), nil
	}
}

// ListPlayDomainStreamInfoInvoker 查询播放域名下的流数据
func (c *LiveClient) ListPlayDomainStreamInfoInvoker(request *model.ListPlayDomainStreamInfoRequest) *ListPlayDomainStreamInfoInvoker {
	requestDef := GenReqDefForListPlayDomainStreamInfo()
	return &ListPlayDomainStreamInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryHttpCode 查询直播拉流HTTP状态码接口
//
// 查询直播拉流HTTP状态码接口。  获取加速域名1分钟粒度的HTTP返回码  最大查询跨度不能超过24小时，最大查询周期7天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListQueryHttpCode(request *model.ListQueryHttpCodeRequest) (*model.ListQueryHttpCodeResponse, error) {
	requestDef := GenReqDefForListQueryHttpCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryHttpCodeResponse), nil
	}
}

// ListQueryHttpCodeInvoker 查询直播拉流HTTP状态码接口
func (c *LiveClient) ListQueryHttpCodeInvoker(request *model.ListQueryHttpCodeRequest) *ListQueryHttpCodeInvoker {
	requestDef := GenReqDefForListQueryHttpCode()
	return &ListQueryHttpCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordData 查询录制用量接口
//
// 查询直播租户每小时录制的最大并发数，计算1小时内每分钟的并发总路数，取最大值做为统计值。  最大查询跨度31天，最大查询周期90天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListRecordData(request *model.ListRecordDataRequest) (*model.ListRecordDataResponse, error) {
	requestDef := GenReqDefForListRecordData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordDataResponse), nil
	}
}

// ListRecordDataInvoker 查询录制用量接口
func (c *LiveClient) ListRecordDataInvoker(request *model.ListRecordDataRequest) *ListRecordDataInvoker {
	requestDef := GenReqDefForListRecordData()
	return &ListRecordDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotData 查询截图用量接口
//
// 查询直播域名每小时的截图数量。  最大查询跨度31天，最大查询周期90天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListSnapshotData(request *model.ListSnapshotDataRequest) (*model.ListSnapshotDataResponse, error) {
	requestDef := GenReqDefForListSnapshotData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotDataResponse), nil
	}
}

// ListSnapshotDataInvoker 查询截图用量接口
func (c *LiveClient) ListSnapshotDataInvoker(request *model.ListSnapshotDataRequest) *ListSnapshotDataInvoker {
	requestDef := GenReqDefForListSnapshotData()
	return &ListSnapshotDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTranscodeData 查询转码用量接口
//
// 查询直播域名每小时的转码时长数据。  最大查询跨度31天，最大查询周期90天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListTranscodeData(request *model.ListTranscodeDataRequest) (*model.ListTranscodeDataResponse, error) {
	requestDef := GenReqDefForListTranscodeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodeDataResponse), nil
	}
}

// ListTranscodeDataInvoker 查询转码用量接口
func (c *LiveClient) ListTranscodeDataInvoker(request *model.ListTranscodeDataRequest) *ListTranscodeDataInvoker {
	requestDef := GenReqDefForListTranscodeData()
	return &ListTranscodeDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsersOfStream 查询观众趋势接口
//
// 查询观众趋势。  最大查询跨度31天，最大查询周期一年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListUsersOfStream(request *model.ListUsersOfStreamRequest) (*model.ListUsersOfStreamResponse, error) {
	requestDef := GenReqDefForListUsersOfStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersOfStreamResponse), nil
	}
}

// ListUsersOfStreamInvoker 查询观众趋势接口
func (c *LiveClient) ListUsersOfStreamInvoker(request *model.ListUsersOfStreamRequest) *ListUsersOfStreamInvoker {
	requestDef := GenReqDefForListUsersOfStream()
	return &ListUsersOfStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStreamCount 查询域名维度推流路数接口
//
// 查询域名维度推流路数接口。  最大查询跨度31天，最大查询周期1年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ShowStreamCount(request *model.ShowStreamCountRequest) (*model.ShowStreamCountResponse, error) {
	requestDef := GenReqDefForShowStreamCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamCountResponse), nil
	}
}

// ShowStreamCountInvoker 查询域名维度推流路数接口
func (c *LiveClient) ShowStreamCountInvoker(request *model.ShowStreamCountRequest) *ShowStreamCountInvoker {
	requestDef := GenReqDefForShowStreamCount()
	return &ShowStreamCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStreamPortrait 查询播放画像信息接口
//
// 查询播放画像信息。  最大查询跨度1天，最大查询周期31天。
// 不统计协议为HLS的播放时长（play_duration）信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ShowStreamPortrait(request *model.ShowStreamPortraitRequest) (*model.ShowStreamPortraitResponse, error) {
	requestDef := GenReqDefForShowStreamPortrait()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamPortraitResponse), nil
	}
}

// ShowStreamPortraitInvoker 查询播放画像信息接口
func (c *LiveClient) ShowStreamPortraitInvoker(request *model.ShowStreamPortraitRequest) *ShowStreamPortraitInvoker {
	requestDef := GenReqDefForShowStreamPortrait()
	return &ShowStreamPortraitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpBandwidth 查询上行带宽数据接口
//
// 查询上行带宽数据。  最大查询跨度31天，最大查询周期1年。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ShowUpBandwidth(request *model.ShowUpBandwidthRequest) (*model.ShowUpBandwidthResponse, error) {
	requestDef := GenReqDefForShowUpBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpBandwidthResponse), nil
	}
}

// ShowUpBandwidthInvoker 查询上行带宽数据接口
func (c *LiveClient) ShowUpBandwidthInvoker(request *model.ShowUpBandwidthRequest) *ShowUpBandwidthInvoker {
	requestDef := GenReqDefForShowUpBandwidth()
	return &ShowUpBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSingleStreamBitrate 查询推流码率数据接口
//
// 查询推流监控码率数据接口。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的码率数据列表粒度为1秒钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamBitrate(request *model.ListSingleStreamBitrateRequest) (*model.ListSingleStreamBitrateResponse, error) {
	requestDef := GenReqDefForListSingleStreamBitrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamBitrateResponse), nil
	}
}

// ListSingleStreamBitrateInvoker 查询推流码率数据接口
func (c *LiveClient) ListSingleStreamBitrateInvoker(request *model.ListSingleStreamBitrateRequest) *ListSingleStreamBitrateInvoker {
	requestDef := GenReqDefForListSingleStreamBitrate()
	return &ListSingleStreamBitrateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSingleStreamDetail 查询推流监控数据接口
//
// 查询流监控数据接口，包括帧率码率断流情况。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的码率数据列表粒度为1秒钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamDetail(request *model.ListSingleStreamDetailRequest) (*model.ListSingleStreamDetailResponse, error) {
	requestDef := GenReqDefForListSingleStreamDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamDetailResponse), nil
	}
}

// ListSingleStreamDetailInvoker 查询推流监控数据接口
func (c *LiveClient) ListSingleStreamDetailInvoker(request *model.ListSingleStreamDetailRequest) *ListSingleStreamDetailInvoker {
	requestDef := GenReqDefForListSingleStreamDetail()
	return &ListSingleStreamDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSingleStreamFramerate 查询推流帧率数据接口
//
// 查询推流帧率数据接口。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的帧率数据列表粒度为1秒钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamFramerate(request *model.ListSingleStreamFramerateRequest) (*model.ListSingleStreamFramerateResponse, error) {
	requestDef := GenReqDefForListSingleStreamFramerate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamFramerateResponse), nil
	}
}

// ListSingleStreamFramerateInvoker 查询推流帧率数据接口
func (c *LiveClient) ListSingleStreamFramerateInvoker(request *model.ListSingleStreamFramerateRequest) *ListSingleStreamFramerateInvoker {
	requestDef := GenReqDefForListSingleStreamFramerate()
	return &ListSingleStreamFramerateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpStreamDetail 查询CDN上行推流质量数据接口
//
// 查询CDN上行推流质量数据。
//
// 最大查询跨度1天，最大查询周期7天。
//
// 返回的CDN上行推流质量数据列表粒度为1分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LiveClient) ListUpStreamDetail(request *model.ListUpStreamDetailRequest) (*model.ListUpStreamDetailResponse, error) {
	requestDef := GenReqDefForListUpStreamDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpStreamDetailResponse), nil
	}
}

// ListUpStreamDetailInvoker 查询CDN上行推流质量数据接口
func (c *LiveClient) ListUpStreamDetailInvoker(request *model.ListUpStreamDetailRequest) *ListUpStreamDetailInvoker {
	requestDef := GenReqDefForListUpStreamDetail()
	return &ListUpStreamDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
