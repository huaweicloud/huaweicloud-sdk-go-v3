package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableVolumeResult struct {

	// **参数解释**: 表的大小。 **取值范围**: 不涉及。
	TableSize *string `json:"table_size,omitempty"`

	// **参数解释**: 表ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 表名称。 **取值范围**: 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **取值范围**: 不涉及。
	TableOwner *string `json:"table_owner,omitempty"`

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**: 表或者索引是否具有分区表的性质。 **取值范围**: 不涉及。
	IsPartType *bool `json:"is_part_type,omitempty"`

	// **参数解释**: 是否包含hash分区列信息。 **取值范围**: 不涉及。
	IsHashClusterKey *bool `json:"is_hash_cluster_key,omitempty"`

	// **参数解释**: 表中行的数目。 **取值范围**: 不涉及。
	Tuples *string `json:"tuples,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**: 修改时间。 **取值范围**: 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**: 表大小平均值(totalsize/DN个数，该值为平均分布的理想情况下，表在各DN占用空间大小)。 **取值范围**: 不涉及。
	AverageSize *string `json:"average_size,omitempty"`

	// **参数解释**: 单DN表大小最大值占比（表在各DN占用空间的最大值/totalsize）。 **取值范围**: 不涉及。
	MaxRatio *string `json:"max_ratio,omitempty"`

	// **参数解释**: 单DN表大小最小值占比（表在各DN占用空间的最小值/totalsize）。 **取值范围**: 不涉及。
	MinRatio *string `json:"min_ratio,omitempty"`

	// **参数解释**: 表分布倾斜值（单DN表大小最大值 - 单DN表大小最小值）。 **取值范围**: 不涉及。
	SkewSize *string `json:"skew_size,omitempty"`

	// **参数解释**: 表分布倾斜率（skewsize/totalsize）。 **取值范围**: 不涉及。
	SkewRatio *string `json:"skew_ratio,omitempty"`

	// **参数解释**: 表分布标准方差（在表大小一定的情况下，该值越大表明表的整体分布情况越倾斜）。。 **取值范围**: 不涉及。
	SkewStddev *string `json:"skew_stddev,omitempty"`
}

func (o TableVolumeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableVolumeResult struct{}"
	}

	return strings.Join([]string{"TableVolumeResult", string(data)}, " ")
}
