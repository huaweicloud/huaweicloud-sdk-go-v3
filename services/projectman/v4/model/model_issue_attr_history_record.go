package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueAttrHistoryRecord struct {

	// 操作的字段
	FieldKey *string `json:"field_key,omitempty" xml:"field_key"`

	// 操作字段的含义
	FieldName *string `json:"field_name,omitempty" xml:"field_name"`

	// 历史记录id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 工作项id
	IssueId *int64 `json:"issue_id,omitempty" xml:"issue_id"`

	// 变更后的值,json字符串
	NewValue *string `json:"new_value,omitempty" xml:"new_value"`

	// 变更前的值,json字符串
	OldValue *string `json:"old_value,omitempty" xml:"old_value"`

	// 变更的时间
	OperatedTime *int64 `json:"operated_time,omitempty" xml:"operated_time"`

	// 操作类型,新建，修改，删除
	Operation *string `json:"operation,omitempty" xml:"operation"`

	Operator *IssueUser `json:"operator,omitempty" xml:"operator"`

	// 变更的属性
	Property *string `json:"property,omitempty" xml:"property"`
}

func (o IssueAttrHistoryRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueAttrHistoryRecord struct{}"
	}

	return strings.Join([]string{"IssueAttrHistoryRecord", string(data)}, " ")
}
