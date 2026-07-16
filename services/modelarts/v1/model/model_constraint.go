package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Constraint 条件。
type Constraint struct {

	// 条件属性，参数的某个字段值。
	Attribute *string `json:"attribute,omitempty"`

	// 操作，当前只支持equal操作。
	Operator *string `json:"operator,omitempty"`

	// 取值。
	Value *interface{} `json:"value,omitempty"`
}

func (o Constraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Constraint struct{}"
	}

	return strings.Join([]string{"Constraint", string(data)}, " ")
}
