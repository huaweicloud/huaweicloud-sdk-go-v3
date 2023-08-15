package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoConfig
type VideoConfig struct {

	// 背景id
	BackgroundId string `json:"background_id"`

	// 图标id
	LogoId *string `json:"logo_id,omitempty"`

	// 是否显示字幕 默认：false
	ShowSubtitles *bool `json:"show_subtitles,omitempty"`

	// 画面分辨率： 0: 宽屏landscape（默认） 1: 竖屏portrait
	ResolutionType *int32 `json:"resolution_type,omitempty"`
}

func (o VideoConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoConfig struct{}"
	}

	return strings.Join([]string{"VideoConfig", string(data)}, " ")
}
