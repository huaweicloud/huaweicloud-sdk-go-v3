package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyRequest Request Object
type UpdatePolicyRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`

	// System-generated unique identifier for the policy.
	PolicyId string `json:"policy_id"`

	Body *UpdatePolicyReqBody `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
