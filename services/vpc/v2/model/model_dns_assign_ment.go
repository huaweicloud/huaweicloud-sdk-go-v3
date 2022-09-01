package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DnsAssignMent struct {

	// 端口hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 端口IP地址
	IpAddress *string `json:"ip_address,omitempty" xml:"ip_address"`

	// 端口内网fqdn
	Fqdn *string `json:"fqdn,omitempty" xml:"fqdn"`
}

func (o DnsAssignMent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsAssignMent struct{}"
	}

	return strings.Join([]string{"DnsAssignMent", string(data)}, " ")
}
