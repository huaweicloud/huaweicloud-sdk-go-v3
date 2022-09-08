package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssociatedIssuesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项ID
	IssueId int32 `json:"issue_id"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAssociatedIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedIssuesRequest struct{}"
	}

	return strings.Join([]string{"ListAssociatedIssuesRequest", string(data)}, " ")
}
