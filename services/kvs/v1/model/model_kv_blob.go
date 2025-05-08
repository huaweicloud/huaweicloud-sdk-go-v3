package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"strings"
)

type KvBlob struct {

	// 用户自定义的主键名及值。 > 内容字段：主键字段名和值，组合索引多个元素。
	PrimaryKey *bson.D `bson:"primary_key"`

	// 属性信息，最大2kb。
	Xblob *primitive.Binary `bson:"xblob,omitempty"`

	// 非结构化数据内容。
	Xattr *primitive.Binary `bson:"xattr,omitempty"`
}

func (o KvBlob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KvBlob struct{}"
	}

	return strings.Join([]string{"KvBlob", string(data)}, " ")
}
