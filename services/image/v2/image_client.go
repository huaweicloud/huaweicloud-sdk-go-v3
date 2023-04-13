package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/model"
)

type ImageClient struct {
	HcClient *http_client.HcHttpClient
}

func NewImageClient(hcClient *http_client.HcHttpClient) *ImageClient {
	return &ImageClient{HcClient: hcClient}
}

func ImageClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateImageHighresolutionMattingTask 创建图像高清抠图任务
//
// 创建图像高清抠图任务，将输入的高清图像中的商品主体从原图中扣取出来，输出商品主体图片或者蒙版。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateImageHighresolutionMattingTask(request *model.CreateImageHighresolutionMattingTaskRequest) (*model.CreateImageHighresolutionMattingTaskResponse, error) {
	requestDef := GenReqDefForCreateImageHighresolutionMattingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageHighresolutionMattingTaskResponse), nil
	}
}

// CreateImageHighresolutionMattingTaskInvoker 创建图像高清抠图任务
func (c *ImageClient) CreateImageHighresolutionMattingTaskInvoker(request *model.CreateImageHighresolutionMattingTaskRequest) *CreateImageHighresolutionMattingTaskInvoker {
	requestDef := GenReqDefForCreateImageHighresolutionMattingTask()
	return &CreateImageHighresolutionMattingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImageTranslateTask 创建任务
//
// Create Task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateImageTranslateTask(request *model.CreateImageTranslateTaskRequest) (*model.CreateImageTranslateTaskResponse, error) {
	requestDef := GenReqDefForCreateImageTranslateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageTranslateTaskResponse), nil
	}
}

// CreateImageTranslateTaskInvoker 创建任务
func (c *ImageClient) CreateImageTranslateTaskInvoker(request *model.CreateImageTranslateTaskRequest) *CreateImageTranslateTaskInvoker {
	requestDef := GenReqDefForCreateImageTranslateTask()
	return &CreateImageTranslateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoCoverAnalysisTask 创建任务
//
// 创建视频封面任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateVideoCoverAnalysisTask(request *model.CreateVideoCoverAnalysisTaskRequest) (*model.CreateVideoCoverAnalysisTaskResponse, error) {
	requestDef := GenReqDefForCreateVideoCoverAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoCoverAnalysisTaskResponse), nil
	}
}

// CreateVideoCoverAnalysisTaskInvoker 创建任务
func (c *ImageClient) CreateVideoCoverAnalysisTaskInvoker(request *model.CreateVideoCoverAnalysisTaskRequest) *CreateVideoCoverAnalysisTaskInvoker {
	requestDef := GenReqDefForCreateVideoCoverAnalysisTask()
	return &CreateVideoCoverAnalysisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoObjectMaskingTask 创建视频脱敏任务
//
// 创建视频脱敏除任务，将输入视频的图像敏感信息模糊化，包括视频中的人脸和车牌
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateVideoObjectMaskingTask(request *model.CreateVideoObjectMaskingTaskRequest) (*model.CreateVideoObjectMaskingTaskResponse, error) {
	requestDef := GenReqDefForCreateVideoObjectMaskingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoObjectMaskingTaskResponse), nil
	}
}

// CreateVideoObjectMaskingTaskInvoker 创建视频脱敏任务
func (c *ImageClient) CreateVideoObjectMaskingTaskInvoker(request *model.CreateVideoObjectMaskingTaskRequest) *CreateVideoObjectMaskingTaskInvoker {
	requestDef := GenReqDefForCreateVideoObjectMaskingTask()
	return &CreateVideoObjectMaskingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoShotSplitTask 创建任务
//
// 创建视频拆条任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateVideoShotSplitTask(request *model.CreateVideoShotSplitTaskRequest) (*model.CreateVideoShotSplitTaskResponse, error) {
	requestDef := GenReqDefForCreateVideoShotSplitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoShotSplitTaskResponse), nil
	}
}

// CreateVideoShotSplitTaskInvoker 创建任务
func (c *ImageClient) CreateVideoShotSplitTaskInvoker(request *model.CreateVideoShotSplitTaskRequest) *CreateVideoShotSplitTaskInvoker {
	requestDef := GenReqDefForCreateVideoShotSplitTask()
	return &CreateVideoShotSplitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoSummarizationAnalysisTask 创建任务
//
// Create Task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateVideoSummarizationAnalysisTask(request *model.CreateVideoSummarizationAnalysisTaskRequest) (*model.CreateVideoSummarizationAnalysisTaskResponse, error) {
	requestDef := GenReqDefForCreateVideoSummarizationAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoSummarizationAnalysisTaskResponse), nil
	}
}

// CreateVideoSummarizationAnalysisTaskInvoker 创建任务
func (c *ImageClient) CreateVideoSummarizationAnalysisTaskInvoker(request *model.CreateVideoSummarizationAnalysisTaskRequest) *CreateVideoSummarizationAnalysisTaskInvoker {
	requestDef := GenReqDefForCreateVideoSummarizationAnalysisTask()
	return &CreateVideoSummarizationAnalysisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoTranslateTask 创建任务
//
// Create Task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) CreateVideoTranslateTask(request *model.CreateVideoTranslateTaskRequest) (*model.CreateVideoTranslateTaskResponse, error) {
	requestDef := GenReqDefForCreateVideoTranslateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoTranslateTaskResponse), nil
	}
}

// CreateVideoTranslateTaskInvoker 创建任务
func (c *ImageClient) CreateVideoTranslateTaskInvoker(request *model.CreateVideoTranslateTaskRequest) *CreateVideoTranslateTaskInvoker {
	requestDef := GenReqDefForCreateVideoTranslateTask()
	return &CreateVideoTranslateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCelebrityRecognition 名人识别
//
// 分析并识别图片中包含的政治人物、明星及网红人物，返回人物信息及人脸坐标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunCelebrityRecognition(request *model.RunCelebrityRecognitionRequest) (*model.RunCelebrityRecognitionResponse, error) {
	requestDef := GenReqDefForRunCelebrityRecognition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCelebrityRecognitionResponse), nil
	}
}

// RunCelebrityRecognitionInvoker 名人识别
func (c *ImageClient) RunCelebrityRecognitionInvoker(request *model.RunCelebrityRecognitionRequest) *RunCelebrityRecognitionInvoker {
	requestDef := GenReqDefForRunCelebrityRecognition()
	return &RunCelebrityRecognitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunDeleteCustomTags 删除媒资图像标签
//
// 用于用户删除自定义的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunDeleteCustomTags(request *model.RunDeleteCustomTagsRequest) (*model.RunDeleteCustomTagsResponse, error) {
	requestDef := GenReqDefForRunDeleteCustomTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDeleteCustomTagsResponse), nil
	}
}

// RunDeleteCustomTagsInvoker 删除媒资图像标签
func (c *ImageClient) RunDeleteCustomTagsInvoker(request *model.RunDeleteCustomTagsRequest) *RunDeleteCustomTagsInvoker {
	requestDef := GenReqDefForRunDeleteCustomTags()
	return &RunDeleteCustomTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageDescription 图像描述
//
// 图像描述
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageDescription(request *model.RunImageDescriptionRequest) (*model.RunImageDescriptionResponse, error) {
	requestDef := GenReqDefForRunImageDescription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageDescriptionResponse), nil
	}
}

// RunImageDescriptionInvoker 图像描述
func (c *ImageClient) RunImageDescriptionInvoker(request *model.RunImageDescriptionRequest) *RunImageDescriptionInvoker {
	requestDef := GenReqDefForRunImageDescription()
	return &RunImageDescriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageMainObjectDetection 主体识别
//
// 检测图像中的主要内容，返回主要内容的坐标信息，这里的主要内容包括两方面：bounding_box和main_object_box
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageMainObjectDetection(request *model.RunImageMainObjectDetectionRequest) (*model.RunImageMainObjectDetectionResponse, error) {
	requestDef := GenReqDefForRunImageMainObjectDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageMainObjectDetectionResponse), nil
	}
}

// RunImageMainObjectDetectionInvoker 主体识别
func (c *ImageClient) RunImageMainObjectDetectionInvoker(request *model.RunImageMainObjectDetectionRequest) *RunImageMainObjectDetectionInvoker {
	requestDef := GenReqDefForRunImageMainObjectDetection()
	return &RunImageMainObjectDetectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageMediaTagging 标签识别
//
// 自然图像的语义内容非常丰富，一个图像包含多个标签内容，图像标签服务准确识别自然图片中数百种场景、上千种通用物体及其属性，让智能相册管理、照片检索和分类、基于场景内容或者物体的广告推荐等功能更加直观。使用时用户发送待处理图片，返回图片标签内容及相应置信度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageMediaTagging(request *model.RunImageMediaTaggingRequest) (*model.RunImageMediaTaggingResponse, error) {
	requestDef := GenReqDefForRunImageMediaTagging()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageMediaTaggingResponse), nil
	}
}

// RunImageMediaTaggingInvoker 标签识别
func (c *ImageClient) RunImageMediaTaggingInvoker(request *model.RunImageMediaTaggingRequest) *RunImageMediaTaggingInvoker {
	requestDef := GenReqDefForRunImageMediaTagging()
	return &RunImageMediaTaggingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageMediaTaggingDet 媒资图像标签检测
//
// 自然图像的语义内容非常丰富，一个图像包含多个标签内容，图像标签服务准确识别自然图片中数百种场景、上千种通用物体及其属性，让智能相册管理、照片检索和分类、基于场景内容或者物体的广告推荐等功能更加直观。使用时用户发送待处理图片，返回图片标签内容及相应的位置坐标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageMediaTaggingDet(request *model.RunImageMediaTaggingDetRequest) (*model.RunImageMediaTaggingDetResponse, error) {
	requestDef := GenReqDefForRunImageMediaTaggingDet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageMediaTaggingDetResponse), nil
	}
}

// RunImageMediaTaggingDetInvoker 媒资图像标签检测
func (c *ImageClient) RunImageMediaTaggingDetInvoker(request *model.RunImageMediaTaggingDetRequest) *RunImageMediaTaggingDetInvoker {
	requestDef := GenReqDefForRunImageMediaTaggingDet()
	return &RunImageMediaTaggingDetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageSuperResolution 图像超分
//
// 图像数据，base64编码，输入图像范围200px ~ 1080px，支持JPG/PNG/BMP/JPEG/WEBP格式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageSuperResolution(request *model.RunImageSuperResolutionRequest) (*model.RunImageSuperResolutionResponse, error) {
	requestDef := GenReqDefForRunImageSuperResolution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageSuperResolutionResponse), nil
	}
}

// RunImageSuperResolutionInvoker 图像超分
func (c *ImageClient) RunImageSuperResolutionInvoker(request *model.RunImageSuperResolutionRequest) *RunImageSuperResolutionInvoker {
	requestDef := GenReqDefForRunImageSuperResolution()
	return &RunImageSuperResolutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunImageTagging 图像标签
//
// 自然图像的语义内容非常丰富，一个图像包含多个标签内容，图像标签服务准确识别自然图片中数百种场景、上千种通用物体及其属性，让智能相册管理、照片检索和分类、基于场景内容或者物体的广告推荐等功能更加直观。使用时用户发送待处理图片，返回图片标签内容及相应置信度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunImageTagging(request *model.RunImageTaggingRequest) (*model.RunImageTaggingResponse, error) {
	requestDef := GenReqDefForRunImageTagging()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunImageTaggingResponse), nil
	}
}

// RunImageTaggingInvoker 图像标签
func (c *ImageClient) RunImageTaggingInvoker(request *model.RunImageTaggingRequest) *RunImageTaggingInvoker {
	requestDef := GenReqDefForRunImageTagging()
	return &RunImageTaggingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQueryCustomTags 查询媒资图像标签
//
// 用于用户自查是否存在自定义的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunQueryCustomTags(request *model.RunQueryCustomTagsRequest) (*model.RunQueryCustomTagsResponse, error) {
	requestDef := GenReqDefForRunQueryCustomTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryCustomTagsResponse), nil
	}
}

// RunQueryCustomTagsInvoker 查询媒资图像标签
func (c *ImageClient) RunQueryCustomTagsInvoker(request *model.RunQueryCustomTagsRequest) *RunQueryCustomTagsInvoker {
	requestDef := GenReqDefForRunQueryCustomTags()
	return &RunQueryCustomTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunRecaptureDetect 翻拍识别
//
// 零售行业通常根据零售店的销售量进行销售奖励，拍摄售出商品的条形码上传后台是常用的统计方式。翻拍识别利用深度神经网络算法判断条形码图片为原始拍摄，还是经过二次翻拍、打印翻拍等手法二次处理的图片。利用翻拍识别，可以检测出经过二次处理的不合规范图片，使得统计数据更准确、有效。。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) RunRecaptureDetect(request *model.RunRecaptureDetectRequest) (*model.RunRecaptureDetectResponse, error) {
	requestDef := GenReqDefForRunRecaptureDetect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunRecaptureDetectResponse), nil
	}
}

// RunRecaptureDetectInvoker 翻拍识别
func (c *ImageClient) RunRecaptureDetectInvoker(request *model.RunRecaptureDetectRequest) *RunRecaptureDetectInvoker {
	requestDef := GenReqDefForRunRecaptureDetect()
	return &RunRecaptureDetectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageHighresolutionMattingTask 查询图像高清抠图任务
//
// 查询图像高清抠图任务，返回参数配置以及任务状态信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowImageHighresolutionMattingTask(request *model.ShowImageHighresolutionMattingTaskRequest) (*model.ShowImageHighresolutionMattingTaskResponse, error) {
	requestDef := GenReqDefForShowImageHighresolutionMattingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageHighresolutionMattingTaskResponse), nil
	}
}

// ShowImageHighresolutionMattingTaskInvoker 查询图像高清抠图任务
func (c *ImageClient) ShowImageHighresolutionMattingTaskInvoker(request *model.ShowImageHighresolutionMattingTaskRequest) *ShowImageHighresolutionMattingTaskInvoker {
	requestDef := GenReqDefForShowImageHighresolutionMattingTask()
	return &ShowImageHighresolutionMattingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageTranslateTask 查询任务
//
// show task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowImageTranslateTask(request *model.ShowImageTranslateTaskRequest) (*model.ShowImageTranslateTaskResponse, error) {
	requestDef := GenReqDefForShowImageTranslateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageTranslateTaskResponse), nil
	}
}

// ShowImageTranslateTaskInvoker 查询任务
func (c *ImageClient) ShowImageTranslateTaskInvoker(request *model.ShowImageTranslateTaskRequest) *ShowImageTranslateTaskInvoker {
	requestDef := GenReqDefForShowImageTranslateTask()
	return &ShowImageTranslateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoCoverAnalysisTask 查询任务
//
// 查询视频封面任务完成状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowVideoCoverAnalysisTask(request *model.ShowVideoCoverAnalysisTaskRequest) (*model.ShowVideoCoverAnalysisTaskResponse, error) {
	requestDef := GenReqDefForShowVideoCoverAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoCoverAnalysisTaskResponse), nil
	}
}

// ShowVideoCoverAnalysisTaskInvoker 查询任务
func (c *ImageClient) ShowVideoCoverAnalysisTaskInvoker(request *model.ShowVideoCoverAnalysisTaskRequest) *ShowVideoCoverAnalysisTaskInvoker {
	requestDef := GenReqDefForShowVideoCoverAnalysisTask()
	return &ShowVideoCoverAnalysisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoObjectMaskingTask 查询视频脱敏任务详情
//
// show task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowVideoObjectMaskingTask(request *model.ShowVideoObjectMaskingTaskRequest) (*model.ShowVideoObjectMaskingTaskResponse, error) {
	requestDef := GenReqDefForShowVideoObjectMaskingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoObjectMaskingTaskResponse), nil
	}
}

// ShowVideoObjectMaskingTaskInvoker 查询视频脱敏任务详情
func (c *ImageClient) ShowVideoObjectMaskingTaskInvoker(request *model.ShowVideoObjectMaskingTaskRequest) *ShowVideoObjectMaskingTaskInvoker {
	requestDef := GenReqDefForShowVideoObjectMaskingTask()
	return &ShowVideoObjectMaskingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoShotSplitTask 查询任务
//
// 查询拆条任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowVideoShotSplitTask(request *model.ShowVideoShotSplitTaskRequest) (*model.ShowVideoShotSplitTaskResponse, error) {
	requestDef := GenReqDefForShowVideoShotSplitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoShotSplitTaskResponse), nil
	}
}

// ShowVideoShotSplitTaskInvoker 查询任务
func (c *ImageClient) ShowVideoShotSplitTaskInvoker(request *model.ShowVideoShotSplitTaskRequest) *ShowVideoShotSplitTaskInvoker {
	requestDef := GenReqDefForShowVideoShotSplitTask()
	return &ShowVideoShotSplitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoSummarizationAnalysisTask 查询任务
//
// show task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowVideoSummarizationAnalysisTask(request *model.ShowVideoSummarizationAnalysisTaskRequest) (*model.ShowVideoSummarizationAnalysisTaskResponse, error) {
	requestDef := GenReqDefForShowVideoSummarizationAnalysisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoSummarizationAnalysisTaskResponse), nil
	}
}

// ShowVideoSummarizationAnalysisTaskInvoker 查询任务
func (c *ImageClient) ShowVideoSummarizationAnalysisTaskInvoker(request *model.ShowVideoSummarizationAnalysisTaskRequest) *ShowVideoSummarizationAnalysisTaskInvoker {
	requestDef := GenReqDefForShowVideoSummarizationAnalysisTask()
	return &ShowVideoSummarizationAnalysisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoTranslateTask 查询任务
//
// show task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageClient) ShowVideoTranslateTask(request *model.ShowVideoTranslateTaskRequest) (*model.ShowVideoTranslateTaskResponse, error) {
	requestDef := GenReqDefForShowVideoTranslateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoTranslateTaskResponse), nil
	}
}

// ShowVideoTranslateTaskInvoker 查询任务
func (c *ImageClient) ShowVideoTranslateTaskInvoker(request *model.ShowVideoTranslateTaskRequest) *ShowVideoTranslateTaskInvoker {
	requestDef := GenReqDefForShowVideoTranslateTask()
	return &ShowVideoTranslateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
