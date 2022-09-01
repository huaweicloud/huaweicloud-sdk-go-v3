package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIssueAssociatedCommitsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 工作项ID
	IssueId int32 `json:"issue_id" xml:"issue_id"`

	// 查询类型：commit（提交记录） || branch（分支记录）
	Type *string `json:"type,omitempty" xml:"type"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListIssueAssociatedCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueAssociatedCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListIssueAssociatedCommitsRequest", string(data)}, " ")
}
