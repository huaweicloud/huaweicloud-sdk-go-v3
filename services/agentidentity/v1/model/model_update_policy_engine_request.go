package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyEngineRequest Request Object
type UpdatePolicyEngineRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`

	Body *UpdatePolicyEngineReqBody `json:"body,omitempty"`
}

func (o UpdatePolicyEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyEngineRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyEngineRequest", string(data)}, " ")
}
