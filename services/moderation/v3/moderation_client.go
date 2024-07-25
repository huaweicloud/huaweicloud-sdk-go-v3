package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v3/model"
)

type ModerationClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewModerationClient(hcClient *httpclient.HcHttpClient) *ModerationClient {
	return &ModerationClient{HcClient: hcClient}
}

func ModerationClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCheckImageSync 图像审核批量同步接口
//
// 图像审核批量同步接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) BatchCheckImageSync(request *model.BatchCheckImageSyncRequest) (*model.BatchCheckImageSyncResponse, error) {
	requestDef := GenReqDefForBatchCheckImageSync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckImageSyncResponse), nil
	}
}

// BatchCheckImageSyncInvoker 图像审核批量同步接口
func (c *ModerationClient) BatchCheckImageSyncInvoker(request *model.BatchCheckImageSyncRequest) *BatchCheckImageSyncInvoker {
	requestDef := GenReqDefForBatchCheckImageSync()
	return &BatchCheckImageSyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckImageModeration 图像内容审核
//
// 分析并识别用户上传的图像内容是否有敏感内容（如涉及暴恐元素、涉黄内容等），并将识别结果返回给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) CheckImageModeration(request *model.CheckImageModerationRequest) (*model.CheckImageModerationResponse, error) {
	requestDef := GenReqDefForCheckImageModeration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckImageModerationResponse), nil
	}
}

// CheckImageModerationInvoker 图像内容审核
func (c *ModerationClient) CheckImageModerationInvoker(request *model.CheckImageModerationRequest) *CheckImageModerationInvoker {
	requestDef := GenReqDefForCheckImageModeration()
	return &CheckImageModerationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCloseAudioStreamModerationJob 关闭音频流内容审核作业
//
// 关闭音频流内容审核作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCloseAudioStreamModerationJob(request *model.RunCloseAudioStreamModerationJobRequest) (*model.RunCloseAudioStreamModerationJobResponse, error) {
	requestDef := GenReqDefForRunCloseAudioStreamModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCloseAudioStreamModerationJobResponse), nil
	}
}

// RunCloseAudioStreamModerationJobInvoker 关闭音频流内容审核作业
func (c *ModerationClient) RunCloseAudioStreamModerationJobInvoker(request *model.RunCloseAudioStreamModerationJobRequest) *RunCloseAudioStreamModerationJobInvoker {
	requestDef := GenReqDefForRunCloseAudioStreamModerationJob()
	return &RunCloseAudioStreamModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCloseVideoStreamModerationJob 关闭视频流内容审核作业
//
// 关闭视频流审核接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCloseVideoStreamModerationJob(request *model.RunCloseVideoStreamModerationJobRequest) (*model.RunCloseVideoStreamModerationJobResponse, error) {
	requestDef := GenReqDefForRunCloseVideoStreamModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCloseVideoStreamModerationJobResponse), nil
	}
}

// RunCloseVideoStreamModerationJobInvoker 关闭视频流内容审核作业
func (c *ModerationClient) RunCloseVideoStreamModerationJobInvoker(request *model.RunCloseVideoStreamModerationJobRequest) *RunCloseVideoStreamModerationJobInvoker {
	requestDef := GenReqDefForRunCloseVideoStreamModerationJob()
	return &RunCloseVideoStreamModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCreateAudioModerationJob 创建音频内容审核作业
//
// 分析并识别用户上传的音频内容是否有敏感内容（如色情、政治等），并将识别结果返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// RunCreateAudioStreamModerationJob 创建音频流内容审核作业
//
// 创建音频流内容审核作业，创建成功会将作业ID返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCreateAudioStreamModerationJob(request *model.RunCreateAudioStreamModerationJobRequest) (*model.RunCreateAudioStreamModerationJobResponse, error) {
	requestDef := GenReqDefForRunCreateAudioStreamModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateAudioStreamModerationJobResponse), nil
	}
}

// RunCreateAudioStreamModerationJobInvoker 创建音频流内容审核作业
func (c *ModerationClient) RunCreateAudioStreamModerationJobInvoker(request *model.RunCreateAudioStreamModerationJobRequest) *RunCreateAudioStreamModerationJobInvoker {
	requestDef := GenReqDefForRunCreateAudioStreamModerationJob()
	return &RunCreateAudioStreamModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCreateDocumentModerationJob 创建文档内容审核作业
//
// 创建文档内容审核作业，创建成功会将作业ID返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCreateDocumentModerationJob(request *model.RunCreateDocumentModerationJobRequest) (*model.RunCreateDocumentModerationJobResponse, error) {
	requestDef := GenReqDefForRunCreateDocumentModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateDocumentModerationJobResponse), nil
	}
}

// RunCreateDocumentModerationJobInvoker 创建文档内容审核作业
func (c *ModerationClient) RunCreateDocumentModerationJobInvoker(request *model.RunCreateDocumentModerationJobRequest) *RunCreateDocumentModerationJobInvoker {
	requestDef := GenReqDefForRunCreateDocumentModerationJob()
	return &RunCreateDocumentModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCreateVideoModerationJob 创建视频内容审核作业
//
// 创建视频内容审核作业，创建成功会将作业ID返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCreateVideoModerationJob(request *model.RunCreateVideoModerationJobRequest) (*model.RunCreateVideoModerationJobResponse, error) {
	requestDef := GenReqDefForRunCreateVideoModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateVideoModerationJobResponse), nil
	}
}

// RunCreateVideoModerationJobInvoker 创建视频内容审核作业
func (c *ModerationClient) RunCreateVideoModerationJobInvoker(request *model.RunCreateVideoModerationJobRequest) *RunCreateVideoModerationJobInvoker {
	requestDef := GenReqDefForRunCreateVideoModerationJob()
	return &RunCreateVideoModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCreateVideoStreamModerationJob 创建视频流内容审核作业
//
// 创建视频流内容审核作业，创建成功会将作业ID返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunCreateVideoStreamModerationJob(request *model.RunCreateVideoStreamModerationJobRequest) (*model.RunCreateVideoStreamModerationJobResponse, error) {
	requestDef := GenReqDefForRunCreateVideoStreamModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateVideoStreamModerationJobResponse), nil
	}
}

// RunCreateVideoStreamModerationJobInvoker 创建视频流内容审核作业
func (c *ModerationClient) RunCreateVideoStreamModerationJobInvoker(request *model.RunCreateVideoStreamModerationJobRequest) *RunCreateVideoStreamModerationJobInvoker {
	requestDef := GenReqDefForRunCreateVideoStreamModerationJob()
	return &RunCreateVideoStreamModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQueryAudioModerationJob 查询音频内容审核作业
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// RunQueryDocumentModerationJob 查询文档审核作业结果
//
// 查询文档审核结果接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunQueryDocumentModerationJob(request *model.RunQueryDocumentModerationJobRequest) (*model.RunQueryDocumentModerationJobResponse, error) {
	requestDef := GenReqDefForRunQueryDocumentModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryDocumentModerationJobResponse), nil
	}
}

// RunQueryDocumentModerationJobInvoker 查询文档审核作业结果
func (c *ModerationClient) RunQueryDocumentModerationJobInvoker(request *model.RunQueryDocumentModerationJobRequest) *RunQueryDocumentModerationJobInvoker {
	requestDef := GenReqDefForRunQueryDocumentModerationJob()
	return &RunQueryDocumentModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQueryVideoModerationJob 查询视频内容审核作业
//
// 查询视频审核作业处理状态与结果，并将识别结果返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModerationClient) RunQueryVideoModerationJob(request *model.RunQueryVideoModerationJobRequest) (*model.RunQueryVideoModerationJobResponse, error) {
	requestDef := GenReqDefForRunQueryVideoModerationJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryVideoModerationJobResponse), nil
	}
}

// RunQueryVideoModerationJobInvoker 查询视频内容审核作业
func (c *ModerationClient) RunQueryVideoModerationJobInvoker(request *model.RunQueryVideoModerationJobRequest) *RunQueryVideoModerationJobInvoker {
	requestDef := GenReqDefForRunQueryVideoModerationJob()
	return &RunQueryVideoModerationJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTextModeration 文本内容审核
//
// 分析并识别用户上传的文本内容是否有敏感内容（如色情、政治等），并将识别结果返回给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
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
