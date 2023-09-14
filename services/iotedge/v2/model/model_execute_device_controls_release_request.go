package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDeviceControlsReleaseRequest Request Object
type ExecuteDeviceControlsReleaseRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 设备ID
	DeviceId string `json:"device_id"`

	Body *DeviceControlReleaseReqDto `json:"body,omitempty"`
}

func (o ExecuteDeviceControlsReleaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeviceControlsReleaseRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDeviceControlsReleaseRequest", string(data)}, " ")
}
