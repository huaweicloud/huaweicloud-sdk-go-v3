package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemDefaultDnsConfigRequestBody 配置负载均衡器系统默认域名化相关配置参数。
type SystemDefaultDnsConfigRequestBody struct {

	// **参数解释**：公网解析记录集超时时间。解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PublicDnsRecordSetTtl *int32 `json:"public_dns_record_set_ttl,omitempty"`

	// **参数解释**：是否配置私网域名。  **约束限制**：不涉及  **取值范围**： true：开启私网域名 false：关闭私网域名  **默认取值**：false
	PrivateDomainNameEnable *bool `json:"private_domain_name_enable,omitempty"`

	// **参数解释**：是否配置公网域名。  **约束限制**：不涉及  **取值范围**： true：开启公网域名 false：关闭公网域名  **默认取值**：false
	PublicDomainNameEnable *bool `json:"public_domain_name_enable,omitempty"`

	// **参数解释**：私网解析记录集超时时间。解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PrivateDnsRecordSetTtl *int32 `json:"private_dns_record_set_ttl,omitempty"`
}

func (o SystemDefaultDnsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemDefaultDnsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"SystemDefaultDnsConfigRequestBody", string(data)}, " ")
}
