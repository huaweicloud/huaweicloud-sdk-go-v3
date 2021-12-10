package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTaskRulesetRes struct {
	// 规则集id

	TemplateId *string `json:"template_id,omitempty"`
	// 规则集语言

	Language *string `json:"language,omitempty"`
	// 规则集名称

	TemplateName *string `json:"template_name,omitempty"`
	// 规则集状态optional：可选，selected：已选

	Type *string `json:"type,omitempty"`
	// 规则集属性0 是默认用户规则集,1 是系统默认规则集

	Status *string `json:"status,omitempty"`
}

func (o ListTaskRulesetRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetRes struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetRes", string(data)}, " ")
}
