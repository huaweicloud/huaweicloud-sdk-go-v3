package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCesHierarchyResponse Response Object
type ShowCesHierarchyResponse struct {

	// **参数解释**： 监控维度。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Dimensions *[]ShowCeshierarchyRespDimensions `json:"dimensions,omitempty"`

	// **参数解释**： 实例信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceIds *[]ShowCesHierarchyRespInstanceIds `json:"instance_ids,omitempty"`

	// **参数解释**： 节点信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Nodes *[]ShowCesHierarchyRespNodes `json:"nodes,omitempty"`

	// **参数解释**： 队列信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topics *[]ShowCeshierarchyRespTopics `json:"topics,omitempty"`

	// **参数解释**： 死信队列。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Dlq *[]ShowCeshierarchyRespDlq `json:"dlq,omitempty"`

	// **参数解释**： 消费组信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups         *[]ShowCeshierarchyRespGroups `json:"groups,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCesHierarchyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyResponse struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyResponse", string(data)}, " ")
}
