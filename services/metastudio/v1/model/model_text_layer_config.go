package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextLayerConfig 素材文字图层配置。
type TextLayerConfig struct {

	// **参数解释**： 文字图层的文本，内容需做Base64编码。  示例：若想添加文字水印“测试文字水印”，那么text_context的值为：5rWL6K+V5paH5a2X5rC05Y2w **约束限制**： 不涉及。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	TextContext *string `json:"text_context,omitempty"`

	// **参数解释**： 字体。当前支持的字体： **约束限制**： 不涉及。 **取值范围**： 支持的字体请参考[服务支持的字体](metastudio_02_0041.xml)
	FontName *string `json:"font_name,omitempty"`

	// **参数解释**： 字体大小（像素）。接口的取值范围为0-120，实际业务使用的取值范围要求为4-120，请以业务实际使用要求为准。  **约束限制**： 不涉及。
	FontSize *int32 `json:"font_size,omitempty"`

	// **参数解释**： 字体颜色。RGB颜色值。  **约束限制**： 不涉及。 **取值范围**： 字符长度0-16位
	FontColor *string `json:"font_color,omitempty"`
}

func (o TextLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextLayerConfig struct{}"
	}

	return strings.Join([]string{"TextLayerConfig", string(data)}, " ")
}
