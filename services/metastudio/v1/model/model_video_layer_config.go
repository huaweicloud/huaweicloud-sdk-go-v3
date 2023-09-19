package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoLayerConfig 素材视频图层配置。
type VideoLayerConfig struct {

	// 视频文件的URL。
	VideoUrl *string `json:"video_url,omitempty"`

	// 视频封面文件的URL。
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`
}

func (o VideoLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoLayerConfig struct{}"
	}

	return strings.Join([]string{"VideoLayerConfig", string(data)}, " ")
}
