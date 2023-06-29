package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointRequest Request Object
type CreateEndpointRequest struct {
	Body *EndpointCreateReq `json:"body,omitempty"`
}

func (o CreateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequest", string(data)}, " ")
}
