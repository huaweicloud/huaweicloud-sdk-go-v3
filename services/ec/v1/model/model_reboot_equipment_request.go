package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootEquipmentRequest Request Object
type RebootEquipmentRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`
}

func (o RebootEquipmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootEquipmentRequest struct{}"
	}

	return strings.Join([]string{"RebootEquipmentRequest", string(data)}, " ")
}
