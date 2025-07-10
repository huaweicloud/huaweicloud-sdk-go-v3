package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDisplayOptions 显示级别共同控制的选项。
type PoliciesDisplayOptions struct {

	// 带宽（Kbps）。取值范围为[256-25000]。默认：20000。
	DisplayBandwidth *int32 `json:"display_bandwidth,omitempty"`

	// 帧率（fps）。取值范围为[1-60]。默认：25。
	FrameRate *int32 `json:"frame_rate,omitempty"`

	// 视频帧率（fps）。取值范围为[1-60]。默认：30。
	VideoFrameRate *int32 `json:"video_frame_rate,omitempty"`

	// 图像缓存最低容量（MB）。取值范围[0-300]，默认：200。
	MinImageCache *int32 `json:"min_image_cache,omitempty"`

	// 有损压缩识别阈值。取值范围为[0-255]。默认：60。
	SmoothingFactor *int32 `json:"smoothing_factor,omitempty"`

	// 无损压缩模式。取值为： Basic Compression：表示初级压缩。 Deep Compression：表示深度压缩。
	LosslessCompressionMode *string `json:"lossless_compression_mode,omitempty"`

	DeepCompressionOptions *PoliciesDisplayOptionsDeepCompressionOptions `json:"deep_compression_options,omitempty"`

	// 有损压缩质量。取值范围为[20-100]。默认：85。
	LossyCompressionQuality *int32 `json:"lossy_compression_quality,omitempty"`

	// 办公场景色彩增强：取值为： false：表示关闭。 true：表示开启。
	ColorEnhancementEnable *bool `json:"color_enhancement_enable,omitempty"`

	// 质量或带宽优先。取值为： Quality First：表示初级压缩。 Bandwidth First：表示深度压缩。
	QualityBandwidthFirst *string `json:"quality_bandwidth_first,omitempty"`

	VideoBitRateOptions *PoliciesDisplayOptionsVideoBitRateOptions `json:"video_bit_rate_options,omitempty"`

	// 视频峰值码率（Kbps）。取值范围为[256-100000]。默认：18000。
	PeakVideoBitRate *int32 `json:"peak_video_bit_rate,omitempty"`

	VideoQualityOptions *PoliciesDisplayOptionsVideoQualityOptions `json:"video_quality_options,omitempty"`

	// GOP大小。取值范围为[0-65535]。默认：100。
	GopSize *int32 `json:"gop_size,omitempty"`

	// 编码预置。取值为： 预置1：Preset 1。 预置2：Preset 2。 预置3：Preset 3。 预置4：Preset 4。 预置5：Preset 5。 预置6：Preset 6。 预置7：Preset 7。
	EncodingPreset *string `json:"encoding_preset,omitempty"`
}

func (o PoliciesDisplayOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayOptions", string(data)}, " ")
}
