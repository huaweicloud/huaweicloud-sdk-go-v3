package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 关联工作项详情
type AssociateIssueDetail struct {

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 工作项ID
	IssueId *int32 `json:"issue_id,omitempty"`

	Project *SimpleProject `json:"project,omitempty"`

	User *SimpleUser `json:"user,omitempty"`

	Status *StatusVo `json:"status,omitempty"`
}

func (o AssociateIssueDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssueDetail struct{}"
	}

	return strings.Join([]string{"AssociateIssueDetail", string(data)}, " ")
}
