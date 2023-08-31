package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEquipmentLanConfigRequest Request Object
type DeleteEquipmentLanConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 智能企业网关设备ID
	EquipmentId string `json:"equipment_id"`

	// 接口名字
	InterfaceName string `json:"interface_name"`

	// VlanID
	VlanId *string `json:"vlan_id,omitempty"`
}

func (o DeleteEquipmentLanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEquipmentLanConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteEquipmentLanConfigRequest", string(data)}, " ")
}
