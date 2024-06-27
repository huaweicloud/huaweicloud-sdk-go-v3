package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLinesUsingResponse Response Object
type ListLinesUsingResponse struct {

	// 返回结果
	List *[]SubTaskCaseSuccessLineVo `json:"list,omitempty"`

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

func (o ListLinesUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinesUsingResponse struct{}"
	}

	return strings.Join([]string{"ListLinesUsingResponse", string(data)}, " ")
}
