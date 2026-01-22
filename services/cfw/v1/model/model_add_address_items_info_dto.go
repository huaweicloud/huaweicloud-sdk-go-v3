package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressItemsInfoDto 地址组成员信息
type AddAddressItemsInfoDto struct {

	// 地址组id，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 地址组成员列表
	AddressItems *[]AddAddressItemsInfoDtoAddressItems `json:"address_items,omitempty"`
}

func (o AddAddressItemsInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemsInfoDto struct{}"
	}

	return strings.Join([]string{"AddAddressItemsInfoDto", string(data)}, " ")
}
