package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowConnectionStatisticsResponse struct {
	// 总连接数，包括内部连接与外部连接。

	TotalConnections *int32 `json:"total_connections,omitempty"`
	// 内部总连接数。

	TotalInnerConnections *int32 `json:"total_inner_connections,omitempty"`
	// 外部总连接数。

	TotalOuterConnections *int32 `json:"total_outer_connections,omitempty"`
	// 内部连接统计信息数组，最大记录数为200条。

	InnerConnections *[]QueryConnectionsResponse `json:"inner_connections,omitempty"`
	// 外部连接统计信息数组，最大记录数为200条。

	OuterConnections *[]QueryConnectionsResponse `json:"outer_connections,omitempty"`
	HttpStatusCode   int                         `json:"-"`
}

func (o ShowConnectionStatisticsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConnectionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionStatisticsResponse", string(data)}, " ")
}
