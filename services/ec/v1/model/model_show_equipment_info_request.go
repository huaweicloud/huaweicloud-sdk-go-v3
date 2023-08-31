package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentInfoRequest Request Object
type ShowEquipmentInfoRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`
}

func (o ShowEquipmentInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEquipmentInfoRequest", string(data)}, " ")
}
