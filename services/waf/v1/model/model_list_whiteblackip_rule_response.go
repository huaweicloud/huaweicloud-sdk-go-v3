package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWhiteblackipRuleResponse struct {
	// 规则条数

	Total *int32 `json:"total,omitempty"`
	// 规则

	Items          *[]WhiteBlackIpResponseBody `json:"items,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWhiteblackipRuleResponse", string(data)}, " ")
}
