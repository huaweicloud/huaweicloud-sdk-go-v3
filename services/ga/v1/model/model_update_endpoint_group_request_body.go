package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointGroupRequestBody update Endpoint Group request
type UpdateEndpointGroupRequestBody struct {
	EndpointGroup *UpdateEndpointGroupOption `json:"endpoint_group"`
}

func (o UpdateEndpointGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointGroupRequestBody", string(data)}, " ")
}
