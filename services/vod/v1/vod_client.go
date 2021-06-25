package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vod/v1/model"
)

type VodClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVodClient(hcClient *http_client.HcHttpClient) *VodClient {
	return &VodClient{HcClient: hcClient}
}

func VodClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//## 典型场景 ##   取消媒资转码任务。<br/>  ## 接口功能 ##   取消媒资转码任务，只能取消排队中的转码任务。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CancelAssetTranscodeTask(request *model.CancelAssetTranscodeTaskRequest) (*model.CancelAssetTranscodeTaskResponse, error) {
	requestDef := GenReqDefForCancelAssetTranscodeTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAssetTranscodeTaskResponse), nil
	}
}

//## 典型场景 ##   取消提取音频任务调用此接口<br/>  ## 接口功能 ##   取消提取音频任务，只有排队中的音频任务才可以取消<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CancelExtractAudioTask(request *model.CancelExtractAudioTaskRequest) (*model.CancelExtractAudioTaskResponse, error) {
	requestDef := GenReqDefForCancelExtractAudioTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelExtractAudioTaskResponse), nil
	}
}

//查询音视频MD5是否重复
func (c *VodClient) CheckMd5Duplication(request *model.CheckMd5DuplicationRequest) (*model.CheckMd5DuplicationResponse, error) {
	requestDef := GenReqDefForCheckMd5Duplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckMd5DuplicationResponse), nil
	}
}

//## 典型场景 ##   确认媒资上传时调用此接口。<br/>  ## 接口功能 ##   确认媒资上传。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) ConfirmAssetUpload(request *model.ConfirmAssetUploadRequest) (*model.ConfirmAssetUploadResponse, error) {
	requestDef := GenReqDefForConfirmAssetUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmAssetUploadResponse), nil
	}
}

//## 典型场景 ##   确认水印图片上传调用此接口<br/>  ## 接口功能 ##   确认水印图片是否已经上传至对象存储<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) ConfirmImageUpload(request *model.ConfirmImageUploadRequest) (*model.ConfirmImageUploadResponse, error) {
	requestDef := GenReqDefForConfirmImageUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmImageUploadResponse), nil
	}
}

//上传方式创建媒资。
func (c *VodClient) CreateAssetByFileUpload(request *model.CreateAssetByFileUploadRequest) (*model.CreateAssetByFileUploadResponse, error) {
	requestDef := GenReqDefForCreateAssetByFileUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetByFileUploadResponse), nil
	}
}

//## 典型场景 ##   创建媒资分类。<br/>  ## 接口功能 ##   创建媒资分类。<br/>  ## 接口约束 ##   最大支持三级分类，每个分类最多支持创建128个子分类。<br/>
func (c *VodClient) CreateAssetCategory(request *model.CreateAssetCategoryRequest) (*model.CreateAssetCategoryResponse, error) {
	requestDef := GenReqDefForCreateAssetCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetCategoryResponse), nil
	}
}

//## 典型场景 ##   媒资处理。<br/>  ## 接口功能 ##   媒资处理。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CreateAssetProcessTask(request *model.CreateAssetProcessTaskRequest) (*model.CreateAssetProcessTaskResponse, error) {
	requestDef := GenReqDefForCreateAssetProcessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetProcessTaskResponse), nil
	}
}

//媒资审核接口
func (c *VodClient) CreateAssetReviewTask(request *model.CreateAssetReviewTaskRequest) (*model.CreateAssetReviewTaskResponse, error) {
	requestDef := GenReqDefForCreateAssetReviewTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssetReviewTaskResponse), nil
	}
}

//## 典型场景 ##   从媒资中提取音频调用此接口<br/>  ## 接口功能 ##   视频媒资提取音频<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CreateExtractAudioTask(request *model.CreateExtractAudioTaskRequest) (*model.CreateExtractAudioTaskResponse, error) {
	requestDef := GenReqDefForCreateExtractAudioTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExtractAudioTaskResponse), nil
	}
}

//## 典型场景 ##   创建自定义模板组时调用此接口。<br/>  ## 接口功能 ##   创建模板组。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CreateTemplateGroup(request *model.CreateTemplateGroupRequest) (*model.CreateTemplateGroupResponse, error) {
	requestDef := GenReqDefForCreateTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateGroupResponse), nil
	}
}

//## 典型场景 ##   创建水印模板调用此接口<br/>  ## 接口功能 ##   创建水印模板<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CreateWatermarkTemplate(request *model.CreateWatermarkTemplateRequest) (*model.CreateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForCreateWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWatermarkTemplateResponse), nil
	}
}

//## 典型场景 ##   删除媒资分类。<br/>  ## 接口功能 ##   删除媒资分类。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) DeleteAssetCategory(request *model.DeleteAssetCategoryRequest) (*model.DeleteAssetCategoryResponse, error) {
	requestDef := GenReqDefForDeleteAssetCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetCategoryResponse), nil
	}
}

//## 典型场景 ##   删除媒资，支持批量删除。<br/>  ## 接口功能 ##   删除媒资，支持批量删除。<br/>  ## 接口约束 ##   最多删除十个。<br/>
func (c *VodClient) DeleteAssets(request *model.DeleteAssetsRequest) (*model.DeleteAssetsResponse, error) {
	requestDef := GenReqDefForDeleteAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetsResponse), nil
	}
}

//## 典型场景 ##   删除自定义模板组接口。<br/>  ## 接口功能 ##   删除自定义模板组。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) DeleteTemplateGroup(request *model.DeleteTemplateGroupRequest) (*model.DeleteTemplateGroupResponse, error) {
	requestDef := GenReqDefForDeleteTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateGroupResponse), nil
	}
}

//## 典型场景 ##   删除水印模板<br/>  ## 接口功能 ##   删除水印模板<br/>  ## 接口约束 ##   无<br/>
func (c *VodClient) DeleteWatermarkTemplate(request *model.DeleteWatermarkTemplateRequest) (*model.DeleteWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForDeleteWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWatermarkTemplateResponse), nil
	}
}

//## 典型场景 ##   查询指定分类信息，及其子分类（即下一级分类）的列表。<br/>  ## 接口功能 ##   查询指定分类信息。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) ListAssetCategory(request *model.ListAssetCategoryRequest) (*model.ListAssetCategoryResponse, error) {
	requestDef := GenReqDefForListAssetCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetCategoryResponse), nil
	}
}

//查询媒资列表
func (c *VodClient) ListAssetList(request *model.ListAssetListRequest) (*model.ListAssetListResponse, error) {
	requestDef := GenReqDefForListAssetList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetListResponse), nil
	}
}

//## 典型场景 ##   查询模板组列表调用此接口。<br/>  ## 接口功能 ##   查询模板组列表。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) ListTemplateGroup(request *model.ListTemplateGroupRequest) (*model.ListTemplateGroupResponse, error) {
	requestDef := GenReqDefForListTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateGroupResponse), nil
	}
}

//## 典型场景 ##  查询TopN播放视频信息 。<br/>  ## 接口功能 ##  查询TopN播放视频信息 。<br/>  ##  接口约束 ##   无。<br/>
func (c *VodClient) ListTopStatistics(request *model.ListTopStatisticsRequest) (*model.ListTopStatisticsResponse, error) {
	requestDef := GenReqDefForListTopStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopStatisticsResponse), nil
	}
}

//## 典型场景 ##   查询水印模板<br/>  ## 接口功能 ##   查询水印模板<br/>  ## 接口约束 ##   查询所有水印<br/>
func (c *VodClient) ListWatermarkTemplate(request *model.ListWatermarkTemplateRequest) (*model.ListWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForListWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWatermarkTemplateResponse), nil
	}
}

//## 典型场景 ##   从OBS转存媒资,一键发布。<br/>  ## 接口功能 ##    在OBS中的媒资一键发布到VOD。<br/>  ## 接口约束 ##   OBS的桶必须先授权给VOD服务用户。<br/>
func (c *VodClient) PublishAssetFromObs(request *model.PublishAssetFromObsRequest) (*model.PublishAssetFromObsResponse, error) {
	requestDef := GenReqDefForPublishAssetFromObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishAssetFromObsResponse), nil
	}
}

//## 典型场景 ##   媒资发布,支持批量发布。<br/>  ## 接口功能 ##   媒资发布,支持批量发布。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) PublishAssets(request *model.PublishAssetsRequest) (*model.PublishAssetsResponse, error) {
	requestDef := GenReqDefForPublishAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishAssetsResponse), nil
	}
}

//查询指定媒资的详细信息
func (c *VodClient) ShowAssetDetail(request *model.ShowAssetDetailRequest) (*model.ShowAssetDetailResponse, error) {
	requestDef := GenReqDefForShowAssetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetDetailResponse), nil
	}
}

//## 典型场景 ##   查询媒资信息。<br/>  ## 接口功能 ##   查询媒资信息。<br/>  ## 接口约束 ##   最多同时查询10个。<br/>
func (c *VodClient) ShowAssetMeta(request *model.ShowAssetMetaRequest) (*model.ShowAssetMetaResponse, error) {
	requestDef := GenReqDefForShowAssetMeta()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetMetaResponse), nil
	}
}

//客户端请求创建媒资时，如果媒资文件超过100MB，需采用分段的方式向OBS上传，在每次与OBS交互前，客户端需通过此接口获取到授权方可与OBS交互。
func (c *VodClient) ShowAssetTempAuthority(request *model.ShowAssetTempAuthorityRequest) (*model.ShowAssetTempAuthorityResponse, error) {
	requestDef := GenReqDefForShowAssetTempAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetTempAuthorityResponse), nil
	}
}

//## 典型场景 ##  查询域名的cdn数据信息 。<br/>  ## 接口功能 ##  查询域名的cdn数据信息 。<br/>  ##  接口约束 ##   无。<br/>
func (c *VodClient) ShowCdnStatistics(request *model.ShowCdnStatisticsRequest) (*model.ShowCdnStatisticsResponse, error) {
	requestDef := GenReqDefForShowCdnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCdnStatisticsResponse), nil
	}
}

//## 典型场景 ##  用于查询点播源站相关的统计数据，包括流量、存储空间、转码时长等 。<br/>  ## 接口功能 ##  用于查询点播源站相关的统计数据，包括流量、存储空间、转码时长等 。<br/>  ##  接口约束 ##   无。<br/>
func (c *VodClient) ShowVodStatistics(request *model.ShowVodStatisticsRequest) (*model.ShowVodStatisticsResponse, error) {
	requestDef := GenReqDefForShowVodStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVodStatisticsResponse), nil
	}
}

//## 典型场景 ##   媒资取消发布。<br/>  ## 接口功能 ##   媒资取消发布。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UnpublishAssets(request *model.UnpublishAssetsRequest) (*model.UnpublishAssetsResponse, error) {
	requestDef := GenReqDefForUnpublishAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnpublishAssetsResponse), nil
	}
}

//## 典型场景 ##   媒资更新,单独上传封面、更新视频文件或更新已有封面。<br/>  ## 接口功能 ##   媒资更新。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateAsset(request *model.UpdateAssetRequest) (*model.UpdateAssetResponse, error) {
	requestDef := GenReqDefForUpdateAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetResponse), nil
	}
}

//## 典型场景 ##   修改媒资分类。<br/>  ## 接口功能 ##   修改媒资分类。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateAssetCategory(request *model.UpdateAssetCategoryRequest) (*model.UpdateAssetCategoryResponse, error) {
	requestDef := GenReqDefForUpdateAssetCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetCategoryResponse), nil
	}
}

//## 典型场景 ##   更新媒资信息。<br/>  ## 接口功能 ##   更新媒资信息。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateAssetMeta(request *model.UpdateAssetMetaRequest) (*model.UpdateAssetMetaResponse, error) {
	requestDef := GenReqDefForUpdateAssetMeta()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssetMetaResponse), nil
	}
}

//## 典型场景 ##   用户通过桶策略方式授权给VOD用户操作桶的权限。<br/>  ## 接口功能 ##   用户通过桶策略方式授权给VOD用户操作桶的权限。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateBucketAuthorized(request *model.UpdateBucketAuthorizedRequest) (*model.UpdateBucketAuthorizedResponse, error) {
	requestDef := GenReqDefForUpdateBucketAuthorized()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBucketAuthorizedResponse), nil
	}
}

//## 典型场景 ##   租户选择截图，来更新封面。<br/>  ## 接口功能 ##   租户选择截图，来更新封面 。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateCoverByThumbnail(request *model.UpdateCoverByThumbnailRequest) (*model.UpdateCoverByThumbnailResponse, error) {
	requestDef := GenReqDefForUpdateCoverByThumbnail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCoverByThumbnailResponse), nil
	}
}

//## 典型场景 ##  修改模板组接口。<br/>  ## 接口功能 ##   修改模板组。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) UpdateTemplateGroup(request *model.UpdateTemplateGroupRequest) (*model.UpdateTemplateGroupResponse, error) {
	requestDef := GenReqDefForUpdateTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateGroupResponse), nil
	}
}

//## 典型场景 ##   修改水印模板<br/>  ## 接口功能 ##   修改水印模板<br/>  ## 接口约束 ##   无<br/>
func (c *VodClient) UpdateWatermarkTemplate(request *model.UpdateWatermarkTemplateRequest) (*model.UpdateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForUpdateWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWatermarkTemplateResponse), nil
	}
}

//## 典型场景 ##    创建媒资：URL拉取注入。<br/>  ## 接口功能 ##    创建媒资：URL拉取注入。<br/>  ## 接口约束 ##    无<br/>
func (c *VodClient) UploadMetaDataByUrl(request *model.UploadMetaDataByUrlRequest) (*model.UploadMetaDataByUrlResponse, error) {
	requestDef := GenReqDefForUploadMetaDataByUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadMetaDataByUrlResponse), nil
	}
}

//## 典型场景 ##   点播发布后可向CDN预热。<br/>  ## 接口功能 ##   CDN预热。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) CreatePreheatingAsset(request *model.CreatePreheatingAssetRequest) (*model.CreatePreheatingAssetResponse, error) {
	requestDef := GenReqDefForCreatePreheatingAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePreheatingAssetResponse), nil
	}
}

//## 典型场景 ##   向CDN查询预热。<br/>  ## 接口功能 ##   查询CDN预热。<br/>  ## 接口约束 ##   无。<br/>
func (c *VodClient) ShowPreheatingAsset(request *model.ShowPreheatingAssetRequest) (*model.ShowPreheatingAssetResponse, error) {
	requestDef := GenReqDefForShowPreheatingAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPreheatingAssetResponse), nil
	}
}

//## 典型场景 ##   终端在尝试播放加密HLS时会先向租户Server请求密钥，租户Server则对终端身份进行认证，认证通过后再向VOD查询密钥<br/>  ## 接口功能 ##   查询密钥<br/>  ## 接口约束 ##   暂无<br/>
func (c *VodClient) ShowAssetCipher(request *model.ShowAssetCipherRequest) (*model.ShowAssetCipherResponse, error) {
	requestDef := GenReqDefForShowAssetCipher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetCipherResponse), nil
	}
}

//## 典型场景 ##   创建托管任务<br/>  ## 接口功能 ##   <br/>  ## 接口约束 ##<br/>
func (c *VodClient) CreateTakeOverTask(request *model.CreateTakeOverTaskRequest) (*model.CreateTakeOverTaskResponse, error) {
	requestDef := GenReqDefForCreateTakeOverTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTakeOverTaskResponse), nil
	}
}

//## 典型场景 ##   查询托管任务<br/>  ## 接口功能 ##   查询托管任务<br/>  ## 接口约束 ##<br/>
func (c *VodClient) ListTakeOverTask(request *model.ListTakeOverTaskRequest) (*model.ListTakeOverTaskResponse, error) {
	requestDef := GenReqDefForListTakeOverTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTakeOverTaskResponse), nil
	}
}

//## 典型场景 ##   查询托管媒资详情<br/>  ## 接口功能 ##   查询托管媒资详情<br/>  ## 接口约束 ##<br/>
func (c *VodClient) ShowTakeOverAssetDetails(request *model.ShowTakeOverAssetDetailsRequest) (*model.ShowTakeOverAssetDetailsResponse, error) {
	requestDef := GenReqDefForShowTakeOverAssetDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTakeOverAssetDetailsResponse), nil
	}
}

//## 典型场景 ##   查询托管任务详情<br/>  ## 接口功能 ##   查询托管任务详情<br/>  ## 接口约束 ##<br/>
func (c *VodClient) ShowTakeOverTaskDetails(request *model.ShowTakeOverTaskDetailsRequest) (*model.ShowTakeOverTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowTakeOverTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTakeOverTaskDetailsResponse), nil
	}
}
