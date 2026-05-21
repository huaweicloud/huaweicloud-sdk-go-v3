package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserDefinedDnsConfigRequestBody struct {

	// **参数解释**：是否启用公网域名解析。  **约束限制**：不涉及  **取值范围**： - true：开启公网域名解析。 - false：关闭公网域名解析。  **默认取值**：false
	PublicDomainNameEnable *bool `json:"public_domain_name_enable,omitempty"`

	// **参数解释**：公网域名解析所使用的根域名。  **约束限制**： - 公网域名解析只能选择公网类型的根域名。 - 若启用公网域名解析（public_domain_name_enable=true），则公网根域名不能为空，且必须在云解析服务已注册。  **取值范围**：不涉及  **默认取值**：不涉及
	PublicDnsZoneName *string `json:"public_dns_zone_name,omitempty"`

	// **参数解释**：公网域名解析记录在本地DNS服务器的缓存超时时间，单位：秒。域名解析信息更新后，需要等待DNS服务器上的缓存超时才会生效。如果您的域名解析信息经常变更，建议TTL值设置相对小些，反之建议设置相对大些。  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PublicDnsRecordSetTtl *int32 `json:"public_dns_record_set_ttl,omitempty"`

	// **参数解释**：是否启用私网域名解析。  **约束限制**：不涉及  **取值范围**： true：开启私网域名 false：关闭私网域名  **默认取值**：false
	PrivateDomainNameEnable *bool `json:"private_domain_name_enable,omitempty"`

	// **参数解释**：私网域名解析所使用的根域名。  **约束限制**： - 私网域名解析可以选择私网类型的根域名，也可以选择公网类型的根域名。需要在private_dns_zone_type字段中明确指定。 - 若启用私网域名解析（private_domain_name_enable=true），则私网根域名不能为空，且必须在云解析服务已注册。  **取值范围**：不涉及  **默认取值**：不涉及
	PrivateDnsZoneName *string `json:"private_dns_zone_name,omitempty"`

	// **参数解释**：私网域名解析所使用的根域名的类型。  **约束限制**：不涉及  **取值范围**： - private: 私网根域名。 - public: 公网根域名。  **默认取值**：private
	PrivateDnsZoneType *string `json:"private_dns_zone_type,omitempty"`

	// **参数解释**：私网域名解析记录在本地DNS服务器的缓存超时时间，单位：秒。域名解析信息更新后，需要等待DNS服务器上的缓存超时才会生效。如果您的域名解析信息经常变更，建议TTL值设置相对小些，反之建议设置相对大些。  **约束限制**：不涉及  **取值范围**：1-2147483647  **默认取值**：300
	PrivateDnsRecordSetTtl *int32 `json:"private_dns_record_set_ttl,omitempty"`
}

func (o UserDefinedDnsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDefinedDnsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UserDefinedDnsConfigRequestBody", string(data)}, " ")
}
