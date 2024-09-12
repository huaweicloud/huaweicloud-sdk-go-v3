package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoConfig 视频输出配置。
type VideoConfig struct {

	// **参数解释**： 输出视频的剪辑方式。 **约束限制**： 不涉及。 **取值范围**： * RESIZE：视频缩放。 * CROP：视频裁剪。
	ClipMode *VideoConfigClipMode `json:"clip_mode,omitempty"`

	// **参数解释**： 视频编码格式及视频文件格式。 **约束限制**： 仅分身数字人视频制作支持VP8编码。  **取值范围**： * H264：h264编码，输出mp4文件。 * VP8：vp8编码，输出webm文件。  **默认取值**： 不涉及
	Codec VideoConfigCodec `json:"codec"`

	// **参数解释**： 输出平均码率。单位：kbps。 **约束限制**： * 分身数字人视频制作采用质量优先，可能会超过设置的码率。 * 分身数字人直播码率范围[1000, 8000]。  **默认取值**： 不涉及
	Bitrate int32 `json:"bitrate"`

	// **参数解释**： 视频宽度。单位：像素。 **约束限制**： * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280、3840x2160、2160x3840六种分辨率。4K分辨率视频需要分身数字人模型支持4K的情况下才能使用。 * clip_mode=CROP，裁剪后视频，（dx,dy）为原点，保留视频像宽度为width。 * 分身数字人直播目前只支持1080x1920、1920x1080。  **默认取值**： 不涉及
	Width int32 `json:"width"`

	// **参数解释**： 视频高度。  单位：像素。 **约束限制**： * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280、3840x2160、2160x3840六种分辨率分辨率。 * clip_mode=CROP，裁剪后视频，（dx,dy）为原点，保留视频像高度为height。 * 分身数字人直播目前只支持1080x1920、1920x1080。  **默认取值**： 不涉及
	Height int32 `json:"height"`

	// **参数解释**： 帧率。单位：FPS。 **约束限制**： 分身数字人视频固定25FPS。
	FrameRate *VideoConfigFrameRate `json:"frame_rate,omitempty"`

	// **参数解释**： 输出的视频是否带字幕。 **约束限制**： 分身数字人直播暂时不支持字幕。  **取值范围**： * true: 打开字幕 * false: 关闭字幕
	IsSubtitleEnable *bool `json:"is_subtitle_enable,omitempty"`

	SubtitleConfig *SubtitleConfig `json:"subtitle_config,omitempty"`

	// **参数解释**： 裁剪视频左上角像素点横坐标。 > 以模特分辨率为画布大小，比如1920*1080分辨率的模特，dx最小值是0，最大值是1920。  **约束限制**： clip_mode= CROP时生效。 **默认取值**： 不涉及
	Dx *int32 `json:"dx,omitempty"`

	// **参数解释**： 裁剪视频左上角像素点纵坐标。 > 以模特分辨率为画布大小，比如1920*1080分辨率的模特，dy最小值是0，最大值是1080  **约束限制**： clip_mode= CROP时生效。 **默认取值**： 不涉及
	Dy *int32 `json:"dy,omitempty"`

	// **参数解释**： 视频是否开启超分。 **约束限制**： 仅分身数字人视频制作支持。 **取值范围** * true: 开启 * false: 不开启
	IsEnableSuperResolution *bool `json:"is_enable_super_resolution,omitempty"`
}

func (o VideoConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoConfig struct{}"
	}

	return strings.Join([]string{"VideoConfig", string(data)}, " ")
}

type VideoConfigClipMode struct {
	value string
}

type VideoConfigClipModeEnum struct {
	RESIZE VideoConfigClipMode
	CROP   VideoConfigClipMode
}

func GetVideoConfigClipModeEnum() VideoConfigClipModeEnum {
	return VideoConfigClipModeEnum{
		RESIZE: VideoConfigClipMode{
			value: "RESIZE",
		},
		CROP: VideoConfigClipMode{
			value: "CROP",
		},
	}
}

func (c VideoConfigClipMode) Value() string {
	return c.value
}

func (c VideoConfigClipMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoConfigClipMode) UnmarshalJSON(b []byte) error {
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

type VideoConfigCodec struct {
	value string
}

type VideoConfigCodecEnum struct {
	H264 VideoConfigCodec
	VP8  VideoConfigCodec
	VP9  VideoConfigCodec
}

func GetVideoConfigCodecEnum() VideoConfigCodecEnum {
	return VideoConfigCodecEnum{
		H264: VideoConfigCodec{
			value: "H264",
		},
		VP8: VideoConfigCodec{
			value: "VP8",
		},
		VP9: VideoConfigCodec{
			value: "VP9",
		},
	}
}

func (c VideoConfigCodec) Value() string {
	return c.value
}

func (c VideoConfigCodec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoConfigCodec) UnmarshalJSON(b []byte) error {
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

type VideoConfigFrameRate struct {
	value string
}

type VideoConfigFrameRateEnum struct {
	E_24 VideoConfigFrameRate
	E_25 VideoConfigFrameRate
	E_30 VideoConfigFrameRate
}

func GetVideoConfigFrameRateEnum() VideoConfigFrameRateEnum {
	return VideoConfigFrameRateEnum{
		E_24: VideoConfigFrameRate{
			value: "24",
		},
		E_25: VideoConfigFrameRate{
			value: "25",
		},
		E_30: VideoConfigFrameRate{
			value: "30",
		},
	}
}

func (c VideoConfigFrameRate) Value() string {
	return c.value
}

func (c VideoConfigFrameRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoConfigFrameRate) UnmarshalJSON(b []byte) error {
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
