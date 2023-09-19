package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextLayerConfig 素材文字图层配置。
type TextLayerConfig struct {

	// 文字水印内容，内容需做Base64编码，若类型为文字水印 (type字段为Text)，则此配置项不能为空  示例：若想添加文字水印“测试文字水印”，那么Content的值为：5rWL6K+V5paH5a2X5rC05Y2w
	TextContext *string `json:"text_context,omitempty"`

	// 字体，当前支持fzyouh
	FontName *string `json:"font_name,omitempty"`

	// 字体大小。  取值范围：[4, 120]
	FontSize *int32 `json:"font_size,omitempty"`

	// 字体颜色。 目前颜色支持 black，blue，white，green，red，yellow，brown，gold，pink，orange，purple。
	FontColor *string `json:"font_color,omitempty"`
}

func (o TextLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextLayerConfig struct{}"
	}

	return strings.Join([]string{"TextLayerConfig", string(data)}, " ")
}
