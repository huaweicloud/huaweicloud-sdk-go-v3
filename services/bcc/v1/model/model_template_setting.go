package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateSetting 模板配置
type TemplateSetting struct {

	// 模板名称
	TemplateName string `json:"template_name"`

	// 模板内容列表
	TemplateContents []TemplateItemEnum `json:"template_contents"`
}

func (o TemplateSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateSetting struct{}"
	}

	return strings.Join([]string{"TemplateSetting", string(data)}, " ")
}
