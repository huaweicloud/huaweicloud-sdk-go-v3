package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Parameters 训练作业的运行参数。
type Parameters struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 参数描述信息。
	Description *string `json:"description,omitempty"`

	Constraint *ParametersConstraint `json:"constraint,omitempty"`

	I18nDescription *I18nDescription `json:"i18n_description,omitempty"`
}

func (o Parameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parameters struct{}"
	}

	return strings.Join([]string{"Parameters", string(data)}, " ")
}
