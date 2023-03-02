package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图像合成视频参数配置信息
type ImageToVideoInfo struct {

	// 背景音乐url
	BgmUrl *string `json:"bgm_url,omitempty"`

	// 生成视频或gif开关，1生成gif，0生成视频，默认为视频
	ToGif *int32 `json:"to_gif,omitempty"`

	// gif分辨率压缩率，gif最终分辨率为最终合成视频分辨率*gif_compress
	GifCompress *float32 `json:"gif_compress,omitempty"`

	// 图像展示时间List
	ImageDurations *[]float32 `json:"image_durations,omitempty"`

	// 分辨率
	Resolution *[]int32 `json:"resolution,omitempty"`

	// 视频帧率，默认30
	Fps *int32 `json:"fps,omitempty"`

	// 动画转场时间，默认1s
	AnimationDuration *float32 `json:"animation_duration,omitempty"`
}

func (o ImageToVideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageToVideoInfo struct{}"
	}

	return strings.Join([]string{"ImageToVideoInfo", string(data)}, " ")
}
