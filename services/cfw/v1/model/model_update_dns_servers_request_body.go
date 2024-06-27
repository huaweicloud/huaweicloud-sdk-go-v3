package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBody struct {

	// DNS服务器
	DnsServer []UpdateDnsServersRequestBodyDnsServer `json:"dns_server"`

	// 健康检查域名
	HealthCheckDomainName *string `json:"health_check_domain_name,omitempty"`
}

func (o UpdateDnsServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBody", string(data)}, " ")
}
