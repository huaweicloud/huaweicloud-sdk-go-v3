package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTaskQueryByPageParams struct {
	More *bool `json:"more,omitempty"`

	// 未设置发布版本
	NotSetReleaseDev *bool `json:"notSetReleaseDev,omitempty"`

	// 页码
	PageNumber *int32 `json:"page_number,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 父任务id
	ParentSubTaskId *string `json:"parent_sub_task_id,omitempty"`

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty"`

	// -| 发布的版本，空数组：代表所有未设置的； null或者无此字段，搜索所有版本 有内容：搜索所有版本
	ReleaseDev *[]string `json:"release_dev,omitempty"`

	Results *[]int32 `json:"results,omitempty"`

	// 任务执行第一次时间
	StartTimeBegin *int64 `json:"start_time_begin,omitempty"`

	// 任务执行最后一次时间
	StartTimeEnd *int64 `json:"start_time_end,omitempty"`

	// 状态
	State *int32 `json:"state,omitempty"`

	// 子任务状态列表
	States *[]int32 `json:"states,omitempty"`

	// 子任务任务id
	SubTaskId *string `json:"sub_task_id,omitempty"`

	// 测试套类型
	SuiteType *int32 `json:"suite_type,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型，1=拨测，2=冒烟
	TaskTypeId *string `json:"task_type_id,omitempty"`

	// 项目id
	TestServiceId *string `json:"test_service_id,omitempty"`
}

func (o SubTaskQueryByPageParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskQueryByPageParams struct{}"
	}

	return strings.Join([]string{"SubTaskQueryByPageParams", string(data)}, " ")
}
