package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiFieldExpression struct {

	// 多字段条件，多个相同优先级的单字段条件。
	Logic *string `bson:"logic,omitempty"`

	// 多个相同逻辑操作的单字段条件。
	Expressions []SingleFieldExpression `bson:"expressions"`
}

func (o MultiFieldExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiFieldExpression struct{}"
	}

	return strings.Join([]string{"MultiFieldExpression", string(data)}, " ")
}
