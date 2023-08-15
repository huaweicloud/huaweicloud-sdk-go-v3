package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckConstraintInput struct {

	// 列名
	ColumnName string `json:"column_name"`

	// constraint Name
	ConstraintName string `json:"constraint_name"`

	// 检查条件表达式
	CheckExpression *string `json:"check_expression,omitempty"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o CheckConstraintInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConstraintInput struct{}"
	}

	return strings.Join([]string{"CheckConstraintInput", string(data)}, " ")
}
