package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterConstraint 参数属性。
type ParameterConstraint struct {

	// 参数种类。
	Type *string `json:"type,omitempty"`

	// 是否可编辑。
	Editable *bool `json:"editable,omitempty"`

	// 是否必须。
	Required *bool `json:"required,omitempty"`

	// 是否敏感。该功能暂未实现。
	Sensitive *bool `json:"sensitive,omitempty"`

	// 有效种类。
	ValidType *string `json:"valid_type,omitempty"`

	// 有效范围。
	ValidRange *[]string `json:"valid_range,omitempty"`
}

func (o ParameterConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterConstraint struct{}"
	}

	return strings.Join([]string{"ParameterConstraint", string(data)}, " ")
}
