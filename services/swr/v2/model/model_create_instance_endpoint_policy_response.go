package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceEndpointPolicyResponse Response Object
type CreateInstanceEndpointPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceEndpointPolicyResponse", string(data)}, " ")
}
