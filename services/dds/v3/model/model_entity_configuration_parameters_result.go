package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EntityConfigurationParametersResult struct {

	// 参数名称。
	Name string `json:"name" xml:"name"`

	// 参数值。
	Value string `json:"value" xml:"value"`

	// 参数值范围。示例：Integer类型取值范围为0~1、Boolean类型取值为“true”或“false”。
	ValueRange string `json:"value_range" xml:"value_range"`

	// 参数是否需要重启。 - 取值为“true”，需要重启。 - 取值为“false”，不需要重启。
	RestartRequired bool `json:"restart_required" xml:"restart_required"`

	// 是否只读。 - 取值为“false”，非只读参数。 - 取值为“true”，只读参数。
	Readonly bool `json:"readonly" xml:"readonly"`

	// 参数类型，取值为“integer”，“string”，“boolean”，“float”或“list”。
	Type string `json:"type" xml:"type"`

	// 参数描述。
	Description string `json:"description" xml:"description"`
}

func (o EntityConfigurationParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityConfigurationParametersResult struct{}"
	}

	return strings.Join([]string{"EntityConfigurationParametersResult", string(data)}, " ")
}
