package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowInstanceDetailRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	// 组件ID。
	ComponentId string `json:"component_id"`

	// 组件实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailRequest", string(data)}, " ")
}
