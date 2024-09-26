package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyCentralNetworkPolicyResponse Response Object
type ApplyCentralNetworkPolicyResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkPolicy *CentralNetworkPolicy `json:"central_network_policy"`

	// 网络策略变化列表。
	CentralNetworkPolicyChangeSet []CentralNetworkElementChange `json:"central_network_policy_change_set"`
	HttpStatusCode                int                           `json:"-"`
}

func (o ApplyCentralNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCentralNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"ApplyCentralNetworkPolicyResponse", string(data)}, " ")
}
