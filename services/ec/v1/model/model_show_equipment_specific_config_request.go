package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentSpecificConfigRequest Request Object
type ShowEquipmentSpecificConfigRequest struct {

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`
}

func (o ShowEquipmentSpecificConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentSpecificConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowEquipmentSpecificConfigRequest", string(data)}, " ")
}
