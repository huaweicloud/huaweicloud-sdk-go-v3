package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionsByTagsResponse Response Object
type ListCloudConnectionsByTagsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 云连接实例列表。
	CloudConnections []CloudConnection `json:"cloud_connections"`
	HttpStatusCode   int               `json:"-"`
}

func (o ListCloudConnectionsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsByTagsResponse", string(data)}, " ")
}
