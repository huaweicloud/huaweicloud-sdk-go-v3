package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionParameter 用户可配置参数
type ExtensionParameter struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数显示标签
	Label *string `json:"label,omitempty"`

	Validation *ExtensionParameterValidation `json:"validation,omitempty"`

	// 默认值
	DefaultValue *string `json:"defaultValue,omitempty"`

	// 帮助文档(markdown格式)。
	HelpMarkdown *string `json:"helpMarkdown,omitempty"`

	DisplaySettings *ExtensionParameterDisplaySettings `json:"displaySettings,omitempty"`
}

func (o ExtensionParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionParameter struct{}"
	}

	return strings.Join([]string{"ExtensionParameter", string(data)}, " ")
}
