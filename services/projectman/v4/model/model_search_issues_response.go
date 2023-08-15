package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchIssuesResponse Response Object
type SearchIssuesResponse struct {

	// 工作项信息列表
	IssueList *[]WorkTableIssuseListResponseBodyIssueList `json:"issue_list,omitempty"`

	// 工作项总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIssuesResponse struct{}"
	}

	return strings.Join([]string{"SearchIssuesResponse", string(data)}, " ")
}
