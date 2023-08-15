package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugRuleResponse Response Object
type DebugRuleResponse struct {

	// 规则测试输出结果
	TestResult     *string `json:"test_result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DebugRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugRuleResponse struct{}"
	}

	return strings.Join([]string{"DebugRuleResponse", string(data)}, " ")
}
