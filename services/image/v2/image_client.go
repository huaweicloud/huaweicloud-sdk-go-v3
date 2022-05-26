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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// RunImageTagging 图像标签
//
// 自然图像的语义内容非常丰富，一个图像包含多个标签内容，图像标签服务准确识别自然图片中数百种场景、上千种通用物体及其属性，让智能相册管理、照片检索和分类、基于场景内容或者物体的广告推荐等功能更加直观。使用时用户发送待处理图片，返回图片标签内容及相应置信度。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
