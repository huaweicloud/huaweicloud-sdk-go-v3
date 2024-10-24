package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeEndpointRequest Request Object
type SubscribeEndpointRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// Endpoint的ID
	EndpointId string `json:"endpoint_id"`
}

func (o SubscribeEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeEndpointRequest struct{}"
	}

	return strings.Join([]string{"SubscribeEndpointRequest", string(data)}, " ")
}
