package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkPolicyChangeSetRequest Request Object
type ListCentralNetworkPolicyChangeSetRequest struct {

	// 网络策略ID。
	PolicyId string `json:"policy_id"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o ListCentralNetworkPolicyChangeSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkPolicyChangeSetRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkPolicyChangeSetRequest", string(data)}, " ")
}
