package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NicIpv6Bandwidth 绑定的共享带宽信息，详情请参见 ipv6_bandwidth数据结构说明。
type NicIpv6Bandwidth struct {

	// ipv6绑定的共享带宽ID。
	BandWidthId string `json:"band_width_id"`
}

func (o NicIpv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NicIpv6Bandwidth struct{}"
	}

	return strings.Join([]string{"NicIpv6Bandwidth", string(data)}, " ")
}
