package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type ScanKvRequestBody struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// create_table时指定的索引名，默认空表示主索引。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	HintIndexName *string `bson:"hint_index_name,omitempty"`

	// 数据量不超过1MB时，返回的文档个数，最大100个，默认1MB或者100个文档。
	Limit *int32 `bson:"limit,omitempty"`

	// 起始主键或索引键值。 - 默认空，表示从头遍历，左闭。 > 分页返回时，该值使用上次响应返回的cursor_key。
	StartKey *bson.D `bson:"start_key,omitempty"`

	// 终止主键或索引键值。 - 默认空，表示直到最后，右开。
	EndKey *bson.D `bson:"end_key,omitempty"`

	FilterExpression *ConditionExpression `bson:"filter_expression,omitempty"`

	// 对表进行采样，尽最大努力保证返回的段列表均分整张表。举例：sample_segments_count=4，返回的段列表[MinKey, KV1)、[KV1,KV2)、[KV2,KV3)和[KV3,MaxKey)。用户可以使用scan-kv对这四个分区执行并发扫描，提高遍历效率。 - 范围: [1, 10000]。默认值：不执行采样。 - sample_segments_count仅能和table_name、start_key和end_key字段配合使用。Range分区模式下支持全表采样和范围采样；Hash分区模式仅支持全表扫描。 - 仅支持对Primary key进行采样，不支持本地/全局二级索引。 - 返回的段列表仅包含主键，不包含键值；且段列表是编码后的数据 ，不可修改。
	SampleSegmentsCount *int32 `bson:"sample_segments_count,omitempty"`

	// 返回查询条件对应的KV总数. - 当KV总数小于limit条件时，返回KV查询结果和KV总数。 - 当KV总数多于limit条件时，只返回KV总数。
	ReturnCountOnly *bool `bson:"return_count_only,omitempty"`
}

func (o ScanKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanKvRequestBody struct{}"
	}

	return strings.Join([]string{"ScanKvRequestBody", string(data)}, " ")
}
