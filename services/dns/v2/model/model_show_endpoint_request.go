package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEndpointRequest struct {

	// 待查询终端节点的ID。
	EndpointId string `json:"endpoint_id"`
}

func (o ShowEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointRequest struct{}"
	}

	return strings.Join([]string{"ShowEndpointRequest", string(data)}, " ")
}
