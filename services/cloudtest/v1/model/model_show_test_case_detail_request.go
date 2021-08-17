package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTestCaseDetailRequest struct {
	// 测试用例唯一标识，固定长度32位字符

	TestcaseId string `json:"testcase_id"`
}

func (o ShowTestCaseDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTestCaseDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTestCaseDetailRequest", string(data)}, " ")
}
