package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentOspfRequest Request Object
type UpdateEquipmentOspfRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentOspfItem `json:"body,omitempty"`
}

func (o UpdateEquipmentOspfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentOspfRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentOspfRequest", string(data)}, " ")
}
