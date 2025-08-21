package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Pagination 分页
type Pagination struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目
	Limit *int32 `json:"limit,omitempty"`
}

func (o Pagination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pagination struct{}"
	}

	return strings.Join([]string{"Pagination", string(data)}, " ")
}
