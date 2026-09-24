package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternalEndpointConnectionsResponse Response Object
type ListInternalEndpointConnectionsResponse struct {

	// 连接列表
	Connections *[]ConnectionItem `json:"connections,omitempty"`

	// 满足查询条件的连接总条数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInternalEndpointConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternalEndpointConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListInternalEndpointConnectionsResponse", string(data)}, " ")
}
