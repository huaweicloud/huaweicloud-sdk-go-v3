package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceItemListResponseDtoDataRecords service item
type ServiceItemListResponseDtoDataRecords struct {

	// 服务成员id
	ItemId *string `json:"item_id,omitempty"`

	// 协议类型:TCP为6,UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *int32 `json:"protocol,omitempty"`

	// 源端口
	SourcePort *string `json:"source_port,omitempty"`

	// 目的端口
	DestPort *string `json:"dest_port,omitempty"`

	// 服务成员描述
	Description *string `json:"description,omitempty"`
}

func (o ServiceItemListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItemListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"ServiceItemListResponseDtoDataRecords", string(data)}, " ")
}
