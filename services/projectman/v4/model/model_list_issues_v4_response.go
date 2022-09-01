package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIssuesV4Response struct {

	// 工作项列表
	Issues *[]ListIssueItemResponse `json:"issues,omitempty" xml:"issues"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssuesV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuesV4Response struct{}"
	}

	return strings.Join([]string{"ListIssuesV4Response", string(data)}, " ")
}
