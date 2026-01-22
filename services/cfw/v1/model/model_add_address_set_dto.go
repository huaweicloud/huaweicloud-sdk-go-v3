package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressSetDto 添加地址组请求体
type AddAddressSetDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得，此处只能使用type为0的防护对象id
	ObjectId string `json:"object_id"`

	// 地址组名称
	Name string `json:"name"`

	// 地址组描述
	Description *string `json:"description,omitempty"`

	// **参数解释**： 地址类型 **约束限制**： 不涉及 **取值范围**： - 0：IPv4 - 1：IPv6 **默认取值**： 0：IPv4
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o AddAddressSetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressSetDto struct{}"
	}

	return strings.Join([]string{"AddAddressSetDto", string(data)}, " ")
}
