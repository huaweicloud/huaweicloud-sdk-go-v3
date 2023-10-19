package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyCentralNetworkPolicyRequest Request Object
type ApplyCentralNetworkPolicyRequest struct {

	// 网络策略ID。
	PolicyId string `json:"policy_id"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o ApplyCentralNetworkPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCentralNetworkPolicyRequest struct{}"
	}

	return strings.Join([]string{"ApplyCentralNetworkPolicyRequest", string(data)}, " ")
}
