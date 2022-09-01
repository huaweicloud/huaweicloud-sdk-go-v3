package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentTempV2 struct {

	// 模板id
	IncidentTemplateId *string `json:"incident_template_id,omitempty" xml:"incident_template_id"`

	// 模板名称
	IncidentTemplateName *string `json:"incident_template_name,omitempty" xml:"incident_template_name"`

	// 模板内容
	IncidentTemplateContent *string `json:"incident_template_content,omitempty" xml:"incident_template_content"`
}

func (o IncidentTempV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentTempV2 struct{}"
	}

	return strings.Join([]string{"IncidentTempV2", string(data)}, " ")
}
