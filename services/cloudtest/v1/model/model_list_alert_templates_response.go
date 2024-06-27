package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertTemplatesResponse Response Object
type ListAlertTemplatesResponse struct {

	// 当前页面最后一个元素在数据库中的行号
	EndRow *int64 `json:"end_row,omitempty"`

	// 是否有下一页
	HasNextPage *bool `json:"has_next_page,omitempty"`

	// 是否有前一页
	HasPreviousPage *bool `json:"has_previous_page,omitempty"`

	// 是否为第一页
	IsFirstPage *bool `json:"is_first_page,omitempty"`

	// 是否为最后一页
	IsLastPage *bool `json:"is_last_page,omitempty"`

	// 返回结果
	List *[]AlertTemplateVo `json:"list,omitempty"`

	// 导航条上的第一页
	NavigateFirstPage *int32 `json:"navigate_first_page,omitempty"`

	// 导航条上的最后一页
	NavigateLastPage *int32 `json:"navigate_last_page,omitempty"`

	// 导航页码数
	NavigatePages *int32 `json:"navigate_pages,omitempty"`

	// 所有导航页号
	NavigatepageNums *[]int32 `json:"navigatepage_nums,omitempty"`

	// 下一页
	NextPage *int32 `json:"next_page,omitempty"`

	// 当前页
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页的数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 总页数
	Pages *int32 `json:"pages,omitempty"`

	// 前一页
	PrePage *int32 `json:"pre_page,omitempty"`

	// 当前页的数量
	Size *int32 `json:"size,omitempty"`

	// 当前页面第一个元素在数据库中的行号
	StartRow *int64 `json:"start_row,omitempty"`

	// 总条数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlertTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertTemplatesResponse", string(data)}, " ")
}
