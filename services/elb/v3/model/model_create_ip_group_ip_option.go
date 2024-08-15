package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupIpOption 参数解释：IP地址组中的包含的IP信息。
type CreateIpGroupIpOption struct {

	// IP地址或IP地址段。支持IPv4、IPv6。IP地址段格式为ip-ip，例如192.168.1.2-192.168.2.253或者2001:0DB8:02de::0e12-2001:0DB8:02de::0e13，终止IP需要大于起始IP.  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Ip string `json:"ip"`

	// 参数解释：备注信息。
	Description *string `json:"description,omitempty"`
}

func (o CreateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupIpOption", string(data)}, " ")
}
