package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkPolicyChangeSetResponse Response Object
type ListCentralNetworkPolicyChangeSetResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 网络策略变化列表。
	CentralNetworkPolicyChangeSet []CentralNetworkElementChange `json:"central_network_policy_change_set"`
	HttpStatusCode                int                           `json:"-"`
}

func (o ListCentralNetworkPolicyChangeSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkPolicyChangeSetResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkPolicyChangeSetResponse", string(data)}, " ")
}
