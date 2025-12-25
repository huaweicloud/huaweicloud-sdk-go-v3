package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DpeMappingFieldDetail struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 映射id
	MappingId *string `json:"mapping_id,omitempty"`

	// 映射id
	MapperId *string `json:"mapper_id,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 目标字段
	TargetKey *string `json:"target_key,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o DpeMappingFieldDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeMappingFieldDetail struct{}"
	}

	return strings.Join([]string{"DpeMappingFieldDetail", string(data)}, " ")
}
