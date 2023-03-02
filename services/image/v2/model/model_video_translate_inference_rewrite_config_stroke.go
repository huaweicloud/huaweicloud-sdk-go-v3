package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 字体描边配置
type VideoTranslateInferenceRewriteConfigStroke struct {

	// 文本描边颜色
	StrokeColor *[]int32 `json:"stroke_color,omitempty"`

	// 文本字体颜色
	FontColor *[]int32 `json:"font_color,omitempty"`

	// 描边宽度
	StrokeRatio *float32 `json:"stroke_ratio,omitempty"`
}

func (o VideoTranslateInferenceRewriteConfigStroke) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateInferenceRewriteConfigStroke struct{}"
	}

	return strings.Join([]string{"VideoTranslateInferenceRewriteConfigStroke", string(data)}, " ")
}
