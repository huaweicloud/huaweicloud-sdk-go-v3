package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBlackWhiteListDto
type AddBlackWhiteListDto struct {

	// 防护对象id
	ObjectId string `json:"object_id"`

	// 黑白名单类型4：黑名单，5：白名单
	ListType int32 `json:"list_type"`

	// 地址方向0：源地址1：目的地址
	Direction int32 `json:"direction"`

	// Ip地址类型 0：ipv4,1:ipv6,2:domain
	AddressType int32 `json:"address_type"`

	// 地址类型
	Address string `json:"address"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol int32 `json:"protocol"`

	// 目的端口
	Port string `json:"port"`
}

func (o AddBlackWhiteListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteListDto struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteListDto", string(data)}, " ")
}
