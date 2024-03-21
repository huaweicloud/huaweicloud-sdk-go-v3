package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PredefinedTemplate 预置模板信息。
type PredefinedTemplate struct {

	// 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 模板描述。
	TemplateDescription *string `json:"template_description,omitempty"`

	// 模板文件名称。
	TemplateFileName *string `json:"template_file_name,omitempty"`
}

func (o PredefinedTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PredefinedTemplate struct{}"
	}

	return strings.Join([]string{"PredefinedTemplate", string(data)}, " ")
}
