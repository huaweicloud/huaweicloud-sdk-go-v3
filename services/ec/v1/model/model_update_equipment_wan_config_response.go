package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentWanConfigResponse Response Object
type UpdateEquipmentWanConfigResponse struct {

	// 设备WAN口配置列表
	WanInterfaces  *[]EquipmentWanItem `json:"wan_interfaces,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateEquipmentWanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentWanConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentWanConfigResponse", string(data)}, " ")
}
