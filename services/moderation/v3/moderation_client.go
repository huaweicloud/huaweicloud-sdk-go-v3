package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v3/model"
)

type ModerationClient struct {
	HcClient *http_client.HcHttpClient
}

func NewModerationClient(hcClient *http_client.HcHttpClient) *ModerationClient {
	return &ModerationClient{HcClient: hcClient}
}

func ModerationClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// RunCreateAudioModerationJob 创建音频内容审核作业
//
// 分析并识别用户上传的音频内容是否有敏感内容（如色情、政治等），并将识别结果返回给用户
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunCreateAudioModerationJob(request *model.RunCreateAudioModerationJobRequest) (*model.RunCreateAudioModerationJobResponse, error) {
	requestDef := GenReqDefForRunCreateAudioModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateAudioModerationJobResponse), nil
	}
}

// RunCreateAudioModerationJobInvoker 创建音频内容审核作业
func (c *ModerationClient) RunCreateAudioModerationJobInvoker(request *model.RunCreateAudioModerationJobRequest) *RunCreateAudioModerationJobInvoker {
	requestDef := GenReqDefForRunCreateAudioModerationJob()
	return &RunCreateAudioModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQueryAudioModerationJob 查询音频内容审核作业
//
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunQueryAudioModerationJob(request *model.RunQueryAudioModerationJobRequest) (*model.RunQueryAudioModerationJobResponse, error) {
	requestDef := GenReqDefForRunQueryAudioModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryAudioModerationJobResponse), nil
	}
}

// RunQueryAudioModerationJobInvoker 查询音频内容审核作业
func (c *ModerationClient) RunQueryAudioModerationJobInvoker(request *model.RunQueryAudioModerationJobRequest) *RunQueryAudioModerationJobInvoker {
	requestDef := GenReqDefForRunQueryAudioModerationJob()
	return &RunQueryAudioModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTextModeration 文本内容审核
//
// 分析并识别用户上传的文本内容是否有敏感内容（如色情、政治等），并将识别结果返回给用户
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunTextModeration(request *model.RunTextModerationRequest) (*model.RunTextModerationResponse, error) {
	requestDef := GenReqDefForRunTextModeration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTextModerationResponse), nil
	}
}

// RunTextModerationInvoker 文本内容审核
func (c *ModerationClient) RunTextModerationInvoker(request *model.RunTextModerationRequest) *RunTextModerationInvoker {
	requestDef := GenReqDefForRunTextModeration()
	return &RunTextModerationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
