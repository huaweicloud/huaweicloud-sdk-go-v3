package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatamapLineageRequest Request Object
type ShowDatamapLineageRequest struct {

	// 实例id
	Instance string `json:"instance"`

	// 资产guid
	Guid string `json:"guid"`

	// 血缘查询方向，默认值:BOTH。枚举值[IN OUT BOTH]
	Direction *string `json:"direction,omitempty"`

	// 关联关系类型列表，默认空
	RelationshipTypes *[]string `json:"relationship_types,omitempty"`

	// 关联关系类型类别，默认空。血缘查询使用DATA_FLOW
	RelationshipTypeCategories *[]string `json:"relationship_type_categories,omitempty"`

	// 关联实体类型，默认空
	RelatedEntityTypes *[]string `json:"related_entity_types,omitempty"`

	// 是否扩展数据，默认false
	ExtendProcessDataFlow *bool `json:"extend_process_data_flow,omitempty"`

	// 是否包含父类实体，默认false
	IncludeParentEntity *bool `json:"include_parent_entity,omitempty"`
}

func (o ShowDatamapLineageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatamapLineageRequest struct{}"
	}

	return strings.Join([]string{"ShowDatamapLineageRequest", string(data)}, " ")
}
