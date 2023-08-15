package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlarmTemplateRequestBody struct {

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName string `json:"template_name"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription *string `json:"template_description,omitempty"`

	// 告警模板策略列表
	Policies []Policies `json:"policies"`
}

func (o UpdateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequestBody", string(data)}, " ")
}
