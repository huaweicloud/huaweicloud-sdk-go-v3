package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEndpointRequest Request Object
type ShowEndpointRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`
}

func (o ShowEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointRequest struct{}"
	}

	return strings.Join([]string{"ShowEndpointRequest", string(data)}, " ")
}
