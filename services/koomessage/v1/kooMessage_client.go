package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/koomessage/v1/model"
)

type KooMessageClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewKooMessageClient(hcClient *httpclient.HcHttpClient) *KooMessageClient {
	return &KooMessageClient{HcClient: hcClient}
}

func KooMessageClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddCallBack 注册智能信息回执URL
//
// 用户根据要求创建回执接口后，可以调用此接口进行注册，注意：此接口仅允许te_admin角色用户调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) AddCallBack(request *model.AddCallBackRequest) (*model.AddCallBackResponse, error) {
	requestDef := GenReqDefForAddCallBack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCallBackResponse), nil
	}
}

// AddCallBackInvoker 注册智能信息回执URL
func (c *KooMessageClient) AddCallBackInvoker(request *model.AddCallBackRequest) *AddCallBackInvoker {
	requestDef := GenReqDefForAddCallBack()
	return &AddCallBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimCallbacks 查询用户已注册回执接口
//
// 用户注册回执接口之后，可以根据此接口查询所有已注册回执接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimCallbacks(request *model.ListAimCallbacksRequest) (*model.ListAimCallbacksResponse, error) {
	requestDef := GenReqDefForListAimCallbacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimCallbacksResponse), nil
	}
}

// ListAimCallbacksInvoker 查询用户已注册回执接口
func (c *KooMessageClient) ListAimCallbacksInvoker(request *model.ListAimCallbacksRequest) *ListAimCallbacksInvoker {
	requestDef := GenReqDefForListAimCallbacks()
	return &ListAimCallbacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckMobileCapability 查询手机号智能信息解析能力
//
// 用户在下发智能信息前，通过此接口批量查询对应手机的智能信息解析能力。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CheckMobileCapability(request *model.CheckMobileCapabilityRequest) (*model.CheckMobileCapabilityResponse, error) {
	requestDef := GenReqDefForCheckMobileCapability()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckMobileCapabilityResponse), nil
	}
}

// CheckMobileCapabilityInvoker 查询手机号智能信息解析能力
func (c *KooMessageClient) CheckMobileCapabilityInvoker(request *model.CheckMobileCapabilityRequest) *CheckMobileCapabilityInvoker {
	requestDef := GenReqDefForCheckMobileCapability()
	return &CheckMobileCapabilityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResolveTask 生成解析任务
//
// 生成解析的短链。一次最多生成100个解析的短链。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateResolveTask(request *model.CreateResolveTaskRequest) (*model.CreateResolveTaskResponse, error) {
	requestDef := GenReqDefForCreateResolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResolveTaskResponse), nil
	}
}

// CreateResolveTaskInvoker 生成解析任务
func (c *KooMessageClient) CreateResolveTaskInvoker(request *model.CreateResolveTaskRequest) *CreateResolveTaskInvoker {
	requestDef := GenReqDefForCreateResolveTask()
	return &CreateResolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimResolveDetails 查询解析明细
//
// 根据用户提供的过滤条件查询个性化解析明细，包括：发送任务ID、发送手机号码等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimResolveDetails(request *model.ListAimResolveDetailsRequest) (*model.ListAimResolveDetailsResponse, error) {
	requestDef := GenReqDefForListAimResolveDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimResolveDetailsResponse), nil
	}
}

// ListAimResolveDetailsInvoker 查询解析明细
func (c *KooMessageClient) ListAimResolveDetailsInvoker(request *model.ListAimResolveDetailsRequest) *ListAimResolveDetailsInvoker {
	requestDef := GenReqDefForListAimResolveDetails()
	return &ListAimResolveDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResolveTasks 查询解析任务
//
// 创建解析任务后，客户可以查询解析任务状态信息。
//
// 通过模板ID等过滤条件，查询解析任务列表。
//
// 如需查询明细，建议使用查询解析明细接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListResolveTasks(request *model.ListResolveTasksRequest) (*model.ListResolveTasksResponse, error) {
	requestDef := GenReqDefForListResolveTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResolveTasksResponse), nil
	}
}

// ListResolveTasksInvoker 查询解析任务
func (c *KooMessageClient) ListResolveTasksInvoker(request *model.ListResolveTasksRequest) *ListResolveTasksInvoker {
	requestDef := GenReqDefForListResolveTasks()
	return &ListResolveTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAimSendTask 发送智能信息
//
// 根据客户的参数发送任务名称、智能信息模板ID等进行智能信息发送。一次最多发送100个智能信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateAimSendTask(request *model.CreateAimSendTaskRequest) (*model.CreateAimSendTaskResponse, error) {
	requestDef := GenReqDefForCreateAimSendTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAimSendTaskResponse), nil
	}
}

// CreateAimSendTaskInvoker 发送智能信息
func (c *KooMessageClient) CreateAimSendTaskInvoker(request *model.CreateAimSendTaskRequest) *CreateAimSendTaskInvoker {
	requestDef := GenReqDefForCreateAimSendTask()
	return &CreateAimSendTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimSendDetails 查询智能信息发送明细
//
// 根据用户提供的过滤条件查询发送明细列表，包括：发送任务ID、发送手机号码等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimSendDetails(request *model.ListAimSendDetailsRequest) (*model.ListAimSendDetailsResponse, error) {
	requestDef := GenReqDefForListAimSendDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimSendDetailsResponse), nil
	}
}

// ListAimSendDetailsInvoker 查询智能信息发送明细
func (c *KooMessageClient) ListAimSendDetailsInvoker(request *model.ListAimSendDetailsRequest) *ListAimSendDetailsInvoker {
	requestDef := GenReqDefForListAimSendDetails()
	return &ListAimSendDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimSendReports 查询智能信息发送报表
//
// 查询智能信息发送报表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimSendReports(request *model.ListAimSendReportsRequest) (*model.ListAimSendReportsResponse, error) {
	requestDef := GenReqDefForListAimSendReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimSendReportsResponse), nil
	}
}

// ListAimSendReportsInvoker 查询智能信息发送报表
func (c *KooMessageClient) ListAimSendReportsInvoker(request *model.ListAimSendReportsRequest) *ListAimSendReportsInvoker {
	requestDef := GenReqDefForListAimSendReports()
	return &ListAimSendReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimSendTasks 查询智能信息发送任务
//
// 根据用户提供的过滤条件查询智能信息发送任务列表，包括：发送任务名称、智能信息模板ID等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimSendTasks(request *model.ListAimSendTasksRequest) (*model.ListAimSendTasksResponse, error) {
	requestDef := GenReqDefForListAimSendTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimSendTasksResponse), nil
	}
}

// ListAimSendTasksInvoker 查询智能信息发送任务
func (c *KooMessageClient) ListAimSendTasksInvoker(request *model.ListAimSendTasksRequest) *ListAimSendTasksInvoker {
	requestDef := GenReqDefForListAimSendTasks()
	return &ListAimSendTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAimPersonalTemplate 创建个人模板
//
// 用于用户创建个人模板。
//
// &gt; 请求中所有字符串不允许携带“&lt;”、“&gt;”、“\\&amp;amp;amp;”或多个空格。
// &gt; 模板内容需加“拒收请回复R”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateAimPersonalTemplate(request *model.CreateAimPersonalTemplateRequest) (*model.CreateAimPersonalTemplateResponse, error) {
	requestDef := GenReqDefForCreateAimPersonalTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAimPersonalTemplateResponse), nil
	}
}

// CreateAimPersonalTemplateInvoker 创建个人模板
func (c *KooMessageClient) CreateAimPersonalTemplateInvoker(request *model.CreateAimPersonalTemplateRequest) *CreateAimPersonalTemplateInvoker {
	requestDef := GenReqDefForCreateAimPersonalTemplate()
	return &CreateAimPersonalTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAimPersonalTemplate 删除模板实例
//
// 根据用户提供的模板ID，删除智能信息个人模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) DeleteAimPersonalTemplate(request *model.DeleteAimPersonalTemplateRequest) (*model.DeleteAimPersonalTemplateResponse, error) {
	requestDef := GenReqDefForDeleteAimPersonalTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAimPersonalTemplateResponse), nil
	}
}

// DeleteAimPersonalTemplateInvoker 删除模板实例
func (c *KooMessageClient) DeleteAimPersonalTemplateInvoker(request *model.DeleteAimPersonalTemplateRequest) *DeleteAimPersonalTemplateInvoker {
	requestDef := GenReqDefForDeleteAimPersonalTemplate()
	return &DeleteAimPersonalTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplateMaterial 删除模板素材
//
// 根据用户提供的模板ID，删除模板素材。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) DeleteTemplateMaterial(request *model.DeleteTemplateMaterialRequest) (*model.DeleteTemplateMaterialResponse, error) {
	requestDef := GenReqDefForDeleteTemplateMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateMaterialResponse), nil
	}
}

// DeleteTemplateMaterialInvoker 删除模板素材
func (c *KooMessageClient) DeleteTemplateMaterialInvoker(request *model.DeleteTemplateMaterialRequest) *DeleteTemplateMaterialInvoker {
	requestDef := GenReqDefForDeleteTemplateMaterial()
	return &DeleteTemplateMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimTemplateMaterials 查询智能消息模板素材列表
//
// 根据用户提供的过滤条件，查询模板素材列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimTemplateMaterials(request *model.ListAimTemplateMaterialsRequest) (*model.ListAimTemplateMaterialsResponse, error) {
	requestDef := GenReqDefForListAimTemplateMaterials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimTemplateMaterialsResponse), nil
	}
}

// ListAimTemplateMaterialsInvoker 查询智能消息模板素材列表
func (c *KooMessageClient) ListAimTemplateMaterialsInvoker(request *model.ListAimTemplateMaterialsRequest) *ListAimTemplateMaterialsInvoker {
	requestDef := GenReqDefForListAimTemplateMaterials()
	return &ListAimTemplateMaterialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimTemplateReports 查询模板报表
//
// 根据用户指定过滤条件查询指定智能信息模板的解析次数。当日数据需要次日16:00之后才能查询到。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimTemplateReports(request *model.ListAimTemplateReportsRequest) (*model.ListAimTemplateReportsResponse, error) {
	requestDef := GenReqDefForListAimTemplateReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimTemplateReportsResponse), nil
	}
}

// ListAimTemplateReportsInvoker 查询模板报表
func (c *KooMessageClient) ListAimTemplateReportsInvoker(request *model.ListAimTemplateReportsRequest) *ListAimTemplateReportsInvoker {
	requestDef := GenReqDefForListAimTemplateReports()
	return &ListAimTemplateReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimTemplates 查询模板
//
// 根据客户提供的过滤条件查询智能信息模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimTemplates(request *model.ListAimTemplatesRequest) (*model.ListAimTemplatesResponse, error) {
	requestDef := GenReqDefForListAimTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimTemplatesResponse), nil
	}
}

// ListAimTemplatesInvoker 查询模板
func (c *KooMessageClient) ListAimTemplatesInvoker(request *model.ListAimTemplatesRequest) *ListAimTemplatesInvoker {
	requestDef := GenReqDefForListAimTemplates()
	return &ListAimTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetPrimaryVideoThumbnail 设置视频模板封面图
//
// 根据用户提供的视频封面图资源ID和AIM视频资源ID设置视频模板的封面图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) SetPrimaryVideoThumbnail(request *model.SetPrimaryVideoThumbnailRequest) (*model.SetPrimaryVideoThumbnailResponse, error) {
	requestDef := GenReqDefForSetPrimaryVideoThumbnail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetPrimaryVideoThumbnailResponse), nil
	}
}

// SetPrimaryVideoThumbnailInvoker 设置视频模板封面图
func (c *KooMessageClient) SetPrimaryVideoThumbnailInvoker(request *model.SetPrimaryVideoThumbnailRequest) *SetPrimaryVideoThumbnailInvoker {
	requestDef := GenReqDefForSetPrimaryVideoThumbnail()
	return &SetPrimaryVideoThumbnailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateVideoThumbnail 查询视频模板封面图
//
// 根据用户提供的过滤条件，查询视频模板封面图资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ShowTemplateVideoThumbnail(request *model.ShowTemplateVideoThumbnailRequest) (*model.ShowTemplateVideoThumbnailResponse, error) {
	requestDef := GenReqDefForShowTemplateVideoThumbnail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateVideoThumbnailResponse), nil
	}
}

// ShowTemplateVideoThumbnailInvoker 查询视频模板封面图
func (c *KooMessageClient) ShowTemplateVideoThumbnailInvoker(request *model.ShowTemplateVideoThumbnailRequest) *ShowTemplateVideoThumbnailInvoker {
	requestDef := GenReqDefForShowTemplateVideoThumbnail()
	return &ShowTemplateVideoThumbnailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePersonalTemplateState 启用或禁用模板实例
//
// 根据用户提供的模板ID，启用或禁用智能信息个人模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdatePersonalTemplateState(request *model.UpdatePersonalTemplateStateRequest) (*model.UpdatePersonalTemplateStateResponse, error) {
	requestDef := GenReqDefForUpdatePersonalTemplateState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePersonalTemplateStateResponse), nil
	}
}

// UpdatePersonalTemplateStateInvoker 启用或禁用模板实例
func (c *KooMessageClient) UpdatePersonalTemplateStateInvoker(request *model.UpdatePersonalTemplateStateRequest) *UpdatePersonalTemplateStateInvoker {
	requestDef := GenReqDefForUpdatePersonalTemplateState()
	return &UpdatePersonalTemplateStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAimTemplateMaterial 上传智能信息模板素材
//
// 支持用户上传模板使用的图片或者视频。
//
// &gt; 单个租户资源空间为10GB，包括回收站内容，请及时清理无用资源，防止资源浪费。
// &gt; 请求中所有字符串不允许携带“&lt;”、“&gt;”或多个空格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UploadAimTemplateMaterial(request *model.UploadAimTemplateMaterialRequest) (*model.UploadAimTemplateMaterialResponse, error) {
	requestDef := GenReqDefForUploadAimTemplateMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAimTemplateMaterialResponse), nil
	}
}

// UploadAimTemplateMaterialInvoker 上传智能信息模板素材
func (c *KooMessageClient) UploadAimTemplateMaterialInvoker(request *model.UploadAimTemplateMaterialRequest) *UploadAimTemplateMaterialInvoker {
	requestDef := GenReqDefForUploadAimTemplateMaterial()
	return &UploadAimTemplateMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMenus 查询智能信息服务号菜单
//
// 根据用户提供的过滤条件查询服务号菜单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListMenus(request *model.ListMenusRequest) (*model.ListMenusResponse, error) {
	requestDef := GenReqDefForListMenus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMenusResponse), nil
	}
}

// ListMenusInvoker 查询智能信息服务号菜单
func (c *KooMessageClient) ListMenusInvoker(request *model.ListMenusRequest) *ListMenusInvoker {
	requestDef := GenReqDefForListMenus()
	return &ListMenusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMenu 修改智能信息服务号菜单
//
// 支持用户修改所属企业的指定菜单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdateMenu(request *model.UpdateMenuRequest) (*model.UpdateMenuResponse, error) {
	requestDef := GenReqDefForUpdateMenu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMenuResponse), nil
	}
}

// UpdateMenuInvoker 修改智能信息服务号菜单
func (c *KooMessageClient) UpdateMenuInvoker(request *model.UpdateMenuRequest) *UpdateMenuInvoker {
	requestDef := GenReqDefForUpdateMenu()
	return &UpdateMenuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePortInfo 删除通道号
//
// 删除通道号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) DeletePortInfo(request *model.DeletePortInfoRequest) (*model.DeletePortInfoResponse, error) {
	requestDef := GenReqDefForDeletePortInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePortInfoResponse), nil
	}
}

// DeletePortInfoInvoker 删除通道号
func (c *KooMessageClient) DeletePortInfoInvoker(request *model.DeletePortInfoRequest) *DeletePortInfoInvoker {
	requestDef := GenReqDefForDeletePortInfo()
	return &DeletePortInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortInfos 查询通道号列表
//
// 支持查询通道号列表和通道号绑定信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListPortInfos(request *model.ListPortInfosRequest) (*model.ListPortInfosResponse, error) {
	requestDef := GenReqDefForListPortInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortInfosResponse), nil
	}
}

// ListPortInfosInvoker 查询通道号列表
func (c *KooMessageClient) ListPortInfosInvoker(request *model.ListPortInfosRequest) *ListPortInfosInvoker {
	requestDef := GenReqDefForListPortInfos()
	return &ListPortInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockPort 通道号绑定服务号
//
// 通道号绑定服务号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) LockPort(request *model.LockPortRequest) (*model.LockPortResponse, error) {
	requestDef := GenReqDefForLockPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockPortResponse), nil
	}
}

// LockPortInvoker 通道号绑定服务号
func (c *KooMessageClient) LockPortInvoker(request *model.LockPortRequest) *LockPortInvoker {
	requestDef := GenReqDefForLockPort()
	return &LockPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterPort 注册通道号
//
// 注册通道号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) RegisterPort(request *model.RegisterPortRequest) (*model.RegisterPortResponse, error) {
	requestDef := GenReqDefForRegisterPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterPortResponse), nil
	}
}

// RegisterPortInvoker 注册通道号
func (c *KooMessageClient) RegisterPortInvoker(request *model.RegisterPortRequest) *RegisterPortInvoker {
	requestDef := GenReqDefForRegisterPort()
	return &RegisterPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockPort 通道号解绑服务号
//
// 通道号解绑服务号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UnlockPort(request *model.UnlockPortRequest) (*model.UnlockPortResponse, error) {
	requestDef := GenReqDefForUnlockPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockPortResponse), nil
	}
}

// UnlockPortInvoker 通道号解绑服务号
func (c *KooMessageClient) UnlockPortInvoker(request *model.UnlockPortRequest) *UnlockPortInvoker {
	requestDef := GenReqDefForUnlockPort()
	return &UnlockPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPortalInfos 查询主页列表
//
// 根据用户提供的过滤条件查找用户管理的主页列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListPortalInfos(request *model.ListPortalInfosRequest) (*model.ListPortalInfosResponse, error) {
	requestDef := GenReqDefForListPortalInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortalInfosResponse), nil
	}
}

// ListPortalInfosInvoker 查询主页列表
func (c *KooMessageClient) ListPortalInfosInvoker(request *model.ListPortalInfosRequest) *ListPortalInfosInvoker {
	requestDef := GenReqDefForListPortalInfos()
	return &ListPortalInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePortalInfo 修改主页信息
//
// 用户对已创建的主页进行信息的修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdatePortalInfo(request *model.UpdatePortalInfoRequest) (*model.UpdatePortalInfoResponse, error) {
	requestDef := GenReqDefForUpdatePortalInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortalInfoResponse), nil
	}
}

// UpdatePortalInfoInvoker 修改主页信息
func (c *KooMessageClient) UpdatePortalInfoInvoker(request *model.UpdatePortalInfoRequest) *UpdatePortalInfoInvoker {
	requestDef := GenReqDefForUpdatePortalInfo()
	return &UpdatePortalInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// FreezePub 冻结服务号
//
// 支持用户通过此接口冻结服务号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) FreezePub(request *model.FreezePubRequest) (*model.FreezePubResponse, error) {
	requestDef := GenReqDefForFreezePub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezePubResponse), nil
	}
}

// FreezePubInvoker 冻结服务号
func (c *KooMessageClient) FreezePubInvoker(request *model.FreezePubRequest) *FreezePubInvoker {
	requestDef := GenReqDefForFreezePub()
	return &FreezePubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPubInfos 查询服务号列表
//
// 支持根据用户提供的过滤条件查询服务号列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListPubInfos(request *model.ListPubInfosRequest) (*model.ListPubInfosResponse, error) {
	requestDef := GenReqDefForListPubInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPubInfosResponse), nil
	}
}

// ListPubInfosInvoker 查询服务号列表
func (c *KooMessageClient) ListPubInfosInvoker(request *model.ListPubInfosRequest) *ListPubInfosInvoker {
	requestDef := GenReqDefForListPubInfos()
	return &ListPubInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnfreezePub 解冻服务号
//
// 服务号解结，冻结服务号。需审核，审核通过生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UnfreezePub(request *model.UnfreezePubRequest) (*model.UnfreezePubResponse, error) {
	requestDef := GenReqDefForUnfreezePub()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezePubResponse), nil
	}
}

// UnfreezePubInvoker 解冻服务号
func (c *KooMessageClient) UnfreezePubInvoker(request *model.UnfreezePubRequest) *UnfreezePubInvoker {
	requestDef := GenReqDefForUnfreezePub()
	return &UnfreezePubInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePubInfo 更新服务号信息
//
// 支持用户更新服务号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdatePubInfo(request *model.UpdatePubInfoRequest) (*model.UpdatePubInfoResponse, error) {
	requestDef := GenReqDefForUpdatePubInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePubInfoResponse), nil
	}
}

// UpdatePubInfoInvoker 更新服务号信息
func (c *KooMessageClient) UpdatePubInfoInvoker(request *model.UpdatePubInfoRequest) *UpdatePubInfoInvoker {
	requestDef := GenReqDefForUpdatePubInfo()
	return &UpdatePubInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePubInfo 一站式服务号创建
//
// 一站式服务号创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreatePubInfo(request *model.CreatePubInfoRequest) (*model.CreatePubInfoResponse, error) {
	requestDef := GenReqDefForCreatePubInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePubInfoResponse), nil
	}
}

// CreatePubInfoInvoker 一站式服务号创建
func (c *KooMessageClient) CreatePubInfoInvoker(request *model.CreatePubInfoRequest) *CreatePubInfoInvoker {
	requestDef := GenReqDefForCreatePubInfo()
	return &CreatePubInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushMenuInfo 催审菜单
//
// 支持用户通过此接口根据菜单ID催审。菜单需要在与其关联的服务号审核通过之后才能催审。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) PushMenuInfo(request *model.PushMenuInfoRequest) (*model.PushMenuInfoResponse, error) {
	requestDef := GenReqDefForPushMenuInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushMenuInfoResponse), nil
	}
}

// PushMenuInfoInvoker 催审菜单
func (c *KooMessageClient) PushMenuInfoInvoker(request *model.PushMenuInfoRequest) *PushMenuInfoInvoker {
	requestDef := GenReqDefForPushMenuInfo()
	return &PushMenuInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushPortalInfo 催审主页
//
// 支持用户通过此接口根据主页ID催审。主页需要在与其关联的服务号审核通过之后才能催审。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) PushPortalInfo(request *model.PushPortalInfoRequest) (*model.PushPortalInfoResponse, error) {
	requestDef := GenReqDefForPushPortalInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushPortalInfoResponse), nil
	}
}

// PushPortalInfoInvoker 催审主页
func (c *KooMessageClient) PushPortalInfoInvoker(request *model.PushPortalInfoRequest) *PushPortalInfoInvoker {
	requestDef := GenReqDefForPushPortalInfo()
	return &PushPortalInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadMedia 上传智能信息服务号图片资源
//
// 支持用户上传图片资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UploadMedia(request *model.UploadMediaRequest) (*model.UploadMediaResponse, error) {
	requestDef := GenReqDefForUploadMedia()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadMediaResponse), nil
	}
}

// UploadMediaInvoker 上传智能信息服务号图片资源
func (c *KooMessageClient) UploadMediaInvoker(request *model.UploadMediaRequest) *UploadMediaInvoker {
	requestDef := GenReqDefForUploadMedia()
	return &UploadMediaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSmsApp 创建短信应用
//
// 该接口用于用户创建短信应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateSmsApp(request *model.CreateSmsAppRequest) (*model.CreateSmsAppResponse, error) {
	requestDef := GenReqDefForCreateSmsApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSmsAppResponse), nil
	}
}

// CreateSmsAppInvoker 创建短信应用
func (c *KooMessageClient) CreateSmsAppInvoker(request *model.CreateSmsAppRequest) *CreateSmsAppInvoker {
	requestDef := GenReqDefForCreateSmsApp()
	return &CreateSmsAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimMsgApp 查询短信应用
//
// 该接口用于用户查询已创建的短信应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimMsgApp(request *model.ListAimMsgAppRequest) (*model.ListAimMsgAppResponse, error) {
	requestDef := GenReqDefForListAimMsgApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimMsgAppResponse), nil
	}
}

// ListAimMsgAppInvoker 查询短信应用
func (c *KooMessageClient) ListAimMsgAppInvoker(request *model.ListAimMsgAppRequest) *ListAimMsgAppInvoker {
	requestDef := GenReqDefForListAimMsgApp()
	return &ListAimMsgAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimMsgAppDetail 获取短信应用详情
//
// 该接口用于用户获取已创建的短信应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimMsgAppDetail(request *model.ListAimMsgAppDetailRequest) (*model.ListAimMsgAppDetailResponse, error) {
	requestDef := GenReqDefForListAimMsgAppDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimMsgAppDetailResponse), nil
	}
}

// ListAimMsgAppDetailInvoker 获取短信应用详情
func (c *KooMessageClient) ListAimMsgAppDetailInvoker(request *model.ListAimMsgAppDetailRequest) *ListAimMsgAppDetailInvoker {
	requestDef := GenReqDefForListAimMsgAppDetail()
	return &ListAimMsgAppDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAimMsgApp 修改短信应用
//
// 该接口用于用户修改短信应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdateAimMsgApp(request *model.UpdateAimMsgAppRequest) (*model.UpdateAimMsgAppResponse, error) {
	requestDef := GenReqDefForUpdateAimMsgApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAimMsgAppResponse), nil
	}
}

// UpdateAimMsgAppInvoker 修改短信应用
func (c *KooMessageClient) UpdateAimMsgAppInvoker(request *model.UpdateAimMsgAppRequest) *UpdateAimMsgAppInvoker {
	requestDef := GenReqDefForUpdateAimMsgApp()
	return &UpdateAimMsgAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendAimBatchDifferentMessages 发送分批短信
//
// 该接口用于向不同用户发送不同内容的短信。
//
// - 前提条件
//
// 1. 已创建短信应用。
// 2. 已申请短信签名，获取签名通道号。
// 3. 已申请短信模板，获取模板ID。
//
// - 注意事项
//
// 单条短信最多允许携带500个号码，每个号码最大长度为21位。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) SendAimBatchDifferentMessages(request *model.SendAimBatchDifferentMessagesRequest) (*model.SendAimBatchDifferentMessagesResponse, error) {
	requestDef := GenReqDefForSendAimBatchDifferentMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendAimBatchDifferentMessagesResponse), nil
	}
}

// SendAimBatchDifferentMessagesInvoker 发送分批短信
func (c *KooMessageClient) SendAimBatchDifferentMessagesInvoker(request *model.SendAimBatchDifferentMessagesRequest) *SendAimBatchDifferentMessagesInvoker {
	requestDef := GenReqDefForSendAimBatchDifferentMessages()
	return &SendAimBatchDifferentMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendAimBatchMessages 发送短信
//
// 向单个或多个用户发送相同内容的短信。
//
// - 前提条件
//
// 1. 已创建短信应用。
// 2. 已申请短信签名，获取签名通道号。
// 3. 已申请短信模板，获取模板ID。
//
// - 注意事项
//
// 最多允许携带500个号码，每个号码最大长度为21位。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) SendAimBatchMessages(request *model.SendAimBatchMessagesRequest) (*model.SendAimBatchMessagesResponse, error) {
	requestDef := GenReqDefForSendAimBatchMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendAimBatchMessagesResponse), nil
	}
}

// SendAimBatchMessagesInvoker 发送短信
func (c *KooMessageClient) SendAimBatchMessagesInvoker(request *model.SendAimBatchMessagesRequest) *SendAimBatchMessagesInvoker {
	requestDef := GenReqDefForSendAimBatchMessages()
	return &SendAimBatchMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAimMsgSignature 创建短信签名
//
// 该接口用于用户创建短信签名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) AddAimMsgSignature(request *model.AddAimMsgSignatureRequest) (*model.AddAimMsgSignatureResponse, error) {
	requestDef := GenReqDefForAddAimMsgSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAimMsgSignatureResponse), nil
	}
}

// AddAimMsgSignatureInvoker 创建短信签名
func (c *KooMessageClient) AddAimMsgSignatureInvoker(request *model.AddAimMsgSignatureRequest) *AddAimMsgSignatureInvoker {
	requestDef := GenReqDefForAddAimMsgSignature()
	return &AddAimMsgSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAimMsgSignature 删除短信签名
//
// 该接口用于用户删除已创建的短信签名。删除已审核通过的签名，会同步删除该签名对应的通道号和该签名下的所有短信模板，请谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) DeleteAimMsgSignature(request *model.DeleteAimMsgSignatureRequest) (*model.DeleteAimMsgSignatureResponse, error) {
	requestDef := GenReqDefForDeleteAimMsgSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAimMsgSignatureResponse), nil
	}
}

// DeleteAimMsgSignatureInvoker 删除短信签名
func (c *KooMessageClient) DeleteAimMsgSignatureInvoker(request *model.DeleteAimMsgSignatureRequest) *DeleteAimMsgSignatureInvoker {
	requestDef := GenReqDefForDeleteAimMsgSignature()
	return &DeleteAimMsgSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimMsgSignature 查询短信签名
//
// 该接口用于用户查询已创建的短信签名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimMsgSignature(request *model.ListAimMsgSignatureRequest) (*model.ListAimMsgSignatureResponse, error) {
	requestDef := GenReqDefForListAimMsgSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimMsgSignatureResponse), nil
	}
}

// ListAimMsgSignatureInvoker 查询短信签名
func (c *KooMessageClient) ListAimMsgSignatureInvoker(request *model.ListAimMsgSignatureRequest) *ListAimMsgSignatureInvoker {
	requestDef := GenReqDefForListAimMsgSignature()
	return &ListAimMsgSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimMsgSignatureDetail 获取短信签名详情
//
// 该接口用于用户获取已创建的短信签名详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimMsgSignatureDetail(request *model.ListAimMsgSignatureDetailRequest) (*model.ListAimMsgSignatureDetailResponse, error) {
	requestDef := GenReqDefForListAimMsgSignatureDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimMsgSignatureDetailResponse), nil
	}
}

// ListAimMsgSignatureDetailInvoker 获取短信签名详情
func (c *KooMessageClient) ListAimMsgSignatureDetailInvoker(request *model.ListAimMsgSignatureDetailRequest) *ListAimMsgSignatureDetailInvoker {
	requestDef := GenReqDefForListAimMsgSignatureDetail()
	return &ListAimMsgSignatureDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAimMsgSignatureFileInfo 查询申请文件
//
// 该接口用于用户查询创建短信签名时上传的营业执照/授权委托书文件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ShowAimMsgSignatureFileInfo(request *model.ShowAimMsgSignatureFileInfoRequest) (*model.ShowAimMsgSignatureFileInfoResponse, error) {
	requestDef := GenReqDefForShowAimMsgSignatureFileInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAimMsgSignatureFileInfoResponse), nil
	}
}

// ShowAimMsgSignatureFileInfoInvoker 查询申请文件
func (c *KooMessageClient) ShowAimMsgSignatureFileInfoInvoker(request *model.ShowAimMsgSignatureFileInfoRequest) *ShowAimMsgSignatureFileInfoInvoker {
	requestDef := GenReqDefForShowAimMsgSignatureFileInfo()
	return &ShowAimMsgSignatureFileInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAimMsgSignature 修改短信签名
//
// 该接口用于用户更新短信签名信息，目前仅支持审核不通过的短信签名重新修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdateAimMsgSignature(request *model.UpdateAimMsgSignatureRequest) (*model.UpdateAimMsgSignatureResponse, error) {
	requestDef := GenReqDefForUpdateAimMsgSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAimMsgSignatureResponse), nil
	}
}

// UpdateAimMsgSignatureInvoker 修改短信签名
func (c *KooMessageClient) UpdateAimMsgSignatureInvoker(request *model.UpdateAimMsgSignatureRequest) *UpdateAimMsgSignatureInvoker {
	requestDef := GenReqDefForUpdateAimMsgSignature()
	return &UpdateAimMsgSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAimMsgSignatureFile 上传申请文件
//
// 该接口用于用户上传创建短信签名所需的营业执照/授权委托书文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UploadAimMsgSignatureFile(request *model.UploadAimMsgSignatureFileRequest) (*model.UploadAimMsgSignatureFileResponse, error) {
	requestDef := GenReqDefForUploadAimMsgSignatureFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAimMsgSignatureFileResponse), nil
	}
}

// UploadAimMsgSignatureFileInvoker 上传申请文件
func (c *KooMessageClient) UploadAimMsgSignatureFileInvoker(request *model.UploadAimMsgSignatureFileRequest) *UploadAimMsgSignatureFileInvoker {
	requestDef := GenReqDefForUploadAimMsgSignatureFile()
	return &UploadAimMsgSignatureFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAimMsgTemplate 创建短信模板
//
// 该接口用于用户创建短信模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateAimMsgTemplate(request *model.CreateAimMsgTemplateRequest) (*model.CreateAimMsgTemplateResponse, error) {
	requestDef := GenReqDefForCreateAimMsgTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAimMsgTemplateResponse), nil
	}
}

// CreateAimMsgTemplateInvoker 创建短信模板
func (c *KooMessageClient) CreateAimMsgTemplateInvoker(request *model.CreateAimMsgTemplateRequest) *CreateAimMsgTemplateInvoker {
	requestDef := GenReqDefForCreateAimMsgTemplate()
	return &CreateAimMsgTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAimMsgTemplate 删除短信模板
//
// 该接口用于用户删除已创建的短信模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) DeleteAimMsgTemplate(request *model.DeleteAimMsgTemplateRequest) (*model.DeleteAimMsgTemplateResponse, error) {
	requestDef := GenReqDefForDeleteAimMsgTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAimMsgTemplateResponse), nil
	}
}

// DeleteAimMsgTemplateInvoker 删除短信模板
func (c *KooMessageClient) DeleteAimMsgTemplateInvoker(request *model.DeleteAimMsgTemplateRequest) *DeleteAimMsgTemplateInvoker {
	requestDef := GenReqDefForDeleteAimMsgTemplate()
	return &DeleteAimMsgTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAimMsgTemplate 查询短信模板
//
// 该接口用于用户查询已创建的短信模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListAimMsgTemplate(request *model.ListAimMsgTemplateRequest) (*model.ListAimMsgTemplateResponse, error) {
	requestDef := GenReqDefForListAimMsgTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAimMsgTemplateResponse), nil
	}
}

// ListAimMsgTemplateInvoker 查询短信模板
func (c *KooMessageClient) ListAimMsgTemplateInvoker(request *model.ListAimMsgTemplateRequest) *ListAimMsgTemplateInvoker {
	requestDef := GenReqDefForListAimMsgTemplate()
	return &ListAimMsgTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAimMsgTemplateDetail 获取短信模板详情
//
// 该接口用于用户获取已创建的短信模板详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ShowAimMsgTemplateDetail(request *model.ShowAimMsgTemplateDetailRequest) (*model.ShowAimMsgTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowAimMsgTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAimMsgTemplateDetailResponse), nil
	}
}

// ShowAimMsgTemplateDetailInvoker 获取短信模板详情
func (c *KooMessageClient) ShowAimMsgTemplateDetailInvoker(request *model.ShowAimMsgTemplateDetailRequest) *ShowAimMsgTemplateDetailInvoker {
	requestDef := GenReqDefForShowAimMsgTemplateDetail()
	return &ShowAimMsgTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAimMsgTemplateVariable 查询短信模板变量
//
// 该接口用于用户查询短信模板变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ShowAimMsgTemplateVariable(request *model.ShowAimMsgTemplateVariableRequest) (*model.ShowAimMsgTemplateVariableResponse, error) {
	requestDef := GenReqDefForShowAimMsgTemplateVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAimMsgTemplateVariableResponse), nil
	}
}

// ShowAimMsgTemplateVariableInvoker 查询短信模板变量
func (c *KooMessageClient) ShowAimMsgTemplateVariableInvoker(request *model.ShowAimMsgTemplateVariableRequest) *ShowAimMsgTemplateVariableInvoker {
	requestDef := GenReqDefForShowAimMsgTemplateVariable()
	return &ShowAimMsgTemplateVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAimMsgTemplate 修改短信模板
//
// 该接口用于用户修改短信模板信息，目前仅支持审核不通过的短信模板重新修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) UpdateAimMsgTemplate(request *model.UpdateAimMsgTemplateRequest) (*model.UpdateAimMsgTemplateResponse, error) {
	requestDef := GenReqDefForUpdateAimMsgTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAimMsgTemplateResponse), nil
	}
}

// UpdateAimMsgTemplateInvoker 修改短信模板
func (c *KooMessageClient) UpdateAimMsgTemplateInvoker(request *model.UpdateAimMsgTemplateRequest) *UpdateAimMsgTemplateInvoker {
	requestDef := GenReqDefForUpdateAimMsgTemplate()
	return &UpdateAimMsgTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVmsCallBack 注册智能信息基础版回执URL
//
// 用户根据要求创建智能信息基础版回执接口后，可以调用此接口进行注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) AddVmsCallBack(request *model.AddVmsCallBackRequest) (*model.AddVmsCallBackResponse, error) {
	requestDef := GenReqDefForAddVmsCallBack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVmsCallBackResponse), nil
	}
}

// AddVmsCallBackInvoker 注册智能信息基础版回执URL
func (c *KooMessageClient) AddVmsCallBackInvoker(request *model.AddVmsCallBackRequest) *AddVmsCallBackInvoker {
	requestDef := GenReqDefForAddVmsCallBack()
	return &AddVmsCallBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVmsSendTask 发送智能信息基础版任务
//
// 支持用户通过此接口进行智能信息基础版任务发送。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateVmsSendTask(request *model.CreateVmsSendTaskRequest) (*model.CreateVmsSendTaskResponse, error) {
	requestDef := GenReqDefForCreateVmsSendTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVmsSendTaskResponse), nil
	}
}

// CreateVmsSendTaskInvoker 发送智能信息基础版任务
func (c *KooMessageClient) CreateVmsSendTaskInvoker(request *model.CreateVmsSendTaskRequest) *CreateVmsSendTaskInvoker {
	requestDef := GenReqDefForCreateVmsSendTask()
	return &CreateVmsSendTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVmsCallbacks 查询用户已注册智能信息基础版回执接口
//
// 查询所有已注册的智能信息基础版回执接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListVmsCallbacks(request *model.ListVmsCallbacksRequest) (*model.ListVmsCallbacksResponse, error) {
	requestDef := GenReqDefForListVmsCallbacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVmsCallbacksResponse), nil
	}
}

// ListVmsCallbacksInvoker 查询用户已注册智能信息基础版回执接口
func (c *KooMessageClient) ListVmsCallbacksInvoker(request *model.ListVmsCallbacksRequest) *ListVmsCallbacksInvoker {
	requestDef := GenReqDefForListVmsCallbacks()
	return &ListVmsCallbacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVmsSendTasks 查询智能信息基础版发送任务
//
// 支持用户根据过滤条件查询智能信息基础版任务列表。包括：发送任务名称、智能信息基础版模板ID等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListVmsSendTasks(request *model.ListVmsSendTasksRequest) (*model.ListVmsSendTasksResponse, error) {
	requestDef := GenReqDefForListVmsSendTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVmsSendTasksResponse), nil
	}
}

// ListVmsSendTasksInvoker 查询智能信息基础版发送任务
func (c *KooMessageClient) ListVmsSendTasksInvoker(request *model.ListVmsSendTasksRequest) *ListVmsSendTasksInvoker {
	requestDef := GenReqDefForListVmsSendTasks()
	return &ListVmsSendTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVmsTemplate 新建智能信息基础版模板
//
// 支持用户通过此接口创建智能信息基础版模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) CreateVmsTemplate(request *model.CreateVmsTemplateRequest) (*model.CreateVmsTemplateResponse, error) {
	requestDef := GenReqDefForCreateVmsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVmsTemplateResponse), nil
	}
}

// CreateVmsTemplateInvoker 新建智能信息基础版模板
func (c *KooMessageClient) CreateVmsTemplateInvoker(request *model.CreateVmsTemplateRequest) *CreateVmsTemplateInvoker {
	requestDef := GenReqDefForCreateVmsTemplate()
	return &CreateVmsTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVmsTemplateStatus 查询智能信息基础版模板状态
//
// 根据用户提供的过滤条件查询智能信息基础版模板状态列表。
// 包括：模板ID、模板名称等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KooMessageClient) ListVmsTemplateStatus(request *model.ListVmsTemplateStatusRequest) (*model.ListVmsTemplateStatusResponse, error) {
	requestDef := GenReqDefForListVmsTemplateStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVmsTemplateStatusResponse), nil
	}
}

// ListVmsTemplateStatusInvoker 查询智能信息基础版模板状态
func (c *KooMessageClient) ListVmsTemplateStatusInvoker(request *model.ListVmsTemplateStatusRequest) *ListVmsTemplateStatusInvoker {
	requestDef := GenReqDefForListVmsTemplateStatus()
	return &ListVmsTemplateStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
