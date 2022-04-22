package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sis/v1/model"
)

type SisClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSisClient(hcClient *http_client.HcHttpClient) *SisClient {
	return &SisClient{HcClient: hcClient}
}

func SisClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 获取录音文件识别结果
//
// 该接口用于获取录音文件识别结果及识别状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) CollectTranscriberJob(request *model.CollectTranscriberJobRequest) (*model.CollectTranscriberJobResponse, error) {
	requestDef := GenReqDefForCollectTranscriberJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectTranscriberJobResponse), nil
	}
}

// 创建热词表
//
// 新建一个热词表，创建成功返回id。每个用户限制创建10个热词表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) CreateVocabulary(request *model.CreateVocabularyRequest) (*model.CreateVocabularyResponse, error) {
	requestDef := GenReqDefForCreateVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVocabularyResponse), nil
	}
}

// 删除热词表
//
// 通过热词表id删除热词表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) DeleteVocabulary(request *model.DeleteVocabularyRequest) (*model.DeleteVocabularyResponse, error) {
	requestDef := GenReqDefForDeleteVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVocabularyResponse), nil
	}
}

// 提交录音文件识别任务
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) PushTranscriberJobs(request *model.PushTranscriberJobsRequest) (*model.PushTranscriberJobsResponse, error) {
	requestDef := GenReqDefForPushTranscriberJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushTranscriberJobsResponse), nil
	}
}

// 录音文件识别极速版
//
// 极速版ASR(Restful API 接口, 适用于音频(文件大小&lt;&#x3D;100M,语音时长&lt;&#x3D;30分钟)文件的同步识别。
// 此接口以POST方式一次性上传整个音频或从华为OBS中下载音频， 识别结果将在请求响应中即刻返回，用于语音文件极速转写，质检分析的离线场景。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) RecognizeFlashAsr(request *model.RecognizeFlashAsrRequest) (*model.RecognizeFlashAsrResponse, error) {
	requestDef := GenReqDefForRecognizeFlashAsr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeFlashAsrResponse), nil
	}
}

// 一句话识别
//
// 一句话识别接口，用于短语音的同步识别。一次性上传整个音频，响应中即返回识别结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) RecognizeShortAudio(request *model.RecognizeShortAudioRequest) (*model.RecognizeShortAudioResponse, error) {
	requestDef := GenReqDefForRecognizeShortAudio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeShortAudioResponse), nil
	}
}

// 语音评测
//
// 口语评测接口，基于一小段朗读语音和预期文本，评价朗读者发音质量。当前仅支持华北-北京四。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) RunAudioAssessment(request *model.RunAudioAssessmentRequest) (*model.RunAudioAssessmentResponse, error) {
	requestDef := GenReqDefForRunAudioAssessment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAudioAssessmentResponse), nil
	}
}

// 多模态评测
//
// 多模态评测接口，根据朗读视频数据、视频对应的音频数据和试题文本，综合给出朗读者口语的评测分数。当前仅支持华北-北京四。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) RunMultiModalAssessment(request *model.RunMultiModalAssessmentRequest) (*model.RunMultiModalAssessmentResponse, error) {
	requestDef := GenReqDefForRunMultiModalAssessment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunMultiModalAssessmentResponse), nil
	}
}

// 语音合成
//
// 语音合成，是一种将文本转换成逼真语音的服务。用户通过实时访问和调用API获取语音合成结果，将用户输入的文字合成为音频。通过音色选择、自定义音量、语速，为企业和个人提供个性化的发音服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) RunTts(request *model.RunTtsRequest) (*model.RunTtsResponse, error) {
	requestDef := GenReqDefForRunTts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTtsResponse), nil
	}
}

// 查询热词表列表
//
// 查询用户所有热词表列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) ShowVocabularies(request *model.ShowVocabulariesRequest) (*model.ShowVocabulariesResponse, error) {
	requestDef := GenReqDefForShowVocabularies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVocabulariesResponse), nil
	}
}

// 查询热词表信息
//
// 通过热词表id查询热词表的信息和内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) ShowVocabulary(request *model.ShowVocabularyRequest) (*model.ShowVocabularyResponse, error) {
	requestDef := GenReqDefForShowVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVocabularyResponse), nil
	}
}

// 更新热词表
//
// 更新一个热词表，更新成功返回id。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SisClient) UpdateVocabulary(request *model.UpdateVocabularyRequest) (*model.UpdateVocabularyResponse, error) {
	requestDef := GenReqDefForUpdateVocabulary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVocabularyResponse), nil
	}
}
