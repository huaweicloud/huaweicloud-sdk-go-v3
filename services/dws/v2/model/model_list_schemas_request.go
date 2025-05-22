package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasRequest Request Object
type ListSchemasRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// **参数解释**： 排序字段。 **约束限制**： 不涉及。 **取值范围**： schemaName：模式名称排序。 **默认取值**： 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序字段。 **约束限制**： 不涉及。 **取值范围**： ASC：表示按升序排序。  DESC：表示按降序排序。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 查询关键词。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Keywords *string `json:"keywords,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasRequest struct{}"
	}

	return strings.Join([]string{"ListSchemasRequest", string(data)}, " ")
}
