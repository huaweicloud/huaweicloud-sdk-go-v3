package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateAgentInvocationsRequest Request Object
type BatchCreateAgentInvocationsRequest struct {
	Body *BatchCreateInvocationRequestBody `json:"body,omitempty"`
}

func (o BatchCreateAgentInvocationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAgentInvocationsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateAgentInvocationsRequest", string(data)}, " ")
}
