package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWhiteblackipRuleResponse struct {

	// 黑白名单规则条数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 黑白名单规则列表信息
	Items          *[]WhiteBlackIpResponseBody `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWhiteblackipRuleResponse", string(data)}, " ")
}
