package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivacyRuleResponse Response Object
type ListPrivacyRuleResponse struct {

	// 规则条数
	Total *int32 `json:"total,omitempty"`

	// 规则详情数组
	Items          *[]PrivacyResponseBody `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"ListPrivacyRuleResponse", string(data)}, " ")
}
