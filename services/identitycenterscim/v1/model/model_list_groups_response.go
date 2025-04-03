package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsResponse Response Object
type ListGroupsResponse struct {

	// 总结果数
	TotalResults *int32 `json:"totalResults,omitempty"`

	// 每页的元素个数
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`

	// 起始索引
	StartIndex *string `json:"startIndex,omitempty"`

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 包含用户组信息的列表
	Resources      *[]GetGroupResp `json:"Resources,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsResponse", string(data)}, " ")
}
