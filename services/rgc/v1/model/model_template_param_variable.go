package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateParamVariable 模板的部署参数的变量。
type TemplateParamVariable struct {

	// 变量默认值。
	Default *interface{} `json:"default,omitempty"`

	// 变量名称。
	Name *string `json:"name,omitempty"`

	// 变量描述。
	Description *string `json:"description,omitempty"`

	// 变量是否可以为空。
	Nullable *bool `json:"nullable,omitempty"`

	// 变量是否为敏感字段。
	Sensitive *bool `json:"sensitive,omitempty"`

	// 变量类型。
	Type *string `json:"type,omitempty"`

	Validations *[]TemplateParamVariableValidation `json:"validations,omitempty"`
}

func (o TemplateParamVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateParamVariable struct{}"
	}

	return strings.Join([]string{"TemplateParamVariable", string(data)}, " ")
}
