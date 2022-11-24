package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 地址组成员信息
type AddAddressItemsInfoDto struct {

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 地址组成员信息
	AddressItems *[]AddAddressItemsInfoDtoAddressItems `json:"address_items,omitempty"`
}

func (o AddAddressItemsInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemsInfoDto struct{}"
	}

	return strings.Join([]string{"AddAddressItemsInfoDto", string(data)}, " ")
}
