package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationParameter struct {

	// 参数名称。
	Name string `json:"name"`

	// 参数值。
	Value string `json:"value"`

	// 是否需要重启。  “false”表示否。 “true”表示是。
	RestartRequired bool `json:"restart_required"`

	// 是否只读。  “false”表示否。 “true”表示是。
	Readonly bool `json:"readonly"`

	// 参数值范围，如integer取值0-1、boolean取值true或者false等。
	ValueRange *string `json:"value_range,omitempty"`

	// 参数类型，取值为“string”、“integer”、“boolean”、“list”或“float”之一。
	Type string `json:"type"`

	// 参数描述。
	Description string `json:"description"`
}

func (o ConfigurationParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameter struct{}"
	}

	return strings.Join([]string{"ConfigurationParameter", string(data)}, " ")
}
