package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComposedExpression struct {

	// 逻辑关系，取值如\"$and\", \"$or\", \"$nor\"。
	Logic *string `bson:"logic,omitempty"`

	// 多个相同优先级且相同逻辑的单字段或多字段条件。
	Expressions []Expression `bson:"expressions"`
}

func (o ComposedExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComposedExpression struct{}"
	}

	return strings.Join([]string{"ComposedExpression", string(data)}, " ")
}
