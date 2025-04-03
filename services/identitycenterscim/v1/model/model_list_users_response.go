package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersResponse Response Object
type ListUsersResponse struct {

	// 总结果数
	TotalResults *int32 `json:"totalResults,omitempty"`

	// 每页的元素个数
	ItemsPerPage *string `json:"itemsPerPage,omitempty"`

	// 起始索引
	StartIndex *string `json:"startIndex,omitempty"`

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 包含用户信息的列表
	Resources      *[]GetUserResp `json:"Resources,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
