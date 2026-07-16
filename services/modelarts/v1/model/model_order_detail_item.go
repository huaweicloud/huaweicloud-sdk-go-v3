package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderDetailItem 订单关联的资源信息。
type OrderDetailItem struct {

	// **参数解释**：资源的ID，取值自资源详情的metadata.name字段。 **取值范围**：不涉及。
	ResourceName *string `json:"resourceName,omitempty"`

	// **参数解释**：订单关联的资源变更动作类型。 **取值范围**：可选值如下： - createPool：创建资源池。 - deletePool：删除资源池。 - createNode：创建节点。 - deleteNode：删除节点，主要是包周期节点退订场景。 - renew：续费。 - toPeriodic：转包周期。
	Action string `json:"action"`

	// **参数解释**：订单关联资源的变更状态。 **取值范围**：可选值如下： - processing：处理中，资源正在处理中。 - succeed：成功，资源处理成功。 - failed：失败，资源处理失败。
	Status string `json:"status"`

	// **参数解释**：资源开始变更时间戳，形如1744285793000，单位毫秒。 **取值范围**：不涉及。
	BeginTimestamp *string `json:"beginTimestamp,omitempty"`

	// **参数解释**：资源变更最后更新时间戳，形如1744285793000，单位毫秒。 **取值范围**：不涉及。
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	// **参数解释**：资源变更的执行信息，如失败原因。 **取值范围**：不涉及。
	Message *string `json:"message,omitempty"`
}

func (o OrderDetailItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderDetailItem struct{}"
	}

	return strings.Join([]string{"OrderDetailItem", string(data)}, " ")
}
