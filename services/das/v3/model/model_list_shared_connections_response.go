package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharedConnectionsResponse Response Object
type ListSharedConnectionsResponse struct {

	// 共享连接列表总数
	Total *int32 `json:"total,omitempty"`

	// 共享连接列表信息
	Data           *[]ConnSharedInfo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSharedConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListSharedConnectionsResponse", string(data)}, " ")
}
