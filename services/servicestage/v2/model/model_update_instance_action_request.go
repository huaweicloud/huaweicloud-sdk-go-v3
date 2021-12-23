package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceActionRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`

	Body *InstanceAction `json:"body,omitempty"`
}

func (o UpdateInstanceActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceActionRequest", string(data)}, " ")
}
