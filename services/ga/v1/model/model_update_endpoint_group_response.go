package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointGroupResponse Response Object
type UpdateEndpointGroupResponse struct {
	EndpointGroup *EndpointGroupDetail `json:"endpoint_group,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEndpointGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointGroupResponse", string(data)}, " ")
}
