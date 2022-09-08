package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBranchesByRepositoryIdRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 分页页数
	Page *string `json:"page,omitempty"`

	// 每页数据数
	PerPage *string `json:"per_page,omitempty"`

	// 匹配条件
	Match *string `json:"match,omitempty"`
}

func (o ListBranchesByRepositoryIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesByRepositoryIdRequest struct{}"
	}

	return strings.Join([]string{"ListBranchesByRepositoryIdRequest", string(data)}, " ")
}
