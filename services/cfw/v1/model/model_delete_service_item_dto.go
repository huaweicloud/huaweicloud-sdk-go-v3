package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServiceItemDto struct {

	// 服务组id
	SetId string `json:"set_id"`

	// 服务组成员id列表
	ServiceItemIds []string `json:"service_item_ids"`
}

func (o DeleteServiceItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemDto struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemDto", string(data)}, " ")
}
