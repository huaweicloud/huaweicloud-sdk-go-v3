package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsConfigDto struct {

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// 域名解析对应IP
	Ip *string `json:"ip,omitempty"`
}

func (o DnsConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsConfigDto struct{}"
	}

	return strings.Join([]string{"DnsConfigDto", string(data)}, " ")
}
