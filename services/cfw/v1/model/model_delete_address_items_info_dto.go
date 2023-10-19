package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAddressItemsInfoDto struct {

	// 地址组id
	SetId string `json:"set_id"`

	// 地址组成员id列表
	AddressItemIds []string `json:"address_item_ids"`
}

func (o DeleteAddressItemsInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemsInfoDto struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemsInfoDto", string(data)}, " ")
}
