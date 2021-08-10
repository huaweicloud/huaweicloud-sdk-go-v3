package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/frs/v1/model"
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

//添加人脸到人脸库中，检测到传入的单张图片中存在多少张人脸，则增加多少张人脸到人脸库当中。
func (c *FrsClient) AddFacesByBase64(request *model.AddFacesByBase64Request) (*model.AddFacesByBase64Response, error) {
	requestDef := GenReqDefForAddFacesByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByBase64Response), nil
	}
}

//添加人脸到人脸库中，检测到传入的单张图片中存在多少张人脸，则增加多少张人脸到人脸库当中。
func (c *FrsClient) AddFacesByFile(request *model.AddFacesByFileRequest) (*model.AddFacesByFileResponse, error) {
	requestDef := GenReqDefForAddFacesByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByFileResponse), nil
	}
}

//添加人脸到人脸库中，检测到传入的单张图片中存在多少张人脸，则增加多少张人脸到人脸库当中。
func (c *FrsClient) AddFacesByUrl(request *model.AddFacesByUrlRequest) (*model.AddFacesByUrlResponse, error) {
	requestDef := GenReqDefForAddFacesByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFacesByUrlResponse), nil
	}
}

//用于查询服务的开通状态。
func (c *FrsClient) AuthorizeFaceRecognitionService(request *model.AuthorizeFaceRecognitionServiceRequest) (*model.AuthorizeFaceRecognitionServiceResponse, error) {
	requestDef := GenReqDefForAuthorizeFaceRecognitionService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeFaceRecognitionServiceResponse), nil
	}
}

//自定义筛选条件，批量删除人脸库中的符合指定条件的多张人脸。
func (c *FrsClient) BatchDeleteFaces(request *model.BatchDeleteFacesRequest) (*model.BatchDeleteFacesResponse, error) {
	requestDef := GenReqDefForBatchDeleteFaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFacesResponse), nil
	}
}

//人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
func (c *FrsClient) CompareFaceByBase64(request *model.CompareFaceByBase64Request) (*model.CompareFaceByBase64Response, error) {
	requestDef := GenReqDefForCompareFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByBase64Response), nil
	}
}

//人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
func (c *FrsClient) CompareFaceByFile(request *model.CompareFaceByFileRequest) (*model.CompareFaceByFileResponse, error) {
	requestDef := GenReqDefForCompareFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByFileResponse), nil
	}
}

//人脸比对是将两个人脸进行比对，来判断是否为同一个人，返回比对置信度。如果传入的图片中包含多个人脸，选取最大的人脸进行比对。
func (c *FrsClient) CompareFaceByUrl(request *model.CompareFaceByUrlRequest) (*model.CompareFaceByUrlResponse, error) {
	requestDef := GenReqDefForCompareFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareFaceByUrlResponse), nil
	}
}

//创建用于存储人脸特征的人脸库。您最多可以创建10个人脸库，每个人脸库最大容量为10万个人脸特征。如有更大规格的需求请联系客服。
func (c *FrsClient) CreateFaceSet(request *model.CreateFaceSetRequest) (*model.CreateFaceSetResponse, error) {
	requestDef := GenReqDefForCreateFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFaceSetResponse), nil
	}
}

//根据external_image_id删除人脸。
func (c *FrsClient) DeleteFaceByExternalImageId(request *model.DeleteFaceByExternalImageIdRequest) (*model.DeleteFaceByExternalImageIdResponse, error) {
	requestDef := GenReqDefForDeleteFaceByExternalImageId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceByExternalImageIdResponse), nil
	}
}

//根据face_id删除人脸。
func (c *FrsClient) DeleteFaceByFaceId(request *model.DeleteFaceByFaceIdRequest) (*model.DeleteFaceByFaceIdResponse, error) {
	requestDef := GenReqDefForDeleteFaceByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceByFaceIdResponse), nil
	}
}

//删除人脸库以及其中所有的人脸。
func (c *FrsClient) DeleteFaceSet(request *model.DeleteFaceSetRequest) (*model.DeleteFaceSetResponse, error) {
	requestDef := GenReqDefForDeleteFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFaceSetResponse), nil
	}
}

//人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
func (c *FrsClient) DetectFaceByBase64(request *model.DetectFaceByBase64Request) (*model.DetectFaceByBase64Response, error) {
	requestDef := GenReqDefForDetectFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByBase64Response), nil
	}
}

//人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
func (c *FrsClient) DetectFaceByFile(request *model.DetectFaceByFileRequest) (*model.DetectFaceByFileResponse, error) {
	requestDef := GenReqDefForDetectFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByFileResponse), nil
	}
}

//人脸检测是对输入图片进行人脸检测和分析，输出人脸在图像中的位置、人脸关键点位置和人脸关键属性。
func (c *FrsClient) DetectFaceByUrl(request *model.DetectFaceByUrlRequest) (*model.DetectFaceByUrlResponse, error) {
	requestDef := GenReqDefForDetectFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectFaceByUrlResponse), nil
	}
}

//动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
func (c *FrsClient) DetectLiveByBase64(request *model.DetectLiveByBase64Request) (*model.DetectLiveByBase64Response, error) {
	requestDef := GenReqDefForDetectLiveByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByBase64Response), nil
	}
}

//动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
func (c *FrsClient) DetectLiveByFile(request *model.DetectLiveByFileRequest) (*model.DetectLiveByFileResponse, error) {
	requestDef := GenReqDefForDetectLiveByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByFileResponse), nil
	}
}

//动作活体检测是通过判断视频中的人物动作与传入动作列表是否一致来识别视频中人物是否为活体。如果有多张人脸出现，则选取最大的人脸进行判定。
func (c *FrsClient) DetectLiveByUrl(request *model.DetectLiveByUrlRequest) (*model.DetectLiveByUrlResponse, error) {
	requestDef := GenReqDefForDetectLiveByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveByUrlResponse), nil
	}
}

//静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
func (c *FrsClient) DetectLiveFaceByBase64(request *model.DetectLiveFaceByBase64Request) (*model.DetectLiveFaceByBase64Response, error) {
	requestDef := GenReqDefForDetectLiveFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByBase64Response), nil
	}
}

//静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
func (c *FrsClient) DetectLiveFaceByFile(request *model.DetectLiveFaceByFileRequest) (*model.DetectLiveFaceByFileResponse, error) {
	requestDef := GenReqDefForDetectLiveFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByFileResponse), nil
	}
}

//静默活体检测是基于人脸图片中可能存在的畸变、摩尔纹、反光、倒影、边框等信息，判断图片中的人脸是否来自于真人活体，有效抵御纸质翻拍照、电子翻拍照以及视频翻拍等各种攻击方式。静默活体检测支持单张图片，不支持多人脸图片。
func (c *FrsClient) DetectLiveFaceByUrl(request *model.DetectLiveFaceByUrlRequest) (*model.DetectLiveFaceByUrlResponse, error) {
	requestDef := GenReqDefForDetectLiveFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetectLiveFaceByUrlResponse), nil
	}
}

//人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
func (c *FrsClient) SearchFaceByBase64(request *model.SearchFaceByBase64Request) (*model.SearchFaceByBase64Response, error) {
	requestDef := GenReqDefForSearchFaceByBase64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByBase64Response), nil
	}
}

//人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
func (c *FrsClient) SearchFaceByFaceId(request *model.SearchFaceByFaceIdRequest) (*model.SearchFaceByFaceIdResponse, error) {
	requestDef := GenReqDefForSearchFaceByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByFaceIdResponse), nil
	}
}

//人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
func (c *FrsClient) SearchFaceByFile(request *model.SearchFaceByFileRequest) (*model.SearchFaceByFileResponse, error) {
	requestDef := GenReqDefForSearchFaceByFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByFileResponse), nil
	}
}

//人脸搜索是指在已有的人脸库中，查询与目标人脸相似的一张或者多张人脸，并返回相应的置信度。 支持传入图片或者faceID进行人脸搜索，如果传入的是多张人脸图片，选取图片中检测到的最大尺寸人脸作为检索的输入。
func (c *FrsClient) SearchFaceByUrl(request *model.SearchFaceByUrlRequest) (*model.SearchFaceByUrlResponse, error) {
	requestDef := GenReqDefForSearchFaceByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFaceByUrlResponse), nil
	}
}

//查询当前用户所有人脸库的状态信息。
func (c *FrsClient) ShowAllFaceSets(request *model.ShowAllFaceSetsRequest) (*model.ShowAllFaceSetsResponse, error) {
	requestDef := GenReqDefForShowAllFaceSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllFaceSetsResponse), nil
	}
}

//查询人脸库当前的状态。
func (c *FrsClient) ShowFaceSet(request *model.ShowFaceSetRequest) (*model.ShowFaceSetResponse, error) {
	requestDef := GenReqDefForShowFaceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFaceSetResponse), nil
	}
}

//查询指定人脸库中人脸信息。
func (c *FrsClient) ShowFacesByFaceId(request *model.ShowFacesByFaceIdRequest) (*model.ShowFacesByFaceIdResponse, error) {
	requestDef := GenReqDefForShowFacesByFaceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFacesByFaceIdResponse), nil
	}
}

//查询指定人脸库中人脸信息。
func (c *FrsClient) ShowFacesByLimit(request *model.ShowFacesByLimitRequest) (*model.ShowFacesByLimitResponse, error) {
	requestDef := GenReqDefForShowFacesByLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFacesByLimitResponse), nil
	}
}

//根据人脸ID（face_id）更新单张人脸信息。
func (c *FrsClient) UpdateFace(request *model.UpdateFaceRequest) (*model.UpdateFaceResponse, error) {
	requestDef := GenReqDefForUpdateFace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFaceResponse), nil
	}
}
