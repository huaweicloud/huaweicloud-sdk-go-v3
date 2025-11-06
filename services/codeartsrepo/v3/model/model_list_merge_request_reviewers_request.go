package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestReviewersRequest Request Object
type ListMergeRequestReviewersRequest struct {

	// 仓库的主键id
	RepositoryId string `json:"repository_id"`

	// 合并请求的短id
	MergeRequestIid int32 `json:"merge_request_iid"`

	// 页码
	Page *int32 `json:"page,omitempty"`

	// 每页条数
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListMergeRequestReviewersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestReviewersRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestReviewersRequest", string(data)}, " ")
}
