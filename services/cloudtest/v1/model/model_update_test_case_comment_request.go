package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestCaseCommentRequest Request Object
type UpdateTestCaseCommentRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例uri
	TestcaseId string `json:"testcase_id"`

	// 评论uri
	CommentId string `json:"comment_id"`

	Body *TestCaseCommentInfo `json:"body,omitempty"`
}

func (o UpdateTestCaseCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseCommentRequest", string(data)}, " ")
}
