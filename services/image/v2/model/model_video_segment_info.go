package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 视频剪切分段参数信息
type VideoSegmentInfo struct {

	// 视频分段起始时间
	StartTime float32 `json:"start_time"`

	// 视频分段持续时间
	Duration float32 `json:"duration"`

	// 视频剪切服务生成视频或gif开关，1生成gif，0生成视频，默认为视频
	ToGif *int32 `json:"to_gif,omitempty"`

	// 视频或gif倍速，默认1
	Speed *float32 `json:"speed,omitempty"`

	// gif分辨率压缩率，gif最终分辨率为最终合成视频分辨率*gif_compress
	GifCompress *float32 `json:"gif_compress,omitempty"`
}

func (o VideoSegmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSegmentInfo struct{}"
	}

	return strings.Join([]string{"VideoSegmentInfo", string(data)}, " ")
}
