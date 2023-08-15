package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionsResponse Response Object
type ListCloudConnectionsResponse struct {

	// 云连接实例列表。
	CloudConnections *[]CloudConnection `json:"cloud_connections,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsResponse", string(data)}, " ")
}
