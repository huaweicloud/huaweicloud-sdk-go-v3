package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEndpointRequest Request Object
type ShowEndpointRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// Endpoint的ID
	EndpointId string `json:"endpoint_id"`

	// 查询endpoint信息的版本：CURRENT-当前版本，PREVIOUS-上个版本
	Version *string `json:"version,omitempty"`
}

func (o ShowEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointRequest struct{}"
	}

	return strings.Join([]string{"ShowEndpointRequest", string(data)}, " ")
}
