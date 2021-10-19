package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListServiceConnectionsResponse struct {
	// 连接列表。

	Connections *[]Endpoints `json:"connections,omitempty"`
	// 满足查询条件的终端节点总条数，不受分 页（即limit、offset参数）影响。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServiceConnectionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsResponse", string(data)}, " ")
}
