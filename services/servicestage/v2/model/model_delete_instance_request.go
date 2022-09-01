package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id" xml:"application_id"`

	// 组件ID。
	ComponentId string `json:"component_id" xml:"component_id"`

	// 组件实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 是否强制删除。
	Force *bool `json:"force,omitempty" xml:"force"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}
