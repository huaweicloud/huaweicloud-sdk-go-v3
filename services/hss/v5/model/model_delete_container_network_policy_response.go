package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteContainerNetworkPolicyResponse Response Object
type DeleteContainerNetworkPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteContainerNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteContainerNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteContainerNetworkPolicyResponse", string(data)}, " ")
}
