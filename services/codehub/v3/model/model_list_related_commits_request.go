package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedCommitsRequest Request Object
type ListRelatedCommitsRequest struct {

	// 仓库长id
	RepositoryUuid string `json:"repository_uuid"`

	// 关联工作项类型
	Type *int32 `json:"type,omitempty"`

	// 查询关键字
	Search *string `json:"search,omitempty"`

	// 页码
	Page *int32 `json:"page,omitempty"`

	// 每页数量
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListRelatedCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListRelatedCommitsRequest", string(data)}, " ")
}
