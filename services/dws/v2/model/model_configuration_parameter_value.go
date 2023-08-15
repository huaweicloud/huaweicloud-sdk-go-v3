package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameterValue 集群参数配置详情。
type ConfigurationParameterValue struct {

	// 参数类型，包括：cn、dn。
	Type string `json:"type"`

	// 参数名称。
	Name string `json:"name"`

	// 参数值。
	Value string `json:"value"`
}

func (o ConfigurationParameterValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterValue struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterValue", string(data)}, " ")
}
