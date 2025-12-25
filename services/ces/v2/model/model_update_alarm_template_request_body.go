package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlarmTemplateRequestBody struct {

	// **参数解释**： 告警模板的名称。 **约束限制**： 不涉及。 **取值范围**： 以字母或汉字开头，可包含字母、数字、汉字、_、-，长度为[1,128]个字符。           **默认取值**： 不涉及。
	TemplateName string `json:"template_name"`

	TemplateType *TemplateTypeUpdate `json:"template_type,omitempty"`

	// **参数解释**： 告警模板的描述     **约束限制**： 不涉及。 **取值范围**： 长度范围[0,256]。          **默认取值**： 空字符串。
	TemplateDescription *string `json:"template_description,omitempty"`

	// 告警模板策略列表
	Policies []UpdateAlarmTemplatePolicies `json:"policies"`
}

func (o UpdateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequestBody", string(data)}, " ")
}
