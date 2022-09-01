package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id" xml:"application_id"`

	// 组件ID。
	ComponentId string `json:"component_id" xml:"component_id"`

	Body *InstanceCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
