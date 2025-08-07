package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 索引位置，偏移量。  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为50。
	Limit *int32 `json:"limit,omitempty"`

	// 任务ID，按任务ID过滤。
	Id *string `json:"id,omitempty"`

	// 实例ID，按实例ID过滤。
	InstanceId *string `json:"instance_id,omitempty"`

	// 订单ID，按订单ID过滤。
	OrderId *string `json:"order_id,omitempty"`

	// 任务名称，按任务名称称过滤。
	Name *string `json:"name,omitempty"`

	// 任务状态，按任务状态过滤。 Running:运行中 Completed:已完成 Failed:已失败
	Status *ListTasksRequestStatus `json:"status,omitempty"`

	// 任务的创建时间，按时间范围进行过滤，“start_time”有值时，“end_time”必选。格式为UTC时间戳。
	StartTime *string `json:"start_time,omitempty"`

	// 任务的结束时间，按时间范围进行过滤，“start_time”有值时，“end_time”必选。格式为UTC时间戳。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestStatus struct {
	value string
}

type ListTasksRequestStatusEnum struct {
	RUNNING   ListTasksRequestStatus
	COMPLETED ListTasksRequestStatus
	FAILED    ListTasksRequestStatus
}

func GetListTasksRequestStatusEnum() ListTasksRequestStatusEnum {
	return ListTasksRequestStatusEnum{
		RUNNING: ListTasksRequestStatus{
			value: "Running",
		},
		COMPLETED: ListTasksRequestStatus{
			value: "Completed",
		},
		FAILED: ListTasksRequestStatus{
			value: "Failed",
		},
	}
}

func (c ListTasksRequestStatus) Value() string {
	return c.value
}

func (c ListTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
