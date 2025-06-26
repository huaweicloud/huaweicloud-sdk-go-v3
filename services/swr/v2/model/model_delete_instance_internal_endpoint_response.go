package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceInternalEndpointResponse Response Object
type DeleteInstanceInternalEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceInternalEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceInternalEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceInternalEndpointResponse", string(data)}, " ")
}
