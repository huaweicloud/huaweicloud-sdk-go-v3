package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointResponse Response Object
type CreateEndpointResponse struct {
	Endpoint       *EndpointObjResp `json:"endpoint,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}
