package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdatePoliciesPriorityRequest struct {
	Body *BatchUpdatePoliciesPriorityRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchUpdatePoliciesPriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequest", string(data)}, " ")
}
