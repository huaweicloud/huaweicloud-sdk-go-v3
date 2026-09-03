package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyEngineResponse Response Object
type UpdatePolicyEngineResponse struct {
	PolicyEngine   *PolicyEngine `json:"policy_engine,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdatePolicyEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyEngineResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyEngineResponse", string(data)}, " ")
}
