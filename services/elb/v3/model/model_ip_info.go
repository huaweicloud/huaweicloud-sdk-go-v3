package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpInfo ip地址组中的包含的ip 信息对象
type IpInfo struct {

	// 参数解释：IP地址组中的IP地址。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Ip string `json:"ip"`

	// 参数解释：IP地址组中ip的备注信息。
	Description string `json:"description"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
