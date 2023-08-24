package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptOrRejectEndpointConnectionsResponse Response Object
type AcceptOrRejectEndpointConnectionsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 连接列表
	Connections *[]EndpointConnection `json:"connections,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptOrRejectEndpointConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointConnectionsResponse struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointConnectionsResponse", string(data)}, " ")
}
