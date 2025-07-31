package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNetworkPolicyRequestInfo struct {

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	PolicyContent *NetworkPolicyBody `json:"policy_content"`
}

func (o CreateNetworkPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateNetworkPolicyRequestInfo", string(data)}, " ")
}
