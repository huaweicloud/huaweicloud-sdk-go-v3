package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRulesResponse struct {
	// 规则配置

	Rules *[]RuleResponse `json:"rules,omitempty"`
	// 满足条件的规则个数

	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRulesResponse", string(data)}, " ")
}
