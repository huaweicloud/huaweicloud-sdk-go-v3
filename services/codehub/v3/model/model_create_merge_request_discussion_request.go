package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeRequestDiscussionRequest Request Object
type CreateMergeRequestDiscussionRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 合并请求iid
	MergeRequestIid int32 `json:"merge_request_iid"`

	Body *CreateMergeRequestDiscussionBodyDto `json:"body,omitempty"`
}

func (o CreateMergeRequestDiscussionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionRequest struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionRequest", string(data)}, " ")
}
