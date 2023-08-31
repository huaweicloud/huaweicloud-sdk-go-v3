package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentStaticRouteConfigRequest Request Object
type UpdateEquipmentStaticRouteConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *StaticRouteItem `json:"body,omitempty"`
}

func (o UpdateEquipmentStaticRouteConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentStaticRouteConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentStaticRouteConfigRequest", string(data)}, " ")
}
