package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetAuthenticationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id"`

	Body *ResetAuthenticationRequestBody `json:"body,omitempty"`
}

func (o ResetAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"ResetAuthenticationRequest", string(data)}, " ")
}
