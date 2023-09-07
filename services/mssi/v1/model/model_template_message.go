package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateMessage struct {

	// 模板描述
	TemplateDescription *string `json:"template_description,omitempty"`

	// 模板标签
	TemplateLabel string `json:"template_label"`

	// 模板名称
	TemplateName string `json:"template_name"`
}

func (o TemplateMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateMessage struct{}"
	}

	return strings.Join([]string{"TemplateMessage", string(data)}, " ")
}
