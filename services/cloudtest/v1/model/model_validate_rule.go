package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateRule struct {

	// 依赖信息列表
	DependentInfo *[]ConditionInstance `json:"dependent_info,omitempty"`

	// 是否启用的标识符
	Enabled *bool `json:"enabled,omitempty"`

	// 配置依赖的标识符
	IsConfigDep *bool `json:"is_configDep,omitempty"`

	// 是否依赖的标识符
	IsDependent *bool `json:"is_dependent,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 结果
	Result *string `json:"result,omitempty"`

	// 特殊字符
	SpecialChar *string `json:"special_char,omitempty"`

	// 特殊字符的有效性
	SpecialCharValid *string `json:"special_char_valid,omitempty"`

	// 示例
	XExample *string `json:"x_example,omitempty"`
}

func (o ValidateRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRule struct{}"
	}

	return strings.Join([]string{"ValidateRule", string(data)}, " ")
}
