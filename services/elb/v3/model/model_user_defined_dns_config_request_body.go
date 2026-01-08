package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserDefinedDnsConfigRequestBody 配置负载均衡器用户自定义域名化相关配置参数。
type UserDefinedDnsConfigRequestBody struct {

	// **参数解释**：是否配置公网域名。  **约束限制**：不涉及  **取值范围**： true：开启公网域名 false：关闭公网域名  **默认取值**：false
	PublicDomainNameEnable *bool `json:"public_domain_name_enable,omitempty"`

	// **参数解释**：公网域名所使用的zone名称。  **约束限制**：公网域名只能使用公网类型的zone，当配置公网域名开关打开时，该字段不能置空，所填的公网zone必须在云解析服务已注册过。  **取值范围**：不涉及  **默认取值**：不涉及
	PublicDnsZoneName *string `json:"public_dns_zone_name,omitempty"`

	// **参数解释**：公网解析记录集超时时间。解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PublicDnsRecordSetTtl *int32 `json:"public_dns_record_set_ttl,omitempty"`

	// **参数解释**：是否配置私网域名。  **约束限制**：不涉及  **取值范围**： true：开启私网域名 false：关闭私网域名  **默认取值**：false
	PrivateDomainNameEnable *bool `json:"private_domain_name_enable,omitempty"`

	// **参数解释**：私网域名所使用的zone的名称。  **约束限制**：   只有当private_domain_name_enable打开时，该字段才有效。   当private_domain_name_enable打开时，该字段不能置空。   所填的私网zone必须在云解析服务已注册过。   私网域名既能使用公网zone，也能使用私网zone，zone的类型在private_dns_zone_type字段中指定。  **取值范围**：不涉及  **默认取值**：不涉及
	PrivateDnsZoneName *string `json:"private_dns_zone_name,omitempty"`

	// **参数解释**：私网域名所使用的zone的类型。  **约束限制**：不涉及  **取值范围**： private: 使用私网zone public: 使用公网zone  **默认取值**：private
	PrivateDnsZoneType *string `json:"private_dns_zone_type,omitempty"`

	// **参数解释**：私网解析记录集超时时间。解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PrivateDnsRecordSetTtl *int32 `json:"private_dns_record_set_ttl,omitempty"`
}

func (o UserDefinedDnsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDefinedDnsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UserDefinedDnsConfigRequestBody", string(data)}, " ")
}
