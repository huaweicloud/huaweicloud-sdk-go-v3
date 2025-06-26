package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesStatisticRequest Request Object
type ListTablesStatisticRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 查询类型，固定取值。 **约束限制**： 不涉及。 **取值范围**： skew：表倾斜率。 dirtyPage：表脏页率。 **默认取值**： 不涉及。
	RateType string `json:"rate_type"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset int32 `json:"offset"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 不限制。
	Limit int32 `json:"limit"`

	// **参数解释**： 排序字段，固定取值。 **约束限制**： 不涉及。 **取值范围**： table_size：表大小。 rate：表倾斜率或脏页率。 **默认取值**： 不涉及。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**： 正序还是倒序排序，固定取值。 **约束限制**： 不涉及。 **取值范围**： ASC：正序。 DESC：倒序。 **默认取值**： 不涉及。
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释**： 查询条件，固定取值。 **约束限制**： 不涉及。 **取值范围**： db_name：数据库名称。 schema_name：schema名称。 table_name：表名。 table_owner：所属用户。 **默认取值**： 不涉及。
	Filter *string `json:"filter,omitempty"`

	// **参数解释**： 过滤条件的值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o ListTablesStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListTablesStatisticRequest", string(data)}, " ")
}
