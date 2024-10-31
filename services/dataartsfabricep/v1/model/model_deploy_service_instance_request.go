package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployServiceInstanceRequest Request Object
type DeployServiceInstanceRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	Body *ServiceInstanceInput `json:"body,omitempty"`
}

func (o DeployServiceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployServiceInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeployServiceInstanceRequest", string(data)}, " ")
}
