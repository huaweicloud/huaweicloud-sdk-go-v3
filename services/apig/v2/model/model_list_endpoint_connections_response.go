package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointConnectionsResponse Response Object
type ListEndpointConnectionsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 连接列表
	Connections *[]EndpointConnection `json:"connections,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointConnectionsResponse", string(data)}, " ")
}
