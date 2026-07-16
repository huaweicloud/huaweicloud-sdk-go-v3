package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StageInfoWithSub 作业流程阶段信息列表的主阶段元信息，包含子阶段。
type StageInfoWithSub struct {

	// **参数解释**：作业ID。 **取值范围**：不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：主阶段名称。  **取值范围**： - scheduling：作业调度 - preparing：环境准备 - running：作业运行 - end：作业结束
	Name *string `json:"name,omitempty"`

	// **参数解释**：主阶段英文描述信息。  **取值范围**：不涉及。
	EnMessage *string `json:"en_message,omitempty"`

	// **参数解释**：主阶段中文描述信息。  **取值范围**：不涉及。
	ZhMessage *string `json:"zh_message,omitempty"`

	// **参数解释**：主阶段开始时间。  **取值范围**：不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：主阶段结束时间。  **取值范围**：不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：主阶段序号。 **取值范围**：[1,4]。
	StageOrder *int32 `json:"stage_order,omitempty"`

	// **参数解释**：子阶段信息列表。
	SubStages *[]SubStage `json:"sub_stages,omitempty"`
}

func (o StageInfoWithSub) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StageInfoWithSub struct{}"
	}

	return strings.Join([]string{"StageInfoWithSub", string(data)}, " ")
}
