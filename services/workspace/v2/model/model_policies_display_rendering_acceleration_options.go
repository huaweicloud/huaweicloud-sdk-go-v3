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

	// 视频识别阈值。取值范围为[0-500]。默认：10。
	VideoRecognitionThreshold *int32 `json:"video_recognition_threshold,omitempty"`

	// 帧率统计长度。取值范围为[2-100]。默认：4。
	FrameRateStatisticalLength *int32 `json:"frame_rate_statistical_length,omitempty"`

	// 图像质量阈值。取值范围为[0-100]。默认：0。
	ImageQualityThreshold *int32 `json:"image_quality_threshold,omitempty"`

	// 刷新率阈值。取值范围为[1-100]。默认：3。
	RefreshFrequencyThreshold *int32 `json:"refresh_frequency_threshold,omitempty"`

	// 退出视频区域阈值。取值范围为[0-100]。默认：8。
	ExitingVideoAreaThreshold *int32 `json:"exiting_video_area_threshold,omitempty"`

	// 识别为视频的最小宽。取值范围为[0-1280]。默认：191。
	MinVideoWidth *int32 `json:"min_video_width,omitempty"`

	// 识别为视频的最小高。取值范围为[0-1280]。默认：191。
	MinVideoLength *int32 `json:"min_video_length,omitempty"`

	// 单帧自然图像块占比阈值。取值范围为[0.000001-1]。默认：0.3。
	SingleFrameNaturalPercentage *string `json:"single_frame_natural_percentage,omitempty"`

	// 周期自然图像数目占比阈值。取值范围为[0-100]。默认：2。
	CyclicalNaturalImagesNumber *int32 `json:"cyclical_natural_images_number,omitempty"`

	// 非自然图面积占比阈值。取值范围为[0.000001-1]。默认：0.85。
	NonNaturalImagePercentage *string `json:"non_natural_image_percentage,omitempty"`

	// 非自然图数目占比阈值。取值范围为[0-100]。默认：25。
	NonNaturalImagesNumber *int32 `json:"non_natural_images_number,omitempty"`
}

func (o PoliciesDisplayRenderingAccelerationOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayRenderingAccelerationOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayRenderingAccelerationOptions", string(data)}, " ")
}
