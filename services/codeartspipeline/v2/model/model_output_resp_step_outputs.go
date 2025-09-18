package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputRespStepOutputs struct {

	// **参数解释**： 步骤ID。 **取值范围**： 32位字符，由数字和字母组成。
	StepRunId *string `json:"step_run_id,omitempty"`

	// **参数解释**： 步骤输出。 **取值范围**： 不涉及。
	OutputResult *[]OutputRespOutputResult `json:"output_result,omitempty"`
}

func (o OutputRespStepOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputRespStepOutputs struct{}"
	}

	return strings.Join([]string{"OutputRespStepOutputs", string(data)}, " ")
}
