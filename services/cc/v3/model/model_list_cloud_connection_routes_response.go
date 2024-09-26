package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionRoutesResponse Response Object
type ListCloudConnectionRoutesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 云连接路由实例列表。
	CloudConnectionRoutes []CloudConnectionRoute `json:"cloud_connection_routes"`
	HttpStatusCode        int                    `json:"-"`
}

func (o ListCloudConnectionRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionRoutesResponse", string(data)}, " ")
}
