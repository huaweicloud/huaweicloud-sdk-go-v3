package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BlackWhiteListResponseDataRecords items
type BlackWhiteListResponseDataRecords struct {

	// 黑白名单列表id
	ListId *string `json:"list_id,omitempty"`

	// 黑白地址方向0：源地址1：目的地址
	Direction *int32 `json:"direction,omitempty"`

	// IP地址类型0：ipv4,1:ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 协议类型:TCP为6,UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *int32 `json:"protocol,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o BlackWhiteListResponseDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListResponseDataRecords struct{}"
	}

	return strings.Join([]string{"BlackWhiteListResponseDataRecords", string(data)}, " ")
}
