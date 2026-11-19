package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EntityResult struct {

	// **参数解释**: 实例ID或节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EntityId *string `json:"entity_id,omitempty"`

	// **参数解释**: 实例名称或节点名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EntityName *string `json:"entity_name,omitempty"`
}

func (o EntityResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityResult struct{}"
	}

	return strings.Join([]string{"EntityResult", string(data)}, " ")
}
