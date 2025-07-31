package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateContainerNetworkPolicyResponse Response Object
type UpdateContainerNetworkPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateContainerNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateContainerNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateContainerNetworkPolicyResponse", string(data)}, " ")
}
