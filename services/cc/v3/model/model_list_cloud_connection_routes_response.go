package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCloudConnectionRoutesResponse struct {
	// 云连接路由实例列表。

	CloudConnectionRoutes *[]CloudConnectionRoute `json:"cloud_connection_routes,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`
	// 请求ID。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudConnectionRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionRoutesResponse", string(data)}, " ")
}
