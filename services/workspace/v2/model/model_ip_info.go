package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpInfo struct {

	// ip地址。
	IpAddress string `json:"ip_address"`

	// 子网掩码。
	SubnetMask string `json:"subnet_mask"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
