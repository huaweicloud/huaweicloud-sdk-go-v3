package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfo struct {
	EndRow *int64 `json:"endRow,omitempty"`

	HasNextPage *bool `json:"hasNextPage,omitempty"`

	HasPreviousPage *bool `json:"hasPreviousPage,omitempty"`

	IsFirstPage *bool `json:"isFirstPage,omitempty"`

	IsLastPage *bool `json:"isLastPage,omitempty"`

	List *[]interface{} `json:"list,omitempty"`

	NavigateFirstPage *int32 `json:"navigateFirstPage,omitempty"`

	NavigateLastPage *int32 `json:"navigateLastPage,omitempty"`

	NavigatePages *int32 `json:"navigatePages,omitempty"`

	NavigatepageNums *[]int32 `json:"navigatepageNums,omitempty"`

	NextPage *int32 `json:"nextPage,omitempty"`

	PageNum *int32 `json:"pageNum,omitempty"`

	PageSize *int32 `json:"pageSize,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	PrePage *int32 `json:"prePage,omitempty"`

	Size *int32 `json:"size,omitempty"`

	StartRow *int64 `json:"startRow,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
