package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTicketParams struct {

	// 字符串搜索条件
	StringFilters []ObjectFilter `json:"string_filters"`

	SortFilter *ObjectFilter `json:"sort_filter,omitempty"`
}

func (o ListTicketParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketParams struct{}"
	}

	return strings.Join([]string{"ListTicketParams", string(data)}, " ")
}
