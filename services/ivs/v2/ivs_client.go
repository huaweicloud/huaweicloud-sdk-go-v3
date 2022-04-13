package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/ivs/v2/model"
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

//使用姓名、身份证号码二要素进行身份审核。身份验证时，传入的数据为身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
func (c *IvsClient) DetectExtentionByIdCardImage(request *model.DetectExtentionByIdCardImageRequest) (*model.DetectExtentionByIdCardImageResponse, error) {
	requestDef := GenReqDefForDetectExtentionByIdCardImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectExtentionByIdCardImageResponse), nil
	}
}

//使用姓名、身份证号码二要素进行身份审核。身份验证时，传入的数据为身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
func (c *IvsClient) DetectExtentionByNameAndId(request *model.DetectExtentionByNameAndIdRequest) (*model.DetectExtentionByNameAndIdResponse, error) {
	requestDef := GenReqDefForDetectExtentionByNameAndId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectExtentionByNameAndIdResponse), nil
	}
}

//使用姓名、身份证号码、人脸图片三要素进行身份审核。 身份验证时，传入的数据为人脸图片、身份证信息。提取身份证信息时，可以使用身份证正反面图片，也可以直接输入姓名、身份证号文本。
func (c *IvsClient) DetectStandardByIdCardImage(request *model.DetectStandardByIdCardImageRequest) (*model.DetectStandardByIdCardImageResponse, error) {
	requestDef := GenReqDefForDetectStandardByIdCardImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectStandardByIdCardImageResponse), nil
	}
}

//校验用户上传的身份证图片支持正反面同时上传 中的信息的真实性，输出最终的审核结果。 该接口也支持用户直接上传姓名和身份证号码进行合法性校验 。
func (c *IvsClient) DetectStandardByNameAndId(request *model.DetectStandardByNameAndIdRequest) (*model.DetectStandardByNameAndIdResponse, error) {
	requestDef := GenReqDefForDetectStandardByNameAndId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectStandardByNameAndIdResponse), nil
	}
}
