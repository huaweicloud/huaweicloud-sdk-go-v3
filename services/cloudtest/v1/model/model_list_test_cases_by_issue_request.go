package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCasesByIssueRequest Request Object
type ListTestCasesByIssueRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 需求id
	IssueId string `json:"issue_id"`

	Body *QueryTestCasesByIssueInfo `json:"body,omitempty"`
}

func (o ListTestCasesByIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesByIssueRequest struct{}"
	}

	return strings.Join([]string{"ListTestCasesByIssueRequest", string(data)}, " ")
}
