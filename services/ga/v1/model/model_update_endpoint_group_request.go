package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointGroupRequest Request Object
type UpdateEndpointGroupRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	Body *UpdateEndpointGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointGroupRequest", string(data)}, " ")
}
