package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateForwardRuleRequest Request Object
type UpdateForwardRuleRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 规则id
	RuleId string `json:"rule_id"`

	// 高防ip
	Ip string `json:"ip"`

	Body *UpdateForwardRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateForwardRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateForwardRuleRequest", string(data)}, " ")
}
