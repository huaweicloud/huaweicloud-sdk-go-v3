package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type GetKvRequestBody struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// 用户自定义的主键名及值。
	PrimaryKey *bson.D `bson:"primary_key"`
}

func (o GetKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKvRequestBody struct{}"
	}

	return strings.Join([]string{"GetKvRequestBody", string(data)}, " ")
}
