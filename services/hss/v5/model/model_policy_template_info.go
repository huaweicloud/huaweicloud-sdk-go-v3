package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTemplateInfo struct {

	// 模板ID
	Id string `json:"id"`

	// 模板名称
	TemplateName string `json:"template_name"`

	// 模板类型
	TemplateType string `json:"template_type"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 策略模板应用资源类型，多个资源类型通过分号份隔连接
	TargetKind string `json:"target_kind"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 推荐级别
	Level *string `json:"level,omitempty"`

	// 策略模板内容
	ConstraintTemplate string `json:"constraint_template"`
}

func (o PolicyTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTemplateInfo struct{}"
	}

	return strings.Join([]string{"PolicyTemplateInfo", string(data)}, " ")
}
