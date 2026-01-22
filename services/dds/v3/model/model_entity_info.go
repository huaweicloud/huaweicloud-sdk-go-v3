package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EntityInfo struct {

	// **参数解释：** 组ID或节点ID。 **取值范围：** 不涉及。
	EntityId string `json:"entity_id"`

	// **参数解释：** 组名称或节点名称。 **取值范围：** 不涉及。
	EntityName string `json:"entity_name"`
}

func (o EntityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityInfo struct{}"
	}

	return strings.Join([]string{"EntityInfo", string(data)}, " ")
}
