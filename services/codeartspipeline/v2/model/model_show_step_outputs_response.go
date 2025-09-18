package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStepOutputsResponse Response Object
type ShowStepOutputsResponse struct {

	// **参数解释**： 步骤输出列表，记录每个步骤输出的ID和结果信息。 **约束限制**： 不涉及。
	StepOutputs *[]OutputRespStepOutputs `json:"step_outputs,omitempty"`

	// **参数解释**： 当前系统时间。 **取值范围**： 不涉及。
	CurrentSystemTime *int64 `json:"current_system_time,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowStepOutputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStepOutputsResponse struct{}"
	}

	return strings.Join([]string{"ShowStepOutputsResponse", string(data)}, " ")
}
