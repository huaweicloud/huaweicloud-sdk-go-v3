package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoAssetMeta 视频元数据，自动提取获得。
type VideoAssetMeta struct {

	// 视频编码格式
	VideoCodec *string `json:"video_codec,omitempty"`

	// 视频宽度
	Width *int32 `json:"width,omitempty"`

	// 视频高度
	Height *int32 `json:"height,omitempty"`

	// 帧率
	FrameRate *string `json:"frame_rate,omitempty"`

	// 视频平均码率,单位kbps
	VideoBitRate *int32 `json:"video_bit_rate,omitempty"`

	// 时长,单位秒
	Duration *int32 `json:"duration,omitempty"`

	// 音频编码格式
	AudioCodec *string `json:"audio_codec,omitempty"`

	// 音频平均码率,单位kbps
	AudioBitRate *int32 `json:"audio_bit_rate,omitempty"`

	// 音频声道数
	AudioChannels *int32 `json:"audio_channels,omitempty"`

	// 采样率,HZ
	Sample *int32 `json:"sample,omitempty"`

	// Horizontal=横向；Vertical=纵向
	Mode *VideoAssetMetaMode `json:"mode,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o VideoAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoAssetMeta struct{}"
	}

	return strings.Join([]string{"VideoAssetMeta", string(data)}, " ")
}

type VideoAssetMetaMode struct {
	value string
}

type VideoAssetMetaModeEnum struct {
	HORIZONTAL VideoAssetMetaMode
	VERTICAL   VideoAssetMetaMode
}

func GetVideoAssetMetaModeEnum() VideoAssetMetaModeEnum {
	return VideoAssetMetaModeEnum{
		HORIZONTAL: VideoAssetMetaMode{
			value: "Horizontal",
		},
		VERTICAL: VideoAssetMetaMode{
			value: "Vertical",
		},
	}
}

func (c VideoAssetMetaMode) Value() string {
	return c.value
}

func (c VideoAssetMetaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoAssetMetaMode) UnmarshalJSON(b []byte) error {
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
