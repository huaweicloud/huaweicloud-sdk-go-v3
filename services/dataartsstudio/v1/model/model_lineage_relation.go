package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LineageRelation struct {

	// 血缘来源
	FromEntityId *string `json:"from_entity_id,omitempty"`

	// 关系id
	RelationshipId *string `json:"relationship_id,omitempty"`

	// 血缘流向
	ToEntityId *string `json:"to_entity_id,omitempty"`
}

func (o LineageRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineageRelation struct{}"
	}

	return strings.Join([]string{"LineageRelation", string(data)}, " ")
}
