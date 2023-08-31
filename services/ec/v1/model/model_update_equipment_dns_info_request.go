package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentDnsInfoRequest Request Object
type UpdateEquipmentDnsInfoRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *EquipmentDnsItem `json:"body,omitempty"`
}

func (o UpdateEquipmentDnsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentDnsInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentDnsInfoRequest", string(data)}, " ")
}
