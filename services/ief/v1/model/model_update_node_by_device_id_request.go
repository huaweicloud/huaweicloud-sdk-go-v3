package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNodeByDeviceIdRequest struct {
	// 设备ID

	DeviceId string `json:"device_id"`

	Body *EdgeNodeUpdateByDevice `json:"body,omitempty"`
}

func (o UpdateNodeByDeviceIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeByDeviceIdRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeByDeviceIdRequest", string(data)}, " ")
}
