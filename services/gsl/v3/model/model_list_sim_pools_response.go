package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSimPoolsResponse struct {
	// 每页记录数

	Limit *int64 `json:"limit,omitempty"`
	// 页码

	Offset *int64 `json:"offset,omitempty"`
	// 当前查询条件的流量池总数

	Count *int64 `json:"count,omitempty"`
	// 当前页的流量池记录列表

	Pools          *[]SimPoolVo `json:"pools,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListSimPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListSimPoolsResponse", string(data)}, " ")
}
