package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMsgInfosUsingResponse Response Object
type ListMsgInfosUsingResponse struct {

	// 返回结果
	List *[]MsgInfoVo `json:"list,omitempty"`

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

func (o ListMsgInfosUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMsgInfosUsingResponse struct{}"
	}

	return strings.Join([]string{"ListMsgInfosUsingResponse", string(data)}, " ")
}
