package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEndpointRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	Body *CreateEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequest", string(data)}, " ")
}
