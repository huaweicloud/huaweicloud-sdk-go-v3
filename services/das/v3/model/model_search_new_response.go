package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchNewResponse Response Object
type SearchNewResponse struct {

	// SQL列表
	SqlItemDtoList *[]SqlItemDto `json:"sql_item_dto_list,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchNewResponse struct{}"
	}

	return strings.Join([]string{"SearchNewResponse", string(data)}, " ")
}
