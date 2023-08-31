package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentLanInfoResponse Response Object
type ShowEquipmentLanInfoResponse struct {

	// 设备LAN口配置列表
	LanInterfaces  *[]EquipmentLanItem `json:"lan_interfaces,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowEquipmentLanInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentLanInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentLanInfoResponse", string(data)}, " ")
}
