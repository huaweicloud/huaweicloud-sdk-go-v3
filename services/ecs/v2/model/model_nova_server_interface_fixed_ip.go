package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaServerInterfaceFixedIp
type NovaServerInterfaceFixedIp struct {

	// 网卡私网IP信息。
	IpAddress string `json:"ip_address"`

	// 网卡私网IP对应子网信息。
	SubnetId string `json:"subnet_id"`
}

func (o NovaServerInterfaceFixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"NovaServerInterfaceFixedIp", string(data)}, " ")
}
