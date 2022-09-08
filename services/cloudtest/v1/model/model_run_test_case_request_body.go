package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行测试用例请求体
type RunTestCaseRequestBody struct {

	// 注册测试类型服务接口返回的服务id
	ServiceId int32 `json:"service_id"`

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty"`

	// 测试用例执行信息，数组长度小于等于50
	ExecuteList []TestCaseExecuteBean `json:"execute_list"`
}

func (o RunTestCaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTestCaseRequestBody struct{}"
	}

	return strings.Join([]string{"RunTestCaseRequestBody", string(data)}, " ")
}
