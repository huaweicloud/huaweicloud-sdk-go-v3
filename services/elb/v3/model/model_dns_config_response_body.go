package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsConfigResponseBody 配置负载均衡器响应体相关配置参数。
type DnsConfigResponseBody struct {

	// **参数解释**：负载均衡器的IPv4虚拟IP地址。
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**：双栈类型负载均衡器的IPv6地址。  **约束限制**：[不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// **参数解释**：负载均衡器绑定的EIP。  注：该字段与publicips一致。
	Eips *[]EipInfo `json:"eips,omitempty"`

	// **参数解释**：负载均衡器绑定的GEIP。
	GlobalEips *[]GlobalEipInfo `json:"global_eips,omitempty"`

	// **参数解释**：是否配置公网域名。 **取值范围**：   true：开启公网域名   false：关闭公网域名
	PublicDomainNameEnable *bool `json:"public_domain_name_enable,omitempty"`

	// **参数解释**：公网域名所使用的zone名称。 **约束限制**：   公网域名只能使用公网类型的zone。   当配置公网域名开关打开时，该字段不能置空。   所填的公网zone必须在云解析服务已注册过。
	PublicDnsZoneName *string `json:"public_dns_zone_name,omitempty"`

	// **参数解释**：   公网域名所使用的zone对应的id。   根据传入的公网zone 名称查询得出。
	PublicDnsZoneId *string `json:"public_dns_zone_id,omitempty"`

	// **参数解释**：   负载均衡实例的公网域名。 **约束限制**：   根据负载均衡实例id，局点id和zone信息以如下格式生成：   {lb_id}.elb.{region_id}.{zone_name}
	PublicDomainName *string `json:"public_domain_name,omitempty"`

	// 参数解释:   公网解析记录集超时时间。   解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。   如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。 **取值范围**：   1-2147483647 **默认取值**：   300
	PublicDnsRecordSetTtl *int32 `json:"public_dns_record_set_ttl,omitempty"`

	// **参数解释**：   是否配置私网域名。 **取值范围**：   true：开启私网域名   false：关闭私网域名
	PrivateDomainNameEnable *bool `json:"private_domain_name_enable,omitempty"`

	// **参数解释**：   私网域名所使用的zone的名称。 **约束限制**：   私网域名既能使用公网zone，也能使用私网zone，zone的类型在private_dns_zone_type字段中指定。   当配置私网域名开关打开时，该字段不能置空。   所填的私网zone必须在云解析服务已注册过。
	PrivateDnsZoneName *string `json:"private_dns_zone_name,omitempty"`

	// **参数解释**：   私网域名所使用的zone对应的id。 **约束限制**：   根据传入的私网zone 名称查询得出。
	PrivateDnsZoneId *string `json:"private_dns_zone_id,omitempty"`

	// **参数解释**：负载均衡实例的私网域名。 **约束限制**：   根据负载均衡实例id，局点id和zone信息以如下格式生成：   {lb_id}-internal.elb.{region_id}.{zone_name}
	PrivateDomainName *string `json:"private_domain_name,omitempty"`

	// **参数解释**：私网域名所使用的zone的类型。 **约束限制**：不涉及 **取值范围**：private public **默认取值**：private
	PrivateDnsZoneType *string `json:"private_dns_zone_type,omitempty"`

	// **参数解释**：   私网解析记录集超时时间。   解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。   如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。 **取值范围**：   1-2147483647 **默认取值**：   300
	PrivateDnsRecordSetTtl *int32 `json:"private_dns_record_set_ttl,omitempty"`
}

func (o DnsConfigResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsConfigResponseBody struct{}"
	}

	return strings.Join([]string{"DnsConfigResponseBody", string(data)}, " ")
}
