package model

import (
	"encoding/json"

	"strings"
)

// ip地址组中的包含的ip 信息对象
type UpadateIpGroupIpOption struct {
	// ip地址组中的包含的ip。 支持ipv4、ipv6的ip

	Ip string `json:"ip"`
	// IP地址组中ip的备注信息

	Description *string `json:"description,omitempty"`
}

func (o UpadateIpGroupIpOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpadateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"UpadateIpGroupIpOption", string(data)}, " ")
}
