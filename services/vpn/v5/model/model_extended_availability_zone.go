package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendedAvailabilityZone struct {

	// 可用区名称
	Name *string `json:"name,omitempty"`

	// 公共边界组
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 该可用区下可选的VPN网关规格组合
	AvailableSpecs *[]AvailableSpec `json:"available_specs,omitempty"`
}

func (o ExtendedAvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendedAvailabilityZone struct{}"
	}

	return strings.Join([]string{"ExtendedAvailabilityZone", string(data)}, " ")
}
