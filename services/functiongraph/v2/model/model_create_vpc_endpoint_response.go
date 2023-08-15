package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcEndpointResponse Response Object
type CreateVpcEndpointResponse struct {

	// Ip列表
	Endpoints *[]string `json:"endpoints,omitempty"`

	// 域名地址
	Address        *string `json:"address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVpcEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointResponse", string(data)}, " ")
}
