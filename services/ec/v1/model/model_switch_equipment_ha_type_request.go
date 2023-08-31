package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchEquipmentHaTypeRequest Request Object
type SwitchEquipmentHaTypeRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	Body *SwitchHaTypeBody `json:"body,omitempty"`
}

func (o SwitchEquipmentHaTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchEquipmentHaTypeRequest struct{}"
	}

	return strings.Join([]string{"SwitchEquipmentHaTypeRequest", string(data)}, " ")
}
