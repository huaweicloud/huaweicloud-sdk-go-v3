package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkPolicyResponse Response Object
type CreateCentralNetworkPolicyResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkPolicy *CentralNetworkPolicy `json:"central_network_policy"`
	HttpStatusCode       int                   `json:"-"`
}

func (o CreateCentralNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkPolicyResponse", string(data)}, " ")
}
