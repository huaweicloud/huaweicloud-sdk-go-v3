package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoConfigResp
type VideoConfigResp struct {

	// 背景id
	BackgroundId string `json:"background_id"`

	// 图标id
	LogoId *string `json:"logo_id,omitempty"`

	// 是否显示字幕 默认：false
	ShowSubtitles *bool `json:"show_subtitles,omitempty"`

	// 画面分辨率： 0: 宽屏landscape（默认） 1: 竖屏portrait
	ResolutionType *int32 `json:"resolution_type,omitempty"`

	// 背景图片地址，取默认背景的第一张
	BackgroundUrl string `json:"background_url"`

	// 播报框地址 和background绑定，如果使用用户自定义背景，则使用演播厅框
	ImageFrameUrl string `json:"image_frame_url"`

	// logo地址
	LogoUrl *string `json:"logo_url,omitempty"`
}

func (o VideoConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoConfigResp struct{}"
	}

	return strings.Join([]string{"VideoConfigResp", string(data)}, " ")
}
