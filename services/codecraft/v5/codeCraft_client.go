package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codecraft/v5/model"
)

type CodeCraftClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeCraftClient(hcClient *httpclient.HcHttpClient) *CodeCraftClient {
	return &CodeCraftClient{HcClient: hcClient}
}

func CodeCraftClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateCompetitionScore 登记第三方提交的作品信息（得分回调）
//
// 针对在第三方提交作品的场景：第三方服务对作品完成判分后，调用该接口将作品信息及作品得分返回给大赛平台
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeCraftClient) CreateCompetitionScore(request *model.CreateCompetitionScoreRequest) (*model.CreateCompetitionScoreResponse, error) {
	requestDef := GenReqDefForCreateCompetitionScore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompetitionScoreResponse), nil
	}
}

// CreateCompetitionScoreInvoker 登记第三方提交的作品信息（得分回调）
func (c *CodeCraftClient) CreateCompetitionScoreInvoker(request *model.CreateCompetitionScoreRequest) *CreateCompetitionScoreInvoker {
	requestDef := GenReqDefForCreateCompetitionScore()
	return &CreateCompetitionScoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCompetitionWorks 获取指定时间内选手提交的作品
//
// 第三方服务获取某个大赛某个阶段中一段时间内提交的作品信息。其中以请求参数read_time作为结束时间，定义向前一天或一小时内的时间作为查询范围
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeCraftClient) ListCompetitionWorks(request *model.ListCompetitionWorksRequest) (*model.ListCompetitionWorksResponse, error) {
	requestDef := GenReqDefForListCompetitionWorks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompetitionWorksResponse), nil
	}
}

// ListCompetitionWorksInvoker 获取指定时间内选手提交的作品
func (c *CodeCraftClient) ListCompetitionWorksInvoker(request *model.ListCompetitionWorksRequest) *ListCompetitionWorksInvoker {
	requestDef := GenReqDefForListCompetitionWorks()
	return &ListCompetitionWorksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterCompetitionInfo 验证用户报名信息和团队信息
//
// 第三方服务验证用户是否在大赛平台报名、是否组建团队、是否可以提交作品。如果已经报名但是未组建团队，则创建一个虚拟团队，设置为允许提交作品。如果已经组建团队则根据大赛报名截止时间判断是否可以提交作品。返回团队ID、是否可以提交作品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeCraftClient) RegisterCompetitionInfo(request *model.RegisterCompetitionInfoRequest) (*model.RegisterCompetitionInfoResponse, error) {
	requestDef := GenReqDefForRegisterCompetitionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterCompetitionInfoResponse), nil
	}
}

// RegisterCompetitionInfoInvoker 验证用户报名信息和团队信息
func (c *CodeCraftClient) RegisterCompetitionInfoInvoker(request *model.RegisterCompetitionInfoRequest) *RegisterCompetitionInfoInvoker {
	requestDef := GenReqDefForRegisterCompetitionInfo()
	return &RegisterCompetitionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCompetitionScore 修改平台提交的作品分数（得分回调）
//
// 针对在大赛平台提交作品的场景：第三方服务对作品完成判分后，根据作品ID调用该接口将作品分数、作品状态等信息返回给大赛平台
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeCraftClient) UpdateCompetitionScore(request *model.UpdateCompetitionScoreRequest) (*model.UpdateCompetitionScoreResponse, error) {
	requestDef := GenReqDefForUpdateCompetitionScore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCompetitionScoreResponse), nil
	}
}

// UpdateCompetitionScoreInvoker 修改平台提交的作品分数（得分回调）
func (c *CodeCraftClient) UpdateCompetitionScoreInvoker(request *model.UpdateCompetitionScoreRequest) *UpdateCompetitionScoreInvoker {
	requestDef := GenReqDefForUpdateCompetitionScore()
	return &UpdateCompetitionScoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
