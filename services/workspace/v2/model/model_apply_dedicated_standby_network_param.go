package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDedicatedStandbyNetworkParam 开通专线备用线路入参
type ApplyDedicatedStandbyNetworkParam struct {

	// 租户指定的专线备用地址
	Address *string `json:"address,omitempty"`

	// 租户指定的专线备用地址端口
	Port *int32 `json:"port,omitempty"`

	// 开通服务资源使用的可用分区，默认随机使用2个可用区
	AvailabilityZone *[]string `json:"availability_zone,omitempty"`
}

func (o ApplyDedicatedStandbyNetworkParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDedicatedStandbyNetworkParam struct{}"
	}

	return strings.Join([]string{"ApplyDedicatedStandbyNetworkParam", string(data)}, " ")
}
