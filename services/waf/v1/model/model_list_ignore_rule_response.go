package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIgnoreRuleResponse struct {

	// 该策略下全局白名单规则数量
	Total *int32 `json:"total,omitempty"`

	// 全局白名单规则信息数组
	Items          *[]IgnoreRuleBody `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"ListIgnoreRuleResponse", string(data)}, " ")
}
