package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePoliciesPriorityResponse Response Object
type BatchUpdatePoliciesPriorityResponse struct {

	// **参数解释**：请求ID。  **取值范围**：不涉及
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdatePoliciesPriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityResponse", string(data)}, " ")
}
