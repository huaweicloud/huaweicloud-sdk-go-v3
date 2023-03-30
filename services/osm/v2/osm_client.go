package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/osm/v2/model"
)

type OsmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewOsmClient(hcClient *http_client.HcHttpClient) *OsmClient {
	return &OsmClient{HcClient: hcClient}
}

func OsmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CheckHosts 验证授权主机
//
// 验证授权主机密码是否正确
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CheckHosts(request *model.CheckHostsRequest) (*model.CheckHostsResponse, error) {
	requestDef := GenReqDefForCheckHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckHostsResponse), nil
	}
}

// CheckHostsInvoker 验证授权主机
func (c *OsmClient) CheckHostsInvoker(request *model.CheckHostsRequest) *CheckHostsInvoker {
	requestDef := GenReqDefForCheckHosts()
	return &CheckHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckNeedVerify 是否需要验证
//
// 是否需要验证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CheckNeedVerify(request *model.CheckNeedVerifyRequest) (*model.CheckNeedVerifyResponse, error) {
	requestDef := GenReqDefForCheckNeedVerify()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNeedVerifyResponse), nil
	}
}

// CheckNeedVerifyInvoker 是否需要验证
func (c *OsmClient) CheckNeedVerifyInvoker(request *model.CheckNeedVerifyRequest) *CheckNeedVerifyInvoker {
	requestDef := GenReqDefForCheckNeedVerify()
	return &CheckNeedVerifyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVerifyCodes 验证联系方式
//
// 验证联系方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CheckVerifyCodes(request *model.CheckVerifyCodesRequest) (*model.CheckVerifyCodesResponse, error) {
	requestDef := GenReqDefForCheckVerifyCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVerifyCodesResponse), nil
	}
}

// CheckVerifyCodesInvoker 验证联系方式
func (c *OsmClient) CheckVerifyCodesInvoker(request *model.CheckVerifyCodesRequest) *CheckVerifyCodesInvoker {
	requestDef := GenReqDefForCheckVerifyCodes()
	return &CheckVerifyCodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmAuthorizations 租户确认授权
//
// 租户确认授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ConfirmAuthorizations(request *model.ConfirmAuthorizationsRequest) (*model.ConfirmAuthorizationsResponse, error) {
	requestDef := GenReqDefForConfirmAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmAuthorizationsResponse), nil
	}
}

// ConfirmAuthorizationsInvoker 租户确认授权
func (c *OsmClient) ConfirmAuthorizationsInvoker(request *model.ConfirmAuthorizationsRequest) *ConfirmAuthorizationsInvoker {
	requestDef := GenReqDefForConfirmAuthorizations()
	return &ConfirmAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAskQuestion 语料提问
//
// 基于语料的一次问答
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateAskQuestion(request *model.CreateAskQuestionRequest) (*model.CreateAskQuestionResponse, error) {
	requestDef := GenReqDefForCreateAskQuestion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAskQuestionResponse), nil
	}
}

// CreateAskQuestionInvoker 语料提问
func (c *OsmClient) CreateAskQuestionInvoker(request *model.CreateAskQuestionRequest) *CreateAskQuestionInvoker {
	requestDef := GenReqDefForCreateAskQuestion()
	return &CreateAskQuestionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorization 创建授权
//
// 创建授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateAuthorization(request *model.CreateAuthorizationRequest) (*model.CreateAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizationResponse), nil
	}
}

// CreateAuthorizationInvoker 创建授权
func (c *OsmClient) CreateAuthorizationInvoker(request *model.CreateAuthorizationRequest) *CreateAuthorizationInvoker {
	requestDef := GenReqDefForCreateAuthorization()
	return &CreateAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCaseExtendsParam 提交工单扩展参数
//
// 提交工单扩展参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateCaseExtendsParam(request *model.CreateCaseExtendsParamRequest) (*model.CreateCaseExtendsParamResponse, error) {
	requestDef := GenReqDefForCreateCaseExtendsParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCaseExtendsParamResponse), nil
	}
}

// CreateCaseExtendsParamInvoker 提交工单扩展参数
func (c *OsmClient) CreateCaseExtendsParamInvoker(request *model.CreateCaseExtendsParamRequest) *CreateCaseExtendsParamInvoker {
	requestDef := GenReqDefForCreateCaseExtendsParam()
	return &CreateCaseExtendsParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCaseLabels 添加工单关联标签接口
//
// 添加工单关联标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateCaseLabels(request *model.CreateCaseLabelsRequest) (*model.CreateCaseLabelsResponse, error) {
	requestDef := GenReqDefForCreateCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCaseLabelsResponse), nil
	}
}

// CreateCaseLabelsInvoker 添加工单关联标签接口
func (c *OsmClient) CreateCaseLabelsInvoker(request *model.CreateCaseLabelsRequest) *CreateCaseLabelsInvoker {
	requestDef := GenReqDefForCreateCaseLabels()
	return &CreateCaseLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCases 创建工单
//
// 创建工单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateCases(request *model.CreateCasesRequest) (*model.CreateCasesResponse, error) {
	requestDef := GenReqDefForCreateCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCasesResponse), nil
	}
}

// CreateCasesInvoker 创建工单
func (c *OsmClient) CreateCasesInvoker(request *model.CreateCasesRequest) *CreateCasesInvoker {
	requestDef := GenReqDefForCreateCases()
	return &CreateCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnoseFeedback 用户反馈是否有帮助
//
// 用户反馈是否有帮助
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateDiagnoseFeedback(request *model.CreateDiagnoseFeedbackRequest) (*model.CreateDiagnoseFeedbackResponse, error) {
	requestDef := GenReqDefForCreateDiagnoseFeedback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnoseFeedbackResponse), nil
	}
}

// CreateDiagnoseFeedbackInvoker 用户反馈是否有帮助
func (c *OsmClient) CreateDiagnoseFeedbackInvoker(request *model.CreateDiagnoseFeedbackRequest) *CreateDiagnoseFeedbackInvoker {
	requestDef := GenReqDefForCreateDiagnoseFeedback()
	return &CreateDiagnoseFeedbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnoseJob 开始一键诊断
//
// 开始一键诊断
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateDiagnoseJob(request *model.CreateDiagnoseJobRequest) (*model.CreateDiagnoseJobResponse, error) {
	requestDef := GenReqDefForCreateDiagnoseJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnoseJobResponse), nil
	}
}

// CreateDiagnoseJobInvoker 开始一键诊断
func (c *OsmClient) CreateDiagnoseJobInvoker(request *model.CreateDiagnoseJobRequest) *CreateDiagnoseJobInvoker {
	requestDef := GenReqDefForCreateDiagnoseJob()
	return &CreateDiagnoseJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvaluate 问答满意度评价
//
// 一次问答完毕后, 针对这一次问答提交满意度评价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateEvaluate(request *model.CreateEvaluateRequest) (*model.CreateEvaluateResponse, error) {
	requestDef := GenReqDefForCreateEvaluate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEvaluateResponse), nil
	}
}

// CreateEvaluateInvoker 问答满意度评价
func (c *OsmClient) CreateEvaluateInvoker(request *model.CreateEvaluateRequest) *CreateEvaluateInvoker {
	requestDef := GenReqDefForCreateEvaluate()
	return &CreateEvaluateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFeedback 创建举报反馈
//
// 创建举报反馈
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateFeedback(request *model.CreateFeedbackRequest) (*model.CreateFeedbackResponse, error) {
	requestDef := GenReqDefForCreateFeedback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFeedbackResponse), nil
	}
}

// CreateFeedbackInvoker 创建举报反馈
func (c *OsmClient) CreateFeedbackInvoker(request *model.CreateFeedbackRequest) *CreateFeedbackInvoker {
	requestDef := GenReqDefForCreateFeedback()
	return &CreateFeedbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLabels 创建标签
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateLabels(request *model.CreateLabelsRequest) (*model.CreateLabelsResponse, error) {
	requestDef := GenReqDefForCreateLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLabelsResponse), nil
	}
}

// CreateLabelsInvoker 创建标签
func (c *OsmClient) CreateLabelsInvoker(request *model.CreateLabelsRequest) *CreateLabelsInvoker {
	requestDef := GenReqDefForCreateLabels()
	return &CreateLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMessages 提交留言
//
// 提交留言
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateMessages(request *model.CreateMessagesRequest) (*model.CreateMessagesResponse, error) {
	requestDef := GenReqDefForCreateMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessagesResponse), nil
	}
}

// CreateMessagesInvoker 提交留言
func (c *OsmClient) CreateMessagesInvoker(request *model.CreateMessagesRequest) *CreateMessagesInvoker {
	requestDef := GenReqDefForCreateMessages()
	return &CreateMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivileges 创建授权
//
// 创建授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreatePrivileges(request *model.CreatePrivilegesRequest) (*model.CreatePrivilegesResponse, error) {
	requestDef := GenReqDefForCreatePrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivilegesResponse), nil
	}
}

// CreatePrivilegesInvoker 创建授权
func (c *OsmClient) CreatePrivilegesInvoker(request *model.CreatePrivilegesRequest) *CreatePrivilegesInvoker {
	requestDef := GenReqDefForCreatePrivileges()
	return &CreatePrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQaAsk 新问答接口
//
// 支持多轮流程问答接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateQaAsk(request *model.CreateQaAskRequest) (*model.CreateQaAskResponse, error) {
	requestDef := GenReqDefForCreateQaAsk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQaAskResponse), nil
	}
}

// CreateQaAskInvoker 新问答接口
func (c *OsmClient) CreateQaAskInvoker(request *model.CreateQaAskRequest) *CreateQaAskInvoker {
	requestDef := GenReqDefForCreateQaAsk()
	return &CreateQaAskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQaFeedbacks 反馈评价
//
// 提交/取消反馈评价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateQaFeedbacks(request *model.CreateQaFeedbacksRequest) (*model.CreateQaFeedbacksResponse, error) {
	requestDef := GenReqDefForCreateQaFeedbacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQaFeedbacksResponse), nil
	}
}

// CreateQaFeedbacksInvoker 反馈评价
func (c *OsmClient) CreateQaFeedbacksInvoker(request *model.CreateQaFeedbacksRequest) *CreateQaFeedbacksInvoker {
	requestDef := GenReqDefForCreateQaFeedbacks()
	return &CreateQaFeedbacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQuestionInSession 提问（在制定的会话中）
//
// 提问（在制定的会话中）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateQuestionInSession(request *model.CreateQuestionInSessionRequest) (*model.CreateQuestionInSessionResponse, error) {
	requestDef := GenReqDefForCreateQuestionInSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQuestionInSessionResponse), nil
	}
}

// CreateQuestionInSessionInvoker 提问（在制定的会话中）
func (c *OsmClient) CreateQuestionInSessionInvoker(request *model.CreateQuestionInSessionRequest) *CreateQuestionInSessionInvoker {
	requestDef := GenReqDefForCreateQuestionInSession()
	return &CreateQuestionInSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRelations 创建关联
//
// 创建关联，一个工单最多支持3个关联
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateRelations(request *model.CreateRelationsRequest) (*model.CreateRelationsResponse, error) {
	requestDef := GenReqDefForCreateRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRelationsResponse), nil
	}
}

// CreateRelationsInvoker 创建关联
func (c *OsmClient) CreateRelationsInvoker(request *model.CreateRelationsRequest) *CreateRelationsInvoker {
	requestDef := GenReqDefForCreateRelations()
	return &CreateRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScores 提交评分
//
// 提交评分
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateScores(request *model.CreateScoresRequest) (*model.CreateScoresResponse, error) {
	requestDef := GenReqDefForCreateScores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScoresResponse), nil
	}
}

// CreateScoresInvoker 提交评分
func (c *OsmClient) CreateScoresInvoker(request *model.CreateScoresRequest) *CreateScoresInvoker {
	requestDef := GenReqDefForCreateScores()
	return &CreateScoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSession 创建问答会话
//
// 用于创建问答会话, 创建会话后可开始问答
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) CreateSession(request *model.CreateSessionRequest) (*model.CreateSessionResponse, error) {
	requestDef := GenReqDefForCreateSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSessionResponse), nil
	}
}

// CreateSessionInvoker 创建问答会话
func (c *OsmClient) CreateSessionInvoker(request *model.CreateSessionRequest) *CreateSessionInvoker {
	requestDef := GenReqDefForCreateSession()
	return &CreateSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccessories 删除附件
//
// 删除附件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DeleteAccessories(request *model.DeleteAccessoriesRequest) (*model.DeleteAccessoriesResponse, error) {
	requestDef := GenReqDefForDeleteAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessoriesResponse), nil
	}
}

// DeleteAccessoriesInvoker 删除附件
func (c *OsmClient) DeleteAccessoriesInvoker(request *model.DeleteAccessoriesRequest) *DeleteAccessoriesInvoker {
	requestDef := GenReqDefForDeleteAccessories()
	return &DeleteAccessoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCaseLabels 删除工单关联标签接口
//
// 删除指定工单关联标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DeleteCaseLabels(request *model.DeleteCaseLabelsRequest) (*model.DeleteCaseLabelsResponse, error) {
	requestDef := GenReqDefForDeleteCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCaseLabelsResponse), nil
	}
}

// DeleteCaseLabelsInvoker 删除工单关联标签接口
func (c *OsmClient) DeleteCaseLabelsInvoker(request *model.DeleteCaseLabelsRequest) *DeleteCaseLabelsInvoker {
	requestDef := GenReqDefForDeleteCaseLabels()
	return &DeleteCaseLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLabels 删除标签
//
// 删除标签，同时会删除工单与标签关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DeleteLabels(request *model.DeleteLabelsRequest) (*model.DeleteLabelsResponse, error) {
	requestDef := GenReqDefForDeleteLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLabelsResponse), nil
	}
}

// DeleteLabelsInvoker 删除标签
func (c *OsmClient) DeleteLabelsInvoker(request *model.DeleteLabelsRequest) *DeleteLabelsInvoker {
	requestDef := GenReqDefForDeleteLabels()
	return &DeleteLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRelation 删除关联
//
// 删除关联
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DeleteRelation(request *model.DeleteRelationRequest) (*model.DeleteRelationResponse, error) {
	requestDef := GenReqDefForDeleteRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRelationResponse), nil
	}
}

// DeleteRelationInvoker 删除关联
func (c *OsmClient) DeleteRelationInvoker(request *model.DeleteRelationRequest) *DeleteRelationInvoker {
	requestDef := GenReqDefForDeleteRelation()
	return &DeleteRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAccessories 下载附件
//
// 下载附件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DownloadAccessories(request *model.DownloadAccessoriesRequest) (*model.DownloadAccessoriesResponse, error) {
	requestDef := GenReqDefForDownloadAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAccessoriesResponse), nil
	}
}

// DownloadAccessoriesInvoker 下载附件
func (c *OsmClient) DownloadAccessoriesInvoker(request *model.DownloadAccessoriesRequest) *DownloadAccessoriesInvoker {
	requestDef := GenReqDefForDownloadAccessories()
	return &DownloadAccessoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadCases 工单导出
//
// 工单导出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DownloadCases(request *model.DownloadCasesRequest) (*model.DownloadCasesResponse, error) {
	requestDef := GenReqDefForDownloadCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadCasesResponse), nil
	}
}

// DownloadCasesInvoker 工单导出
func (c *OsmClient) DownloadCasesInvoker(request *model.DownloadCasesRequest) *DownloadCasesInvoker {
	requestDef := GenReqDefForDownloadCases()
	return &DownloadCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadImages 图片展示
//
// 返回图片内容，用于页面直接展示
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) DownloadImages(request *model.DownloadImagesRequest) (*model.DownloadImagesResponse, error) {
	requestDef := GenReqDefForDownloadImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadImagesResponse), nil
	}
}

// DownloadImagesInvoker 图片展示
func (c *OsmClient) DownloadImagesInvoker(request *model.DownloadImagesRequest) *DownloadImagesInvoker {
	requestDef := GenReqDefForDownloadImages()
	return &DownloadImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessoryAccessUrls 租户批量获取下载链接
//
// 租户批量获取下载链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListAccessoryAccessUrls(request *model.ListAccessoryAccessUrlsRequest) (*model.ListAccessoryAccessUrlsResponse, error) {
	requestDef := GenReqDefForListAccessoryAccessUrls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessoryAccessUrlsResponse), nil
	}
}

// ListAccessoryAccessUrlsInvoker 租户批量获取下载链接
func (c *OsmClient) ListAccessoryAccessUrlsInvoker(request *model.ListAccessoryAccessUrlsRequest) *ListAccessoryAccessUrlsInvoker {
	requestDef := GenReqDefForListAccessoryAccessUrls()
	return &ListAccessoryAccessUrlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgencies 查询委托
//
// 查询委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesResponse), nil
	}
}

// ListAgenciesInvoker 查询委托
func (c *OsmClient) ListAgenciesInvoker(request *model.ListAgenciesRequest) *ListAgenciesInvoker {
	requestDef := GenReqDefForListAgencies()
	return &ListAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAreaCodes 查询国家码
//
// 查询国家码，用于提交工单页面填写联系方式使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListAreaCodes(request *model.ListAreaCodesRequest) (*model.ListAreaCodesResponse, error) {
	requestDef := GenReqDefForListAreaCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAreaCodesResponse), nil
	}
}

// ListAreaCodesInvoker 查询国家码
func (c *OsmClient) ListAreaCodesInvoker(request *model.ListAreaCodesRequest) *ListAreaCodesInvoker {
	requestDef := GenReqDefForListAreaCodes()
	return &ListAreaCodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArticles 查询案例
//
// 查询满足指定条件的案例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListArticles(request *model.ListArticlesRequest) (*model.ListArticlesResponse, error) {
	requestDef := GenReqDefForListArticles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArticlesResponse), nil
	}
}

// ListArticlesInvoker 查询案例
func (c *OsmClient) ListArticlesInvoker(request *model.ListArticlesRequest) *ListArticlesInvoker {
	requestDef := GenReqDefForListArticles()
	return &ListArticlesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizations 查看授权列表
//
// 查询授权列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListAuthorizations(request *model.ListAuthorizationsRequest) (*model.ListAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizationsResponse), nil
	}
}

// ListAuthorizationsInvoker 查看授权列表
func (c *OsmClient) ListAuthorizationsInvoker(request *model.ListAuthorizationsRequest) *ListAuthorizationsInvoker {
	requestDef := GenReqDefForListAuthorizations()
	return &ListAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseCategories 查询工单类目列表
//
// 查询工单类目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseCategories(request *model.ListCaseCategoriesRequest) (*model.ListCaseCategoriesResponse, error) {
	requestDef := GenReqDefForListCaseCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCategoriesResponse), nil
	}
}

// ListCaseCategoriesInvoker 查询工单类目列表
func (c *OsmClient) ListCaseCategoriesInvoker(request *model.ListCaseCategoriesRequest) *ListCaseCategoriesInvoker {
	requestDef := GenReqDefForListCaseCategories()
	return &ListCaseCategoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseCcEmails 查询工单抄送邮箱
//
// 查询工单抄送邮箱
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseCcEmails(request *model.ListCaseCcEmailsRequest) (*model.ListCaseCcEmailsResponse, error) {
	requestDef := GenReqDefForListCaseCcEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCcEmailsResponse), nil
	}
}

// ListCaseCcEmailsInvoker 查询工单抄送邮箱
func (c *OsmClient) ListCaseCcEmailsInvoker(request *model.ListCaseCcEmailsRequest) *ListCaseCcEmailsInvoker {
	requestDef := GenReqDefForListCaseCcEmails()
	return &ListCaseCcEmailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseCounts 统计各状态工单数量
//
// 统计各状态工单数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseCounts(request *model.ListCaseCountsRequest) (*model.ListCaseCountsResponse, error) {
	requestDef := GenReqDefForListCaseCounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCountsResponse), nil
	}
}

// ListCaseCountsInvoker 统计各状态工单数量
func (c *OsmClient) ListCaseCountsInvoker(request *model.ListCaseCountsRequest) *ListCaseCountsInvoker {
	requestDef := GenReqDefForListCaseCounts()
	return &ListCaseCountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseLabels 查询工单关联标签接口
//
// 查询工单关联标签接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseLabels(request *model.ListCaseLabelsRequest) (*model.ListCaseLabelsResponse, error) {
	requestDef := GenReqDefForListCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseLabelsResponse), nil
	}
}

// ListCaseLabelsInvoker 查询工单关联标签接口
func (c *OsmClient) ListCaseLabelsInvoker(request *model.ListCaseLabelsRequest) *ListCaseLabelsInvoker {
	requestDef := GenReqDefForListCaseLabels()
	return &ListCaseLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseLimits 查询工单限制，比如抄送邮箱个数等
//
// 查询工单限制，比如抄送邮箱个数等
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseLimits(request *model.ListCaseLimitsRequest) (*model.ListCaseLimitsResponse, error) {
	requestDef := GenReqDefForListCaseLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseLimitsResponse), nil
	}
}

// ListCaseLimitsInvoker 查询工单限制，比如抄送邮箱个数等
func (c *OsmClient) ListCaseLimitsInvoker(request *model.ListCaseLimitsRequest) *ListCaseLimitsInvoker {
	requestDef := GenReqDefForListCaseLimits()
	return &ListCaseLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseOperateLogs 查询工单操作日志
//
// 查询工单操作日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseOperateLogs(request *model.ListCaseOperateLogsRequest) (*model.ListCaseOperateLogsResponse, error) {
	requestDef := GenReqDefForListCaseOperateLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseOperateLogsResponse), nil
	}
}

// ListCaseOperateLogsInvoker 查询工单操作日志
func (c *OsmClient) ListCaseOperateLogsInvoker(request *model.ListCaseOperateLogsRequest) *ListCaseOperateLogsInvoker {
	requestDef := GenReqDefForListCaseOperateLogs()
	return &ListCaseOperateLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseQuotas 查询工单配额
//
// 查询工单配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseQuotas(request *model.ListCaseQuotasRequest) (*model.ListCaseQuotasResponse, error) {
	requestDef := GenReqDefForListCaseQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseQuotasResponse), nil
	}
}

// ListCaseQuotasInvoker 查询工单配额
func (c *OsmClient) ListCaseQuotasInvoker(request *model.ListCaseQuotasRequest) *ListCaseQuotasInvoker {
	requestDef := GenReqDefForListCaseQuotas()
	return &ListCaseQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCaseTemplates 查询问题类型对应模板
//
// 查询问题类型对应模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCaseTemplates(request *model.ListCaseTemplatesRequest) (*model.ListCaseTemplatesResponse, error) {
	requestDef := GenReqDefForListCaseTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseTemplatesResponse), nil
	}
}

// ListCaseTemplatesInvoker 查询问题类型对应模板
func (c *OsmClient) ListCaseTemplatesInvoker(request *model.ListCaseTemplatesRequest) *ListCaseTemplatesInvoker {
	requestDef := GenReqDefForListCaseTemplates()
	return &ListCaseTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCases 查询工单列表接口
//
// 查询工单列表接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCases(request *model.ListCasesRequest) (*model.ListCasesResponse, error) {
	requestDef := GenReqDefForListCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCasesResponse), nil
	}
}

// ListCasesInvoker 查询工单列表接口
func (c *OsmClient) ListCasesInvoker(request *model.ListCasesRequest) *ListCasesInvoker {
	requestDef := GenReqDefForListCases()
	return &ListCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomersRegions 查询用户关联的region
//
// 查询用户关联的region
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListCustomersRegions(request *model.ListCustomersRegionsRequest) (*model.ListCustomersRegionsResponse, error) {
	requestDef := GenReqDefForListCustomersRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomersRegionsResponse), nil
	}
}

// ListCustomersRegionsInvoker 查询用户关联的region
func (c *OsmClient) ListCustomersRegionsInvoker(request *model.ListCustomersRegionsRequest) *ListCustomersRegionsInvoker {
	requestDef := GenReqDefForListCustomersRegions()
	return &ListCustomersRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnoseItems 根据itemIds查询指定的诊断检查项执行结果
//
// 根据itemIds查询指定的诊断检查项执行结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListDiagnoseItems(request *model.ListDiagnoseItemsRequest) (*model.ListDiagnoseItemsResponse, error) {
	requestDef := GenReqDefForListDiagnoseItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnoseItemsResponse), nil
	}
}

// ListDiagnoseItemsInvoker 根据itemIds查询指定的诊断检查项执行结果
func (c *OsmClient) ListDiagnoseItemsInvoker(request *model.ListDiagnoseItemsRequest) *ListDiagnoseItemsInvoker {
	requestDef := GenReqDefForListDiagnoseItems()
	return &ListDiagnoseItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnoseJob 查询诊断任务执行结果
//
// 查询诊断任务执行结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListDiagnoseJob(request *model.ListDiagnoseJobRequest) (*model.ListDiagnoseJobResponse, error) {
	requestDef := GenReqDefForListDiagnoseJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnoseJobResponse), nil
	}
}

// ListDiagnoseJobInvoker 查询诊断任务执行结果
func (c *OsmClient) ListDiagnoseJobInvoker(request *model.ListDiagnoseJobRequest) *ListDiagnoseJobInvoker {
	requestDef := GenReqDefForListDiagnoseJob()
	return &ListDiagnoseJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnoseRecords 查询诊断历史记录列表
//
// 查询诊断历史记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListDiagnoseRecords(request *model.ListDiagnoseRecordsRequest) (*model.ListDiagnoseRecordsResponse, error) {
	requestDef := GenReqDefForListDiagnoseRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnoseRecordsResponse), nil
	}
}

// ListDiagnoseRecordsInvoker 查询诊断历史记录列表
func (c *OsmClient) ListDiagnoseRecordsInvoker(request *model.ListDiagnoseRecordsRequest) *ListDiagnoseRecordsInvoker {
	requestDef := GenReqDefForListDiagnoseRecords()
	return &ListDiagnoseRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExtendsParams 查询附加参数
//
// 提单时，根据不同的产品或者问题类型，会存在不同的一些附加参数填写
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListExtendsParams(request *model.ListExtendsParamsRequest) (*model.ListExtendsParamsResponse, error) {
	requestDef := GenReqDefForListExtendsParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExtendsParamsResponse), nil
	}
}

// ListExtendsParamsInvoker 查询附加参数
func (c *OsmClient) ListExtendsParamsInvoker(request *model.ListExtendsParamsRequest) *ListExtendsParamsInvoker {
	requestDef := GenReqDefForListExtendsParams()
	return &ListExtendsParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeedbackOption 查询反馈选项
//
// 查询符合条件的反馈选项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListFeedbackOption(request *model.ListFeedbackOptionRequest) (*model.ListFeedbackOptionResponse, error) {
	requestDef := GenReqDefForListFeedbackOption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeedbackOptionResponse), nil
	}
}

// ListFeedbackOptionInvoker 查询反馈选项
func (c *OsmClient) ListFeedbackOptionInvoker(request *model.ListFeedbackOptionRequest) *ListFeedbackOptionInvoker {
	requestDef := GenReqDefForListFeedbackOption()
	return &ListFeedbackOptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHasVerifiedContacts 查询已验证的列表
//
// 查询已验证的列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListHasVerifiedContacts(request *model.ListHasVerifiedContactsRequest) (*model.ListHasVerifiedContactsResponse, error) {
	requestDef := GenReqDefForListHasVerifiedContacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHasVerifiedContactsResponse), nil
	}
}

// ListHasVerifiedContactsInvoker 查询已验证的列表
func (c *OsmClient) ListHasVerifiedContactsInvoker(request *model.ListHasVerifiedContactsRequest) *ListHasVerifiedContactsInvoker {
	requestDef := GenReqDefForListHasVerifiedContacts()
	return &ListHasVerifiedContactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryOperateLogs 查询堡垒机历史操作记录
//
// 查询堡垒机历史操作记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListHistoryOperateLogs(request *model.ListHistoryOperateLogsRequest) (*model.ListHistoryOperateLogsResponse, error) {
	requestDef := GenReqDefForListHistoryOperateLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOperateLogsResponse), nil
	}
}

// ListHistoryOperateLogsInvoker 查询堡垒机历史操作记录
func (c *OsmClient) ListHistoryOperateLogsInvoker(request *model.ListHistoryOperateLogsRequest) *ListHistoryOperateLogsInvoker {
	requestDef := GenReqDefForListHistoryOperateLogs()
	return &ListHistoryOperateLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistorySessions 查询堡垒机历史会话列表
//
// 查询堡垒机历史会话列
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListHistorySessions(request *model.ListHistorySessionsRequest) (*model.ListHistorySessionsResponse, error) {
	requestDef := GenReqDefForListHistorySessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistorySessionsResponse), nil
	}
}

// ListHistorySessionsInvoker 查询堡垒机历史会话列表
func (c *OsmClient) ListHistorySessionsInvoker(request *model.ListHistorySessionsRequest) *ListHistorySessionsInvoker {
	requestDef := GenReqDefForListHistorySessions()
	return &ListHistorySessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLabels 查询标签
//
// 查询标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListLabels(request *model.ListLabelsRequest) (*model.ListLabelsResponse, error) {
	requestDef := GenReqDefForListLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLabelsResponse), nil
	}
}

// ListLabelsInvoker 查询标签
func (c *OsmClient) ListLabelsInvoker(request *model.ListLabelsRequest) *ListLabelsInvoker {
	requestDef := GenReqDefForListLabels()
	return &ListLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessages 查询留言
//
// 查询留言
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListMessages(request *model.ListMessagesRequest) (*model.ListMessagesResponse, error) {
	requestDef := GenReqDefForListMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesResponse), nil
	}
}

// ListMessagesInvoker 查询留言
func (c *OsmClient) ListMessagesInvoker(request *model.ListMessagesRequest) *ListMessagesInvoker {
	requestDef := GenReqDefForListMessages()
	return &ListMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMoreInstantMessages 查询更多留言
//
// 查询更多留言
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListMoreInstantMessages(request *model.ListMoreInstantMessagesRequest) (*model.ListMoreInstantMessagesResponse, error) {
	requestDef := GenReqDefForListMoreInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMoreInstantMessagesResponse), nil
	}
}

// ListMoreInstantMessagesInvoker 查询更多留言
func (c *OsmClient) ListMoreInstantMessagesInvoker(request *model.ListMoreInstantMessagesRequest) *ListMoreInstantMessagesInvoker {
	requestDef := GenReqDefForListMoreInstantMessages()
	return &ListMoreInstantMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNewInstantMessages 轮询查询即时消息
//
// 轮询查询即时消息接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListNewInstantMessages(request *model.ListNewInstantMessagesRequest) (*model.ListNewInstantMessagesResponse, error) {
	requestDef := GenReqDefForListNewInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNewInstantMessagesResponse), nil
	}
}

// ListNewInstantMessagesInvoker 轮询查询即时消息
func (c *OsmClient) ListNewInstantMessagesInvoker(request *model.ListNewInstantMessagesRequest) *ListNewInstantMessagesInvoker {
	requestDef := GenReqDefForListNewInstantMessages()
	return &ListNewInstantMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotices 查询公告
//
// 查询满足指定条件的公告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListNotices(request *model.ListNoticesRequest) (*model.ListNoticesResponse, error) {
	requestDef := GenReqDefForListNotices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNoticesResponse), nil
	}
}

// ListNoticesInvoker 查询公告
func (c *OsmClient) ListNoticesInvoker(request *model.ListNoticesRequest) *ListNoticesInvoker {
	requestDef := GenReqDefForListNotices()
	return &ListNoticesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivileges 查询工单权限
//
// 查询工单权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListPrivileges(request *model.ListPrivilegesRequest) (*model.ListPrivilegesResponse, error) {
	requestDef := GenReqDefForListPrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivilegesResponse), nil
	}
}

// ListPrivilegesInvoker 查询工单权限
func (c *OsmClient) ListPrivilegesInvoker(request *model.ListPrivilegesRequest) *ListPrivilegesInvoker {
	requestDef := GenReqDefForListPrivileges()
	return &ListPrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProblemTypes 查询问题类型列表
//
// 提交工单时，选择产品类型之后选择对应的问题列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListProblemTypes(request *model.ListProblemTypesRequest) (*model.ListProblemTypesResponse, error) {
	requestDef := GenReqDefForListProblemTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProblemTypesResponse), nil
	}
}

// ListProblemTypesInvoker 查询问题类型列表
func (c *OsmClient) ListProblemTypesInvoker(request *model.ListProblemTypesRequest) *ListProblemTypesInvoker {
	requestDef := GenReqDefForListProblemTypes()
	return &ListProblemTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProductCategories 查询产品类型列表
//
// 查询产品类型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListProductCategories(request *model.ListProductCategoriesRequest) (*model.ListProductCategoriesResponse, error) {
	requestDef := GenReqDefForListProductCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductCategoriesResponse), nil
	}
}

// ListProductCategoriesInvoker 查询产品类型列表
func (c *OsmClient) ListProductCategoriesInvoker(request *model.ListProductCategoriesRequest) *ListProductCategoriesInvoker {
	requestDef := GenReqDefForListProductCategories()
	return &ListProductCategoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecommendWords 查询推荐热词
//
// 查询指定条件的推荐热词
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListRecommendWords(request *model.ListRecommendWordsRequest) (*model.ListRecommendWordsResponse, error) {
	requestDef := GenReqDefForListRecommendWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecommendWordsResponse), nil
	}
}

// ListRecommendWordsInvoker 查询推荐热词
func (c *OsmClient) ListRecommendWordsInvoker(request *model.ListRecommendWordsRequest) *ListRecommendWordsInvoker {
	requestDef := GenReqDefForListRecommendWords()
	return &ListRecommendWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查询区域列表
//
// 查询区域列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查询区域列表
func (c *OsmClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRelation 查询关联工单
//
// 查询工单的关联，返回关联工单的简要信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListRelation(request *model.ListRelationRequest) (*model.ListRelationResponse, error) {
	requestDef := GenReqDefForListRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelationResponse), nil
	}
}

// ListRelationInvoker 查询关联工单
func (c *OsmClient) ListRelationInvoker(request *model.ListRelationRequest) *ListRelationInvoker {
	requestDef := GenReqDefForListRelation()
	return &ListRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSatisfactionDimensions 工单满意度分类列表
//
// 工单满意度分类列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListSatisfactionDimensions(request *model.ListSatisfactionDimensionsRequest) (*model.ListSatisfactionDimensionsResponse, error) {
	requestDef := GenReqDefForListSatisfactionDimensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSatisfactionDimensionsResponse), nil
	}
}

// ListSatisfactionDimensionsInvoker 工单满意度分类列表
func (c *OsmClient) ListSatisfactionDimensionsInvoker(request *model.ListSatisfactionDimensionsRequest) *ListSatisfactionDimensionsInvoker {
	requestDef := GenReqDefForListSatisfactionDimensions()
	return &ListSatisfactionDimensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSeverities 查询问题严重性列表
//
// 查询问题严重性列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListSeverities(request *model.ListSeveritiesRequest) (*model.ListSeveritiesResponse, error) {
	requestDef := GenReqDefForListSeverities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSeveritiesResponse), nil
	}
}

// ListSeveritiesInvoker 查询问题严重性列表
func (c *OsmClient) ListSeveritiesInvoker(request *model.ListSeveritiesRequest) *ListSeveritiesInvoker {
	requestDef := GenReqDefForListSeverities()
	return &ListSeveritiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomers 查询子用户信息
//
// 查询子用户信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

// ListSubCustomersInvoker 查询子用户信息
func (c *OsmClient) ListSubCustomersInvoker(request *model.ListSubCustomersRequest) *ListSubCustomersInvoker {
	requestDef := GenReqDefForListSubCustomers()
	return &ListSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTools 查询工具
//
// 查询满足指定条件的工具列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListTools(request *model.ListToolsRequest) (*model.ListToolsResponse, error) {
	requestDef := GenReqDefForListTools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListToolsResponse), nil
	}
}

// ListToolsInvoker 查询工具
func (c *OsmClient) ListToolsInvoker(request *model.ListToolsRequest) *ListToolsInvoker {
	requestDef := GenReqDefForListTools()
	return &ListToolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransportHistories 查询堡垒机文件传输记录
//
// 查询堡垒机文件传输记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListTransportHistories(request *model.ListTransportHistoriesRequest) (*model.ListTransportHistoriesResponse, error) {
	requestDef := GenReqDefForListTransportHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransportHistoriesResponse), nil
	}
}

// ListTransportHistoriesInvoker 查询堡垒机文件传输记录
func (c *OsmClient) ListTransportHistoriesInvoker(request *model.ListTransportHistoriesRequest) *ListTransportHistoriesInvoker {
	requestDef := GenReqDefForListTransportHistories()
	return &ListTransportHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUnreadNewInstantMessages 查询未读消息
//
// 查询未读消息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ListUnreadNewInstantMessages(request *model.ListUnreadNewInstantMessagesRequest) (*model.ListUnreadNewInstantMessagesResponse, error) {
	requestDef := GenReqDefForListUnreadNewInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnreadNewInstantMessagesResponse), nil
	}
}

// ListUnreadNewInstantMessagesInvoker 查询未读消息
func (c *OsmClient) ListUnreadNewInstantMessagesInvoker(request *model.ListUnreadNewInstantMessagesRequest) *ListUnreadNewInstantMessagesInvoker {
	requestDef := GenReqDefForListUnreadNewInstantMessages()
	return &ListUnreadNewInstantMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokeMessage 撤回留言
//
// 撤回留言
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) RevokeMessage(request *model.RevokeMessageRequest) (*model.RevokeMessageResponse, error) {
	requestDef := GenReqDefForRevokeMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeMessageResponse), nil
	}
}

// RevokeMessageInvoker 撤回留言
func (c *OsmClient) RevokeMessageInvoker(request *model.RevokeMessageRequest) *RevokeMessageInvoker {
	requestDef := GenReqDefForRevokeMessage()
	return &RevokeMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVerifyCodes 获取验证码
//
// 获取验证码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) SendVerifyCodes(request *model.SendVerifyCodesRequest) (*model.SendVerifyCodesResponse, error) {
	requestDef := GenReqDefForSendVerifyCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerifyCodesResponse), nil
	}
}

// SendVerifyCodesInvoker 获取验证码
func (c *OsmClient) SendVerifyCodesInvoker(request *model.SendVerifyCodesRequest) *SendVerifyCodesInvoker {
	requestDef := GenReqDefForSendVerifyCodes()
	return &SendVerifyCodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessoryLimits 附件限制
//
// 查询附件的一下限制，比如大小，数量，文件类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowAccessoryLimits(request *model.ShowAccessoryLimitsRequest) (*model.ShowAccessoryLimitsResponse, error) {
	requestDef := GenReqDefForShowAccessoryLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessoryLimitsResponse), nil
	}
}

// ShowAccessoryLimitsInvoker 附件限制
func (c *OsmClient) ShowAccessoryLimitsInvoker(request *model.ShowAccessoryLimitsRequest) *ShowAccessoryLimitsInvoker {
	requestDef := GenReqDefForShowAccessoryLimits()
	return &ShowAccessoryLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssociatedQuestions 查询联想问题
//
// 根据当前输入, 联想相关的问题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowAssociatedQuestions(request *model.ShowAssociatedQuestionsRequest) (*model.ShowAssociatedQuestionsResponse, error) {
	requestDef := GenReqDefForShowAssociatedQuestions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssociatedQuestionsResponse), nil
	}
}

// ShowAssociatedQuestionsInvoker 查询联想问题
func (c *OsmClient) ShowAssociatedQuestionsInvoker(request *model.ShowAssociatedQuestionsRequest) *ShowAssociatedQuestionsInvoker {
	requestDef := GenReqDefForShowAssociatedQuestions()
	return &ShowAssociatedQuestionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthorizationDetail 查询授权详情
//
// 查询授权详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowAuthorizationDetail(request *model.ShowAuthorizationDetailRequest) (*model.ShowAuthorizationDetailResponse, error) {
	requestDef := GenReqDefForShowAuthorizationDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthorizationDetailResponse), nil
	}
}

// ShowAuthorizationDetailInvoker 查询授权详情
func (c *OsmClient) ShowAuthorizationDetailInvoker(request *model.ShowAuthorizationDetailRequest) *ShowAuthorizationDetailInvoker {
	requestDef := GenReqDefForShowAuthorizationDetail()
	return &ShowAuthorizationDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCaseDetail 查询工单详情
//
// 查询工单详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowCaseDetail(request *model.ShowCaseDetailRequest) (*model.ShowCaseDetailResponse, error) {
	requestDef := GenReqDefForShowCaseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCaseDetailResponse), nil
	}
}

// ShowCaseDetailInvoker 查询工单详情
func (c *OsmClient) ShowCaseDetailInvoker(request *model.ShowCaseDetailRequest) *ShowCaseDetailInvoker {
	requestDef := GenReqDefForShowCaseDetail()
	return &ShowCaseDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCaseExtendsParam 查询工单扩展参数
//
// 查询工单扩展参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowCaseExtendsParam(request *model.ShowCaseExtendsParamRequest) (*model.ShowCaseExtendsParamResponse, error) {
	requestDef := GenReqDefForShowCaseExtendsParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCaseExtendsParamResponse), nil
	}
}

// ShowCaseExtendsParamInvoker 查询工单扩展参数
func (c *OsmClient) ShowCaseExtendsParamInvoker(request *model.ShowCaseExtendsParamRequest) *ShowCaseExtendsParamInvoker {
	requestDef := GenReqDefForShowCaseExtendsParam()
	return &ShowCaseExtendsParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCaseStatus 查询某个工单状态
//
// 查询某个工单状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowCaseStatus(request *model.ShowCaseStatusRequest) (*model.ShowCaseStatusResponse, error) {
	requestDef := GenReqDefForShowCaseStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCaseStatusResponse), nil
	}
}

// ShowCaseStatusInvoker 查询某个工单状态
func (c *OsmClient) ShowCaseStatusInvoker(request *model.ShowCaseStatusRequest) *ShowCaseStatusInvoker {
	requestDef := GenReqDefForShowCaseStatus()
	return &ShowCaseStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfiguration 查询配置
//
// 查询配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowConfiguration(request *model.ShowConfigurationRequest) (*model.ShowConfigurationResponse, error) {
	requestDef := GenReqDefForShowConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationResponse), nil
	}
}

// ShowConfigurationInvoker 查询配置
func (c *OsmClient) ShowConfigurationInvoker(request *model.ShowConfigurationRequest) *ShowConfigurationInvoker {
	requestDef := GenReqDefForShowConfiguration()
	return &ShowConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerPrivilegePolicy 查询提单权限
//
// 查询提单权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowCustomerPrivilegePolicy(request *model.ShowCustomerPrivilegePolicyRequest) (*model.ShowCustomerPrivilegePolicyResponse, error) {
	requestDef := GenReqDefForShowCustomerPrivilegePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerPrivilegePolicyResponse), nil
	}
}

// ShowCustomerPrivilegePolicyInvoker 查询提单权限
func (c *OsmClient) ShowCustomerPrivilegePolicyInvoker(request *model.ShowCustomerPrivilegePolicyRequest) *ShowCustomerPrivilegePolicyInvoker {
	requestDef := GenReqDefForShowCustomerPrivilegePolicy()
	return &ShowCustomerPrivilegePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDownloadAccessoryUrl 附件下载地址
//
// 附件下载地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowDownloadAccessoryUrl(request *model.ShowDownloadAccessoryUrlRequest) (*model.ShowDownloadAccessoryUrlResponse, error) {
	requestDef := GenReqDefForShowDownloadAccessoryUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDownloadAccessoryUrlResponse), nil
	}
}

// ShowDownloadAccessoryUrlInvoker 附件下载地址
func (c *OsmClient) ShowDownloadAccessoryUrlInvoker(request *model.ShowDownloadAccessoryUrlRequest) *ShowDownloadAccessoryUrlInvoker {
	requestDef := GenReqDefForShowDownloadAccessoryUrl()
	return &ShowDownloadAccessoryUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLatestPublishedAgreement 查看最新发布版本协议详情
//
// 查看最新发布版本协议详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowLatestPublishedAgreement(request *model.ShowLatestPublishedAgreementRequest) (*model.ShowLatestPublishedAgreementResponse, error) {
	requestDef := GenReqDefForShowLatestPublishedAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestPublishedAgreementResponse), nil
	}
}

// ShowLatestPublishedAgreementInvoker 查看最新发布版本协议详情
func (c *OsmClient) ShowLatestPublishedAgreementInvoker(request *model.ShowLatestPublishedAgreementRequest) *ShowLatestPublishedAgreementInvoker {
	requestDef := GenReqDefForShowLatestPublishedAgreement()
	return &ShowLatestPublishedAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartnersCasesPrivilege 查询伙伴工单权限
//
// 查询伙伴工单权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowPartnersCasesPrivilege(request *model.ShowPartnersCasesPrivilegeRequest) (*model.ShowPartnersCasesPrivilegeResponse, error) {
	requestDef := GenReqDefForShowPartnersCasesPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartnersCasesPrivilegeResponse), nil
	}
}

// ShowPartnersCasesPrivilegeInvoker 查询伙伴工单权限
func (c *OsmClient) ShowPartnersCasesPrivilegeInvoker(request *model.ShowPartnersCasesPrivilegeRequest) *ShowPartnersCasesPrivilegeInvoker {
	requestDef := GenReqDefForShowPartnersCasesPrivilege()
	return &ShowPartnersCasesPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartnersServiceInfo 查询关联伙伴服务信息
//
// 查询关联伙伴服务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowPartnersServiceInfo(request *model.ShowPartnersServiceInfoRequest) (*model.ShowPartnersServiceInfoResponse, error) {
	requestDef := GenReqDefForShowPartnersServiceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartnersServiceInfoResponse), nil
	}
}

// ShowPartnersServiceInfoInvoker 查询关联伙伴服务信息
func (c *OsmClient) ShowPartnersServiceInfoInvoker(request *model.ShowPartnersServiceInfoRequest) *ShowPartnersServiceInfoInvoker {
	requestDef := GenReqDefForShowPartnersServiceInfo()
	return &ShowPartnersServiceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQaPairDetail 查询语料详情
//
// 查询指定语料的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowQaPairDetail(request *model.ShowQaPairDetailRequest) (*model.ShowQaPairDetailResponse, error) {
	requestDef := GenReqDefForShowQaPairDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQaPairDetailResponse), nil
	}
}

// ShowQaPairDetailInvoker 查询语料详情
func (c *OsmClient) ShowQaPairDetailInvoker(request *model.ShowQaPairDetailRequest) *ShowQaPairDetailInvoker {
	requestDef := GenReqDefForShowQaPairDetail()
	return &ShowQaPairDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQaPairs 查询语料
//
// 查询满足指定条件的语料列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowQaPairs(request *model.ShowQaPairsRequest) (*model.ShowQaPairsResponse, error) {
	requestDef := GenReqDefForShowQaPairs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQaPairsResponse), nil
	}
}

// ShowQaPairsInvoker 查询语料
func (c *OsmClient) ShowQaPairsInvoker(request *model.ShowQaPairsRequest) *ShowQaPairsInvoker {
	requestDef := GenReqDefForShowQaPairs()
	return &ShowQaPairsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSignedLatestPublishedAgreement 查询用户是否签署最新协议
//
// 查询用户是否签署最新协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowSignedLatestPublishedAgreement(request *model.ShowSignedLatestPublishedAgreementRequest) (*model.ShowSignedLatestPublishedAgreementResponse, error) {
	requestDef := GenReqDefForShowSignedLatestPublishedAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSignedLatestPublishedAgreementResponse), nil
	}
}

// ShowSignedLatestPublishedAgreementInvoker 查询用户是否签署最新协议
func (c *OsmClient) ShowSignedLatestPublishedAgreementInvoker(request *model.ShowSignedLatestPublishedAgreementRequest) *ShowSignedLatestPublishedAgreementInvoker {
	requestDef := GenReqDefForShowSignedLatestPublishedAgreement()
	return &ShowSignedLatestPublishedAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTheme 产品类型转为主题
//
// 产品类型转为主题
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) ShowTheme(request *model.ShowThemeRequest) (*model.ShowThemeResponse, error) {
	requestDef := GenReqDefForShowTheme()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowThemeResponse), nil
	}
}

// ShowThemeInvoker 产品类型转为主题
func (c *OsmClient) ShowThemeInvoker(request *model.ShowThemeRequest) *ShowThemeInvoker {
	requestDef := GenReqDefForShowTheme()
	return &ShowThemeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SignPublishedAgreement 签署协议
//
// 签署协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) SignPublishedAgreement(request *model.SignPublishedAgreementRequest) (*model.SignPublishedAgreementResponse, error) {
	requestDef := GenReqDefForSignPublishedAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignPublishedAgreementResponse), nil
	}
}

// SignPublishedAgreementInvoker 签署协议
func (c *OsmClient) SignPublishedAgreementInvoker(request *model.SignPublishedAgreementRequest) *SignPublishedAgreementInvoker {
	requestDef := GenReqDefForSignPublishedAgreement()
	return &SignPublishedAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuthorizations 拒绝|撤销授权
//
// 拒绝|撤销授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UpdateAuthorizations(request *model.UpdateAuthorizationsRequest) (*model.UpdateAuthorizationsResponse, error) {
	requestDef := GenReqDefForUpdateAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthorizationsResponse), nil
	}
}

// UpdateAuthorizationsInvoker 拒绝|撤销授权
func (c *OsmClient) UpdateAuthorizationsInvoker(request *model.UpdateAuthorizationsRequest) *UpdateAuthorizationsInvoker {
	requestDef := GenReqDefForUpdateAuthorizations()
	return &UpdateAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCaseContactInfo 修改联系方式
//
// 修改联系方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UpdateCaseContactInfo(request *model.UpdateCaseContactInfoRequest) (*model.UpdateCaseContactInfoResponse, error) {
	requestDef := GenReqDefForUpdateCaseContactInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCaseContactInfoResponse), nil
	}
}

// UpdateCaseContactInfoInvoker 修改联系方式
func (c *OsmClient) UpdateCaseContactInfoInvoker(request *model.UpdateCaseContactInfoRequest) *UpdateCaseContactInfoInvoker {
	requestDef := GenReqDefForUpdateCaseContactInfo()
	return &UpdateCaseContactInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCases 工单操作
//
// 工单操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UpdateCases(request *model.UpdateCasesRequest) (*model.UpdateCasesResponse, error) {
	requestDef := GenReqDefForUpdateCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCasesResponse), nil
	}
}

// UpdateCasesInvoker 工单操作
func (c *OsmClient) UpdateCasesInvoker(request *model.UpdateCasesRequest) *UpdateCasesInvoker {
	requestDef := GenReqDefForUpdateCases()
	return &UpdateCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLabels 修改标签
//
// 修改标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UpdateLabels(request *model.UpdateLabelsRequest) (*model.UpdateLabelsResponse, error) {
	requestDef := GenReqDefForUpdateLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLabelsResponse), nil
	}
}

// UpdateLabelsInvoker 修改标签
func (c *OsmClient) UpdateLabelsInvoker(request *model.UpdateLabelsRequest) *UpdateLabelsInvoker {
	requestDef := GenReqDefForUpdateLabels()
	return &UpdateLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNewInstantMessagesRead 设置消息已读
//
// 设置消息已读
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UpdateNewInstantMessagesRead(request *model.UpdateNewInstantMessagesReadRequest) (*model.UpdateNewInstantMessagesReadResponse, error) {
	requestDef := GenReqDefForUpdateNewInstantMessagesRead()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNewInstantMessagesReadResponse), nil
	}
}

// UpdateNewInstantMessagesReadInvoker 设置消息已读
func (c *OsmClient) UpdateNewInstantMessagesReadInvoker(request *model.UpdateNewInstantMessagesReadRequest) *UpdateNewInstantMessagesReadInvoker {
	requestDef := GenReqDefForUpdateNewInstantMessagesRead()
	return &UpdateNewInstantMessagesReadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadJsonAccessories 上传附件
//
// 上传附件给SDK使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OsmClient) UploadJsonAccessories(request *model.UploadJsonAccessoriesRequest) (*model.UploadJsonAccessoriesResponse, error) {
	requestDef := GenReqDefForUploadJsonAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJsonAccessoriesResponse), nil
	}
}

// UploadJsonAccessoriesInvoker 上传附件
func (c *OsmClient) UploadJsonAccessoriesInvoker(request *model.UploadJsonAccessoriesRequest) *UploadJsonAccessoriesInvoker {
	requestDef := GenReqDefForUploadJsonAccessories()
	return &UploadJsonAccessoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
