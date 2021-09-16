package model

import (
	"encoding/json"

	"strings"
)

// DNS配置
type DnsAssignment struct {
	// 端口内网fqdn

	Fqdn *string `json:"fqdn,omitempty"`
	// 端口hostname

	Hostname *string `json:"hostname,omitempty"`
	// 端口IP地址

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o DnsAssignment) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DnsAssignment struct{}"
	}

	return strings.Join([]string{"DnsAssignment", string(data)}, " ")
}
