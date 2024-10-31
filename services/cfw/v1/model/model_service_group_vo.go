package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceGroupVo struct {

	// 服务组名称
	Name *string `json:"name,omitempty"`

	// 协议列表，协议类型：TCP为6，UDP为17，ICMP为1，ICMPV6为58，ANY为-1
	Protocols *[]int32 `json:"protocols,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示预定义服务组
	ServiceSetType *int32 `json:"service_set_type,omitempty"`

	// 服务组id，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId *string `json:"set_id,omitempty"`
}

func (o ServiceGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceGroupVo struct{}"
	}

	return strings.Join([]string{"ServiceGroupVo", string(data)}, " ")
}
