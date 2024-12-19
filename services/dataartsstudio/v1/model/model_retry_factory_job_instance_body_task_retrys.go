package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetryFactoryJobInstanceBodyTaskRetrys struct {

	// 作业id，当前重跑实例的作业id。
	JobId int64 `json:"job_id"`

	// 实例计划时间，13位时间戳，可通过查询作业实例列表接口获取。
	PlanTime int64 `json:"plan_time"`

	// 实例提交时间，13位时间戳，可通过查询作业实例列表接口获取。
	SubmitTime int64 `json:"submit_time"`
}

func (o RetryFactoryJobInstanceBodyTaskRetrys) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryFactoryJobInstanceBodyTaskRetrys struct{}"
	}

	return strings.Join([]string{"RetryFactoryJobInstanceBodyTaskRetrys", string(data)}, " ")
}
