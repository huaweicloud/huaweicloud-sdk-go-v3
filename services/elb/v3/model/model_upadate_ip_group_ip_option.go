package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpadateIpGroupIpOption IP地址更新参数。
type UpadateIpGroupIpOption struct {

	// 参数解释：IP地址或者IP地址段。支持IPv4、IPv6。IP地址段格式为ip-ip，例如192.168.1.2-192.168.2.253或者2001:0DB8:02de::0e12-2001:0DB8:02de::0e13，终止IP需要大于起始IP。 若传入IP地址不存在，则新增；否则更新已有IP地址的描述信息。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
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
