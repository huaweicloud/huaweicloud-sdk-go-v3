package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 地址组成员列表
type AddressItemListResponseDtoDataRecords struct {

	// 地址组成员id
	ItemId *string `json:"item_id,omitempty"`

	// 地址组成员name
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 地址组类型，0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 地址组
	Address *string `json:"address,omitempty"`
}

func (o AddressItemListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"AddressItemListResponseDtoDataRecords", string(data)}, " ")
}
