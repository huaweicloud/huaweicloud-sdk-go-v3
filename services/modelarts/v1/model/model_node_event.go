package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeEvent **参数解释**：事件关联节点信息列表详情。 **取值范围**：不涉及。
type NodeEvent struct {

	// **参数解释**：Iaas计划事件Id。 **取值范围**：不涉及。
	EventId *string `json:"eventId,omitempty"`

	// **参数解释**：是否故障节点。 **取值范围**：不涉及。
	IsFaultNode *bool `json:"isFaultNode,omitempty"`

	// **参数解释**：节点名称。 **取值范围**：不涉及。
	NodeName *string `json:"nodeName,omitempty"`

	// **参数解释**：节点ID。 **取值范围**：不涉及。
	InstanceId *string `json:"instanceId,omitempty"`

	// **参数解释**：实例类型。 **取值范围**：不涉及。
	InstanceType *string `json:"instanceType,omitempty"`

	// **参数解释**：底层实例Id。 **取值范围**：不涉及。
	UnderLyingInstanceId *string `json:"underLyingInstanceId,omitempty"`

	// **参数解释**：节点规格是否支持重部署。 **取值范围**：不涉及。
	FlavorSupportRedeploy *bool `json:"flavorSupportRedeploy,omitempty"`

	// **参数解释**：资源池名称。 **取值范围**：不涉及。
	PoolName *string `json:"poolName,omitempty"`

	// **参数解释**：资源池显示名称。 **取值范围**：不涉及。
	PoolDisplayName *string `json:"poolDisplayName,omitempty"`

	// **参数解释**：Iaas计划事件状态。 **取值范围**：不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**：Iaas事件发布时间。 **取值范围**：不涉及。
	PublishTime *string `json:"publishTime,omitempty"`

	// **参数解释**：Iaas事件开始时间。 **取值范围**：不涉及。
	StartTime *string `json:"startTime,omitempty"`

	// **参数解释**：Iaas事件完成时间。 **取值范围**：不涉及。
	FinishTime *string `json:"finishTime,omitempty"`

	// **参数解释**：重部署进度。 **取值范围**：不涉及。
	JobProgress *string `json:"jobProgress,omitempty"`
}

func (o NodeEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeEvent struct{}"
	}

	return strings.Join([]string{"NodeEvent", string(data)}, " ")
}
