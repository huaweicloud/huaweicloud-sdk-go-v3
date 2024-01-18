package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatamapLineageResponse Response Object
type ShowDatamapLineageResponse struct {

	// 查询命中总条数
	Total float32 `json:"total,omitempty"`

	// 关系列表
	Relationships *[]Lineage `json:"relationships,omitempty"`

	// 关系类型
	RelationshipTypes *interface{} `json:"relationship_types,omitempty"`

	// 资产信息
	Entities *[]Entity `json:"entities,omitempty"`

	// 实体类型
	EntityTypes    *interface{} `json:"entity_types,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDatamapLineageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatamapLineageResponse struct{}"
	}

	return strings.Join([]string{"ShowDatamapLineageResponse", string(data)}, " ")
}
