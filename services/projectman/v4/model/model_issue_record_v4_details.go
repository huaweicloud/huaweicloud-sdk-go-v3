package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueRecordV4Details struct {
	// 操作属性

	Property *string `json:"property,omitempty"`
	// 上次的记录

	OldValue *string `json:"old_value,omitempty"`
	// 当前值

	NewValue *string `json:"new_value,omitempty"`
	// 操作

	Operation *string `json:"operation,omitempty"`
	// 操作记录的id

	Id *int32 `json:"id,omitempty"`
	// 操作的字段

	Name *string `json:"name,omitempty"`
}

func (o IssueRecordV4Details) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRecordV4Details struct{}"
	}

	return strings.Join([]string{"IssueRecordV4Details", string(data)}, " ")
}
