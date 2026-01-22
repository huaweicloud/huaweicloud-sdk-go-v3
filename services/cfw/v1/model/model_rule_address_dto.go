package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleAddressDto **参数解释**： 规则地址dto **约束限制**： 不涉及
type RuleAddressDto struct {

	// **参数解释**： 地址输入类型，用于区分不同的输入类型 **约束限制**： 当规则type=0（互联网规则）或者type= 2（NAT规则）时方向值（direction）必填 **取值范围**： 0手动输入，1关联IP地址组，2域名，3地理位置，4域名组-应用型，5多对象，6域名组-网络型，7域名-应用型。 **默认取值**： 不涉及
	Type int32 `json:"type"`

	// **参数解释**： IP地址互联网协议类型，用于区分不同互联网协议 **约束限制**： 当type为0手动输入时，此处不能为空 **取值范围**： 地址类型0 IPv4，1 IPv6。 **默认取值**： 不涉及
	AddressType *int32 `json:"address_type,omitempty"`

	// **参数解释**： IP地址信息，用于明确规则IP地址 **约束限制**： 当type为0手动输入时，此处不能为空 **取值范围**： 不涉及 **默认取值**： 不涉及
	Address *string `json:"address,omitempty"`

	// **参数解释**： 关联IP地址组ID，用于明确规则IP地址组id，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 当type为1关联IP地址组时，此处不能为空 **取值范围**： 不涉及 **默认取值**： 不涉及
	AddressSetId *string `json:"address_set_id,omitempty"`

	// **参数解释**： 关联IP地址组名称，用于明确规则IP地址组名称，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。 **约束限制**： 当type为1关联IP地址组时，此处不能为空 **取值范围**： 不涉及 **默认取值**： 不涉及。
	AddressSetName *string `json:"address_set_name,omitempty"`

	// **参数解释**： 域名名称或引用域名组名称，用于明确规则引用域名或域名组名称 **约束限制**： 当type为2或7时不能为空，长度为0-255 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainAddressName *string `json:"domain_address_name,omitempty"`

	// **参数解释**： 规则地域列表json值，用于明确规则引用地域名称列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RegionListJson *string `json:"region_list_json,omitempty"`

	// **参数解释**： 规则地域列表传输值 **约束限制**： 不涉及
	RegionList *[]IpRegionDto `json:"region_list,omitempty"`

	// **参数解释**： 域名组id，用于明确规则引用域名组。可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： type为4（应用型域名组）或6（网络域名组）时不能为空。 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainSetId *string `json:"domain_set_id,omitempty"`

	// **参数解释**： 域名组名称，用于明确规则引用域名组。可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。 **约束限制**： type为4（应用型域名组）或6（网络域名组）时不能为空。 **取值范围**： 不涉及 **默认取值**： 不涉及
	DomainSetName *string `json:"domain_set_name,omitempty"`

	// **参数解释**： IP地址列表，用于明确规则引用IP地址列表。 **约束限制**： 当type为5（多对象）时不能为空。 **取值范围**： 不涉及 **默认取值**： 不涉及
	IpAddress *[]string `json:"ip_address,omitempty"`

	// **参数解释**： 地址组ID列表，用于明确规则引用地址组id列表。地址组id可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。查询条件中query_address_set_type需要设置为0自定义地址组。 **约束限制**： 当type为5（多对象）时不能为空。 **取值范围**： 不涉及 **默认取值**： 不涉及
	AddressGroup *[]string `json:"address_group,omitempty"`

	// **参数解释**： 地址组名称列表 **约束限制**： 不涉及
	AddressGroupNames *[]AddressGroupVo `json:"address_group_names,omitempty"`

	// **参数解释**： 地址组类型，用于明确规则引用地址组类型。 **约束限制**： 当address的type为1（关联IP地址组）时，此处不能为空。 **取值范围**： 0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组 **默认取值**： 不涉及
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// **参数解释**： 预定义地址组id列表，用于明确规则引用预定义地址组id列表。地址组id可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。查询条件中query_address_set_type需要设置为1预定义地址组。 **约束限制**： 当type为5（多对象）时不能为空。 **取值范围**： 0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组 **默认取值**： 不涉及
	PredefinedGroup *[]string `json:"predefined_group,omitempty"`
}

func (o RuleAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddressDto struct{}"
	}

	return strings.Join([]string{"RuleAddressDto", string(data)}, " ")
}
