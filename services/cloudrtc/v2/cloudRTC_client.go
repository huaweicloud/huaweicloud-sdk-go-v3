package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudrtc/v2/model"
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

//调用此接口接口启动单流任务。
func (c *CloudRTCClient) CreateIndividualStreamJob(request *model.CreateIndividualStreamJobRequest) (*model.CreateIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForCreateIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndividualStreamJobResponse), nil
	}
}

//调用此接口创建合流转码任务。
func (c *CloudRTCClient) CreateMixJob(request *model.CreateMixJobRequest) (*model.CreateMixJobResponse, error) {
	requestDef := GenReqDefForCreateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMixJobResponse), nil
	}
}

//调用此接口查询单流任务状态。  租户的OBS桶内的情况，暂不支持查询。
func (c *CloudRTCClient) ShowIndividualStreamJob(request *model.ShowIndividualStreamJobRequest) (*model.ShowIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForShowIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndividualStreamJobResponse), nil
	}
}

//调用此接口查询合流转码任务状态。
func (c *CloudRTCClient) ShowMixJob(request *model.ShowMixJobRequest) (*model.ShowMixJobResponse, error) {
	requestDef := GenReqDefForShowMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMixJobResponse), nil
	}
}

//调用此接口停止单流任务
func (c *CloudRTCClient) StopIndividualStreamJob(request *model.StopIndividualStreamJobRequest) (*model.StopIndividualStreamJobResponse, error) {
	requestDef := GenReqDefForStopIndividualStreamJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopIndividualStreamJobResponse), nil
	}
}

//调用此接口停止已下发的合流转码任务。
func (c *CloudRTCClient) StopMixJob(request *model.StopMixJobRequest) (*model.StopMixJobResponse, error) {
	requestDef := GenReqDefForStopMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMixJobResponse), nil
	}
}

//调用此接口更新合流任务布局。
func (c *CloudRTCClient) UpdateMixJob(request *model.UpdateMixJobRequest) (*model.UpdateMixJobResponse, error) {
	requestDef := GenReqDefForUpdateMixJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMixJobResponse), nil
	}
}
