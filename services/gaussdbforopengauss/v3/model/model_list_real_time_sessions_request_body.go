package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRealTimeSessionsRequestBody struct {

	// **参数解释**： 节点ID。 **约束限制**： 仅支持非日志类型的CN或DN节点。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 组件ID。 **约束限制**： 仅支持非日志类型的CN或DN节点。 **取值范围**： 不涉及。 **默认取值**: 不涉及。
	ComponentId string `json:"component_id"`

	// **参数解释**： 是否展示本地会话。 **约束限制**： 不涉及。 **取值范围**： - true：展示本地会话 - false：不展示本地会话  **默认取值**: false。
	SupportLocal *bool `json:"support_local,omitempty"`

	QueryInfo *SessionQueryInfoOption `json:"query_info,omitempty"`
}

func (o ListRealTimeSessionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRealTimeSessionsRequestBody struct{}"
	}

	return strings.Join([]string{"ListRealTimeSessionsRequestBody", string(data)}, " ")
}
