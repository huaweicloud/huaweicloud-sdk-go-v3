package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleAddressDtoForResponse 规则地址dto
type RuleAddressDtoForResponse struct {

	// 地址类型0手工输入，1关联IP地址组，2域名，3地理位置，4域名组，5多对象，6域名组-DNS解析，7域名组-应用型。
	Type int32 `json:"type"`

	// 地址类型0 ipv4，1 ipv6，当type为0手动输入类型时不能为空
	AddressType *int32 `json:"address_type,omitempty"`

	// IP地址信息
	Address *string `json:"address,omitempty"`

	// 关联IP地址组ID
	AddressSetId *string `json:"address_set_id,omitempty"`

	// 地址组名称
	AddressSetName *string `json:"address_set_name,omitempty"`

	// 域名地址名称
	DomainAddressName *string `json:"domain_address_name,omitempty"`

	// 规则地域列表json值
	RegionListJson *string `json:"region_list_json,omitempty"`

	// 规则地域列表
	RegionList *[]IpRegionDto `json:"region_list,omitempty"`

	// 域名组id
	DomainSetId *string `json:"domain_set_id,omitempty"`

	// 域名组名称
	DomainSetName *string `json:"domain_set_name,omitempty"`

	// IP地址列表
	IpAddress *[]string `json:"ip_address,omitempty"`

	// 地址组id列表
	AddressGroup *[]string `json:"address_group,omitempty"`

	// 地址组名称列表
	AddressGroupNames *[]AddressGroupVo `json:"address_group_names,omitempty"`

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`
}

func (o RuleAddressDtoForResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddressDtoForResponse struct{}"
	}

	return strings.Join([]string{"RuleAddressDtoForResponse", string(data)}, " ")
}
