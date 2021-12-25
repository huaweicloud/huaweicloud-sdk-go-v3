package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IP地址更新参数。
type UpadateIpGroupIpOption struct {
	// IP地址。支持IPv4、IPv6。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)

	Ip string `json:"ip"`
	// 备注信息。

	Description *string `json:"description,omitempty"`
}

func (o UpadateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpadateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"UpadateIpGroupIpOption", string(data)}, " ")
}
