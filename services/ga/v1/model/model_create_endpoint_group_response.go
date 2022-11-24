package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEndpointGroupResponse struct {
	EndpointGroup *EndpointGroupDetail `json:"endpoint_group,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEndpointGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointGroupResponse", string(data)}, " ")
}
