package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作详情
type IterationHistoryDetails struct {
	// 变更的字段

	OperateFieldName *string `json:"operate_field_name,omitempty"`
	// 操作后的值

	NewValue *string `json:"new_value,omitempty"`
	// 操作前的值

	OldValue *string `json:"old_value,omitempty"`
}

func (o IterationHistoryDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IterationHistoryDetails struct{}"
	}

	return strings.Join([]string{"IterationHistoryDetails", string(data)}, " ")
}
