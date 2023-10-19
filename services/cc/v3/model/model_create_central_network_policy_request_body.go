package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkPolicyRequestBody 创建中心网络策略的请求体。
type CreateCentralNetworkPolicyRequestBody struct {
	CentralNetworkPolicyDocument *CentralNetworkPolicyDocument `json:"central_network_policy_document"`
}

func (o CreateCentralNetworkPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkPolicyRequestBody", string(data)}, " ")
}
