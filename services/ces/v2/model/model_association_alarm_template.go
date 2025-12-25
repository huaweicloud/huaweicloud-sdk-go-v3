package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociationAlarmTemplate struct {

	// **参数解释** 告警模板ID。 **取值范围** 以at开头，后跟字母、数字，长度在[2,64]区间内。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释** 告警模板名称。 **取值范围** 以字母或汉字开头，可包含字母、数字、汉字、_、-，长度在[1,128]区间内。
	TemplateName *string `json:"template_name,omitempty"`
}

func (o AssociationAlarmTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationAlarmTemplate struct{}"
	}

	return strings.Join([]string{"AssociationAlarmTemplate", string(data)}, " ")
}
