package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIssueAssociatedCommitsResponse struct {

	// 提交记录列表
	Commits *[]CommitRecordDetail `json:"commits,omitempty" xml:"commits"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssueAssociatedCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueAssociatedCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueAssociatedCommitsResponse", string(data)}, " ")
}
