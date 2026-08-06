package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTaskCaseQuery struct {

	// 用例创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 用例结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	KeyWord *string `json:"keyWord,omitempty"`

	// 执行机区域ID
	LocationId *string `json:"location_id,omitempty"`

	More *bool `json:"more,omitempty"`

	// 分页时页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页时每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 告警策略选择失败后重试时有值
	Pid *string `json:"pid,omitempty"`

	// cloudTest任务执行结果列表
	Results *[]int32 `json:"results,omitempty"`

	// 排序字段
	SortBy *string `json:"sortBy,omitempty"`

	// 用例所处的阶段 0：前置， 1：测试用例 2：后置用例
	Stage *int32 `json:"stage,omitempty"`

	// 用例开始时间
	StartTime *int64 `json:"startTime,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 子任务ID列表
	SubtaskIds *[]string `json:"subtask_ids,omitempty"`

	// 子任务ID
	SubtaskId *string `json:"subtask_id,omitempty"`

	SuiteType *int32 `json:"suiteType,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务ID列表
	TaskIds *[]string `json:"task_ids,omitempty"`

	// 任务类型
	TaskTypeId *string `json:"taskTypeId,omitempty"`

	// 服务ID
	TestServiceId *string `json:"test_service_id,omitempty"`

	// 用例ID
	TestcaseId *string `json:"testcase_id,omitempty"`
}

func (o SubTaskCaseQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskCaseQuery struct{}"
	}

	return strings.Join([]string{"SubTaskCaseQuery", string(data)}, " ")
}
