package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseStepInfo 对外测试步骤
type TestCaseStepInfo struct {

	// 测试步骤
	TestStep *string `json:"test_step,omitempty"`

	// 预期结果
	ExpectResult *string `json:"expect_result,omitempty"`

	// 步骤的实际结果
	StepActual *string `json:"step_actual,omitempty"`

	// 步骤结果
	StepResult *string `json:"step_result,omitempty"`
}

func (o TestCaseStepInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseStepInfo struct{}"
	}

	return strings.Join([]string{"TestCaseStepInfo", string(data)}, " ")
}
