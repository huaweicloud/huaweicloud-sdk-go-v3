package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTestCaseCommentRequest Request Object
type AddTestCaseCommentRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例uri
	TestcaseId string `json:"testcase_id"`

	Body *TestCaseCommentInfo `json:"body,omitempty"`
}

func (o AddTestCaseCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTestCaseCommentRequest struct{}"
	}

	return strings.Join([]string{"AddTestCaseCommentRequest", string(data)}, " ")
}
