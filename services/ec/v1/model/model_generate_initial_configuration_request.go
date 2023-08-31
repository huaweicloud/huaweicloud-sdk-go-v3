package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateInitialConfigurationRequest Request Object
type GenerateInitialConfigurationRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *InitialConfigurationReq `json:"body,omitempty"`
}

func (o GenerateInitialConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateInitialConfigurationRequest struct{}"
	}

	return strings.Join([]string{"GenerateInitialConfigurationRequest", string(data)}, " ")
}
