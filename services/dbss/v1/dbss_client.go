package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
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

// SwitchAgent 开启关闭Agent
//
// 用于开启和关闭agent的功能，当开启后，开始抓取用户的访问信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DbssClient) SwitchAgent(request *model.SwitchAgentRequest) (*model.SwitchAgentResponse, error) {
	requestDef := GenReqDefForSwitchAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAgentResponse), nil
	}
}

// SwitchAgentInvoker 开启关闭Agent
func (c *DbssClient) SwitchAgentInvoker(request *model.SwitchAgentRequest) *SwitchAgentInvoker {
	requestDef := GenReqDefForSwitchAgent()
	return &SwitchAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchRiskRule 开启关闭风险规则
//
// 开启关闭风险规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DbssClient) SwitchRiskRule(request *model.SwitchRiskRuleRequest) (*model.SwitchRiskRuleResponse, error) {
	requestDef := GenReqDefForSwitchRiskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRiskRuleResponse), nil
	}
}

// SwitchRiskRuleInvoker 开启关闭风险规则
func (c *DbssClient) SwitchRiskRuleInvoker(request *model.SwitchRiskRuleRequest) *SwitchRiskRuleInvoker {
	requestDef := GenReqDefForSwitchRiskRule()
	return &SwitchRiskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
