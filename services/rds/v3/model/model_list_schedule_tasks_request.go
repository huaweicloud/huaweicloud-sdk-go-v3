package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScheduleTasksRequest Request Object
type ListScheduleTasksRequest struct {

	// 索引位置，偏移量。  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为50。
	Limit *int32 `json:"limit,omitempty"`

	// 实例id，按实例id过滤。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称，按实例名称过滤。
	InstanceName *string `json:"instance_name,omitempty"`

	// 任务状态，按任务状态过滤。 Initing:初始化中 Pending:挂起 Running:运行中 Completed:已完成 Failed:已失败 Unauthorized:未授权 Canceled:已取消
	Status *ListScheduleTasksRequestStatus `json:"status,omitempty"`

	// 查询任务创建的开始时间，“start_time”有值时，“end_time”必选。格式为UTC时间戳。
	StartTime *string `json:"start_time,omitempty"`

	// 查询任务创建的结束时间，“start_time”有值时，“end_time”必选。格式为UTC时间戳。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListScheduleTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleTasksRequest", string(data)}, " ")
}

type ListScheduleTasksRequestStatus struct {
	value string
}

type ListScheduleTasksRequestStatusEnum struct {
	INITING      ListScheduleTasksRequestStatus
	PENDING      ListScheduleTasksRequestStatus
	RUNNING      ListScheduleTasksRequestStatus
	COMPLETED    ListScheduleTasksRequestStatus
	FAILED       ListScheduleTasksRequestStatus
	UNAUTHORIZED ListScheduleTasksRequestStatus
	CANCELED     ListScheduleTasksRequestStatus
}

func GetListScheduleTasksRequestStatusEnum() ListScheduleTasksRequestStatusEnum {
	return ListScheduleTasksRequestStatusEnum{
		INITING: ListScheduleTasksRequestStatus{
			value: "Initing",
		},
		PENDING: ListScheduleTasksRequestStatus{
			value: "Pending",
		},
		RUNNING: ListScheduleTasksRequestStatus{
			value: "Running",
		},
		COMPLETED: ListScheduleTasksRequestStatus{
			value: "Completed",
		},
		FAILED: ListScheduleTasksRequestStatus{
			value: "Failed",
		},
		UNAUTHORIZED: ListScheduleTasksRequestStatus{
			value: "Unauthorized",
		},
		CANCELED: ListScheduleTasksRequestStatus{
			value: "Canceled",
		},
	}
}

func (c ListScheduleTasksRequestStatus) Value() string {
	return c.value
}

func (c ListScheduleTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduleTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
