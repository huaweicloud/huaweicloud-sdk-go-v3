package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisReportResponse Response Object
type ShowDiagnosisReportResponse struct {

	// **参数解释**： 报告ID。 **取值范围**： 不涉及。
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 消费组名称。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 消费者数量。 **取值范围**： 不涉及。
	ConsumerNums *int32 `json:"consumer_nums,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 生成时间。 **取值范围**： 不涉及。
	CreatAt *int64 `json:"creat_at,omitempty"`

	// **参数解释**： 异常项数量。 **取值范围**： 不涉及。
	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`

	// **参数解释**： 异常节点数量。 **取值范围**： 不涉及。
	FaultedNodeSum *int32 `json:"faulted_node_sum,omitempty"`

	// **参数解释**： 是否在线。 **取值范围**： - true: 在线。 - false: 不在线。
	Online *bool `json:"online,omitempty"`

	// **参数解释**： 消息堆积数。 **取值范围**： 不涉及。
	MessageAccumulation *int32 `json:"message_accumulation,omitempty"`

	// **参数解释**： 订阅一致性。 **取值范围**： - true: 订阅关系一致。 - false: 订阅关系不一致。
	SubscriptionConsistency *bool `json:"subscription_consistency,omitempty"`

	// **参数解释**： 订阅者列表。 **取值范围**： 不涉及。
	Subscriptions *[]string `json:"subscriptions,omitempty"`

	// **参数解释**： 诊断节点报告列表。 **取值范围**： 不涉及。
	DiagnosisNodeReportList *[]string `json:"diagnosis_node_report_list,omitempty"`
	HttpStatusCode          int       `json:"-"`
}

func (o ShowDiagnosisReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisReportResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisReportResponse", string(data)}, " ")
}
