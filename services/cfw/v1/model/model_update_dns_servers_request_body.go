package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBody struct {

	// DNS服务器列表
	DnsServer []UpdateDnsServersRequestBodyDnsServer `json:"dns_server"`

	// 健康检查域名，可通过[查询dns服务器列表接口](ListDnsServers.xml)查询获得，通过返回值中的data.health_check_domain_name（.表示各对象之间层级的区分）获得。
	HealthCheckDomainName *string `json:"health_check_domain_name,omitempty"`
}

func (o UpdateDnsServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBody", string(data)}, " ")
}
