package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAddressItemsInfoDto struct {

	// 地址组id，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 地址组成员id列表，地址组成员id可通过[查询地址组成员接口](ListAddressItems.xml)查询获得，通过返回值中的data.records.item_id（.表示各对象之间层级的区分）获得。
	AddressItemIds []string `json:"address_item_ids"`
}

func (o DeleteAddressItemsInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemsInfoDto struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemsInfoDto", string(data)}, " ")
}
