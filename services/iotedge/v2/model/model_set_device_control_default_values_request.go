package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDeviceControlDefaultValuesRequest Request Object
type SetDeviceControlDefaultValuesRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	Body *DeviceControlDefaultValuesReqDto `json:"body,omitempty"`
}

func (o SetDeviceControlDefaultValuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDeviceControlDefaultValuesRequest struct{}"
	}

	return strings.Join([]string{"SetDeviceControlDefaultValuesRequest", string(data)}, " ")
}
