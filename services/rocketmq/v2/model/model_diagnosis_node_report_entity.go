package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisNodeReportEntity struct {

	// **参数解释**： 节点ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**： 是否故障。 **约束限制**： 不涉及。 **取值范围**： - true：故障 - false：没有故障。 **默认取值**： 不涉及。
	IsFaulted *bool `json:"is_faulted,omitempty"`

	// **参数解释**： 异常项总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`

	// **参数解释**： 消息堆积数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MessageAccumulation *int32 `json:"message_accumulation,omitempty"`

	// **参数解释**： 是否为死锁。 **约束限制**： 不涉及。 **取值范围**： - true：是死锁。 - false：不是死锁。 **默认取值**： 不涉及。
	DeadLock *bool `json:"dead_lock,omitempty"`

	// **参数解释**： 死锁线程。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DeadlockThread *string `json:"deadlock_thread,omitempty"`

	// **参数解释**： 线程ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StackId *string `json:"stack_id,omitempty"`

	// **参数解释**： 是否为pop消费模式。 **约束限制**： 不涉及。 **取值范围**： - true：是pop消费模式。 - false：不是pop消费模式。 **默认取值**： 不涉及。
	IsPop *bool `json:"is_pop,omitempty"`

	// **参数解释**： 消费类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumeType *string `json:"consume_type,omitempty"`
}

func (o DiagnosisNodeReportEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisNodeReportEntity struct{}"
	}

	return strings.Join([]string{"DiagnosisNodeReportEntity", string(data)}, " ")
}
