package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIgnorePolicyRulesResponse Response Object
type ListIgnorePolicyRulesResponse struct {

	// 该策略下全局白名单规则数量
	Total *int32 `json:"total,omitempty"`

	// 全局白名单规则信息数组
	Items          *[]IgnoreRuleBody `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListIgnorePolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIgnorePolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListIgnorePolicyRulesResponse", string(data)}, " ")
}
