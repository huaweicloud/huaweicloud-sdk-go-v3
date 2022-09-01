package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateGroup struct {

	// 模板组id<br/>
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 模板组名称<br/>
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否默认<br/>
	Status *string `json:"status,omitempty" xml:"status"`

	// 模板组类型<br/>
	Type *string `json:"type,omitempty" xml:"type"`

	// 是否自动加密。  取值如下： - 0：表示不加密。 - 1：表示需要加密。  默认值：0。  加密与转码必须要一起进行，当需要加密时，转码参数不能为空，且转码输出格式必须要为HLS。
	AutoEncrypt *int32 `json:"auto_encrypt,omitempty" xml:"auto_encrypt"`

	// 画质配置信息列表<br/>
	QualityInfoList *[]QualityInfo `json:"quality_info_list,omitempty" xml:"quality_info_list"`

	// 绑定的水印模板组ID数组<br/>
	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty" xml:"watermark_template_ids"`

	// 模板介绍<br/>
	Description *string `json:"description,omitempty" xml:"description"`

	Common *Common `json:"common,omitempty" xml:"common"`
}

func (o TemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateGroup struct{}"
	}

	return strings.Join([]string{"TemplateGroup", string(data)}, " ")
}
