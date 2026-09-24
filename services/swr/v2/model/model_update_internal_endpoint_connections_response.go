package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternalEndpointConnectionsResponse Response Object
type UpdateInternalEndpointConnectionsResponse struct {

	// 终端节点连接列表
	Connections    *[]ConnectionItem `json:"connections,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateInternalEndpointConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointConnectionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointConnectionsResponse", string(data)}, " ")
}
