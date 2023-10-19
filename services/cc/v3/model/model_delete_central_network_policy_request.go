package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkPolicyRequest Request Object
type DeleteCentralNetworkPolicyRequest struct {

	// 网络策略ID。
	PolicyId string `json:"policy_id"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o DeleteCentralNetworkPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkPolicyRequest", string(data)}, " ")
}
