package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceInstanceRequest Request Object
type DeleteServiceInstanceRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 实例的Id
	InstanceId string `json:"instance_id"`
}

func (o DeleteServiceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceInstanceRequest", string(data)}, " ")
}
