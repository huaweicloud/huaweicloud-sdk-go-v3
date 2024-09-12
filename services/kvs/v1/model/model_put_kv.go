package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type PutKv struct {

	// 请求内的操作编码，未成功的操作返回该标识。
	OperId int32 `bson:"oper_id"`

	// 用户文档。
	KvDoc *bson.D `bson:"kv_doc,omitempty"`
}

func (o PutKv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutKv struct{}"
	}

	return strings.Join([]string{"PutKv", string(data)}, " ")
}
