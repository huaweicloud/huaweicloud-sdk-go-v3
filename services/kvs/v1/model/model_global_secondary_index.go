package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalSecondaryIndex 全局二级索引定义，元素key为数组下标
type GlobalSecondaryIndex struct {

	// 二级索引名称，表内唯一。
	IndexName string `bson:"index_name"`

	// 分区键字段名数组，顺序组合。
	ShardKeyFields []Field `bson:"shard_key_fields"`

	// 分区模式。
	ShardMode *string `bson:"shard_mode,omitempty"`

	// 排序键字段名数组，顺序组合。
	SortKeyFields *[]Field `bson:"sort_key_fields,omitempty"`

	// 摘要字段名数组。
	AbstractFields *[]string `bson:"abstract_fields,omitempty"`
}

func (o GlobalSecondaryIndex) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalSecondaryIndex struct{}"
	}

	return strings.Join([]string{"GlobalSecondaryIndex", string(data)}, " ")
}
