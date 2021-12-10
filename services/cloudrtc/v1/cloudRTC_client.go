package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

//查询用户通话质量指标数据。  可查询5天内的数据，mid 不为null，查询实时数据时，查询起止时间不超过24个小时，每次查询单个用户时，支持跨天查询。
func (c *CloudRTCClient) ListRtcClientQosDetails(request *model.ListRtcClientQosDetailsRequest) (*model.ListRtcClientQosDetailsResponse, error) {
	requestDef := GenReqDefForListRtcClientQosDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcClientQosDetailsResponse), nil
	}
}

//查询质量指标过去每天的体验数据，可查询最近31天的数据。当天未结束，无法查询到当天的体验数据
func (c *CloudRTCClient) ListRtcHistoryQuality(request *model.ListRtcHistoryQualityRequest) (*model.ListRtcHistoryQualityResponse, error) {
	requestDef := GenReqDefForListRtcHistoryQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryQualityResponse), nil
	}
}

//查询指标过去每天的规模数量，可查询最近31天的数据。当天未结束，无法查到当天的房间数与用户数
func (c *CloudRTCClient) ListRtcHistoryScale(request *model.ListRtcHistoryScaleRequest) (*model.ListRtcHistoryScaleResponse, error) {
	requestDef := GenReqDefForListRtcHistoryScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryScaleResponse), nil
	}
}

//查询过去的某一时间段内各种业务的用量数据
func (c *CloudRTCClient) ListRtcHistoryUsage(request *model.ListRtcHistoryUsageRequest) (*model.ListRtcHistoryUsageResponse, error) {
	requestDef := GenReqDefForListRtcHistoryUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcHistoryUsageResponse), nil
	}
}

//获取质量数据相关指标在某一时间段内每分钟的统计数据
func (c *CloudRTCClient) ListRtcRealtimeNetwork(request *model.ListRtcRealtimeNetworkRequest) (*model.ListRtcRealtimeNetworkResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeNetworkResponse), nil
	}
}

//获取实时质量数据的相关指标在某一时间段内每分钟的统计数据
func (c *CloudRTCClient) ListRtcRealtimeQuality(request *model.ListRtcRealtimeQualityRequest) (*model.ListRtcRealtimeQualityResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeQualityResponse), nil
	}
}

//获取规模相关的指标在某一时间段内每分钟的统计数据
func (c *CloudRTCClient) ListRtcRealtimeScale(request *model.ListRtcRealtimeScaleRequest) (*model.ListRtcRealtimeScaleResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeScaleResponse), nil
	}
}

//对规模相关的数据，根据指定维度按在线用户数排名，获取规模相关的指标在指定维度下的统计数据
func (c *CloudRTCClient) ListRtcRealtimeScaleDimension(request *model.ListRtcRealtimeScaleDimensionRequest) (*model.ListRtcRealtimeScaleDimensionResponse, error) {
	requestDef := GenReqDefForListRtcRealtimeScaleDimension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRealtimeScaleDimensionResponse), nil
	}
}

//指定时间范围查询这段期间创建的房间列表
func (c *CloudRTCClient) ListRtcRoomList(request *model.ListRtcRoomListRequest) (*model.ListRtcRoomListResponse, error) {
	requestDef := GenReqDefForListRtcRoomList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcRoomListResponse), nil
	}
}

//指定时间范围查询这段期间加入房间的用户列表
func (c *CloudRTCClient) ListRtcUserList(request *model.ListRtcUserListRequest) (*model.ListRtcUserListResponse, error) {
	requestDef := GenReqDefForListRtcUserList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRtcUserListResponse), nil
	}
}
