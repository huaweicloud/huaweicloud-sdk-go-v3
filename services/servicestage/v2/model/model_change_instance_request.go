package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeInstanceRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`

	Body *InstanceModify `json:"body,omitempty"`
}

func (o ChangeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceRequest", string(data)}, " ")
}
