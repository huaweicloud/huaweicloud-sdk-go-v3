package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ivs/v2/model"
)

type IvsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIvsClient(hcClient *http_client.HcHttpClient) *IvsClient {
	return &IvsClient{HcClient: hcClient}
}

func IvsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// DetectExtentionByIdCardImage 人证核身证件版（二要素）
//
// 使用姓名、身份证号码二要素进行身份审核。身份验证时，传入的数据为身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IvsClient) DetectExtentionByIdCardImage(request *model.DetectExtentionByIdCardImageRequest) (*model.DetectExtentionByIdCardImageResponse, error) {
	requestDef := GenReqDefForDetectExtentionByIdCardImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectExtentionByIdCardImageResponse), nil
	}
}

// DetectExtentionByIdCardImageInvoker 人证核身证件版（二要素）
func (c *IvsClient) DetectExtentionByIdCardImageInvoker(request *model.DetectExtentionByIdCardImageRequest) *DetectExtentionByIdCardImageInvoker {
	requestDef := GenReqDefForDetectExtentionByIdCardImage()
	return &DetectExtentionByIdCardImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectExtentionByNameAndId 人证核身证件版（二要素）
//
// 使用姓名、身份证号码二要素进行身份审核。身份验证时，传入的数据为身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IvsClient) DetectExtentionByNameAndId(request *model.DetectExtentionByNameAndIdRequest) (*model.DetectExtentionByNameAndIdResponse, error) {
	requestDef := GenReqDefForDetectExtentionByNameAndId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectExtentionByNameAndIdResponse), nil
	}
}

// DetectExtentionByNameAndIdInvoker 人证核身证件版（二要素）
func (c *IvsClient) DetectExtentionByNameAndIdInvoker(request *model.DetectExtentionByNameAndIdRequest) *DetectExtentionByNameAndIdInvoker {
	requestDef := GenReqDefForDetectExtentionByNameAndId()
	return &DetectExtentionByNameAndIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectStandardByIdCardImage 人证核身标准版（三要素）
//
// 使用姓名、身份证号码、人脸图片三要素进行身份审核。
// 身份验证时，传入的数据为人脸图片、身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IvsClient) DetectStandardByIdCardImage(request *model.DetectStandardByIdCardImageRequest) (*model.DetectStandardByIdCardImageResponse, error) {
	requestDef := GenReqDefForDetectStandardByIdCardImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectStandardByIdCardImageResponse), nil
	}
}

// DetectStandardByIdCardImageInvoker 人证核身标准版（三要素）
func (c *IvsClient) DetectStandardByIdCardImageInvoker(request *model.DetectStandardByIdCardImageRequest) *DetectStandardByIdCardImageInvoker {
	requestDef := GenReqDefForDetectStandardByIdCardImage()
	return &DetectStandardByIdCardImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectStandardByNameAndId 人证核身标准版（三要素）
//
// 校验用户上传的身份证图片支持正反面同时上传 中的信息的真实性，输出最终的审核结果。 该接口也支持用户直接上传姓名和身份证号码进行合法性校验 。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IvsClient) DetectStandardByNameAndId(request *model.DetectStandardByNameAndIdRequest) (*model.DetectStandardByNameAndIdResponse, error) {
	requestDef := GenReqDefForDetectStandardByNameAndId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectStandardByNameAndIdResponse), nil
	}
}

// DetectStandardByNameAndIdInvoker 人证核身标准版（三要素）
func (c *IvsClient) DetectStandardByNameAndIdInvoker(request *model.DetectStandardByNameAndIdRequest) *DetectStandardByNameAndIdInvoker {
	requestDef := GenReqDefForDetectStandardByNameAndId()
	return &DetectStandardByNameAndIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
