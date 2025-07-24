package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkInterface 弹性网卡
type NetworkInterface struct {

	// subnet id
	SubnetId string `json:"subnet_id"`

	// 弹性网卡私有IPv4地址
	Ipv4Address *string `json:"ipv4_address,omitempty"`
}

func (o NetworkInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkInterface struct{}"
	}

	return strings.Join([]string{"NetworkInterface", string(data)}, " ")
}
