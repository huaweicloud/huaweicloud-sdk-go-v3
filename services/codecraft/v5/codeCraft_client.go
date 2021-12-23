package v5

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/codecraft/v5/model"
)

type CodeCraftClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCodeCraftClient(hcClient *http_client.HcHttpClient) *CodeCraftClient {
	return &CodeCraftClient{HcClient: hcClient}
}

func CodeCraftClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//针对在第三方提交作品的场景：第三方服务对作品完成判分后，调用该接口将作品信息及作品得分返回给大赛平台
func (c *CodeCraftClient) CreateCompetitionScore(request *model.CreateCompetitionScoreRequest) (*model.CreateCompetitionScoreResponse, error) {
	requestDef := GenReqDefForCreateCompetitionScore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCompetitionScoreResponse), nil
	}
}

//第三方服务获取某个大赛某个阶段中一段时间内提交的作品信息。其中以请求参数read_time作为结束时间，定义向前一天或一小时内的时间作为查询范围
func (c *CodeCraftClient) ListCompetitionWorks(request *model.ListCompetitionWorksRequest) (*model.ListCompetitionWorksResponse, error) {
	requestDef := GenReqDefForListCompetitionWorks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompetitionWorksResponse), nil
	}
}

//第三方服务验证用户是否在大赛平台报名、是否组建团队、是否可以提交作品。如果已经报名但是未组建团队，则创建一个虚拟团队，设置为允许提交作品。如果已经组建团队则根据大赛报名截止时间判断是否可以提交作品。返回团队ID、是否可以提交作品
func (c *CodeCraftClient) RegisterCompetitionInfo(request *model.RegisterCompetitionInfoRequest) (*model.RegisterCompetitionInfoResponse, error) {
	requestDef := GenReqDefForRegisterCompetitionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterCompetitionInfoResponse), nil
	}
}

//针对在大赛平台提交作品的场景：第三方服务对作品完成判分后，根据作品ID调用该接口将作品分数、作品状态等信息返回给大赛平台
func (c *CodeCraftClient) UpdateCompetitionScore(request *model.UpdateCompetitionScoreRequest) (*model.UpdateCompetitionScoreResponse, error) {
	requestDef := GenReqDefForUpdateCompetitionScore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCompetitionScoreResponse), nil
	}
}
