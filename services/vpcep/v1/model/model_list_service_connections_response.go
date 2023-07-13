package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceConnectionsResponse Response Object
type ListServiceConnectionsResponse struct {

	// 连接列表。
	Connections *[]ConnectionEndpoints `json:"connections,omitempty"`

	// 满足查询条件的终端节点总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServiceConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsResponse", string(data)}, " ")
}
