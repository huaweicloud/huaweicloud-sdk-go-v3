package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainInfo struct {

	// 域名地址id
	DomainAddressId *string `json:"domain_address_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 域名服务器列表
	DnsIps *[]string `json:"dns_ips,omitempty"`
}

func (o DomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainInfo struct{}"
	}

	return strings.Join([]string{"DomainInfo", string(data)}, " ")
}
