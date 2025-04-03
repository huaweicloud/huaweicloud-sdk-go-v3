package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubtitleConfig 字幕配置。
type SubtitleConfig struct {

	// **参数解释**： 字幕框左下角像素点坐标。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	Dx *int32 `json:"dx,omitempty"`

	// **参数解释**： 字幕框左下角像素点坐标。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	Dy *int32 `json:"dy,omitempty"`

	// **参数解释**： 字幕框高度。 **约束限制**： 参数h用于方便前端计算字幕框左上角坐标，后台不使用该参数。
	H *int32 `json:"h,omitempty"`

	// **参数解释**： 字幕框宽度。 **约束限制**： * 字幕框宽度固定为屏幕宽度的80% * 参数w用于方便前端计算字幕框左上角坐标，后台不使用该参数
	W *int32 `json:"w,omitempty"`

	// **参数解释**： 字体。当前支持的字体请参考[服务支持的字体](metastudio_02_0041.xml) **约束限制**： 不涉及。 **取值范围**： 字符长度0-64位
	FontName *string `json:"font_name,omitempty"`

	// **参数解释**： 字体大小。接口的取值范围为0-120，实际业务使用的取值范围要求为24-120，请以业务实际使用要求为准。 **约束限制**： 不涉及。
	FontSize *int32 `json:"font_size,omitempty"`

	// **参数解释**： 字幕字体颜色的RGB颜色值。 **约束限制**： 无 **取值范围**： 字符长度0-7位，固定长度
	FontColor *string `json:"font_color,omitempty"`

	// **参数解释**： 字幕字体描边颜色的RGB颜色值。 **约束限制**： 无 **取值范围**： 字符长度0-7位，固定长度
	StrokeColor *string `json:"stroke_color,omitempty"`

	// **参数解释**： 字幕字体描边粗细像素值。 **约束限制**： 无 **取值范围**： 0-50
	StrokeThickness *float32 `json:"stroke_thickness,omitempty"`

	// **参数解释**： 字幕字体不透明度，0表示完全透明，1表示完全不透明。默认1。 **约束限制**： 无 **取值范围**： 0-1
	Opacity *float32 `json:"opacity,omitempty"`
}

func (o SubtitleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleConfig struct{}"
	}

	return strings.Join([]string{"SubtitleConfig", string(data)}, " ")
}
