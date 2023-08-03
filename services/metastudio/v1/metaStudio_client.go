package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type MetaStudioClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMetaStudioClient(hcClient *http_client.HcHttpClient) *MetaStudioClient {
	return &MetaStudioClient{HcClient: hcClient}
}

func MetaStudioClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateDigitalAsset 创建资产
//
// 该接口用于在资产库中添加上传新的媒体资产。可上传的资产类型包括：数字人模型、素材、视频、图片、场景等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateDigitalAsset(request *model.CreateDigitalAssetRequest) (*model.CreateDigitalAssetResponse, error) {
	requestDef := GenReqDefForCreateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDigitalAssetResponse), nil
	}
}

// CreateDigitalAssetInvoker 创建资产
func (c *MetaStudioClient) CreateDigitalAssetInvoker(request *model.CreateDigitalAssetRequest) *CreateDigitalAssetInvoker {
	requestDef := GenReqDefForCreateDigitalAsset()
	return &CreateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAsset 删除资产
//
// 该接口用于删除资产库中的媒体资产。第一次调用删除接口，将指定资产放入回收站；第二次调用删除接口，将指定资产彻底删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// DeleteAssetInvoker 删除资产
func (c *MetaStudioClient) DeleteAssetInvoker(request *model.DeleteAssetRequest) *DeleteAssetInvoker {
	requestDef := GenReqDefForDeleteAsset()
	return &DeleteAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssetSummary 查询资产概要
//
// 该接口用于查询媒体资产库中指定的多个资产的概要信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssetSummary(request *model.ListAssetSummaryRequest) (*model.ListAssetSummaryResponse, error) {
	requestDef := GenReqDefForListAssetSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetSummaryResponse), nil
	}
}

// ListAssetSummaryInvoker 查询资产概要
func (c *MetaStudioClient) ListAssetSummaryInvoker(request *model.ListAssetSummaryRequest) *ListAssetSummaryInvoker {
	requestDef := GenReqDefForListAssetSummary()
	return &ListAssetSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssets 查询资产列表
//
// 该接口用于查询资产库中的媒体资产列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListAssets(request *model.ListAssetsRequest) (*model.ListAssetsResponse, error) {
	requestDef := GenReqDefForListAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetsResponse), nil
	}
}

// ListAssetsInvoker 查询资产列表
func (c *MetaStudioClient) ListAssetsInvoker(request *model.ListAssetsRequest) *ListAssetsInvoker {
	requestDef := GenReqDefForListAssets()
	return &ListAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreAsset 恢复被删除的资产
//
// 该接口用于恢复被删除至回收站的媒体资产。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) RestoreAsset(request *model.RestoreAssetRequest) (*model.RestoreAssetResponse, error) {
	requestDef := GenReqDefForRestoreAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreAssetResponse), nil
	}
}

// RestoreAssetInvoker 恢复被删除的资产
func (c *MetaStudioClient) RestoreAssetInvoker(request *model.RestoreAssetRequest) *RestoreAssetInvoker {
	requestDef := GenReqDefForRestoreAsset()
	return &RestoreAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsset 查询资产详情
//
// 该接口用于查询资产库中指定媒体资产的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 查询资产详情
func (c *MetaStudioClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDigitalAsset 更新资产
//
// 该接口用于更新资产库中的媒体资产信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) UpdateDigitalAsset(request *model.UpdateDigitalAssetRequest) (*model.UpdateDigitalAssetResponse, error) {
	requestDef := GenReqDefForUpdateDigitalAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDigitalAssetResponse), nil
	}
}

// UpdateDigitalAssetInvoker 更新资产
func (c *MetaStudioClient) UpdateDigitalAssetInvoker(request *model.UpdateDigitalAssetRequest) *UpdateDigitalAssetInvoker {
	requestDef := GenReqDefForUpdateDigitalAsset()
	return &UpdateDigitalAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmFileUpload 确认文件已上传
//
// 资产文件上传完毕后，通过该接口确认上传完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ConfirmFileUpload(request *model.ConfirmFileUploadRequest) (*model.ConfirmFileUploadResponse, error) {
	requestDef := GenReqDefForConfirmFileUpload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmFileUploadResponse), nil
	}
}

// ConfirmFileUploadInvoker 确认文件已上传
func (c *MetaStudioClient) ConfirmFileUploadInvoker(request *model.ConfirmFileUploadRequest) *ConfirmFileUploadInvoker {
	requestDef := GenReqDefForConfirmFileUpload()
	return &ConfirmFileUploadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建文件并获取上传URL
//
// 该接口用于创建文件并获取上传URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建文件并获取上传URL
func (c *MetaStudioClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除文件
//
// 该接口用于删除媒体资产库中指定的文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除文件
func (c *MetaStudioClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingByUrlJob 基于图片URL创建照片建模任务
//
// 该接口用于从URL中获取图片进行照片建模任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingByUrlJob(request *model.CreatePictureModelingByUrlJobRequest) (*model.CreatePictureModelingByUrlJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingByUrlJobResponse), nil
	}
}

// CreatePictureModelingByUrlJobInvoker 基于图片URL创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingByUrlJobInvoker(request *model.CreatePictureModelingByUrlJobRequest) *CreatePictureModelingByUrlJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingByUrlJob()
	return &CreatePictureModelingByUrlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePictureModelingJob 创建照片建模任务
//
// 该接口用于创建风格化照片建模任务。通过上传照片，生成风格化数字人模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreatePictureModelingJob(request *model.CreatePictureModelingJobRequest) (*model.CreatePictureModelingJobResponse, error) {
	requestDef := GenReqDefForCreatePictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePictureModelingJobResponse), nil
	}
}

// CreatePictureModelingJobInvoker 创建照片建模任务
func (c *MetaStudioClient) CreatePictureModelingJobInvoker(request *model.CreatePictureModelingJobRequest) *CreatePictureModelingJobInvoker {
	requestDef := GenReqDefForCreatePictureModelingJob()
	return &CreatePictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPictureModelingJobs 照片建模任务列表查询
//
// 该接口用于查询风格化照片建模任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListPictureModelingJobs(request *model.ListPictureModelingJobsRequest) (*model.ListPictureModelingJobsResponse, error) {
	requestDef := GenReqDefForListPictureModelingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPictureModelingJobsResponse), nil
	}
}

// ListPictureModelingJobsInvoker 照片建模任务列表查询
func (c *MetaStudioClient) ListPictureModelingJobsInvoker(request *model.ListPictureModelingJobsRequest) *ListPictureModelingJobsInvoker {
	requestDef := GenReqDefForListPictureModelingJobs()
	return &ListPictureModelingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPictureModelingJob 照片建模任务详情查询
//
// 该接口用于风格化查询照片建模任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowPictureModelingJob(request *model.ShowPictureModelingJobRequest) (*model.ShowPictureModelingJobResponse, error) {
	requestDef := GenReqDefForShowPictureModelingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPictureModelingJobResponse), nil
	}
}

// ShowPictureModelingJobInvoker 照片建模任务详情查询
func (c *MetaStudioClient) ShowPictureModelingJobInvoker(request *model.ShowPictureModelingJobRequest) *ShowPictureModelingJobInvoker {
	requestDef := GenReqDefForShowPictureModelingJob()
	return &ShowPictureModelingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStyles 查询数字人风格列表
//
// 查询数字人风格列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListStyles(request *model.ListStylesRequest) (*model.ListStylesResponse, error) {
	requestDef := GenReqDefForListStyles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStylesResponse), nil
	}
}

// ListStylesInvoker 查询数字人风格列表
func (c *MetaStudioClient) ListStylesInvoker(request *model.ListStylesRequest) *ListStylesInvoker {
	requestDef := GenReqDefForListStyles()
	return &ListStylesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTtsa 创建语音驱动任务
//
// 该接口用于创建驱动数字人表情、动作及语音的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateTtsa(request *model.CreateTtsaRequest) (*model.CreateTtsaResponse, error) {
	requestDef := GenReqDefForCreateTtsa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTtsaResponse), nil
	}
}

// CreateTtsaInvoker 创建语音驱动任务
func (c *MetaStudioClient) CreateTtsaInvoker(request *model.CreateTtsaRequest) *CreateTtsaInvoker {
	requestDef := GenReqDefForCreateTtsa()
	return &CreateTtsaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtsaData 获取语音驱动数据
//
// 该接口用于获取生成的数字人驱动数据，包括语音、表情、动作等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtsaData(request *model.ListTtsaDataRequest) (*model.ListTtsaDataResponse, error) {
	requestDef := GenReqDefForListTtsaData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtsaDataResponse), nil
	}
}

// ListTtsaDataInvoker 获取语音驱动数据
func (c *MetaStudioClient) ListTtsaDataInvoker(request *model.ListTtsaDataRequest) *ListTtsaDataInvoker {
	requestDef := GenReqDefForListTtsaData()
	return &ListTtsaDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTtsaJobs 获取语音驱动任务列表
//
// 该接口用于查询驱动数字人表情、动作及语音的任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListTtsaJobs(request *model.ListTtsaJobsRequest) (*model.ListTtsaJobsResponse, error) {
	requestDef := GenReqDefForListTtsaJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTtsaJobsResponse), nil
	}
}

// ListTtsaJobsInvoker 获取语音驱动任务列表
func (c *MetaStudioClient) ListTtsaJobsInvoker(request *model.ListTtsaJobsRequest) *ListTtsaJobsInvoker {
	requestDef := GenReqDefForListTtsaJobs()
	return &ListTtsaJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVideoMotionCaptureJob 创建视频驱动任务
//
// 该接口用于创建视频驱动任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) CreateVideoMotionCaptureJob(request *model.CreateVideoMotionCaptureJobRequest) (*model.CreateVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForCreateVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVideoMotionCaptureJobResponse), nil
	}
}

// CreateVideoMotionCaptureJobInvoker 创建视频驱动任务
func (c *MetaStudioClient) CreateVideoMotionCaptureJobInvoker(request *model.CreateVideoMotionCaptureJobRequest) *CreateVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForCreateVideoMotionCaptureJob()
	return &CreateVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteVideoMotionCaptureCommand 控制数字人驱动
//
// 该接口用于控制数字人驱动。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ExecuteVideoMotionCaptureCommand(request *model.ExecuteVideoMotionCaptureCommandRequest) (*model.ExecuteVideoMotionCaptureCommandResponse, error) {
	requestDef := GenReqDefForExecuteVideoMotionCaptureCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteVideoMotionCaptureCommandResponse), nil
	}
}

// ExecuteVideoMotionCaptureCommandInvoker 控制数字人驱动
func (c *MetaStudioClient) ExecuteVideoMotionCaptureCommandInvoker(request *model.ExecuteVideoMotionCaptureCommandRequest) *ExecuteVideoMotionCaptureCommandInvoker {
	requestDef := GenReqDefForExecuteVideoMotionCaptureCommand()
	return &ExecuteVideoMotionCaptureCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVideoMotionCaptureJobs 查询视频驱动任务列表
//
// 该接口用于查询视频驱动任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ListVideoMotionCaptureJobs(request *model.ListVideoMotionCaptureJobsRequest) (*model.ListVideoMotionCaptureJobsResponse, error) {
	requestDef := GenReqDefForListVideoMotionCaptureJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVideoMotionCaptureJobsResponse), nil
	}
}

// ListVideoMotionCaptureJobsInvoker 查询视频驱动任务列表
func (c *MetaStudioClient) ListVideoMotionCaptureJobsInvoker(request *model.ListVideoMotionCaptureJobsRequest) *ListVideoMotionCaptureJobsInvoker {
	requestDef := GenReqDefForListVideoMotionCaptureJobs()
	return &ListVideoMotionCaptureJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVideoMotionCaptureJob 查询视频驱动任务详情
//
// 该接口用于查询视频驱动任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) ShowVideoMotionCaptureJob(request *model.ShowVideoMotionCaptureJobRequest) (*model.ShowVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForShowVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVideoMotionCaptureJobResponse), nil
	}
}

// ShowVideoMotionCaptureJobInvoker 查询视频驱动任务详情
func (c *MetaStudioClient) ShowVideoMotionCaptureJobInvoker(request *model.ShowVideoMotionCaptureJobRequest) *ShowVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForShowVideoMotionCaptureJob()
	return &ShowVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopVideoMotionCaptureJob 停止视频驱动任务
//
// 该接口用于停止视频驱动任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MetaStudioClient) StopVideoMotionCaptureJob(request *model.StopVideoMotionCaptureJobRequest) (*model.StopVideoMotionCaptureJobResponse, error) {
	requestDef := GenReqDefForStopVideoMotionCaptureJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopVideoMotionCaptureJobResponse), nil
	}
}

// StopVideoMotionCaptureJobInvoker 停止视频驱动任务
func (c *MetaStudioClient) StopVideoMotionCaptureJobInvoker(request *model.StopVideoMotionCaptureJobRequest) *StopVideoMotionCaptureJobInvoker {
	requestDef := GenReqDefForStopVideoMotionCaptureJob()
	return &StopVideoMotionCaptureJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
