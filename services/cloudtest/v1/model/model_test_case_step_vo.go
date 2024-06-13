package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseStepVo 用例测试步骤和预期结果
type TestCaseStepVo struct {

	// 测试步骤
	TestStep *string `json:"test_step,omitempty"`

	// 预期结果
	ExpectResult *string `json:"expect_result,omitempty"`
}

func (o TestCaseStepVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseStepVo struct{}"
	}

	return strings.Join([]string{"TestCaseStepVo", string(data)}, " ")
}
