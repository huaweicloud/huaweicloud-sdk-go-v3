package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointGroupRequestBody create Endpoint Group request
type CreateEndpointGroupRequestBody struct {
	EndpointGroup *CreateEndpointGroupOption `json:"endpoint_group"`
}

func (o CreateEndpointGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointGroupRequestBody", string(data)}, " ")
}
