package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcEndpointResponse Response Object
type DeleteVpcEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcEndpointResponse", string(data)}, " ")
}
