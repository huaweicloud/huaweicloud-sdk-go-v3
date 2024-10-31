package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultStepVo 测试步骤信息
type ResultStepVo struct {

	// 结果
	Result *string `json:"result,omitempty"`

	// 测试步骤
	TestStep *string `json:"test_step,omitempty"`

	// 预期结果
	ExpectResult *string `json:"expect_result,omitempty"`

	// 实际结果
	ActualResult *string `json:"actual_result,omitempty"`

	// 测试结果名称
	ResultName *string `json:"result_name,omitempty"`
}

func (o ResultStepVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultStepVo struct{}"
	}

	return strings.Join([]string{"ResultStepVo", string(data)}, " ")
}
