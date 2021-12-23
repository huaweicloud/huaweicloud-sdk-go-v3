package v2

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/osm/v2/model"
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

//验证授权主机密码是否正确
func (c *OsmClient) CheckHosts(request *model.CheckHostsRequest) (*model.CheckHostsResponse, error) {
	requestDef := GenReqDefForCheckHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckHostsResponse), nil
	}
}

//是否需要验证
func (c *OsmClient) CheckNeedVerify(request *model.CheckNeedVerifyRequest) (*model.CheckNeedVerifyResponse, error) {
	requestDef := GenReqDefForCheckNeedVerify()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNeedVerifyResponse), nil
	}
}

//验证联系方式
func (c *OsmClient) CheckVerifyCodes(request *model.CheckVerifyCodesRequest) (*model.CheckVerifyCodesResponse, error) {
	requestDef := GenReqDefForCheckVerifyCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVerifyCodesResponse), nil
	}
}

//租户确认授权
func (c *OsmClient) ConfirmAuthorizations(request *model.ConfirmAuthorizationsRequest) (*model.ConfirmAuthorizationsResponse, error) {
	requestDef := GenReqDefForConfirmAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmAuthorizationsResponse), nil
	}
}

//添加工单关联标签接口
func (c *OsmClient) CreateCaseLabels(request *model.CreateCaseLabelsRequest) (*model.CreateCaseLabelsResponse, error) {
	requestDef := GenReqDefForCreateCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCaseLabelsResponse), nil
	}
}

//创建工单
func (c *OsmClient) CreateCases(request *model.CreateCasesRequest) (*model.CreateCasesResponse, error) {
	requestDef := GenReqDefForCreateCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCasesResponse), nil
	}
}

//创建标签
func (c *OsmClient) CreateLabels(request *model.CreateLabelsRequest) (*model.CreateLabelsResponse, error) {
	requestDef := GenReqDefForCreateLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLabelsResponse), nil
	}
}

//提交留言
func (c *OsmClient) CreateMessages(request *model.CreateMessagesRequest) (*model.CreateMessagesResponse, error) {
	requestDef := GenReqDefForCreateMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessagesResponse), nil
	}
}

//创建授权
func (c *OsmClient) CreatePrivileges(request *model.CreatePrivilegesRequest) (*model.CreatePrivilegesResponse, error) {
	requestDef := GenReqDefForCreatePrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivilegesResponse), nil
	}
}

//创建关联，一个工单最多支持3个关联
func (c *OsmClient) CreateRelations(request *model.CreateRelationsRequest) (*model.CreateRelationsResponse, error) {
	requestDef := GenReqDefForCreateRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRelationsResponse), nil
	}
}

//提交评分
func (c *OsmClient) CreateScores(request *model.CreateScoresRequest) (*model.CreateScoresResponse, error) {
	requestDef := GenReqDefForCreateScores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScoresResponse), nil
	}
}

//删除附件
func (c *OsmClient) DeleteAccessories(request *model.DeleteAccessoriesRequest) (*model.DeleteAccessoriesResponse, error) {
	requestDef := GenReqDefForDeleteAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessoriesResponse), nil
	}
}

//删除指定工单关联标签接口
func (c *OsmClient) DeleteCaseLabels(request *model.DeleteCaseLabelsRequest) (*model.DeleteCaseLabelsResponse, error) {
	requestDef := GenReqDefForDeleteCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCaseLabelsResponse), nil
	}
}

//删除标签，同时会删除工单与标签关联关系
func (c *OsmClient) DeleteLabels(request *model.DeleteLabelsRequest) (*model.DeleteLabelsResponse, error) {
	requestDef := GenReqDefForDeleteLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLabelsResponse), nil
	}
}

//删除关联
func (c *OsmClient) DeleteRelation(request *model.DeleteRelationRequest) (*model.DeleteRelationResponse, error) {
	requestDef := GenReqDefForDeleteRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRelationResponse), nil
	}
}

//下载附件
func (c *OsmClient) DownloadAccessories(request *model.DownloadAccessoriesRequest) (*model.DownloadAccessoriesResponse, error) {
	requestDef := GenReqDefForDownloadAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAccessoriesResponse), nil
	}
}

//工单导出
func (c *OsmClient) DownloadCases(request *model.DownloadCasesRequest) (*model.DownloadCasesResponse, error) {
	requestDef := GenReqDefForDownloadCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadCasesResponse), nil
	}
}

//返回图片内容，用于页面直接展示
func (c *OsmClient) DownloadImages(request *model.DownloadImagesRequest) (*model.DownloadImagesResponse, error) {
	requestDef := GenReqDefForDownloadImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadImagesResponse), nil
	}
}

//查询委托
func (c *OsmClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesResponse), nil
	}
}

//查询国家码，用于提交工单页面填写联系方式使用
func (c *OsmClient) ListAreaCodes(request *model.ListAreaCodesRequest) (*model.ListAreaCodesResponse, error) {
	requestDef := GenReqDefForListAreaCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAreaCodesResponse), nil
	}
}

//查询授权列表
func (c *OsmClient) ListAuthorizations(request *model.ListAuthorizationsRequest) (*model.ListAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizationsResponse), nil
	}
}

//查询工单类目列表
func (c *OsmClient) ListCaseCategories(request *model.ListCaseCategoriesRequest) (*model.ListCaseCategoriesResponse, error) {
	requestDef := GenReqDefForListCaseCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCategoriesResponse), nil
	}
}

//查询工单抄送邮箱
func (c *OsmClient) ListCaseCcEmails(request *model.ListCaseCcEmailsRequest) (*model.ListCaseCcEmailsResponse, error) {
	requestDef := GenReqDefForListCaseCcEmails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCcEmailsResponse), nil
	}
}

//统计各状态工单数量
func (c *OsmClient) ListCaseCounts(request *model.ListCaseCountsRequest) (*model.ListCaseCountsResponse, error) {
	requestDef := GenReqDefForListCaseCounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseCountsResponse), nil
	}
}

//查询工单关联标签接口
func (c *OsmClient) ListCaseLabels(request *model.ListCaseLabelsRequest) (*model.ListCaseLabelsResponse, error) {
	requestDef := GenReqDefForListCaseLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseLabelsResponse), nil
	}
}

//查询工单限制，比如抄送邮箱个数等
func (c *OsmClient) ListCaseLimits(request *model.ListCaseLimitsRequest) (*model.ListCaseLimitsResponse, error) {
	requestDef := GenReqDefForListCaseLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseLimitsResponse), nil
	}
}

//查询工单操作日志
func (c *OsmClient) ListCaseOperateLogs(request *model.ListCaseOperateLogsRequest) (*model.ListCaseOperateLogsResponse, error) {
	requestDef := GenReqDefForListCaseOperateLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseOperateLogsResponse), nil
	}
}

//查询工单配额
func (c *OsmClient) ListCaseQuotas(request *model.ListCaseQuotasRequest) (*model.ListCaseQuotasResponse, error) {
	requestDef := GenReqDefForListCaseQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseQuotasResponse), nil
	}
}

//查询问题类型对应模板
func (c *OsmClient) ListCaseTemplates(request *model.ListCaseTemplatesRequest) (*model.ListCaseTemplatesResponse, error) {
	requestDef := GenReqDefForListCaseTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCaseTemplatesResponse), nil
	}
}

//查询工单列表接口
func (c *OsmClient) ListCases(request *model.ListCasesRequest) (*model.ListCasesResponse, error) {
	requestDef := GenReqDefForListCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCasesResponse), nil
	}
}

//提单时，根据不同的产品或者问题类型，会存在不同的一些附加参数填写
func (c *OsmClient) ListExtendsParams(request *model.ListExtendsParamsRequest) (*model.ListExtendsParamsResponse, error) {
	requestDef := GenReqDefForListExtendsParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExtendsParamsResponse), nil
	}
}

//查询已验证的列表
func (c *OsmClient) ListHasVerifiedContacts(request *model.ListHasVerifiedContactsRequest) (*model.ListHasVerifiedContactsResponse, error) {
	requestDef := GenReqDefForListHasVerifiedContacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHasVerifiedContactsResponse), nil
	}
}

//查询堡垒机历史操作记录
func (c *OsmClient) ListHistoryOperateLogs(request *model.ListHistoryOperateLogsRequest) (*model.ListHistoryOperateLogsResponse, error) {
	requestDef := GenReqDefForListHistoryOperateLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOperateLogsResponse), nil
	}
}

//查询堡垒机历史会话列
func (c *OsmClient) ListHistorySessions(request *model.ListHistorySessionsRequest) (*model.ListHistorySessionsResponse, error) {
	requestDef := GenReqDefForListHistorySessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistorySessionsResponse), nil
	}
}

//查询标签
func (c *OsmClient) ListLabels(request *model.ListLabelsRequest) (*model.ListLabelsResponse, error) {
	requestDef := GenReqDefForListLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLabelsResponse), nil
	}
}

//查询留言
func (c *OsmClient) ListMessages(request *model.ListMessagesRequest) (*model.ListMessagesResponse, error) {
	requestDef := GenReqDefForListMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessagesResponse), nil
	}
}

//查询更多留言
func (c *OsmClient) ListMoreInstantMessages(request *model.ListMoreInstantMessagesRequest) (*model.ListMoreInstantMessagesResponse, error) {
	requestDef := GenReqDefForListMoreInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMoreInstantMessagesResponse), nil
	}
}

//轮询查询即时消息接口
func (c *OsmClient) ListNewInstantMessages(request *model.ListNewInstantMessagesRequest) (*model.ListNewInstantMessagesResponse, error) {
	requestDef := GenReqDefForListNewInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNewInstantMessagesResponse), nil
	}
}

//查询工单权限
func (c *OsmClient) ListPrivileges(request *model.ListPrivilegesRequest) (*model.ListPrivilegesResponse, error) {
	requestDef := GenReqDefForListPrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivilegesResponse), nil
	}
}

//提交工单时，选择产品类型之后选择对应的问题列表
func (c *OsmClient) ListProblemTypes(request *model.ListProblemTypesRequest) (*model.ListProblemTypesResponse, error) {
	requestDef := GenReqDefForListProblemTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProblemTypesResponse), nil
	}
}

//查询产品类型列表
func (c *OsmClient) ListProductCategories(request *model.ListProductCategoriesRequest) (*model.ListProductCategoriesResponse, error) {
	requestDef := GenReqDefForListProductCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductCategoriesResponse), nil
	}
}

//查询区域列表
func (c *OsmClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

//查询工单的关联，返回关联工单的简要信息
func (c *OsmClient) ListRelation(request *model.ListRelationRequest) (*model.ListRelationResponse, error) {
	requestDef := GenReqDefForListRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelationResponse), nil
	}
}

//工单满意度分类列表
func (c *OsmClient) ListSatisfactionDimensions(request *model.ListSatisfactionDimensionsRequest) (*model.ListSatisfactionDimensionsResponse, error) {
	requestDef := GenReqDefForListSatisfactionDimensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSatisfactionDimensionsResponse), nil
	}
}

//查询问题严重性列表
func (c *OsmClient) ListSeverities(request *model.ListSeveritiesRequest) (*model.ListSeveritiesResponse, error) {
	requestDef := GenReqDefForListSeverities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSeveritiesResponse), nil
	}
}

//查询子用户信息
func (c *OsmClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

//查询堡垒机文件传输记录
func (c *OsmClient) ListTransportHistories(request *model.ListTransportHistoriesRequest) (*model.ListTransportHistoriesResponse, error) {
	requestDef := GenReqDefForListTransportHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransportHistoriesResponse), nil
	}
}

//查询未读消息
func (c *OsmClient) ListUnreadNewInstantMessages(request *model.ListUnreadNewInstantMessagesRequest) (*model.ListUnreadNewInstantMessagesResponse, error) {
	requestDef := GenReqDefForListUnreadNewInstantMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnreadNewInstantMessagesResponse), nil
	}
}

//获取验证码
func (c *OsmClient) SendVerifyCodes(request *model.SendVerifyCodesRequest) (*model.SendVerifyCodesResponse, error) {
	requestDef := GenReqDefForSendVerifyCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerifyCodesResponse), nil
	}
}

//查询附件的一下限制，比如大小，数量，文件类型
func (c *OsmClient) ShowAccessoryLimits(request *model.ShowAccessoryLimitsRequest) (*model.ShowAccessoryLimitsResponse, error) {
	requestDef := GenReqDefForShowAccessoryLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessoryLimitsResponse), nil
	}
}

//查询授权详情
func (c *OsmClient) ShowAuthorizationDetail(request *model.ShowAuthorizationDetailRequest) (*model.ShowAuthorizationDetailResponse, error) {
	requestDef := GenReqDefForShowAuthorizationDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthorizationDetailResponse), nil
	}
}

//查询工单详情
func (c *OsmClient) ShowCaseDetail(request *model.ShowCaseDetailRequest) (*model.ShowCaseDetailResponse, error) {
	requestDef := GenReqDefForShowCaseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCaseDetailResponse), nil
	}
}

//查询某个工单状态
func (c *OsmClient) ShowCaseStatus(request *model.ShowCaseStatusRequest) (*model.ShowCaseStatusResponse, error) {
	requestDef := GenReqDefForShowCaseStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCaseStatusResponse), nil
	}
}

//查询伙伴工单权限
func (c *OsmClient) ShowPartnersCasesPrivilege(request *model.ShowPartnersCasesPrivilegeRequest) (*model.ShowPartnersCasesPrivilegeResponse, error) {
	requestDef := GenReqDefForShowPartnersCasesPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartnersCasesPrivilegeResponse), nil
	}
}

//查询关联伙伴服务信息
func (c *OsmClient) ShowPartnersServiceInfo(request *model.ShowPartnersServiceInfoRequest) (*model.ShowPartnersServiceInfoResponse, error) {
	requestDef := GenReqDefForShowPartnersServiceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartnersServiceInfoResponse), nil
	}
}

//拒绝|撤销授权
func (c *OsmClient) UpdateAuthorizations(request *model.UpdateAuthorizationsRequest) (*model.UpdateAuthorizationsResponse, error) {
	requestDef := GenReqDefForUpdateAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthorizationsResponse), nil
	}
}

//工单操作
func (c *OsmClient) UpdateCases(request *model.UpdateCasesRequest) (*model.UpdateCasesResponse, error) {
	requestDef := GenReqDefForUpdateCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCasesResponse), nil
	}
}

//修改标签
func (c *OsmClient) UpdateLabels(request *model.UpdateLabelsRequest) (*model.UpdateLabelsResponse, error) {
	requestDef := GenReqDefForUpdateLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLabelsResponse), nil
	}
}

//设置消息已读
func (c *OsmClient) UpdateNewInstantMessagesRead(request *model.UpdateNewInstantMessagesReadRequest) (*model.UpdateNewInstantMessagesReadResponse, error) {
	requestDef := GenReqDefForUpdateNewInstantMessagesRead()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNewInstantMessagesReadResponse), nil
	}
}

//上传附件给SDK使用
func (c *OsmClient) UploadJsonAccessories(request *model.UploadJsonAccessoriesRequest) (*model.UploadJsonAccessoriesResponse, error) {
	requestDef := GenReqDefForUploadJsonAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJsonAccessoriesResponse), nil
	}
}
