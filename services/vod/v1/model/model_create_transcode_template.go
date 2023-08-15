package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTranscodeTemplate struct {

	// 模板组名称<br/>
	Name string `json:"name"`

	// 是否设置成默认转码模板<br/>
	IsDefault *bool `json:"is_default,omitempty"`

	// 是否开启加密
	IsAutoEncrypt *bool `json:"is_auto_encrypt,omitempty"`

	// 画质配置信息列表<br/>
	QualityInfoList []QualityInfoList `json:"quality_info_list"`

	Common *CommonInfo `json:"common"`

	// 绑定的水印模板组ID数组<br/>
	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty"`

	// 模板介绍<br/>
	Description *string `json:"description,omitempty"`
}

func (o CreateTranscodeTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodeTemplate struct{}"
	}

	return strings.Join([]string{"CreateTranscodeTemplate", string(data)}, " ")
}
