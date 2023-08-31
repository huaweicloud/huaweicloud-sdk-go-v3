package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEquipmentLanConfigRequest Request Object
type CreateEquipmentLanConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentLanItem `json:"body,omitempty"`
}

func (o CreateEquipmentLanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentLanConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateEquipmentLanConfigRequest", string(data)}, " ")
}
