package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type GetKv struct {

	// 请求内的操作编码，未成功的操作返回该标识。
	OperId int32 `bson:"oper_id"`

	// 用户自定义的主键名及值。
	PrimaryKey *bson.D `bson:"primary_key"`
}

func (o GetKv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKv struct{}"
	}

	return strings.Join([]string{"GetKv", string(data)}, " ")
}
