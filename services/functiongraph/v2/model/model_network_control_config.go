package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkControlConfig 函数网络配置。
type NetworkControlConfig struct {

	// 禁止公网访问开关。
	DisablePublicNetwork *bool `json:"disable_public_network,omitempty"`

	// 指定触发函数vpc配置。
	TriggerAccessVpcs *[]VpcConfig `json:"trigger_access_vpcs,omitempty"`
}

func (o NetworkControlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkControlConfig struct{}"
	}

	return strings.Join([]string{"NetworkControlConfig", string(data)}, " ")
}
