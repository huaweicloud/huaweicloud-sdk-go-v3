package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateObjectConfigDesc struct {

	// 成员描述
	Description *string `json:"description,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 成员id，地址组成员id，可通过[查询地址组成员接口](ListAddressItems.xml)查询获得，通过返回值中的data.records.item_id（.表示各对象之间层级的区分）获得。服务组成员id，可通过[查询服务成员列表接口](ListServiceItems.xml)查询获得，通过返回值中的data.records.item_id（.表示各对象之间层级的区分）获得。域名id可通过[获取域名组下域名列表接口](ListDomains.xml)查询获得，通过返回值中的data.records.domain_address_id（.表示各对象之间层级的区分）获得。
	ItemId *string `json:"item_id,omitempty"`

	// 组id，地址组id，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。服务组id，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。域名组id，可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId *string `json:"set_id,omitempty"`

	// 组类型，包含：地址组ADDR_SET，服务组SERV_SET，域名组DOMAIN_SET，URL组URL_SET
	Type *string `json:"type,omitempty"`
}

func (o UpdateObjectConfigDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObjectConfigDesc struct{}"
	}

	return strings.Join([]string{"UpdateObjectConfigDesc", string(data)}, " ")
}
