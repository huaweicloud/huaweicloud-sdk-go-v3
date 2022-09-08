package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueAttrHistoryRecord struct {

	// 操作的字段
	FieldKey *string `json:"field_key,omitempty"`

	// 操作字段的含义
	FieldName *string `json:"field_name,omitempty"`

	// 历史记录id
	Id *int64 `json:"id,omitempty"`

	// 工作项id
	IssueId *int64 `json:"issue_id,omitempty"`

	// 变更后的值,json字符串
	NewValue *string `json:"new_value,omitempty"`

	// 变更前的值,json字符串
	OldValue *string `json:"old_value,omitempty"`

	// 变更的时间
	OperatedTime *int64 `json:"operated_time,omitempty"`

	// 操作类型,新建，修改，删除
	Operation *string `json:"operation,omitempty"`

	Operator *IssueUser `json:"operator,omitempty"`

	// 变更的属性
	Property *string `json:"property,omitempty"`
}

func (o IssueAttrHistoryRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueAttrHistoryRecord struct{}"
	}

	return strings.Join([]string{"IssueAttrHistoryRecord", string(data)}, " ")
}
