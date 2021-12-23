package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteComponentRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 是否强制删除。

	Force *bool `json:"force,omitempty"`
}

func (o DeleteComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentRequest struct{}"
	}

	return strings.Join([]string{"DeleteComponentRequest", string(data)}, " ")
}
