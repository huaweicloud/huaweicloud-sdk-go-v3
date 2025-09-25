package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentPayInfoV2 struct {

	// |参数名称：订单代理支付人的客户账号ID。| |参数的约束及描述：订单代理支付人的客户账号ID|
	PayingAgentId *string `json:"paying_agent_id,omitempty"`

	// |参数名称：是否代理支付。| |参数的约束及描述：是否代理支付。true：代理支付，不为ture时为非代理支付订单|
	IsAgentPay *bool `json:"is_agent_pay,omitempty"`
}

func (o AgentPayInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentPayInfoV2 struct{}"
	}

	return strings.Join([]string{"AgentPayInfoV2", string(data)}, " ")
}
