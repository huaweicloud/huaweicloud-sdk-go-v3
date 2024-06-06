package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStepOutputsResponse Response Object
type ShowStepOutputsResponse struct {
	StepOutputs *[]OutputRespStepOutputs `json:"step_outputs,omitempty"`

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
