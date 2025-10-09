package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessComputeAbilityPolicyResponse Response Object
type UpdateServerlessComputeAbilityPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerlessComputeAbilityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessComputeAbilityPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerlessComputeAbilityPolicyResponse", string(data)}, " ")
}
