package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentEsnRequest Request Object
type UpdateEquipmentEsnRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentEsn `json:"body,omitempty"`
}

func (o UpdateEquipmentEsnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentEsnRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentEsnRequest", string(data)}, " ")
}
