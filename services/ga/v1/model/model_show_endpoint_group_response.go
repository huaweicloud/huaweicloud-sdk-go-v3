package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEndpointGroupResponse struct {
	EndpointGroup *EndpointGroupDetail `json:"endpoint_group,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEndpointGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowEndpointGroupResponse", string(data)}, " ")
}
