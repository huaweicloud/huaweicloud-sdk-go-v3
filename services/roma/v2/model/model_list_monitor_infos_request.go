package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListMonitorInfosRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于1

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 需要搜索的任务名称，支持模糊搜索，大小写敏感，非必填参数，如果为空，搜索所有任务

	TaskName *string `json:"task_name,omitempty"`
	// 需要搜索任务的执行状态, 只允许如下枚举值：UNSTARTED-未启动, WAITING-等待执行,RUNNING-执行中, SUCCESS-执行成功, CANCELLED-任务取消, ERROR-执行异常</br> 非必填参数，如果为空，搜索所有任务

	ExecuteStatus *ListMonitorInfosRequestExecuteStatus `json:"execute_status,omitempty"`
}

func (o ListMonitorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListMonitorInfosRequest", string(data)}, " ")
}

type ListMonitorInfosRequestExecuteStatus struct {
	value string
}

type ListMonitorInfosRequestExecuteStatusEnum struct {
	UNSTARTED ListMonitorInfosRequestExecuteStatus
	WAITING   ListMonitorInfosRequestExecuteStatus
	RUNNING   ListMonitorInfosRequestExecuteStatus
	SUCCESS   ListMonitorInfosRequestExecuteStatus
	CANCELLED ListMonitorInfosRequestExecuteStatus
	ERROR     ListMonitorInfosRequestExecuteStatus
}

func GetListMonitorInfosRequestExecuteStatusEnum() ListMonitorInfosRequestExecuteStatusEnum {
	return ListMonitorInfosRequestExecuteStatusEnum{
		UNSTARTED: ListMonitorInfosRequestExecuteStatus{
			value: "UNSTARTED",
		},
		WAITING: ListMonitorInfosRequestExecuteStatus{
			value: "WAITING",
		},
		RUNNING: ListMonitorInfosRequestExecuteStatus{
			value: "RUNNING",
		},
		SUCCESS: ListMonitorInfosRequestExecuteStatus{
			value: "SUCCESS",
		},
		CANCELLED: ListMonitorInfosRequestExecuteStatus{
			value: "CANCELLED",
		},
		ERROR: ListMonitorInfosRequestExecuteStatus{
			value: "ERROR",
		},
	}
}

func (c ListMonitorInfosRequestExecuteStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMonitorInfosRequestExecuteStatus) UnmarshalJSON(b []byte) error {
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
