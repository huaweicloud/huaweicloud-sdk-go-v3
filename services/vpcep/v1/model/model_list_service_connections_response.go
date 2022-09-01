package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListServiceConnectionsResponse struct {

	// 连接列表。
	Connections *[]Connection `json:"connections,omitempty" xml:"connections"`

	// 满足查询条件的终端节点总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServiceConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsResponse", string(data)}, " ")
}
