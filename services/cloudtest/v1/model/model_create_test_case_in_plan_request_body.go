package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计划中批量添加测试用例请求体
type CreateTestCaseInPlanRequestBody struct {
	// 注册测试类型服务接口返回的服务id

	ServiceId int32 `json:"service_id"`
	// 计划下包含的用例个数，数组长度小于50个，只能包含一种测试类型

	TestcaseIdList []string `json:"testcase_id_list"`
}

func (o CreateTestCaseInPlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseInPlanRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTestCaseInPlanRequestBody", string(data)}, " ")
}
