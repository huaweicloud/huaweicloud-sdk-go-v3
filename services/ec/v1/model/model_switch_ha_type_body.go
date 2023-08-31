package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchHaTypeBody struct {

	// 主设备ID
	ActiveEquipmentId string `json:"active_equipment_id"`

	// 备设备ID
	StandbyEquipmentId string `json:"standby_equipment_id"`
}

func (o SwitchHaTypeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHaTypeBody struct{}"
	}

	return strings.Join([]string{"SwitchHaTypeBody", string(data)}, " ")
}
