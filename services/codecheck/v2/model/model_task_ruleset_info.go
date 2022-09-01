package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskRulesetInfo struct {

	// 规则集id
	TemplateId *string `json:"template_id,omitempty" xml:"template_id"`

	// 规则集语言
	Language *string `json:"language,omitempty" xml:"language"`

	// 规则集名称
	TemplateName *string `json:"template_name,omitempty" xml:"template_name"`

	// 规则集状态optional：可选，selected：已选
	Type *string `json:"type,omitempty" xml:"type"`

	// 规则集属性0 是默认用户规则集,1 是系统默认规则集
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o TaskRulesetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskRulesetInfo struct{}"
	}

	return strings.Join([]string{"TaskRulesetInfo", string(data)}, " ")
}
