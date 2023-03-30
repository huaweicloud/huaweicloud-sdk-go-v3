package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
)

type CloudIDEClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudIDEClient(hcClient *http_client.HcHttpClient) *CloudIDEClient {
	return &CloudIDEClient{HcClient: hcClient}
}

func CloudIDEClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddExtensionEvaluation 添加插件评论
//
// 添加插件评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) AddExtensionEvaluation(request *model.AddExtensionEvaluationRequest) (*model.AddExtensionEvaluationResponse, error) {
	requestDef := GenReqDefForAddExtensionEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddExtensionEvaluationResponse), nil
	}
}

// AddExtensionEvaluationInvoker 添加插件评论
func (c *CloudIDEClient) AddExtensionEvaluationInvoker(request *model.AddExtensionEvaluationRequest) *AddExtensionEvaluationInvoker {
	requestDef := GenReqDefForAddExtensionEvaluation()
	return &AddExtensionEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddExtensionEvaluationReply 添加评论回复、回复评论回复
//
// 添加评论回复、回复评论回复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) AddExtensionEvaluationReply(request *model.AddExtensionEvaluationReplyRequest) (*model.AddExtensionEvaluationReplyResponse, error) {
	requestDef := GenReqDefForAddExtensionEvaluationReply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddExtensionEvaluationReplyResponse), nil
	}
}

// AddExtensionEvaluationReplyInvoker 添加评论回复、回复评论回复
func (c *CloudIDEClient) AddExtensionEvaluationReplyInvoker(request *model.AddExtensionEvaluationReplyRequest) *AddExtensionEvaluationReplyInvoker {
	requestDef := GenReqDefForAddExtensionEvaluationReply()
	return &AddExtensionEvaluationReplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddExtensionStar 添加新评星
//
// 添加新评星
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) AddExtensionStar(request *model.AddExtensionStarRequest) (*model.AddExtensionStarResponse, error) {
	requestDef := GenReqDefForAddExtensionStar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddExtensionStarResponse), nil
	}
}

// AddExtensionStarInvoker 添加新评星
func (c *CloudIDEClient) AddExtensionStarInvoker(request *model.AddExtensionStarRequest) *AddExtensionStarInvoker {
	requestDef := GenReqDefForAddExtensionStar()
	return &AddExtensionStarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckMaliciousExtensionEvaluation 举报评论,举报回复
//
// 举报评论,举报回复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CheckMaliciousExtensionEvaluation(request *model.CheckMaliciousExtensionEvaluationRequest) (*model.CheckMaliciousExtensionEvaluationResponse, error) {
	requestDef := GenReqDefForCheckMaliciousExtensionEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckMaliciousExtensionEvaluationResponse), nil
	}
}

// CheckMaliciousExtensionEvaluationInvoker 举报评论,举报回复
func (c *CloudIDEClient) CheckMaliciousExtensionEvaluationInvoker(request *model.CheckMaliciousExtensionEvaluationRequest) *CheckMaliciousExtensionEvaluationInvoker {
	requestDef := GenReqDefForCheckMaliciousExtensionEvaluation()
	return &CheckMaliciousExtensionEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExtensionAuthorization 设置ide实例对插件的授权
//
// 设置ide实例对插件的授权。同意、不同意、未知（下次重新询问）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateExtensionAuthorization(request *model.CreateExtensionAuthorizationRequest) (*model.CreateExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExtensionAuthorizationResponse), nil
	}
}

// CreateExtensionAuthorizationInvoker 设置ide实例对插件的授权
func (c *CloudIDEClient) CreateExtensionAuthorizationInvoker(request *model.CreateExtensionAuthorizationRequest) *CreateExtensionAuthorizationInvoker {
	requestDef := GenReqDefForCreateExtensionAuthorization()
	return &CreateExtensionAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvaluation 删除评论
//
// 删除评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) DeleteEvaluation(request *model.DeleteEvaluationRequest) (*model.DeleteEvaluationResponse, error) {
	requestDef := GenReqDefForDeleteEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEvaluationResponse), nil
	}
}

// DeleteEvaluationInvoker 删除评论
func (c *CloudIDEClient) DeleteEvaluationInvoker(request *model.DeleteEvaluationRequest) *DeleteEvaluationInvoker {
	requestDef := GenReqDefForDeleteEvaluation()
	return &DeleteEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvaluationReply 删除回复
//
// 删除回复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) DeleteEvaluationReply(request *model.DeleteEvaluationReplyRequest) (*model.DeleteEvaluationReplyResponse, error) {
	requestDef := GenReqDefForDeleteEvaluationReply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEvaluationReplyResponse), nil
	}
}

// DeleteEvaluationReplyInvoker 删除回复
func (c *CloudIDEClient) DeleteEvaluationReplyInvoker(request *model.DeleteEvaluationReplyRequest) *DeleteEvaluationReplyInvoker {
	requestDef := GenReqDefForDeleteEvaluationReply()
	return &DeleteEvaluationReplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExtensions 查询插件列表
//
// 查询插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListExtensions(request *model.ListExtensionsRequest) (*model.ListExtensionsResponse, error) {
	requestDef := GenReqDefForListExtensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExtensionsResponse), nil
	}
}

// ListExtensionsInvoker 查询插件列表
func (c *CloudIDEClient) ListExtensionsInvoker(request *model.ListExtensionsRequest) *ListExtensionsInvoker {
	requestDef := GenReqDefForListExtensions()
	return &ListExtensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTemplates 查询技术栈模板工程
//
// 查询技术栈模板工程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListProjectTemplates(request *model.ListProjectTemplatesRequest) (*model.ListProjectTemplatesResponse, error) {
	requestDef := GenReqDefForListProjectTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTemplatesResponse), nil
	}
}

// ListProjectTemplatesInvoker 查询技术栈模板工程
func (c *CloudIDEClient) ListProjectTemplatesInvoker(request *model.ListProjectTemplatesRequest) *ListProjectTemplatesInvoker {
	requestDef := GenReqDefForListProjectTemplates()
	return &ListProjectTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublisher 获取当前用户下的发布商列表
//
// 获取当前用户下的发布商列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListPublisher(request *model.ListPublisherRequest) (*model.ListPublisherResponse, error) {
	requestDef := GenReqDefForListPublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublisherResponse), nil
	}
}

// ListPublisherInvoker 获取当前用户下的发布商列表
func (c *CloudIDEClient) ListPublisherInvoker(request *model.ListPublisherRequest) *ListPublisherInvoker {
	requestDef := GenReqDefForListPublisher()
	return &ListPublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStacks 按region获取标签所有技术栈
//
// 按region获取标签所有技术栈
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListStacks(request *model.ListStacksRequest) (*model.ListStacksResponse, error) {
	requestDef := GenReqDefForListStacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStacksResponse), nil
	}
}

// ListStacksInvoker 按region获取标签所有技术栈
func (c *CloudIDEClient) ListStacksInvoker(request *model.ListStacksRequest) *ListStacksInvoker {
	requestDef := GenReqDefForListStacks()
	return &ListStacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishExtension 插件发布
//
// 插件发布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) PublishExtension(request *model.PublishExtensionRequest) (*model.PublishExtensionResponse, error) {
	requestDef := GenReqDefForPublishExtension()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishExtensionResponse), nil
	}
}

// PublishExtensionInvoker 插件发布
func (c *CloudIDEClient) PublishExtensionInvoker(request *model.PublishExtensionRequest) *PublishExtensionInvoker {
	requestDef := GenReqDefForPublishExtension()
	return &PublishExtensionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccountStatus 查询当前帐号访问权限
//
// 查询当前帐号访问权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowAccountStatus(request *model.ShowAccountStatusRequest) (*model.ShowAccountStatusResponse, error) {
	requestDef := GenReqDefForShowAccountStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccountStatusResponse), nil
	}
}

// ShowAccountStatusInvoker 查询当前帐号访问权限
func (c *CloudIDEClient) ShowAccountStatusInvoker(request *model.ShowAccountStatusRequest) *ShowAccountStatusInvoker {
	requestDef := GenReqDefForShowAccountStatus()
	return &ShowAccountStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCategoryList 查询插件分类
//
// 查询插件分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowCategoryList(request *model.ShowCategoryListRequest) (*model.ShowCategoryListResponse, error) {
	requestDef := GenReqDefForShowCategoryList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCategoryListResponse), nil
	}
}

// ShowCategoryListInvoker 查询插件分类
func (c *CloudIDEClient) ShowCategoryListInvoker(request *model.ShowCategoryListRequest) *ShowCategoryListInvoker {
	requestDef := GenReqDefForShowCategoryList()
	return &ShowCategoryListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionAuthorization 查询ide实例对插件的授权情况
//
// 查询ide实例对插件的授权情况，同意授权的插件能在ide实例内携带登陆用户的token调用第三方服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionAuthorization(request *model.ShowExtensionAuthorizationRequest) (*model.ShowExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForShowExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionAuthorizationResponse), nil
	}
}

// ShowExtensionAuthorizationInvoker 查询ide实例对插件的授权情况
func (c *CloudIDEClient) ShowExtensionAuthorizationInvoker(request *model.ShowExtensionAuthorizationRequest) *ShowExtensionAuthorizationInvoker {
	requestDef := GenReqDefForShowExtensionAuthorization()
	return &ShowExtensionAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionDetail 查询插件详细信息
//
// 查询插件详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionDetail(request *model.ShowExtensionDetailRequest) (*model.ShowExtensionDetailResponse, error) {
	requestDef := GenReqDefForShowExtensionDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionDetailResponse), nil
	}
}

// ShowExtensionDetailInvoker 查询插件详细信息
func (c *CloudIDEClient) ShowExtensionDetailInvoker(request *model.ShowExtensionDetailRequest) *ShowExtensionDetailInvoker {
	requestDef := GenReqDefForShowExtensionDetail()
	return &ShowExtensionDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionEvaluation 查询插件评价
//
// 查询插件评价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionEvaluation(request *model.ShowExtensionEvaluationRequest) (*model.ShowExtensionEvaluationResponse, error) {
	requestDef := GenReqDefForShowExtensionEvaluation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionEvaluationResponse), nil
	}
}

// ShowExtensionEvaluationInvoker 查询插件评价
func (c *CloudIDEClient) ShowExtensionEvaluationInvoker(request *model.ShowExtensionEvaluationRequest) *ShowExtensionEvaluationInvoker {
	requestDef := GenReqDefForShowExtensionEvaluation()
	return &ShowExtensionEvaluationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionEvaluationStar 查询插件评星
//
// 查询插件评星
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionEvaluationStar(request *model.ShowExtensionEvaluationStarRequest) (*model.ShowExtensionEvaluationStarResponse, error) {
	requestDef := GenReqDefForShowExtensionEvaluationStar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionEvaluationStarResponse), nil
	}
}

// ShowExtensionEvaluationStarInvoker 查询插件评星
func (c *CloudIDEClient) ShowExtensionEvaluationStarInvoker(request *model.ShowExtensionEvaluationStarRequest) *ShowExtensionEvaluationStarInvoker {
	requestDef := GenReqDefForShowExtensionEvaluationStar()
	return &ShowExtensionEvaluationStarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionTestingResult 获取插件检测结果
//
// 获取插件检测结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionTestingResult(request *model.ShowExtensionTestingResultRequest) (*model.ShowExtensionTestingResultResponse, error) {
	requestDef := GenReqDefForShowExtensionTestingResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionTestingResultResponse), nil
	}
}

// ShowExtensionTestingResultInvoker 获取插件检测结果
func (c *CloudIDEClient) ShowExtensionTestingResultInvoker(request *model.ShowExtensionTestingResultRequest) *ShowExtensionTestingResultInvoker {
	requestDef := GenReqDefForShowExtensionTestingResult()
	return &ShowExtensionTestingResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrice 获取技术栈计费信息
//
// 获取技术栈计费信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowPrice(request *model.ShowPriceRequest) (*model.ShowPriceResponse, error) {
	requestDef := GenReqDefForShowPrice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPriceResponse), nil
	}
}

// ShowPriceInvoker 获取技术栈计费信息
func (c *CloudIDEClient) ShowPriceInvoker(request *model.ShowPriceRequest) *ShowPriceInvoker {
	requestDef := GenReqDefForShowPrice()
	return &ShowPriceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadExtensionFile 上传插件
//
// 上传插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) UploadExtensionFile(request *model.UploadExtensionFileRequest) (*model.UploadExtensionFileResponse, error) {
	requestDef := GenReqDefForUploadExtensionFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadExtensionFileResponse), nil
	}
}

// UploadExtensionFileInvoker 上传插件
func (c *CloudIDEClient) UploadExtensionFileInvoker(request *model.UploadExtensionFileRequest) *UploadExtensionFileInvoker {
	requestDef := GenReqDefForUploadExtensionFile()
	return &UploadExtensionFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFilePublisher 文件上传归一化
//
// 文件上传归一化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) UploadFilePublisher(request *model.UploadFilePublisherRequest) (*model.UploadFilePublisherResponse, error) {
	requestDef := GenReqDefForUploadFilePublisher()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFilePublisherResponse), nil
	}
}

// UploadFilePublisherInvoker 文件上传归一化
func (c *CloudIDEClient) UploadFilePublisherInvoker(request *model.UploadFilePublisherRequest) *UploadFilePublisherInvoker {
	requestDef := GenReqDefForUploadFilePublisher()
	return &UploadFilePublisherInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAcceptance CreateAcceptance接口
//
// create a acceptance
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateAcceptance(request *model.CreateAcceptanceRequest) (*model.CreateAcceptanceResponse, error) {
	requestDef := GenReqDefForCreateAcceptance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAcceptanceResponse), nil
	}
}

// CreateAcceptanceInvoker CreateAcceptance接口
func (c *CloudIDEClient) CreateAcceptanceInvoker(request *model.CreateAcceptanceRequest) *CreateAcceptanceInvoker {
	requestDef := GenReqDefForCreateAcceptance()
	return &CreateAcceptanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApply CreateJoinRequest接口
//
// create a join-request
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateApply(request *model.CreateApplyRequest) (*model.CreateApplyResponse, error) {
	requestDef := GenReqDefForCreateApply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplyResponse), nil
	}
}

// CreateApplyInvoker CreateJoinRequest接口
func (c *CloudIDEClient) CreateApplyInvoker(request *model.CreateApplyRequest) *CreateApplyInvoker {
	requestDef := GenReqDefForCreateApply()
	return &CreateApplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvent CreateEvent接口
//
// create an event
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateEvent(request *model.CreateEventRequest) (*model.CreateEventResponse, error) {
	requestDef := GenReqDefForCreateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventResponse), nil
	}
}

// CreateEventInvoker CreateEvent接口
func (c *CloudIDEClient) CreateEventInvoker(request *model.CreateEventRequest) *CreateEventInvoker {
	requestDef := GenReqDefForCreateEvent()
	return &CreateEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogin CreateLogin接口
//
// create a login
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateLogin(request *model.CreateLoginRequest) (*model.CreateLoginResponse, error) {
	requestDef := GenReqDefForCreateLogin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoginResponse), nil
	}
}

// CreateLoginInvoker CreateLogin接口
func (c *CloudIDEClient) CreateLoginInvoker(request *model.CreateLoginRequest) *CreateLoginInvoker {
	requestDef := GenReqDefForCreateLogin()
	return &CreateLoginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRequest Create Request接口
//
// create a code generation request.
//
// if agent receives a code generation request, it will:
// - record the request into mysql,
// - decompose the request into &#x60;topn&#x60; tasks.
// - send the tasks to kafka.
//
// if agent receives a duplicated code generation request, it will not decompose request and send to kafka.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateRequest(request *model.CreateRequestRequest) (*model.CreateRequestResponse, error) {
	requestDef := GenReqDefForCreateRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRequestResponse), nil
	}
}

// CreateRequestInvoker Create Request接口
func (c *CloudIDEClient) CreateRequestInvoker(request *model.CreateRequestRequest) *CreateRequestInvoker {
	requestDef := GenReqDefForCreateRequest()
	return &CreateRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResult Show Result接口
//
// get the result of the code generation request.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowResult(request *model.ShowResultRequest) (*model.ShowResultResponse, error) {
	requestDef := GenReqDefForShowResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResultResponse), nil
	}
}

// ShowResultInvoker Show Result接口
func (c *CloudIDEClient) ShowResultInvoker(request *model.ShowResultRequest) *ShowResultInvoker {
	requestDef := GenReqDefForShowResult()
	return &ShowResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckInstanceAccess 查询用户是否有权限访问某个IDE实例
//
// 查询用户是否有权限访问某个IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CheckInstanceAccess(request *model.CheckInstanceAccessRequest) (*model.CheckInstanceAccessResponse, error) {
	requestDef := GenReqDefForCheckInstanceAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckInstanceAccessResponse), nil
	}
}

// CheckInstanceAccessInvoker 查询用户是否有权限访问某个IDE实例
func (c *CloudIDEClient) CheckInstanceAccessInvoker(request *model.CheckInstanceAccessRequest) *CheckInstanceAccessInvoker {
	requestDef := GenReqDefForCheckInstanceAccess()
	return &CheckInstanceAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckName 查询IDE实例名是否重复
//
// 查询IDE实例名是否重复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CheckName(request *model.CheckNameRequest) (*model.CheckNameResponse, error) {
	requestDef := GenReqDefForCheckName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNameResponse), nil
	}
}

// CheckNameInvoker 查询IDE实例名是否重复
func (c *CloudIDEClient) CheckNameInvoker(request *model.CheckNameRequest) *CheckNameInvoker {
	requestDef := GenReqDefForCheckName()
	return &CheckNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建IDE实例
//
// 创建IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建IDE实例
func (c *CloudIDEClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceBy3rd 外部第三方集成商创建IDE实例
//
// 创建IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) CreateInstanceBy3rd(request *model.CreateInstanceBy3rdRequest) (*model.CreateInstanceBy3rdResponse, error) {
	requestDef := GenReqDefForCreateInstanceBy3rd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceBy3rdResponse), nil
	}
}

// CreateInstanceBy3rdInvoker 外部第三方集成商创建IDE实例
func (c *CloudIDEClient) CreateInstanceBy3rdInvoker(request *model.CreateInstanceBy3rdRequest) *CreateInstanceBy3rdInvoker {
	requestDef := GenReqDefForCreateInstanceBy3rd()
	return &CreateInstanceBy3rdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除IDE实例
//
// 删除IDE实例（同时删除磁盘数据）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除IDE实例
func (c *CloudIDEClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询IDE实例列表
//
// 查询IDE实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询IDE实例列表
func (c *CloudIDEClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrgInstances 查询某个租户下的IDE实例列表
//
// 查询某个租户下的IDE实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ListOrgInstances(request *model.ListOrgInstancesRequest) (*model.ListOrgInstancesResponse, error) {
	requestDef := GenReqDefForListOrgInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrgInstancesResponse), nil
	}
}

// ListOrgInstancesInvoker 查询某个租户下的IDE实例列表
func (c *CloudIDEClient) ListOrgInstancesInvoker(request *model.ListOrgInstancesRequest) *ListOrgInstancesInvoker {
	requestDef := GenReqDefForListOrgInstances()
	return &ListOrgInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询某个IDE实例
//
// 查询某个IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询某个IDE实例
func (c *CloudIDEClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceStatusInfo 查询某个IDE实例的状态
//
// 查询某个IDE实例的状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) ShowInstanceStatusInfo(request *model.ShowInstanceStatusInfoRequest) (*model.ShowInstanceStatusInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceStatusInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusInfoResponse), nil
	}
}

// ShowInstanceStatusInfoInvoker 查询某个IDE实例的状态
func (c *CloudIDEClient) ShowInstanceStatusInfoInvoker(request *model.ShowInstanceStatusInfoRequest) *ShowInstanceStatusInfoInvoker {
	requestDef := GenReqDefForShowInstanceStatusInfo()
	return &ShowInstanceStatusInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInstance 启动IDE实例
//
// 启动IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

// StartInstanceInvoker 启动IDE实例
func (c *CloudIDEClient) StartInstanceInvoker(request *model.StartInstanceRequest) *StartInstanceInvoker {
	requestDef := GenReqDefForStartInstance()
	return &StartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInstance 停止IDE实例
//
// 停止IDE实例（不删除磁盘数据）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

// StopInstanceInvoker 停止IDE实例
func (c *CloudIDEClient) StopInstanceInvoker(request *model.StopInstanceRequest) *StopInstanceInvoker {
	requestDef := GenReqDefForStopInstance()
	return &StopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改IDE实例
//
// 修改IDE实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改IDE实例
func (c *CloudIDEClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceActivity 刷新IDE实例活跃状态
//
// 刷新IDE实例活跃状态，超过该实例设置的过期时间后实例自动关闭。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudIDEClient) UpdateInstanceActivity(request *model.UpdateInstanceActivityRequest) (*model.UpdateInstanceActivityResponse, error) {
	requestDef := GenReqDefForUpdateInstanceActivity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceActivityResponse), nil
	}
}

// UpdateInstanceActivityInvoker 刷新IDE实例活跃状态
func (c *CloudIDEClient) UpdateInstanceActivityInvoker(request *model.UpdateInstanceActivityRequest) *UpdateInstanceActivityInvoker {
	requestDef := GenReqDefForUpdateInstanceActivity()
	return &UpdateInstanceActivityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
