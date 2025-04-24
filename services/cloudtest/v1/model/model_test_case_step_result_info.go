package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseStepResultInfo 用例步骤结果信息
type TestCaseStepResultInfo struct {

	// 步骤结果值
	Result *string `json:"result,omitempty"`

	// 步骤实际结果
	ActualResult *string `json:"actual_result,omitempty"`

	// 步骤期望结果
	ExpectResult *string `json:"expect_result,omitempty"`

	// 用例操作步骤
	TestStep *string `json:"test_step,omitempty"`
}

func (o TestCaseStepResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseStepResultInfo struct{}"
	}

	return strings.Join([]string{"TestCaseStepResultInfo", string(data)}, " ")
}
