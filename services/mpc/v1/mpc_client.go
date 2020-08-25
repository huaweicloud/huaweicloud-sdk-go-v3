package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mpc/v1/model"
)

type MpcClient struct {
	hcClient *http_client.HcHttpClient
}

func NewMpcClient(hcClient *http_client.HcHttpClient) *MpcClient {
	return &MpcClient{hcClient: hcClient}
}

func MpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建动图任务，用于将完整的视频文件或视频文件中的一部分转换为动态图文件，暂只支持输出GIF文件。 待转动图的视频文件需要存储在与媒体处理服务同区域的OBS桶中，且该OBS桶已授权。
func (c *MpcClient) CreateAnimatedGraphicsTask(request *model.CreateAnimatedGraphicsTaskRequest) (*model.CreateAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForCreateAnimatedGraphicsTask(request)
	resp, responseDef := GenRespForCreateAnimatedGraphicsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消已下发的生成动图任务，仅支持取消正在排队中的任务。
func (c *MpcClient) DeleteAnimatedGraphicsTask(request *model.DeleteAnimatedGraphicsTaskRequest) (*model.DeleteAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForDeleteAnimatedGraphicsTask(request)
	resp, responseDef := GenRespForDeleteAnimatedGraphicsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询动图任务的状态。
func (c *MpcClient) ListAnimatedGraphicsTask(request *model.ListAnimatedGraphicsTaskRequest) (*model.ListAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForListAnimatedGraphicsTask(request)
	resp, responseDef := GenRespForListAnimatedGraphicsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//支持独立加密，包括创建、查询、删除独立加密任务。  约束：   - 只支持转码后的文件进行加密。   - 加密的文件必须是m3u8或者mpd结尾的文件。
func (c *MpcClient) CreateEncryptTask(request *model.CreateEncryptTaskRequest) (*model.CreateEncryptTaskResponse, error) {
	requestDef := GenReqDefForCreateEncryptTask(request)
	resp, responseDef := GenRespForCreateEncryptTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消独立加密任务。  约束：    只能取消正在任务队列中排队的任务。已开始加密或已完成的加密任务不能取消。
func (c *MpcClient) DeleteEncryptTask(request *model.DeleteEncryptTaskRequest) (*model.DeleteEncryptTaskResponse, error) {
	requestDef := GenReqDefForDeleteEncryptTask(request)
	resp, responseDef := GenRespForDeleteEncryptTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询独立加密任务状态。返回任务执行结果或当前状态。
func (c *MpcClient) ListEncryptTask(request *model.ListEncryptTaskRequest) (*model.ListEncryptTaskResponse, error) {
	requestDef := GenReqDefForListEncryptTask(request)
	resp, responseDef := GenRespForListEncryptTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建视频解析任务，解析视频元数据。
func (c *MpcClient) CreateExtractTask(request *model.CreateExtractTaskRequest) (*model.CreateExtractTaskResponse, error) {
	requestDef := GenReqDefForCreateExtractTask(request)
	resp, responseDef := GenRespForCreateExtractTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消已下发的视频解析任务，仅支持取消正在排队中的任务。
func (c *MpcClient) DeleteExtractTask(request *model.DeleteExtractTaskRequest) (*model.DeleteExtractTaskResponse, error) {
	requestDef := GenReqDefForDeleteExtractTask(request)
	resp, responseDef := GenRespForDeleteExtractTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询解析任务的状态和结果。
func (c *MpcClient) ListExtractTask(request *model.ListExtractTaskRequest) (*model.ListExtractTaskResponse, error) {
	requestDef := GenReqDefForListExtractTask(request)
	resp, responseDef := GenRespForListExtractTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询媒资转码详情
func (c *MpcClient) ListTranscodeDetail(request *model.ListTranscodeDetailRequest) (*model.ListTranscodeDetailResponse, error) {
	requestDef := GenReqDefForListTranscodeDetail(request)
	resp, responseDef := GenRespForListTranscodeDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消已下发的转封装任务，仅支持取消正在排队中的任务。。
func (c *MpcClient) CancelRemuxTask(request *model.CancelRemuxTaskRequest) (*model.CancelRemuxTaskResponse, error) {
	requestDef := GenReqDefForCancelRemuxTask(request)
	resp, responseDef := GenRespForCancelRemuxTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建转封装任务，转换音视频文件的格式，但不改变其分辨率和码率。 待转封装的媒资文件需要存储在与媒体处理服务同区域的OBS桶中，且该OBS桶已授权。
func (c *MpcClient) CreateRemuxTask(request *model.CreateRemuxTaskRequest) (*model.CreateRemuxTaskResponse, error) {
	requestDef := GenReqDefForCreateRemuxTask(request)
	resp, responseDef := GenRespForCreateRemuxTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//对失败的转封装任务进行重试。
func (c *MpcClient) CreateRetryRemuxTask(request *model.CreateRetryRemuxTaskRequest) (*model.CreateRetryRemuxTaskResponse, error) {
	requestDef := GenReqDefForCreateRetryRemuxTask(request)
	resp, responseDef := GenRespForCreateRetryRemuxTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除转封装任务
func (c *MpcClient) DeleteRemuxTask(request *model.DeleteRemuxTaskRequest) (*model.DeleteRemuxTaskResponse, error) {
	requestDef := GenReqDefForDeleteRemuxTask(request)
	resp, responseDef := GenRespForDeleteRemuxTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转封装任务状态。
func (c *MpcClient) ListRemuxTask(request *model.ListRemuxTaskRequest) (*model.ListRemuxTaskResponse, error) {
	requestDef := GenReqDefForListRemuxTask(request)
	resp, responseDef := GenRespForListRemuxTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//新建转码模板组，最多支持一进六出。
func (c *MpcClient) CreateTemplateGroup(request *model.CreateTemplateGroupRequest) (*model.CreateTemplateGroupResponse, error) {
	requestDef := GenReqDefForCreateTemplateGroup(request)
	resp, responseDef := GenRespForCreateTemplateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除转码模板组。
func (c *MpcClient) DeleteTemplateGroup(request *model.DeleteTemplateGroupRequest) (*model.DeleteTemplateGroupResponse, error) {
	requestDef := GenReqDefForDeleteTemplateGroup(request)
	resp, responseDef := GenRespForDeleteTemplateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转码模板组列表。
func (c *MpcClient) ListTemplateGroup(request *model.ListTemplateGroupRequest) (*model.ListTemplateGroupResponse, error) {
	requestDef := GenReqDefForListTemplateGroup(request)
	resp, responseDef := GenRespForListTemplateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//修改模板组接口。
func (c *MpcClient) UpdateTemplateGroup(request *model.UpdateTemplateGroupRequest) (*model.UpdateTemplateGroupResponse, error) {
	requestDef := GenReqDefForUpdateTemplateGroup(request)
	resp, responseDef := GenRespForUpdateTemplateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//新建截图任务，视频截图将从首帧开始，按设置的时间间隔截图，最后截取末帧。 待截图的视频文件需要存储在与媒体处理服务同区域的OBS桶中，且该OBS桶已授权。  约束：   暂只支持生成JPG格式的图片文件。
func (c *MpcClient) CreateThumbnailsTask(request *model.CreateThumbnailsTaskRequest) (*model.CreateThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForCreateThumbnailsTask(request)
	resp, responseDef := GenRespForCreateThumbnailsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消已下发截图任务。 只能取消已接受尚在队列中等待处理的任务，已完成或正在执行阶段的任务不能取消。
func (c *MpcClient) DeleteThumbnailsTask(request *model.DeleteThumbnailsTaskRequest) (*model.DeleteThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForDeleteThumbnailsTask(request)
	resp, responseDef := GenRespForDeleteThumbnailsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询截图任务状态。返回任务执行结果，包括状态、输入、输出等信息。
func (c *MpcClient) ListThumbnailsTask(request *model.ListThumbnailsTaskRequest) (*model.ListThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForListThumbnailsTask(request)
	resp, responseDef := GenRespForListThumbnailsTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//新建转码任务可以将视频进行转码，并在转码过程中压制水印、内容质检、视频截图等。视频转码前需要配置转码模板。 待转码的音视频需要存储在与媒体处理服务同区域的OBS桶中，且该OBS桶已授权。
func (c *MpcClient) CreateTranscodingTask(request *model.CreateTranscodingTaskRequest) (*model.CreateTranscodingTaskResponse, error) {
	requestDef := GenReqDefForCreateTranscodingTask(request)
	resp, responseDef := GenRespForCreateTranscodingTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//取消已下发转码任务。 只能取消正在转码任务队列中排队的转码任务。已开始转码或已完成的转码任务不能取消。
func (c *MpcClient) DeleteTranscodingTask(request *model.DeleteTranscodingTaskRequest) (*model.DeleteTranscodingTaskResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingTask(request)
	resp, responseDef := GenRespForDeleteTranscodingTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转码任务状态。
func (c *MpcClient) ListTranscodingTask(request *model.ListTranscodingTaskRequest) (*model.ListTranscodingTaskResponse, error) {
	requestDef := GenReqDefForListTranscodingTask(request)
	resp, responseDef := GenRespForListTranscodingTask()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//新建转码模板，采用自定义的模板转码。
func (c *MpcClient) CreateTransTemplate(request *model.CreateTransTemplateRequest) (*model.CreateTransTemplateResponse, error) {
	requestDef := GenReqDefForCreateTransTemplate(request)
	resp, responseDef := GenRespForCreateTransTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除转码模板。
func (c *MpcClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate(request)
	resp, responseDef := GenRespForDeleteTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询用户自定义转码配置模板。 支持指定模板ID查询，或分页全量查询。转码配置模板ID，最多10个。
func (c *MpcClient) ListTemplate(request *model.ListTemplateRequest) (*model.ListTemplateResponse, error) {
	requestDef := GenReqDefForListTemplate(request)
	resp, responseDef := GenRespForListTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新转码模板。
func (c *MpcClient) UpdateTransTemplate(request *model.UpdateTransTemplateRequest) (*model.UpdateTransTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTransTemplate(request)
	resp, responseDef := GenRespForUpdateTransTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//自定义水印模板。
func (c *MpcClient) CreateWatermarkTemplate(request *model.CreateWatermarkTemplateRequest) (*model.CreateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForCreateWatermarkTemplate(request)
	resp, responseDef := GenRespForCreateWatermarkTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除自定义水印模板。
func (c *MpcClient) DeleteWatermarkTemplate(request *model.DeleteWatermarkTemplateRequest) (*model.DeleteWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForDeleteWatermarkTemplate(request)
	resp, responseDef := GenRespForDeleteWatermarkTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询自定义水印模板。支持指定模板ID查询，或分页全量查询。
func (c *MpcClient) ListWatermarkTemplate(request *model.ListWatermarkTemplateRequest) (*model.ListWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForListWatermarkTemplate(request)
	resp, responseDef := GenRespForListWatermarkTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新自定义水印模板。
func (c *MpcClient) UpdateWatermarkTemplate(request *model.UpdateWatermarkTemplateRequest) (*model.UpdateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForUpdateWatermarkTemplate(request)
	resp, responseDef := GenRespForUpdateWatermarkTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
