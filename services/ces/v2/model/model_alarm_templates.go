package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplates struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId string `json:"template_id"`

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName string `json:"template_name"`

	TemplateType *TemplateType `json:"template_type"`

	// 告警模板的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription *string `json:"template_description,omitempty"`
}

func (o AlarmTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplates struct{}"
	}

	return strings.Join([]string{"AlarmTemplates", string(data)}, " ")
}
