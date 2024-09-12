package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// GetKvResponse Response Object
type GetKvResponse struct {

	// 对kv_doc有效。 > 内容字段：主键字段+投影字段或者全部字段。
	KvDoc          *bson.D `bson:"kv_doc,omitempty"`
	HttpStatusCode int     `bson:"-"`
}

func (o GetKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKvResponse struct{}"
	}

	return strings.Join([]string{"GetKvResponse", string(data)}, " ")
}
