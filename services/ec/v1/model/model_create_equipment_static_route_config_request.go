package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEquipmentStaticRouteConfigRequest Request Object
type CreateEquipmentStaticRouteConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	Body *StaticRouteItem `json:"body,omitempty"`
}

func (o CreateEquipmentStaticRouteConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentStaticRouteConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateEquipmentStaticRouteConfigRequest", string(data)}, " ")
}
