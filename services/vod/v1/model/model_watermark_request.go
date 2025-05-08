package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WatermarkRequest struct {

	// 水印模板ID
	TemplateId *string `json:"template_id,omitempty"`

	Input *ObsInfo `json:"input,omitempty"`

	ImageWatermark *ImageWatermark `json:"image_watermark,omitempty"`

	// 文字水印内容，内容需做Base64编码，若类型为文字水印 (type字段为Text)，则此配置项不能为空
	TextContext *string `json:"text_context,omitempty"`

	TextWatermark *TextWatermark `json:"text_watermark,omitempty"`

	// svg水印的内容
	SvgContext *string `json:"svg_context,omitempty"`

	SvgWatermark *SvgWatermark `json:"svg_watermark,omitempty"`
}

func (o WatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WatermarkRequest struct{}"
	}

	return strings.Join([]string{"WatermarkRequest", string(data)}, " ")
}
