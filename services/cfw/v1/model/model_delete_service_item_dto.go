package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServiceItemDto struct {

	// 服务组id，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 服务组成员id列表，服务组成员id可通过[查询服务成员列表接口](ListServiceItems.xml)查询获得，通过返回值中的data.records.item_id（.表示各对象之间层级的区分）获得。
	ServiceItemIds []string `json:"service_item_ids"`
}

func (o DeleteServiceItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemDto struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemDto", string(data)}, " ")
}
