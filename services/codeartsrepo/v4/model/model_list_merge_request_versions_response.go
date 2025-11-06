package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestVersionsResponse Response Object
type ListMergeRequestVersionsResponse struct {

	// **参数解释：** diff主键ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** head commit节点。
	HeadCommitSha *string `json:"head_commit_sha,omitempty"`

	// **参数解释：** base commit节点。
	BaseCommitSha *string `json:"base_commit_sha,omitempty"`

	// **参数解释：** tart commit节点。
	StartCommitSha *string `json:"start_commit_sha,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** MR主键ID。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 状态。
	State *string `json:"state,omitempty"`

	// **参数解释：** diff大小。
	RealSize       *string `json:"real_size,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeRequestVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestVersionsResponse", string(data)}, " ")
}
