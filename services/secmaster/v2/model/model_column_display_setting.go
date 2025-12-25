package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnDisplaySetting 表格列展示设置
type ColumnDisplaySetting struct {

	// 映射是否必填
	MappingRequired *bool `json:"mapping_required,omitempty"`

	// 组顺序编号
	GroupSequenceNumber *int32 `json:"group_sequence_number,omitempty"`

	// 组内顺序编号
	IntraGroupSequenceNumber *int32 `json:"intra_group_sequence_number,omitempty"`

	// 值类型
	ValueType *string `json:"value_type,omitempty"`

	// 合法值
	ValueQualified *string `json:"value_qualified,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 显示描述
	DisplayDescription *string `json:"display_description,omitempty"`

	// 组名
	GroupName *string `json:"group_name,omitempty"`
}

func (o ColumnDisplaySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnDisplaySetting struct{}"
	}

	return strings.Join([]string{"ColumnDisplaySetting", string(data)}, " ")
}
