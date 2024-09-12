package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type PutKvRequestBody struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	ConditionExpression *ConditionExpression `bson:"condition_expression,omitempty"`

	// 用户文档。
	KvDoc *bson.D `bson:"kv_doc,omitempty"`
}

func (o PutKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutKvRequestBody struct{}"
	}

	return strings.Join([]string{"PutKvRequestBody", string(data)}, " ")
}
