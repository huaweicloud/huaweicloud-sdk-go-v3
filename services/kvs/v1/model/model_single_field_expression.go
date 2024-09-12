package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type SingleFieldExpression struct {

	// 条件字段。
	Field string `bson:"field"`

	// 条件函数，取值如\"$gt\", $lt\",\"$gte\", $lte\" \"$eq\", \"$ne\", \"$prefix\", \"$exists\"。
	Func string `bson:"func"`

	// value和value_array二选一。 - value条件值，适用于除\"$in\", \"$nin\"外的func。 - 字段名无意义，可以传空，也可以传字段名。 - $exists值为true/false。 > $prefix操作只适用于string和binary类型。
	Value *bson.D `bson:"value,omitempty"`

	// \"value\"和\"value_array\"二选一。 - \"value_array\" 条件值列表, 值用于\"$in\", \"$nin\"。
	ValueArray *[]bson.D `bson:"value_array,omitempty"`
}

func (o SingleFieldExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleFieldExpression struct{}"
	}

	return strings.Join([]string{"SingleFieldExpression", string(data)}, " ")
}
