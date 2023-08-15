package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleConfig struct {

	// 规则配置ID
	Id *int32 `json:"id,omitempty"`

	// 规则ID
	RuleId *int32 `json:"rule_id,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 当前
	OptionValue *string `json:"option_value,omitempty"`

	// 当前规则配置项key
	OptionKey *string `json:"option_key,omitempty"`

	// 当前规则配置项名称
	OptionName *string `json:"option_name,omitempty"`

	// 规则集id
	TemplateId *string `json:"template_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o RuleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleConfig struct{}"
	}

	return strings.Join([]string{"RuleConfig", string(data)}, " ")
}
