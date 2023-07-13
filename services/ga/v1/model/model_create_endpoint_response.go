package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointResponse Response Object
type CreateEndpointResponse struct {
	Endpoint *EndpointDetail `json:"endpoint,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}
