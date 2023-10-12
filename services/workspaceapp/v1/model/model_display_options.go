package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisplayOptions 显示级别共同控制的选项。
type DisplayOptions struct {

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
	LosslessCompressionMode *DisplayOptionsLosslessCompressionMode `json:"lossless_compression_mode,omitempty"`

	DeepCompressionOptions *DisplayOptionsDeepCompressionOptions `json:"deep_compression_options,omitempty"`

	// 有损压缩质量。取值范围为[20-100]。默认：85。
	LossyCompressionQuality *int32 `json:"lossy_compression_quality,omitempty"`

	// 办公场景色彩增强：取值为： false：表示关闭。 true：表示开启。
	ColorEnhancementEnable *bool `json:"color_enhancement_enable,omitempty"`

	// 质量或带宽优先。取值为： Quality First：表示初级压缩。 Bandwidth First：表示深度压缩。
	QualityBandwidthFirst *DisplayOptionsQualityBandwidthFirst `json:"quality_bandwidth_first,omitempty"`

	VideoBitRateOptions *DisplayOptionsVideoBitRateOptions `json:"video_bit_rate_options,omitempty"`

	// 视频峰值码率（Kbps）。取值范围为[256-100000]。默认：18000。
	PeakVideoBitRate *int32 `json:"peak_video_bit_rate,omitempty"`

	VideoQualityOptions *DisplayOptionsVideoQualityOptions `json:"video_quality_options,omitempty"`

	// GOP大小。取值范围为[0-65535]。默认：100。
	GopSize *int32 `json:"gop_size,omitempty"`

	// 编码预置。取值为： 预置1：Preset 1。 预置2：Preset 2。 预置3：Preset 3。 预置4：Preset 4。 预置5：Preset 5。 预置6：Preset 6。 预置7：Preset 7。
	EncodingPreset *DisplayOptionsEncodingPreset `json:"encoding_preset,omitempty"`
}

func (o DisplayOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayOptions struct{}"
	}

	return strings.Join([]string{"DisplayOptions", string(data)}, " ")
}

type DisplayOptionsLosslessCompressionMode struct {
	value string
}

type DisplayOptionsLosslessCompressionModeEnum struct {
	BASIC_COMPRESSION DisplayOptionsLosslessCompressionMode
	DEEP_COMPRESSION  DisplayOptionsLosslessCompressionMode
}

func GetDisplayOptionsLosslessCompressionModeEnum() DisplayOptionsLosslessCompressionModeEnum {
	return DisplayOptionsLosslessCompressionModeEnum{
		BASIC_COMPRESSION: DisplayOptionsLosslessCompressionMode{
			value: "Basic Compression",
		},
		DEEP_COMPRESSION: DisplayOptionsLosslessCompressionMode{
			value: "Deep Compression",
		},
	}
}

func (c DisplayOptionsLosslessCompressionMode) Value() string {
	return c.value
}

func (c DisplayOptionsLosslessCompressionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisplayOptionsLosslessCompressionMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type DisplayOptionsQualityBandwidthFirst struct {
	value string
}

type DisplayOptionsQualityBandwidthFirstEnum struct {
	QUALITY_FIRST   DisplayOptionsQualityBandwidthFirst
	BANDWIDTH_FIRST DisplayOptionsQualityBandwidthFirst
}

func GetDisplayOptionsQualityBandwidthFirstEnum() DisplayOptionsQualityBandwidthFirstEnum {
	return DisplayOptionsQualityBandwidthFirstEnum{
		QUALITY_FIRST: DisplayOptionsQualityBandwidthFirst{
			value: "Quality First",
		},
		BANDWIDTH_FIRST: DisplayOptionsQualityBandwidthFirst{
			value: "Bandwidth First",
		},
	}
}

func (c DisplayOptionsQualityBandwidthFirst) Value() string {
	return c.value
}

func (c DisplayOptionsQualityBandwidthFirst) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisplayOptionsQualityBandwidthFirst) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type DisplayOptionsEncodingPreset struct {
	value string
}

type DisplayOptionsEncodingPresetEnum struct {
	PRESET_1 DisplayOptionsEncodingPreset
	PRESET_2 DisplayOptionsEncodingPreset
	PRESET_3 DisplayOptionsEncodingPreset
	PRESET_4 DisplayOptionsEncodingPreset
	PRESET_5 DisplayOptionsEncodingPreset
	PRESET_6 DisplayOptionsEncodingPreset
	PRESET_7 DisplayOptionsEncodingPreset
}

func GetDisplayOptionsEncodingPresetEnum() DisplayOptionsEncodingPresetEnum {
	return DisplayOptionsEncodingPresetEnum{
		PRESET_1: DisplayOptionsEncodingPreset{
			value: "Preset 1",
		},
		PRESET_2: DisplayOptionsEncodingPreset{
			value: "Preset 2",
		},
		PRESET_3: DisplayOptionsEncodingPreset{
			value: "Preset 3",
		},
		PRESET_4: DisplayOptionsEncodingPreset{
			value: "Preset 4",
		},
		PRESET_5: DisplayOptionsEncodingPreset{
			value: "Preset 5",
		},
		PRESET_6: DisplayOptionsEncodingPreset{
			value: "Preset 6",
		},
		PRESET_7: DisplayOptionsEncodingPreset{
			value: "Preset 7",
		},
	}
}

func (c DisplayOptionsEncodingPreset) Value() string {
	return c.value
}

func (c DisplayOptionsEncodingPreset) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisplayOptionsEncodingPreset) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
