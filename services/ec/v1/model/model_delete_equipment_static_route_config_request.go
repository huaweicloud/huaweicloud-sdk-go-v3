package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEquipmentStaticRouteConfigRequest Request Object
type DeleteEquipmentStaticRouteConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	// 目标网络
	Prefix string `json:"prefix"`

	// 下一跳地址
	NextHop string `json:"next_hop"`

	// 接口名字
	InterfaceName string `json:"interface_name"`
}

func (o DeleteEquipmentStaticRouteConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEquipmentStaticRouteConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteEquipmentStaticRouteConfigRequest", string(data)}, " ")
}
