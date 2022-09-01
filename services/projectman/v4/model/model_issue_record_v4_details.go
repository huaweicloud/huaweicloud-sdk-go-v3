package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueRecordV4Details struct {

	// 操作属性
	Property *string `json:"property,omitempty" xml:"property"`

	// 上次的记录
	OldValue *string `json:"old_value,omitempty" xml:"old_value"`

	// 当前值
	NewValue *string `json:"new_value,omitempty" xml:"new_value"`

	// 操作
	Operation *string `json:"operation,omitempty" xml:"operation"`

	// 操作记录的id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 操作的字段
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o IssueRecordV4Details) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRecordV4Details struct{}"
	}

	return strings.Join([]string{"IssueRecordV4Details", string(data)}, " ")
}
