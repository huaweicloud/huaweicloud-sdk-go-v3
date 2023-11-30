package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublisherResponse Response Object
type ListPublisherResponse struct {

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询总数
	Total *int64 `json:"total,omitempty"`

	// 数据列表
	Data           *[]PublisherVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublisherResponse struct{}"
	}

	return strings.Join([]string{"ListPublisherResponse", string(data)}, " ")
}
