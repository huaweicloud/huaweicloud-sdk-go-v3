package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 回写配置
type VideoTranslateInferenceRewriteConfig struct {

	// 字体类型
	FontType *string `json:"font_type,omitempty"`

	// 字幕字体行间距
	RewriteRowInterval *float32 `json:"rewrite_row_interval,omitempty"`

	Stroke *VideoTranslateInferenceRewriteConfigStroke `json:"stroke,omitempty"`

	Background *VideoTranslateInferenceRewriteConfigBackground `json:"background,omitempty"`
}

func (o VideoTranslateInferenceRewriteConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateInferenceRewriteConfig struct{}"
	}

	return strings.Join([]string{"VideoTranslateInferenceRewriteConfig", string(data)}, " ")
}
