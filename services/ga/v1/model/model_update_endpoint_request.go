package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointRequest Request Object
type UpdateEndpointRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`

	Body *UpdateEndpointRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequest", string(data)}, " ")
}
