package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群使用的参数配置项详情。
type ConfigurationParameter struct {

	// 参数名称。
	Name string `json:"name"`

	// 参数值。
	Values []ConfigurationParameterUnit `json:"values"`

	// 参数单位。
	Unit string `json:"unit"`

	// 参数类型，包括boolean、string、integer、float、list。
	Type string `json:"type"`

	// 是否只读。
	Readonly bool `json:"readonly"`

	// 参数值范围。
	ValueRange string `json:"value_range"`

	// 是否需要重启。
	RestartRequired bool `json:"restart_required"`

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
