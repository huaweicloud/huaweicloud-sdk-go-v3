package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeRequestDiscussionNoteRequest Request Object
type CreateMergeRequestDiscussionNoteRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 合并请求iid
	MergeRequestIid int32 `json:"merge_request_iid"`

	// 评论id
	DiscussionId string `json:"discussion_id"`

	Body *CreateMergeRequestDiscussionNoteDto `json:"body,omitempty"`
}

func (o CreateMergeRequestDiscussionNoteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionNoteRequest struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionNoteRequest", string(data)}, " ")
}
