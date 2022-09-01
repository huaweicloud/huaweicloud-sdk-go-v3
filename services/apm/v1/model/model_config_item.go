package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigItem struct {

	// 配置项名称
	ConfigName *string `json:"config_name,omitempty" xml:"config_name"`

	// 配置项值
	ConfigValue *string `json:"config_value,omitempty" xml:"config_value"`

	// 是否重写
	ShouldOverride *bool `json:"should_override,omitempty" xml:"should_override"`
}

func (o ConfigItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigItem struct{}"
	}

	return strings.Join([]string{"ConfigItem", string(data)}, " ")
}
