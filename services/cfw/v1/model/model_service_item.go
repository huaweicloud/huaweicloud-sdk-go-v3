package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceItem struct {

	// 协议类型：TCP为6，UDP为17，ICMP为1，ICMPV6为58，ANY为-1,RuleServiceDto.type为0时不能为空。
	Protocol *int32 `json:"protocol,omitempty"`

	// 源端口
	SourcePort *string `json:"source_port,omitempty"`

	// 目的端口
	DestPort *string `json:"dest_port,omitempty"`

	// 服务成员描述
	Description *string `json:"description,omitempty"`

	// 服务成员名称
	Name *string `json:"name,omitempty"`
}

func (o ServiceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItem struct{}"
	}

	return strings.Join([]string{"ServiceItem", string(data)}, " ")
}
