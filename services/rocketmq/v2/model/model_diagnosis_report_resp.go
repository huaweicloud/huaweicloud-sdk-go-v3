package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisReportResp struct {

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
}

func (o DiagnosisReportResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisReportResp struct{}"
	}

	return strings.Join([]string{"DiagnosisReportResp", string(data)}, " ")
}
