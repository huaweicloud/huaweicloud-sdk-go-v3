package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthenticationRequest Request Object
type ShowAuthenticationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id"`
}

func (o ShowAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthenticationRequest", string(data)}, " ")
}
