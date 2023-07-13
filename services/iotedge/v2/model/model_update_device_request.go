package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceRequest Request Object
type UpdateDeviceRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 设备ID
	DeviceId string `json:"device_id"`

	Body *UpdateDesireds `json:"body,omitempty"`
}

func (o UpdateDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequest", string(data)}, " ")
}
