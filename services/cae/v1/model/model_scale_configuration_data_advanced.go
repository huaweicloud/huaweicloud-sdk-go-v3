package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleConfigurationDataAdvanced 伸缩策略高级配置。
type ScaleConfigurationDataAdvanced struct {

	// 缩容冷却时间窗。
	ScaledownStabilizationSeconds *int32 `json:"scaledown_stabilization_seconds,omitempty"`

	// 缩容步长。
	ScaledownRate *int32 `json:"scaledown_rate,omitempty"`

	// 扩容冷却时间窗。
	ScaleupStabilizationSeconds *int32 `json:"scaleup_stabilization_seconds,omitempty"`

	// 扩容步长。
	ScaleupRate *int32 `json:"scaleup_rate,omitempty"`

	// 是否禁用自动缩容。
	DisableScaledown *bool `json:"disable_scaledown,omitempty"`
}

func (o ScaleConfigurationDataAdvanced) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleConfigurationDataAdvanced struct{}"
	}

	return strings.Join([]string{"ScaleConfigurationDataAdvanced", string(data)}, " ")
}
