package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"strings"
)

type UpdateBlob struct {

	// 属性信息。
	Xattr *primitive.Binary `bson:"xattr,omitempty"`

	// value部分的偏移位置。 > - 超过value当前size无效 > - \"offset\"与\"len\"与\"blob_data\" 要么都带，要么都不带。
	Offset *int32 `bson:"offset,omitempty"`

	// 更新内容长度。
	Len *int32 `bson:"len,omitempty"`

	// 二进制内容。
	BlobData *primitive.Binary `bson:"blob_data,omitempty"`
}

func (o UpdateBlob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlob struct{}"
	}

	return strings.Join([]string{"UpdateBlob", string(data)}, " ")
}
