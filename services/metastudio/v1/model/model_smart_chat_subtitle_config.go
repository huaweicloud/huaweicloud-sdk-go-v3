package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartChatSubtitleConfig 字幕配置。
type SmartChatSubtitleConfig struct {

	// 字幕左上角像素点坐标。  > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dx *int32 `json:"dx,omitempty"`

	// 字幕左上角像素点坐标。  > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dy *int32 `json:"dy,omitempty"`

	// 图层图片宽度像素值（相对画布大小）。 > 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。
	Width *int32 `json:"width,omitempty"`

	// 图层图片高度像素值（相对画布大小）。 > 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。
	Height *int32 `json:"height,omitempty"`
}

func (o SmartChatSubtitleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatSubtitleConfig struct{}"
	}

	return strings.Join([]string{"SmartChatSubtitleConfig", string(data)}, " ")
}
