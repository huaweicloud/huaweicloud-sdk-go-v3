package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkPoliciesResponse Response Object
type ListCentralNetworkPoliciesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 中心网络策略列表。
	CentralNetworkPolicies []CentralNetworkPolicy `json:"central_network_policies"`
	HttpStatusCode         int                    `json:"-"`
}

func (o ListCentralNetworkPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkPoliciesResponse", string(data)}, " ")
}
