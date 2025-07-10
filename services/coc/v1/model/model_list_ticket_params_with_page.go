package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTicketParamsWithPage 带分页信息的查询
type ListTicketParamsWithPage struct {

	// 字符串搜索条件，可根据该条件搜索到具体的工单。
	StringFilters []ObjectFilter `json:"string_filters"`

	SortFilter *ObjectFilter `json:"sort_filter,omitempty"`

	// 当前页
	Page *int32 `json:"page,omitempty"`

	// 每页数量
	PerPage *int32 `json:"per_page,omitempty"`

	// 是否返回所有数据
	ContainAll *bool `json:"contain_all,omitempty"`

	// 是否返回总数
	ContainTotal *bool `json:"contain_total,omitempty"`

	// 是否包含子单
	ContainSubTicket *bool `json:"contain_sub_ticket,omitempty"`

	// 搜索的工单类型
	TicketTypes *[]string `json:"ticket_types,omitempty"`
}

func (o ListTicketParamsWithPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketParamsWithPage struct{}"
	}

	return strings.Join([]string{"ListTicketParamsWithPage", string(data)}, " ")
}
