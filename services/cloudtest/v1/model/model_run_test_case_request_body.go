package model

import (
	"encoding/json"

	"strings"
)

// 执行测试用例请求体
type RunTestCaseRequestBody struct {
	// 注册结果返回的服务id

	ServiceId int32 `json:"service_id"`
	// 测试计划id

	PlanId *string `json:"plan_id,omitempty"`
	// 测试用例执行信息，数组长度小于等于50

	ExecuteList []TestCaseExecuteBean `json:"execute_list"`
}

func (o RunTestCaseRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunTestCaseRequestBody struct{}"
	}

	return strings.Join([]string{"RunTestCaseRequestBody", string(data)}, " ")
}
