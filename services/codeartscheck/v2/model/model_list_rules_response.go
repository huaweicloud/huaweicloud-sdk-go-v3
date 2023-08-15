package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRulesResponse Response Object
type ListRulesResponse struct {

	// 规则信息
	Info *[]RuleListItem `json:"info,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRulesResponse", string(data)}, " ")
}
