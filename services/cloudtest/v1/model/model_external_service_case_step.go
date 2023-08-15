package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalServiceCaseStep 测试步骤，数组长度小于10
type ExternalServiceCaseStep struct {

	// 测试用例预期结果信息，长度为[0-500]位字符
	ExpectResult *string `json:"expect_result,omitempty"`

	// 测试步骤描述信息，长度为[0-500]位字符
	TestStep *string `json:"test_step,omitempty"`
}

func (o ExternalServiceCaseStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalServiceCaseStep struct{}"
	}

	return strings.Join([]string{"ExternalServiceCaseStep", string(data)}, " ")
}
