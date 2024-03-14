package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointConnectionsDescRequest Request Object
type UpdateEndpointConnectionsDescRequest struct {

	// 终端节点服务ID
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *UpdateEndpointConnectionsDescRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointConnectionsDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointConnectionsDescRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointConnectionsDescRequest", string(data)}, " ")
}
