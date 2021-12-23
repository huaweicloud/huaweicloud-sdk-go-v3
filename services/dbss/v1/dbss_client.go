package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type DbssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDbssClient(hcClient *http_client.HcHttpClient) *DbssClient {
	return &DbssClient{HcClient: hcClient}
}

func DbssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//用于开启和关闭agent的功能，当开启后，开始抓取用户的访问信息。
func (c *DbssClient) SwitchAgent(request *model.SwitchAgentRequest) (*model.SwitchAgentResponse, error) {
	requestDef := GenReqDefForSwitchAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAgentResponse), nil
	}
}

//开启关闭风险规则
func (c *DbssClient) SwitchRiskRule(request *model.SwitchRiskRuleRequest) (*model.SwitchRiskRuleResponse, error) {
	requestDef := GenReqDefForSwitchRiskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRiskRuleResponse), nil
	}
}
