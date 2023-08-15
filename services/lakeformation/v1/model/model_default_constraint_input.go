package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefaultConstraintInput struct {

	// 列名
	ColumnName string `json:"column_name"`

	// constraint Name
	ConstraintName string `json:"constraint_name"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o DefaultConstraintInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultConstraintInput struct{}"
	}

	return strings.Join([]string{"DefaultConstraintInput", string(data)}, " ")
}
