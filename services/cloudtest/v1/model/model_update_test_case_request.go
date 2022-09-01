package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTestCaseRequest struct {

	// 测试用例唯一标识，固定长度32位字符
	TestcaseId string `json:"testcase_id" xml:"testcase_id"`

	Body *UpdateTestCaseRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseRequest", string(data)}, " ")
}
