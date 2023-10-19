package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkPolicyResponse Response Object
type DeleteCentralNetworkPolicyResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCentralNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkPolicyResponse", string(data)}, " ")
}
