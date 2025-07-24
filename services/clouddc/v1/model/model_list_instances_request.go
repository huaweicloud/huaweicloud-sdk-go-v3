package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// 分页查询时每页行数。最大值为 1000。  默认值： 当不设置值或设置的值小于 10 时，默认值为 10。 当设置的值大于 1000 时，默认值为 1000。
	Limit *string `json:"limit,omitempty"`

	// 分页游标
	Marker *string `json:"marker,omitempty"`

	// 实例状态
	InstanceState *ListInstancesRequestInstanceState `json:"instance_state,omitempty"`

	// 实例 ID。取值可以由多个实例 ID 组成数组，最多支持 100 个 ID，ID 之间用半角逗号（,）隔开，示例：uuid1,uuid2,uuid3。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例 ID。取值可以由多个服务器 ID 组成数组，最多支持 100 个 ID，ID 之间用半角逗号（,）隔开，示例：uuid1,uuid2,uuid3。与instance_id_set查询条件互斥。
	ServerId *string `json:"server_id,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestInstanceState struct {
	value string
}

type ListInstancesRequestInstanceStateEnum struct {
	PENDING       ListInstancesRequestInstanceState
	RUNNING       ListInstancesRequestInstanceState
	STOPPING      ListInstancesRequestInstanceState
	STOPPED       ListInstancesRequestInstanceState
	REINSTALLING  ListInstancesRequestInstanceState
	SHUTTING_DOWN ListInstancesRequestInstanceState
	TERMINATED    ListInstancesRequestInstanceState
	FAILED        ListInstancesRequestInstanceState
}

func GetListInstancesRequestInstanceStateEnum() ListInstancesRequestInstanceStateEnum {
	return ListInstancesRequestInstanceStateEnum{
		PENDING: ListInstancesRequestInstanceState{
			value: "pending",
		},
		RUNNING: ListInstancesRequestInstanceState{
			value: "running",
		},
		STOPPING: ListInstancesRequestInstanceState{
			value: "stopping",
		},
		STOPPED: ListInstancesRequestInstanceState{
			value: "stopped",
		},
		REINSTALLING: ListInstancesRequestInstanceState{
			value: "reinstalling",
		},
		SHUTTING_DOWN: ListInstancesRequestInstanceState{
			value: "shutting-down",
		},
		TERMINATED: ListInstancesRequestInstanceState{
			value: "terminated",
		},
		FAILED: ListInstancesRequestInstanceState{
			value: "failed",
		},
	}
}

func (c ListInstancesRequestInstanceState) Value() string {
	return c.value
}

func (c ListInstancesRequestInstanceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestInstanceState) UnmarshalJSON(b []byte) error {
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
