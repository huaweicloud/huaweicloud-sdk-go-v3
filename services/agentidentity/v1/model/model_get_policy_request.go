package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyRequest Request Object
type GetPolicyRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`

	// System-generated unique identifier for the policy.
	PolicyId string `json:"policy_id"`
}

func (o GetPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyRequest struct{}"
	}

	return strings.Join([]string{"GetPolicyRequest", string(data)}, " ")
}
