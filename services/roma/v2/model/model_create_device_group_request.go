package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeviceGroupRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *CreateDeviceGroupRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateDeviceGroupRequest", string(data)}, " ")
}
