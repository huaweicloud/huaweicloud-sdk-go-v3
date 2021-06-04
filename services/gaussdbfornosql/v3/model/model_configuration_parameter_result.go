package model

import (
	"encoding/json"

	"strings"
)

type ConfigurationParameterResult struct {
	// 参数名称。

	Name string `json:"name"`
	// 参数值。

	Value string `json:"value"`
	// 是否需要重启。 - 取值为“false”，不需要重启。 - 取值为“true”，需要重启。

	RestartRequired bool `json:"restart_required"`
	// 是否只读。 - 取值为“false”，非只读参数。 - 取值为“true”，只读参数。

	Readonly bool `json:"readonly"`
	// 参数值范围，如integer取值0-1、boolean取值true|false等。

	ValueRange string `json:"value_range"`
	// 参数类型，取值为“string”、“integer”、“boolean”、“list”或“float”之一。

	Type string `json:"type"`
	// 参数描述。

	Description string `json:"description"`
}

func (o ConfigurationParameterResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfigurationParameterResult struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterResult", string(data)}, " ")
}
