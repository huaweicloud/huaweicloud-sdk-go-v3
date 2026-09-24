package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternalEndpointConnectionsRequest Request Object
type UpdateInternalEndpointConnectionsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateInternalEndpointConnectionsRequestBody `json:"body,omitempty"`
}

func (o UpdateInternalEndpointConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointConnectionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointConnectionsRequest", string(data)}, " ")
}
