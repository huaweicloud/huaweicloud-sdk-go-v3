package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisReportForRocketMqResponse Response Object
type ShowDiagnosisReportForRocketMqResponse struct {

	// **参数解释**： 报告ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 消费者数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumerNums *int32 `json:"consumer_nums,omitempty"`

	// **参数解释**： 状态。 **约束限制**： 不涉及。 **取值范围**： - diagnosing：诊断中。 - failed：诊断失败。 - deleted：手动删除。 - finished：诊断完成。 - normal：诊断结果正常。 - abnormal：诊断结果异常。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 生成时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**： 异常项数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`

	// **参数解释**： 异常节点数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FaultedNodeSum *int32 `json:"faulted_node_sum,omitempty"`

	// **参数解释**： 是否在线。 **取值范围**： - True：在线。 - False：不在线。
	Online *bool `json:"online,omitempty"`

	// **参数解释**： 消息堆积数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MessageAccumulation *int32 `json:"message_accumulation,omitempty"`

	// **参数解释**： 订阅一致性。 **约束限制**： 不涉及。 **取值范围**： - True：订阅关系一致。 - False：订阅关系不一致。 **默认取值**： 不涉及。
	SubscriptionConsistency *bool `json:"subscription_consistency,omitempty"`

	// **参数解释**： 是否存在重复的客户端ID。 **约束限制**： 不涉及。 **取值范围**： - True：存在重复的客户端ID。 - False：不存在重复的客户端ID。 **默认取值**： 不涉及。
	DuplicateClientId *bool `json:"duplicate_client_id,omitempty"`

	// **参数解释**： 是否存在不一致的消费类型。 **约束限制**： 不涉及。 **取值范围**： - True：存在不一致的消费类型。 - False：不存在不一致的消费类型。 **默认取值**： 不涉及。
	DifferentConsumerType *bool `json:"different_consumer_type,omitempty"`

	// **参数解释**： 订阅者列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Subscriptions *[]SubscriptionEntity `json:"subscriptions,omitempty"`

	// **参数解释**： 诊断节点报告列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DiagnosisNodeReportList *[]DiagnosisNodeReportEntity `json:"diagnosis_node_report_list,omitempty"`
	HttpStatusCode          int                          `json:"-"`
}

func (o ShowDiagnosisReportForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisReportForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisReportForRocketMqResponse", string(data)}, " ")
}
