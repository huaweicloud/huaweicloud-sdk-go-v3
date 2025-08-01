package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateContainerNetworkPolicyResponse Response Object
type CreateContainerNetworkPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateContainerNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateContainerNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateContainerNetworkPolicyResponse", string(data)}, " ")
}
