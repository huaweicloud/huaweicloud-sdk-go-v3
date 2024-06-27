package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsingGetResponse Response Object
type ListUsingGetResponse struct {

	// 返回结果
	List *[]DashboardDto `json:"list,omitempty"`

	// 页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总条数
	TotalSize      *int64 `json:"total_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListUsingGetResponse", string(data)}, " ")
}
