package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociationAlarmTemplate struct {

	// 告警模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 告警模板名称
	TemplateName *string `json:"template_name,omitempty"`
}

func (o AssociationAlarmTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationAlarmTemplate struct{}"
	}

	return strings.Join([]string{"AssociationAlarmTemplate", string(data)}, " ")
}
