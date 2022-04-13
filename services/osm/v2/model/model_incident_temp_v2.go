package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentTempV2 struct {
	// 模板id

	IncidentTemplateId *string `json:"incident_template_id,omitempty"`
	// 模板名称

	IncidentTemplateName *string `json:"incident_template_name,omitempty"`
	// 模板内容

	IncidentTemplateContent *string `json:"incident_template_content,omitempty"`
}

func (o IncidentTempV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentTempV2 struct{}"
	}

	return strings.Join([]string{"IncidentTempV2", string(data)}, " ")
}
