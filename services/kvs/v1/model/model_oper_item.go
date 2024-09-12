package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperItem put_kv或delete_kv或get_kv操作。
type OperItem struct {
	PutKv *PutKv `bson:"put_kv,omitempty"`

	DeleteKv *DeleteKv `bson:"delete_kv,omitempty"`
}

func (o OperItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperItem struct{}"
	}

	return strings.Join([]string{"OperItem", string(data)}, " ")
}
