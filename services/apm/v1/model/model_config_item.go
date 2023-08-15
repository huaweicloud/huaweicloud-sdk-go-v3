package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigItem struct {

	// 配置项名称。
	ConfigName *string `json:"config_name,omitempty"`

	// 配置项值。
	ConfigValue *string `json:"config_value,omitempty"`

	// 是否重写。
	ShouldOverride *bool `json:"should_override,omitempty"`
}

func (o ConfigItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigItem struct{}"
	}

	return strings.Join([]string{"ConfigItem", string(data)}, " ")
}
