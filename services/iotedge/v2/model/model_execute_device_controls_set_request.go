package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDeviceControlsSetRequest Request Object
type ExecuteDeviceControlsSetRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 设备ID
	DeviceId string `json:"device_id"`

	Body *DeviceControlSetReqDto `json:"body,omitempty"`
}

func (o ExecuteDeviceControlsSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeviceControlsSetRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDeviceControlsSetRequest", string(data)}, " ")
}
