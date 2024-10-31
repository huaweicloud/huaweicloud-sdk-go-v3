package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsServersResponseDto struct {

	// 域名服务器id
	Id *int32 `json:"id,omitempty"`

	// 域名服务器是否应用，0否 1是
	IsApplied *int32 `json:"is_applied,omitempty"`

	// 域名服务器是否是用户自定义的dns服务器，0否 1是
	IsCustomized *int32 `json:"is_customized,omitempty"`

	// DNS服务器IP
	ServerIp *string `json:"server_ip,omitempty"`

	// 健康检查域名
	HealthCheckDomainName *string `json:"health_check_domain_name,omitempty"`
}

func (o DnsServersResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsServersResponseDto struct{}"
	}

	return strings.Join([]string{"DnsServersResponseDto", string(data)}, " ")
}
