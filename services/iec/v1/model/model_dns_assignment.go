package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DNS配置
type DnsAssignment struct {

	// 端口内网fqdn
	Fqdn *string `json:"fqdn,omitempty" xml:"fqdn"`

	// 端口hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 端口IP地址
	IpAddress *string `json:"ip_address,omitempty" xml:"ip_address"`
}

func (o DnsAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsAssignment struct{}"
	}

	return strings.Join([]string{"DnsAssignment", string(data)}, " ")
}
