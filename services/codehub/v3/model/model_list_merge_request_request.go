package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestRequest Request Object
type ListMergeRequestRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 匹配条件
	State *string `json:"state,omitempty"`

	// 分页页数
	Page *string `json:"page,omitempty"`

	// 每页数据数
	PerPage *string `json:"per_page,omitempty"`

	// 匹配条件
	Search *string `json:"search,omitempty"`
}

func (o ListMergeRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestRequest", string(data)}, " ")
}
