package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetryFactoryJobInstanceBody struct {

	// 重跑开始位置，取值有firstNode、errorNode和specifiedNode。
	RetryLocation string `json:"retry_location"`

	// 节点名称。
	NodeName *string `json:"node_name,omitempty"`

	// 使用的作业参数，取值有original_version和current_version。
	RetryTaskVersion string `json:"retry_task_version"`

	// 是否忽略obs监听，默认为true。
	IgnoreObsMonitor *bool `json:"ignore_obs_monitor,omitempty"`

	// 作业实例重跑参数，当重跑当前实例类型时，需要指定该参数的重跑信息，重跑当前作业及其上下游作业实例类型不需要指定该参数。
	TaskRetrys *[]RetryFactoryJobInstanceBodyTaskRetrys `json:"task_retrys,omitempty"`

	// 实例开始时间，当重跑当前作业及其上下游作业实例类型时，该参数有效。
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 实例结束时间，当重跑当前作业及其上下游作业实例类型时，该参数有效。
	EndTime *int64 `json:"end_time,omitempty"`

	// 作业实例重跑参数，当重跑当前作业及其上下游作业实例类型时，需要指定该参数的重跑信息，重跑当前实例类型不需要指定该参数。
	Jobs *[]RetryFactoryJobInstanceBodyJobs `json:"jobs,omitempty"`

	// 并行度，当重跑当前作业及其上下游作业实例类型时，该参数有效，默认值为1，取值范围为1到128。
	Concurrent *int32 `json:"concurrent,omitempty"`
}

func (o RetryFactoryJobInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryFactoryJobInstanceBody struct{}"
	}

	return strings.Join([]string{"RetryFactoryJobInstanceBody", string(data)}, " ")
}
