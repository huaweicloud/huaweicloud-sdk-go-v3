package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointRequestBody create Endpoint request
type CreateEndpointRequestBody struct {
	Endpoint *CreateEndpointOption `json:"endpoint"`
}

func (o CreateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequestBody", string(data)}, " ")
}
