package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImmutableRulesResponse Response Object
type ListImmutableRulesResponse struct {

	// 策略列表
	ImmutableRules *[]ImmutableRule `json:"immutable_rules,omitempty"`

	// 策略总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImmutableRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImmutableRulesResponse struct{}"
	}

	return strings.Join([]string{"ListImmutableRulesResponse", string(data)}, " ")
}
