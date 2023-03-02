package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 视频合成服务推理参数
type VideoSynthesisInference struct {
	VideoConfig *VideoSynthesisInfo `json:"video_config"`
}

func (o VideoSynthesisInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSynthesisInference struct{}"
	}

	return strings.Join([]string{"VideoSynthesisInference", string(data)}, " ")
}
