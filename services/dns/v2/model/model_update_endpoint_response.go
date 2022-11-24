package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEndpointResponse struct {
	Endpoints      *[]EndpointResp `json:"endpoints,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointResponse", string(data)}, " ")
}
