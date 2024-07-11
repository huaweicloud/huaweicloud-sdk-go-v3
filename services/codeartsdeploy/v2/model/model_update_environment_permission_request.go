package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvironmentPermissionRequest Request Object
type UpdateEnvironmentPermissionRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`

	Body *EnvironmentPermissionV2Body `json:"body,omitempty"`
}

func (o UpdateEnvironmentPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentPermissionRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentPermissionRequest", string(data)}, " ")
}
