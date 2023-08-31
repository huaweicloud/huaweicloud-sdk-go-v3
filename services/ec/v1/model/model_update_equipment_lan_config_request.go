package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentLanConfigRequest Request Object
type UpdateEquipmentLanConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentLanItem `json:"body,omitempty"`
}

func (o UpdateEquipmentLanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentLanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentLanConfigRequest", string(data)}, " ")
}
