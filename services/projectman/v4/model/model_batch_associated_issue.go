package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociatedIssue 关联工作项详情
type BatchAssociatedIssue struct {

	// 关联工作项ID
	SourceIssueId *int32 `json:"source_issue_id,omitempty"`

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 工作项ID
	IssueId *int32 `json:"issue_id,omitempty"`

	Project *SimpleProject `json:"project,omitempty"`

	User *SimpleUser `json:"user,omitempty"`

	Status *StatusVo `json:"status,omitempty"`
}

func (o BatchAssociatedIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociatedIssue struct{}"
	}

	return strings.Join([]string{"BatchAssociatedIssue", string(data)}, " ")
}
