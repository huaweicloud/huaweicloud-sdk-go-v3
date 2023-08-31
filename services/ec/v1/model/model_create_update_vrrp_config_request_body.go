package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUpdateVrrpConfigRequestBody struct {

	// 虚IP
	VirtualIp string `json:"virtual_ip"`

	// 主设备ID
	ActiveEquipmentId string `json:"active_equipment_id"`

	// 主设备接口名字
	ActiveInterfaceName string `json:"active_interface_name"`

	// 备设备ID
	StandbyEquipmentId string `json:"standby_equipment_id"`

	// 备设备接口名字
	StandbyInterfaceName string `json:"standby_interface_name"`
}

func (o CreateUpdateVrrpConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpdateVrrpConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUpdateVrrpConfigRequestBody", string(data)}, " ")
}
