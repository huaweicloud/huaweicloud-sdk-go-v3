package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseObjectsRequest Request Object
type ListDatabaseObjectsRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 对象类型。 **约束限制**： 不涉及。 **取值范围**： DATABASE、SCHEMA、TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE、NODEGROUP **默认取值**： null
	Type string `json:"type"`

	// **参数解释**： 对象名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 数据库名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**： 模式名。 **约束限制**： 当对象类型为TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE时必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Schema *string `json:"schema,omitempty"`

	// **参数解释**： 表名。 **约束限制**： 对象类型为COLUMN时必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Table *string `json:"table,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 是否细粒度容灾。 **约束限制**： 不涉及。 **取值范围**： true|false **默认取值**： 不涉及。
	IsFineGrainedDisaster *string `json:"is_fine_grained_disaster,omitempty"`
}

func (o ListDatabaseObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseObjectsRequest", string(data)}, " ")
}
