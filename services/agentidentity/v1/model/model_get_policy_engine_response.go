package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyEngineResponse Response Object
type GetPolicyEngineResponse struct {
	PolicyEngine   *PolicyEngine `json:"policy_engine,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o GetPolicyEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyEngineResponse struct{}"
	}

	return strings.Join([]string{"GetPolicyEngineResponse", string(data)}, " ")
}
