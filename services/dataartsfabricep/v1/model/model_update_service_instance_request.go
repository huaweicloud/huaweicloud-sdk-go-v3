package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceInstanceRequest Request Object
type UpdateServiceInstanceRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 实例的Id
	InstanceId string `json:"instance_id"`

	Body *UpdateServiceInstanceInput `json:"body,omitempty"`
}

func (o UpdateServiceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServiceInstanceRequest", string(data)}, " ")
}
