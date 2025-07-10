package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySharedVpcDedicatedParam 共享Vpc专线入参。
type ApplySharedVpcDedicatedParam struct {

	// 租户指定的共享Vpc专线地址。
	Address *string `json:"address,omitempty"`

	// 租户指定共享Vpc专线地址端口。
	Port *int32 `json:"port,omitempty"`

	// 开通服务资源使用的可用分区，默认随机使用2个可用区。
	AvailabilityZone *[]string `json:"availability_zone,omitempty"`
}

func (o ApplySharedVpcDedicatedParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySharedVpcDedicatedParam struct{}"
	}

	return strings.Join([]string{"ApplySharedVpcDedicatedParam", string(data)}, " ")
}
