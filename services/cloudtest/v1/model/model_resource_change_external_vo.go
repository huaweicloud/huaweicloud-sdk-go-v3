package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceChangeExternalVo 历史记录字段变更列表
type ResourceChangeExternalVo struct {

	// 变更字段
	FieldName *string `json:"field_name,omitempty"`

	// 测试用例自定义字段类型
	CustomFieldType *string `json:"custom_field_type,omitempty"`

	OldChangeInfo *ElementResourceChangeExternalVo `json:"old_change_info,omitempty"`

	NewChangeInfo *ElementResourceChangeExternalVo `json:"new_change_info,omitempty"`
}

func (o ResourceChangeExternalVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceChangeExternalVo struct{}"
	}

	return strings.Join([]string{"ResourceChangeExternalVo", string(data)}, " ")
}
