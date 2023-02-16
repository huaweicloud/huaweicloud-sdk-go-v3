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
