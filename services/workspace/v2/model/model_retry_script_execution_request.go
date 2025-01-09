package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryScriptExecutionRequest Request Object
type RetryScriptExecutionRequest struct {
	Body *RetryScriptExecutionReq `json:"body,omitempty"`
}

func (o RetryScriptExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryScriptExecutionRequest struct{}"
	}

	return strings.Join([]string{"RetryScriptExecutionRequest", string(data)}, " ")
}
