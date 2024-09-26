package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseDetailResponse Response Object
type ShowTestCaseDetailResponse struct {

	// 测试用例唯一标识
	TestcaseId *string `json:"testcase_id,omitempty"`

	// 软开云项目唯一标识
	ProjectId *string `json:"project_id,omitempty"`

	// 注册测试类型服务接口返回的服务id
	ServiceId *int32 `json:"service_id,omitempty"`

	// 测试用例名称
	Name *string `json:"name,omitempty"`

	// 测试用例编号
	TestcaseNumber *string `json:"testcase_number,omitempty"`

	// 测试用例等级
	RankId *int32 `json:"rank_id,omitempty"`

	// 测试用例状态
	StatusId *string `json:"status_id,omitempty"`

	AssignedUser *AssignedUserInfo `json:"assigned_user,omitempty"`

	// 测试用例执行次数
	ExecuteCount *int32 `json:"execute_count,omitempty"`

	// 测试用例执行结果
	ResultId *string `json:"result_id,omitempty"`

	ExtendInfo *ExtendInfo `json:"extend_info,omitempty"`

	// 接口调用失败错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 接口调用失败错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTestCaseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTestCaseDetailResponse", string(data)}, " ")
}
