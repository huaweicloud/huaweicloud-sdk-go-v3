package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEndpointGroupRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`
}

func (o ShowEndpointGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowEndpointGroupRequest", string(data)}, " ")
}
