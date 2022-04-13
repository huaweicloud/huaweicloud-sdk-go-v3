package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdatePoliciesPriorityRequest struct {
	Body *BatchUpdatePoliciesPriorityRequestBody `json:"body,omitempty"`
}

func (o BatchUpdatePoliciesPriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequest", string(data)}, " ")
}
