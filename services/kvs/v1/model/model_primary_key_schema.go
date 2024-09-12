package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrimaryKeySchema struct {

	// 分区键字段名数组，顺序组合。
	ShardKeyFields []Field `bson:"shard_key_fields"`

	// 分区模式。
	ShardMode *string `bson:"shard_mode,omitempty"`

	// 排序键字段名数组，顺序组合。
	SortKeyFields *[]Field `bson:"sort_key_fields,omitempty"`
}

func (o PrimaryKeySchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrimaryKeySchema struct{}"
	}

	return strings.Join([]string{"PrimaryKeySchema", string(data)}, " ")
}
