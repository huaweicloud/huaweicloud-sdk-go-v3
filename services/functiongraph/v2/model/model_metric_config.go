package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricConfig 流量配置
type MetricConfig struct {

	// 流量配置名称
	Name *string `json:"name,omitempty"`

	// 流量配置类型
	Type *string `json:"type,omitempty"`

	// 流量阈值
	Threshold *int32 `json:"threshold,omitempty"`

	// 流量最小值
	Min *int32 `json:"min,omitempty"`
}

func (o MetricConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricConfig struct{}"
	}

	return strings.Join([]string{"MetricConfig", string(data)}, " ")
}
