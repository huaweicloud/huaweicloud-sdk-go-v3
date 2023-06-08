package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVpcEndpointRequest struct {
	Body *CreateVpcEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateVpcEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointRequest", string(data)}, " ")
}
