package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcEndpointResponse Response Object
type CreateVpcEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVpcEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointResponse", string(data)}, " ")
}
