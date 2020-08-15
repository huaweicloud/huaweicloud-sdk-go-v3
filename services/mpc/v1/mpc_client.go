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

//云服务操作异步查询接口：云运营系统通过此接口，异步查询云服务的操作结果。
func (c *MpcClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus(request)
	resp, responseDef := GenRespForShowJobStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询桶列表（仅供Console调用）。
func (c *MpcClient) ListAllBuckets(request *model.ListAllBucketsRequest) (*model.ListAllBucketsResponse, error) {
	requestDef := GenReqDefForListAllBuckets(request)
	resp, responseDef := GenRespForListAllBuckets()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//桶授权或取消授权（仅供Console调用）。
func (c *MpcClient) UpdateBucketAuthorized(request *model.UpdateBucketAuthorizedRequest) (*model.UpdateBucketAuthorizedResponse, error) {
	requestDef := GenReqDefForUpdateBucketAuthorized(request)
	resp, responseDef := GenRespForUpdateBucketAuthorized()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询桶里的object（仅供Console调用）。
func (c *MpcClient) ListAllObsObjList(request *model.ListAllObsObjListRequest) (*model.ListAllObsObjListResponse, error) {
	requestDef := GenReqDefForListAllObsObjList(request)
	resp, responseDef := GenRespForListAllObsObjList()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建视频增强模板
func (c *MpcClient) CreateQualityEnhanceTemplate(request *model.CreateQualityEnhanceTemplateRequest) (*model.CreateQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForCreateQualityEnhanceTemplate(request)
	resp, responseDef := GenRespForCreateQualityEnhanceTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除用户视频增强模板。
func (c *MpcClient) DeleteQualityEnhanceTemplate(request *model.DeleteQualityEnhanceTemplateRequest) (*model.DeleteQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForDeleteQualityEnhanceTemplate(request)
	resp, responseDef := GenRespForDeleteQualityEnhanceTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询视频增强预置模板，返回所有结果。
func (c *MpcClient) ListQualityEnhanceDefaultTemplate(request *model.ListQualityEnhanceDefaultTemplateRequest) (*model.ListQualityEnhanceDefaultTemplateResponse, error) {
	requestDef := GenReqDefForListQualityEnhanceDefaultTemplate(request)
	resp, responseDef := GenRespForListQualityEnhanceDefaultTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询用户自定义视频增强模板。 支持指定模板ID查询，或分页全量查询。模板ID，最多10个。
func (c *MpcClient) ListQualityEnhanceTemplate(request *model.ListQualityEnhanceTemplateRequest) (*model.ListQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForListQualityEnhanceTemplate(request)
	resp, responseDef := GenRespForListQualityEnhanceTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新视频增强模板。
func (c *MpcClient) UpdateQualityEnhanceTemplate(request *model.UpdateQualityEnhanceTemplateRequest) (*model.UpdateQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForUpdateQualityEnhanceTemplate(request)
	resp, responseDef := GenRespForUpdateQualityEnhanceTemplate()

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

//## 典型场景 ##   转码Console查询转码服务端所有事件，并将查询到的事件展示在页面供用户配置  ## 接口功能 ##   查询转码服务端所有事件 。  ## 接口约束 ##   无。
func (c *MpcClient) ListNotifyEvent(request *model.ListNotifyEventRequest) (*model.ListNotifyEventResponse, error) {
	requestDef := GenReqDefForListNotifyEvent(request)
	resp, responseDef := GenRespForListNotifyEvent()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//## 典型场景 ##   查询转码服务端事件通知。  ## 接口功能 ##   查询转码服务端事件通知。  ## 接口约束 ##   无。
func (c *MpcClient) ListNotifySmnTopicConfig(request *model.ListNotifySmnTopicConfigRequest) (*model.ListNotifySmnTopicConfigResponse, error) {
	requestDef := GenReqDefForListNotifySmnTopicConfig(request)
	resp, responseDef := GenRespForListNotifySmnTopicConfig()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询规定的消息通知模板内容。
func (c *MpcClient) ListNotifyTemplate(request *model.ListNotifyTemplateRequest) (*model.ListNotifyTemplateResponse, error) {
	requestDef := GenReqDefForListNotifyTemplate(request)
	resp, responseDef := GenRespForListNotifyTemplate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//配置转码服务端事件通知。
func (c *MpcClient) NotifySmnTopicConfig(request *model.NotifySmnTopicConfigRequest) (*model.NotifySmnTopicConfigResponse, error) {
	requestDef := GenReqDefForNotifySmnTopicConfig(request)
	resp, responseDef := GenRespForNotifySmnTopicConfig()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询点播概览信息（仅供Console调用）。
func (c *MpcClient) ListStatSummary(request *model.ListStatSummaryRequest) (*model.ListStatSummaryResponse, error) {
	requestDef := GenReqDefForListStatSummary(request)
	resp, responseDef := GenRespForListStatSummary()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转码概览信息，（仅供Console调用）。 转码Console有个概览页面，用于展示登录租户最近一个月或最近一周转码时长（分钟）、转码任务数量。仅展示转码成功的，按结束时间算。
func (c *MpcClient) ListSummary(request *model.ListSummaryRequest) (*model.ListSummaryResponse, error) {
	requestDef := GenReqDefForListSummary(request)
	resp, responseDef := GenRespForListSummary()

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

//## 典型场景 ##   更新租户状态。  ## 接口功能 ##   更新租户状态。  ## 接口约束 ##   无。
func (c *MpcClient) UpdateTenantStatus(request *model.UpdateTenantStatusRequest) (*model.UpdateTenantStatusResponse, error) {
	requestDef := GenReqDefForUpdateTenantStatus(request)
	resp, responseDef := GenRespForUpdateTenantStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//## 典型场景 ##   查询租户信息，查询租户的欠费、冻结状态、是否实名认证、是否开通服务  ## 接口功能 ##   查询租户信息，查询租户的欠费、冻结状态、是否实名认证、是否开通服务  ## 接口约束 ##   无
func (c *MpcClient) ShowTenantInfo(request *model.ShowTenantInfoRequest) (*model.ShowTenantInfoResponse, error) {
	requestDef := GenReqDefForShowTenantInfo(request)
	resp, responseDef := GenRespForShowTenantInfo()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//## 典型场景 ##   修改租户信息，如开通点播服务  ## 接口功能 ##   修改租户信息，如开通点播服务  ## 接口约束 ##   无
func (c *MpcClient) UpdateTenantInfo(request *model.UpdateTenantInfoRequest) (*model.UpdateTenantInfoResponse, error) {
	requestDef := GenReqDefForUpdateTenantInfo(request)
	resp, responseDef := GenRespForUpdateTenantInfo()

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

//删除转码任务(仅供Console调用)。
func (c *MpcClient) DeleteTranscodingTaskByConsole(request *model.DeleteTranscodingTaskByConsoleRequest) (*model.DeleteTranscodingTaskByConsoleResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingTaskByConsole(request)
	resp, responseDef := GenRespForDeleteTranscodingTaskByConsole()

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
