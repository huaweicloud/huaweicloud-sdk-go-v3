package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharedConnectionsRequest Request Object
type ListSharedConnectionsRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 搜索关键字
	Keywords *string `json:"keywords,omitempty"`

	// 当前页码
	CurPage *string `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *string `json:"per_page,omitempty"`
}

func (o ListSharedConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedConnectionsRequest", string(data)}, " ")
}
