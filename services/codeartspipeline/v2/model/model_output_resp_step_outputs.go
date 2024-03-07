package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputRespStepOutputs struct {
	StepRunId *string `json:"step_run_id,omitempty"`

	OutputResult *[]OutputRespOutputResult `json:"output_result,omitempty"`
}

func (o OutputRespStepOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputRespStepOutputs struct{}"
	}

	return strings.Join([]string{"OutputRespStepOutputs", string(data)}, " ")
}
