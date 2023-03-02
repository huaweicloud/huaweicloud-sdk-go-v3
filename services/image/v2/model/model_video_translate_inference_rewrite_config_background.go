package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 背景框配置
type VideoTranslateInferenceRewriteConfigBackground struct {

	// 文本背景框颜色
	BoxColor *[]int32 `json:"box_color,omitempty"`

	// 文本字体颜色
	BackgroundFontColor *[]int32 `json:"background_font_color,omitempty"`
}

func (o VideoTranslateInferenceRewriteConfigBackground) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateInferenceRewriteConfigBackground struct{}"
	}

	return strings.Join([]string{"VideoTranslateInferenceRewriteConfigBackground", string(data)}, " ")
}
