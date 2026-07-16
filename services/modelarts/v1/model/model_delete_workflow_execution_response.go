package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowExecutionResponse Response Object
type DeleteWorkflowExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowExecutionResponse", string(data)}, " ")
}
