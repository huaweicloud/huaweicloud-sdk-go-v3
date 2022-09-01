package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeInstanceRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id" xml:"application_id"`

	// 组件ID。
	ComponentId string `json:"component_id" xml:"component_id"`

	// 组件实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *InstanceModify `json:"body,omitempty" xml:"body"`
}

func (o ChangeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceRequest", string(data)}, " ")
}
