package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointGroupResponse Response Object
type DeleteEndpointGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointGroupResponse", string(data)}, " ")
}
