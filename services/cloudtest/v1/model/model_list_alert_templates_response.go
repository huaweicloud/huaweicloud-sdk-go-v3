package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertTemplatesResponse Response Object
type ListAlertTemplatesResponse struct {
	EndRow *int64 `json:"end_row,omitempty"`

	HasNextPage *bool `json:"has_next_page,omitempty"`

	HasPreviousPage *bool `json:"has_previous_page,omitempty"`

	IsFirstPage *bool `json:"is_first_page,omitempty"`

	IsLastPage *bool `json:"is_last_page,omitempty"`

	List *[]AlertTemplateVo `json:"list,omitempty"`

	NavigateFirstPage *int32 `json:"navigateFirstPage,omitempty"`

	NavigateLastPage *int32 `json:"navigateLastPage,omitempty"`

	NavigatePages *int32 `json:"navigatePages,omitempty"`

	NavigatepageNums *[]int32 `json:"navigatepageNums,omitempty"`

	NextPage *int32 `json:"next_page,omitempty"`

	PageNum *int32 `json:"page_num,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	PrePage *int32 `json:"prePage,omitempty"`

	Size *int32 `json:"size,omitempty"`

	StartRow *int64 `json:"startRow,omitempty"`

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
