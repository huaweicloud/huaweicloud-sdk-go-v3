package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplates struct {

	// **参数解释**： 告警模板的ID。     **约束限制**： 不涉及。 **取值范围**： 以at开头，后跟字母、数字，长度为[2,64]个字符。           **默认取值**： 不涉及。
	TemplateId string `json:"template_id"`

	// **参数解释**： 告警模板的名称。 **约束限制**： 不涉及。 **取值范围**： 以字母或汉字开头，可包含字母、数字、汉字、_、-，长度为[1,128]个字符。           **默认取值**： 不涉及。
	TemplateName string `json:"template_name"`

	TemplateType *TemplateType `json:"template_type"`

	// **参数解释**： 告警模板的创建时间 **取值范围**： 不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// **参数解释**： 告警模板的描述     **约束限制**： 不涉及。 **取值范围**： 长度范围[0,256]。          **默认取值**： 空字符串。
	TemplateDescription *string `json:"template_description,omitempty"`
}

func (o AlarmTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplates struct{}"
	}

	return strings.Join([]string{"AlarmTemplates", string(data)}, " ")
}
