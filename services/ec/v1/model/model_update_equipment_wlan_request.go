package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentWlanRequest Request Object
type UpdateEquipmentWlanRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentWlanItem `json:"body,omitempty"`
}

func (o UpdateEquipmentWlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentWlanRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentWlanRequest", string(data)}, " ")
}
