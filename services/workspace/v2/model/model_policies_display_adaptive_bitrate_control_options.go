package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDisplayAdaptiveBitrateControlOptions 自适应流控控制选项。
type PoliciesDisplayAdaptiveBitrateControlOptions struct {

	// 峰值带宽抑制：取值为： false：表示关闭。 true：表示开启。
	PeakBandwidthSuppressionEnable *bool `json:"peak_bandwidth_suppression_enable,omitempty"`

	// 网络平均期望延时。取值范围为[1-2000]。默认：160。
	ExpectedAverageNetworkLatency *int32 `json:"expected_average_network_latency,omitempty"`

	// 网络延时阈值1（ms）。取值范围为[1-2000]。默认：160。
	NetworkLatencyThreshold1 *int32 `json:"network_latency_threshold1,omitempty"`

	// 网络延时阈值2（ms）。取值范围为[1-2000]。默认：300。
	NetworkLatencyThreshold2 *int32 `json:"network_latency_threshold2,omitempty"`

	// 最小动态帧率（fps）。取值范围为[1-60]。默认：17。
	MinDynamicFrameRate *int32 `json:"min_dynamic_frame_rate,omitempty"`

	// 最小动态帧率Lv1（fps）。取值范围为[1-60]。默认：17。
	MinDynamicFrameRateLv1 *int32 `json:"min_dynamic_frame_rate_lv1,omitempty"`

	// 最小动态帧率Lv2（fps）。取值范围为[1-60]。默认：10。
	MinDynamicFrameRateLv2 *int32 `json:"min_dynamic_frame_rate_lv2,omitempty"`

	// 帧率控制参数1。取值范围为[0-1000]。默认：20。
	RttThreshold *int32 `json:"rtt_threshold,omitempty"`

	// 帧率控制参数2。取值范围为[0-120]。默认：8。
	MinAddFramerate *int32 `json:"min_add_framerate,omitempty"`

	// 帧率控制参数3。取值范围为[0-120]。默认：20。
	MaxAddFramerate *int32 `json:"max_add_framerate,omitempty"`

	// 帧率控制参数4。取值范围为[0-120]。默认：25。
	SubFramerate *int32 `json:"sub_framerate,omitempty"`

	// 自适应带宽下限。取值范围为[100-20000]。默认：500。
	AdaptiveBandwidthLowerLimit *int32 `json:"adaptive_bandwidth_lower_limit,omitempty"`

	// 自适应有损压缩质量下限。取值范围为[1-100]。默认：60。
	AdaptiveCompressionQualityLowerLimit *int32 `json:"adaptive_compression_quality_lower_limit,omitempty"`

	// 自适应有损压缩质量上限。取值范围为[1-100]。默认：85。
	AdaptiveCompressionQualityUpperLimit *int32 `json:"adaptive_compression_quality_upper_limit,omitempty"`

	// 自适应有损压缩质量增限。取值范围为[1-100]。默认：5。
	AdaptiveCompressionQualityIncreaseLimit *int32 `json:"adaptive_compression_quality_increase_limit,omitempty"`

	// 自适应有损压缩质量减限。取值范围为[1-100]。默认：10。
	AdaptiveCompressionQualityDecreaseLimit *int32 `json:"adaptive_compression_quality_decrease_limit,omitempty"`

	// 自适应视频平均质量下限。取值范围为[5-50]。默认：15。
	AdaptiveAverageQualityLowerLimit *int32 `json:"adaptive_average_quality_lower_limit,omitempty"`

	// 自适应视频平均质量上限。取值范围为[5-50]。默认：25。
	AdaptiveAverageQualityUpperLimit *int32 `json:"adaptive_average_quality_upper_limit,omitempty"`

	// 自适应视频平均质量增限。取值范围为[1-50]。默认：3。
	AdaptiveAverageQualityIncreaseLimit *int32 `json:"adaptive_average_quality_increase_limit,omitempty"`

	// 自适应视频平均质量减限。取值范围为[1-50]。默认：1。
	AdaptiveAverageQualityDecreaseLimit *int32 `json:"adaptive_average_quality_decrease_limit,omitempty"`

	// 自适应视频最低质量下限。取值范围为[5-69]。默认：25。
	AdaptiveMinQualityLowerLimit *int32 `json:"adaptive_min_quality_lower_limit,omitempty"`

	// 自适应视频最低质量上限。取值范围为[5-69]。默认：30。
	AdaptiveMinQualityUpperLimit *int32 `json:"adaptive_min_quality_upper_limit,omitempty"`

	// 自适应视频最低质量增限。取值范围为[1-50]。默认：3。
	AdaptiveMinQualityIncreaseLimit *int32 `json:"adaptive_min_quality_increase_limit,omitempty"`

	// 自适应视频最低质量减限。取值范围为[1-50]。默认：1。
	AdaptiveMinQualityDecreaseLimit *int32 `json:"adaptive_min_quality_decrease_limit,omitempty"`
}

func (o PoliciesDisplayAdaptiveBitrateControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayAdaptiveBitrateControlOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayAdaptiveBitrateControlOptions", string(data)}, " ")
}
