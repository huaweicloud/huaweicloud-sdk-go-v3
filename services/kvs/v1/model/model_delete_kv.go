package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type DeleteKv struct {

	// 请求内的操作编码，未成功的操作返回该标识。
	OperId int32 `bson:"oper_id"`

	// 用户自定义的主键名及值。
	PrimaryKey *bson.D `bson:"primary_key"`
}

func (o DeleteKv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKv struct{}"
	}

	return strings.Join([]string{"DeleteKv", string(data)}, " ")
}
