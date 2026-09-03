package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdsConnectionStatResponse Response Object
type ShowDdsConnectionStatResponse struct {

	// 总连接数
	TotalConnections *int32 `json:"total_connections,omitempty"`

	// 内部连接总数
	TotalInnerConnections *int32 `json:"total_inner_connections,omitempty"`

	// 外部连接总数
	TotalOuterConnections *int32 `json:"total_outer_connections,omitempty"`

	// 内部连接详情
	InnerConnections *[]ConnectionDetail `json:"inner_connections,omitempty"`

	// 外部连接详情
	OuterConnections *[]ConnectionDetail `json:"outer_connections,omitempty"`

	// 内部连接详情总数
	InnerConnectionsSize *int32 `json:"inner_connections_size,omitempty"`

	// 外部连接详情总数
	OuterConnectionsSize *int32 `json:"outer_connections_size,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowDdsConnectionStatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdsConnectionStatResponse struct{}"
	}

	return strings.Join([]string{"ShowDdsConnectionStatResponse", string(data)}, " ")
}
