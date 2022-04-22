package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v2/model"
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

// 处理结果查询
//
// 分析并识别用户上传的图像内容是否有敏感内容（如涉及政治人物、暴恐元素、涉黄内容等），并将识别结果返回给用户。
// &gt; 任务最长保留时间为30分钟，过期后会被清理掉。建议在任务提交后，每30s进行一次周期查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunCheckResult(request *model.RunCheckResultRequest) (*model.RunCheckResultResponse, error) {
	requestDef := GenReqDefForRunCheckResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCheckResultResponse), nil
	}
}

// 任务列表查询
//
// 查询批量图像内容审核任务列表，可通过指定任务状态查询来对任务列表进行过滤。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunCheckTaskJobs(request *model.RunCheckTaskJobsRequest) (*model.RunCheckTaskJobsResponse, error) {
	requestDef := GenReqDefForRunCheckTaskJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCheckTaskJobsResponse), nil
	}
}

// 图像内容审核（批量）
//
// 分析并识别用户上传的图像内容是否有敏感内容（如涉及政治人物、暴恐元素、涉黄内容等），并将识别结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunImageBatchModeration(request *model.RunImageBatchModerationRequest) (*model.RunImageBatchModerationResponse, error) {
	requestDef := GenReqDefForRunImageBatchModeration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageBatchModerationResponse), nil
	}
}

// 图像内容审核
//
// 分析并识别用户上传的图像内容是否有敏感内容（如涉及政治人物、暴恐元素、涉黄内容等），并将识别结果返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunImageModeration(request *model.RunImageModerationRequest) (*model.RunImageModerationResponse, error) {
	requestDef := GenReqDefForRunImageModeration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageModerationResponse), nil
	}
}

// 语音内容审核
//
// 分析并识别用户上传的语音内容是否有敏感内容（如色情、政治等），并将识别结果 返回给用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunModerationAudio(request *model.RunModerationAudioRequest) (*model.RunModerationAudioResponse, error) {
	requestDef := GenReqDefForRunModerationAudio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunModerationAudioResponse), nil
	}
}

// 任务提交
//
// 提交批量图像内容审核任务，返回任务标识，任务标识可用于查询任务结果。此接口为异步接口，相对于批量接口，支持更大图片列表批次。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ModerationClient) RunTaskSumbit(request *model.RunTaskSumbitRequest) (*model.RunTaskSumbitResponse, error) {
	requestDef := GenReqDefForRunTaskSumbit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTaskSumbitResponse), nil
	}
}

// 文本内容审核
//
// 分析并识别用户上传的文本内容是否有敏感内容（如色情、政治等），并将识别结果返回给用户。
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
