package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyTransTemplate struct {

	// 模板组名称<br/>
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 模板组名称<br/>
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否是默认转码模板<br/>
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`

	// 是否开启加密
	IsAutoEncrypt *bool `json:"is_auto_encrypt,omitempty" xml:"is_auto_encrypt"`

	// 画质配置信息列表<br/>
	QualityInfoList *[]QualityInfoList `json:"quality_info_list,omitempty" xml:"quality_info_list"`

	// 绑定的水印模板组ID数组<br/>
	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty" xml:"watermark_template_ids"`

	// 模板介绍<br/>
	Description *string `json:"description,omitempty" xml:"description"`

	Common *CommonInfo `json:"common,omitempty" xml:"common"`
}

func (o ModifyTransTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTransTemplate struct{}"
	}

	return strings.Join([]string{"ModifyTransTemplate", string(data)}, " ")
}
