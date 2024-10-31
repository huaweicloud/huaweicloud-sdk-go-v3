package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleAddressDto 规则地址dto
type RuleAddressDto struct {

	// 地址类型0手工输入，1关联IP地址组，2域名，3地理位置，4域名组，5多对象，6域名组-网络型，7域名组-应用型。
	Type int32 `json:"type"`

	// 地址类型0 ipv4，1 ipv6，当type为0手动输入类型时不能为空
	AddressType *int32 `json:"address_type,omitempty"`

	// IP地址信息，当type为0手动输入类型时不能为空
	Address *string `json:"address,omitempty"`

	// 关联IP地址组ID，当type为1关联IP地址组类型时不能为空，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	AddressSetId *string `json:"address_set_id,omitempty"`

	// 关联IP地址组名称，当type为1关联IP地址组类型时不能为空，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。
	AddressSetName *string `json:"address_set_name,omitempty"`

	// type为2（域名）和7（应用域名组）具体内容根据type中7修改后的类型名称
	DomainAddressName *string `json:"domain_address_name,omitempty"`

	// 规则地域列表json值
	RegionListJson *string `json:"region_list_json,omitempty"`

	// 规则地域列表
	RegionList *[]IpRegionDto `json:"region_list,omitempty"`

	// 域名组id，type为4（域名组）或7（域名组-应用型）时不能为空。可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	DomainSetId *string `json:"domain_set_id,omitempty"`

	// 域名组名称，type为4（域名组）或7（域名组-应用型）时不能为空。可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。
	DomainSetName *string `json:"domain_set_name,omitempty"`

	// IP地址列表，当type为5（多对象）时不能为空。
	IpAddress *[]string `json:"ip_address,omitempty"`

	// 地址组id列表，当type为5（多对象）时不能为空。地址组id可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。查询条件中query_address_set_type需要设置为0自定义地址组。
	AddressGroup *[]string `json:"address_group,omitempty"`

	// 地址组名称列表
	AddressGroupNames *[]AddressGroupVo `json:"address_group_names,omitempty"`

	// 地址组类型，当type为1（关联IP地址组）时不能为空。0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// 预定义地址组id列表，当type为5（多对象）时不能为空。地址组id可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。查询条件中query_address_set_type需要设置为1预定义地址组。
	PredefinedGroup *[]string `json:"predefined_group,omitempty"`
}

func (o RuleAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddressDto struct{}"
	}

	return strings.Join([]string{"RuleAddressDto", string(data)}, " ")
}
