package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointResponse Response Object
type DeleteEndpointResponse struct {
	Endpoint       *EndpointResp `json:"endpoint,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DeleteEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointResponse", string(data)}, " ")
}
