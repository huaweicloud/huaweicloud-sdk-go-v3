package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLineageResponse Response Object
type ShowLineageResponse struct {

	// 当前资产的guid
	BaseEntityGuid *string `json:"base_entity_guid,omitempty"`

	// 实体集合Map(String, OpenEntityHeader)
	GuidEntityMap *interface{} `json:"guid_entity_map,omitempty"`

	// 血缘关系
	Relations *[]LineageRelation `json:"relations,omitempty"`

	// 相关实体集合Map(String, OpenEntity)
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowLineageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLineageResponse struct{}"
	}

	return strings.Join([]string{"ShowLineageResponse", string(data)}, " ")
}
