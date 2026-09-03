package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyEngineRequest Request Object
type GetPolicyEngineRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`
}

func (o GetPolicyEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyEngineRequest struct{}"
	}

	return strings.Join([]string{"GetPolicyEngineRequest", string(data)}, " ")
}
