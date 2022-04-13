package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeviceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateDeviceRequestBody `json:"body,omitempty"`
}

func (o CreateDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceRequest struct{}"
	}

	return strings.Join([]string{"CreateDeviceRequest", string(data)}, " ")
}
