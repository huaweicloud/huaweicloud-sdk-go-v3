package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuesResponse Response Object
type ListQueuesResponse struct {

	// 当前显示数量
	Size *int32 `json:"size,omitempty"`

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// 查询详情
	Items          *[]QueueDetails `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuesResponse", string(data)}, " ")
}
