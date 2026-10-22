package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiGuardRuleResponse Response Object
type ListAiGuardRuleResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 规则列表
	Items          *[]AiGuardRuleItem `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"ListAiGuardRuleResponse", string(data)}, " ")
}
