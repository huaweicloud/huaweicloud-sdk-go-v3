package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sis/v1/model"
)

type SisClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSisClient(hcClient *httpclient.HcHttpClient) *SisClient {
	return &SisClient{HcClient: hcClient}
}

func SisClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CollectTranscriberJob 获取录音文件识别结果
//
// 该接口用于获取录音文件识别结果及识别状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) CollectTranscriberJob(request *model.CollectTranscriberJobRequest) (*model.CollectTranscriberJobResponse, error) {
	requestDef := GenReqDefForCollectTranscriberJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectTranscriberJobResponse), nil
	}
}

// CollectTranscriberJobInvoker 获取录音文件识别结果
func (c *SisClient) CollectTranscriberJobInvoker(request *model.CollectTranscriberJobRequest) *CollectTranscriberJobInvoker {
	requestDef := GenReqDefForCollectTranscriberJob()
	return &CollectTranscriberJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVocabulary 创建热词表
//
// 新建一个热词表，创建成功返回id。每个用户限制创建10个热词表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) CreateVocabulary(request *model.CreateVocabularyRequest) (*model.CreateVocabularyResponse, error) {
	requestDef := GenReqDefForCreateVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVocabularyResponse), nil
	}
}

// CreateVocabularyInvoker 创建热词表
func (c *SisClient) CreateVocabularyInvoker(request *model.CreateVocabularyRequest) *CreateVocabularyInvoker {
	requestDef := GenReqDefForCreateVocabulary()
	return &CreateVocabularyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVoice 注册接口
//
// 客户上传一段录音，并指定voice_name，在系统中注册声音。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) CreateVoice(request *model.CreateVoiceRequest) (*model.CreateVoiceResponse, error) {
	requestDef := GenReqDefForCreateVoice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVoiceResponse), nil
	}
}

// CreateVoiceInvoker 注册接口
func (c *SisClient) CreateVoiceInvoker(request *model.CreateVoiceRequest) *CreateVoiceInvoker {
	requestDef := GenReqDefForCreateVoice()
	return &CreateVoiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVocabulary 删除热词表
//
// 通过热词表id删除热词表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) DeleteVocabulary(request *model.DeleteVocabularyRequest) (*model.DeleteVocabularyResponse, error) {
	requestDef := GenReqDefForDeleteVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVocabularyResponse), nil
	}
}

// DeleteVocabularyInvoker 删除热词表
func (c *SisClient) DeleteVocabularyInvoker(request *model.DeleteVocabularyRequest) *DeleteVocabularyInvoker {
	requestDef := GenReqDefForDeleteVocabulary()
	return &DeleteVocabularyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GenerateSpeech 合成接口
//
// 用户指定一个声色名称，并指定对应的文本，合成对应的复刻的声音
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) GenerateSpeech(request *model.GenerateSpeechRequest) (*model.GenerateSpeechResponse, error) {
	requestDef := GenReqDefForGenerateSpeech()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GenerateSpeechResponse), nil
	}
}

// GenerateSpeechInvoker 合成接口
func (c *SisClient) GenerateSpeechInvoker(request *model.GenerateSpeechRequest) *GenerateSpeechInvoker {
	requestDef := GenReqDefForGenerateSpeech()
	return &GenerateSpeechInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVoices 查询接口
//
// 查询已注册的声音列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) ListVoices(request *model.ListVoicesRequest) (*model.ListVoicesResponse, error) {
	requestDef := GenReqDefForListVoices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVoicesResponse), nil
	}
}

// ListVoicesInvoker 查询接口
func (c *SisClient) ListVoicesInvoker(request *model.ListVoicesRequest) *ListVoicesInvoker {
	requestDef := GenReqDefForListVoices()
	return &ListVoicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushTranscriberJobs 提交录音文件识别任务
//
// **录音文件识别**
// 录音文件识别接口，用于识别长录音文件，录音文件放在华为云OBS（对象存储服务）上。
//
// 由于录音文件识别通常会需要较长的时间，因此识别是异步的，也即接口分为创建识别任务和查询任务状态两个接口。创建识别任务接口创建任务完成后返回，然后用户通过调用查询任务状态接口来获得转写状态和结果。
//
// **功能介绍**
// 该接口用于提交录音文件识别任务，其中录音文件保存在用户的OBS桶中。用户开通录音识别服务时，需授权录音文件引擎读取用户OBS桶权限。
//
// 接口约束
// 录音时长不超过5小时，文件大小不超过300M，识别结果保存72小时（从识别完成的时间算起）。72小时后如果再访问，将会返回 \&quot;task id is not found\&quot;错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) PushTranscriberJobs(request *model.PushTranscriberJobsRequest) (*model.PushTranscriberJobsResponse, error) {
	requestDef := GenReqDefForPushTranscriberJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushTranscriberJobsResponse), nil
	}
}

// PushTranscriberJobsInvoker 提交录音文件识别任务
func (c *SisClient) PushTranscriberJobsInvoker(request *model.PushTranscriberJobsRequest) *PushTranscriberJobsInvoker {
	requestDef := GenReqDefForPushTranscriberJobs()
	return &PushTranscriberJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeFlashAsr 录音文件识别极速版
//
// 极速版ASR(Restful API 接口, 适用于音频(文件大小&lt;&#x3D;100M,语音时长&lt;&#x3D;30分钟)文件的同步识别。
// 此接口以POST方式一次性上传整个音频或从华为OBS中下载音频， 识别结果将在请求响应中即刻返回，用于语音文件极速转写，质检分析的离线场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) RecognizeFlashAsr(request *model.RecognizeFlashAsrRequest) (*model.RecognizeFlashAsrResponse, error) {
	requestDef := GenReqDefForRecognizeFlashAsr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFlashAsrResponse), nil
	}
}

// RecognizeFlashAsrInvoker 录音文件识别极速版
func (c *SisClient) RecognizeFlashAsrInvoker(request *model.RecognizeFlashAsrRequest) *RecognizeFlashAsrInvoker {
	requestDef := GenReqDefForRecognizeFlashAsr()
	return &RecognizeFlashAsrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeShortAudio 一句话识别
//
// 一句话识别接口，用于短语音的同步识别。一次性上传整个音频，响应中即返回识别结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) RecognizeShortAudio(request *model.RecognizeShortAudioRequest) (*model.RecognizeShortAudioResponse, error) {
	requestDef := GenReqDefForRecognizeShortAudio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeShortAudioResponse), nil
	}
}

// RecognizeShortAudioInvoker 一句话识别
func (c *SisClient) RecognizeShortAudioInvoker(request *model.RecognizeShortAudioRequest) *RecognizeShortAudioInvoker {
	requestDef := GenReqDefForRecognizeShortAudio()
	return &RecognizeShortAudioInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTts 语音合成
//
// 语音合成，是一种将文本转换成逼真语音的服务。用户通过实时访问和调用API获取语音合成结果，将用户输入的文字合成为音频。通过音色选择、自定义音量、语速，为企业和个人提供个性化的发音服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) RunTts(request *model.RunTtsRequest) (*model.RunTtsResponse, error) {
	requestDef := GenReqDefForRunTts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTtsResponse), nil
	}
}

// RunTtsInvoker 语音合成
func (c *SisClient) RunTtsInvoker(request *model.RunTtsRequest) *RunTtsInvoker {
	requestDef := GenReqDefForRunTts()
	return &RunTtsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVocabularies 查询热词表列表
//
// 查询用户所有热词表列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) ShowVocabularies(request *model.ShowVocabulariesRequest) (*model.ShowVocabulariesResponse, error) {
	requestDef := GenReqDefForShowVocabularies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVocabulariesResponse), nil
	}
}

// ShowVocabulariesInvoker 查询热词表列表
func (c *SisClient) ShowVocabulariesInvoker(request *model.ShowVocabulariesRequest) *ShowVocabulariesInvoker {
	requestDef := GenReqDefForShowVocabularies()
	return &ShowVocabulariesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVocabulary 查询热词表信息
//
// 通过热词表id查询热词表的信息和内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) ShowVocabulary(request *model.ShowVocabularyRequest) (*model.ShowVocabularyResponse, error) {
	requestDef := GenReqDefForShowVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVocabularyResponse), nil
	}
}

// ShowVocabularyInvoker 查询热词表信息
func (c *SisClient) ShowVocabularyInvoker(request *model.ShowVocabularyRequest) *ShowVocabularyInvoker {
	requestDef := GenReqDefForShowVocabulary()
	return &ShowVocabularyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVocabulary 更新热词表
//
// 更新一个热词表，更新成功返回id。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SisClient) UpdateVocabulary(request *model.UpdateVocabularyRequest) (*model.UpdateVocabularyResponse, error) {
	requestDef := GenReqDefForUpdateVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVocabularyResponse), nil
	}
}

// UpdateVocabularyInvoker 更新热词表
func (c *SisClient) UpdateVocabularyInvoker(request *model.UpdateVocabularyRequest) *UpdateVocabularyInvoker {
	requestDef := GenReqDefForUpdateVocabulary()
	return &UpdateVocabularyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
