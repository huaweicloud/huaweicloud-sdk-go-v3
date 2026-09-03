package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCaseTemplateVo struct {

	// 用例对应的告警模板id
	AlertTemplateId *string `json:"alertTemplateId,omitempty"`

	// 测试用例id
	TestcaseId *string `json:"testcase_id,omitempty"`
}

func (o TestCaseTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseTemplateVo struct{}"
	}

	return strings.Join([]string{"TestCaseTemplateVo", string(data)}, " ")
}
