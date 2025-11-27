package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkConfig struct {

	// 容器网段
	PodCIDR *string `json:"podCIDR,omitempty"`

	// 服务网段
	ServiceCIDR *string `json:"serviceCIDR,omitempty"`
}

func (o NetworkConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkConfig struct{}"
	}

	return strings.Join([]string{"NetworkConfig", string(data)}, " ")
}
