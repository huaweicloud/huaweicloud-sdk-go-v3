package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleResponse Response Object
type ListRuleResponse struct {

	// 静态规则列表
	Data *[]Rule `json:"data,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleResponse struct{}"
	}

	return strings.Join([]string{"ListRuleResponse", string(data)}, " ")
}
