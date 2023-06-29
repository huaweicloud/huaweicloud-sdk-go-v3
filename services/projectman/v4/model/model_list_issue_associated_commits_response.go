package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueAssociatedCommitsResponse Response Object
type ListIssueAssociatedCommitsResponse struct {

	// 提交记录列表
	Commits *[]CommitRecordDetail `json:"commits,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssueAssociatedCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueAssociatedCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueAssociatedCommitsResponse", string(data)}, " ")
}
