package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceEndpointPolicyResponse Response Object
type UpdateInstanceEndpointPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceEndpointPolicyResponse", string(data)}, " ")
}
