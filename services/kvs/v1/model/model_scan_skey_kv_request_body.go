package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type ScanSkeyKvRequestBody struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// create_table时指定的索引名。
	HintIndexName *string `bson:"hint_index_name,omitempty"`

	// 数据量不超过1MB时，返回的文档个数，最大100个，默认1MB或者100个文档。
	Limit *int32 `bson:"limit,omitempty"`

	// 要遍历的指定分区键下的kv。
	ShardKey *bson.D `bson:"shard_key"`

	// 起始排序键值，默认空表示从头遍历，左闭。 > 分页返回时，该值使用上次响应返回的cursor_sort_key
	StartSortKey *bson.D `bson:"start_sort_key,omitempty"`

	// 终止主键或索引键值，默认空表示直到最后，右开。
	EndSortKey *bson.D `bson:"end_sort_key,omitempty"`

	FilterExpression *ConditionExpression `bson:"filter_expression,omitempty"`

	// 返回查询条件对应的KV总数. - 当KV总数小于limit条件时，返回KV查询结果和KV总数。 - 当KV总数多于limit条件时，只返回KV总数。
	ReturnCountOnly *bool `bson:"return_count_only,omitempty"`
}

func (o ScanSkeyKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanSkeyKvRequestBody struct{}"
	}

	return strings.Join([]string{"ScanSkeyKvRequestBody", string(data)}, " ")
}
