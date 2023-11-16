package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceIpRuleRequest Request Object
type UpdateInstanceIpRuleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 单个 IP
	Ip string `json:"ip"`

	// 规则ID
	RuleId string `json:"rule_id"`

	Body *TransferRuleBody `json:"body,omitempty"`
}

func (o UpdateInstanceIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceIpRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceIpRuleRequest", string(data)}, " ")
}
