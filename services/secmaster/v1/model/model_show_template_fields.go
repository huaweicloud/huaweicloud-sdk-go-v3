package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTemplateFields struct {

	// 相关描述信息
	Fields *[]TemplateField `json:"fields,omitempty"`

	// UUID
	TemplateId *string `json:"template_id,omitempty"`
}

func (o ShowTemplateFields) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateFields struct{}"
	}

	return strings.Join([]string{"ShowTemplateFields", string(data)}, " ")
}
