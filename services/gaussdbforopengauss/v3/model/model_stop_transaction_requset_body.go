package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StopTransactionRequsetBody struct {

	// **参数解释**: 节点ID，仅支持非日志类型的CN或DN节点。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 组件ID，仅支持非日志类型的CN或DN节点。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ComponentId string `json:"component_id"`

	// **参数解释**: 查杀事务的ID列表。 **约束限制**: 不涉及。
	SessionIds []int32 `json:"session_ids"`
}

func (o StopTransactionRequsetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTransactionRequsetBody struct{}"
	}

	return strings.Join([]string{"StopTransactionRequsetBody", string(data)}, " ")
}
