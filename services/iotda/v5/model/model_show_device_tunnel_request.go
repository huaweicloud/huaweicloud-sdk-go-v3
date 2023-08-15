package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceTunnelRequest Request Object
type ShowDeviceTunnelRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// 隧道ID
	TunnelId string `json:"tunnel_id"`
}

func (o ShowDeviceTunnelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTunnelRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceTunnelRequest", string(data)}, " ")
}
