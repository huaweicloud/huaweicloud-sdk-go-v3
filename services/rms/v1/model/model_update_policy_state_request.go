package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePolicyStateRequest struct {
	Body *PolicyStateRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyStateRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyStateRequest", string(data)}, " ")
}
