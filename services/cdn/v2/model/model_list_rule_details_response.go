package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleDetailsResponse Response Object
type ListRuleDetailsResponse struct {

	// 规则列表
	Rules          *[]RuleResponse `json:"rules,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRuleDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListRuleDetailsResponse", string(data)}, " ")
}
