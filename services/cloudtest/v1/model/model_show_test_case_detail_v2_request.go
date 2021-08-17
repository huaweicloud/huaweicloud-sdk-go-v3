package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTestCaseDetailV2Request struct {
	// 用例编号，长度为[3-128]位字符

	TestcaseNumber *string `json:"testcase_number,omitempty"`
	// 测试用例唯一标识，固定长度32位字符

	TestcaseId *string `json:"testcase_id,omitempty"`
}

func (o ShowTestCaseDetailV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTestCaseDetailV2Request struct{}"
	}

	return strings.Join([]string{"ShowTestCaseDetailV2Request", string(data)}, " ")
}
