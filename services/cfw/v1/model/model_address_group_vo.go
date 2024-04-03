package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressGroupVo struct {

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 协议列表，协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocols *[]int32 `json:"protocols,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示常用WEB服务，2表示常用远程登录和PING，3表示常用数据库
	ServiceSetType *int32 `json:"service_set_type,omitempty"`
}

func (o AddressGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroupVo struct{}"
	}

	return strings.Join([]string{"AddressGroupVo", string(data)}, " ")
}
