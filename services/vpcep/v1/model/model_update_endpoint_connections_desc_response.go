package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointConnectionsDescResponse Response Object
type UpdateEndpointConnectionsDescResponse struct {

	// 连接列表
	Connections    *[]ConnectionEndpoints `json:"connections,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateEndpointConnectionsDescResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointConnectionsDescResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointConnectionsDescResponse", string(data)}, " ")
}
