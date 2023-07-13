package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDisplayRenderingAccelerationOptions 渲染加速控制选项。
type PoliciesDisplayRenderingAccelerationOptions struct {

	// 视频加速增强配置。取值为： false：表示关闭。 true：表示开启。
	VideoAccelerationEnhancementEnable *bool `json:"video_acceleration_enhancement_enable,omitempty"`

	// 是否开启视频场景优化。取值为： false：表示关闭。 true：表示开启。
	VideoOptimizationEnable *bool `json:"video_optimization_enable,omitempty"`

	// 是否开启GPU色彩优化。取值为： false：表示关闭。 true：表示开启。
	GpuColorOptimizationEnable *bool `json:"gpu_color_optimization_enable,omitempty"`
}

func (o PoliciesDisplayRenderingAccelerationOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayRenderingAccelerationOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayRenderingAccelerationOptions", string(data)}, " ")
}
