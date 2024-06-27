package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBlackWhiteListDto UpdateBlackWhiteListDto
type UpdateBlackWhiteListDto struct {

	// 地址方向0：源地址1：目的地址
	Direction *int32 `json:"direction,omitempty"`

	// 地址类型0：ipv4,1:ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address string `json:"address"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1
	Protocol *int32 `json:"protocol,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateBlackWhiteListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListDto struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListDto", string(data)}, " ")
}
