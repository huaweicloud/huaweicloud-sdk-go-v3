package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameterUnit 集群使用的参数配置值。
type ConfigurationParameterUnit struct {

	// 参数类型，包括：cn、dn。
	Type string `json:"type"`

	// 参数值。
	Value string `json:"value"`

	// 参数默认值。
	DefaultValue string `json:"default_value"`
}

func (o ConfigurationParameterUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterUnit struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterUnit", string(data)}, " ")
}
