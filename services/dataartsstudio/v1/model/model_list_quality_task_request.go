package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListQualityTaskRequest Request Object
type ListQualityTaskRequest struct {

	// 目录ID
	CategoryId *int64 `json:"category_id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 调度状态 UNKNOWN:未知,NOT_START:未启动,SCHEDULING:调度中,FINISH_SUCCESS:正常结束,KILL:手动停止,RUNNING_EXCEPTION:运行失败
	ScheduleStatus *ListQualityTaskRequestScheduleStatus `json:"schedule_status,omitempty"`

	// 最近运行时间查询区间的开始时间,13位时间戳(精确到毫秒)
	StartTime *int64 `json:"start_time,omitempty"`

	// 最近运行时间查询区间的结束时间,13位时间戳(精确到毫秒)
	EndTime *int64 `json:"end_time,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 分页条数,最大值为100
	Limit *int64 `json:"limit,omitempty"`

	// 分页偏移量,最小值0
	Offset *int64 `json:"offset,omitempty"`

	// workspace 信息
	Workspace string `json:"workspace"`
}

func (o ListQualityTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTaskRequest struct{}"
	}

	return strings.Join([]string{"ListQualityTaskRequest", string(data)}, " ")
}

type ListQualityTaskRequestScheduleStatus struct {
	value string
}

type ListQualityTaskRequestScheduleStatusEnum struct {
	UNKNOWN           ListQualityTaskRequestScheduleStatus
	NOT_START         ListQualityTaskRequestScheduleStatus
	SCHEDULING        ListQualityTaskRequestScheduleStatus
	FINISH_SUCCESS    ListQualityTaskRequestScheduleStatus
	KILL              ListQualityTaskRequestScheduleStatus
	RUNNING_EXCEPTION ListQualityTaskRequestScheduleStatus
}

func GetListQualityTaskRequestScheduleStatusEnum() ListQualityTaskRequestScheduleStatusEnum {
	return ListQualityTaskRequestScheduleStatusEnum{
		UNKNOWN: ListQualityTaskRequestScheduleStatus{
			value: "UNKNOWN",
		},
		NOT_START: ListQualityTaskRequestScheduleStatus{
			value: "NOT_START",
		},
		SCHEDULING: ListQualityTaskRequestScheduleStatus{
			value: "SCHEDULING",
		},
		FINISH_SUCCESS: ListQualityTaskRequestScheduleStatus{
			value: "FINISH_SUCCESS",
		},
		KILL: ListQualityTaskRequestScheduleStatus{
			value: "KILL",
		},
		RUNNING_EXCEPTION: ListQualityTaskRequestScheduleStatus{
			value: "RUNNING_EXCEPTION",
		},
	}
}

func (c ListQualityTaskRequestScheduleStatus) Value() string {
	return c.value
}

func (c ListQualityTaskRequestScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQualityTaskRequestScheduleStatus) UnmarshalJSON(b []byte) error {
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
