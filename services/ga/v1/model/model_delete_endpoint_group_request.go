package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointGroupRequest Request Object
type DeleteEndpointGroupRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`
}

func (o DeleteEndpointGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointGroupRequest", string(data)}, " ")
}
