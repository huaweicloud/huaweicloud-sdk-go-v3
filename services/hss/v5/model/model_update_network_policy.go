package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNetworkPolicy struct {

	// 策略Id
	PolicyId string `json:"policy_id"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	PolicyContent *NetworkPolicyBody `json:"policy_content"`
}

func (o UpdateNetworkPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkPolicy struct{}"
	}

	return strings.Join([]string{"UpdateNetworkPolicy", string(data)}, " ")
}
