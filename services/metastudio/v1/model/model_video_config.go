package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoConfig 视频输出配置。
type VideoConfig struct {

	// 输出视频的剪辑方式。 * RESIZE：视频缩放。 * CROP：视频裁剪。
	ClipMode *VideoConfigClipMode `json:"clip_mode,omitempty"`

	// 视频编码格式及视频文件格式。 * H264: h264编码，输出mp4文件 * VP8：vp8编码，输出webm文件
	Codec VideoConfigCodec `json:"codec"`

	// 输出平均码率。  单位：kbps。  最小值40，最大值30000。 > * 分身数字人视频制作采用质量优先，可能会超过设置的码率。 > * 分身数字人直播码率范围[1000, 8000]。
	Bitrate int32 `json:"bitrate"`

	// 视频宽度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280四种分辨率。 > * clip_mode=CROP，视频保留中间width宽度，裁掉左右两边。 > * 分身数字人直播目前只支持1080x1920。
	Width int32 `json:"width"`

	// 视频高度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280四种分辨率。 > * clip_mode=CROP，视频保留底部height高度，裁掉顶部。 > * 分身数字人直播目前只支持1080x1920。
	Height int32 `json:"height"`

	// 帧率。  单位：FPS。 > * 分身数字人帧率目前只支持25。
	FrameRate *VideoConfigFrameRate `json:"frame_rate,omitempty"`

	// 输出的视频是否带字幕。 > true: 打开字幕 > false: 关闭字幕
	IsSubtitleEnable *bool `json:"is_subtitle_enable,omitempty"`

	// 输出的视频是否关闭系统水印。目前该参数需要白名单的租户才起作用。 > true: 关闭系统水印 > false: 不关闭系统水印
	DisableSystemWatermark *bool `json:"disable_system_watermark,omitempty"`

	// 裁剪视频左上角像素点坐标。  clip_mode= CROP时生效。 > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dx *int32 `json:"dx,omitempty"`

	// 裁剪视频左上角像素点坐标。  clip_mode= CROP时生效。 > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dy *int32 `json:"dy,omitempty"`
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
}

func GetVideoConfigCodecEnum() VideoConfigCodecEnum {
	return VideoConfigCodecEnum{
		H264: VideoConfigCodec{
			value: "H264",
		},
		VP8: VideoConfigCodec{
			value: "VP8",
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
