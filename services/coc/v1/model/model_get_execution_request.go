package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetExecutionRequest Request Object
type GetExecutionRequest struct {
	ExecutionId string `json:"execution_id"`
}

func (o GetExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetExecutionRequest struct{}"
	}

	return strings.Join([]string{"GetExecutionRequest", string(data)}, " ")
}
