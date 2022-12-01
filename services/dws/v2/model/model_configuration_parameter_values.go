package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群参数配置列表信息
type ConfigurationParameterValues struct {

	// 集群参数配置列表
	Configurations []ConfigurationParameterValue `json:"configurations"`
}

func (o ConfigurationParameterValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterValues struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterValues", string(data)}, " ")
}
