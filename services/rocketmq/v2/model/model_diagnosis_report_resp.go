package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisReportResp struct {

	// **参数解释**： 报告ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 消费者数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumerNums *int32 `json:"consumer_nums,omitempty"`

	// **参数解释**： 状态。 **约束限制**： 不涉及。 **取值范围**： - diagnosing：诊断中。 - failed：诊断失败。 - deleted：手动删除。 - finished：诊断完成。 - normal：诊断结果正常。 - abnormal：诊断结果异常。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 异常项数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`

	// **参数解释**： 异常节点数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FaultedNodeSum *int32 `json:"faulted_node_sum,omitempty"`
}

func (o DiagnosisReportResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisReportResp struct{}"
	}

	return strings.Join([]string{"DiagnosisReportResp", string(data)}, " ")
}
