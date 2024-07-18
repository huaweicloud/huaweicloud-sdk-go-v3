package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpadateIpGroupIpOption IP地址更新参数。
type UpadateIpGroupIpOption struct {

	// 参数解释：IP地址。支持IPv4、IPv6。若传入IP地址不存在，则新增；否则更新已有IP地址的描述信息。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Ip string `json:"ip"`

	// 参数解释：备注信息。
	Description *string `json:"description,omitempty"`
}

func (o UpadateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpadateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"UpadateIpGroupIpOption", string(data)}, " ")
}
