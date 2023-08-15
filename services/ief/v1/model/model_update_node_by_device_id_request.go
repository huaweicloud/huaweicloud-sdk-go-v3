package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeByDeviceIdRequest Request Object
type UpdateNodeByDeviceIdRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

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
