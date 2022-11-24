package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEndpointResponse struct {
	Endpoint *EndpointDetail `json:"endpoint,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointResponse struct{}"
	}

	return strings.Join([]string{"ShowEndpointResponse", string(data)}, " ")
}
