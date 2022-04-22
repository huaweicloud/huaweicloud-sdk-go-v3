package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
)

type LiveClient struct {
	HcClient *http_client.HcHttpClient
}

func NewLiveClient(hcClient *http_client.HcHttpClient) *LiveClient {
	return &LiveClient{HcClient: hcClient}
}

func LiveClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 查询直播各区域指标分布接口
//
// 查询直播全球区域维度的详细数据接口。
//
// 最大查询跨度1天，最大查询周期90天。
//
// 支持查询当天，当前数据延时少于1分钟。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListAreaDetail(request *model.ListAreaDetailRequest) (*model.ListAreaDetailResponse, error) {
	requestDef := GenReqDefForListAreaDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAreaDetailResponse), nil
	}
}

// 查询播放带宽趋势接口
//
// 查询播放域名带宽数据。  最大查询跨度31天，最大查询周期一年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListBandwidthDetail(request *model.ListBandwidthDetailRequest) (*model.ListBandwidthDetailResponse, error) {
	requestDef := GenReqDefForListBandwidthDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthDetailResponse), nil
	}
}

// 查询播放带宽峰值接口
//
// 查询指定时间范围内播放带宽峰值。  最大查询跨度31天，最大查询周期一年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListDomainBandwidthPeak(request *model.ListDomainBandwidthPeakRequest) (*model.ListDomainBandwidthPeakResponse, error) {
	requestDef := GenReqDefForListDomainBandwidthPeak()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainBandwidthPeakResponse), nil
	}
}

// 查询播放流量趋势接口
//
// 查询播放域名流量数据。  最大查询跨度31天，最大查询周期一年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListDomainTrafficDetail(request *model.ListDomainTrafficDetailRequest) (*model.ListDomainTrafficDetailResponse, error) {
	requestDef := GenReqDefForListDomainTrafficDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficDetailResponse), nil
	}
}

// 查询播放流量汇总接口
//
// 查询指定时间范围内流量汇总量。  最大查询跨度31天，最大查询周期一年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListDomainTrafficSummary(request *model.ListDomainTrafficSummaryRequest) (*model.ListDomainTrafficSummaryResponse, error) {
	requestDef := GenReqDefForListDomainTrafficSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficSummaryResponse), nil
	}
}

// 查询历史推流列表接口
//
// 查询历史推流列表。
//
// 不能查询现推流。
//
// 最大查询跨度1天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListHistoryStreams(request *model.ListHistoryStreamsRequest) (*model.ListHistoryStreamsResponse, error) {
	requestDef := GenReqDefForListHistoryStreams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryStreamsResponse), nil
	}
}

// 查询直播拉流HTTP状态码接口
//
// 查询直播拉流HTTP状态码接口。  获取加速域名1分钟粒度的HTTP返回码  最大查询跨度不能超过24小时，最大查询周期7天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListQueryHttpCode(request *model.ListQueryHttpCodeRequest) (*model.ListQueryHttpCodeResponse, error) {
	requestDef := GenReqDefForListQueryHttpCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryHttpCodeResponse), nil
	}
}

// 查询录制用量接口
//
// 查询直播租户每小时录制的最大并发数，计算1小时内每分钟的并发总路数，取最大值做为统计值。  最大查询跨度31天，最大查询周期90天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListRecordData(request *model.ListRecordDataRequest) (*model.ListRecordDataResponse, error) {
	requestDef := GenReqDefForListRecordData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordDataResponse), nil
	}
}

// 查询截图用量接口
//
// 查询直播域名每小时的截图数量。  最大查询跨度31天，最大查询周期90天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListSnapshotData(request *model.ListSnapshotDataRequest) (*model.ListSnapshotDataResponse, error) {
	requestDef := GenReqDefForListSnapshotData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotDataResponse), nil
	}
}

// 查询转码用量接口
//
// 查询直播域名每小时的转码时长数据。  最大查询跨度31天，最大查询周期90天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListTranscodeData(request *model.ListTranscodeDataRequest) (*model.ListTranscodeDataResponse, error) {
	requestDef := GenReqDefForListTranscodeData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodeDataResponse), nil
	}
}

// 查询直播转码任务数接口
//
// 查询5分钟粒度的各档位转码任务数。
//
// 仅支持查询视频转码任务数。
//
// 最大查询跨度7天，最大查询周期90天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListTranscodeTaskCount(request *model.ListTranscodeTaskCountRequest) (*model.ListTranscodeTaskCountResponse, error) {
	requestDef := GenReqDefForListTranscodeTaskCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodeTaskCountResponse), nil
	}
}

// 查询观众趋势接口
//
// 查询观众趋势。  最大查询跨度31天，最大查询周期一年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListUsersOfStream(request *model.ListUsersOfStreamRequest) (*model.ListUsersOfStreamResponse, error) {
	requestDef := GenReqDefForListUsersOfStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersOfStreamResponse), nil
	}
}

// 查询域名维度推流路数接口
//
// 查询域名维度推流路数接口。  最大查询跨度31天，最大查询周期1年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ShowStreamCount(request *model.ShowStreamCountRequest) (*model.ShowStreamCountResponse, error) {
	requestDef := GenReqDefForShowStreamCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamCountResponse), nil
	}
}

// 查询播放画像信息接口
//
// 查询播放画像信息。  最大查询跨度1天，最大查询周期31天。
// 不统计协议为HLS的播放时长（play_duration）信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ShowStreamPortrait(request *model.ShowStreamPortraitRequest) (*model.ShowStreamPortraitResponse, error) {
	requestDef := GenReqDefForShowStreamPortrait()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamPortraitResponse), nil
	}
}

// 查询上行带宽数据接口
//
// 查询上行带宽数据。  最大查询跨度31天，最大查询周期1年。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ShowUpBandwidth(request *model.ShowUpBandwidthRequest) (*model.ShowUpBandwidthResponse, error) {
	requestDef := GenReqDefForShowUpBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpBandwidthResponse), nil
	}
}

// 查询推流码率数据接口
//
// 查询推流监控码率数据接口。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的码率数据列表粒度为1秒钟。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamBitrate(request *model.ListSingleStreamBitrateRequest) (*model.ListSingleStreamBitrateResponse, error) {
	requestDef := GenReqDefForListSingleStreamBitrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamBitrateResponse), nil
	}
}

// 查询流监控数据接口
//
// 查询流监控数据接口，包括帧率码率断流情况。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的码率数据列表粒度为1秒钟。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamDetail(request *model.ListSingleStreamDetailRequest) (*model.ListSingleStreamDetailResponse, error) {
	requestDef := GenReqDefForListSingleStreamDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamDetailResponse), nil
	}
}

// 查询推流帧率数据接口
//
// 查询推流帧率数据接口。
//
// 最大查询跨度1天，最大查询周期1个月。
//
// 返回的帧率数据列表粒度为1秒钟。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *LiveClient) ListSingleStreamFramerate(request *model.ListSingleStreamFramerateRequest) (*model.ListSingleStreamFramerateResponse, error) {
	requestDef := GenReqDefForListSingleStreamFramerate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamFramerateResponse), nil
	}
}
