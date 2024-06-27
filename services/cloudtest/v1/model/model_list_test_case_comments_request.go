package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseCommentsRequest Request Object
type ListTestCaseCommentsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例uri
	TestcaseId string `json:"testcase_id"`

	// 页数
	PageNo int32 `json:"page_no"`

	// 页数大小
	PageSize int32 `json:"page_size"`

	// 分支或者测试计划uri
	VersionUri *string `json:"version_uri,omitempty"`
}

func (o ListTestCaseCommentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseCommentsRequest struct{}"
	}

	return strings.Join([]string{"ListTestCaseCommentsRequest", string(data)}, " ")
}
