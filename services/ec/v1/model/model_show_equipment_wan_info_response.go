package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentWanInfoResponse Response Object
type ShowEquipmentWanInfoResponse struct {

	// 设备WAN口配置列表
	WanInterfaces  *[]EquipmentWanItem `json:"wan_interfaces,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowEquipmentWanInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentWanInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentWanInfoResponse", string(data)}, " ")
}
