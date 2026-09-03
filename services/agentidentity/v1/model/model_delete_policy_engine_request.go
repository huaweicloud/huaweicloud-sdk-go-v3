package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyEngineRequest Request Object
type DeletePolicyEngineRequest struct {

	// System-generated unique identifier for the policy engine.
	PolicyEngineId string `json:"policy_engine_id"`
}

func (o DeletePolicyEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyEngineRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyEngineRequest", string(data)}, " ")
}
