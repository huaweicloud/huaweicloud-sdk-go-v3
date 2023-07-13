package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssuesSfV4Response Response Object
type ListIssuesSfV4Response struct {

	// 工作项
	Issues *[]IssueItemSfV4 `json:"issues,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssuesSfV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuesSfV4Response struct{}"
	}

	return strings.Join([]string{"ListIssuesSfV4Response", string(data)}, " ")
}
