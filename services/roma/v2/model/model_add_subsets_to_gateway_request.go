package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddSubsetsToGatewayRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备ID

	DeviceId int32 `json:"device_id"`

	Body *AddSubsetsToGatewayRequestBody `json:"body,omitempty"`
}

func (o AddSubsetsToGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubsetsToGatewayRequest struct{}"
	}

	return strings.Join([]string{"AddSubsetsToGatewayRequest", string(data)}, " ")
}
