package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesStatisticRequest Request Object
type ListTablesStatisticRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 查询类型。固定取值。 skew：表倾斜率。 dirtyPage：表脏页率。
	RateType string `json:"rate_type"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset int32 `json:"offset"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 不限制。
	Limit int32 `json:"limit"`

	// 排序字段。固定取值。 table_size：表大小。rate：表倾斜率或脏页率。
	OrderBy *string `json:"order_by,omitempty"`

	// 正序还是倒叙。固定取值。 ASC：正序。DESC：倒序
	SortBy *string `json:"sort_by,omitempty"`

	// 查询条件。固定取值。 db_name：数据库名称。 schema_name：schema名称。 table_name：表名。 table_owner：所属用户。
	Filter *string `json:"filter,omitempty"`

	// 查询条件的取值。
	Value *string `json:"value,omitempty"`
}

func (o ListTablesStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListTablesStatisticRequest", string(data)}, " ")
}
