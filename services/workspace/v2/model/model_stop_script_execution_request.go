package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopScriptExecutionRequest Request Object
type StopScriptExecutionRequest struct {
	Body *StopScriptExecutionReq `json:"body,omitempty"`
}

func (o StopScriptExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopScriptExecutionRequest struct{}"
	}

	return strings.Join([]string{"StopScriptExecutionRequest", string(data)}, " ")
}
