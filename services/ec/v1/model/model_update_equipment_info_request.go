package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentInfoRequest Request Object
type UpdateEquipmentInfoRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *UpdateEquipmentInfoBody `json:"body,omitempty"`
}

func (o UpdateEquipmentInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentInfoRequest", string(data)}, " ")
}
