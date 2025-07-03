package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTicketParamsWithPageV2 带分页信息的查询
type ListTicketParamsWithPageV2 struct {

	// 返回字段
	Fields *[]string `json:"fields,omitempty"`

	// 批量计数
	CountFilters *[]ListTicketParamsV2CountFilters `json:"count_filters,omitempty"`

	// 字符串搜索条件
	StringFilters *[]ObjectFilterV2 `json:"string_filters,omitempty"`

	GroupByFilter *ObjectFilterV2 `json:"group_by_filter,omitempty"`

	// 整形过滤器
	IntFilters *[]ObjectFilterV2 `json:"int_filters,omitempty"`

	SortFilter *ObjectFilterV2 `json:"sort_filter,omitempty"`

	// 表达式，对复杂表达式进行组装，可以用英文括号()、与&、或|。示例：( filterName1 & filterName2) | filterName3。filterName*：取自string_filters ObjectFilter.name。如果为空，string_filters中的条件默认是与的关系
	Condition *string `json:"condition,omitempty"`

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

func (o ListTicketParamsWithPageV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketParamsWithPageV2 struct{}"
	}

	return strings.Join([]string{"ListTicketParamsWithPageV2", string(data)}, " ")
}
