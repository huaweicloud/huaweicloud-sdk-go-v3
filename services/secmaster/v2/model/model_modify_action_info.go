package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyActionInfo Action info
type ModifyActionInfo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 类型，默认AOP_WORKFLOW.
	ActionType *string `json:"action_type,omitempty"`

	// 剧本动作ID
	ActionId *string `json:"action_id,omitempty"`

	// 排序方式
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o ModifyActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyActionInfo struct{}"
	}

	return strings.Join([]string{"ModifyActionInfo", string(data)}, " ")
}
