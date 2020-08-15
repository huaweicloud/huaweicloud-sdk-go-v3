package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v1/model"
)

type LiveAPIClient struct {
	hcClient *http_client.HcHttpClient
}

func NewLiveAPIClient(hcClient *http_client.HcHttpClient) *LiveAPIClient {
	return &LiveAPIClient{hcClient: hcClient}
}

func LiveAPIClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建录制配置接口
func (c *LiveAPIClient) CreateRecordConfig(request *model.CreateRecordConfigRequest) (*model.CreateRecordConfigResponse, error) {
	requestDef := GenReqDefForCreateRecordConfig(request)
	resp, responseDef := GenRespForCreateRecordConfig()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//禁止直播推流
func (c *LiveAPIClient) CreateStreamForbidden(request *model.CreateStreamForbiddenRequest) (*model.CreateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForCreateStreamForbidden(request)
	resp, responseDef := GenRespForCreateStreamForbidden()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建直播转码模板
func (c *LiveAPIClient) CreateTranscodingsTemplate(request *model.CreateTranscodingsTemplateRequest) (*model.CreateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForCreateTranscodingsTemplate(request)
	resp, responseDef := GenRespForCreateTranscodingsTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除录制配置接口
func (c *LiveAPIClient) DeleteRecordConfig(request *model.DeleteRecordConfigRequest) (*model.DeleteRecordConfigResponse, error) {
	requestDef := GenReqDefForDeleteRecordConfig(request)
	resp, responseDef := GenRespForDeleteRecordConfig()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//恢复直播推流接口
func (c *LiveAPIClient) DeleteStreamForbidden(request *model.DeleteStreamForbiddenRequest) (*model.DeleteStreamForbiddenResponse, error) {
	requestDef := GenReqDefForDeleteStreamForbidden(request)
	resp, responseDef := GenRespForDeleteStreamForbidden()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除直播转码模板
func (c *LiveAPIClient) DeleteTranscodingsTemplate(request *model.DeleteTranscodingsTemplateRequest) (*model.DeleteTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingsTemplate(request)
	resp, responseDef := GenRespForDeleteTranscodingsTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询录制配置接口
func (c *LiveAPIClient) ListRecordConfigs(request *model.ListRecordConfigsRequest) (*model.ListRecordConfigsResponse, error) {
	requestDef := GenReqDefForListRecordConfigs(request)
	resp, responseDef := GenRespForListRecordConfigs()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询禁播黑名单列表
func (c *LiveAPIClient) ListStreamForbidden(request *model.ListStreamForbiddenRequest) (*model.ListStreamForbiddenResponse, error) {
	requestDef := GenReqDefForListStreamForbidden(request)
	resp, responseDef := GenRespForListStreamForbidden()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询直播加速的播流域名网络带宽监控数据
func (c *LiveAPIClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth(request)
	resp, responseDef := GenRespForShowBandwidth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询加速的直播播放在线人数
func (c *LiveAPIClient) ShowOnlineUsers(request *model.ShowOnlineUsersRequest) (*model.ShowOnlineUsersResponse, error) {
	requestDef := GenReqDefForShowOnlineUsers(request)
	resp, responseDef := GenRespForShowOnlineUsers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询直播加速的播流域名网络流量监控数据
func (c *LiveAPIClient) ShowTraffic(request *model.ShowTrafficRequest) (*model.ShowTrafficResponse, error) {
	requestDef := GenReqDefForShowTraffic(request)
	resp, responseDef := GenRespForShowTraffic()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询直播转码模板
func (c *LiveAPIClient) ShowTranscodingsTemplate(request *model.ShowTranscodingsTemplateRequest) (*model.ShowTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForShowTranscodingsTemplate(request)
	resp, responseDef := GenRespForShowTranscodingsTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//修改禁推属性
func (c *LiveAPIClient) UpdateStreamForbidden(request *model.UpdateStreamForbiddenRequest) (*model.UpdateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForUpdateStreamForbidden(request)
	resp, responseDef := GenRespForUpdateStreamForbidden()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//修改直播转码模板
func (c *LiveAPIClient) UpdateTranscodingsTemplate(request *model.UpdateTranscodingsTemplateRequest) (*model.UpdateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTranscodingsTemplate(request)
	resp, responseDef := GenRespForUpdateTranscodingsTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
