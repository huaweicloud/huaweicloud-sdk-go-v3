package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskCaseResponseTimeDetailVo struct {

	// 用例ID
	CaseId *string `json:"case_id,omitempty"`

	// 用例名称
	CaseName *string `json:"case_name,omitempty"`

	// 用例结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 用例状态
	ErrorReason *string `json:"error_reason,omitempty"`

	// 用例响应时间
	ResponseTime *int64 `json:"response_time,omitempty"`

	// 服务ID
	ServiceId *string `json:"service_id,omitempty"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 用例开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 用例状态
	State *string `json:"state,omitempty"`

	// 子任务ID
	SubTaskId *string `json:"sub_task_id,omitempty"`

	// 已废弃
	TagId *string `json:"tag_id,omitempty"`

	// 已废弃
	TagName *string `json:"tag_name,omitempty"`

	// 任务和用例关联关系的ID
	TaskCaseId *string `json:"task_case_id,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型，1=拨测，2=冒烟，4=部署测试，5=小网拨测
	TaskTypeId *string `json:"task_type_id,omitempty"`

	// 测试套类型
	TestSuiteType *int32 `json:"test_suite_type,omitempty"`
}

func (o TaskCaseResponseTimeDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCaseResponseTimeDetailVo struct{}"
	}

	return strings.Join([]string{"TaskCaseResponseTimeDetailVo", string(data)}, " ")
}
