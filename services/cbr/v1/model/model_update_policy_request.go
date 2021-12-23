package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePolicyRequest struct {
	// 策略ID

	PolicyId string `json:"policy_id"`

	Body *PolicyUpdateReq `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
