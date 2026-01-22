package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBody struct {

	// **参数解释**： DNS服务器列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DnsServer []UpdateDnsServersRequestBodyDnsServer `json:"dns_server"`

	// **参数解释**： 健康检查域名，可通过[查询DNS服务器列表接口](ListDnsServers.xml)查询获得，通过返回值中的data.health_check_domain_name（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HealthCheckDomainName *string `json:"health_check_domain_name,omitempty"`
}

func (o UpdateDnsServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBody", string(data)}, " ")
}
