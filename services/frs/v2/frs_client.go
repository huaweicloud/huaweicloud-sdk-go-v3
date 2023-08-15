package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/frs/v2/model"
)

type FrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewFrsClient(hcClient *http_client.HcHttpClient) *FrsClient {
	return &FrsClient{HcClient: hcClient}
}

func FrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddFacesByBase64 添加人脸
//
// 添加人脸到人脸库中。将单张图片中的人脸添加至人脸库中，支持添加最大人脸或所有人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) AddFacesByBase64(request *model.AddFacesByBase64Request) (*model.AddFacesByBase64Response, error) {
	requestDef := GenReqDefForAddFacesByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByBase64Response), nil
	}
}

// AddFacesByBase64Invoker 添加人脸
func (c *FrsClient) AddFacesByBase64Invoker(request *model.AddFacesByBase64Request) *AddFacesByBase64Invoker {
	requestDef := GenReqDefForAddFacesByBase64()
	return &AddFacesByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddFacesByFile 添加人脸
//
// 添加人脸到人脸库中。将单张图片中的人脸添加至人脸库中，支持添加最大人脸或所有人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) AddFacesByFile(request *model.AddFacesByFileRequest) (*model.AddFacesByFileResponse, error) {
	requestDef := GenReqDefForAddFacesByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByFileResponse), nil
	}
}

// AddFacesByFileInvoker 添加人脸
func (c *FrsClient) AddFacesByFileInvoker(request *model.AddFacesByFileRequest) *AddFacesByFileInvoker {
	requestDef := GenReqDefForAddFacesByFile()
	return &AddFacesByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddFacesByUrl 添加人脸
//
// 添加人脸到人脸库中。将单张图片中的人脸添加至人脸库中，支持添加最大人脸或所有人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) AddFacesByUrl(request *model.AddFacesByUrlRequest) (*model.AddFacesByUrlResponse, error) {
	requestDef := GenReqDefForAddFacesByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByUrlResponse), nil
	}
}

// AddFacesByUrlInvoker 添加人脸
func (c *FrsClient) AddFacesByUrlInvoker(request *model.AddFacesByUrlRequest) *AddFacesByUrlInvoker {
	requestDef := GenReqDefForAddFacesByUrl()
	return &AddFacesByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteFaces 批量删除人脸
//
// 自定义筛选条件，批量删除人脸库中的符合指定条件的多张人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) BatchDeleteFaces(request *model.BatchDeleteFacesRequest) (*model.BatchDeleteFacesResponse, error) {
	requestDef := GenReqDefForBatchDeleteFaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFacesResponse), nil
	}
}

// BatchDeleteFacesInvoker 批量删除人脸
func (c *FrsClient) BatchDeleteFacesInvoker(request *model.BatchDeleteFacesRequest) *BatchDeleteFacesInvoker {
	requestDef := GenReqDefForBatchDeleteFaces()
	return &BatchDeleteFacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareFaceByBase64 人脸比对
//
// 人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) CompareFaceByBase64(request *model.CompareFaceByBase64Request) (*model.CompareFaceByBase64Response, error) {
	requestDef := GenReqDefForCompareFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByBase64Response), nil
	}
}

// CompareFaceByBase64Invoker 人脸比对
func (c *FrsClient) CompareFaceByBase64Invoker(request *model.CompareFaceByBase64Request) *CompareFaceByBase64Invoker {
	requestDef := GenReqDefForCompareFaceByBase64()
	return &CompareFaceByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareFaceByFile 人脸比对
//
// 人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) CompareFaceByFile(request *model.CompareFaceByFileRequest) (*model.CompareFaceByFileResponse, error) {
	requestDef := GenReqDefForCompareFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByFileResponse), nil
	}
}

// CompareFaceByFileInvoker 人脸比对
func (c *FrsClient) CompareFaceByFileInvoker(request *model.CompareFaceByFileRequest) *CompareFaceByFileInvoker {
	requestDef := GenReqDefForCompareFaceByFile()
	return &CompareFaceByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareFaceByUrl 人脸比对
//
// 人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) CompareFaceByUrl(request *model.CompareFaceByUrlRequest) (*model.CompareFaceByUrlResponse, error) {
	requestDef := GenReqDefForCompareFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByUrlResponse), nil
	}
}

// CompareFaceByUrlInvoker 人脸比对
func (c *FrsClient) CompareFaceByUrlInvoker(request *model.CompareFaceByUrlRequest) *CompareFaceByUrlInvoker {
	requestDef := GenReqDefForCompareFaceByUrl()
	return &CompareFaceByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFaceSet 创建人脸库
//
// 创建用于存储人脸特征的人脸库。您最多可以创建10个人脸库，每个人脸库最大容量为10万个人脸特征。如有更大规格的需求请联系客服。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) CreateFaceSet(request *model.CreateFaceSetRequest) (*model.CreateFaceSetResponse, error) {
	requestDef := GenReqDefForCreateFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFaceSetResponse), nil
	}
}

// CreateFaceSetInvoker 创建人脸库
func (c *FrsClient) CreateFaceSetInvoker(request *model.CreateFaceSetRequest) *CreateFaceSetInvoker {
	requestDef := GenReqDefForCreateFaceSet()
	return &CreateFaceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFaceByExternalImageId 删除人脸
//
// 根据external_image_id删除人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DeleteFaceByExternalImageId(request *model.DeleteFaceByExternalImageIdRequest) (*model.DeleteFaceByExternalImageIdResponse, error) {
	requestDef := GenReqDefForDeleteFaceByExternalImageId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceByExternalImageIdResponse), nil
	}
}

// DeleteFaceByExternalImageIdInvoker 删除人脸
func (c *FrsClient) DeleteFaceByExternalImageIdInvoker(request *model.DeleteFaceByExternalImageIdRequest) *DeleteFaceByExternalImageIdInvoker {
	requestDef := GenReqDefForDeleteFaceByExternalImageId()
	return &DeleteFaceByExternalImageIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFaceByFaceId 删除人脸
//
// 根据face_id删除人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DeleteFaceByFaceId(request *model.DeleteFaceByFaceIdRequest) (*model.DeleteFaceByFaceIdResponse, error) {
	requestDef := GenReqDefForDeleteFaceByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceByFaceIdResponse), nil
	}
}

// DeleteFaceByFaceIdInvoker 删除人脸
func (c *FrsClient) DeleteFaceByFaceIdInvoker(request *model.DeleteFaceByFaceIdRequest) *DeleteFaceByFaceIdInvoker {
	requestDef := GenReqDefForDeleteFaceByFaceId()
	return &DeleteFaceByFaceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFaceSet 删除人脸库
//
// 删除人脸库以及其中所有的人脸。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DeleteFaceSet(request *model.DeleteFaceSetRequest) (*model.DeleteFaceSetResponse, error) {
	requestDef := GenReqDefForDeleteFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceSetResponse), nil
	}
}

// DeleteFaceSetInvoker 删除人脸库
func (c *FrsClient) DeleteFaceSetInvoker(request *model.DeleteFaceSetRequest) *DeleteFaceSetInvoker {
	requestDef := GenReqDefForDeleteFaceSet()
	return &DeleteFaceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByBase64 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByBase64(request *model.DetectFaceByBase64Request) (*model.DetectFaceByBase64Response, error) {
	requestDef := GenReqDefForDetectFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByBase64Response), nil
	}
}

// DetectFaceByBase64Invoker 人脸检测
func (c *FrsClient) DetectFaceByBase64Invoker(request *model.DetectFaceByBase64Request) *DetectFaceByBase64Invoker {
	requestDef := GenReqDefForDetectFaceByBase64()
	return &DetectFaceByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByBase64Intl 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByBase64Intl(request *model.DetectFaceByBase64IntlRequest) (*model.DetectFaceByBase64IntlResponse, error) {
	requestDef := GenReqDefForDetectFaceByBase64Intl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByBase64IntlResponse), nil
	}
}

// DetectFaceByBase64IntlInvoker 人脸检测
func (c *FrsClient) DetectFaceByBase64IntlInvoker(request *model.DetectFaceByBase64IntlRequest) *DetectFaceByBase64IntlInvoker {
	requestDef := GenReqDefForDetectFaceByBase64Intl()
	return &DetectFaceByBase64IntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByFile 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByFile(request *model.DetectFaceByFileRequest) (*model.DetectFaceByFileResponse, error) {
	requestDef := GenReqDefForDetectFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByFileResponse), nil
	}
}

// DetectFaceByFileInvoker 人脸检测
func (c *FrsClient) DetectFaceByFileInvoker(request *model.DetectFaceByFileRequest) *DetectFaceByFileInvoker {
	requestDef := GenReqDefForDetectFaceByFile()
	return &DetectFaceByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByFileIntl 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByFileIntl(request *model.DetectFaceByFileIntlRequest) (*model.DetectFaceByFileIntlResponse, error) {
	requestDef := GenReqDefForDetectFaceByFileIntl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByFileIntlResponse), nil
	}
}

// DetectFaceByFileIntlInvoker 人脸检测
func (c *FrsClient) DetectFaceByFileIntlInvoker(request *model.DetectFaceByFileIntlRequest) *DetectFaceByFileIntlInvoker {
	requestDef := GenReqDefForDetectFaceByFileIntl()
	return &DetectFaceByFileIntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByUrl 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByUrl(request *model.DetectFaceByUrlRequest) (*model.DetectFaceByUrlResponse, error) {
	requestDef := GenReqDefForDetectFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByUrlResponse), nil
	}
}

// DetectFaceByUrlInvoker 人脸检测
func (c *FrsClient) DetectFaceByUrlInvoker(request *model.DetectFaceByUrlRequest) *DetectFaceByUrlInvoker {
	requestDef := GenReqDefForDetectFaceByUrl()
	return &DetectFaceByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectFaceByUrlIntl 人脸检测
//
// 人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectFaceByUrlIntl(request *model.DetectFaceByUrlIntlRequest) (*model.DetectFaceByUrlIntlResponse, error) {
	requestDef := GenReqDefForDetectFaceByUrlIntl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByUrlIntlResponse), nil
	}
}

// DetectFaceByUrlIntlInvoker 人脸检测
func (c *FrsClient) DetectFaceByUrlIntlInvoker(request *model.DetectFaceByUrlIntlRequest) *DetectFaceByUrlIntlInvoker {
	requestDef := GenReqDefForDetectFaceByUrlIntl()
	return &DetectFaceByUrlIntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByBase64 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByBase64(request *model.DetectLiveByBase64Request) (*model.DetectLiveByBase64Response, error) {
	requestDef := GenReqDefForDetectLiveByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByBase64Response), nil
	}
}

// DetectLiveByBase64Invoker 动作活体检测
func (c *FrsClient) DetectLiveByBase64Invoker(request *model.DetectLiveByBase64Request) *DetectLiveByBase64Invoker {
	requestDef := GenReqDefForDetectLiveByBase64()
	return &DetectLiveByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByBase64Intl 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByBase64Intl(request *model.DetectLiveByBase64IntlRequest) (*model.DetectLiveByBase64IntlResponse, error) {
	requestDef := GenReqDefForDetectLiveByBase64Intl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByBase64IntlResponse), nil
	}
}

// DetectLiveByBase64IntlInvoker 动作活体检测
func (c *FrsClient) DetectLiveByBase64IntlInvoker(request *model.DetectLiveByBase64IntlRequest) *DetectLiveByBase64IntlInvoker {
	requestDef := GenReqDefForDetectLiveByBase64Intl()
	return &DetectLiveByBase64IntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByFile 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByFile(request *model.DetectLiveByFileRequest) (*model.DetectLiveByFileResponse, error) {
	requestDef := GenReqDefForDetectLiveByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByFileResponse), nil
	}
}

// DetectLiveByFileInvoker 动作活体检测
func (c *FrsClient) DetectLiveByFileInvoker(request *model.DetectLiveByFileRequest) *DetectLiveByFileInvoker {
	requestDef := GenReqDefForDetectLiveByFile()
	return &DetectLiveByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByFileIntl 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByFileIntl(request *model.DetectLiveByFileIntlRequest) (*model.DetectLiveByFileIntlResponse, error) {
	requestDef := GenReqDefForDetectLiveByFileIntl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByFileIntlResponse), nil
	}
}

// DetectLiveByFileIntlInvoker 动作活体检测
func (c *FrsClient) DetectLiveByFileIntlInvoker(request *model.DetectLiveByFileIntlRequest) *DetectLiveByFileIntlInvoker {
	requestDef := GenReqDefForDetectLiveByFileIntl()
	return &DetectLiveByFileIntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByUrl 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByUrl(request *model.DetectLiveByUrlRequest) (*model.DetectLiveByUrlResponse, error) {
	requestDef := GenReqDefForDetectLiveByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByUrlResponse), nil
	}
}

// DetectLiveByUrlInvoker 动作活体检测
func (c *FrsClient) DetectLiveByUrlInvoker(request *model.DetectLiveByUrlRequest) *DetectLiveByUrlInvoker {
	requestDef := GenReqDefForDetectLiveByUrl()
	return &DetectLiveByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveByUrlIntl 动作活体检测
//
// 动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveByUrlIntl(request *model.DetectLiveByUrlIntlRequest) (*model.DetectLiveByUrlIntlResponse, error) {
	requestDef := GenReqDefForDetectLiveByUrlIntl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByUrlIntlResponse), nil
	}
}

// DetectLiveByUrlIntlInvoker 动作活体检测
func (c *FrsClient) DetectLiveByUrlIntlInvoker(request *model.DetectLiveByUrlIntlRequest) *DetectLiveByUrlIntlInvoker {
	requestDef := GenReqDefForDetectLiveByUrlIntl()
	return &DetectLiveByUrlIntlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveFaceByBase64 静默活体检测
//
// 静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveFaceByBase64(request *model.DetectLiveFaceByBase64Request) (*model.DetectLiveFaceByBase64Response, error) {
	requestDef := GenReqDefForDetectLiveFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByBase64Response), nil
	}
}

// DetectLiveFaceByBase64Invoker 静默活体检测
func (c *FrsClient) DetectLiveFaceByBase64Invoker(request *model.DetectLiveFaceByBase64Request) *DetectLiveFaceByBase64Invoker {
	requestDef := GenReqDefForDetectLiveFaceByBase64()
	return &DetectLiveFaceByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveFaceByFile 静默活体检测
//
// 静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveFaceByFile(request *model.DetectLiveFaceByFileRequest) (*model.DetectLiveFaceByFileResponse, error) {
	requestDef := GenReqDefForDetectLiveFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByFileResponse), nil
	}
}

// DetectLiveFaceByFileInvoker 静默活体检测
func (c *FrsClient) DetectLiveFaceByFileInvoker(request *model.DetectLiveFaceByFileRequest) *DetectLiveFaceByFileInvoker {
	requestDef := GenReqDefForDetectLiveFaceByFile()
	return &DetectLiveFaceByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetectLiveFaceByUrl 静默活体检测
//
// 静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) DetectLiveFaceByUrl(request *model.DetectLiveFaceByUrlRequest) (*model.DetectLiveFaceByUrlResponse, error) {
	requestDef := GenReqDefForDetectLiveFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByUrlResponse), nil
	}
}

// DetectLiveFaceByUrlInvoker 静默活体检测
func (c *FrsClient) DetectLiveFaceByUrlInvoker(request *model.DetectLiveFaceByUrlRequest) *DetectLiveFaceByUrlInvoker {
	requestDef := GenReqDefForDetectLiveFaceByUrl()
	return &DetectLiveFaceByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFaceByBase64 人脸搜索
//
// 人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。
// 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) SearchFaceByBase64(request *model.SearchFaceByBase64Request) (*model.SearchFaceByBase64Response, error) {
	requestDef := GenReqDefForSearchFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByBase64Response), nil
	}
}

// SearchFaceByBase64Invoker 人脸搜索
func (c *FrsClient) SearchFaceByBase64Invoker(request *model.SearchFaceByBase64Request) *SearchFaceByBase64Invoker {
	requestDef := GenReqDefForSearchFaceByBase64()
	return &SearchFaceByBase64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFaceByFaceId 人脸搜索
//
// 人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。
// 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) SearchFaceByFaceId(request *model.SearchFaceByFaceIdRequest) (*model.SearchFaceByFaceIdResponse, error) {
	requestDef := GenReqDefForSearchFaceByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByFaceIdResponse), nil
	}
}

// SearchFaceByFaceIdInvoker 人脸搜索
func (c *FrsClient) SearchFaceByFaceIdInvoker(request *model.SearchFaceByFaceIdRequest) *SearchFaceByFaceIdInvoker {
	requestDef := GenReqDefForSearchFaceByFaceId()
	return &SearchFaceByFaceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFaceByFile 人脸搜索
//
// 人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。
// 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) SearchFaceByFile(request *model.SearchFaceByFileRequest) (*model.SearchFaceByFileResponse, error) {
	requestDef := GenReqDefForSearchFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByFileResponse), nil
	}
}

// SearchFaceByFileInvoker 人脸搜索
func (c *FrsClient) SearchFaceByFileInvoker(request *model.SearchFaceByFileRequest) *SearchFaceByFileInvoker {
	requestDef := GenReqDefForSearchFaceByFile()
	return &SearchFaceByFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFaceByUrl 人脸搜索
//
// 人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。
// 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) SearchFaceByUrl(request *model.SearchFaceByUrlRequest) (*model.SearchFaceByUrlResponse, error) {
	requestDef := GenReqDefForSearchFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByUrlResponse), nil
	}
}

// SearchFaceByUrlInvoker 人脸搜索
func (c *FrsClient) SearchFaceByUrlInvoker(request *model.SearchFaceByUrlRequest) *SearchFaceByUrlInvoker {
	requestDef := GenReqDefForSearchFaceByUrl()
	return &SearchFaceByUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllFaceSets 查询所有人脸库
//
// 查询当前用户所有人脸库的状态信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) ShowAllFaceSets(request *model.ShowAllFaceSetsRequest) (*model.ShowAllFaceSetsResponse, error) {
	requestDef := GenReqDefForShowAllFaceSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllFaceSetsResponse), nil
	}
}

// ShowAllFaceSetsInvoker 查询所有人脸库
func (c *FrsClient) ShowAllFaceSetsInvoker(request *model.ShowAllFaceSetsRequest) *ShowAllFaceSetsInvoker {
	requestDef := GenReqDefForShowAllFaceSets()
	return &ShowAllFaceSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFaceSet 查询人脸库
//
// 查询人脸库当前的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) ShowFaceSet(request *model.ShowFaceSetRequest) (*model.ShowFaceSetResponse, error) {
	requestDef := GenReqDefForShowFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFaceSetResponse), nil
	}
}

// ShowFaceSetInvoker 查询人脸库
func (c *FrsClient) ShowFaceSetInvoker(request *model.ShowFaceSetRequest) *ShowFaceSetInvoker {
	requestDef := GenReqDefForShowFaceSet()
	return &ShowFaceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFacesByFaceId 查询人脸
//
// 查询指定人脸库中人脸信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) ShowFacesByFaceId(request *model.ShowFacesByFaceIdRequest) (*model.ShowFacesByFaceIdResponse, error) {
	requestDef := GenReqDefForShowFacesByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFacesByFaceIdResponse), nil
	}
}

// ShowFacesByFaceIdInvoker 查询人脸
func (c *FrsClient) ShowFacesByFaceIdInvoker(request *model.ShowFacesByFaceIdRequest) *ShowFacesByFaceIdInvoker {
	requestDef := GenReqDefForShowFacesByFaceId()
	return &ShowFacesByFaceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFacesByLimit 查询人脸
//
// 查询指定人脸库中人脸信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) ShowFacesByLimit(request *model.ShowFacesByLimitRequest) (*model.ShowFacesByLimitResponse, error) {
	requestDef := GenReqDefForShowFacesByLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFacesByLimitResponse), nil
	}
}

// ShowFacesByLimitInvoker 查询人脸
func (c *FrsClient) ShowFacesByLimitInvoker(request *model.ShowFacesByLimitRequest) *ShowFacesByLimitInvoker {
	requestDef := GenReqDefForShowFacesByLimit()
	return &ShowFacesByLimitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFace 更新人脸
//
// 根据人脸ID（face_id）更新单张人脸信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FrsClient) UpdateFace(request *model.UpdateFaceRequest) (*model.UpdateFaceResponse, error) {
	requestDef := GenReqDefForUpdateFace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFaceResponse), nil
	}
}

// UpdateFaceInvoker 更新人脸
func (c *FrsClient) UpdateFaceInvoker(request *model.UpdateFaceRequest) *UpdateFaceInvoker {
	requestDef := GenReqDefForUpdateFace()
	return &UpdateFaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
