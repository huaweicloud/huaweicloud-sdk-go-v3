package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointRequest Request Object
type DeleteEndpointRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`
}

func (o DeleteEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointRequest", string(data)}, " ")
}
