package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRepoMembersRequest struct {
	// 第几页

	PageIndex *int32 `json:"page_index,omitempty"`
	// 每页显示size

	PageSize *int32 `json:"page_size,omitempty"`
	// 仓库uuid

	RepositoryUuid string `json:"repository_uuid"`
	// 搜索关键字

	Subject *string `json:"subject,omitempty"`
}

func (o ListRepoMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoMembersRequest struct{}"
	}

	return strings.Join([]string{"ListRepoMembersRequest", string(data)}, " ")
}
