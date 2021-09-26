package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTestCaseRequest struct {
	// 测试用例唯一标识，固定长度32位字符

	TestcaseId string `json:"testcase_id"`

	Body *UpdateTestCaseRequestBody `json:"body,omitempty"`
}

func (o UpdateTestCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTestCaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseRequest", string(data)}, " ")
}