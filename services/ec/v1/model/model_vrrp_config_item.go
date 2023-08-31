package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VrrpConfigItem struct {

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// 虚路由ID
	VirtualRouterId *int32 `json:"virtual_router_id,omitempty"`

	// 虚IP
	VirtualIp *string `json:"virtual_ip,omitempty"`

	// 主设备ID
	ActiveEquipmentId *string `json:"active_equipment_id,omitempty"`

	// 主设备接口名字
	ActiveInterfaceName *string `json:"active_interface_name,omitempty"`

	// 备设备ID
	StandbyEquipmentId *string `json:"standby_equipment_id,omitempty"`

	// 备设备接口名字
	StandbyInterfaceName *string `json:"standby_interface_name,omitempty"`
}

func (o VrrpConfigItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VrrpConfigItem struct{}"
	}

	return strings.Join([]string{"VrrpConfigItem", string(data)}, " ")
}
