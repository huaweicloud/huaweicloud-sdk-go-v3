package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 视频合成参数配置信息
type VideoSynthesisInfo struct {

	// 背景音乐url
	BgmUrl *string `json:"bgm_url,omitempty"`

	// 视频分辨率
	Resolution *[]int32 `json:"resolution,omitempty"`

	// 视频帧率
	Fps *int32 `json:"fps,omitempty"`

	// 转场动画时间
	AnimationDuration *float32 `json:"animation_duration,omitempty"`
}

func (o VideoSynthesisInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSynthesisInfo struct{}"
	}

	return strings.Join([]string{"VideoSynthesisInfo", string(data)}, " ")
}
