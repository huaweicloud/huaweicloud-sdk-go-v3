package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceItemsUsingPostRequestBodyServiceItems service item
type AddServiceItemsUsingPostRequestBodyServiceItems struct {

	// 服务成员id
	ItemId *string `json:"item_id,omitempty"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol int32 `json:"protocol"`

	// 源端口
	SourcePort string `json:"source_port"`

	// 目的端口
	DestPort string `json:"dest_port"`

	// 服务成员名称
	Name *string `json:"name,omitempty"`

	// 服务成员描述
	Description *string `json:"description,omitempty"`
}

func (o AddServiceItemsUsingPostRequestBodyServiceItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceItemsUsingPostRequestBodyServiceItems struct{}"
	}

	return strings.Join([]string{"AddServiceItemsUsingPostRequestBodyServiceItems", string(data)}, " ")
}
