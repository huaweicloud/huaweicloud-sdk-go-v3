package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvironmentsResponse struct {
	// 环境分组列表

	Environments *[]Environment `json:"environments,omitempty"`
	// 偏移量，表示从此偏移量开始查询，offset大于等于0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量,最大支持200条

	Limit *int64 `json:"limit,omitempty"`
	// 环境分组总条数

	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvironmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsResponse", string(data)}, " ")
}
