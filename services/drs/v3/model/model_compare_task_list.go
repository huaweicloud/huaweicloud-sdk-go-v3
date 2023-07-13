package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CompareTaskList struct {

	// 对比任务的id。
	CompareTaskId string `json:"compare_task_id"`

	// 对比任务的类型。
	CompareType string `json:"compare_type"`

	// 对比任务的状态。 - RUNNING-运行中 - WAITING_FOR_RUNNING-等待启动中 - SUCCESSFUL-完成 - FAILED-失败 - CANCELLED-已取消 - TIMEOUT_INTERRUPT-超时中断 - FULL_DOING-全量校验中 - INCRE_DOING-增量校验中
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
			value: "RUNNING",
		},
		WAITING_FOR_RUNNING: CompareTaskListCompareTaskStatus{
			value: "WAITING_FOR_RUNNING",
		},
		SUCCESSFUL: CompareTaskListCompareTaskStatus{
			value: "SUCCESSFUL",
		},
		FAILED: CompareTaskListCompareTaskStatus{
			value: "FAILED",
		},
		CANCELLED: CompareTaskListCompareTaskStatus{
			value: "CANCELLED",
		},
		TIMEOUT_INTERRUPT: CompareTaskListCompareTaskStatus{
			value: "TIMEOUT_INTERRUPT",
		},
		FULL_DOING: CompareTaskListCompareTaskStatus{
			value: "FULL_DOING",
		},
		INCRE_DOING: CompareTaskListCompareTaskStatus{
			value: "INCRE_DOING",
		},
	}
}

func (c CompareTaskListCompareTaskStatus) Value() string {
	return c.value
}

func (c CompareTaskListCompareTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompareTaskListCompareTaskStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
