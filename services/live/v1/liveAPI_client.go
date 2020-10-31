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
	requestDef := GenReqDefForCreateRecordConfig()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordConfigResponse), nil
	}
}

//禁止直播推流
func (c *LiveAPIClient) CreateStreamForbidden(request *model.CreateStreamForbiddenRequest) (*model.CreateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForCreateStreamForbidden()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamForbiddenResponse), nil
	}
}

//创建直播转码模板
func (c *LiveAPIClient) CreateTranscodingsTemplate(request *model.CreateTranscodingsTemplateRequest) (*model.CreateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForCreateTranscodingsTemplate()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTranscodingsTemplateResponse), nil
	}
}

//删除录制配置接口
func (c *LiveAPIClient) DeleteRecordConfig(request *model.DeleteRecordConfigRequest) (*model.DeleteRecordConfigResponse, error) {
	requestDef := GenReqDefForDeleteRecordConfig()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordConfigResponse), nil
	}
}

//恢复直播推流接口
func (c *LiveAPIClient) DeleteStreamForbidden(request *model.DeleteStreamForbiddenRequest) (*model.DeleteStreamForbiddenResponse, error) {
	requestDef := GenReqDefForDeleteStreamForbidden()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStreamForbiddenResponse), nil
	}
}

//删除直播转码模板
func (c *LiveAPIClient) DeleteTranscodingsTemplate(request *model.DeleteTranscodingsTemplateRequest) (*model.DeleteTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingsTemplate()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTranscodingsTemplateResponse), nil
	}
}

//查询录制配置接口
func (c *LiveAPIClient) ListRecordConfigs(request *model.ListRecordConfigsRequest) (*model.ListRecordConfigsResponse, error) {
	requestDef := GenReqDefForListRecordConfigs()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordConfigsResponse), nil
	}
}

//查询禁播黑名单列表
func (c *LiveAPIClient) ListStreamForbidden(request *model.ListStreamForbiddenRequest) (*model.ListStreamForbiddenResponse, error) {
	requestDef := GenReqDefForListStreamForbidden()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStreamForbiddenResponse), nil
	}
}

//查询直播加速的播流域名网络带宽监控数据
func (c *LiveAPIClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthResponse), nil
	}
}

//查询加速的直播播放在线人数
func (c *LiveAPIClient) ShowOnlineUsers(request *model.ShowOnlineUsersRequest) (*model.ShowOnlineUsersResponse, error) {
	requestDef := GenReqDefForShowOnlineUsers()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOnlineUsersResponse), nil
	}
}

//查询直播加速的播流域名网络流量监控数据
func (c *LiveAPIClient) ShowTraffic(request *model.ShowTrafficRequest) (*model.ShowTrafficResponse, error) {
	requestDef := GenReqDefForShowTraffic()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrafficResponse), nil
	}
}

//查询直播转码模板
func (c *LiveAPIClient) ShowTranscodingsTemplate(request *model.ShowTranscodingsTemplateRequest) (*model.ShowTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForShowTranscodingsTemplate()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTranscodingsTemplateResponse), nil
	}
}

//修改禁推属性
func (c *LiveAPIClient) UpdateStreamForbidden(request *model.UpdateStreamForbiddenRequest) (*model.UpdateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForUpdateStreamForbidden()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStreamForbiddenResponse), nil
	}
}

//修改直播转码模板
func (c *LiveAPIClient) UpdateTranscodingsTemplate(request *model.UpdateTranscodingsTemplateRequest) (*model.UpdateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTranscodingsTemplate()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTranscodingsTemplateResponse), nil
	}
}
