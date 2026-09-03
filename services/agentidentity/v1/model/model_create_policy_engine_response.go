package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyEngineResponse Response Object
type CreatePolicyEngineResponse struct {
	PolicyEngine   *PolicyEngine `json:"policy_engine,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreatePolicyEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyEngineResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyEngineResponse", string(data)}, " ")
}
