package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleAddressDto 规则地址dto
type RuleAddressDto struct {

	// 源类型0手工输入,1关联IP地址组,2域名
	Type int32 `json:"type"`

	// 源类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 源IP，手动类型不能为空，自动及domain类型为空
	Address *string `json:"address,omitempty"`

	// 关联IP地址组ID，自动类型不能为空，手动类型合domain类型为空
	AddressSetId *string `json:"address_set_id,omitempty"`

	// 地址组名称
	AddressSetName *string `json:"address_set_name,omitempty"`

	// 域名地址名称，域名类型时不能为空，手动类型及自动类型时为空
	DomainAddressName *string `json:"domain_address_name,omitempty"`
}

func (o RuleAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAddressDto struct{}"
	}

	return strings.Join([]string{"RuleAddressDto", string(data)}, " ")
}
