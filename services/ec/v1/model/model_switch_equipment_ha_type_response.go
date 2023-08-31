package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchEquipmentHaTypeResponse Response Object
type SwitchEquipmentHaTypeResponse struct {

	// 主设备ID
	ActiveEquipmentId *string `json:"active_equipment_id,omitempty"`

	// 备设备ID
	StandbyEquipmentId *string `json:"standby_equipment_id,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o SwitchEquipmentHaTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchEquipmentHaTypeResponse struct{}"
	}

	return strings.Join([]string{"SwitchEquipmentHaTypeResponse", string(data)}, " ")
}
