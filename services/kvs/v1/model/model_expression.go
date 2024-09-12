package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Expression struct {
	SingleFieldExpression *SingleFieldExpression `bson:"single_field_expression,omitempty"`

	MultiFieldExpression *MultiFieldExpression `bson:"multi_field_expression,omitempty"`

	ComposedExpression *ComposedExpression `bson:"composed_expression,omitempty"`
}

func (o Expression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Expression struct{}"
	}

	return strings.Join([]string{"Expression", string(data)}, " ")
}
