package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentWanConfigRequest Request Object
type UpdateEquipmentWanConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentWanItemList `json:"body,omitempty"`
}

func (o UpdateEquipmentWanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentWanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentWanConfigRequest", string(data)}, " ")
}
