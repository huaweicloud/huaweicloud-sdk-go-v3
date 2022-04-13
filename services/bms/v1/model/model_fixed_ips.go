package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// fixed_ips字段数据结构说明
type FixedIps struct {
	// 网卡私网IP对应子网的子网ID（subnet_id）。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 网卡私网IP信息

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o FixedIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixedIps struct{}"
	}

	return strings.Join([]string{"FixedIps", string(data)}, " ")
}
