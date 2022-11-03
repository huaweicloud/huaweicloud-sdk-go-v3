package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// update Endpoint request
type UpdateEndpointRequestBody struct {
	Endpoint *UpdateEndpointOption `json:"endpoint"`
}

func (o UpdateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequestBody", string(data)}, " ")
}
