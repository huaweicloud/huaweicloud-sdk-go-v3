package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressItemsInfoDtoAddressItems 成员信息
type AddAddressItemsInfoDtoAddressItems struct {

	// 地址类型0 ipv4，1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address string `json:"address"`

	// 地址组成员描述
	Description *string `json:"description,omitempty"`
}

func (o AddAddressItemsInfoDtoAddressItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemsInfoDtoAddressItems struct{}"
	}

	return strings.Join([]string{"AddAddressItemsInfoDtoAddressItems", string(data)}, " ")
}
