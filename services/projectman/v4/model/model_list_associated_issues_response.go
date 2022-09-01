package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssociatedIssuesResponse struct {

	// 关联的工作项列表
	Issues *[]AssociateIssueDetail `json:"issues,omitempty" xml:"issues"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAssociatedIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedIssuesResponse struct{}"
	}

	return strings.Join([]string{"ListAssociatedIssuesResponse", string(data)}, " ")
}
