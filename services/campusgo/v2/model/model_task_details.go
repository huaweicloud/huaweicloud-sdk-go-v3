package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type TaskDetails struct {
	// 作业ID

	Id string `json:"id"`
	// 作业的名称

	Name string `json:"name"`
	// 作业创建者的用户名

	Creator string `json:"creator"`
	// 作业创建者的项目ID

	ProjectId string `json:"project_id"`
	// 作业的描述

	Description *string `json:"description,omitempty"`
	// 作业对应服务的名称

	ServiceName string `json:"service_name"`
	// 作业对应服务的版本号

	ServiceVersion string `json:"service_version"`

	ServiceTitle *TaskDetailsServiceTitle `json:"service_title"`
	// 仅边缘作业会出现，作业运行所在的边缘运行池ID

	EdgePoolId *string `json:"edge_pool_id,omitempty"`
	// 作业指定的算法能力包包周期订单ID

	ResourceOrderId *string `json:"resource_order_id,omitempty"`
	// 作业创建的时间

	CreatedAt *sdktime.SdkTime `json:"created_at"`
	// 作业最近一次状态更新的时间

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
	// 作业当前的状态，分别为PENDING（等待中），RECOVERING（恢复中），STARTING（启动中），UPGRADING（升级中），CREATE_FAILED（创建失败），START_FAILED（启动失败），RUNNING（运行中），STOPPING（停止中），STOPPED（已停止），ABNORMAL（异常），SUCCEEDED（运行成功），FAILED（运行失败），DELETING（删除中），FREEZING（冻结中），FROZEN（已冻结）

	State TaskDetailsState `json:"state"`
	// 作业状态的详情信息，仅部分状态会有详情信息

	Status *string `json:"status,omitempty"`

	Error *TaskDetailsError `json:"error,omitempty"`

	Input *TaskInput `json:"input"`

	Output *TaskOutputForDisplay `json:"output"`

	ServiceConfig *TaskServiceConfig `json:"service_config,omitempty"`
}

func (o TaskDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetails struct{}"
	}

	return strings.Join([]string{"TaskDetails", string(data)}, " ")
}

type TaskDetailsState struct {
	value string
}

type TaskDetailsStateEnum struct {
	PENDING       TaskDetailsState
	RECOVERING    TaskDetailsState
	STARTING      TaskDetailsState
	UPGRADING     TaskDetailsState
	CREATE_FAILED TaskDetailsState
	START_FAILED  TaskDetailsState
	RUNNING       TaskDetailsState
	STOPPING      TaskDetailsState
	STOPPED       TaskDetailsState
	ABNORMAL      TaskDetailsState
	SUCCEEDED     TaskDetailsState
	FAILED        TaskDetailsState
	DELETING      TaskDetailsState
	FREEZING      TaskDetailsState
	FROZEN        TaskDetailsState
}

func GetTaskDetailsStateEnum() TaskDetailsStateEnum {
	return TaskDetailsStateEnum{
		PENDING: TaskDetailsState{
			value: "PENDING",
		},
		RECOVERING: TaskDetailsState{
			value: "RECOVERING",
		},
		STARTING: TaskDetailsState{
			value: "STARTING",
		},
		UPGRADING: TaskDetailsState{
			value: "UPGRADING",
		},
		CREATE_FAILED: TaskDetailsState{
			value: "CREATE_FAILED",
		},
		START_FAILED: TaskDetailsState{
			value: "START_FAILED",
		},
		RUNNING: TaskDetailsState{
			value: "RUNNING",
		},
		STOPPING: TaskDetailsState{
			value: "STOPPING",
		},
		STOPPED: TaskDetailsState{
			value: "STOPPED",
		},
		ABNORMAL: TaskDetailsState{
			value: "ABNORMAL",
		},
		SUCCEEDED: TaskDetailsState{
			value: "SUCCEEDED",
		},
		FAILED: TaskDetailsState{
			value: "FAILED",
		},
		DELETING: TaskDetailsState{
			value: "DELETING",
		},
		FREEZING: TaskDetailsState{
			value: "FREEZING",
		},
		FROZEN: TaskDetailsState{
			value: "FROZEN",
		},
	}
}

func (c TaskDetailsState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskDetailsState) UnmarshalJSON(b []byte) error {
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
