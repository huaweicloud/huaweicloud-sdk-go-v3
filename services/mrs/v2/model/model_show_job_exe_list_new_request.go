package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowJobExeListNewRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。

	ClusterId string `json:"cluster_id"`
	// 作业名称，只能由字母、数字、中划线和下划线组成，并且长度为1～36个字符。

	JobName *string `json:"job_name,omitempty"`
	// 作业ID，只能由字母、数字、中划线(-)组成，并且长度为1~36字符。

	JobId *string `json:"job_id,omitempty"`
	// 用户名称、只能由字母、数字、特殊字符(-_.)组成，且不能以数字开头，并且长度为1～32字符。

	User *string `json:"user,omitempty"`
	// 作业类型。 - MapReduce - SparkSubmit - SparkSubmit：SparkPython类型的作业在查询时作业类型请选择SparkSubmit。 - HiveScript - HiveSql - DistCp，导入、导出数据。 - SparkScript - SparkSql - Flink

	JobType *string `json:"job_type,omitempty"`
	// 作业运行状态。 - FAILED：失败 - KILLED：已终止 - NEW：已创建 - NEW_SAVING：已创建保存中 - SUBMITTED：已提交 - ACCEPTED：已接受 - RUNNING：运行中 - FINISHED：已完成

	JobState *ShowJobExeListNewRequestJobState `json:"job_state,omitempty"`
	// 作业运行结果。 - FAILED：执行失败的作业。 - KILLED：执行中被手动终止的作业。 - UNDEFINED：正在执行的作业。 - SUCCEEDED：执行成功的作业。

	JobResult *ShowJobExeListNewRequestJobResult `json:"job_result,omitempty"`
	// 作业的资源对列类型名称，作业的资源对列类型名称，只能由数字、字母和特殊字符(-_)组成, 并且长度为1～64字符。

	Queue *string `json:"queue,omitempty"`
	// 返回结果中每页显示条数。缺省值：10

	Limit *string `json:"limit,omitempty"`
	// 表示作业列表从该偏移量开始查询。缺省值：1

	Offset *string `json:"offset,omitempty"`
	// 返回结果的排序方式，默认值为desc。 - asc：按升序排列 - desc：按降序排列

	SortBy *string `json:"sort_by,omitempty"`
	// 查询该时间之后提交的作业，UTC的毫秒时间戳。例如：1562032041362。

	SubmittedTimeBegin *int64 `json:"submitted_time_begin,omitempty"`
	// 查询该时间之前提交的作业UTC的毫秒时间戳。例如：1562032041362。

	SubmittedTimeEnd *int64 `json:"submitted_time_end,omitempty"`
}

func (o ShowJobExeListNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExeListNewRequest struct{}"
	}

	return strings.Join([]string{"ShowJobExeListNewRequest", string(data)}, " ")
}

type ShowJobExeListNewRequestJobState struct {
	value string
}

type ShowJobExeListNewRequestJobStateEnum struct {
	FAILED     ShowJobExeListNewRequestJobState
	KILLED     ShowJobExeListNewRequestJobState
	NEW        ShowJobExeListNewRequestJobState
	NEW_SAVING ShowJobExeListNewRequestJobState
	SUBMITTED  ShowJobExeListNewRequestJobState
	ACCEPTED   ShowJobExeListNewRequestJobState
	RUNNING    ShowJobExeListNewRequestJobState
	FINISHED   ShowJobExeListNewRequestJobState
}

func GetShowJobExeListNewRequestJobStateEnum() ShowJobExeListNewRequestJobStateEnum {
	return ShowJobExeListNewRequestJobStateEnum{
		FAILED: ShowJobExeListNewRequestJobState{
			value: "FAILED：失败",
		},
		KILLED: ShowJobExeListNewRequestJobState{
			value: "KILLED：已终止",
		},
		NEW: ShowJobExeListNewRequestJobState{
			value: "NEW：已创建",
		},
		NEW_SAVING: ShowJobExeListNewRequestJobState{
			value: "NEW_SAVING：已创建保存中",
		},
		SUBMITTED: ShowJobExeListNewRequestJobState{
			value: "SUBMITTED：已提交",
		},
		ACCEPTED: ShowJobExeListNewRequestJobState{
			value: "ACCEPTED：已接受",
		},
		RUNNING: ShowJobExeListNewRequestJobState{
			value: "RUNNING：运行中",
		},
		FINISHED: ShowJobExeListNewRequestJobState{
			value: "FINISHED：已完成",
		},
	}
}

func (c ShowJobExeListNewRequestJobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobExeListNewRequestJobState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowJobExeListNewRequestJobResult struct {
	value string
}

type ShowJobExeListNewRequestJobResultEnum struct {
	FAILED    ShowJobExeListNewRequestJobResult
	KILLED    ShowJobExeListNewRequestJobResult
	UNDEFINED ShowJobExeListNewRequestJobResult
	SUCCEEDED ShowJobExeListNewRequestJobResult
}

func GetShowJobExeListNewRequestJobResultEnum() ShowJobExeListNewRequestJobResultEnum {
	return ShowJobExeListNewRequestJobResultEnum{
		FAILED: ShowJobExeListNewRequestJobResult{
			value: "FAILED：执行失败的作业。",
		},
		KILLED: ShowJobExeListNewRequestJobResult{
			value: "KILLED：执行中被手动终止的作业。",
		},
		UNDEFINED: ShowJobExeListNewRequestJobResult{
			value: "UNDEFINED：正在执行的作业。",
		},
		SUCCEEDED: ShowJobExeListNewRequestJobResult{
			value: "SUCCEEDED：执行成功的作业。",
		},
	}
}

func (c ShowJobExeListNewRequestJobResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobExeListNewRequestJobResult) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
