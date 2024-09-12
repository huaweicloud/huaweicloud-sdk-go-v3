package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConditionExpression struct {
	SingleFieldExpression *SingleFieldExpression `bson:"single_field_expression,omitempty"`

	MultiFieldExpression *MultiFieldExpression `bson:"multi_field_expression,omitempty"`

	ComposedExpression *ComposedExpression `bson:"composed_expression,omitempty"`
}

func (o ConditionExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionExpression struct{}"
	}

	return strings.Join([]string{"ConditionExpression", string(data)}, " ")
}
