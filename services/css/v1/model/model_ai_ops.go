package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiOps struct {

	// **参数解释**： 集群风险检测任务ID。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 检测范围 **取值范围**： - full_detection：全量检测 - unavailability_detection：集群不可用检测 - partial_detection：部分检测
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释**： 触发类型 **取值范围**： - manual：手动 - auto：自动
	TriggerType *string `json:"trigger_type,omitempty"`

	// **参数解释**： 集群风险检测任务名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 集群风险检测任务描述。 **取值范围**： 不涉及
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 任务执行状态。 **取值范围**： - 150：未开启。 - 200：已开启。 - 300：已发送。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 集群风险检测任务创建时间。 **取值范围**： 格式：Unix时间戳格式。
	CreateTime *string `json:"create_time,omitempty"`

	Summary *SummaryInfo `json:"summary,omitempty"`

	// **参数解释**： 集群风险项详情。 **取值范围**： 不涉及
	TaskRisks *[]AiOpsRiskObject `json:"task_risks,omitempty"`
}

func (o AiOps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOps struct{}"
	}

	return strings.Join([]string{"AiOps", string(data)}, " ")
}
