package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetaVideoInfo struct {

	// 视频码率，单位：bit/s
	Bitrate *int64 `json:"bitrate,omitempty"`

	// 视频编码格式
	Codec *string `json:"codec,omitempty"`

	// 帧率
	Fps *int32 `json:"fps,omitempty"`

	// 视频高度
	Height *int32 `json:"height,omitempty"`

	// 视频宽度
	Width *int32 `json:"width,omitempty"`
}

func (o MetaVideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaVideoInfo struct{}"
	}

	return strings.Join([]string{"MetaVideoInfo", string(data)}, " ")
}
