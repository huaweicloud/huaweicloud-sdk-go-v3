package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartVideoLayerConfig 素材视频图层配置。
type SmartVideoLayerConfig struct {

	// 视频文件的URL。
	VideoUrl string `json:"video_url"`

	// 视频封面文件的URL。
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`

	// 视频显示时长。单位s * 0：表示一直显示。
	DisplayDuration *int32 `json:"display_duration,omitempty"`
}

func (o SmartVideoLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartVideoLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartVideoLayerConfig", string(data)}, " ")
}
