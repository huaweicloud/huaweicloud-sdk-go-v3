package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EquipmentWanItemList struct {

	// 设备WAN口配置列表
	WanInterfaces *[]EquipmentWanItem `json:"wan_interfaces,omitempty"`
}

func (o EquipmentWanItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentWanItemList struct{}"
	}

	return strings.Join([]string{"EquipmentWanItemList", string(data)}, " ")
}
