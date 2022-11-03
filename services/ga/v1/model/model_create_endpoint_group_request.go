package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEndpointGroupRequest struct {
	Body *CreateEndpointGroupRequestBody `json:"body,omitempty"`
}

func (o CreateEndpointGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointGroupRequest", string(data)}, " ")
}
