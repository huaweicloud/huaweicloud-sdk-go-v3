package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiConcatInfo struct {

	// 拼接任务输入源地址。
	Inputs []ObsObjInfo `json:"inputs" xml:"inputs"`

	// 拼接完成后转码对应的转码模板ID
	TransTemplateIds *[]int32 `json:"trans_template_ids,omitempty" xml:"trans_template_ids"`

	// 转码参数。 设置“trans_template_id”和此参数，则优先使用此参数进行转码。
	AvParameters *[]AvParameters `json:"av_parameters,omitempty" xml:"av_parameters"`

	Output *ObsObjInfo `json:"output" xml:"output"`

	// 水印信息。
	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty" xml:"image_watermark_settings"`
}

func (o MultiConcatInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiConcatInfo struct{}"
	}

	return strings.Join([]string{"MultiConcatInfo", string(data)}, " ")
}
