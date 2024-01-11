package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ipv6BandwidthForNic struct {

	// IPv6带宽的ID。
	Id *string `json:"id,omitempty"`
}

func (o Ipv6BandwidthForNic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipv6BandwidthForNic struct{}"
	}

	return strings.Join([]string{"Ipv6BandwidthForNic", string(data)}, " ")
}
