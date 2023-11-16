package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleResponse Response Object
type CreateRuleResponse struct {

	// 创建状态
	Status *bool `json:"status,omitempty"`

	// 规则ID
	RuleId         *string `json:"rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleResponse", string(data)}, " ")
}
