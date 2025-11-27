package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleResponse Response Object
type ListRuleResponse struct {

	// 权限策略列表
	Items *[]Rule `json:"items,omitempty"`

	// 所有页的结果的总数
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
