package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkPolicyRequest Request Object
type CreateCentralNetworkPolicyRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *CreateCentralNetworkPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateCentralNetworkPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkPolicyRequest", string(data)}, " ")
}
