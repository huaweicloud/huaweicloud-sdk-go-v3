package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUpgradeTasksRequest Request Object
type ListUpgradeTasksRequest struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称（支持模糊查询）
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型：0-云桌面 1-应用服务器 2-镜像
	TaskType *int32 `json:"task_type,omitempty"`

	// 执行周期类型：FIXED_TIME-指定时间 DAY-按天 WEEK-按周 MONTH-按月
	ScheduledType *ListUpgradeTasksRequestScheduledType `json:"scheduled_type,omitempty"`

	// 启用状态：0-未启用 1-启用
	IsEnable *int32 `json:"is_enable,omitempty"`

	// 上次执行状态
	LastExecuteStatus *string `json:"last_execute_status,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，默认10，最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUpgradeTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpgradeTasksRequest struct{}"
	}

	return strings.Join([]string{"ListUpgradeTasksRequest", string(data)}, " ")
}

type ListUpgradeTasksRequestScheduledType struct {
	value string
}

type ListUpgradeTasksRequestScheduledTypeEnum struct {
	FIXED_TIME ListUpgradeTasksRequestScheduledType
	DAY        ListUpgradeTasksRequestScheduledType
	WEEK       ListUpgradeTasksRequestScheduledType
	MONTH      ListUpgradeTasksRequestScheduledType
}

func GetListUpgradeTasksRequestScheduledTypeEnum() ListUpgradeTasksRequestScheduledTypeEnum {
	return ListUpgradeTasksRequestScheduledTypeEnum{
		FIXED_TIME: ListUpgradeTasksRequestScheduledType{
			value: "FIXED_TIME",
		},
		DAY: ListUpgradeTasksRequestScheduledType{
			value: "DAY",
		},
		WEEK: ListUpgradeTasksRequestScheduledType{
			value: "WEEK",
		},
		MONTH: ListUpgradeTasksRequestScheduledType{
			value: "MONTH",
		},
	}
}

func (c ListUpgradeTasksRequestScheduledType) Value() string {
	return c.value
}

func (c ListUpgradeTasksRequestScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUpgradeTasksRequestScheduledType) UnmarshalJSON(b []byte) error {
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
