package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateParamVariableValidation 模板的部署参数变量的校验规则。
type TemplateParamVariableValidation struct {

	// 校验表达式。
	Condition *string `json:"condition,omitempty"`

	// 校验失败后的错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o TemplateParamVariableValidation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateParamVariableValidation struct{}"
	}

	return strings.Join([]string{"TemplateParamVariableValidation", string(data)}, " ")
}
