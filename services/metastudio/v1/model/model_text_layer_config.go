package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextLayerConfig 素材文字图层配置。
type TextLayerConfig struct {

	// 文字图层的文本，内容需做Base64编码。  示例：若想添加文字水印“测试文字水印”，那么text_context的值为：5rWL6K+V5paH5a2X5rC05Y2w
	TextContext *string `json:"text_context,omitempty"`

	// 字体。当前支持的字体： * HarmonyOS_Sans_SC_Black：鸿蒙粗体 * HarmonyOS_Sans_SC_Regular：鸿蒙常规 * HarmonyOS_Sans_SC_Thin：鸿蒙细体 * fzyouh：方正瘦体
	FontName *string `json:"font_name,omitempty"`

	// 字体大小（像素）。  取值范围：[4, 120]
	FontSize *int32 `json:"font_size,omitempty"`

	// 字体颜色。RGB颜色值。
	FontColor *string `json:"font_color,omitempty"`
}

func (o TextLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextLayerConfig struct{}"
	}

	return strings.Join([]string{"TextLayerConfig", string(data)}, " ")
}
