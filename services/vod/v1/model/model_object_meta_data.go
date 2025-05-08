package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectMetaData struct {

	// 总码率，单位：bit/秒
	Bitrate *int64 `json:"bitrate,omitempty"`

	// 封装格式
	Container *string `json:"container,omitempty"`

	// 时长，单位：秒
	Duration *int32 `json:"duration,omitempty"`

	// 时长，单位：秒
	VideoDuration *float32 `json:"video_duration,omitempty"`

	// 时长，单位：秒
	AudioDuration *float32 `json:"audio_duration,omitempty"`

	// 时长，单位：秒
	FloatDuration *float32 `json:"float_duration,omitempty"`

	// 视频高度
	Height *int32 `json:"height,omitempty"`

	// 视频宽度
	Width *int32 `json:"width,omitempty"`

	// 视频md5
	Md5 *string `json:"md5,omitempty"`

	// 视频拍摄角度
	Rotate *float32 `json:"rotate,omitempty"`

	// 文件大小，单位：byte
	Size *int64 `json:"size,omitempty"`

	// 视频流元数据
	VideoStreamList *[]MetaVideoInfo `json:"video_stream_list,omitempty"`

	// 音频流元数据
	AudioStreamList *[]MetaAudioInfo `json:"audio_stream_list,omitempty"`
}

func (o ObjectMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectMetaData struct{}"
	}

	return strings.Join([]string{"ObjectMetaData", string(data)}, " ")
}
