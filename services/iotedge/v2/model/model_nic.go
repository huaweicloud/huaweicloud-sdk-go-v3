package model

import (
	"encoding/json"

	"strings"
)

// 网络网卡信息
type Nic struct {
	// 网卡信息，如eth0,eth1

	Eth *string `json:"eth,omitempty"`
	// 网卡ip

	Ip *string `json:"ip,omitempty"`
}

func (o Nic) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
