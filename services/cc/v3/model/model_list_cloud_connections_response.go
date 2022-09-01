package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCloudConnectionsResponse struct {

	// 云连接实例列表。
	CloudConnections *[]CloudConnection `json:"cloud_connections,omitempty" xml:"cloud_connections"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsResponse", string(data)}, " ")
}
