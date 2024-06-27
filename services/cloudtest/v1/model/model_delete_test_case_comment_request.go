package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTestCaseCommentRequest Request Object
type DeleteTestCaseCommentRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例uri
	TestcaseId string `json:"testcase_id"`

	// 评论uri
	CommentId string `json:"comment_id"`

	// 分支或者测试计划uri
	VersionUri *string `json:"version_uri,omitempty"`
}

func (o DeleteTestCaseCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTestCaseCommentRequest struct{}"
	}

	return strings.Join([]string{"DeleteTestCaseCommentRequest", string(data)}, " ")
}
