package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpInfo **参数解释**：IP地址组中IP列表的IP地址信息。
type IpInfo struct {

	// **参数解释**：IP地址组中的IP地址。  **取值范围**：IP地址段格式为ip-ip，例如192.168.1.2-192.168.2.253或者2001:0DB8:02de::0e12-2001:0DB8:02de::0e13，终止IP需要大于起始IP。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt)
	Ip string `json:"ip"`

	// **参数解释**：IP地址组中ip的备注信息。  **取值范围**：长度为0-255个字符。
	Description string `json:"description"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
