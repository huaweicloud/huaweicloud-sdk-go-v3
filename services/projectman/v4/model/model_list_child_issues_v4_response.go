package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListChildIssuesV4Response struct {

	// 工作项列表
	Issues *[]IssueResponseV4 `json:"issues,omitempty" xml:"issues"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListChildIssuesV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChildIssuesV4Response struct{}"
	}

	return strings.Join([]string{"ListChildIssuesV4Response", string(data)}, " ")
}
