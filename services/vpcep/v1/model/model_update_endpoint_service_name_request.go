package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServiceNameRequest Request Object
type UpdateEndpointServiceNameRequest struct {

	// 终端节点服务ID
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *UpdateEndpointServiceNameRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointServiceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameRequest", string(data)}, " ")
}
