package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Parameter 训练作业、算法依赖参数。
type Parameter struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 参数描述信息。
	Description *string `json:"description,omitempty"`

	Constraint *ParameterConstraint `json:"constraint,omitempty"`

	I18nDescription *ParameterI18nDescription `json:"i18n_description,omitempty"`
}

func (o Parameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parameter struct{}"
	}

	return strings.Join([]string{"Parameter", string(data)}, " ")
}
