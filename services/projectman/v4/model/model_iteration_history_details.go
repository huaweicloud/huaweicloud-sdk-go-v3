package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作详情
type IterationHistoryDetails struct {

	// 变更的字段
	OperateFieldName *string `json:"operate_field_name,omitempty" xml:"operate_field_name"`

	// 操作后的值
	NewValue *string `json:"new_value,omitempty" xml:"new_value"`

	// 操作前的值
	OldValue *string `json:"old_value,omitempty" xml:"old_value"`
}

func (o IterationHistoryDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IterationHistoryDetails struct{}"
	}

	return strings.Join([]string{"IterationHistoryDetails", string(data)}, " ")
}
