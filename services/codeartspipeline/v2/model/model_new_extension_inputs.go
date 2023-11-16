package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewExtensionInputs struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 标签
	Label *string `json:"label,omitempty"`

	// 说明
	Description *string `json:"description,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 必填
	Required *bool `json:"required,omitempty"`

	ExtendProp *ExtensionExtendProp `json:"extend_prop,omitempty"`

	Validation *ExtensionValidation `json:"validation,omitempty"`
}

func (o NewExtensionInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionInputs struct{}"
	}

	return strings.Join([]string{"NewExtensionInputs", string(data)}, " ")
}
