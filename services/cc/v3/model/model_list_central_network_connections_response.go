package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkConnectionsResponse Response Object
type ListCentralNetworkConnectionsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 中心网络连接列表。
	CentralNetworkConnections []CentralNetworkConnection `json:"central_network_connections"`
	HttpStatusCode            int                        `json:"-"`
}

func (o ListCentralNetworkConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkConnectionsResponse", string(data)}, " ")
}
