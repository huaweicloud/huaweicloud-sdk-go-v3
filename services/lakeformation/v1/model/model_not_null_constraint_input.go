package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotNullConstraintInput struct {

	// 列名
	ColumnName string `json:"column_name"`

	// constraint Name
	ConstraintName string `json:"constraint_name"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o NotNullConstraintInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotNullConstraintInput struct{}"
	}

	return strings.Join([]string{"NotNullConstraintInput", string(data)}, " ")
}
