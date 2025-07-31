package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetNetworkPolicy struct {

	// 策略Id
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	PolicyContent *NetworkPolicyBody `json:"policy_content,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 下发策略到cce成功与否
	DeployStatus *bool `json:"deploy_status,omitempty"`
}

func (o GetNetworkPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetNetworkPolicy struct{}"
	}

	return strings.Join([]string{"GetNetworkPolicy", string(data)}, " ")
}
