package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type CompareTaskList struct {
	// 对比任务的id。

	CompareTaskId string `json:"compare_task_id"`
	// 对比任务的类型。

	CompareType string `json:"compare_type"`
	// 对比任务的状态。

	CompareTaskStatus CompareTaskListCompareTaskStatus `json:"compare_task_status"`
	// 对比开始时间。

	CreateTime string `json:"create_time"`
	// 对比结束时间。

	EndTime *string `json:"end_time,omitempty"`
}

func (o CompareTaskList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskList struct{}"
	}

	return strings.Join([]string{"CompareTaskList", string(data)}, " ")
}

type CompareTaskListCompareTaskStatus struct {
	value string
}

type CompareTaskListCompareTaskStatusEnum struct {
	RUNNING             CompareTaskListCompareTaskStatus
	WAITING_FOR_RUNNING CompareTaskListCompareTaskStatus
	SUCCESSFUL          CompareTaskListCompareTaskStatus
	FAILED              CompareTaskListCompareTaskStatus
	CANCELLED           CompareTaskListCompareTaskStatus
	TIMEOUT_INTERRUPT   CompareTaskListCompareTaskStatus
	FULL_DOING          CompareTaskListCompareTaskStatus
	INCRE_DOING         CompareTaskListCompareTaskStatus
}

func GetCompareTaskListCompareTaskStatusEnum() CompareTaskListCompareTaskStatusEnum {
	return CompareTaskListCompareTaskStatusEnum{
		RUNNING: CompareTaskListCompareTaskStatus{
			value: "RUNNING-运行中",
		},
		WAITING_FOR_RUNNING: CompareTaskListCompareTaskStatus{
			value: "WAITING_FOR_RUNNING-等待启动中",
		},
		SUCCESSFUL: CompareTaskListCompareTaskStatus{
			value: "SUCCESSFUL-完成",
		},
		FAILED: CompareTaskListCompareTaskStatus{
			value: "FAILED-失败",
		},
		CANCELLED: CompareTaskListCompareTaskStatus{
			value: "CANCELLED-已取消",
		},
		TIMEOUT_INTERRUPT: CompareTaskListCompareTaskStatus{
			value: "TIMEOUT_INTERRUPT-超时中断",
		},
		FULL_DOING: CompareTaskListCompareTaskStatus{
			value: "FULL_DOING-全量校验中",
		},
		INCRE_DOING: CompareTaskListCompareTaskStatus{
			value: "INCRE_DOING-增量校验中",
		},
	}
}

func (c CompareTaskListCompareTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompareTaskListCompareTaskStatus) UnmarshalJSON(b []byte) error {
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
