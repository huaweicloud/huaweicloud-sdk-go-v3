package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DebugRuleRequestBody struct {
	// 测试的被规则执行的数据

	TestData *string `json:"test_data,omitempty"`
	// 测试的规则

	TestRuleExpress *string `json:"test_rule_express,omitempty"`
}

func (o DebugRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugRuleRequestBody struct{}"
	}

	return strings.Join([]string{"DebugRuleRequestBody", string(data)}, " ")
}
