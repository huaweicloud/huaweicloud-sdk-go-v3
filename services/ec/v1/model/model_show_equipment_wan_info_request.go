package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentWanInfoRequest Request Object
type ShowEquipmentWanInfoRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`
}

func (o ShowEquipmentWanInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentWanInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEquipmentWanInfoRequest", string(data)}, " ")
}
