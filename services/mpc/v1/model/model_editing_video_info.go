package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingVideoInfo struct {

	// 视频宽度
	Width *int32 `json:"width,omitempty"`

	// 视频高度
	Height *int32 `json:"height,omitempty"`

	// 视频码率，单位: kbit/s
	Bitrate *int32 `json:"bitrate,omitempty"`

	// 视频码率，单位: bit/s
	BitrateBps *int64 `json:"bitrate_bps,omitempty"`

	// 输出帧率。
	FrameRate *int32 `json:"frame_rate,omitempty"`

	// 视频编码格式
	Codec *string `json:"codec,omitempty"`
}

func (o EditingVideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingVideoInfo struct{}"
	}

	return strings.Join([]string{"EditingVideoInfo", string(data)}, " ")
}
