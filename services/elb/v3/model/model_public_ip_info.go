package model

import (
	"encoding/json"

	"strings"
)

// 弹性公网EIP信息
type PublicIpInfo struct {
	// 弹性公网ip配置id

	PublicipId string `json:"publicip_id"`
	// IP地址

	PublicipAddress string `json:"publicip_address"`
	// IP版本信息。 取值范围：4和6 4：IPv4 6：IPv6

	IpVersion int32 `json:"ip_version"`
}

func (o PublicIpInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}
