package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleResponse Response Object
type UpdateRuleResponse struct {

	// 创建状态
	Status *bool `json:"status,omitempty"`

	// 规则ID
	RuleId         *string `json:"rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleResponse", string(data)}, " ")
}
