package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

// ReturnedKvItem kv元素。
type ReturnedKvItem struct {

	// 对kv_doc有效。
	KvDoc *bson.D `bson:"kv_doc,omitempty"`
}

func (o ReturnedKvItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReturnedKvItem struct{}"
	}

	return strings.Join([]string{"ReturnedKvItem", string(data)}, " ")
}
