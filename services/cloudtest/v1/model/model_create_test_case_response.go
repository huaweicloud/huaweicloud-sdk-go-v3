package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTestCaseResponse struct {

	// 测试用例唯一标识
	TestcaseId *string `json:"testcase_id,omitempty" xml:"testcase_id"`

	// 软开云项目唯一标识
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 注册测试类型服务接口返回的服务id
	ServiceId *int32 `json:"service_id,omitempty" xml:"service_id"`

	// 测试用例名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 测试用例编号
	TestcaseNumber *string `json:"testcase_number,omitempty" xml:"testcase_number"`

	// 测试用例等级
	RankId *string `json:"rank_id,omitempty" xml:"rank_id"`

	// 测试用例状态
	StatusId *string `json:"status_id,omitempty" xml:"status_id"`

	AssignedUser *AssignedUserInfo `json:"assigned_user,omitempty" xml:"assigned_user"`

	// 测试用例执行次数
	ExecuteCount *int32 `json:"execute_count,omitempty" xml:"execute_count"`

	// 测试用例执行结果
	ResultId *string `json:"result_id,omitempty" xml:"result_id"`

	ExtendInfo *ExtendInfo `json:"extend_info,omitempty" xml:"extend_info"`

	// 接口调用失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 接口调用失败错误信息
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTestCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseResponse struct{}"
	}

	return strings.Join([]string{"CreateTestCaseResponse", string(data)}, " ")
}
