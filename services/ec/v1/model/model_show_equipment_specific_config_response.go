package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentSpecificConfigResponse Response Object
type ShowEquipmentSpecificConfigResponse struct {

	// 设备类型
	Type *string `json:"type,omitempty"`

	// WAN口列表
	WanInterfaces *[]string `json:"wan_interfaces,omitempty"`

	// LTE口列表
	LteInterfaces *[]string `json:"lte_interfaces,omitempty"`

	// LAN口列表
	LanInterfaces  *[]string `json:"lan_interfaces,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowEquipmentSpecificConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentSpecificConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentSpecificConfigResponse", string(data)}, " ")
}
