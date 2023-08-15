package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsResponse Response Object
type ListConnectionsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]ConnectionInfo `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsResponse", string(data)}, " ")
}
