package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudrtc/v1/model"
)

type CloudRTCClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudRTCClient(hcClient *http_client.HcHttpClient) *CloudRTCClient {
	return &CloudRTCClient{HcClient: hcClient}
}

func CloudRTCClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListRtcAbnormalEventDimension 查询异常事件用户分布
//
// 查询指定APP下指定时间内的通话异常明细数据分布情况。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcAbnormalEventDimension(request *model.ListRtcAbnormalEventDimensionRequest) (*model.ListRtcAbnormalEventDimensionResponse, error) {
	requestDef := GenReqDefForListRtcAbnormalEventDimension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcAbnormalEventDimensionResponse), nil
	}
}

// ListRtcAbnormalEventDimensionInvoker 查询异常事件用户分布
func (c *CloudRTCClient) ListRtcAbnormalEventDimensionInvoker(request *model.ListRtcAbnormalEventDimensionRequest) *ListRtcAbnormalEventDimensionInvoker {
	requestDef := GenReqDefForListRtcAbnormalEventDimension()
	return &ListRtcAbnormalEventDimensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcAbnormalEvents 查询用户异常体验事件
//
// 查询指定APP下通话的异常明细数据。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcAbnormalEvents(request *model.ListRtcAbnormalEventsRequest) (*model.ListRtcAbnormalEventsResponse, error) {
	requestDef := GenReqDefForListRtcAbnormalEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcAbnormalEventsResponse), nil
	}
}

// ListRtcAbnormalEventsInvoker 查询用户异常体验事件
func (c *CloudRTCClient) ListRtcAbnormalEventsInvoker(request *model.ListRtcAbnormalEventsRequest) *ListRtcAbnormalEventsInvoker {
	requestDef := GenReqDefForListRtcAbnormalEvents()
	return &ListRtcAbnormalEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcClientQosDetails 查询用户通话指标
//
// 查询用户通话质量指标数据。
//
// 可查询5天内的数据，mid 不为null，查询实时数据时，查询起止时间不超过24个小时，每次查询单个用户时，支持跨天查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcClientQosDetails(request *model.ListRtcClientQosDetailsRequest) (*model.ListRtcClientQosDetailsResponse, error) {
	requestDef := GenReqDefForListRtcClientQosDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcClientQosDetailsResponse), nil
	}
}

// ListRtcClientQosDetailsInvoker 查询用户通话指标
func (c *CloudRTCClient) ListRtcClientQosDetailsInvoker(request *model.ListRtcClientQosDetailsRequest) *ListRtcClientQosDetailsInvoker {
	requestDef := GenReqDefForListRtcClientQosDetails()
	return &ListRtcClientQosDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcHistoryQuality 查询历史质量
//
// 查询质量指标过去每天的体验数据，可查询最近31天的数据。当天未结束，无法查询到当天的体验数据。
//
// 最大查询跨度31天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcHistoryQuality(request *model.ListRtcHistoryQualityRequest) (*model.ListRtcHistoryQualityResponse, error) {
	requestDef := GenReqDefForListRtcHistoryQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryQualityResponse), nil
	}
}

// ListRtcHistoryQualityInvoker 查询历史质量
func (c *CloudRTCClient) ListRtcHistoryQualityInvoker(request *model.ListRtcHistoryQualityRequest) *ListRtcHistoryQualityInvoker {
	requestDef := GenReqDefForListRtcHistoryQuality()
	return &ListRtcHistoryQualityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcHistoryScale 查询历史规模
//
// 查询指标过去每天的规模数量，可查询最近31天的数据。当天未结束，无法查到当天的房间数与用户数。
//
// 最大查询跨度31天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcHistoryScale(request *model.ListRtcHistoryScaleRequest) (*model.ListRtcHistoryScaleResponse, error) {
	requestDef := GenReqDefForListRtcHistoryScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryScaleResponse), nil
	}
}

// ListRtcHistoryScaleInvoker 查询历史规模
func (c *CloudRTCClient) ListRtcHistoryScaleInvoker(request *model.ListRtcHistoryScaleRequest) *ListRtcHistoryScaleInvoker {
	requestDef := GenReqDefForListRtcHistoryScale()
	return &ListRtcHistoryScaleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcHistoryUsage 查询用量
//
// 查询过去的某一时间段内各种业务的用量数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcHistoryUsage(request *model.ListRtcHistoryUsageRequest) (*model.ListRtcHistoryUsageResponse, error) {
	requestDef := GenReqDefForListRtcHistoryUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryUsageResponse), nil
	}
}

// ListRtcHistoryUsageInvoker 查询用量
func (c *CloudRTCClient) ListRtcHistoryUsageInvoker(request *model.ListRtcHistoryUsageRequest) *ListRtcHistoryUsageInvoker {
	requestDef := GenReqDefForListRtcHistoryUsage()
	return &ListRtcHistoryUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcRealtimeNetwork 查询实时网络
//
// 获取实时网络数据相关指标在某一时间段内每分钟的统计数据。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcRealtimeNetwork(request *model.ListRtcRealtimeNetworkRequest) (*model.ListRtcRealtimeNetworkResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeNetworkResponse), nil
	}
}

// ListRtcRealtimeNetworkInvoker 查询实时网络
func (c *CloudRTCClient) ListRtcRealtimeNetworkInvoker(request *model.ListRtcRealtimeNetworkRequest) *ListRtcRealtimeNetworkInvoker {
	requestDef := GenReqDefForListRtcRealtimeNetwork()
	return &ListRtcRealtimeNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcRealtimeQuality 查询实时质量数据
//
// 获取实时质量数据的相关指标在某一时间段内每分钟的统计数据。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcRealtimeQuality(request *model.ListRtcRealtimeQualityRequest) (*model.ListRtcRealtimeQualityResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeQualityResponse), nil
	}
}

// ListRtcRealtimeQualityInvoker 查询实时质量数据
func (c *CloudRTCClient) ListRtcRealtimeQualityInvoker(request *model.ListRtcRealtimeQualityRequest) *ListRtcRealtimeQualityInvoker {
	requestDef := GenReqDefForListRtcRealtimeQuality()
	return &ListRtcRealtimeQualityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcRealtimeScale 查询实时规模
//
// 获取规模相关的指标在某一时间段内每分钟的统计数据。
//
// 最大查询跨度1天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcRealtimeScale(request *model.ListRtcRealtimeScaleRequest) (*model.ListRtcRealtimeScaleResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeScaleResponse), nil
	}
}

// ListRtcRealtimeScaleInvoker 查询实时规模
func (c *CloudRTCClient) ListRtcRealtimeScaleInvoker(request *model.ListRtcRealtimeScaleRequest) *ListRtcRealtimeScaleInvoker {
	requestDef := GenReqDefForListRtcRealtimeScale()
	return &ListRtcRealtimeScaleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcRealtimeScaleDimension 查询实时规模分布
//
// 对规模相关的数据，根据指定维度按在线用户数排名，获取规模相关的指标在指定维度下的统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcRealtimeScaleDimension(request *model.ListRtcRealtimeScaleDimensionRequest) (*model.ListRtcRealtimeScaleDimensionResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeScaleDimension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeScaleDimensionResponse), nil
	}
}

// ListRtcRealtimeScaleDimensionInvoker 查询实时规模分布
func (c *CloudRTCClient) ListRtcRealtimeScaleDimensionInvoker(request *model.ListRtcRealtimeScaleDimensionRequest) *ListRtcRealtimeScaleDimensionInvoker {
	requestDef := GenReqDefForListRtcRealtimeScaleDimension()
	return &ListRtcRealtimeScaleDimensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcRoomList 查询房间列表
//
// 指定事件范围查询这段期间创建的房间列表。
//
// 最大查询跨度90天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcRoomList(request *model.ListRtcRoomListRequest) (*model.ListRtcRoomListResponse, error) {
	requestDef := GenReqDefForListRtcRoomList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRoomListResponse), nil
	}
}

// ListRtcRoomListInvoker 查询房间列表
func (c *CloudRTCClient) ListRtcRoomListInvoker(request *model.ListRtcRoomListRequest) *ListRtcRoomListInvoker {
	requestDef := GenReqDefForListRtcRoomList()
	return &ListRtcRoomListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRtcUserList 查询用户列表
//
// 指定事件范围查询这段期间加入房间的用户列表。
//
// 最大查询跨度90天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudRTCClient) ListRtcUserList(request *model.ListRtcUserListRequest) (*model.ListRtcUserListResponse, error) {
	requestDef := GenReqDefForListRtcUserList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcUserListResponse), nil
	}
}

// ListRtcUserListInvoker 查询用户列表
func (c *CloudRTCClient) ListRtcUserListInvoker(request *model.ListRtcUserListRequest) *ListRtcUserListInvoker {
	requestDef := GenReqDefForListRtcUserList()
	return &ListRtcUserListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
