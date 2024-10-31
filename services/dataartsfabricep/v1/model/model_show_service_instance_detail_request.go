package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServiceInstanceDetailRequest Request Object
type ShowServiceInstanceDetailRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 实例的Id
	InstanceId string `json:"instance_id"`
}

func (o ShowServiceInstanceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowServiceInstanceDetailRequest", string(data)}, " ")
}
